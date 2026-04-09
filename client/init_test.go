// Copyright 2025 - See NOTICE file for copyright holders.
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
	"sync"
	"testing"

	"github.com/sirupsen/logrus"

	plogrus "perun.network/go-perun/log/logrus"
)

// Client tests are expensive, and under the full race suite two concurrent
// simulated setups still overload the runner enough to cause package-level
// timeouts. Keep them serialized for stable CI runtimes.
var heavySimTestSlots = make(chan struct{}, 1)

func acquireHeavySimTestSlot(t *testing.T) func() {
	t.Helper()

	heavySimTestSlots <- struct{}{}
	var once sync.Once
	return func() {
		once.Do(func() {
			<-heavySimTestSlots
		})
	}
}

func init() {
	plogrus.Set(logrus.WarnLevel, &logrus.TextFormatter{ForceColors: true})
}
