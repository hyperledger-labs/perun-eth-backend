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

package test

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"testing"
	"time"

	ethctest "github.com/perun-network/perun-eth-backend/channel/test"
	ethwtest "github.com/perun-network/perun-eth-backend/wallet/test"
	"github.com/stretchr/testify/assert"

	clienttest "perun.network/go-perun/client/test"
	"perun.network/go-perun/watcher/local"
	"perun.network/go-perun/wire"
	perunnet "perun.network/go-perun/wire/net"
	"perun.network/go-perun/wire/net/simple"
	perunio "perun.network/go-perun/wire/perunio/serializer"
	wiretest "perun.network/go-perun/wire/test"
)

const (
	// DefaultTimeout is the default timeout for client tests.
	DefaultTimeout = 20 * time.Second
	// BlockInterval is the default block interval for the simulated chain.
	BlockInterval = 200 * time.Millisecond
	// challenge duration in blocks that is used by MakeRoleSetups.
	challengeDurationBlocks = 90

	// dialerTimeout is the timeout for dialing a connection.
	dialerTimeout = 15 * time.Second
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
			Identity:    wiretest.NewRandomAccount(rng),
			Bus:         bus,
			Funder:      s.Funders[i],
			Adjudicator: s.Adjs[i],
			Watcher:     watcher,
			Wallet:      ethwtest.NewTmpWallet(),
			Timeout:     DefaultTimeout,
			// Scaled due to simbackend automining progressing faster than real time.
			ChallengeDuration: challengeDurationBlocks * uint64(time.Second/BlockInterval),
			Errors:            make(chan error),
			BalanceReader:     s.SimBackend.NewBalanceReader(s.Accs[i].Address()),
		}
	}
	return setups
}

// MakeNetRoleSetups creates a two party client test setup with the provided names and uses the default TLS-bus from go-perun.
func MakeNetRoleSetups(t *testing.T, rng *rand.Rand, s *ethctest.Setup, names []string) []clienttest.RoleSetup {
	t.Helper()
	setups := make([]clienttest.RoleSetup, len(names))
	commonName := "127.0.0.1"
	sans := []string{"127.0.0.1", "localhost"}
	tlsConfigs, err := GenerateSelfSignedCertConfigs(commonName, sans, len(names))
	if err != nil {
		panic("Error generating TLS configs: " + err.Error())
	}

	hosts := make([]string, len(names))
	for i := 0; i < len(names); i++ {
		port, err := findFreePort()
		assert.NoError(t, err, "Error finding free port")
		hosts[i] = fmt.Sprintf("127.0.0.1:%d", port)
	}

	dialers, listeners := makeSimpleDialersListeners(t, tlsConfigs, hosts)

	for i := 0; i < len(setups); i++ {
		watcher, err := local.NewWatcher(s.Adjs[i])
		if err != nil {
			panic("Error initializing watcher: " + err.Error())
		}

		acc := wiretest.NewRandomAccount(rng)

		for j := 0; j < len(setups); j++ {
			dialers[j].Register(acc.Address(), hosts[i])
		}

		bus := perunnet.NewBus(acc, dialers[i], perunio.Serializer())
		go bus.Listen(listeners[i])

		setups[i] = clienttest.RoleSetup{
			Name:        names[i],
			Identity:    acc,
			Bus:         bus,
			Funder:      s.Funders[i],
			Adjudicator: s.Adjs[i],
			Watcher:     watcher,
			Wallet:      ethwtest.NewTmpWallet(),
			Timeout:     DefaultTimeout,
			// Scaled due to simbackend automining progressing faster than real time.
			ChallengeDuration: challengeDurationBlocks * uint64(time.Second/BlockInterval),
			Errors:            make(chan error),
			BalanceReader:     s.SimBackend.NewBalanceReader(s.Accs[i].Address()),
		}
	}
	return setups
}

func makeSimpleDialersListeners(t *testing.T, tlsConfigs []*tls.Config, hosts []string) ([]*simple.Dialer, []*simple.Listener) {
	t.Helper()
	dialers := make([]*simple.Dialer, len(tlsConfigs))
	listeners := make([]*simple.Listener, len(tlsConfigs))

	var err error
	for i, tlsConfig := range tlsConfigs {
		dialers[i] = simple.NewTCPDialer(dialerTimeout, tlsConfig)
		listeners[i], err = simple.NewTCPListener(hosts[i], tlsConfig)
		assert.NoError(t, err, "Error creating listener")
	}

	return dialers, listeners
}

func findFreePort() (int, error) {
	// Create a listener on a random port to get an available port.
	l, err := net.Listen("tcp", "127.0.0.1:0") // Use ":0" to bind to a random free port
	if err != nil {
		return 0, err
	}
	defer l.Close()

	// Get the port from the listener's address
	addr := l.Addr().(*net.TCPAddr)
	return addr.Port, nil
}
