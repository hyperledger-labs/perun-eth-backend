// Copyright 2024 - See NOTICE file for copyright holders.
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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/assetholder"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
	"github.com/perun-network/perun-eth-backend/subscription"
	"github.com/perun-network/perun-eth-backend/wallet"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/log"
)

// Withdraw ensures that a channel has been concluded and the final outcome.
// withdrawn from the asset holders.
func (a *Adjudicator) Withdraw(ctx context.Context, req channel.AdjudicatorReq, subStates channel.StateMap) error {
	if err := a.ensureConcluded(ctx, req, subStates); err != nil {
		return errors.WithMessage(err, "ensure Concluded")
	}
	if err := a.checkConcludedState(ctx, req, subStates); err != nil {
		return errors.WithMessage(err, "check concluded state")
	}
	return errors.WithMessage(a.ensureWithdrawn(ctx, req), "ensure Withdrawn")
}

// ensureWithdrawn ensures that the channel has been withdrawn from the asset.
func (a *Adjudicator) ensureWithdrawn(ctx context.Context, req channel.AdjudicatorReq) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, asset := range filterAssets(req.Tx.Allocation.Assets, a.chainID) {
		index, ok := assetIdx(req.Tx.Allocation.Assets, asset)
		if !ok {
			return errors.New("asset not found in adjudicator request")
		}
		// Skip zero balance withdrawals
		if req.Tx.Allocation.Balances[index][req.Idx].Sign() == 0 {
			a.log.WithFields(log.Fields{"channel": req.Params.ID, "idx": req.Idx}).Debug("Skipped zero withdrawing.")
			continue
		}
		asset := asset // Capture asset locally for usage in closure.
		g.Go(func() error {
			// Create subscription
			contract := bindAssetHolder(a.ContractBackend, asset, index)
			fundingID := FundingIDs(req.Params.ID(), req.Params.Parts[req.Idx])[0]
			events := make(chan *subscription.Event, adjEventBuffSize)
			subErr := make(chan error, 1)
			sub, err := subscription.Subscribe(ctx, a.ContractBackend, contract.contract, withdrawnEventType(fundingID), startBlockOffset, a.txFinalityDepth)
			if err != nil {
				return errors.WithMessage(err, "subscribing")
			}
			defer sub.Close()

			// Check for past event.
			if err := sub.ReadPast(ctx, events); err != nil {
				return errors.WithMessage(err, "reading past events")
			}
			select {
			case <-events:
				return nil
			default:
			}
			log.Println("No withdrawn event found in the past, send transaction.", req.Tx.Sigs)
			// No withdrawn event found in the past, send transaction.
			if err := a.callAssetWithdraw(ctx, req, contract); err != nil {
				return errors.WithMessage(err, "withdrawing assets failed")
			}

			// Wait for event.
			go func() {
				subErr <- sub.Read(ctx, events)
			}()

			select {
			case <-events:
				return nil
			case <-ctx.Done():
				return errors.Wrap(ctx.Err(), "context cancelled")
			case err = <-subErr:
				if err != nil {
					return errors.WithMessage(err, "subscription error")
				}
				return errors.New("subscription closed")
			}
		})
	}
	return g.Wait()
}

func withdrawnEventType(fundingID [32]byte) subscription.EventFactory {
	return func() *subscription.Event {
		return &subscription.Event{
			Name:   bindings.Events.AhWithdrawn,
			Data:   new(assetholder.AssetholderWithdrawn),
			Filter: [][]interface{}{{fundingID}},
		}
	}
}

func bindAssetHolder(cb ContractBackend, asset channel.Asset, assetIndex channel.Index) assetHolder {
	// Decode and set the asset address.
	assetAddr := asset.(*Asset).EthAddress() //nolint:forcetypeassert
	ctr, err := assetholder.NewAssetholder(assetAddr, cb)
	if err != nil {
		log.Panic("Invalid AssetHolder ABI definition.")
	}
	contract := bind.NewBoundContract(assetAddr, bindings.ABI.AssetHolder, cb, cb, cb)
	return assetHolder{ctr, &assetAddr, contract, assetIndex}
}

