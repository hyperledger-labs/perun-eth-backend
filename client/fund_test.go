// Copyright 2022 - See NOTICE file for copyright holders.
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

package client_test

import (
	"context"
	stderrors "errors"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	channeltest "github.com/perun-network/perun-eth-backend/channel/test"
	ethclienttest "github.com/perun-network/perun-eth-backend/client/test"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	ctest "perun.network/go-perun/client/test"
	"perun.network/go-perun/wallet"
	"perun.network/go-perun/wire"
	"polycry.pt/poly-go/test"
)

func TestFundRecovery(t *testing.T) {
	t.Parallel()

	release := acquireHeavySimTestSlot(t)
	defer release()

	rng := test.Prng(t)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	params := ctest.FundSetup{
		ChallengeDuration: 1,
		FridaInitBal:      ethclienttest.EtherToWei(100),
		FredInitBal:       ethclienttest.EtherToWei(50),
		BalanceDelta:      ethclienttest.EtherToWei(0.001),
	}
	setupFn := func() ([2]ctest.RoleSetup, channel.Asset) {
		setup := channeltest.NewSetup(t, rng, 2, ethclienttest.BlockInterval, 1)
		for i, adj := range setup.Adjs {
			adj.Receiver = setup.Accs[i].Account.Address
		}
		roles := ethclienttest.MakeRoleSetups(rng, setup, []string{"Frida", "Fred"})
		for i := range roles {
			roles[i].Timeout = 90 * time.Second
		}
		var rolesArray [2]ctest.RoleSetup
		copy(rolesArray[:], roles)
		return rolesArray, setup.Asset
	}

	t.Run("failing funder proposer", func(t *testing.T) {
		roles, asset := setupFn()
		roles[0].Funder = ctest.FailingFunder{}
		runEthFundRecovery(ctx, t, params, roles, asset)
	})
	t.Run("failing funder proposee", func(t *testing.T) {
		roles, asset := setupFn()
		roles[1].Funder = ctest.FailingFunder{}
		runEthFundRecovery(ctx, t, params, roles, asset)
	})
	t.Run("failing funder both sides", func(t *testing.T) {
		roles, asset := setupFn()
		roles[0].Funder = ctest.FailingFunder{}
		roles[1].Funder = ctest.FailingFunder{}
		runEthFundRecovery(ctx, t, params, roles, asset)
	})
}

func runEthFundRecovery(
	ctx context.Context,
	t *testing.T,
	params ctest.FundSetup,
	setups [2]ctest.RoleSetup,
	asset channel.Asset,
) {
	t.Helper()

	const (
		fridaIdx = 0
		fredIdx  = 1
		numParts = 2
	)

	clients := ctest.NewClients(t, test.Prng(t), setups[:])
	frida, fred := clients[fridaIdx], clients[fredIdx]
	fridaWireAddr := wire.AddressMapfromAccountMap(frida.Identity)
	fredWireAddr := wire.AddressMapfromAccountMap(fred.Identity)
	fridaWalletAddr, fredWalletAddr := frida.WalletAddress, fred.WalletAddress

	balancesBefore := channel.Balances{{
		frida.BalanceReader.Balance(asset),
		fred.BalanceReader.Balance(asset),
	}}

	chsFred := make(chan *client.Channel, 1)
	errsFred := make(chan error, 1)
	go fred.Handle(
		ctest.AlwaysAcceptChannelHandler(ctx, fredWalletAddr, chsFred, errsFred),
		ctest.AlwaysRejectUpdateHandler(ctx, errsFred),
	)

	initAlloc := channel.NewAllocation(numParts, []wallet.BackendID{channeltest.BackendID}, asset)
	initAlloc.SetAssetBalances(asset, []*big.Int{params.FridaInitBal, params.FredInitBal})
	parts := []map[wallet.BackendID]wire.Address{fridaWireAddr, fredWireAddr}
	prop, err := client.NewLedgerChannelProposal(
		params.ChallengeDuration,
		fridaWalletAddr,
		initAlloc,
		parts,
	)
	require.NoError(t, err)

	chFrida, err := frida.ProposeChannel(ctx, prop)
	require.Error(t, err)
	var fundingErr *client.ChannelFundingError
	require.True(t, stderrors.As(err, &fundingErr))
	require.NotNil(t, chFrida)
	require.NoError(t, chFrida.Settle(ctx, false))

	chFred := <-chsFred
	require.NotNil(t, chFred)

	select {
	case err := <-errsFred:
		require.Error(t, err)
		fundingErr = nil
		require.True(t, stderrors.As(err, &fundingErr))
	case <-ctx.Done():
		require.NoError(t, ctx.Err())
	}
	require.NoError(t, chFred.Settle(ctx, false))

	balancesAfter := channel.Balances{{
		frida.BalanceReader.Balance(asset),
		fred.BalanceReader.Balance(asset),
	}}
	balancesDiff := balancesAfter.Sub(balancesBefore)
	expectedBalancesDiff := channel.Balances{{big.NewInt(0), big.NewInt(0)}}
	eq := ctest.EqualBalancesWithDelta(expectedBalancesDiff, balancesDiff, params.BalanceDelta)
	assert.Truef(t, eq, "final ledger balances incorrect: expected balance difference %v +- %v, got %v", expectedBalancesDiff, params.BalanceDelta, balancesDiff)
}
