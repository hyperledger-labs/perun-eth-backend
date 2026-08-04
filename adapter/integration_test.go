// Copyright 2026 - See NOTICE file for copyright holders.
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

//nolint:testpackage // white-box: exercises unexported helpers and mocks
package adapter

import (
	"context"
	"errors"
	"math/big"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/perun-network/perun-eth-backend/bindings/liquiditypool"
	"github.com/perun-network/perun-eth-backend/channel"
	chtest "github.com/perun-network/perun-eth-backend/channel/test"
	"github.com/stretchr/testify/require"
)

const (
	integrationGasLimit   = uint64(600000)
	deployGasLimit        = uint64(6600000)
	integrationTestTimout = 10 * time.Second
)

type flakySubscriptionBackend struct {
	*channel.ContractBackend
	failures int32
}

func (f *flakySubscriptionBackend) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	if atomic.CompareAndSwapInt32(&f.failures, 0, 1) {
		return nil, errors.New("simulated websocket drop")
	}
	return f.ContractBackend.SubscribeFilterLogs(ctx, q, ch)
}

func setupPool(t *testing.T) (*chtest.SimSetup, common.Address, *liquiditypool.LiquidityPool) {
	t.Helper()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := chtest.NewSimSetup(t, rng, 1, 0)

	ctx, cancel := context.WithTimeout(context.Background(), integrationTestTimout)
	defer cancel()
	auth, err := s.CB.NewTransactor(ctx, deployGasLimit, s.TxSender.Account)
	require.NoError(t, err)
	addr, tx, _, err := liquiditypool.DeployLiquidityPool(
		auth, *s.CB, s.TxSender.Account.Address, big.NewInt(3600), big.NewInt(10))
	require.NoError(t, err)
	_, err = s.CB.ConfirmTransaction(ctx, tx, s.TxSender.Account)
	require.NoError(t, err)

	pool, err := liquiditypool.NewLiquidityPool(addr, s.CB)
	require.NoError(t, err)

	// Seed pool with free ETH so fundChannel can transfer principal to operator.
	auth, err = s.CB.NewTransactor(ctx, integrationGasLimit, s.TxSender.Account)
	require.NoError(t, err)
	auth.Value = big.NewInt(100)
	tx, err = pool.Deposit(auth)
	require.NoError(t, err)
	_, err = s.CB.ConfirmTransaction(ctx, tx, s.TxSender.Account)
	require.NoError(t, err)

	// Post operator bond so fundChannel passes the coverage requirement.
	auth, err = s.CB.NewTransactor(ctx, integrationGasLimit, s.TxSender.Account)
	require.NoError(t, err)
	auth.Value = big.NewInt(1000)
	tx, err = pool.BondETH(auth)
	require.NoError(t, err)
	_, err = s.CB.ConfirmTransaction(ctx, tx, s.TxSender.Account)
	require.NoError(t, err)

	return s, addr, pool
}

func TestIntegrationFundSettle_HappyPath(t *testing.T) {
	s, addr, pool := setupPool(t)

	a, err := NewLiquidityPoolAdapter(
		s.CB,
		addr,
		s.TxSender.Account,
		integrationGasLimit,
		WithRetryBackoff(15*time.Millisecond, 40*time.Millisecond),
		WithFinalityDepth(1),
	)
	require.NoError(t, err)
	require.Equal(t, 15*time.Millisecond, a.retryInitial)
	require.Equal(t, 40*time.Millisecond, a.retryMax)
	require.EqualValues(t, 1, a.finality)

	ctx, cancel := context.WithTimeout(context.Background(), integrationTestTimout)
	defer cancel()

	channelID := [32]byte{0xaa}
	principal := big.NewInt(40)
	ret := big.NewInt(45)

	require.NoError(t, a.FundChannel(ctx, channelID, principal))
	locked, err := pool.LockedByChannel(nil, channelID)
	require.NoError(t, err)
	require.Zero(t, locked.Cmp(principal))

	require.NoError(t, a.SettleChannel(ctx, channelID, ret))
	locked, err = pool.LockedByChannel(nil, channelID)
	require.NoError(t, err)
	require.Zero(t, locked.Sign())

	totalLocked, err := pool.TotalLockedETH(nil)
	require.NoError(t, err)
	require.Zero(t, totalLocked.Sign())
}

func TestIntegrationSubscribeChannelFunded_Live(t *testing.T) {
	s, addr, _ := setupPool(t)
	a, err := NewLiquidityPoolAdapter(s.CB, addr, s.TxSender.Account, integrationGasLimit, WithFinalityDepth(1))
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), integrationTestTimout)
	defer cancel()

	ch := make(chan ChannelFundedEvent, 2)
	sub, err := a.SubscribeChannelFunded(ctx, ch)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	channelID := [32]byte{0xbb}
	principal := big.NewInt(11)
	require.NoError(t, a.FundChannel(ctx, channelID, principal))

	select {
	case ev := <-ch:
		require.Equal(t, channelID, ev.ChannelID)
		require.Zero(t, principal.Cmp(ev.Principal))
		require.Equal(t, s.TxSender.Account.Address, ev.Operator)
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for ChannelFunded event")
	}
}

func TestIntegrationSubscribeChannelFunded_ReconnectAfterDrop(t *testing.T) {
	s, addr, _ := setupPool(t)
	a, err := NewLiquidityPoolAdapter(
		s.CB,
		addr,
		s.TxSender.Account,
		integrationGasLimit,
		WithRetryBackoff(10*time.Millisecond, 30*time.Millisecond),
	)
	require.NoError(t, err)

	flaky := &flakySubscriptionBackend{ContractBackend: s.CB}
	a.backend = flaky

	ctx, cancel := context.WithTimeout(context.Background(), integrationTestTimout)
	defer cancel()

	ch := make(chan ChannelFundedEvent, 2)
	sub, err := a.SubscribeChannelFunded(ctx, ch)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	// Let the first subscription attempt fail and trigger backoff/reconnect path.
	time.Sleep(40 * time.Millisecond)

	channelID := [32]byte{0xcc}
	principal := big.NewInt(13)
	require.NoError(t, a.FundChannel(ctx, channelID, principal))

	select {
	case ev := <-ch:
		require.Equal(t, channelID, ev.ChannelID)
		require.Zero(t, principal.Cmp(ev.Principal))
	case <-time.After(3 * time.Second):
		t.Fatal("timed out waiting for ChannelFunded event after reconnect")
	}

	require.GreaterOrEqual(t, atomic.LoadInt32(&flaky.failures), int32(1))
}
