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
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/perun-network/perun-eth-backend/bindings/liquiditypool"
	"github.com/perun-network/perun-eth-backend/channel"
)

// liquidityPoolContract mirrors the generated LiquidityPool binding 1:1, so its
// method count tracks the on-chain ABI rather than any design choice here.
//
//nolint:interfacebloat // intentionally mirrors the on-chain LiquidityPool ABI
type liquidityPoolContract interface {
	WithdrawableETH(opts *bind.CallOpts) (*big.Int, error)
	SharesOf(opts *bind.CallOpts, provider common.Address) (*big.Int, error)
	LockedByChannel(opts *bind.CallOpts, channelID [32]byte) (*big.Int, error)
	MinSettlementValue(opts *bind.CallOpts, channelID [32]byte) (*big.Int, error)
	OperatorBond(opts *bind.CallOpts) (*big.Int, error)
	FundChannel(opts *bind.TransactOpts, channelID [32]byte, amount *big.Int) (*types.Transaction, error)
	SettleChannel(opts *bind.TransactOpts, channelID [32]byte) (*types.Transaction, error)
	BondETH(opts *bind.TransactOpts) (*types.Transaction, error)
	DepositFor(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error)
	Operator(opts *bind.CallOpts) (common.Address, error)
	TotalAssets(opts *bind.CallOpts) (*big.Int, error)
	TotalLockedETH(opts *bind.CallOpts) (*big.Int, error)
	TotalShares(opts *bind.CallOpts) (*big.Int, error)
	PreviewWithdrawETH(opts *bind.CallOpts, sharesAmount *big.Int) (*big.Int, error)
}

type adapterBackend interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	NewTransactor(ctx context.Context, gasLimit uint64, acc accounts.Account) (*bind.TransactOpts, error)
	TxFinalityDepth() uint64
}

type (
	txOptsFactory func(ctx context.Context) (*bind.TransactOpts, error)
	txConfirmer   func(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)
)

// AdapterOption configures runtime behavior of the adapter.
type AdapterOption func(*adapterConfig)

type adapterConfig struct {
	retryInitial time.Duration
	retryMax     time.Duration
	finality     uint64
}

const (
	defaultRetryInitial = 250 * time.Millisecond
	defaultRetryMax     = 5 * time.Second
)

// WithRetryBackoff configures reconnect backoff bounds for event subscriptions.
func WithRetryBackoff(initial, maximum time.Duration) AdapterOption {
	return func(c *adapterConfig) {
		if initial > 0 {
			c.retryInitial = initial
		}
		if maximum > 0 {
			c.retryMax = maximum
		}
	}
}

// WithFinalityDepth overrides transaction finality depth for tx confirmation.
// A zero value keeps backend default finality depth.
func WithFinalityDepth(depth uint64) AdapterOption {
	return func(c *adapterConfig) {
		if depth > 0 {
			c.finality = depth
		}
	}
}

// LiquidityPoolAdapter is the ETH execution adapter between Hub/Websocket and LiquidityPool.
type LiquidityPoolAdapter struct {
	backend      adapterBackend
	contract     liquidityPoolContract
	poolAddress  common.Address
	chainID      *big.Int
	newTxOpts    txOptsFactory
	confirmTx    txConfirmer
	retryInitial time.Duration
	retryMax     time.Duration
	finality     uint64
}

// sendTx builds fresh TransactOpts and submits one transaction via send,
// retrying while the node rejects it for a transient mempool/nonce condition
// (see isTransientSubmissionError). Each attempt rebuilds the opts, so a retry
// picks up a fresh nonce and gas price — retrying with the same opts would just
// reproduce "underpriced".
//
// Only submission is retried. Once a transaction is accepted, confirmation and
// revert handling stay with the caller: resending a mined tx would execute the
// call twice.
//
// Backoff starts at retryInitial and doubles up to retryMax, bounded by ctx.
func (a *LiquidityPoolAdapter) sendTx(
	ctx context.Context,
	mutate func(*bind.TransactOpts),
	send func(*bind.TransactOpts) (*types.Transaction, error),
) (*types.Transaction, error) {
	delay := a.retryInitial
	if delay <= 0 {
		delay = defaultRetryInitial
	}
	maxDelay := a.retryMax
	if maxDelay <= 0 {
		maxDelay = defaultRetryMax
	}

	for {
		opts, err := a.newTxOpts(ctx)
		if err != nil {
			return nil, classifyEthError(err)
		}
		if mutate != nil {
			mutate(opts)
		}
		tx, err := send(opts)
		if err == nil {
			return tx, nil
		}
		classified := classifyEthError(err)
		if !errors.Is(classified, ErrRetriable) {
			return nil, classified
		}

		select {
		case <-ctx.Done():
			// Report what we were retrying on, not just that time ran out.
			return nil, fmt.Errorf("%w: %v (last submission error: %v)", ErrRetriable, ctx.Err(), classified)
		case <-time.After(delay):
		}
		if delay < maxDelay {
			if delay *= 2; delay > maxDelay {
				delay = maxDelay
			}
		}
	}
}

