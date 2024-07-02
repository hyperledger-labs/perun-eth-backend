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

package wire_test

import (
	"math/rand"

	ethwire "github.com/perun-network/perun-eth-backend/wire" // import wire for testing
	"perun.network/go-perun/wire"
	wiretest "perun.network/go-perun/wire/test"
)

func init() {
	wire.SetNewAddressFunc(func() wire.Address {
		return ethwire.NewAddress()
	})

	wiretest.SetNewRandomAddress(func(rng *rand.Rand) wire.Address {
		return ethwire.NewRandomAddress(rng)
	})

	wiretest.SetNewRandomAccount(func(rng *rand.Rand) wire.Account {
		return ethwire.NewRandomAccount(rng)
	})
}
