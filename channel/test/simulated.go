// Copyright 2019 - See NOTICE file for copyright holders.
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

package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	stdsync "sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/perun-network/perun-eth-backend/channel"
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"
	"github.com/pkg/errors"

	perunchannel "perun.network/go-perun/channel"
	"perun.network/go-perun/channel/test"
	"perun.network/go-perun/log"
	"perun.network/go-perun/wallet"
	"polycry.pt/poly-go/sync"
)

const (
	// InitialGasBaseFee is the simulated backend's initial base fee.
	// It should only decrease from the first block onwards, as no gas auctions
	// are taking place.
	//
	// When constructing a transaction manually, GasFeeCap can be set to this
	// value to avoid the error 'max fee per gas less than block base fee'.
	InitialGasBaseFee = 875_000_000
	// internal gas limit of the simulated backend.
	simBackendGasLimit = 8_000_000
	defaultSimChainID  = 1337
)

// SimulatedBackend provides a simulated ethereum blockchain for tests.
type SimulatedBackend struct {
	backend *simulated.Backend
	client  simulated.Client
	sbMtx   sync.Mutex // protects backend operations

	chainID *big.Int
	Signer  types.Signer

	faucetKey     *ecdsa.PrivateKey
	faucetAddr    common.Address
	clockMu       sync.Mutex    // Mutex for clock adjustments. Locked by SimTimeouts.
	mining        chan struct{} // Used for auto-mining blocks.
	stoppedMining chan struct{} // For making sure that mining stopped.
	stopMiningOnce stdsync.Once
	commitTx      bool          // Whether each transaction is committed.
}

// BalanceReader is a balance reader used for testing. It is associated with a
// given account.
type BalanceReader struct {
	b   *SimulatedBackend
	acc wallet.Address
}

// Balance returns the asset balance of the associated account.
func (br *BalanceReader) Balance(asset perunchannel.Asset) perunchannel.Bal {
	return br.b.Balance(br.acc, asset)
}

// NewBalanceReader creates a new balance reader for the given account.
func (s *SimulatedBackend) NewBalanceReader(acc wallet.Address) *BalanceReader {
	return &BalanceReader{
		b:   s,
		acc: acc,
	}
}

// Balance returns the balance of the given address on the simulated backend.
func (s *SimulatedBackend) Balance(addr wallet.Address, _ perunchannel.Asset) perunchannel.Bal {
	ctx := context.Background()
	bal, _ := s.BalanceAt(ctx, ethwallet.AsEthAddr(addr), nil)
	return bal
}

type (
	// Reorder can be used to insert, reorder and exclude transactions in
	// combination with `Reorg`.
	Reorder func([]types.Transactions) []types.Transactions

	// SimBackendOpt represents an optional argument for the sim backend.
	SimBackendOpt func(*simBackendConfig)
)

type simBackendConfig struct {
	chainID  *big.Int
	commitTx bool
	options  []func(nodeConf *node.Config, ethConf *ethconfig.Config)
}

func normalizeCtx(ctx context.Context) context.Context {
	if ctx == nil {
		// Keep older tests working even though callers should prefer explicit
		// contexts so cancellation bugs are easier to spot.
		return context.Background()
	}
	return ctx
}

// NewSimulatedBackend creates a new Simulated Backend.
func NewSimulatedBackend(opts ...SimBackendOpt) *SimulatedBackend {
	cfg := simBackendConfig{
		chainID:  big.NewInt(defaultSimChainID),
		commitTx: true,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	sk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)
	alloc := types.GenesisAlloc{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
		faucetAddr:                       {Balance: new(big.Int).Sub(channel.MaxBalance, big.NewInt(8))},
	}
	backendOpts := []func(nodeConf *node.Config, ethConf *ethconfig.Config){
		simulated.WithBlockGasLimit(simBackendGasLimit),
		func(nodeConf *node.Config, ethConf *ethconfig.Config) {
			chainConfig := *params.AllDevChainProtocolChanges
			chainConfig.ChainID = new(big.Int).Set(cfg.chainID)
			ethConf.Genesis.Config = &chainConfig
		},
	}
	backendOpts = append(backendOpts, cfg.options...)
	b := simulated.NewBackend(alloc, backendOpts...)
	sb := &SimulatedBackend{
		backend:    b,
		client:     b.Client(),
		chainID:    new(big.Int).Set(cfg.chainID),
		faucetKey:  sk,
		faucetAddr: faucetAddr,
		commitTx:   cfg.commitTx,
	}
	sb.Signer = types.LatestSignerForChainID(sb.ChainID())
	return sb
}

