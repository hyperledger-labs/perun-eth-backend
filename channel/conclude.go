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

package channel

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	// concludeWaitSlack absorbs simulated-backend timing jitter between timeout
	// observation and the settle transaction reaching the chain.
	concludeWaitSlack = 12
)

// StateMap represents a channel state tree.
type StateMap map[channel.ID]*channel.State

type disputeInfo struct {
	Timeout           uint64
	ChallengeDuration uint64
	Version           uint64
	Phase             uint8
	StateHash         [32]byte
	HasApp            bool
}

// MakeStateMap creates a new StateMap object.
func MakeStateMap() StateMap {
	return make(map[channel.ID]*channel.State)
}

// Add adds the given states to the state map.
func (m StateMap) Add(states ...*channel.State) {
	for _, s := range states {
		m[s.ID] = s
	}
}

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

	// If the secondary flag is set, we wait for someone else to conclude.
	concluded, err := a.waitConcludedSecondary(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "waiting for secondary conclude")
	} else if concluded {
		return nil
	}

	// Wait until we can conclude.
	err = a.waitConcludable(ctx, req)
	if err != nil {
		return fmt.Errorf("waiting for concludability: %w", err)
	}

	// No conclude event found in the past, send transaction.
	err = a.conclude(ctx, req, subStates)
	if err != nil {
		return errors.WithMessage(err, "concluding")
	}
	// A concurrent conclude can make our own transaction revert. In that case,
	// re-check the current on-chain state before waiting for a fresh event.
	if concluded, err := a.isConcluded(ctx, req.Tx.ID); err != nil {
		return errors.WithMessage(err, "re-checking concluded state")
	} else if concluded {
		return nil
	}
	if concluded, err := a.hasConcludedEvent(ctx, req.Tx.ID); err != nil {
		return errors.WithMessage(err, "checking concluded event")
	} else if concluded {
		return nil
	}

	// Wait for concluded event.
	sub, events, subErr, err := a.createEventSub(ctx, req.Tx.ID, false, startBlockOffset)
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
	states := MakeStateMap()
	states.Add(req.Tx.State)
	for _, v := range subStates {
		states.Add(v)
	}

	validated := make(map[channel.ID]bool, len(states))
	validate := func() error {
		for id, state := range states {
			if validated[id] {
				continue
			}
			dispute, err := a.dispute(ctx, id)
			if err != nil {
				return errors.WithMessage(err, "querying dispute")
			}
			if dispute.Phase != phaseConcluded {
				continue
			}
			if dispute.Version != state.Version {
				return errors.Errorf("wrong version: expected %v, got %v", state.Version, dispute.Version)
			}
			validated[id] = true
		}
		return nil
	}
	if err := validate(); err != nil {
		return err
	}
	if len(validated) == len(states) {
		return nil
	}

	heads := make(chan *types.Header, adjHeaderBuffSize)
	hsub, err := a.SubscribeNewHead(ctx, heads)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return errors.WithMessage(err, "subscribing to new blocks")
	}
	defer hsub.Unsubscribe()

	for {
		select {
		case <-heads:
			if err := validate(); err != nil {
				return err
			}
			log.Debugf("validated: %v/%v", len(validated), len(states))
			if len(validated) == len(states) {
				return nil
			}
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context cancelled")
		case err := <-hsub.Err():
			err = cherrors.CheckIsChainNotReachableError(err)
			return errors.WithMessage(err, "header subscription error")
		}
	}
}

func (a *Adjudicator) waitConcludedSecondary(ctx context.Context, req channel.AdjudicatorReq) (concluded bool, err error) {
	// In final Register calls, as the non-initiator, we optimistically wait for
	// the other party to send the transaction first for
	// `secondaryWaitBlocks + TxFinalityDepth` many blocks.
	if req.Tx.IsFinal && req.Secondary {
		waitBlocks := secondaryWaitBlocks + int(a.txFinalityDepth)
		return waitConcludedForNBlocksPolling(ctx, a, req.Tx.ID, waitBlocks)
	}
	return false, nil
}

