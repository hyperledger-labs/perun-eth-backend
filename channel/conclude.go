// Copyright (c) 2020 Chair of Applied Cryptography, Technische Universität
// Darmstadt, Germany. All rights reserved. This file is part of go-perun. Use
// of this source code is governed by a MIT-style license that can be found in
// the LICENSE file.

package channel // import "perun.network/go-perun/backend/ethereum/channel"

import (
	"context"
	stderrors "errors"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"perun.network/go-perun/backend/ethereum/bindings/adjudicator"
	"perun.network/go-perun/log"

	"perun.network/go-perun/channel"
)

// Error that is returned if an event was not found in the past.
var errConcludedNotFound = stderrors.New("concluded event not found")

// conclude calls the conclude(or ConcludeFinal) function on the adjudicator.
// The call returns whether concluding was successful.
func (a *Adjudicator) conclude(ctx context.Context, params *channel.Params, tx channel.Transaction) error {
	// Listen for blockchain events.
	watchOpts, err := a.newWatchOpts(ctx)
	if err != nil {
		return errors.WithMessage(err, "creating watchOpts")
	}
	concluded := make(chan *adjudicator.AdjudicatorConcluded)
	sub, err := a.contract.WatchConcluded(watchOpts, concluded, [][32]byte{params.ID()})
	if err != nil {
		return errors.Wrap(err, "WatchConcluded failed")
	}
	defer sub.Unsubscribe()

	if err := a.filterConcludedConfirmations(ctx, params.ID()); err != errConcludedNotFound {
		// err might be nil, which is fine
		return errors.WithMessage(err, "filtering old Concluded events")
	}

	// No conclude event found in the past, send transaction.
	ethParams := channelParamsToEthParams(params)
	ethState := channelStateToEthState(tx.State)
	ethTX, err := func() (*types.Transaction, error) {
		if !a.mu.TryLockCtx(ctx) {
			return nil, errors.New("Could not acquire lock in time")
		}
		defer a.mu.Unlock()
		trans, err := a.newTransactor(ctx, big.NewInt(0), GasLimit)
		if err != nil {
			return nil, errors.WithMessage(err, "creating transactor")
		}
		a.log.WithField("IsFinal", tx.State.IsFinal).Debug("calling conclude")
		if tx.State.IsFinal {
			return a.contract.ConcludeFinal(trans, ethParams, ethState, tx.Sigs)
		}
		return a.contract.Conclude(trans, ethParams, ethState)
	}()
	if err != nil {
		return errors.WithMessage(err, "creating transaction")
	}

	if err := confirmTransaction(ctx, a.ContractBackend, ethTX); err != nil {
		log.Warnf("transaction failed: %v", err)
	} else {
		log.Debug("Transaction mined successfully")
	}

	select {
	case <-concluded:
		return nil
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "Waiting for final concluded event cancelled by context")
	case err = <-sub.Err():
		return errors.Wrap(err, "Error while waiting for events")
	}
}

func (a *Adjudicator) filterConcludedConfirmations(ctx context.Context, channelID channel.ID) error {
	// Filter
	filterOpts, err := a.newFilterOpts(ctx)
	if err != nil {
		return err
	}
	iter, err := a.contract.FilterConcluded(filterOpts, [][32]byte{channelID})
	if err != nil {
		return errors.WithStack(err)
	}
	if !iter.Next() {
		return errConcludedNotFound
	}
	// Event found, return nil
	return nil
}
