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

package channel

import (
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"
	ethtestwallet "github.com/perun-network/perun-eth-backend/wallet/test"
	"math/rand"
	"perun.network/go-perun/channel"
)

var _ channel.AppID = new(AppID)

// AppID described an app identifier.
type AppID struct {
	*ethwallet.Address
}

// AppIDKey is the key representation of an app identifier.
type AppIDKey string

// Equal compares two AppID objects for equality.
func (a AppID) Equal(b channel.AppID) bool {
	bTyped, ok := b.(*AppID)
	if !ok {
		return false
	}
	return a.Address.Equal(bTyped.Address)
}

// Key returns the key representation of this app identifier.
func (a AppID) Key() channel.AppIDKey {
	b, err := a.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return channel.AppIDKey(b)

}

// MarshalBinary marshals the contents of AppID into a byte string.
func (a AppID) MarshalBinary() ([]byte, error) {
	data, err := a.Address.MarshalBinary()

	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnmarshalBinary converts a bytestring, representing AppID into the AppID struct.
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

// NewRandomAppID calls NewRandomAddress to generate a random Ethereum test address.
func NewRandomAppID(rng *rand.Rand) *AppID {
	addr := ethtestwallet.NewRandomAddress(rng)
	return &AppID{&addr}
}
