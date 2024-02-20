// Copyright 2019 - See NOTICE file for copyright holders.
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

package channel_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ethchannel "github.com/perun-network/perun-eth-backend/channel"
	"github.com/perun-network/perun-eth-backend/channel/test"
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"
	"perun.network/go-perun/channel"
	channeltest "perun.network/go-perun/channel/test"
	pkgtest "polycry.pt/poly-go/test"
)

func TestAdjudicator_MultipleWithdraws_FinalState(t *testing.T) {
	testParallel := func(n int) {
		t.Run(fmt.Sprintf("Withdraw %d party parallel", n), func(t *testing.T) { withdrawMultipleConcurrentFinal(t, n, true) })
	}
	testSequential := func(n int) {
		t.Run(fmt.Sprintf("Withdraw %d party sequential", n), func(t *testing.T) { withdrawMultipleConcurrentFinal(t, n, false) })
	}

	for _, n := range []int{1, 2, 5} {
		testParallel(n)
		testSequential(n)
	}
}

func withdrawMultipleConcurrentFinal(t *testing.T, numParts int, parallel bool) {
	t.Helper()
	rng := pkgtest.Prng(t)
	// create test setup
	s := test.NewSetup(t, rng, numParts, blockInterval, TxFinalityDepth)
	// create valid state and params
	params, state := channeltest.NewRandomParamsAndState(
		rng,
		channeltest.WithParts(s.Parts...),
		channeltest.WithAssets(s.Asset),
		channeltest.WithIsFinal(false),
		channeltest.WithLedgerChannel(true),
	)
	// we need to properly fund the channel
	fundingCtx, funCancel := context.WithTimeout(context.Background(), defaultTxTimeout*time.Duration(numParts))
	defer funCancel()
	// fund the contract
	ct := pkgtest.NewConcurrent(t)
	for i, funder := range s.Funders {
		sleepTime := time.Millisecond * time.Duration(rng.Int63n(10)+1)
		i, funder := i, funder
		go ct.StageN("funding loop", numParts, func(rt pkgtest.ConcT) {
			time.Sleep(sleepTime)
			req := channel.NewFundingReq(params, state, channel.Index(i), state.Balances)
			require.NoError(rt, funder.Fund(fundingCtx, *req), "funding should succeed")
		})
	}
	ct.Wait("funding loop")
	// manipulate the state
	state.IsFinal = true
	tx := testSignState(t, s.Accs, state)

	// Now test the withdraw function
	ctx, cancel := context.WithTimeout(context.Background(), defaultTxTimeout)
	defer cancel()
	if parallel {
		startBarrier := make(chan struct{})
		var wg sync.WaitGroup
		wg.Add(numParts)
		for i := 0; i < numParts; i++ {
			sleepDuration := time.Duration(rng.Int63n(10)+1) * time.Millisecond
			go func(i int) {
				defer wg.Done()
				<-startBarrier
				time.Sleep(sleepDuration)
				req := channel.AdjudicatorReq{
					Params: params,
					Acc:    s.Accs[i],
					Idx:    channel.Index(i),
					Tx:     tx,
				}
				err := s.Adjs[i].Withdraw(ctx, req, nil)
				assert.NoError(t, err, "Withdrawing should succeed")
			}(i)
		}
		close(startBarrier)
		wg.Wait()
	} else {
		for i := 0; i < numParts; i++ {
			req := channel.AdjudicatorReq{
				Params: params,
				Acc:    s.Accs[i],
				Idx:    channel.Index(i),
				Tx:     tx,
			}
			err := s.Adjs[i].Withdraw(ctx, req, nil)
			assert.NoError(t, err, "Withdrawing should succeed")
		}
	}
	assertHoldingsZero(ctx, t, s.CB, params, state.Assets)
}

func TestWithdrawZeroBalance(t *testing.T) {
	t.Run("1 Participant", func(t *testing.T) { testWithdrawZeroBalance(t, 1) })
	t.Run("2 Participant", func(t *testing.T) { testWithdrawZeroBalance(t, 2) })
	t.Run("5 Participant", func(t *testing.T) { testWithdrawZeroBalance(t, 5) })
}

// shouldFunders decides who should fund. 1 indicates funding, 0 indicates skipping.
//nolint:thelper // Not a helper.
func testWithdrawZeroBalance(t *testing.T, n int) {
	rng := pkgtest.Prng(t)
	s := test.NewSetup(t, rng, n, blockInterval, TxFinalityDepth)
	// create valid state and params
	params, state := channeltest.NewRandomParamsAndState(
		rng,
		channeltest.WithParts(s.Parts...),
		channeltest.WithAssets(s.Asset),
		channeltest.WithIsFinal(true),
		channeltest.WithLedgerChannel(true),
	)
	agreement := state.Balances.Clone()

	for i := range params.Parts {
		if i%2 == 0 {
			state.Balances[0][i].SetInt64(0)
			agreement[0][i].SetInt64(0)
		} // is != 0 otherwise
		t.Logf("Part: %d ShouldFund: %t Bal: %v", i, i%2 == 1, state.Balances[0][i])
	}

	// fund
	ct := pkgtest.NewConcurrent(t)
	for i, funder := range s.Funders {
		i, funder := i, funder
		go ct.StageN("funding loop", n, func(rt pkgtest.ConcT) {
			req := channel.NewFundingReq(params, state, channel.Index(i), agreement)
			require.NoError(rt, funder.Fund(context.Background(), *req), "funding should succeed")
		})
	}
	ct.Wait("funding loop")

	// register
	req := channel.AdjudicatorReq{
		Params: params,
		Acc:    s.Accs[0],
		Tx:     testSignState(t, s.Accs, state),
		Idx:    0,
	}
	require.NoError(t, s.Adjs[0].Register(context.Background(), req, nil))
	// we don't need to wait for a timeout since we registered a final state

	// withdraw
	for i, _adj := range s.Adjs {
		adj := _adj
		req.Acc = s.Accs[i]
		req.Idx = channel.Index(i)
		// check that the nonce stays the same for zero balance withdrawals
		diff, err := test.NonceDiff(s.Accs[i].Address(), adj, func() error {
			return adj.Withdraw(context.Background(), req, nil)
		})
		require.NoError(t, err)
		if i%2 == 0 {
			assert.Zero(t, diff, "Nonce should stay the same")
		} else {
			assert.Equal(t, 1, diff, "Nonce should increase by 1")
		}
	}
	assertHoldingsZero(context.Background(), t, s.CB, params, state.Assets)
}

