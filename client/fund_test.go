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

func TestFundRecovery(t *testing.T) {
	rng := test.Prng(t)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	ctest.TestFundRecovery(
		ctx,
		t,
		ctest.FundSetup{
			ChallengeDuration: 1,
			FridaInitBal:      ethToWei(big.NewFloat(100)),
			FredInitBal:       ethToWei(big.NewFloat(50)),
			BalanceDelta:      ethToWei(big.NewFloat(0.001)),
		},
		func(r *rand.Rand) ([2]ctest.RoleSetup, channel.Asset) {
			setup := channeltest.NewSetup(t, rng, 2, ethclienttest.BlockInterval, 1)
			for i, adj := range setup.Adjs {
				adj.Receiver = setup.Accs[i].Account.Address
			}
			roles := ethclienttest.MakeRoleSetups(rng, setup, [2]string{"Frida", "Fred"})
			return roles, setup.Asset
		},
	)
}

func ethToWei(eth *big.Float) *big.Int {
	wei, _ := new(big.Float).Mul(eth, big.NewFloat(1e18)).Int(nil)
	return wei
}
