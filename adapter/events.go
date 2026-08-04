// Copyright 2026 - See NOTICE file for copyright holders.
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

package adapter

import (
	"context"
	"math/big"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/perun-network/perun-eth-backend/bindings/liquiditypool"
)

const eventStartBlockOffset uint64 = 100

type liquidityPoolEventParser interface {
	ParseChannelFunded(log types.Log) (*liquiditypool.LiquidityPoolChannelFunded, error)
	ParseChannelSettled(log types.Log) (*liquiditypool.LiquidityPoolChannelSettled, error)
	ParseOperatorUpdated(log types.Log) (*liquiditypool.LiquidityPoolOperatorUpdated, error)
}

type managedSubscription struct {
	errCh  chan error
	cancel context.CancelFunc
	once   sync.Once
}

func (s *managedSubscription) Err() <-chan error {
	return s.errCh
}

func (s *managedSubscription) Unsubscribe() {
	s.once.Do(s.cancel)
}

func (a *LiquidityPoolAdapter) parser() (liquidityPoolEventParser, bool) {
	p, ok := a.contract.(liquidityPoolEventParser)
	return p, ok
}

// SubscribeChannelFunded subscribes typed ChannelFunded events with reconnect/backoff.
func (a *LiquidityPoolAdapter) SubscribeChannelFunded(ctx context.Context, ch chan<- ChannelFundedEvent) (Subscription, error) {
	p, ok := a.parser()
	if !ok {
		return nil, classifyEthError(bind.ErrNoCode)
	}
	return a.subscribeLoop(ctx, func(log types.Log) error {
		e, err := p.ParseChannelFunded(log)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- ChannelFundedEvent{
			ChannelID: e.ChannelId,
			Principal: e.Principal,
			Operator:  e.Operator,
			BlockNum:  e.Raw.BlockNumber,
		}:
			return nil
		}
	}, liquiditypool.LiquidityPoolChannelFundedTopic, func() { close(ch) })
}

// SubscribeChannelSettled subscribes typed ChannelSettled events with reconnect/backoff.
func (a *LiquidityPoolAdapter) SubscribeChannelSettled(ctx context.Context, ch chan<- ChannelSettledEvent) (Subscription, error) {
	p, ok := a.parser()
	if !ok {
		return nil, classifyEthError(bind.ErrNoCode)
	}
	return a.subscribeLoop(ctx, func(log types.Log) error {
		e, err := p.ParseChannelSettled(log)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- ChannelSettledEvent{
			ChannelID:     e.ChannelId,
			Principal:     e.Principal,
			TotalReturned: e.TotalReturned,
			FeeGain:       e.FeeGain,
			BlockNum:      e.Raw.BlockNumber,
		}:
			return nil
		}
	}, liquiditypool.LiquidityPoolChannelSettledTopic, func() { close(ch) })
}

// SubscribeOperatorUpdated subscribes typed OperatorUpdated events with reconnect/backoff.
func (a *LiquidityPoolAdapter) SubscribeOperatorUpdated(ctx context.Context, ch chan<- OperatorUpdatedEvent) (Subscription, error) {
	p, ok := a.parser()
	if !ok {
		return nil, classifyEthError(bind.ErrNoCode)
	}
	return a.subscribeLoop(ctx, func(log types.Log) error {
		e, err := p.ParseOperatorUpdated(log)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- OperatorUpdatedEvent{
			Previous: e.PreviousOperator,
			New:      e.NewOperator,
		}:
			return nil
		}
	}, liquiditypool.LiquidityPoolOperatorUpdatedTopic, func() { close(ch) })
}