// SendTransaction executes a transaction.
func (s *SimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ctx = normalizeCtx(ctx)
	s.sbMtx.Lock()
	defer s.sbMtx.Unlock()

	if err := s.client.SendTransaction(ctx, tx); err != nil {
		return errors.WithStack(err)
	}
	if s.commitTx {
		s.backend.Commit()
	}
	return nil
}

// FundAddress funds a given address with `test.MaxBalance` eth from a faucet.
func (s *SimulatedBackend) FundAddress(ctx context.Context, addr common.Address) {
	ctx = normalizeCtx(ctx)
	s.sbMtx.Lock()
	defer s.sbMtx.Unlock()

	nonce, err := s.PendingNonceAt(ctx, s.faucetAddr)
	if err != nil {
		panic(err)
	}
	gasFeeCap := new(big.Int).Add(big.NewInt(InitialGasBaseFee), big.NewInt(params.GWei))
	txdata := &types.DynamicFeeTx{
		ChainID:   s.ChainID(),
		Nonce:     nonce,
		GasTipCap: big.NewInt(params.GWei),
		GasFeeCap: gasFeeCap,
		Gas:       params.TxGas,
		To:        &addr,
		Value:     test.MaxBalance,
	}
	tx, err := types.SignNewTx(s.faucetKey, s.Signer, txdata)
	if err != nil {
		panic(err)
	}
	if err := s.client.SendTransaction(ctx, tx); err != nil {
		panic(err)
	}
	s.backend.Commit()
	if _, err := bind.WaitMined(ctx, s, tx); err != nil {
		panic(err)
	}
}

// StartMining makes the simulated blockchain auto-mine blocks with the given
// interval. Must be stopped with `StopMining`.
// The block time of generated blocks will always increase by 10 seconds.
func (s *SimulatedBackend) StartMining(interval time.Duration) {
	if interval == 0 {
		panic("blockTime can not be zero")
	}

	s.mining = make(chan struct{})
	s.stoppedMining = make(chan struct{})
	s.stopMiningOnce = stdsync.Once{}
	go func() {
		log.Trace("Started mining")
		defer log.Trace("Stopped mining")
		defer close(s.stoppedMining)

		for {
			s.Commit()
			log.Trace("Mined simulated block")

			select {
			case <-time.After(interval):
			case <-s.mining: // stopped
				return
			}
		}
	}()
}

// StopMining stops the auto-mining of the simulated blockchain.
// Must be called exactly once to free resources iff `StartMining` was called.
// Waits until the auto-mining routine terminates.
func (s *SimulatedBackend) StopMining() {
	if s.mining == nil || s.stoppedMining == nil {
		return
	}
	s.stopMiningOnce.Do(func() {
		close(s.mining)
		<-s.stoppedMining
	})
}

// Reorg applies a chain reorg.
func (s *SimulatedBackend) Reorg(ctx context.Context, depth uint64, reorder Reorder) error {
	if !s.sbMtx.TryLockCtx(ctx) {
		return errors.Errorf("locking mutex: %v", ctx.Err())
	}
	defer s.sbMtx.Unlock()

	currentBlock, err := s.BlockByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "retrieving current block")
	}
	parentN := new(big.Int).Sub(currentBlock.Number(), big.NewInt(int64(depth)))
	parent, err := s.BlockByNumber(ctx, parentN)
	if err != nil {
		return errors.Wrap(err, "retrieving reorg parent")
	}

	blocks := make([]types.Transactions, depth)
	for i := uint64(0); i < depth; i++ {
		blockN := new(big.Int).Add(parentN, big.NewInt(int64(i+1)))
		block, err := s.BlockByNumber(ctx, blockN)
		if err != nil {
			return errors.Wrap(err, "retrieving block")
		}
		blocks[i] = block.Transactions()
	}

	newBlocks := reorder(blocks)
	if uint64(len(newBlocks)) <= depth {
		return fmt.Errorf("number of blocks added %d must be greater than number of blocks removed %d", len(newBlocks), depth)
	}

	if err := s.backend.Fork(parent.Hash()); err != nil {
		return errors.Wrap(err, "forking")
	}
	s.backend.Rollback()

	for _, txs := range newBlocks {
		for _, tx := range txs {
			if err := s.client.SendTransaction(ctx, tx); err != nil {
				return errors.Wrap(err, "re-sending transaction")
			}
		}
		s.backend.Commit()
	}
	return nil
}

// Commit seals a block and moves the chain forward.
func (s *SimulatedBackend) Commit() common.Hash {
	s.sbMtx.Lock()
	defer s.sbMtx.Unlock()

	return s.backend.Commit()
}

// Rollback removes all pending transactions.
func (s *SimulatedBackend) Rollback() {
	s.sbMtx.Lock()
	defer s.sbMtx.Unlock()

	s.backend.Rollback()
}

