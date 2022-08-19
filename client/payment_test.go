// Copyright 2019 - See NOTICE file for copyright holders.
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
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/perun-network/perun-eth-backend/channel/test"
	ctest "github.com/perun-network/perun-eth-backend/client/test"
	"github.com/perun-network/perun-eth-backend/wallet"

	chtest "perun.network/go-perun/channel/test"
	"perun.network/go-perun/client"
	clienttest "perun.network/go-perun/client/test"
	"perun.network/go-perun/log"
	"perun.network/go-perun/wire"
	pkgtest "polycry.pt/poly-go/test"
)

const (
	twoPartyTestTimeout = 10 * time.Second
	TxFinalityDepth     = 3
)

func TestPaymentHappy(t *testing.T) {
	log.Info("Starting happy test")
	rng := pkgtest.Prng(t)

	const A, B = 0, 1 // Indices of Alice and Bob
	var (
		name = [2]string{"Alice", "Bob"}
		role [2]clienttest.Executer
	)

	s := test.NewSetup(t, rng, 2, ctest.BlockInterval, TxFinalityDepth)
	setup := ctest.MakeRoleSetups(rng, s, name[:])

	role[A] = clienttest.NewAlice(t, setup[A])
	role[B] = clienttest.NewBob(t, setup[B])
	// enable stages synchronization
	stages := role[A].EnableStages()
	role[B].SetStages(stages)

	execConfig := &clienttest.AliceBobExecConfig{
		BaseExecConfig: clienttest.MakeBaseExecConfig(
			[2]wire.Address{setup[A].Identity.Address(), setup[B].Identity.Address()},
			s.Asset,
			[2]*big.Int{big.NewInt(100), big.NewInt(100)},
			client.WithApp(chtest.NewRandomAppAndData(rng)),
		),
		NumPayments: [2]int{2, 2},
		TxAmounts:   [2]*big.Int{big.NewInt(5), big.NewInt(3)},
	}

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(i int) {
			defer wg.Done()
			log.Infof("Starting %s.Execute", name[i])
			role[i].Execute(execConfig)
		}(i)
	}

	wg.Wait()

	// Assert correct final balances
	aliceToBob := big.NewInt(int64(execConfig.NumPayments[A])*execConfig.TxAmounts[A].Int64() -
		int64(execConfig.NumPayments[B])*execConfig.TxAmounts[B].Int64())
	finalBalAlice := new(big.Int).Sub(execConfig.InitBals()[A], aliceToBob)
	finalBalBob := new(big.Int).Add(execConfig.InitBals()[B], aliceToBob)
	// reset context timeout
	ctx, cancel := context.WithTimeout(context.Background(), ctest.DefaultTimeout)
	defer cancel()
	assertBal := func(addr *wallet.Address, bal *big.Int) {
		b, err := s.SimBackend.BalanceAt(ctx, common.Address(*addr), nil)
		require.NoError(t, err)
		assert.Zero(t, bal.Cmp(b), "ETH balance mismatch")
	}

	assertBal(s.Recvs[A], finalBalAlice)
	assertBal(s.Recvs[B], finalBalBob)

	log.Info("Happy test done")
}

func TestPaymentDispute(t *testing.T) {
	log.Info("Starting dispute test")
	rng := pkgtest.Prng(t)

	const A, B = 0, 1 // Indices of Mallory and Carol
	var (
		name = [2]string{"Mallory", "Carol"}
		role [2]clienttest.Executer
	)

	s := test.NewSetup(t, rng, 2, ctest.BlockInterval, TxFinalityDepth)
	setup := ctest.MakeRoleSetups(rng, s, name[:])

	role[A] = clienttest.NewMallory(t, setup[A])
	role[B] = clienttest.NewCarol(t, setup[B])

	execConfig := &clienttest.MalloryCarolExecConfig{
		BaseExecConfig: clienttest.MakeBaseExecConfig(
			[2]wire.Address{setup[A].Identity.Address(), setup[B].Identity.Address()},
			s.Asset,
			[2]*big.Int{big.NewInt(100), big.NewInt(1)},
			client.WithoutApp(),
		),
		NumPayments: [2]int{5, 0},
		TxAmounts:   [2]*big.Int{big.NewInt(20), big.NewInt(0)},
	}

	ctx, cancel := context.WithTimeout(context.Background(), twoPartyTestTimeout)
	defer cancel()
	clienttest.ExecuteTwoPartyTest(ctx, t, role, execConfig)

	// Assert correct final balances
	netTransfer := big.NewInt(int64(execConfig.NumPayments[A])*execConfig.TxAmounts[A].Int64() -
		int64(execConfig.NumPayments[B])*execConfig.TxAmounts[B].Int64())
	finalBal := [2]*big.Int{
		new(big.Int).Sub(execConfig.InitBals()[A], netTransfer),
		new(big.Int).Add(execConfig.InitBals()[B], netTransfer),
	}
	// reset context timeout
	ctx, cancel = context.WithTimeout(context.Background(), ctest.DefaultTimeout)
	defer cancel()
	for i, bal := range finalBal {
		b, err := s.SimBackend.BalanceAt(ctx, common.Address(*s.Recvs[i]), nil)
		require.NoError(t, err)
		assert.Zero(t, b.Cmp(bal), "ETH balance mismatch")
	}

	log.Info("Dispute test done")
}