// subscribeLoop is the reconnect state machine backing every typed
// subscription: it runs a historical scan (scanAndDeliver), closes the gap to
// the live feed, pumps live logs (pumpLive), and reconnects with backoff. The
// branching lives in those helpers; what remains here is the reconnect loop
// wiring them together.
//
//nolint:funlen,gocognit // reconnect loop; scan/pump/backoff are extracted
func (a *LiquidityPoolAdapter) subscribeLoop(
	parent context.Context,
	onLog func(log types.Log) error,
	topic common.Hash,
	onClose func(),
) (Subscription, error) {
	if a.backend == nil {
		return nil, classifyEthError(bind.ErrNoCode)
	}

	ctx, cancel := context.WithCancel(parent)
	sub := &managedSubscription{errCh: make(chan error, 1), cancel: cancel}

	start, err := a.startBlock(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	go func() {
		defer close(sub.errCh)
		defer onClose()

		nextBlock := start
		delay := a.retryInitial
		maxDelay := a.retryMax
		if delay <= 0 {
			delay = 250 * time.Millisecond
		}
		if maxDelay < delay {
			maxDelay = 5 * time.Second
		}

		for {
			if ctx.Err() != nil {
				return
			}

			// Historical scan. On a scan error, back off and retry the loop.
			if scanErr := a.scanAndDeliver(ctx, sub.errCh, topic, onLog, &nextBlock); scanErr != nil {
				if reportAndBackoff(sub.errCh, scanErr, &delay, maxDelay) {
					return
				}
				continue
			}

			query := ethereum.FilterQuery{
				Addresses: []common.Address{a.poolAddress},
				Topics:    [][]common.Hash{{topic}},
				FromBlock: new(big.Int).SetUint64(nextBlock),
			}
			live := make(chan types.Log, 32)
			liveSub, err := a.backend.SubscribeFilterLogs(ctx, query, live)
			if err != nil {
				if reportAndBackoff(sub.errCh, err, &delay, maxDelay) {
					return
				}
				continue
			}
			delay = a.retryInitial
			if delay <= 0 {
				delay = 250 * time.Millisecond
			}

			// Close the gap between the historical scan and the live feed: an
			// event mined in that window would be delivered by neither. pumpLive
			// skips anything already covered (BlockNumber < nextBlock). A
			// catch-up scan error is best-effort — the live feed recovers it.
			if scanErr := a.scanAndDeliver(ctx, sub.errCh, topic, onLog, &nextBlock); scanErr != nil {
				sendErr(sub.errCh, classifyEthError(scanErr))
			}

			if a.pumpLive(ctx, sub.errCh, liveSub, live, onLog, &nextBlock) {
				return
			}
			time.Sleep(delay)
			delay = minDuration(delay*2, maxDelay)
		}
	}()

	return sub, nil
}

// pumpLive drains the live subscription, delivering new logs and advancing
// *nextBlock, until either the subscription drops (return false, so the caller
// reconnects) or the goroutine must stop (return true: context cancelled or the
// error channel is gone). It unsubscribes liveSub on every exit path.
func (a *LiquidityPoolAdapter) pumpLive(
	ctx context.Context,
	errCh chan<- error,
	liveSub ethereum.Subscription,
	live <-chan types.Log,
	onLog func(log types.Log) error,
	nextBlock *uint64,
) bool {
	for {
		select {
		case <-ctx.Done():
			liveSub.Unsubscribe()
			return true
		case err := <-liveSub.Err():
			liveSub.Unsubscribe()
			if ctx.Err() != nil {
				return true
			}
			if err != nil && !sendErr(errCh, classifyEthError(err)) {
				return true
			}
			return false
		case lg := <-live:
			if lg.Removed || lg.BlockNumber < *nextBlock {
				continue
			}
			if e := onLog(lg); e != nil {
				if ctx.Err() != nil {
					liveSub.Unsubscribe()
					return true
				}
				if !sendErr(errCh, classifyEthError(e)) {
					liveSub.Unsubscribe()
					return true
				}
			}
			*nextBlock = lg.BlockNumber + 1
		}
	}
}

// reportAndBackoff reports err on the error channel and sleeps for the current
// backoff, then doubles it (bounded by maxDelay). It returns true when the
// goroutine should stop because the error channel is gone.
func reportAndBackoff(errCh chan<- error, err error, delay *time.Duration, maxDelay time.Duration) bool {
	if !sendErr(errCh, classifyEthError(err)) {
		return true
	}
	time.Sleep(*delay)
	*delay = minDuration(*delay*2, maxDelay)
	return false
}

// scanAndDeliver runs one FilterLogs sweep from *nextBlock and delivers every
// not-yet-seen matching log via onLog, advancing *nextBlock past them. It backs
// both the initial historical scan and the post-subscribe gap catch-up, and
// returns the scan error (nil on success). Delivery errors are reported but do
// not abort the finite batch; cancellation is handled by the caller's loop.
func (a *LiquidityPoolAdapter) scanAndDeliver(
	ctx context.Context,
	errCh chan<- error,
	topic common.Hash,
	onLog func(log types.Log) error,
	nextBlock *uint64,
) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{a.poolAddress},
		Topics:    [][]common.Hash{{topic}},
		FromBlock: new(big.Int).SetUint64(*nextBlock),
	}
	logs, err := a.backend.FilterLogs(ctx, query)
	if err != nil {
		return err
	}
	for _, lg := range logs {
		if lg.Removed || lg.BlockNumber < *nextBlock {
			continue
		}
		if e := onLog(lg); e != nil {
			sendErr(errCh, classifyEthError(e))
		}
		*nextBlock = lg.BlockNumber + 1
	}
	return nil
}

func (a *LiquidityPoolAdapter) startBlock(ctx context.Context) (uint64, error) {
	head, err := a.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, classifyEthError(err)
	}
	h := head.Number.Uint64()
	if h <= eventStartBlockOffset {
		return 1, nil
	}
	return h - eventStartBlockOffset, nil
}

func minDuration(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}

// sendErr delivers err on the buffered error channel without blocking. It
// returns false when the channel is full (a prior error is still unread),
// signalling the subscription goroutine to stop rather than spin.
func sendErr(ch chan<- error, err error) bool {
	select {
	case ch <- err:
		return true
	default:
		return false
	}
}