func (a *Adjudicator) conclude(ctx context.Context, req channel.AdjudicatorReq, subStates channel.StateMap) error {
	// If the on-chain state resulted from forced execution, we do not have a fully-signed state and cannot call concludeFinal.
	forceExecuted, err := a.isForceExecuted(ctx, req.Params.ID())
	if err != nil {
		return errors.WithMessage(err, "checking force execution")
	}
	dispute, err := a.dispute(ctx, req.Params.ID())
	if err != nil {
		return errors.WithMessage(err, "querying dispute")
	}
	recorded := hasRecordedDispute(dispute)
	// concludeFinal only works as the fast path before the channel entered the
	// regular dispute lifecycle. Once there is already an on-chain dispute, we
	// must use conclude after the timeout even for final states.
	if req.Tx.IsFinal && !forceExecuted && (!recorded || (dispute.Phase != phaseDispute && dispute.Phase != phaseForceExec)) {
		err = errors.WithMessage(a.callConcludeFinal(ctx, req), "calling concludeFinal")
	} else {
		err = errors.WithMessage(a.callConclude(ctx, req, subStates), "calling conclude")
	}
	if isBenignConcludeTxFailure(err) {
		a.log.WithError(err).Warn("Calling conclude(Final) failed, waiting for event anyways...")
	} else if err != nil {
		return err
	}
	return nil
}

func isBenignConcludeTxFailure(err error) bool {
	if !IsErrTxFailed(err) {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "concluded already") || strings.Contains(msg, "already concluded")
}

// isConcluded returns whether a channel is already concluded.
func (a *Adjudicator) isConcluded(ctx context.Context, ch channel.ID) (bool, error) {
	dispute, err := a.dispute(ctx, ch)
	if err != nil {
		return false, errors.WithMessage(err, "querying dispute")
	}
	return hasRecordedDispute(dispute) && dispute.Phase == phaseConcluded, nil
}

func (a *Adjudicator) hasConcludedEvent(ctx context.Context, ch channel.ID) (bool, error) {

	sub, err := subscription.NewEventSub(ctx, a.ContractBackend, a.bound, updateEventType(ch), startBlockOffset)
	if err != nil {
		return false, errors.WithMessage(err, "subscribing")
	}
	defer sub.Close()
	events := make(chan *subscription.Event, adjEventBuffSize)
	subErr := make(chan error, 1)
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

func (a *Adjudicator) createEventSub(
	ctx context.Context,
	ch channel.ID,
	past bool,
	pastBlocks uint64,
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
		pastBlocks,
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
	dispute, err := a.dispute(ctx, req.Tx.ID)
	if err != nil {
		return errors.WithMessage(err, "querying dispute")
	}
	recorded := hasRecordedDispute(dispute)

	// Final states can be concluded immediately if the channel is not already in
	// a dispute. Once a dispute exists, we must respect its timeout and fall back
	// to the regular conclude path afterwards.
	if req.Tx.IsFinal && (!recorded || (dispute.Phase != phaseDispute && dispute.Phase != phaseForceExec)) {
		return nil
	}
	if !recorded {
		return nil
	}

	switch dispute.Phase {
	case phaseDispute:
		t := dispute.Timeout
		if dispute.HasApp && !channel.IsNoApp(req.Params.App) {
			t += dispute.ChallengeDuration
		}
		t += concludeWaitSlack
		return NewBlockTimeout(a.ContractInterface, t).Wait(ctx)
	case phaseForceExec:
		return NewBlockTimeout(a.ContractInterface, dispute.Timeout+concludeWaitSlack).Wait(ctx)
	case phaseConcluded:
		return nil
	default:
		return nil
	}
}

// isForceExecuted returns whether a channel is in the forced execution phase.
func (a *Adjudicator) isForceExecuted(_ctx context.Context, c channel.ID) (bool, error) {
	dispute, err := a.dispute(_ctx, c)
	if err != nil {
		return false, errors.WithMessage(err, "querying dispute")
	}
	return hasRecordedDispute(dispute) && dispute.Phase == phaseForceExec, nil
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

func (a *Adjudicator) dispute(ctx context.Context, ch channel.ID) (disputeInfo, error) {
	return a.contract.Disputes(&bind.CallOpts{Context: ctx}, ch)
}

func hasRecordedDispute(dispute disputeInfo) bool {
	return dispute.Timeout != 0 || dispute.ChallengeDuration != 0 || dispute.StateHash != ([32]byte{})
}

func waitConcludedForNBlocksPolling(ctx context.Context, a *Adjudicator, ch channel.ID, numBlocks int) (bool, error) {
	h := make(chan *types.Header, adjHeaderBuffSize)
	hsub, err := a.SubscribeNewHead(ctx, h)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return false, errors.WithMessage(err, "subscribing to new blocks")
	}
	defer hsub.Unsubscribe()

	for i := 0; i < numBlocks; i++ {
		select {
		case <-h:
			concluded, err := a.isConcluded(ctx, ch)
			if err != nil {
				return false, errors.WithMessage(err, "checking concluded state")
			}
			if concluded {
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
