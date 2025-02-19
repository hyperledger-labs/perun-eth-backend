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

package test

import (
	"math/rand"

	"github.com/perun-network/perun-eth-backend/channel"
	pchannel "perun.network/go-perun/channel"
	"perun.network/go-perun/channel/test"
)

func init() {
	test.SetRandomizer(new(randomizer), 1)
	test.SetNewRandomAppID(func(r *rand.Rand) pchannel.AppID {
		return channel.NewRandomAppID(r)
	}, 1)
}
