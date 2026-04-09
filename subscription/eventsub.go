// Copyright 2021 - See NOTICE file for copyright holders.
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

package subscription

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
	"perun.network/go-perun/log"
	pkgsync "polycry.pt/poly-go/sync"
)

type (
	// EventSub generic event subscription.
	// Can be used on any Contract with any Event.
	// This EventSub does not prevent duplicates.
	EventSub struct {
		closer   pkgsync.Closer
		contract *bind.BoundContract
		eFact    EventFactory

		watchLogs, filterLogs chan types.Log
		watchSub, filterSub   event.Subscription
	}

	// Event is a generic on-chain event.
	Event struct {
		Name   string      // Name of the event. Must match the ABI definition.
		Data   interface{} // Instance of the concrete Event type.
		Filter Filter      // Filters Events by their body.
		Log    types.Log   // Raw original log for additional information.
	}

	// EventFactory is used to create `Event`s.
	// The `Data` and `Name` fields must be set. `Filter` is optional.
	EventFactory func() *Event

	// Filter can be used to filter events.
	// Look at `TestEventSub_Filter` test or the auto generated Filter- and
	// Watch-functions in the bindings/ folder for an example.
	Filter [][]interface{}
)

const transientFilterLogsErr = "failed to retrieve log value pointer"

// NewEventSub creates a new `EventSub`. Should always be closed with `Close`.
// `pastBlocks` can be used to define how many blocks into the past the sub
// should query.
func NewEventSub(ctx context.Context, chain ethereum.ChainReader, contract *bind.BoundContract, eFact EventFactory, pastBlocks uint64) (*EventSub, error) {
	endBlock, err := latestBlockNumber(ctx, chain)
	if err != nil {
		return nil, errors.WithMessage(err, "calculating ending block number")
	}
	// Split the subscription into a stable historical range and a live range.
	// Using endBlock+1 for the live watch avoids replaying the historical logs
	// through the watch stream, which can otherwise delay or starve fresh
	// events behind duplicates under heavier test load.
	watchStart := endBlock + 1
	evt := eFact()
	watchOpts := &bind.WatchOpts{Start: &watchStart}
	watchLogs, watchSub, err := contract.WatchLogs(watchOpts, evt.Name, evt.Filter...)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, errors.WithMessage(err, "watching logs")
	}
	var (
		filterLogs chan types.Log
		filterSub event.Subscription
	)
	if pastBlocks > 0 {
		startBlock, err := calcStartBlock(ctx, chain, pastBlocks)
		if err != nil {
			watchSub.Unsubscribe()
			return nil, errors.WithMessage(err, "calculating starting block number")
		}
		// Read past events.
		filterOpts := &bind.FilterOpts{Start: startBlock, End: &endBlock}
		filterLogs, filterSub, err = filterLogsWithRetry(ctx, contract, filterOpts, evt.Name, evt.Filter...)
		if err != nil {
			watchSub.Unsubscribe()
			err = cherrors.CheckIsChainNotReachableError(err)
			return nil, errors.WithMessage(err, "filtering logs")
		}
	}

	ret := &EventSub{
		contract:   contract,
		eFact:      eFact,
		watchLogs:  watchLogs,
		filterLogs: filterLogs,
		watchSub:   watchSub,
		filterSub:  filterSub,
	}
	ret.closer.OnCloseAlways(func() {
		watchSub.Unsubscribe()
		if filterSub != nil {
			filterSub.Unsubscribe()
		}
	})
	return ret, nil
}

func filterLogsWithRetry(ctx context.Context, contract *bind.BoundContract, opts *bind.FilterOpts, name string, query ...[]interface{}) (chan types.Log, event.Subscription, error) {
	const (
		maxAttempts = 5
		retryDelay  = 20 * time.Millisecond
	)
	var (
		logs chan types.Log
		sub  event.Subscription
		err  error
	)
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		logs, sub, err = contract.FilterLogs(opts, name, query...)
		// Geth v1.17.x can transiently return this internal error while filtering
		// logs around freshly mined blocks. Keep the version-coupled retry localized
		// here until geth exposes a typed error.
		if err == nil || !strings.Contains(err.Error(), transientFilterLogsErr) || attempt == maxAttempts {
			return logs, sub, err
		}
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		case <-time.After(retryDelay):
		}
	}
	return logs, sub, err
}

