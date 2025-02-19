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

package subscription_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/assetholder"
	"github.com/perun-network/perun-eth-backend/bindings/assetholdereth"
	"github.com/perun-network/perun-eth-backend/bindings/peruntoken"
	ethchannel "github.com/perun-network/perun-eth-backend/channel"
	"github.com/perun-network/perun-eth-backend/channel/test"
	"github.com/perun-network/perun-eth-backend/subscription"
	"github.com/perun-network/perun-eth-backend/wallet/keystore"

	channeltest "perun.network/go-perun/channel/test"
	"perun.network/go-perun/log"
	wallettest "perun.network/go-perun/wallet/test"
	pkgtest "polycry.pt/poly-go/test"
)

const (
	txGasLimit      = 100000
	txFinalityDepth = 1
)

// TestEventSub tests the `EventSub` by:
// 1. Emit `1/4 n` events
// 2. Starting up the `EventSub`
// 3. Emit `3/4 n` events
// 4. Checking that `EventSub` contains `n` distinct events.
func TestEventSub(t *testing.T) {
	n := 1000
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	rng := pkgtest.Prng(t)

	// Simulated chain setup.
	sb := test.NewSimulatedBackend()
	ksWallet := wallettest.RandomWallet(1).(*keystore.Wallet)
	account := &ksWallet.NewRandomAccount(rng).(*keystore.Account).Account
	sb.FundAddress(ctx, account.Address)
	cb := ethchannel.NewContractBackend(
		sb,
		ethchannel.MakeChainID(sb.ChainID()),
		keystore.NewTransactor(*ksWallet, sb.Signer),
		txFinalityDepth,
	)

	// Setup Perun Token.
	tokenAddr, err := ethchannel.DeployPerunToken(ctx, cb, *account, []common.Address{account.Address}, channeltest.MaxBalance)
	require.NoError(t, err)
	token, err := peruntoken.NewPeruntoken(tokenAddr, cb)
	log.Println("Token Address: ", tokenAddr, err)
	require.NoError(t, err)
	ct := pkgtest.NewConcurrent(t)

	// Sync channel to ensure that at least n/4 events were sent.
	waitSent := make(chan interface{})
	go ct.Stage("emitter", func(t pkgtest.ConcT) {
		for i := 0; i < n; i++ {
			if i == n/4 {
				close(waitSent)
			}
			log.Println("Sending ", i)
			// Send the transaction.
			opts, err := cb.NewTransactor(ctx, txGasLimit, *account)
			log.Println("Transactor: ", opts, err)
			require.NoError(t, err)
			tx, err := token.IncreaseAllowance(opts, account.Address, big.NewInt(1))
			log.Println("TX: ", tx, err)
			require.NoError(t, err)
			// Wait for the TX to be mined.
			_, err = cb.ConfirmTransaction(ctx, tx, *account)
			log.Println("Confirm TX: ", err)
			require.NoError(t, err)
		}
	})
	sink := make(chan *subscription.Event, 10)
	eFact := func() *subscription.Event {
		return &subscription.Event{
			Name: bindings.Events.PerunTokenApproval,
			Data: new(peruntoken.PeruntokenApproval),
		}
	}
	// Setup the event sub after some events have been sent.
	<-waitSent
	contract := bind.NewBoundContract(tokenAddr, bindings.ABI.PerunToken, cb, cb, cb)
	log.Println("Contract: ", contract)
	sub, err := subscription.NewEventSub(ctx, cb, contract, eFact, 10000)
	log.Println("Sub")
	require.NoError(t, err)
	go ct.Stage("sub", func(t pkgtest.ConcT) {
		defer close(sink)
		require.NoError(t, sub.Read(context.Background(), sink))
	})

	go ct.Stage("receiver", func(t pkgtest.ConcT) {
		var lastTx common.Hash
		// Receive `n` unique events.
		for i := 0; i < n; i++ {
			e := <-sink
			log.Println("Read ", i)
			require.NotNil(t, e)
			// It is possible to receive the same event twice.
			if e.Log.TxHash == lastTx {
				i--
				log.Println("Duplicate ", i)
			}
			lastTx = e.Log.TxHash
			log.Println("TX Hash: ", lastTx)
			want := &peruntoken.PeruntokenApproval{
				Owner:   account.Address,
				Spender: account.Address,
				Value:   big.NewInt(int64(i + 1)),
			}
			log.Println("Want: ", want)
			require.Equal(t, want, e.Data)
			require.False(t, e.Log.Removed)
		}
		sub.Close()
	})

	ct.Wait("emitter", "sub", "receiver")
	// Check that read terminated.
	require.Nil(t, <-sink)
}

// TestEventSub_Filter checks that the EventSub filters transactions.
func TestEventSub_Filter(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	rng := pkgtest.Prng(t)

	// Simulated chain setup.
	sb := test.NewSimulatedBackend()
	ksWallet := wallettest.RandomWallet(1).(*keystore.Wallet)
	account := &ksWallet.NewRandomAccount(rng).(*keystore.Account).Account
	sb.FundAddress(ctx, account.Address)
	cb := ethchannel.NewContractBackend(
		sb,
		ethchannel.MakeChainID(sb.ChainID()),
		keystore.NewTransactor(*ksWallet, sb.Signer),
		txFinalityDepth,
	)

	// Setup ETH AssetHolder.
	adjAddr, err := ethchannel.DeployAdjudicator(ctx, cb, *account)
	require.NoError(t, err)
	ahAddr, err := ethchannel.DeployETHAssetholder(ctx, cb, adjAddr, *account)
	require.NoError(t, err)
	ah, err := assetholdereth.NewAssetholdereth(ahAddr, cb)
	require.NoError(t, err)
	ct := pkgtest.NewConcurrent(t)

	// Send the transaction.
	fundingID := channeltest.NewRandomChannelID(rng, channeltest.WithBackend(1))
	opts, err := cb.NewTransactor(ctx, txGasLimit, *account)
	require.NoError(t, err)
	opts.Value = big.NewInt(1)
	tx, err := ah.Deposit(opts, fundingID, big.NewInt(1))
	require.NoError(t, err)
	// Wait for the TX to be mined.
	_, err = cb.ConfirmTransaction(ctx, tx, *account)
	require.NoError(t, err)

	// Create the filter.
	Filter := []interface{}{fundingID}
	// Setup the event sub.
	sink := make(chan *subscription.Event, 1)
	eFact := func() *subscription.Event {
		return &subscription.Event{
			Name:   bindings.Events.AhDeposited,
			Data:   new(assetholder.AssetholderDeposited),
			Filter: [][]interface{}{Filter},
		}
	}
	contract := bind.NewBoundContract(ahAddr, bindings.ABI.AssetHolder, cb, cb, cb)
	sub, err := subscription.NewEventSub(ctx, cb, contract, eFact, 100)
	require.NoError(t, err)
	go ct.Stage("sub", func(t pkgtest.ConcT) {
		defer close(sink)
		require.NoError(t, sub.Read(context.Background(), sink))
	})

	// Receive 1 event.
	e := <-sink
	require.NotNil(t, e)
	want := &assetholder.AssetholderDeposited{
		FundingID: fundingID,
		Amount:    big.NewInt(int64(1)),
	}
	log.Debug("TX0 Hash: ", e.Log.TxHash)
	require.Equal(t, want, e.Data)
	require.False(t, e.Log.Removed)
	sub.Close()
	// We do not check here that <-sink returns nil, since the EventSub
	// can receive events more than once.
}
