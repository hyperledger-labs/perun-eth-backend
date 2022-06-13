// Copyright 2021 - See NOTICE file for copyright holders.
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

package test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/require"

	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/peruntoken"
	ethchannel "github.com/perun-network/perun-eth-backend/channel"
	"github.com/perun-network/perun-eth-backend/wallet/keystore"

	channeltest "perun.network/go-perun/channel/test"
	wallettest "perun.network/go-perun/wallet/test"
)

// TokenSetup is used to create specific Events and TX easily.
type TokenSetup struct {
	SB         *SimulatedBackend
	CB         ethchannel.ContractBackend
	Token      *peruntoken.Peruntoken
	Contract   *bind.BoundContract
	R          *require.Assertions
	T          *testing.T
	Acc1, Acc2 *accounts.Account

	subApproval, subTransfer event.Subscription
	SinkApproval             chan *peruntoken.PeruntokenApproval
	SinkTransfer             chan *peruntoken.PeruntokenTransfer
}

const (
	eventTimeout  = 100 * time.Millisecond
	txGasLimit    = 100000
	blockInterval = 100 * time.Millisecond
	eventBuffSize = 10
)

// NewTokenSetup creates a new TokenSetup.
func NewTokenSetup(ctx context.Context, t *testing.T, rng *rand.Rand, txFinalityDepth uint64) *TokenSetup {
	t.Helper()
	// Simulated chain setup.
	sb := NewSimulatedBackend()
	ksWallet := wallettest.RandomWallet().(*keystore.Wallet)
	acc1 := &ksWallet.NewRandomAccount(rng).(*keystore.Account).Account
	sb.FundAddress(ctx, acc1.Address)
	acc2 := &ksWallet.NewRandomAccount(rng).(*keystore.Account).Account
	sb.FundAddress(ctx, acc2.Address)
	cb := ethchannel.NewContractBackend(
		sb,
		keystore.NewTransactor(*ksWallet, SimSigner),
		txFinalityDepth,
	)

	// Setup Perun Token.
	sb.StartMining(blockInterval) // Automine to allow reorg-resistant contract deployment.
	defer sb.StopMining()
	tokenAddr, err := ethchannel.DeployPerunToken(ctx, cb, *acc1, []common.Address{acc1.Address}, channeltest.MaxBalance)
	require.NoError(t, err)
	token, err := peruntoken.NewPeruntoken(tokenAddr, cb)
	require.NoError(t, err)
	contract := bind.NewBoundContract(tokenAddr, bindings.ABI.PerunToken, cb, cb, cb)

	return &TokenSetup{
		SB:       sb,
		CB:       cb,
		Token:    token,
		Contract: contract,
		R:        require.New(t),
		T:        t,
		Acc1:     acc1,
		Acc2:     acc2,
	}
}

// StartSubs starts the Approval and Transfer subscriptions.
func (s *TokenSetup) StartSubs() {
	// Approval sub.
	sinkApproval := make(chan *peruntoken.PeruntokenApproval, eventBuffSize)
	subApproval, err := s.Token.WatchApproval(&bind.WatchOpts{}, sinkApproval, nil, nil)
	require.NoError(s.T, err)
	// Transfer sub.
	sinkTransfer := make(chan *peruntoken.PeruntokenTransfer, eventBuffSize)
	subTransfer, err := s.Token.WatchTransfer(&bind.WatchOpts{}, sinkTransfer, nil, nil)
	require.NoError(s.T, err)

	s.subApproval = subApproval
	s.SinkApproval = sinkApproval
	s.subTransfer = subTransfer
	s.SinkTransfer = sinkTransfer
}

// IncAllowance sends an IncreaseAllowance TX.
func (s *TokenSetup) IncAllowance(ctx context.Context) *types.Transaction {
	opts, err := s.CB.NewTransactor(ctx, txGasLimit, *s.Acc1)
	s.R.NoError(err)
	tx, err := s.Token.IncreaseAllowance(opts, s.Acc2.Address, big.NewInt(1))
	s.R.NoError(err)
	return tx
}

// Transfer sends a Transfer TX.
func (s *TokenSetup) Transfer(ctx context.Context) *types.Transaction {
	opts, err := s.CB.NewTransactor(ctx, txGasLimit, *s.Acc2)
	s.R.NoError(err)
	tx, err := s.Token.TransferFrom(opts, s.Acc1.Address, s.Acc2.Address, big.NewInt(1))
	s.R.NoError(err)
	return tx
}

// ConfirmTx confirms that a TX is included in the chain at least once.
func (s *TokenSetup) ConfirmTx(tx *types.Transaction, confirm bool) {
	ctx, cancel := context.WithTimeout(context.Background(), eventTimeout)
	defer cancel()
	_, err := s.CB.ConfirmTransaction(ctx, tx, *s.Acc1)

	if confirm {
		s.R.NoError(err)
	} else {
		s.R.Error(err)
	}
}

// AllowanceEvent waits for an allowance event with value `v`.
// `included` decided whether or not its `Removed` values should not be set.
func (s *TokenSetup) AllowanceEvent(v uint64, included bool) {
	var e *peruntoken.PeruntokenApproval

	select {
	case e = <-s.SinkApproval:
	case <-time.After(eventTimeout):
		s.T.FailNow()
	}

	s.R.NotNil(e)
	s.R.Equal(s.Acc1.Address, e.Owner)
	s.R.Equal(s.Acc2.Address, e.Spender)
	s.R.Equal(v, e.Value.Uint64())
	s.R.Equal(!included, e.Raw.Removed)
}

// TransferEvent waits for a transfer event.
// `included` decided whether or not its `Removed` values should not be set.
func (s *TokenSetup) TransferEvent(included bool) {
	var e *peruntoken.PeruntokenTransfer

	select {
	case e = <-s.SinkTransfer:
	case <-time.After(eventTimeout):
		s.T.FailNow()
	}

	s.R.NotNil(e)
	s.R.Equal(s.Acc1.Address, e.From)
	s.R.Equal(s.Acc2.Address, e.To)
	s.R.Equal(big.NewInt(1), e.Value)
	s.R.Equal(!included, e.Raw.Removed)
}

// NoMoreEvents asserts that no more events should be generated.
func (s *TokenSetup) NoMoreEvents() {
	select {
	case e := <-s.SinkApproval:
		s.R.FailNow("Expected no event but got: ", e)
	case e := <-s.SinkTransfer:
		s.R.FailNow("Expected no event but got: ", e)
	case <-time.After(eventTimeout):
	}
}

// StopSubs stops the event subs.
// Should be called for cleanup iff StartSubs was called.
func (s *TokenSetup) StopSubs() {
	s.subApproval.Unsubscribe()
	s.subTransfer.Unsubscribe()
	close(s.SinkApproval)
	close(s.SinkTransfer)
}
