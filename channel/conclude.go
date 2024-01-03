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
	"fmt"

	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	"github.com/perun-network/perun-eth-backend/subscription"
	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/log"
)

const (
	adjEventBuffSize = 10
)

// ensureConcluded ensures that conclude or concludeFinal (for non-final and
// final states, resp.) is called on the adjudicator.
// - a subscription on Concluded events is established
// - it searches for a past concluded event by calling `isConcluded`
//   - if found, channel is already concluded and success is returned
//   - if none found, conclude/concludeFinal is called on the adjudicator
//
// - it waits for a Concluded event from the blockchain.
func (a *Adjudicator) ensureConcluded(ctx context.Context, req channel.AdjudicatorReq, subStates channel.StateMap) error {
	// Check whether it is already concluded.
	if concluded, err := a.isConcluded(ctx, req.Tx.ID); err != nil {
		return errors.WithMessage(err, "isConcluded")
	} else if concluded {
		return nil
	}

	// Wait until we can conclude.
	err := a.waitConcludable(ctx, req)
	if err != nil {
		return fmt.Errorf("waiting for concludability: %w", err)
	}

	// No conclude event found in the past, send transaction.
	err = a.conclude(ctx, req, subStates)
	if err != nil {
		return errors.WithMessage(err, "concluding")
	}

	// Wait for concluded event.
	sub, events, subErr, err := a.createEventSub(ctx, req.Tx.ID, false)
	if err != nil {
		return errors.WithMessage(err, "subscribing")
	}
	defer sub.Close()
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
func (a *Adjudicator) isConcluded(ctx context.Context, ch channel.ID) (bool, error) {
	sub, events, subErr, err := a.createEventSub(ctx, ch, true)
	if err != nil {
		return false, errors.WithMessage(err, "subscribing")
	}
	defer sub.Close()

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

func (a *Adjudicator) createEventSub(
	ctx context.Context,
	ch channel.ID,
	past bool,
) (
	*subscription.ResistantEventSub,
	<-chan *subscription.Event,
	<-chan error,
	error,
) {
	sub, err := subscription.Subscribe(
		ctx,
		a.ContractBackend,
		a.bound,
		updateEventType(ch),
		startBlockOffset,
		a.txFinalityDepth,
	)
	if err != nil {
		return nil, nil, nil, errors.WithMessage(err, "subscribing")
	}

	events := make(chan *subscription.Event, adjEventBuffSize)
	subErr := make(chan error, 1)
	go func() {
		defer close(events)
		if past {
			subErr <- sub.ReadPast(ctx, events)
		} else {
			subErr <- sub.Read(ctx, events)
		}
	}()
	return sub, events, subErr, nil
}

// waitConcludable waits until the specified channel is concludable.
func (a *Adjudicator) waitConcludable(ctx context.Context, req channel.AdjudicatorReq) error {
	// If final, we can conclude immediately.
	if req.Tx.IsFinal {
		return nil
	}

	sub, events, subErr, err := a.createEventSub(ctx, req.Tx.ID, true)
	if err != nil {
		return errors.WithMessage(err, "subscribing")
	}
	defer sub.Close()

	// Process events.
	var up *adjudicator.AdjudicatorChannelUpdate
	for e := range events {
		var ok bool
		up, ok = e.Data.(*adjudicator.AdjudicatorChannelUpdate)
		if !ok {
			log.Panic("wrong event type")
		}
		if up.Phase == phaseDispute {
			// Wait until channel is concludable. If we have an app channel, we need
			// to wait two phases.
			t := up.Timeout
			if !channel.IsNoApp(req.Params.App) {
				t += req.Params.ChallengeDuration
			}
			timeout := NewBlockTimeout(a.ContractInterface, t)
			err := timeout.Wait(ctx)
			if err != nil {
				return err
			}
		}
	}
	return errors.WithMessage(<-subErr, "reading past events")
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
