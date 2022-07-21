// Copyright 2020 - See NOTICE file for copyright holders.
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
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
	"github.com/perun-network/perun-eth-backend/subscription"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/log"
)

const (
	secondaryWaitBlocks = 2
	adjEventBuffSize    = 10
	adjHeaderBuffSize   = 10
)

// ensureConcluded ensures that conclude or concludeFinal (for non-final and
// final states, resp.) is called on the adjudicator.
// - a subscription on Concluded events is established
// - it searches for a past concluded event by calling `isConcluded`
//   - if found, channel is already concluded and success is returned
//   - if none found, conclude/concludeFinal is called on the adjudicator
// - it waits for a Concluded event from the blockchain.
func (a *Adjudicator) ensureConcluded(ctx context.Context, req channel.AdjudicatorReq, subStates channel.StateMap) error {
	sub, err := subscription.Subscribe(ctx, a.ContractBackend, a.bound, updateEventType(req.Params.ID()), startBlockOffset, a.txFinalityDepth)
	if err != nil {
		return errors.WithMessage(err, "subscribing")
	}
	defer sub.Close()
	// Check whether it is already concluded.
	if concluded, err := a.isConcluded(ctx, sub); err != nil {
		return errors.WithMessage(err, "isConcluded")
	} else if concluded {
		return nil
	}

	events := make(chan *subscription.Event, adjEventBuffSize)
	subErr := make(chan error, 1)
	waitCtx, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()
		subErr <- sub.Read(ctx, events)
	}()

	concluded, err := a.waitConcludedSecondary(waitCtx, req, events)
	if err != nil {
		return errors.WithMessage(err, "waiting for secondary conclude")
	} else if concluded {
		return nil
	}

	// No conclude event found in the past, send transaction.
	err = a.conclude(ctx, req, subStates)
	if err != nil {
		return errors.WithMessage(err, "concluding")
	}

	// Wait for concluded event.
	for {
		select {
		case _e := <-events:
			e, ok := _e.Data.(*adjudicator.AdjudicatorChannelUpdate)
			if !ok {
				log.Panic("wrong event type")
			}
			if e.Phase == phaseConcluded {
				return nil
			}
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context cancelled")
		case err = <-subErr:
			if err != nil {
				return errors.WithMessage(err, "subscription error")
			}
			return errors.New("subscription closed")
		}
	}
}

// checkConcludedState checks whether the concluded state is equal to the
// expected state.
func (a *Adjudicator) checkConcludedState(
	ctx context.Context,
	req channel.AdjudicatorReq,
	subStates channel.StateMap,
) error {
	states := channel.MakeStateMap()
	states.Add(req.Tx.State)
	for _, v := range subStates {
		states.Add(v)
	}

	// Start event subscription for each channel.
	events := make(chan *subscription.Event, adjEventBuffSize)
	subErr := make(chan error, 1)
	for id := range states {
		sub, err := subscription.Subscribe(
			ctx,
			a.ContractBackend,
			a.bound,
			updateEventType(id),
			startBlockOffset,
			a.txFinalityDepth,
		)
		if err != nil {
			return errors.WithMessage(err, "subscribing")
		}
		defer sub.Close()
		go func() {
			subErr <- sub.Read(ctx, events)
		}()
	}

	// Wait for concluded events and check state version.
	validated := make(map[channel.ID]bool, len(states))
	for {
		select {
		case e := <-events:
			if adjEvent, ok := e.Data.(*adjudicator.AdjudicatorChannelUpdate); ok && adjEvent.Phase == phaseConcluded {
				id := adjEvent.ChannelID
				v := states[id].Version
				if adjEvent.Version != v {
					return errors.Errorf("wrong version: expected %v, got %v", v, adjEvent.Version)
				}
				validated[id] = true
				log.Debugf("validated: %v/%v", len(validated), len(states))
				if len(validated) == len(states) {
					return nil
				}
			}
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context cancelled")
		case err := <-subErr:
			if err != nil {
				return errors.WithMessage(err, "subscription error")
			}
			return errors.New("subscription closed")
		}
	}
}

func (a *Adjudicator) waitConcludedSecondary(ctx context.Context, req channel.AdjudicatorReq, events chan *subscription.Event) (concluded bool, err error) {
	// In final Register calls, as the non-initiator, we optimistically wait for
	// the other party to send the transaction first for
	// `secondaryWaitBlocks + TxFinalityDepth` many blocks.
	if req.Tx.IsFinal && req.Secondary {
		waitBlocks := secondaryWaitBlocks + int(a.txFinalityDepth)
		return waitConcludedForNBlocks(ctx, a, events, waitBlocks)
	}
	return false, nil
}