func latestBlockNumber(ctx context.Context, chain ethereum.ChainReader) (uint64, error) {
	current, err := chain.HeaderByNumber(ctx, nil)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return 0, errors.WithMessage(err, "retrieving latest block")
	}
	return current.Number.Uint64(), nil
}

func calcStartBlock(ctx context.Context, chain ethereum.ChainReader, pastBlocks uint64) (uint64, error) {
	endBlock, err := latestBlockNumber(ctx, chain)
	if err != nil {
		return 0, err
	}
	if endBlock <= pastBlocks {
		return 1, nil
	}
	return endBlock - pastBlocks, nil
}

// Read reads all past and future events into `sink`.
// Can be aborted by cancelling `ctx` or `Close()`.
// It is possible that the same event is read more than once.
// After casting the generic event to a specific type, the `Raw` log field
// will be nil. Use the `Log` field of the generic event instead.
func (s *EventSub) Read(ctx context.Context, sink chan<- *Event) error {
	// First read into the past.
	if err := s.readPast(ctx, sink); err != nil {
		return errors.WithMessage(err, "reading logs")
	}
	// Then wait for new events.
	if err := s.readFuture(ctx, sink); err != nil {
		return errors.WithMessage(err, "reading logs")
	}
	return nil
}

// ReadPast reads all past events into `sink`.
// Can be aborted by cancelling `ctx` or `Close()`.
// It is possible that the same event is read more than once.
// After casting the generic event to a specific type, the `Raw` log field
// will be nil. Use the `Log` field of the generic event instead.
func (s *EventSub) ReadPast(ctx context.Context, sink chan<- *Event) error {
	return errors.WithMessage(s.readPast(ctx, sink), "reading logs")
}

func (s *EventSub) readPast(ctx context.Context, sink chan<- *Event) error {
	if s.filterLogs == nil || s.filterSub == nil {
		return nil
	}

	var logs []types.Log
	// Two read loops are needed if the event sub is closed before all events
	// could be read.
read1:
	for {
		select {
		case <-s.closer.Closed():
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		select {
		case log := <-s.filterLogs:
			logs = append(logs, log)
		case err := <-s.filterSub.Err():
			if err != nil {
				err = cherrors.CheckIsChainNotReachableError(err)
				return err
			}
			break read1
		case <-ctx.Done():
			return ctx.Err()
		case <-s.closer.Closed():
			return nil
		}
	}
read2:
	for {
		select {
		case log := <-s.filterLogs:
			logs = append(logs, log)
		case <-s.closer.Closed():
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
			break read2
		}
	}

	for _, log := range logs {
		if err := s.processLog(ctx, log, sink); err != nil {
			return err
		}
	}
	return nil
}

func (s *EventSub) readFuture(ctx context.Context, sink chan<- *Event) error {
	for {
		select {
		case <-s.closer.Closed():
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		select {
		case log := <-s.watchLogs:
			if err := s.processLog(ctx, log, sink); err != nil {
				return err
			}
		case err := <-s.watchSub.Err():
			err = cherrors.CheckIsChainNotReachableError(err)
			return err
		case <-ctx.Done():
			return ctx.Err()
		case <-s.closer.Closed():
			return nil
		}
	}
}

func (s *EventSub) processLog(ctx context.Context, log types.Log, sink chan<- *Event) error {
	event := s.eFact()
	if err := s.contract.UnpackLog(event.Data, event.Name, log); err != nil {
		return errors.WithStack(err)
	}
	event.Log = log

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.closer.Closed():
		return nil
	default:
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.closer.Closed():
		return nil
	case sink <- event:
		return nil
	}
}

// Close closes the sub and frees associated resources.
// Can be called more than once. Is thread safe.
func (s *EventSub) Close() {
	if err := s.closer.Close(); err != nil && !pkgsync.IsAlreadyClosedError(err) {
		log.WithError(err).Error("could not close EventSub")
	}
}
