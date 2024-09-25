// Copyright 2020 - See NOTICE file for copyright holders.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package channel

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/log"
	psync "polycry.pt/poly-go/sync"
)

// Compile time check that we implement the Perun adjudicator interface.
var _ channel.Adjudicator = (*Adjudicator)(nil)

// The Adjudicator struct implements the channel.Adjudicator interface
// It provides all functionality to close a channel.
type Adjudicator struct {
	ContractBackend
	// chainID specifies the chain the funder is living on.
	chainID  AssetID
	contract *adjudicator.Adjudicator
	bound    *bind.BoundContract
	// The address to which we send all funds.
	Receiver common.Address
	// Structured logger
	log log.Logger
	// Transaction mutex
	mu psync.Mutex
	// txSender is sending the TX.
	txSender accounts.Account
	// gasLimit is the gas limit for all transactions to the Adjudicator.
	gasLimit uint64
}

// NewAdjudicator creates a new ethereum adjudicator. The receiver is the
// on-chain address that receives withdrawals.
func NewAdjudicator(backend ContractBackend, contract common.Address, receiver common.Address, txSender accounts.Account, gasLimit uint64) *Adjudicator {
	contr, err := adjudicator.NewAdjudicator(contract, backend)
	if err != nil {
		panic("Could not create a new instance of adjudicator")
	}
	bound := bind.NewBoundContract(contract, bindings.ABI.Adjudicator, backend, backend, backend)
	return &Adjudicator{
		ContractBackend: backend,
		chainID:         backend.chainID,
		contract:        contr,
		bound:           bound,
		Receiver:        receiver,
		txSender:        txSender,
		log:             log.WithField("txSender", txSender.Address),
		gasLimit:        gasLimit,
	}
}

// Progress progresses a channel state on-chain.
func (a *Adjudicator) Progress(ctx context.Context, req channel.ProgressReq) error {
	ethNewState := ToEthState(req.NewState)
	ethActorIndex := big.NewInt(int64(req.Idx))

	progress := func(
		opts *bind.TransactOpts,
		params adjudicator.ChannelParams,
		state adjudicator.ChannelState,
		_ [][]byte,
	) (*types.Transaction, error) {
		return a.contract.Progress(opts, params, state, ethNewState, ethActorIndex, req.Sig)
	}
	return a.call(ctx, req.AdjudicatorReq, progress, Progress)
}

func (a *Adjudicator) callRegister(ctx context.Context, req channel.AdjudicatorReq, subChannels []channel.SignedState) error {
	return a.call(ctx, req,
		func(opts *bind.TransactOpts, params adjudicator.ChannelParams, state adjudicator.ChannelState, sigs [][]byte) (*types.Transaction, error) {
			ch := adjudicator.AdjudicatorSignedState{
				Params: params,
				State:  state,
				Sigs:   sigs,
			}
			sub := toEthSignedStates(subChannels)
			return a.contract.Register(opts, ch, sub)
		}, Register)
}

func toEthSignedStates(subChannels []channel.SignedState) (ethSubChannels []adjudicator.AdjudicatorSignedState) {
	ethSubChannels = make([]adjudicator.AdjudicatorSignedState, len(subChannels))
	for i, x := range subChannels {
		ethSubChannels[i] = adjudicator.AdjudicatorSignedState{
			Params: ToEthParams(x.Params),
			State:  ToEthState(x.State),
			Sigs:   x.Sigs,
		}
		log.Println("Subchannel", ethSubChannels[i].State.ChannelID, i)
	}
	return
}

func (a *Adjudicator) callConclude(ctx context.Context, req channel.AdjudicatorReq, subStates channel.StateMap) error {
	ethSubStates := toEthSubStates(req.Tx.State, subStates)
	log.Println("Concluding channel", req.Params.ID, ethSubStates)

	conclude := func(
		opts *bind.TransactOpts,
		params adjudicator.ChannelParams,
		state adjudicator.ChannelState,
		_ [][]byte,
	) (*types.Transaction, error) {
		return a.contract.Conclude(opts, params, state, ethSubStates)
	}

	return a.call(ctx, req, conclude, Conclude)
}

func (a *Adjudicator) callConcludeFinal(ctx context.Context, req channel.AdjudicatorReq) error {
	return a.call(ctx, req, a.contract.ConcludeFinal, ConcludeFinal)
}

type adjFunc = func(
	opts *bind.TransactOpts,
	params adjudicator.ChannelParams,
	state adjudicator.ChannelState,
	sigs [][]byte,
) (*types.Transaction, error)

// call calls the given contract function `fn` with the data from `req`.
// `fn` should be a method of `a.contract`, like `a.contract.Register`.
// `txType` should be one of the valid transaction types defined in the client package.
func (a *Adjudicator) call(ctx context.Context, req channel.AdjudicatorReq, fn adjFunc, txType OnChainTxType) error {
	ethParams := ToEthParams(req.Params)
	ethState := ToEthState(req.Tx.State)
	tx, err := func() (*types.Transaction, error) {
		if !a.mu.TryLockCtx(ctx) {
			return nil, errors.Wrap(ctx.Err(), "context canceled while acquiring tx lock")
		}
		defer a.mu.Unlock()

		trans, err := a.NewTransactor(ctx, a.gasLimit, a.txSender)
		if err != nil {
			return nil, errors.WithMessage(err, "creating transactor")
		}
		tx, err := fn(trans, ethParams, ethState, req.Tx.Sigs)
		if err != nil {
			err = cherrors.CheckIsChainNotReachableError(err)
			return nil, errors.WithMessage(err, "calling adjudicator function")
		}
		log.Debugf("Sent transaction %v", tx.Hash().Hex())
		return tx, nil
	}()
	if err != nil {
		return err
	}

	_, err = a.ConfirmTransaction(ctx, tx, a.txSender)
	log.Println("Transaction confirmed", err)
	if errors.Is(err, errTxTimedOut) {
		err = client.NewTxTimedoutError(txType.String(), tx.Hash().Hex(), err.Error())
	}
	return errors.WithMessage(err, "mining transaction")
}

// ValidateAdjudicator checks if the bytecode at given address is correct.
// Returns a ContractBytecodeError if the bytecode at given address is invalid.
// This error can be checked with function IsErrInvalidContractCode.
func ValidateAdjudicator(ctx context.Context, backend bind.ContractCaller, adjudicatorAddr common.Address) error {
	return validateContract(ctx, backend, adjudicatorAddr, adjudicator.AdjudicatorBinRuntime)
}

// toEthSubStates generates a channel tree in depth-first order.
func toEthSubStates(state *channel.State, subStates channel.StateMap) (ethSubStates []adjudicator.ChannelState) {
	for _, subAlloc := range state.Locked {
		subState, ok := subStates[channel.IDKey(subAlloc.ID)]
		if !ok {
			log.Panic("sub-state not found")
		}
		ethSubStates = append(ethSubStates, ToEthState(subState))
		if len(subState.Locked) > 0 {
			_subSubStates := toEthSubStates(subState, subStates)
			ethSubStates = append(ethSubStates, _subSubStates...)
		}
	}
	return
}
