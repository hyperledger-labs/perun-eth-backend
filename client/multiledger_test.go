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

package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/perun-network/perun-eth-backend/client/test"

	ctest "perun.network/go-perun/client/test"
)

const (
	challengeDuration = 15 * uint64(time.Second/test.BlockInterval)
	testDuration      = 30 * time.Second
)

func TestMultiLedgerHappy(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	mlt := test.SetupMultiLedgerTest(t, testDuration)
	ctest.TestMultiLedgerHappy(ctx, t, mlt, challengeDuration)
}

func TestMultiLedgerDispute(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), testDuration)
	defer cancel()

	mlt := test.SetupMultiLedgerTest(t, testDuration)
	ctest.TestMultiLedgerDispute(ctx, t, mlt, challengeDuration)
}
