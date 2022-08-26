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
	"math/big"
	"math/rand"
	"testing"
	"time"

	channeltest "github.com/perun-network/perun-eth-backend/channel/test"
	ethclienttest "github.com/perun-network/perun-eth-backend/client/test"
	"perun.network/go-perun/channel"
	ctest "perun.network/go-perun/client/test"
	"polycry.pt/poly-go/test"
)

func TestVirtualChannelOptimistic(t *testing.T) {
	rng := test.Prng(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctest.TestVirtualChannelOptimistic(
		ctx,
		t,
		createVirtualChannelSetup(t, rng),
	)
}

func TestVirtualChannelDispute(t *testing.T) {
	rng := test.Prng(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctest.TestVirtualChannelDispute(
		ctx,
		t,
		createVirtualChannelSetup(t, rng),
	)
}

func createVirtualChannelSetup(t *testing.T, rng *rand.Rand) ctest.VirtualChannelSetup {
	t.Helper()
	clients, asset := setupVirtualChannelClients(t, rng)
	return ctest.VirtualChannelSetup{
		ChallengeDuration: challengeDuration,
		Clients:           clients,
		Balances: ctest.VirtualChannelBalances{
			InitBalsAliceIngrid: []*big.Int{ethclienttest.EtherToWei(10), ethclienttest.EtherToWei(10)},
			InitBalsBobIngrid:   []*big.Int{ethclienttest.EtherToWei(10), ethclienttest.EtherToWei(10)},
			InitBalsAliceBob:    []*big.Int{ethclienttest.EtherToWei(5), ethclienttest.EtherToWei(5)},
			VirtualBalsUpdated:  []*big.Int{ethclienttest.EtherToWei(2), ethclienttest.EtherToWei(8)},
			FinalBalsAlice:      []*big.Int{ethclienttest.EtherToWei(7), ethclienttest.EtherToWei(13)},
			FinalBalsBob:        []*big.Int{ethclienttest.EtherToWei(13), ethclienttest.EtherToWei(7)},
		},
		BalanceDelta:       ethclienttest.EtherToWei(0.001),
		Asset:              asset,
		Rng:                rng,
		WaitWatcherTimeout: 100 * time.Millisecond,
	}
}

func setupVirtualChannelClients(t *testing.T, rng *rand.Rand) ([3]ctest.RoleSetup, channel.Asset) {
	t.Helper()
	setup := channeltest.NewSetup(t, rng, 3, ethclienttest.BlockInterval, 1)
	for i, adj := range setup.Adjs {
		adj.Receiver = setup.Accs[i].Account.Address
	}
	roles := ethclienttest.MakeRoleSetups(rng, setup, []string{"Alice", "Bob", "Ingrid"})
	var rolesArray [3]ctest.RoleSetup
	copy(rolesArray[:], roles)
	return rolesArray, setup.Asset
}