// Close shuts down the simulated backend.
func (s *SimulatedBackend) Close() error {
	s.StopMining()
	return s.backend.Close()
}

// ChainID returns the chainID of the underlying blockchain.
func (s *SimulatedBackend) ChainID() *big.Int {
	return new(big.Int).Set(s.chainID)
}

// WithCommitTx controls whether the simulated backend should automatically
// mine a block after a transaction was sent.
func WithCommitTx(b bool) SimBackendOpt {
	return func(cfg *simBackendConfig) { cfg.commitTx = b }
}

// WithChainID configures the simulated backend to use the specified chain ID.
func WithChainID(chainID *big.Int) SimBackendOpt {
	if chainID == nil || chainID.Sign() < 0 {
		panic("invalid chain ID")
	}
	return func(cfg *simBackendConfig) {
		cfg.chainID = new(big.Int).Set(chainID)
	}
}

// AdjustTime changes the block timestamp and creates a new block.
func (s *SimulatedBackend) AdjustTime(adjustment time.Duration) error {
	s.sbMtx.Lock()
	defer s.sbMtx.Unlock()

	return s.backend.AdjustTime(adjustment)
}

// Delegated methods from simulated.Client to satisfy ethclient interfaces.

func (s *SimulatedBackend) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	ctx = normalizeCtx(ctx)
	return s.client.BalanceAt(ctx, account, blockNumber)
}

func (s *SimulatedBackend) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ctx = normalizeCtx(ctx)
	return s.client.BlockByHash(ctx, hash)
}

func (s *SimulatedBackend) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ctx = normalizeCtx(ctx)
	return s.client.BlockByNumber(ctx, number)
}

func (s *SimulatedBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	ctx = normalizeCtx(ctx)
	return s.client.HeaderByHash(ctx, hash)
}

func (s *SimulatedBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ctx = normalizeCtx(ctx)
	return s.client.HeaderByNumber(ctx, number)
}

func (s *SimulatedBackend) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	ctx = normalizeCtx(ctx)
	return s.client.TransactionCount(ctx, blockHash)
}

func (s *SimulatedBackend) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	ctx = normalizeCtx(ctx)
	return s.client.TransactionInBlock(ctx, blockHash, index)
}

func (s *SimulatedBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	ctx = normalizeCtx(ctx)
	return s.client.CodeAt(ctx, contract, blockNumber)
}

func (s *SimulatedBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ctx = normalizeCtx(ctx)
	return s.client.CallContract(ctx, call, blockNumber)
}

func (s *SimulatedBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	ctx = normalizeCtx(ctx)
	return s.client.PendingCodeAt(ctx, account)
}

func (s *SimulatedBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	ctx = normalizeCtx(ctx)
	return s.client.PendingNonceAt(ctx, account)
}

func (s *SimulatedBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ctx = normalizeCtx(ctx)
	return s.client.SuggestGasPrice(ctx)
}

func (s *SimulatedBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	ctx = normalizeCtx(ctx)
	return s.client.SuggestGasTipCap(ctx)
}

func (s *SimulatedBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	ctx = normalizeCtx(ctx)
	return s.client.EstimateGas(ctx, call)
}

func (s *SimulatedBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	ctx = normalizeCtx(ctx)
	return s.client.FilterLogs(ctx, query)
}

func (s *SimulatedBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	ctx = normalizeCtx(ctx)
	return s.client.SubscribeFilterLogs(ctx, query, ch)
}

func (s *SimulatedBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ctx = normalizeCtx(ctx)
	return s.client.TransactionReceipt(ctx, txHash)
}

func (s *SimulatedBackend) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ctx = normalizeCtx(ctx)
	return s.client.SubscribeNewHead(ctx, ch)
}

func (s *SimulatedBackend) TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	ctx = normalizeCtx(ctx)
	return s.client.TransactionByHash(ctx, txHash)
}

func (s *SimulatedBackend) BlockNumber(ctx context.Context) (uint64, error) {
	ctx = normalizeCtx(ctx)
	return s.client.BlockNumber(ctx)
}

func (s *SimulatedBackend) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	ctx = normalizeCtx(ctx)
	return s.client.NonceAt(ctx, account, blockNumber)
}

func (s *SimulatedBackend) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	ctx = normalizeCtx(ctx)
	return s.client.PendingCallContract(ctx, call)
}

func (s *SimulatedBackend) SubscribeTransactionReceipts(ctx context.Context, q *ethereum.TransactionReceiptsQuery, ch chan<- []*types.Receipt) (ethereum.Subscription, error) {
	ctx = normalizeCtx(ctx)
	return s.client.SubscribeTransactionReceipts(ctx, q, ch)
}
