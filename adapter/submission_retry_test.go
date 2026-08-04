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

//nolint:testpackage // white-box: exercises unexported retry/classify helpers
package adapter

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// Mempool/nonce rejections describe the tx pool at one instant, not the call.
// Classifying them as deterministic makes callers abandon work that would have
// succeeded — a pool-funded swap gets rejected because an unrelated transaction
// from the same operator key was still in flight.
func TestClassifyTransientSubmissionErrorsAreRetriable(t *testing.T) {
	for _, msg := range []string{
		"transaction underpriced",
		"replacement transaction underpriced",
		"nonce too low",
		"already known",
		"txpool is full",
		"max fee per gas less than block base fee",
		// Node wording varies in case; classification must not.
		"Transaction Underpriced",
	} {
		t.Run(msg, func(t *testing.T) {
			err := classifyEthError(errors.New(msg))
			require.ErrorIs(t, err, ErrRetriable,
				"%q is a transient tx-pool condition; a rebuilt tx succeeds", msg)
			require.NotErrorIs(t, err, ErrDeterministic)
		})
	}
}

// A genuine input/state problem must stay deterministic — retrying it would
// just burn the settlement window.
func TestClassifyRealFailuresStayNonRetriable(t *testing.T) {
	require.ErrorIs(t, classifyEthError(errors.New("execution reverted: Invalid beneficiary")), ErrContractRevert)
	require.ErrorIs(t, classifyEthError(errors.New("invalid opcode")), ErrDeterministic)
}

// The live regression: the hub funded a CKB->ETH swap, the node answered
// "transaction underpriced", and the adapter reported it as deterministic, so
// the hub released the reservation and rejected the proposal. Submission must
// instead be retried with fresh opts until the node accepts it.
func TestFundChannelRetriesTransientSubmissionRejection(t *testing.T) {
	m := &mockContract{
		withdrawable:      big.NewInt(100),
		locked:            big.NewInt(0),
		fundTx:            fakeTx([]byte{0x26, 0x0b, 0x60, 0x86}, big.NewInt(0)),
		fundTransientErrs: 2, // rejected twice, accepted on the third attempt
	}

	optsBuilt := 0
	a := newTestAdapter(m, func(context.Context) (*bind.TransactOpts, error) {
		optsBuilt++
		return &bind.TransactOpts{}, nil
	}, func(_ context.Context, _ *types.Transaction) (*types.Receipt, error) {
		return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
	})

	require.NoError(t, a.FundChannel(context.Background(), [32]byte{1}, big.NewInt(42)))
	require.Equal(t, 3, m.fundCalls, "should retry until the node accepts the tx")
	// Each attempt must rebuild the opts: reusing them keeps the stale nonce and
	// gas price, which is what produced "underpriced" in the first place.
	require.Equal(t, 3, optsBuilt, "every attempt must rebuild TransactOpts")
}

// Retrying must not paper over a real rejection.
func TestFundChannelDoesNotRetryDeterministicRejection(t *testing.T) {
	m := &mockContract{
		withdrawable: big.NewInt(100),
		locked:       big.NewInt(0),
		fundErr:      errors.New("invalid opcode"),
	}
	a := newTestAdapter(m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, nil)

	err := a.FundChannel(context.Background(), [32]byte{1}, big.NewInt(42))
	require.ErrorIs(t, err, ErrDeterministic)
	require.Equal(t, 1, m.fundCalls, "a deterministic rejection must not be retried")
}

// A caller's deadline must bound the retry loop, and the reported error must say
// what we were retrying on rather than only that time ran out.
func TestFundChannelRetryHonoursContextDeadline(t *testing.T) {
	m := &mockContract{
		withdrawable:      big.NewInt(100),
		locked:            big.NewInt(0),
		fundTx:            fakeTx([]byte{0x01}, big.NewInt(0)),
		fundTransientErrs: 1 << 30, // never accepts
	}
	a := newTestAdapter(m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 30*1e6) // 30ms
	defer cancel()

	err := a.FundChannel(ctx, [32]byte{1}, big.NewInt(42))
	require.ErrorIs(t, err, ErrRetriable)
	require.Contains(t, err.Error(), "underpriced",
		"the error must name the condition we kept retrying on")
}
