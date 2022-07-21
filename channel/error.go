// Copyright 2022 - See NOTICE file for copyright holders.
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
	stderrors "errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
	"github.com/pkg/errors"
)

// ErrTxFailed signals a failed, i.e., reverted, transaction.
var ErrTxFailed = stderrors.New("transaction failed")

// IsErrTxFailed returns whether the cause of the error was a failed transaction.
func IsErrTxFailed(err error) bool {
	return errors.Is(err, ErrTxFailed)
}

func errorReason(ctx context.Context, b *ContractBackend, tx *types.Transaction, blockNum *big.Int, acc accounts.Account) (string, error) {
	msg := ethereum.CallMsg{
		From:     acc.Address,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	res, err := b.CallContract(ctx, msg, blockNum)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return "", errors.WithMessage(err, "CallContract")
	}
	reason, err := abi.UnpackRevert(res)
	return reason, errors.Wrap(err, "unpacking revert reason")
}

// ErrInvalidContractCode signals invalid bytecode at given address, such as incorrect or no code.
var ErrInvalidContractCode = stderrors.New("invalid bytecode at address")

// IsErrInvalidContractCode returns whether the cause of the error was a invalid bytecode.
func IsErrInvalidContractCode(err error) bool {
	return errors.Is(err, ErrInvalidContractCode)
}
