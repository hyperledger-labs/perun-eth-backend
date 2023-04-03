package channel

import (
	"bytes"
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wire/perunio"
)

var _ channel.AppID = new(AppID)

type AppID struct {
	*ethwallet.Address
}
type AppIDKey string

func (id AppID) Equal(b channel.AppID) bool {
	bTyped, ok := b.(*AppID)
	if !ok {
		return false
	}

	return id.Address.Equal(bTyped.Address)
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
	var buf bytes.Buffer
	err := perunio.Encode(&buf, &a)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary unmarshals the asset from its binary representation.
func (a *AppID) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)
	return perunio.Decode(buf, &a)
}
