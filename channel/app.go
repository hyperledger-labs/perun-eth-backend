package channel

import (
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"
	ethtestwallet "github.com/perun-network/perun-eth-backend/wallet/test"

	"math/rand"

	"perun.network/go-perun/channel"
)

var _ channel.AppID = new(AppID)

type AppID struct {
	*ethwallet.Address
}
type AppIDKey string

func (a AppID) Equal(b channel.AppID) bool {
	bTyped, ok := b.(*AppID)
	if !ok {
		return false
	}
	return a.Address.Equal(bTyped.Address)
}

// Key returns the key representation of this app identifier.
func (id AppID) Key() channel.AppIDKey {
	b, err := id.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return channel.AppIDKey(b)

}

func (a AppID) MarshalBinary() ([]byte, error) {
	data, err := a.Address.MarshalBinary()

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AppID) UnmarshalBinary(data []byte) error {
	addr := &ethwallet.Address{}
	err := addr.UnmarshalBinary(data)
	if err != nil {
		return err
	}
	appaddr := &AppID{addr}
	*a = *appaddr
	return nil
}

func NewRandomAppID(rng *rand.Rand) *AppID {
	addr := ethtestwallet.NewRandomAddress(rng)
	return &AppID{&addr}
}
