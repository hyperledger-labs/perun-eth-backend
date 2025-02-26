// Copyright 2025 - See NOTICE file for copyright holders.
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
	"math/rand"
	"time"

	"perun.network/go-perun/wallet"
	wtest "perun.network/go-perun/wallet/test"

	ethctest "github.com/perun-network/perun-eth-backend/channel/test"
	ethwtest "github.com/perun-network/perun-eth-backend/wallet/test"

	clienttest "perun.network/go-perun/client/test"
	"perun.network/go-perun/watcher/local"
	"perun.network/go-perun/wire"
	wiretest "perun.network/go-perun/wire/test"
)

const (
	// DefaultTimeout is the default timeout for client tests.
	DefaultTimeout = 20 * time.Second
	// BlockInterval is the default block interval for the simulated chain.
	BlockInterval = 200 * time.Millisecond
	// challenge duration in blocks that is used by MakeRoleSetups.
	challengeDurationBlocks = 90
)

// MakeRoleSetups creates a two party client test setup with the provided names.
func MakeRoleSetups(rng *rand.Rand, s *ethctest.Setup, names []string) []clienttest.RoleSetup {
	setups := make([]clienttest.RoleSetup, len(names))
	bus := wire.NewLocalBus()
	for i := 0; i < len(setups); i++ {
		watcher, err := local.NewWatcher(s.Adjs[i])
		if err != nil {
			panic("Error initializing watcher: " + err.Error())
		}
		setups[i] = clienttest.RoleSetup{
			Name:        names[i],
			Identity:    wiretest.NewRandomAccountMap(rng, ethwtest.BackendID),
			Bus:         bus,
			Funder:      s.Funders[i],
			Adjudicator: s.Adjs[i],
			Watcher:     watcher,
			Wallet:      map[wallet.BackendID]wtest.Wallet{ethwtest.BackendID: ethwtest.NewTmpWallet()},
			Timeout:     DefaultTimeout,
			// Scaled due to simbackend automining progressing faster than real time.
			ChallengeDuration: challengeDurationBlocks * uint64(time.Second/BlockInterval),
			Errors:            make(chan error),
			BalanceReader:     s.SimBackend.NewBalanceReader(s.Accs[i].Address()),
		}
	}
	return setups
}