// NewLiquidityPoolAdapter creates a new adapter over a deployed LiquidityPool contract.
func NewLiquidityPoolAdapter(
	backend *channel.ContractBackend,
	poolAddress common.Address,
	txSender accounts.Account,
	gasLimit uint64,
	opts ...AdapterOption,
) (*LiquidityPoolAdapter, error) {
	contract, err := liquiditypool.NewLiquidityPool(poolAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("binding LiquidityPool: %w", err)
	}

	cfg := adapterConfig{
		retryInitial: defaultRetryInitial,
		retryMax:     defaultRetryMax,
		finality:     backend.TxFinalityDepth(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.retryInitial <= 0 {
		cfg.retryInitial = defaultRetryInitial
	}
	if cfg.retryMax < cfg.retryInitial {
		cfg.retryMax = cfg.retryInitial
	}
	if cfg.finality < 1 {
		cfg.finality = 1
	}

	var chainID *big.Int
	if cid := backend.ChainID().Int; cid != nil {
		chainID = new(big.Int).Set(cid)
	}

	return &LiquidityPoolAdapter{
		backend:     backend,
		contract:    contract,
		poolAddress: poolAddress,
		chainID:     chainID,
		newTxOpts: func(ctx context.Context) (*bind.TransactOpts, error) {
			return backend.NewTransactor(ctx, gasLimit, txSender)
		},
		confirmTx: func(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
			return confirmWithFinality(ctx, backend, tx, cfg.finality)
		},
		retryInitial: cfg.retryInitial,
		retryMax:     cfg.retryMax,
		finality:     cfg.finality,
	}, nil
}

func newTestAdapter(
	contract liquidityPoolContract,
	newTxOpts txOptsFactory,
	confirm txConfirmer,
) *LiquidityPoolAdapter {
	return &LiquidityPoolAdapter{
		contract:     contract,
		newTxOpts:    newTxOpts,
		confirmTx:    confirm,
		retryInitial: 5 * time.Millisecond,
		retryMax:     50 * time.Millisecond,
		finality:     1,
	}
}

func confirmWithFinality(ctx context.Context, backend adapterBackend, tx *types.Transaction, finalityDepth uint64) (*types.Receipt, error) {
	if finalityDepth < 1 {
		finalityDepth = 1
	}

	head, err := waitMinedHead(ctx, backend, tx)
	if err != nil {
		return nil, err
	}

	heads := make(chan *types.Header, 10)
	heads <- head
	hsub, err := backend.SubscribeNewHead(ctx, heads)
	if err != nil {
		return nil, err
	}
	defer hsub.Unsubscribe()

	for {
		select {
		case head := <-heads:
			receipt, err := backend.TransactionReceipt(ctx, tx.Hash())
			if err != nil || receipt == nil {
				continue
			}
			if isReceiptFinal(receipt, head, finalityDepth) {
				return receipt, nil
			}
		case err := <-hsub.Err():
			if err != nil {
				return nil, err
			}
			return nil, context.Canceled
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func waitMinedHead(ctx context.Context, backend adapterBackend, tx *types.Transaction) (*types.Header, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		receipt, err := backend.TransactionReceipt(ctx, tx.Hash())
		if err == nil && receipt != nil {
			return backend.HeaderByNumber(ctx, nil)
		}

		timer := time.NewTimer(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}

func isReceiptFinal(receipt *types.Receipt, head *types.Header, finalityDepth uint64) bool {
	depth := new(big.Int).SetUint64(finalityDepth)
	diff := new(big.Int).Sub(head.Number, receipt.BlockNumber)
	included := new(big.Int).Add(diff, big.NewInt(1))
	return included.Cmp(depth) >= 0
}

// FundChannel calls fundChannel(channelId, amount) on LiquidityPool.
func (a *LiquidityPoolAdapter) FundChannel(ctx context.Context, channelID [32]byte, amount *big.Int) error {
	free, err := a.contract.WithdrawableETH(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("%w: cannot read free liquidity: %v", ErrRetriable, err)
	}
	if amount.Cmp(free) > 0 {
		return fmt.Errorf("%w: amount %s > free liquidity %s", ErrDeterministic, amount, free)
	}

	existing, err := a.contract.LockedByChannel(&bind.CallOpts{Context: ctx}, channelID)
	if err != nil {
		return fmt.Errorf("%w: cannot read channel lock: %v", ErrRetriable, err)
	}
	if existing.Sign() > 0 {
		return fmt.Errorf("%w: channel already funded with %s", ErrDeterministic, existing)
	}

	// Mirror the contract's bond-coverage requirement so an under-bonded
	// operator fails deterministically instead of burning gas on a revert.
	bond, err := a.contract.OperatorBond(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("%w: cannot read operator bond: %v", ErrRetriable, err)
	}
	locked, err := a.contract.TotalLockedETH(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("%w: cannot read total locked: %v", ErrRetriable, err)
	}
	if required := new(big.Int).Add(locked, amount); bond.Cmp(required) < 0 {
		return fmt.Errorf("%w: bond %s < locked+amount %s - insufficient coverage", ErrDeterministic, bond, required)
	}

	tx, err := a.sendTx(ctx, nil, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return a.contract.FundChannel(opts, channelID, amount)
	})
	if err != nil {
		return err
	}
	receipt, err := a.confirmTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("%w: tx wait failed: %v", ErrRetriable, err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("%w: fundChannel reverted tx=%s", ErrContractRevert, tx.Hash())
	}
	return nil
}

// SettleChannel calls settleChannel(channelId) payable on LiquidityPool.
func (a *LiquidityPoolAdapter) SettleChannel(ctx context.Context, channelID [32]byte, returnAmount *big.Int) error {
	locked, err := a.contract.LockedByChannel(&bind.CallOpts{Context: ctx}, channelID)
	if err != nil {
		return fmt.Errorf("%w: cannot read channel lock: %v", ErrRetriable, err)
	}
	if locked.Sign() == 0 {
		return fmt.Errorf("%w: channel %x not found in pool", ErrDeterministic, channelID)
	}
	minValue, err := a.contract.MinSettlementValue(&bind.CallOpts{Context: ctx}, channelID)
	if err != nil {
		return fmt.Errorf("%w: cannot read min settlement value: %v", ErrRetriable, err)
	}
	if returnAmount.Cmp(minValue) < 0 {
		return fmt.Errorf("%w: returnAmount %s < principal plus fee floor %s - contract would revert", ErrDeterministic, returnAmount, minValue)
	}

	tx, err := a.sendTx(ctx,
		func(opts *bind.TransactOpts) { opts.Value = returnAmount },
		func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return a.contract.SettleChannel(opts, channelID)
		})
	if err != nil {
		return err
	}
	receipt, err := a.confirmTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("%w: tx wait failed: %v", ErrRetriable, err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("%w: settleChannel reverted tx=%s", ErrContractRevert, tx.Hash())
	}
	return nil
}

// GetOperator returns the delegated operator address from LiquidityPool.
func (a *LiquidityPoolAdapter) GetOperator(ctx context.Context) (common.Address, error) {
	op, err := a.contract.Operator(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, fmt.Errorf("%w: operator: %v", ErrRetriable, err)
	}
	return op, nil
}

// GetPoolState returns total assets and locked ETH values for Hub accounting.
// Both values are wei-denominated big.Ints: pool sizes routinely exceed the
// ~18.45 ETH that fits in a uint64 of wei.
func (a *LiquidityPoolAdapter) GetPoolState(ctx context.Context) (reserve *big.Int, locked *big.Int, err error) {
	total, err := a.contract.TotalAssets(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: totalAssets: %v", ErrRetriable, err)
	}
	lockedBig, err := a.contract.TotalLockedETH(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, fmt.Errorf("%w: totalLockedETH: %v", ErrRetriable, err)
	}
	return total, lockedBig, nil
}

// OperatorBond returns the ETH collateral the operator has posted in the pool.
func (a *LiquidityPoolAdapter) OperatorBond(ctx context.Context) (*big.Int, error) {
	bond, err := a.contract.OperatorBond(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("%w: operatorBond: %v", ErrRetriable, err)
	}
	return bond, nil
}

// BondETH posts amount wei of operator collateral via bondETH(). The tx
// sender must be the pool operator.
func (a *LiquidityPoolAdapter) BondETH(ctx context.Context, amount *big.Int) error {
	if amount == nil || amount.Sign() <= 0 {
		return fmt.Errorf("%w: bond amount must be > 0", ErrDeterministic)
	}
	tx, err := a.sendTx(ctx,
		func(opts *bind.TransactOpts) { opts.Value = amount },
		func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return a.contract.BondETH(opts)
		})
	if err != nil {
		return err
	}
	receipt, err := a.confirmTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("%w: tx wait failed: %v", ErrRetriable, err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("%w: bondETH reverted tx=%s", ErrContractRevert, tx.Hash())
	}
	return nil
}

// DepositFor deposits amount wei into the pool and mints the resulting shares
// to beneficiary rather than to the tx sender. Used to credit a CKB LP cell's
// traded countervalue to that cell's owner-designated beneficiary: the sender
// funds the deposit but keeps no claim on it.
func (a *LiquidityPoolAdapter) DepositFor(ctx context.Context, beneficiary common.Address, amount *big.Int) error {
	if amount == nil || amount.Sign() <= 0 {
		return fmt.Errorf("%w: deposit amount must be > 0", ErrDeterministic)
	}
	if beneficiary == (common.Address{}) {
		return fmt.Errorf("%w: beneficiary must not be the zero address", ErrDeterministic)
	}
	tx, err := a.sendTx(ctx,
		func(opts *bind.TransactOpts) { opts.Value = amount },
		func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return a.contract.DepositFor(opts, beneficiary)
		})
	if err != nil {
		return err
	}
	receipt, err := a.confirmTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("%w: tx wait failed: %v", ErrRetriable, err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("%w: depositFor reverted tx=%s", ErrContractRevert, tx.Hash())
	}
	return nil
}

// SharesOf returns the LP provider's share balance in the pool.
func (a *LiquidityPoolAdapter) SharesOf(ctx context.Context, owner common.Address) (*big.Int, error) {
	shares, err := a.contract.SharesOf(&bind.CallOpts{Context: ctx}, owner)
	if err != nil {
		return nil, fmt.Errorf("%w: sharesOf: %v", ErrRetriable, err)
	}
	return shares, nil
}

// WithdrawableETH returns the pool's free (non-locked) ETH available for withdrawal.
func (a *LiquidityPoolAdapter) WithdrawableETH(ctx context.Context) (*big.Int, error) {
	free, err := a.contract.WithdrawableETH(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("%w: withdrawableETH: %v", ErrRetriable, err)
	}
	return free, nil
}

// PoolMetadata returns the pool address and chain ID for the UI to target the contract.
func (a *LiquidityPoolAdapter) PoolMetadata() (addr common.Address, chainID *big.Int) {
	if a.chainID == nil {
		return a.poolAddress, nil
	}
	return a.poolAddress, new(big.Int).Set(a.chainID)
}

// TotalShares returns the total number of LP shares issued by the pool. Combined
// with SharesOf it lets callers compute an owner's pro-rata fraction of the pool.
func (a *LiquidityPoolAdapter) TotalShares(ctx context.Context) (*big.Int, error) {
	shares, err := a.contract.TotalShares(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("%w: totalShares: %v", ErrRetriable, err)
	}
	return shares, nil
}

// PreviewWithdrawETH returns the ETH the contract would pay out for burning the
// given number of shares. The pool's withdraw reverts if it would touch locked
// ETH, so what is withdrawable right now is min(PreviewWithdrawETH(shares),
// WithdrawableETH()).
func (a *LiquidityPoolAdapter) PreviewWithdrawETH(ctx context.Context, shares *big.Int) (*big.Int, error) {
	out, err := a.contract.PreviewWithdrawETH(&bind.CallOpts{Context: ctx}, shares)
	if err != nil {
		return nil, fmt.Errorf("%w: previewWithdrawETH: %v", ErrRetriable, err)
	}
	return out, nil
}
