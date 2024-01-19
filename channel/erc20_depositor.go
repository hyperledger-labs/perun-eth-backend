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
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"perun.network/go-perun/log"

	"github.com/perun-network/perun-eth-backend/bindings/assetholdererc20"
	"github.com/perun-network/perun-eth-backend/bindings/peruntoken"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
)

// ERC20Depositor deposits tokens into the `AssetHolderERC20` contract.
// It is bound to a token but can be reused to deposit multiple times.
type ERC20Depositor struct {
	Token common.Address
}

// ERC20DepositorTXGasLimit is the limit of Gas that an `ERC20Depositor` will
// spend per transaction when depositing funds.
// An `IncreaseAllowance` uses ~45kGas and a `Deposit` call ~84kGas on average.
const ERC20DepositorTXGasLimit = 0

// Return value of ERC20Depositor.NumTx.
const erc20DepositorNumTx = 2

// Keep track of the increase allowance and deposit processes.
var depositLocksMtx sync.Mutex
var depositLocks = make(map[string]*sync.Mutex)

// DepositResult is created to keep track of the returned values.
type DepositResult struct {
	Transactions types.Transactions
	Error        error
}

// NewERC20Depositor creates a new ERC20Depositor.
func NewERC20Depositor(token common.Address) *ERC20Depositor {
	return &ERC20Depositor{Token: token}
}

// Deposit approves the value to be swapped and calls DepositOnly.
func (d *ERC20Depositor) Deposit(ctx context.Context, req DepositReq) (types.Transactions, error) {
	lockKey := lockKey(req.Account.Address, req.Asset.EthAddress())
	lock := handleLock(lockKey)

	// Bind an `ERC20` instance.
	token, err := peruntoken.NewPeruntoken(d.Token, req.CB)
	if err != nil {
		return nil, errors.Wrapf(err, "binding ERC20 contract at: %x", d.Token)
	}
	callOpts := bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	// variables for the return value.
	var depResult DepositResult
	txApproval, approvalReceived, errApproval := Approve(ctx, lock, req, token, callOpts)
	if errApproval != nil {
		depResult.Error = errApproval
	}
	if approvalReceived {
		txDeposit, err := d.DepositOnly(ctx, req)
		depResult.Transactions = []*types.Transaction{txApproval, txDeposit}
		depResult.Error = errors.WithMessage(err, "AssetHolderERC20 depositing")
	} else {
		depResult.Error = errors.WithMessage(errApproval, "PerunToken was not approved")
	}
	return depResult.Transactions, depResult.Error
}

// DepositOnly deposits ERC20 tokens into the ERC20 AssetHolder specified at the
// requests asset address.
func (d *ERC20Depositor) DepositOnly(ctx context.Context, req DepositReq) (*types.Transaction, error) {
	// Bind a `AssetHolderERC20` instance.
	assetholder, err := assetholdererc20.NewAssetholdererc20(req.Asset.EthAddress(), req.CB)
	if err != nil {
		return nil, errors.Wrapf(err, "binding AssetHolderERC20 contract at: %x", req.Asset)
	}
	// Deposit.
	opts, err := req.CB.NewTransactor(ctx, ERC20DepositorTXGasLimit, req.Account)
	if err != nil {
		return nil, errors.WithMessagef(err, "creating transactor for asset: %x", req.Asset)
	}

	tx2, err := assetholder.Deposit(opts, req.FundingID, req.Balance)
	return tx2, err
}

// NumTX returns 2 since it does IncreaseAllowance and Deposit.
func (*ERC20Depositor) NumTX() uint32 {
	return erc20DepositorNumTx
}

// Create key from account address and asset to only lock the process when hub deposits the same asset at the same time.
func lockKey(account common.Address, asset common.Address) string {
	return fmt.Sprintf("%s-%s", account.Hex(), asset.Hex())
}

// Retrieves Lock for specific key.
func handleLock(lockKey string) *sync.Mutex {
	depositLocksMtx.Lock()
	defer depositLocksMtx.Unlock()

	if lock, exists := depositLocks[lockKey]; exists {
		return lock
	}

	lock := &sync.Mutex{}
	depositLocks[lockKey] = lock
	return lock
}

// Approve locks the lock argument and Approves the requested balance + the current allowance of the requested account.
func Approve(ctx context.Context, lock *sync.Mutex, req DepositReq, token *peruntoken.Peruntoken, callOpts bind.CallOpts) (*types.Transaction, bool, error) {
	lock.Lock()
	defer lock.Unlock()
	allowance, err := token.Allowance(&callOpts, req.Account.Address, req.Asset.EthAddress())
	if err != nil {
		return nil, false, errors.WithMessagef(err, "could not get Allowance for asset: %x", req.Asset)
	}
	result := new(big.Int).Add(req.Balance, allowance)

	// Increase the allowance.
	opts, err := req.CB.NewTransactor(ctx, ERC20DepositorTXGasLimit, req.Account)
	if err != nil {
		return nil, false, errors.WithMessagef(err, "creating transactor for asset: %x", req.Asset)
	}
	// Create a channel for receiving PeruntokenApproval events
	eventSink := make(chan *peruntoken.PeruntokenApproval)

	// Create a channel for receiving the Approval event
	eventReceived := make(chan bool)

	// Watch for Approval events and send them to the eventSink
	subscription, err := token.WatchApproval(&bind.WatchOpts{Start: nil, Context: ctx}, eventSink, []common.Address{req.Account.Address}, []common.Address{req.Asset.EthAddress()})
	if err != nil {
		return nil, false, errors.WithMessagef(err, "Cannot listen for event")
	}
	tx, err := token.Approve(opts, req.Asset.EthAddress(), result)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, false, errors.WithMessagef(err, "increasing allowance for asset: %x", req.Asset)
	}
	var approvalReceived bool
	go func() {
		select {
		case event := <-eventSink:
			log.Printf("Received Approval event: Owner: %s, Spender: %s, Value: %s\n", event.Owner.Hex(), event.Spender.Hex(), event.Value.String())
			eventReceived <- true
		case err := <-subscription.Err():
			log.Println("Subscription error:", err)
		}
	}()
	approvalReceived = <-eventReceived
	return tx, approvalReceived, nil
}
