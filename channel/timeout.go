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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
)

// BlockTimeout is a timeout on an Ethereum blockchain. A ChainReader is used to
// wait for the timeout to pass.
//
// This is much better than a channel.TimeTimeout because the local clock might
// not match the blockchain's timestamp at the point in time when the timeout
// has passed locally.
type BlockTimeout struct {
	Time uint64
	cr   ethereum.ChainReader
}

// NewBlockTimeout creates a new BlockTimeout bound to the provided ChainReader
// and ts as the absolute block.timestamp timeout.
func NewBlockTimeout(cr ethereum.ChainReader, ts uint64) *BlockTimeout {
	return &BlockTimeout{
		Time: ts,
		cr:   cr,
	}
}

// NewBlockTimeoutDuration creates a new BlockTimeout that elapses duration
// blockchain seconds added to the current block timestamp in the future.
func NewBlockTimeoutDuration(ctx context.Context, cr ethereum.ChainReader, duration uint64) (*BlockTimeout, error) {
	h, err := cr.HeaderByNumber(ctx, nil)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, errors.WithMessage(err, "getting latest header")
	}
	return NewBlockTimeout(cr, h.Time+duration), nil
}

// IsElapsed reads the timestamp from the current blockchain header to check
// whether the timeout has passed yet.
func (t *BlockTimeout) IsElapsed(ctx context.Context) bool {
	header, err := t.cr.HeaderByNumber(ctx, nil)
	if err != nil {
		// If there's an error, just return false here. A later Wait on the
		// BlockTimeout will expose the error to the caller.
		return false
	}

	return header.Time >= t.Time
}

// Wait subscribes to new blocks until the timeout is reached.
// It returns the context error if it is canceled before the timeout is reached.
func (t *BlockTimeout) Wait(ctx context.Context) error {
	headers := make(chan *types.Header)
	sub, err := t.cr.SubscribeNewHead(ctx, headers)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return errors.WithMessage(err, "subscribing to new heads")
	}
	defer sub.Unsubscribe()

	for {
		select {
		case header := <-headers:
			if header.Time >= t.Time {
				return nil
			}
		case err := <-sub.Err():
			if err != nil {
				err = cherrors.CheckIsChainNotReachableError(err)
				return errors.WithMessage(err, "sub done")
			}
			// make sure we return a non-nil error if the timeout hasn't passed yet
			return errors.New("sub done before timeout")
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context done")
		}
	}
}

// String returns a string stating the block timeout as a unix timestamp.
func (t *BlockTimeout) String() string {
	return fmt.Sprintf("<Block timeout: %d>", t.Time)
}