func (a *Adjudicator) conclude(ctx context.Context, req channel.AdjudicatorReq, subStates channel.StateMap) error {
	// If the on-chain state resulted from forced execution, we do not have a fully-signed state and cannot call concludeFinal.
	forceExecuted, err := a.isForceExecuted(ctx, req.Params.ID())
	if err != nil {
		return errors.WithMessage(err, "checking force execution")
	}
	if req.Tx.IsFinal && !forceExecuted {
		err = errors.WithMessage(a.callConcludeFinal(ctx, req), "calling concludeFinal")
	} else {
		err = errors.WithMessage(a.callConclude(ctx, req, subStates), "calling conclude")
	}
	if IsErrTxFailed(err) {
		a.log.WithError(err).Warn("Calling conclude(Final) failed, waiting for event anyways...")
	} else if err != nil {
		return err
	}
	return nil
}

// isConcluded returns whether a channel is already concluded.
func (a *Adjudicator) isConcluded(ctx context.Context, sub *subscription.ResistantEventSub) (bool, error) {
	events := make(chan *subscription.Event, adjEventBuffSize)
	subErr := make(chan error, 1)
	// Write the events into events.
	go func() {
		defer close(events)
		subErr <- sub.ReadPast(ctx, events)
	}()
	// Read all events and check for concluded.
	for _e := range events {
		e, ok := _e.Data.(*adjudicator.AdjudicatorChannelUpdate)
		if !ok {
			log.Panic("wrong event type")
		}
		if e.Phase == phaseConcluded {
			return true, nil
		}
	}
	return false, errors.WithMessage(<-subErr, "reading past events")
}

// isForceExecuted returns whether a channel is in the forced execution phase.
func (a *Adjudicator) isForceExecuted(_ctx context.Context, c channel.ID) (bool, error) {
	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()
	sub, err := subscription.NewEventSub(ctx, a.ContractBackend, a.bound, updateEventType(c), startBlockOffset)
	if err != nil {
		return false, errors.WithMessage(err, "subscribing")
	}
	defer sub.Close()
	events := make(chan *subscription.Event, adjEventBuffSize)
	subErr := make(chan error, 1)
	// Write the events into events.
	go func() {
		defer close(events)
		subErr <- sub.ReadPast(ctx, events)
	}()
	// Read all events and check for force execution.
	var lastEvent *subscription.Event
	for _e := range events {
		lastEvent = _e
	}
	if lastEvent != nil {
		e, ok := lastEvent.Data.(*adjudicator.AdjudicatorChannelUpdate)
		if !ok {
			log.Panic("wrong event type")
		}
		if e.Phase == phaseForceExec {
			return true, nil
		}
	}
	return false, errors.WithMessage(<-subErr, "reading past events")
}

func updateEventType(channelID [32]byte) subscription.EventFactory {
	return func() *subscription.Event {
		return &subscription.Event{
			Name: bindings.Events.AdjChannelUpdate,
			Data: new(adjudicator.AdjudicatorChannelUpdate),
			// In the best case we could already filter for 'Concluded' phase only here.
			Filter: [][]interface{}{{channelID}},
		}
	}
}

// waitConcludedForNBlocks waits for up to numBlocks blocks for a Concluded
// event on the concluded channel. If an event is emitted, true is returned.
// Otherwise, if numBlocks blocks have passed, false is returned.
//
// cr is the ChainReader used for setting up a block header subscription. sub is
// the Concluded event subscription instance.
func waitConcludedForNBlocks(ctx context.Context,
	cr ethereum.ChainReader,
	concluded chan *subscription.Event,
	numBlocks int,
) (bool, error) {
	h := make(chan *types.Header, adjHeaderBuffSize)
	hsub, err := cr.SubscribeNewHead(ctx, h)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return false, errors.WithMessage(err, "subscribing to new blocks")
	}
	defer hsub.Unsubscribe()
	for i := 0; i < numBlocks; i++ {
		select {
		case <-h: // do nothing, wait another block
		case _e := <-concluded: // other participant performed transaction
			e, ok := _e.Data.(*adjudicator.AdjudicatorChannelUpdate)
			if !ok {
				log.Panic("wrong event type")
			}
			if e.Phase == phaseConcluded {
				return true, nil
			}
		case <-ctx.Done():
			return false, errors.Wrap(ctx.Err(), "context cancelled")
		case err = <-hsub.Err():
			err = cherrors.CheckIsChainNotReachableError(err)
			return false, errors.WithMessage(err, "header subscription error")
		}
	}
	return false, nil
}