func TestWithdraw(t *testing.T) {
	rng := pkgtest.Prng(t)
	// create test setup
	s := test.NewSetup(t, rng, 1, blockInterval, TxFinalityDepth)
	// create valid state and params
	params, state := channeltest.NewRandomParamsAndState(
		rng,
		channeltest.WithParts(s.Parts...),
		channeltest.WithAssets(s.Asset),
		channeltest.WithIsFinal(false),
		channeltest.WithLedgerChannel(true),
	)
	// we need to properly fund the channel
	fundingCtx, funCancel := context.WithTimeout(context.Background(), defaultTxTimeout)
	defer funCancel()
	// fund the contract
	fundingReq := channel.NewFundingReq(params, state, channel.Index(0), state.Balances)
	require.NoError(t, s.Funders[0].Fund(fundingCtx, *fundingReq), "funding should succeed")
	req := channel.AdjudicatorReq{
		Params: params,
		Acc:    s.Accs[0],
		Idx:    channel.Index(0),
	}

	testWithdraw := func(t *testing.T, shouldWork bool) {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), defaultTxTimeout)
		defer cancel()
		req.Tx = testSignState(t, s.Accs, state)
		err := s.Adjs[0].Withdraw(ctx, req, nil)

		if shouldWork {
			assert.NoError(t, err, "Withdrawing should work")
		} else {
			assert.Error(t, err, "Withdrawing should fail")
		}
	}

	t.Run("Withdraw non-final state", func(t *testing.T) {
		testWithdraw(t, false)
	})

	t.Run("Withdraw final state", func(t *testing.T) {
		state.IsFinal = true
		testWithdraw(t, true)
	})

	t.Run("Withdrawal idempotence", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			// get nonce
			oldNonce, err := s.Adjs[0].PendingNonceAt(context.Background(), ethwallet.AsEthAddr(s.Accs[0].Address()))
			require.NoError(t, err)
			// withdraw
			testWithdraw(t, true)
			// get nonce
			nonce, err := s.Adjs[0].PendingNonceAt(context.Background(), ethwallet.AsEthAddr(s.Accs[0].Address()))
			require.NoError(t, err)
			assert.Equal(t, oldNonce, nonce, "Nonce must not change in subsequent withdrawals")
		}
	})
}

func TestWithdrawNonFinal(t *testing.T) {
	assert := assert.New(t)
	rng := pkgtest.Prng(t)
	// create test setup
	s := test.NewSetup(t, rng, 1, blockInterval, TxFinalityDepth)
	// create valid state and params
	params, state := channeltest.NewRandomParamsAndState(
		rng,
		channeltest.WithChallengeDuration(60),
		channeltest.WithParts(s.Parts...),
		channeltest.WithAssets(s.Asset),
		channeltest.WithIsFinal(false),
		channeltest.WithoutApp(),
		channeltest.WithLedgerChannel(true),
	)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTxTimeout)
	defer cancel()
	fundingReq := channel.NewFundingReq(params, state, channel.Index(0), state.Balances)
	require.NoError(t, s.Funders[0].Fund(ctx, *fundingReq), "funding should succeed")

	// create subscription
	adj := s.Adjs[0]
	sub, err := adj.Subscribe(ctx, params.ID())
	require.NoError(t, err)
	defer sub.Close()

	// register
	req := channel.AdjudicatorReq{
		Params: params,
		Acc:    s.Accs[0],
		Idx:    0,
		Tx:     testSignState(t, s.Accs, state),
	}
	require.NoError(t, adj.Register(ctx, req, nil))
	reg := sub.Next()
	t.Log("Registered ", reg)
	assert.False(reg.Timeout().IsElapsed(ctx),
		"registering non-final state should have non-elapsed timeout")
	assert.NoError(reg.Timeout().Wait(ctx))
	assert.True(reg.Timeout().IsElapsed(ctx), "timeout should have elapsed after Wait()")
	assert.NoError(adj.Withdraw(ctx, req, nil),
		"withdrawing should succeed after waiting for timeout")
}

func assertHoldingsZero(ctx context.Context, t *testing.T, cb *ethchannel.ContractBackend, params *channel.Params, _assets []channel.Asset) {
	t.Helper()
	alloc, err := onChainAllocation(ctx, cb, params, _assets)
	require.NoError(t, err, "Getting on-chain allocs should succeed")
	for i, assetalloc := range alloc {
		for j, a := range assetalloc {
			assert.Zerof(t, a.Sign(), "Allocation of asset[%d] and part[%d] non-zero.", j, i)
		}
	}
}
