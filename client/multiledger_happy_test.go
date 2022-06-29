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

package client

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	ethchannel "github.com/perun-network/perun-eth-backend/channel"
	"github.com/perun-network/perun-eth-backend/wallet"
	"github.com/perun-network/perun-eth-backend/wallet/keystore"
	ethwire "github.com/perun-network/perun-eth-backend/wire"
	"github.com/stretchr/testify/require"

	chtest "github.com/perun-network/perun-eth-backend/channel/test"

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
	challengeDuration = 10
	testDuration      = 10 * time.Second
	txFinalityDepth   = 1
	blockInterval     = 100 * time.Millisecond
)

func TestMultiLedgerHappy(t *testing.T) {
	mlt := SetupMultiLedgerTest(t)
	ctest.TestMultiLedgerHappy(t, mlt, challengeDuration)
}

func SetupMultiLedgerTest(t *testing.T) ctest.MultiLedgerSetup {
	t.Helper()
	rng := test.Prng(t)

	ctx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	l1 := setupLedger(ctx, t, rng, big.NewInt(1337))
	l2 := setupLedger(ctx, t, rng, big.NewInt(1338))

	// Setup message bus.
	bus := wire.NewLocalBus()

	// Setup clients.
	c1 := setupClient(t, rng, l1, l2, bus)
	c2 := setupClient(t, rng, l1, l2, bus)

	// Fund accounts.
	l1.simSetup.SimBackend.FundAddress(ctx, c1.accountAddress())
	l1.simSetup.SimBackend.FundAddress(ctx, c2.accountAddress())
	l2.simSetup.SimBackend.FundAddress(ctx, c1.accountAddress())
	l2.simSetup.SimBackend.FundAddress(ctx, c2.accountAddress())

	return ctest.MultiLedgerSetup{
		Client1:        c1.Client,
		Client2:        c2.Client,
		Adjudicator1:   c1.adjL1,
		Adjudicator2:   c2.adjL2,
		Asset1:         l1.asset,
		Asset2:         l2.asset,
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

// FIXME: This wrapper may should be implemented in go-perun.
type testClient struct {
	ctest.Client
	adjL1 channel.Adjudicator
	adjL2 channel.Adjudicator
}

func (c *testClient) accountAddress() common.Address {
	return wallet.AsEthAddr(c.WalletAddress)
}

func setupClient(t *testing.T, rng *rand.Rand, l1, l2 testLedger, bus wire.Bus) testClient {
	require := require.New(t)

	// Setup wallet and account.
	w := wtest.RandomWallet().(*keystore.Wallet)
	acc := w.NewRandomAccount(rng).(*keystore.Account)

	// Setup contract backends.
	signer1 := chtest.SignerForChainID(l1.ChainID().Int)
	cb1 := ethchannel.NewContractBackend(
		l1.simSetup.CB,
		keystore.NewTransactor(*w, signer1),
		l1.simSetup.CB.TxFinalityDepth(),
	)
	signer2 := chtest.SignerForChainID(l2.ChainID().Int)
	cb2 := ethchannel.NewContractBackend(
		l2.simSetup.CB,
		keystore.NewTransactor(*w, signer2),
		l2.simSetup.CB.TxFinalityDepth(),
	)

	// Setup funder.
	multiFunder := multi.NewFunder()
	funderL1 := ethchannel.NewFunder(cb1, l1.ChainID())
	funderL2 := ethchannel.NewFunder(cb2, l2.ChainID())
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

	return testClient{
		Client: ctest.Client{
			Client:        c,
			WireAddress:   wireAddr,
			WalletAddress: walletAddr,
			Events:        make(chan channel.AdjudicatorEvent),
		},
		adjL1: adjL1,
		adjL2: adjL2,
	}
}
