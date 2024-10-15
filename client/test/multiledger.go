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
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"
	"github.com/perun-network/perun-eth-backend/wallet/keystore"
	ethwire "github.com/perun-network/perun-eth-backend/wire"
	"github.com/stretchr/testify/require"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/channel/multi"
	"perun.network/go-perun/client"
	ctest "perun.network/go-perun/client/test"
	"perun.network/go-perun/wallet"
	wtest "perun.network/go-perun/wallet/test"
	"perun.network/go-perun/watcher/local"
	"perun.network/go-perun/wire"
	"polycry.pt/poly-go/test"
)

const (
	txFinalityDepth    = 1
	blockInterval      = 300 * time.Millisecond
	defaultETHGasLimit = uint64(50000)
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
	l1.simSetup.SimBackend.FundAddress(ctx, ethwallet.AsEthAddr(c1.WalletAddress[1]))
	l1.simSetup.SimBackend.FundAddress(ctx, ethwallet.AsEthAddr(c2.WalletAddress[1]))
	l2.simSetup.SimBackend.FundAddress(ctx, ethwallet.AsEthAddr(c1.WalletAddress[1]))
	l2.simSetup.SimBackend.FundAddress(ctx, ethwallet.AsEthAddr(c2.WalletAddress[1]))

	//nolint:gomnd
	return ctest.MultiLedgerSetup{
		Client1: c1,
		Client2: c2,
		Asset1:  l1.asset,
		Asset2:  l2.asset,
		InitBalances: channel.Balances{
			{EtherToWei(8), EtherToWei(2)}, // Asset 1.
			{EtherToWei(2), EtherToWei(8)}, // Asset 2.
		},
		UpdateBalances1: channel.Balances{
			{EtherToWei(5), EtherToWei(5)}, // Asset 1.
			{EtherToWei(3), EtherToWei(7)}, // Asset 2.
		},
		UpdateBalances2: channel.Balances{
			{EtherToWei(1), EtherToWei(9)}, // Asset 1.
			{EtherToWei(5), EtherToWei(5)}, // Asset 2.
		},
		BalanceDelta: EtherToWei(0.00012),
	}
}

type testLedger struct {
	simSetup    *chtest.SimSetup
	adjudicator common.Address
	assetHolder common.Address
	asset       *ethchannel.Asset
}

func (l testLedger) AssetID() multi.AssetID {
	return ethchannel.MakeAssetID(ethchannel.MakeChainID(l.simSetup.SimBackend.ChainID()).Int)
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

func setupClient(t *testing.T, rng *rand.Rand, l1, l2 testLedger, bus wire.Bus) ctest.MultiLedgerClient {
	t.Helper()
	require := require.New(t)

	// Setup wallet and account.
	w := map[wallet.BackendID]wtest.Wallet{1: wtest.RandomWallet(1).(*keystore.Wallet)}
	acc := w[1].NewRandomAccount(rng).(*keystore.Account)

	// Setup contract backends.
	signer1 := l1.simSetup.SimBackend.Signer
	cb1 := ethchannel.NewContractBackend(
		l1.simSetup.CB,
		ethchannel.MakeAssetID(ethchannel.MakeChainID(l1.simSetup.SimBackend.ChainID()).Int),
		keystore.NewTransactor(*w[1].(*keystore.Wallet), signer1),
		l1.simSetup.CB.TxFinalityDepth(),
	)
	signer2 := l2.simSetup.SimBackend.Signer
	cb2 := ethchannel.NewContractBackend(
		l2.simSetup.CB,
		ethchannel.MakeAssetID(ethchannel.MakeChainID(l2.simSetup.SimBackend.ChainID()).Int),
		keystore.NewTransactor(*w[1].(*keystore.Wallet), signer2),
		l2.simSetup.CB.TxFinalityDepth(),
	)

	// Setup funder.
	multiFunder := multi.NewFunder()
	funderL1 := ethchannel.NewFunder(cb1)
	funderL2 := ethchannel.NewFunder(cb2)
	registered := funderL1.RegisterAsset(*l1.asset, ethchannel.NewETHDepositor(defaultETHGasLimit), acc.Account)
	require.True(registered)
	registered = funderL1.RegisterAsset(*l2.asset, ethchannel.NewNoOpDepositor(), acc.Account)
	require.True(registered)
	registered = funderL2.RegisterAsset(*l1.asset, ethchannel.NewNoOpDepositor(), acc.Account)
	require.True(registered)
	registered = funderL2.RegisterAsset(*l2.asset, ethchannel.NewETHDepositor(defaultETHGasLimit), acc.Account)
	require.True(registered)
	multiFunder.RegisterFunder(l1.AssetID(), funderL1)
	multiFunder.RegisterFunder(l2.AssetID(), funderL2)

	// Setup adjudicator.
	multiAdj := multi.NewAdjudicator()
	adjL1 := chtest.NewSimAdjudicator(*l1.simSetup.CB, l1.adjudicator, acc.Account.Address, acc.Account)
	adjL2 := chtest.NewSimAdjudicator(*l2.simSetup.CB, l2.adjudicator, acc.Account.Address, acc.Account)
	multiAdj.RegisterAdjudicator(l1.AssetID(), adjL1)
	multiAdj.RegisterAdjudicator(l2.AssetID(), adjL2)

	// Setup watcher.
	watcher, err := local.NewWatcher(multiAdj)
	require.NoError(err)

	walletAddr := acc.Address().(*ethwallet.Address)
	wireAddr := &ethwire.Address{Address: walletAddr}
	perunWallet := map[wallet.BackendID]wallet.Wallet{1: w[1]}
	c, err := client.New(
		map[wallet.BackendID]wire.Address{1: wireAddr},
		bus,
		multiFunder,
		multiAdj,
		perunWallet,
		watcher,
	)
	require.NoError(err)

	return ctest.MultiLedgerClient{
		Client:         c,
		Adjudicator1:   adjL1,
		Adjudicator2:   adjL2,
		WireAddress:    map[wallet.BackendID]wire.Address{1: wireAddr},
		WalletAddress:  map[wallet.BackendID]wallet.Address{1: walletAddr},
		Events:         make(chan channel.AdjudicatorEvent),
		BalanceReader1: l1.simSetup.SimBackend.NewBalanceReader(acc.Address()),
		BalanceReader2: l2.simSetup.SimBackend.NewBalanceReader(acc.Address()),
	}
}

// EtherToWei converts eth to wei.
func EtherToWei(eth float64) *big.Int {
	weiFloat := new(big.Float).Mul(big.NewFloat(eth), new(big.Float).SetFloat64(params.Ether))
	wei, _ := weiFloat.Int(nil)
	return wei
}
