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

package client_test

import (
	"context"
	"math/big"
	"testing"

	"perun.network/go-perun/wallet"

	ethchanneltest "github.com/perun-network/perun-eth-backend/channel/test"
	ethclienttest "github.com/perun-network/perun-eth-backend/client/test"
	"perun.network/go-perun/apps/payment"
	chtest "perun.network/go-perun/channel/test"
	"perun.network/go-perun/client"
	clienttest "perun.network/go-perun/client/test"
	"perun.network/go-perun/wire"
	pkgtest "polycry.pt/poly-go/test"
)

func TestSubChannelHappy(t *testing.T) {
	rng := pkgtest.Prng(t)
	const A, B = 0, 1 // Indices of clients.
	s := ethchanneltest.NewSetup(t, rng, 2, ethclienttest.BlockInterval, TxFinalityDepth)
	setups := ethclienttest.MakeRoleSetups(rng, s, []string{"Susie", "Tim"})
	roles := [2]clienttest.Executer{
		clienttest.NewSusie(t, setups[A]),
		clienttest.NewTim(t, setups[B]),
	}
	// Enable stage synchronization.
	stages := roles[A].EnableStages()
	roles[B].SetStages(stages)

	// Build configuration.
	baseCfg := clienttest.MakeBaseExecConfig(
		[2]map[wallet.BackendID]wire.Address{wire.AddressMapfromAccountMap(setups[A].Identity), wire.AddressMapfromAccountMap(setups[B].Identity)},
		s.Asset,
		ethchanneltest.BackendID,
		[2]*big.Int{big.NewInt(100), big.NewInt(100)},
		client.WithoutApp(),
	)
	const (
		numSubChannels    = 2
		numSubSubChannels = 3
	)
	var (
		subChannelFunds = [][2]*big.Int{
			{big.NewInt(10), big.NewInt(10)},
			{big.NewInt(5), big.NewInt(5)},
		}
		subSubChannelFunds = [][2]*big.Int{
			{big.NewInt(3), big.NewInt(3)},
			{big.NewInt(2), big.NewInt(2)},
			{big.NewInt(1), big.NewInt(1)},
		}
		txAmount = big.NewInt(1)
	)
	cfg := clienttest.NewSusieTimExecConfig(
		baseCfg,
		numSubChannels,
		numSubSubChannels,
		subChannelFunds,
		subSubChannelFunds,
		client.WithApp(
			chtest.NewRandomAppAndData(rng, chtest.WithAppRandomizer(new(payment.Randomizer)), chtest.WithBackend(ethchanneltest.BackendID)),
		),
		txAmount,
	)

	ctx, cancel := context.WithTimeout(context.Background(), twoPartyTestTimeout)
	defer cancel()
	clienttest.ExecuteTwoPartyTest(ctx, t, roles, cfg)
}

func TestSubChannelDispute(t *testing.T) {
	rng := pkgtest.Prng(t)

	const A, B = 0, 1 // Indices of clients.
	s := ethchanneltest.NewSetup(t, rng, 2, ethclienttest.BlockInterval, TxFinalityDepth)
	setups := ethclienttest.MakeRoleSetups(rng, s, []string{"DisputeSusie", "DisputeTim"})
	roles := [2]clienttest.Executer{
		clienttest.NewDisputeSusie(t, setups[A]),
		clienttest.NewDisputeTim(t, setups[B]),
	}
	// Enable stage synchronization.
	stages := roles[A].EnableStages()
	roles[B].SetStages(stages)

	baseCfg := clienttest.MakeBaseExecConfig(
		[2]map[wallet.BackendID]wire.Address{wire.AddressMapfromAccountMap(setups[A].Identity), wire.AddressMapfromAccountMap(setups[B].Identity)},
		s.Asset,
		ethchanneltest.BackendID,
		[2]*big.Int{big.NewInt(100), big.NewInt(100)},
		client.WithoutApp(),
	)
	cfg := &clienttest.DisputeSusieTimExecConfig{
		BaseExecConfig:  baseCfg,
		SubChannelFunds: [2]*big.Int{big.NewInt(10), big.NewInt(10)},
		TxAmount:        big.NewInt(1),
	}

	ctx, cancel := context.WithTimeout(context.Background(), twoPartyTestTimeout)
	defer cancel()
	clienttest.ExecuteTwoPartyTest(ctx, t, roles, cfg)
}
