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

package test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/perun-network/perun-eth-backend/channel"
	"github.com/stretchr/testify/require"
	pchannel "perun.network/go-perun/channel"
)

// waitTime is the time to wait before funding a channel when the funder is not egoistic.
const waitTime = 2 * time.Second

// FunderFinishTime is the returned type of FundAll which includes the process time of each funder.
type FunderFinishTime struct {
	Index int
	Time  time.Duration
}

// FundAll calls fund for all funders simultaneously.
func FundAll(ctx context.Context, t *testing.T, funders []*channel.Funder, reqs []*pchannel.FundingReq, egoisticIndex int) ([]FunderFinishTime, error) {
	t.Helper()
	finishTimes := make([]FunderFinishTime, len(funders))
	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(funders))

	for i := range funders {
		i := i
		go func() {
			defer wg.Done()
			if i != egoisticIndex {
				time.Sleep(waitTime)
			}
			startTime := time.Now()
			err := funders[i].Fund(ctx, *reqs[i])
			require.NoError(t, err)
			finishTime := time.Now()
			mutex.Lock()
			finishTimes[i] = FunderFinishTime{
				Index: i,
				Time:  finishTime.Sub(startTime),
			}
			mutex.Unlock()
		}()
	}

	wg.Wait()
	return finishTimes, nil
}
