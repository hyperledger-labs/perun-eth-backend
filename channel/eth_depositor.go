// Copyright 2024 - See NOTICE file for copyright holders.
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

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/perun-network/perun-eth-backend/bindings/assetholdereth"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
)

// ETHDepositor deposits funds into the `AssetHolderETH` contract.
// It has no state and is therefore completely reusable.
type ETHDepositor struct {
	GasLimit uint64
}

// NewETHDepositor creates a new ETHDepositor.
func NewETHDepositor(gasLimit uint64) *ETHDepositor {
	return &ETHDepositor{
		GasLimit: gasLimit,
	}
}

// Deposit deposits ether into the ETH AssetHolder specified at the requests's
// asset address.
func (d *ETHDepositor) Deposit(ctx context.Context, req DepositReq) (types.Transactions, error) {
	// Bind an `AssetHolderETH` instance. Using `AssetHolder` is also possible
	// since we only use the interface functions here.
	contract, err := assetholdereth.NewAssetholdereth(req.Asset.EthAddress(), req.CB)
	if err != nil {
		return nil, errors.Wrapf(err, "binding AssetHolderETH contract at: %v", req.Asset)
	}
	opts, err := req.CB.NewTransactor(ctx, d.GasLimit, req.Account)
	if err != nil {
		return nil, errors.WithMessagef(err, "creating transactor for asset: %v", req.Asset)
	}
	opts.Value = req.Balance

	tx, err := contract.Deposit(opts, req.FundingID, req.Balance)
	err = cherrors.CheckIsChainNotReachableError(err)
	return []*types.Transaction{tx}, errors.WithMessage(err, "AssetHolderETH depositing")
}

// NumTX returns 1 since it only does Deposit.
func (*ETHDepositor) NumTX() uint32 {
	return 1
}
