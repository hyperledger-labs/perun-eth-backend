package test

import (
	"context"
	"github.com/perun-network/perun-eth-backend/channel"
	pchannel "perun.network/go-perun/channel"
	pkgerrors "polycry.pt/poly-go/errors"
	"sync"
	"time"
)

type FunderFinishTime struct {
	Index int           // Funder ID
	Time  time.Duration // Time when Fund method returned
}

func FundAll(ctx context.Context, funders []*channel.Funder, reqs []*pchannel.FundingReq) ([]FunderFinishTime, error) {
	g := pkgerrors.NewGatherer()
	finishTimes := make([]FunderFinishTime, len(funders))
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(len(funders))
	for i := range funders {
		i := i
		g.Go(func() error {
			defer wg.Done()
			startTime := time.Now()
			err := funders[i].Fund(ctx, *reqs[i])
			finishTime := time.Now()
			mutex.Lock()
			finishTimes[i] = FunderFinishTime{
				Index: i,
				Time:  finishTime.Sub(startTime),
			}
			mutex.Unlock()
			return err
		})
	}
	wg.Wait()

	return finishTimes, nil
}
