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

package test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	ethchannel "github.com/perun-network/perun-eth-backend/channel"
	chtest "github.com/perun-network/perun-eth-backend/channel/test"
	"github.com/perun-network/perun-eth-backend/wallet"
	"github.com/perun-network/perun-eth-backend/wallet/keystore"
	ethwire "github.com/perun-network/perun-eth-backend/wire"
	"github.com/stretchr/testify/require"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/channel/multi"
	"perun.network/go-perun/client"
	ctest "perun.network/go-perun/client/test"
	wtest "perun.network/go-perun/wallet/test"
	"perun.network/go-perun/watcher/local"
	"perun.network/go-perun/wire"
	"polycry.pt/poly-go/test"
)

const (
	txFinalityDepth = 1
	blockInterval   = 300 * time.Millisecond
)

// SetupMultiLedgerTest creates a multi-ledger test setup.
func SetupMultiLedgerTest(t *testing.T, testDuration time.Duration) ctest.MultiLedgerSetup {
	t.Helper()
	rng := test.Prng(t)

	ctx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	l1 := setupLedger(ctx, t, rng, big.NewInt(1337)) //nolint:gomnd
	l2 := setupLedger(ctx, t, rng, big.NewInt(1338)) //nolint:gomnd

	// Setup message bus.
	bus := wire.NewLocalBus()

	// Setup clients.
	c1 := setupClient(t, rng, l1, l2, bus)
	c2 := setupClient(t, rng, l1, l2, bus)

	// Fund accounts.
	l1.simSetup.SimBackend.FundAddress(ctx, wallet.AsEthAddr(c1.WalletAddress))
	l1.simSetup.SimBackend.FundAddress(ctx, wallet.AsEthAddr(c2.WalletAddress))
	l2.simSetup.SimBackend.FundAddress(ctx, wallet.AsEthAddr(c1.WalletAddress))
	l2.simSetup.SimBackend.FundAddress(ctx, wallet.AsEthAddr(c2.WalletAddress))

	//nolint:gomnd
	return ctest.MultiLedgerSetup{
		Client1: c1,
		Client2: c2,
		Asset1:  l1.asset,
		Asset2:  l2.asset,
		InitBalances: channel.Balances{
			{etherToWei(10), etherToWei(0)}, // Asset 1.
			{etherToWei(0), etherToWei(10)}, // Asset 2.
		},
		UpdateBalances1: channel.Balances{
			{etherToWei(5), etherToWei(5)}, // Asset 1.
			{etherToWei(3), etherToWei(7)}, // Asset 2.
		},
		UpdateBalances2: channel.Balances{
			{etherToWei(1), etherToWei(9)}, // Asset 1.
			{etherToWei(5), etherToWei(5)}, // Asset 2.
		},
		BalanceDelta:   etherToWei(0.00012),
		BalanceReader1: l1.simSetup.SimBackend,
		BalanceReader2: l2.simSetup.SimBackend,
	}
}

type testLedger struct {
	simSetup    *chtest.SimSetup
	adjudicator common.Address
	assetHolder common.Address
	asset       *ethchannel.Asset
}

func (l testLedger) ChainID() ethchannel.ChainID {
	return ethchannel.MakeChainID(l.simSetup.SimBackend.ChainID())
}

func setupLedger(ctx context.Context, t *testing.T, rng *rand.Rand, chainID *big.Int) testLedger {
	t.Helper()

	// Set chainID for SimulatedBackend.
	cfg := *params.AllEthashProtocolChanges
	cfg.ChainID = new(big.Int).Set(chainID)
	params.AllEthashProtocolChanges = &cfg
	simSetup := chtest.NewSimSetup(t, rng, txFinalityDepth, blockInterval)

	adjudicator, err := ethchannel.DeployAdjudicator(ctx, *simSetup.CB, simSetup.TxSender.Account)
	require.NoError(t, err)
	assetHolder, err := ethchannel.DeployETHAssetholder(ctx, *simSetup.CB, adjudicator, simSetup.TxSender.Account)
	require.NoError(t, err)
	asset := ethchannel.NewAsset(chainID, assetHolder)

	return testLedger{
		simSetup:    simSetup,
		adjudicator: adjudicator,
		assetHolder: assetHolder,
		asset:       asset,
	}
}

func setupClient(t *testing.T, rng *rand.Rand, l1, l2 testLedger, bus wire.Bus) ctest.Client {
	t.Helper()
	require := require.New(t)

	// Setup wallet and account.
	w := wtest.RandomWallet().(*keystore.Wallet)
	acc := w.NewRandomAccount(rng).(*keystore.Account)

	// Setup contract backends.
	signer1 := l1.simSetup.SimBackend.Signer
	cb1 := ethchannel.NewContractBackend(
		l1.simSetup.CB,
		l1.ChainID(),
		keystore.NewTransactor(*w, signer1),
		l1.simSetup.CB.TxFinalityDepth(),
	)
	signer2 := l2.simSetup.SimBackend.Signer
	cb2 := ethchannel.NewContractBackend(
		l2.simSetup.CB,
		l2.ChainID(),
		keystore.NewTransactor(*w, signer2),
		l2.simSetup.CB.TxFinalityDepth(),
	)

	// Setup funder.
	multiFunder := multi.NewFunder()
	funderL1 := ethchannel.NewFunder(cb1)
	funderL2 := ethchannel.NewFunder(cb2)
	registered := funderL1.RegisterAsset(*l1.asset, ethchannel.NewETHDepositor(), acc.Account)
	require.True(registered)
	registered = funderL2.RegisterAsset(*l2.asset, ethchannel.NewETHDepositor(), acc.Account)
	require.True(registered)
	multiFunder.RegisterFunder(l1.ChainID(), funderL1)
	multiFunder.RegisterFunder(l2.ChainID(), funderL2)

	// Setup adjudicator.
	multiAdj := multi.NewAdjudicator()
	adjL1 := chtest.NewSimAdjudicator(*l1.simSetup.CB, l1.adjudicator, acc.Account.Address, acc.Account)
	adjL2 := chtest.NewSimAdjudicator(*l2.simSetup.CB, l2.adjudicator, acc.Account.Address, acc.Account)
	multiAdj.RegisterAdjudicator(l1.ChainID(), adjL1)
	multiAdj.RegisterAdjudicator(l2.ChainID(), adjL2)

	// Setup watcher.
	watcher, err := local.NewWatcher(multiAdj)
	require.NoError(err)

	walletAddr := acc.Address().(*wallet.Address)
	wireAddr := &ethwire.Address{Address: walletAddr}

	c, err := client.New(
		wireAddr,
		bus,
		multiFunder,
		multiAdj,
		w,
		watcher,
	)
	require.NoError(err)

	return ctest.Client{
		Client:        c,
		Adjudicator1:  adjL1,
		Adjudicator2:  adjL2,
		WireAddress:   wireAddr,
		WalletAddress: walletAddr,
		Events:        make(chan channel.AdjudicatorEvent),
	}
}

func etherToWei(eth float64) *big.Int {
	weiFloat := new(big.Float).Mul(big.NewFloat(eth), new(big.Float).SetFloat64(params.Ether))
	wei, _ := weiFloat.Int(nil)
	return wei
}
