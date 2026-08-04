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
	"errors"
	"fmt"
	"strings"

	ch "github.com/perun-network/perun-eth-backend/channel"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
)

var (
	// ErrRetriable is a transient transport/RPC failure.
	ErrRetriable = errors.New("eth adapter: retriable error")
	// ErrDeterministic is a permanent input/state mismatch.
	ErrDeterministic = errors.New("eth adapter: deterministic error")
	// ErrContractRevert is an on-chain revert condition.
	ErrContractRevert = errors.New("eth adapter: contract revert")
)

func classifyEthError(err error) error {
	if err == nil {
		return nil
	}
	if isRPCTimeout(err) || isConnectionRefused(err) || cherrors.IsChainNotReachableError(err) ||
		isTransientSubmissionError(err) {
		return fmt.Errorf("%w: %v", ErrRetriable, err)
	}
	if isRevertError(err) || ch.IsErrTxFailed(err) {
		return fmt.Errorf("%w: %v", ErrContractRevert, err)
	}
	return fmt.Errorf("%w: %v", ErrDeterministic, err)
}

// isTransientSubmissionError reports whether the node rejected the transaction
// for a mempool/nonce condition rather than a permanent input or state problem.
//
// These describe the state of the tx pool at one instant, not the call: the same
// transaction, rebuilt with a fresh nonce and gas price, succeeds. Classifying
// them as deterministic makes a caller abandon work that would have gone through
// — which is how a pool-funded swap ends up rejected because an unrelated
// transaction from the same operator key was still in flight.
func isTransientSubmissionError(err error) bool {
	msg := strings.ToLower(err.Error())
	for _, m := range []string{
		// Gas price under the pool's minimum, or under the +10% bump a
		// replacement needs. Both clear once the pending tx lands.
		"underpriced",
		// The nonce was consumed between building and sending the tx.
		"nonce too low",
		"nonce is too low",
		// Same tx already queued/mined — a rebuild resolves it.
		"already known",
		"known transaction",
		"already imported",
		// Pool is full/busy right now.
		"txpool is full",
		"transaction pool is full",
		// EIP-1559 fee cap below the current base fee.
		"fee cap less than block base fee",
		"max fee per gas less than block base fee",
		// Explicit "retry later" from the node.
		"replacement transaction",
		"transaction pool limit",
	} {
		if strings.Contains(msg, m) {
			return true
		}
	}
	return false
}

func isRPCTimeout(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return true
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "timeout") || strings.Contains(msg, "deadline exceeded")
}

func isConnectionRefused(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "connection refused") || strings.Contains(msg, "connection reset") || strings.Contains(msg, "eof")
}

func isRevertError(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "revert") || strings.Contains(msg, "execution reverted")
}
