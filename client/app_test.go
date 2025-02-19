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

package client_test

import (
	"context"
	"math/big"
	"testing"

	ethchannel "github.com/perun-network/perun-eth-backend/channel"
	"github.com/perun-network/perun-eth-backend/channel/test"
	ctest "github.com/perun-network/perun-eth-backend/client/test"
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"
	"github.com/stretchr/testify/require"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	clienttest "perun.network/go-perun/client/test"
	"perun.network/go-perun/wallet"
	"perun.network/go-perun/wire"
	pkgtest "polycry.pt/poly-go/test"
)

func TestProgression(t *testing.T) {
	rng := pkgtest.Prng(t)

	names := []string{"Paul", "Paula"}
	backendSetup := test.NewSetup(t, rng, 2, ctest.BlockInterval, TxFinalityDepth)
	roleSetups := ctest.MakeRoleSetups(rng, backendSetup, names)
	clients := [2]clienttest.Executer{
		clienttest.NewPaul(t, roleSetups[0]),
		clienttest.NewPaula(t, roleSetups[1]),
	}

	appAddress := deployMockApp(t, backendSetup)
	appAddrBackend := appAddress.(*ethwallet.Address)
	appID := &ethchannel.AppID{Address: appAddrBackend}
	app := channel.NewMockApp(appID)
	channel.RegisterApp(app)

	execConfig := &clienttest.ProgressionExecConfig{
		BaseExecConfig: clienttest.MakeBaseExecConfig(
			clientAddresses(roleSetups),
			backendSetup.Asset,
			ethwallet.BackendID,
			[2]*big.Int{big.NewInt(99), big.NewInt(1)},
			client.WithApp(app, channel.NewMockOp(channel.OpValid)),
		),
	}

	ctx, cancel := context.WithTimeout(context.Background(), twoPartyTestTimeout)
	defer cancel()
	clienttest.ExecuteTwoPartyTest(ctx, t, clients, execConfig)
}

func deployMockApp(t *testing.T, s *test.Setup) wallet.Address {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), ctest.DefaultTimeout)
	defer cancel()
	addr, err := ethchannel.DeployTrivialApp(ctx, *s.CB, s.TxSender.Account)
	require.NoError(t, err)
	return ethwallet.AsWalletAddr(addr)
}

func clientAddresses(roleSetups []clienttest.RoleSetup) (addresses [2]map[wallet.BackendID]wire.Address) {
	for i := 0; i < len(roleSetups); i++ {
		addresses[i] = wire.AddressMapfromAccountMap(roleSetups[i].Identity)
	}
	return
}
