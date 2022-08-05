package client_test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	channeltest "github.com/perun-network/perun-eth-backend/channel/test"
	ethclienttest "github.com/perun-network/perun-eth-backend/client/test"
	ctest "perun.network/go-perun/client/test"
	"polycry.pt/poly-go/test"
)

func TestFund(t *testing.T) {
	rng := test.Prng(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ctest.TestFundRecovery(
		ctx,
		t,
		ctest.FundSetup{
			ChallengeDuration: 1,
			FridaInitBal:      big.NewInt(100),
			FredInitBal:       big.NewInt(50),
		},
		func(r *rand.Rand) [2]ctest.RoleSetup {
			setup := channeltest.NewSetup(t, rng, 2, ethclienttest.BlockInterval, 1)
			roles := ethclienttest.MakeRoleSetups(rng, setup, [2]string{"Frida", "Fred"})
			return roles
		},
	)
}
