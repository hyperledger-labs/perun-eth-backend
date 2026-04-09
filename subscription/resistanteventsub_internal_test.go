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

package subscription

import (
	"context"
	stderrors "errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestResistantEventSub_ProcessHeadKeepsEventOnCanonicalityError(t *testing.T) {
	txHash := common.HexToHash("0x1")
	canonicalHeader := &types.Header{Number: big.NewInt(1)}
	reader := &canonicalityReader{
		headers: []*types.Header{
			nil,
			canonicalHeader,
		},
		errs: []error{
			stderrors.New("temporary rpc error"),
			nil,
		},
	}

	sub := &ResistantEventSub{
		cr:            reader,
		finalityDepth: big.NewInt(2),
		lastBlockNum:  big.NewInt(1),
		events: map[common.Hash]*Event{
			txHash: {
				Log: types.Log{
					TxHash:      txHash,
					BlockHash:   canonicalHeader.Hash(),
					BlockNumber: 1,
				},
			},
		},
	}

	sink := make(chan *Event, 1)
	sub.processHead(context.Background(), &types.Header{Number: big.NewInt(2)}, sink)
	require.Contains(t, sub.events, txHash)
	select {
	case <-sink:
		t.Fatal("event must stay buffered while canonicality checks fail")
	default:
	}

	sub.processHead(context.Background(), &types.Header{Number: big.NewInt(3)}, sink)
	require.NotContains(t, sub.events, txHash)
	require.NotNil(t, <-sink)
}

type canonicalityReader struct {
	headers []*types.Header
	errs    []error
	idx     int
}

func (r *canonicalityReader) BlockByHash(context.Context, common.Hash) (*types.Block, error) {
	panic("unexpected BlockByHash call")
}

func (r *canonicalityReader) BlockByNumber(context.Context, *big.Int) (*types.Block, error) {
	panic("unexpected BlockByNumber call")
}

func (r *canonicalityReader) HeaderByHash(context.Context, common.Hash) (*types.Header, error) {
	panic("unexpected HeaderByHash call")
}

func (r *canonicalityReader) HeaderByNumber(context.Context, *big.Int) (*types.Header, error) {
	header := r.headers[r.idx]
	err := r.errs[r.idx]
	if r.idx < len(r.headers)-1 {
		r.idx++
	}
	return header, err
}

func (r *canonicalityReader) TransactionCount(context.Context, common.Hash) (uint, error) {
	panic("unexpected TransactionCount call")
}

func (r *canonicalityReader) TransactionInBlock(context.Context, common.Hash, uint) (*types.Transaction, error) {
	panic("unexpected TransactionInBlock call")
}

func (r *canonicalityReader) SubscribeNewHead(context.Context, chan<- *types.Header) (ethereum.Subscription, error) {
	panic("unexpected SubscribeNewHead call")
}