func (a *Adjudicator) callAssetWithdraw(ctx context.Context, request channel.AdjudicatorReq, asset assetHolder) error {
	auth, sig, err := a.newWithdrawalAuth(request, asset)
	if err != nil {
		return errors.WithMessage(err, "creating withdrawal auth")
	}
	tx, err := func() (*types.Transaction, error) {
		if !a.mu.TryLockCtx(ctx) {
			return nil, errors.Wrap(ctx.Err(), "context canceled while acquiring tx lock")
		}
		defer a.mu.Unlock()
		trans, err := a.NewTransactor(ctx, a.gasLimit, a.txSender)
		if err != nil {
			return nil, errors.WithMessagef(err, "creating transactor for asset %d", asset.assetIndex)
		}
		tx, err := asset.Withdraw(trans, auth, sig)
		if err != nil {
			err = cherrors.CheckIsChainNotReachableError(err)
			return nil, errors.WithMessagef(err, "withdrawing asset %d with transaction nonce %d", asset.assetIndex, trans.Nonce)
		}
		return tx, nil
	}()
	if err != nil {
		return err
	}
	_, err = a.ConfirmTransaction(ctx, tx, a.txSender)
	if err != nil && errors.Is(err, errTxTimedOut) {
		err = client.NewTxTimedoutError(Withdraw.String(), tx.Hash().Hex(), err.Error())
	}
	return errors.WithMessage(err, "mining transaction")
}

func (a *Adjudicator) newWithdrawalAuth(request channel.AdjudicatorReq, asset assetHolder) (assetholder.AssetHolderWithdrawalAuth, []byte, error) {
	fid := FundingID(request.Tx.ID, request.Params.Parts[request.Idx][1])
	bal, err := asset.Assetholder.Holdings(nil, fid)
	if err != nil {
		return assetholder.AssetHolderWithdrawalAuth{}, nil, fmt.Errorf("getting balance: %w", err)
	}

	auth := assetholder.AssetHolderWithdrawalAuth{
		ChannelID:   request.Params.ID(),
		Participant: wallet.AsChannelParticipant(wallet.AddressMapfromAccountMap(request.Acc)),
		Receiver:    a.Receiver,
		Amount:      bal,
	}
	enc, err := encodeAssetHolderWithdrawalAuth(auth)
	if err != nil {
		return assetholder.AssetHolderWithdrawalAuth{}, nil, errors.WithMessage(err, "encoding withdrawal auth")
	}

	sig, err := request.Acc[1].SignData(enc)
	return auth, sig, errors.WithMessage(err, "sign data")
}

func encodeAssetHolderWithdrawalAuth(a assetholder.AssetHolderWithdrawalAuth) ([]byte, error) {
	// Define the top-level ABI type for the Authorization struct.
	authorizationType, err := abi.NewType("tuple", "tuple(bytes32 channelID, tuple(address ethAddress, bytes ccAddress) participant, address receiver, uint256 amount)", []abi.ArgumentMarshaling{
		{Name: "channelID", Type: "bytes32"},
		{Name: "participant", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "ethAddress", Type: "address"},
			{Name: "ccAddress", Type: "bytes"},
		}},
		{Name: "receiver", Type: "address"},
		{Name: "amount", Type: "uint256"},
	})
	if err != nil {
		return nil, err
	}

	// Define the Arguments.
	args := abi.Arguments{
		{Type: authorizationType},
	}

	// Pack the data for encoding.
	return args.Pack(
		struct {
			ChannelID   [32]byte
			Participant struct {
				EthAddress common.Address
				CcAddress  []byte
			}
			Receiver common.Address
			Amount   *big.Int
		}{
			ChannelID: a.ChannelID,
			Participant: struct {
				EthAddress common.Address
				CcAddress  []byte
			}{
				EthAddress: a.Participant.EthAddress,
				CcAddress:  a.Participant.CcAddress,
			},
			Receiver: a.Receiver,
			Amount:   a.Amount,
		},
	)
}
