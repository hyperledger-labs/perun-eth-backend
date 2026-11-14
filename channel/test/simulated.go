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
	"slices"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
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
	// value plus MinGasTipCap to avoid the error 'max fee per gas less than
	// block base fee'.
	InitialGasBaseFee = params.InitialBaseFee
	// MinGasTipCap is the minimal gas tip that the simulated backend's miner
	// accepts. Transactions with a lower tip are never mined.
	MinGasTipCap = params.GWei / 1000
	// internal gas limit of the simulated backend.
	simBackendGasLimit = 8_000_000
	// block time increase of auto-mined blocks.
	minedBlockTime = 10 * time.Second
	// number of retries of AdjustTime when transactions are pending.
	adjustTimeRetries = 3
)

// SimulatedBackend provides a simulated ethereum blockchain for tests.
type SimulatedBackend struct {
	backends.SimulatedBackend
	sbMtx sync.Mutex // protects SimulatedBackend
	// sealMtx serializes block sealing, which the simulated backend does not
	// support concurrently.
	sealMtx sync.Mutex

	Signer types.Signer

	faucetKey     *ecdsa.PrivateKey
	faucetAddr    common.Address
	clockMu       sync.Mutex    // Mutex for clock adjustments. Locked by SimTimeouts.
	mining        chan struct{} // Used for auto-mining blocks.
	stoppedMining chan struct{} // For making sure that mining stopped.
	commitTx      bool          // Whether each transaction is committed.
	chainID       *big.Int      // Chain ID of the simulated chain, nil for the default.
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
	SimBackendOpt func(*SimulatedBackend)
)

// NewSimulatedBackend creates a new Simulated Backend.
func NewSimulatedBackend(opts ...SimBackendOpt) *SimulatedBackend {
	sk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)
	addr := map[common.Address]types.Account{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
	}
	addr[faucetAddr] = types.Account{Balance: new(big.Int).Sub(channel.MaxBalance, big.NewInt(int64(len(addr))))}

	sb := &SimulatedBackend{
		faucetKey:  sk,
		faucetAddr: faucetAddr,
		commitTx:   true,
	}
	for _, opt := range opts {
		opt(sb)
	}

	b := simulated.NewBackend(types.GenesisAlloc(addr),
		simulated.WithBlockGasLimit(simBackendGasLimit),
		withChainConfig(sb.chainID),
	)
	sb.SimulatedBackend = backends.SimulatedBackend{Backend: b, Client: b.Client()}
	sb.Signer = types.LatestSignerForChainID(sb.ChainID())
	return sb
}

// withChainConfig adjusts the chain config of the simulated backend. Forks that
// are not yet scheduled on mainnet are disabled, as they change gas costs. If
// chainID is not nil, it replaces the default chain ID.
func withChainConfig(chainID *big.Int) func(*node.Config, *ethconfig.Config) {
	return func(_ *node.Config, ethConf *ethconfig.Config) {
		cfg := *ethConf.Genesis.Config
		cfg.BogotaTime = nil
		if chainID != nil {
			cfg.ChainID = new(big.Int).Set(chainID)
		}
		ethConf.Genesis.Config = &cfg
	}
}

// SendTransaction executes a transaction.
func (s *SimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if err := s.SimulatedBackend.SendTransaction(ctx, tx); err != nil {
		return errors.WithStack(err)
	}
	if s.commitTx {
		s.Commit()
	}
	return nil
}

// Commit seals a block containing the pending transactions.
func (s *SimulatedBackend) Commit() common.Hash {
	s.sealMtx.Lock()
	defer s.sealMtx.Unlock()
	return s.SimulatedBackend.Commit()
}

// AdjustTime mines a new block whose time is advanced by the given duration.
// The simulated backend can only adjust the time on empty blocks, so pending
// transactions are committed and the adjustment is retried.
func (s *SimulatedBackend) AdjustTime(adjustment time.Duration) error {
	s.sealMtx.Lock()
	defer s.sealMtx.Unlock()
	for i := 0; ; i++ {
		err := s.SimulatedBackend.AdjustTime(adjustment)
		if err == nil || i == adjustTimeRetries {
			return errors.WithStack(err)
		}
		s.SimulatedBackend.Commit()
	}
}

// Fork sets the head to the given parent block. Transactions of the orphaned
// blocks are not re-added to the transaction pool.
func (s *SimulatedBackend) Fork(ctx context.Context, parentHash common.Hash) error {
	s.sealMtx.Lock()
	defer s.sealMtx.Unlock()
	if err := s.SimulatedBackend.Fork(ctx, parentHash); err != nil {
		return errors.WithStack(err)
	}
	s.Rollback()
	return nil
}

// FundAddress funds a given address with `test.MaxBalance` eth from a faucet.
func (s *SimulatedBackend) FundAddress(ctx context.Context, addr common.Address) {
	nonce, err := s.PendingNonceAt(ctx, s.faucetAddr)
	if err != nil {
		panic(err)
	}
	txdata := &types.DynamicFeeTx{
		Nonce:     nonce,
		GasFeeCap: big.NewInt(InitialGasBaseFee + MinGasTipCap),
		GasTipCap: big.NewInt(MinGasTipCap),
		Gas:       params.TxGas,
		To:        &addr,
		Value:     test.MaxBalance,
	}
	tx, err := types.SignNewTx(s.faucetKey, s.Signer, txdata)
	if err != nil {
		panic(err)
	}
	if err := s.SimulatedBackend.SendTransaction(ctx, tx); err != nil {
		panic(err)
	}
	s.Commit()
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
	go func() {
		log.Trace("Started mining")
		defer log.Trace("Stopped mining")
		defer close(s.stoppedMining)

		for {
			if err := s.AdjustTime(minedBlockTime); err != nil {
				log.Warn("Mining simulated block: ", err)
			} else {
				log.Trace("Mined simulated block")
			}

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
	close(s.mining)
	<-s.stoppedMining
}

// Reorg applies a chain reorg.
// `depth` is the number of blocks to be removed.
// `reorder` is a function that gets as input the removed blocks and outputs a list of blocks that are to be added after the removal.
// It is required that the number of added blocks is greater than `depth` for a reorg to be accepted.
// The nonce prevents transactions of the same account from being re-ordered. Trying to do this will panic.
func (s *SimulatedBackend) Reorg(ctx context.Context, depth uint64, reorder Reorder) error {
	// Lock
	if !s.sbMtx.TryLockCtx(ctx) {
		return errors.Errorf("locking mutex: %v", ctx.Err())
	}
	defer s.sbMtx.Unlock()

	// parent at current - depth.
	head, err := s.HeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "retrieving head")
	}
	parentN := new(big.Int).Sub(head.Number, big.NewInt(int64(depth)))
	parent, err := s.BlockByNumber(ctx, parentN)
	if err != nil {
		return errors.Wrap(err, "retrieving reorg parent")
	}

	// Collect orphaned blocks.
	blocks := make([]types.Transactions, depth)
	hashes := []common.Hash{parent.Hash()}
	for i := uint64(0); i < depth; i++ {
		blockN := new(big.Int).Add(parentN, big.NewInt(int64(i+1)))
		block, err := s.BlockByNumber(ctx, blockN)
		if err != nil {
			return errors.Wrap(err, "retrieving block")
		}
		// Add the TXs from block parent + 1 + i.
		blocks[i] = block.Transactions()
		hashes = append(hashes, block.Hash())
	}

	// Modify the blocks with the reorder callback.
	newBlocks := reorder(slices.Clone(blocks)) // Clone, as blocks is compared below.
	if uint64(len(newBlocks)) <= depth {
		return fmt.Errorf("number of blocks added %d must be greater than number of blocks removed %d", len(newBlocks), depth)
	}

	// Keep the unchanged blocks, so that the chain is only reorganized from
	// the first changed block on.
	keep := 0
	for keep < len(blocks) && sameTxs(blocks[keep], newBlocks[keep]) {
		keep++
	}
	newBlocks = newBlocks[keep:]

	// Reset the chain to the last kept block.
	if keep < len(blocks) {
		if err := s.Fork(ctx, hashes[keep]); err != nil {
			return errors.Wrap(err, "forking")
		}
	}

	// Add modified blocks.
	for _, txs := range newBlocks {
		for _, tx := range txs {
			if err := s.SimulatedBackend.SendTransaction(ctx, tx); err != nil {
				return errors.Wrap(err, "re-sending transaction")
			}
		}
		s.Commit()
	}
	return nil
}

// sameTxs returns whether both lists contain the same transactions in the
// same order.
func sameTxs(a, b types.Transactions) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Hash() != b[i].Hash() {
			return false
		}
	}
	return true
}

// ChainID returns the chainID of the underlying blockchain.
func (s *SimulatedBackend) ChainID() *big.Int {
	chainID, err := s.Client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	return chainID
}

// WithChainID sets the chain ID of the simulated blockchain.
func WithChainID(chainID *big.Int) SimBackendOpt {
	return func(sb *SimulatedBackend) { sb.chainID = chainID }
}

// WithCommitTx controls whether the simulated backend should automatically
// mine a block after a transaction was sent.
func WithCommitTx(b bool) SimBackendOpt {
	return func(sb *SimulatedBackend) { sb.commitTx = b }
}
