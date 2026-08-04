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

//nolint:testpackage // white-box: exercises unexported helpers and mocks
package adapter

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/perun-network/perun-eth-backend/bindings/liquiditypool"
	"github.com/stretchr/testify/require"
)

type mockContract struct {
	withdrawable    *big.Int
	shares          *big.Int
	totalShares     *big.Int
	previewWithdraw *big.Int
	locked          *big.Int
	totalAssets     *big.Int
	totalLocked     *big.Int
	bond            *big.Int
	minSettle       *big.Int
	operator        common.Address

	withdrawErr    error
	sharesErr      error
	totalSharesErr error
	previewErr     error
	lockedErr      error
	fundErr        error
	settleErr      error
	bondCallErr    error
	minSettleErr   error
	bondTxErr      error
	totalErr       error
	lockedTErr     error
	opErr          error

	fundTx     *types.Transaction
	settleTx   *types.Transaction
	bondTx     *types.Transaction
	depositTx  *types.Transaction
	depositErr error

	// fundTransientErrs makes FundChannel reject the submission with a
	// transient mempool error this many times before accepting it, so the
	// send-retry path can be exercised without a node.
	fundTransientErrs int
	fundCalls         int

	lastSettleValue      *big.Int
	lastBondValue        *big.Int
	lastDepositValue     *big.Int
	lastDepositRecipient common.Address
}

func (m *mockContract) WithdrawableETH(_ *bind.CallOpts) (*big.Int, error) {
	if m.withdrawErr != nil {
		return nil, m.withdrawErr
	}
	return new(big.Int).Set(m.withdrawable), nil
}

func (m *mockContract) SharesOf(_ *bind.CallOpts, _ common.Address) (*big.Int, error) {
	if m.sharesErr != nil {
		return nil, m.sharesErr
	}
	return new(big.Int).Set(m.shares), nil
}

func (m *mockContract) TotalShares(_ *bind.CallOpts) (*big.Int, error) {
	if m.totalSharesErr != nil {
		return nil, m.totalSharesErr
	}
	return new(big.Int).Set(m.totalShares), nil
}

func (m *mockContract) PreviewWithdrawETH(_ *bind.CallOpts, _ *big.Int) (*big.Int, error) {
	if m.previewErr != nil {
		return nil, m.previewErr
	}
	return new(big.Int).Set(m.previewWithdraw), nil
}

func (m *mockContract) LockedByChannel(_ *bind.CallOpts, _ [32]byte) (*big.Int, error) {
	if m.lockedErr != nil {
		return nil, m.lockedErr
	}
	return new(big.Int).Set(m.locked), nil
}

func (m *mockContract) FundChannel(_ *bind.TransactOpts, _ [32]byte, _ *big.Int) (*types.Transaction, error) {
	m.fundCalls++
	if m.fundTransientErrs > 0 {
		m.fundTransientErrs--
		return nil, errors.New("transaction underpriced")
	}
	if m.fundErr != nil {
		return nil, m.fundErr
	}
	return m.fundTx, nil
}

func (m *mockContract) SettleChannel(opts *bind.TransactOpts, _ [32]byte) (*types.Transaction, error) {
	if opts.Value != nil {
		m.lastSettleValue = new(big.Int).Set(opts.Value)
	}
	if m.settleErr != nil {
		return nil, m.settleErr
	}
	return m.settleTx, nil
}

// OperatorBond defaults to an ample bond when unset so pre-existing funding
// tests exercise their own concern, not coverage.
func (m *mockContract) OperatorBond(_ *bind.CallOpts) (*big.Int, error) {
	if m.bondCallErr != nil {
		return nil, m.bondCallErr
	}
	if m.bond == nil {
		return new(big.Int).Lsh(big.NewInt(1), 128), nil
	}
	return new(big.Int).Set(m.bond), nil
}

// MinSettlementValue defaults to the locked principal (a zero fee floor)
// when unset.
func (m *mockContract) MinSettlementValue(_ *bind.CallOpts, _ [32]byte) (*big.Int, error) {
	if m.minSettleErr != nil {
		return nil, m.minSettleErr
	}
	if m.minSettle == nil {
		return new(big.Int).Set(m.locked), nil
	}
	return new(big.Int).Set(m.minSettle), nil
}

func (m *mockContract) BondETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	if opts.Value != nil {
		m.lastBondValue = new(big.Int).Set(opts.Value)
	}
	if m.bondTxErr != nil {
		return nil, m.bondTxErr
	}
	return m.bondTx, nil
}

func (m *mockContract) DepositFor(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	if opts.Value != nil {
		m.lastDepositValue = new(big.Int).Set(opts.Value)
	}
	m.lastDepositRecipient = beneficiary
	if m.depositErr != nil {
		return nil, m.depositErr
	}
	return m.depositTx, nil
}

func (m *mockContract) Operator(_ *bind.CallOpts) (common.Address, error) {
	if m.opErr != nil {
		return common.Address{}, m.opErr
	}
	return m.operator, nil
}

func (m *mockContract) TotalAssets(_ *bind.CallOpts) (*big.Int, error) {
	if m.totalErr != nil {
		return nil, m.totalErr
	}
	return new(big.Int).Set(m.totalAssets), nil
}

func (m *mockContract) TotalLockedETH(_ *bind.CallOpts) (*big.Int, error) {
	if m.lockedTErr != nil {
		return nil, m.lockedTErr
	}
	if m.totalLocked == nil {
		return big.NewInt(0), nil
	}
	return new(big.Int).Set(m.totalLocked), nil
}

func fakeTx(selector []byte, value *big.Int) *types.Transaction {
	to := common.HexToAddress("0x1234")
	return types.NewTx(&types.LegacyTx{
		Nonce:    1,
		To:       &to,
		Gas:      21000,
		GasPrice: big.NewInt(1),
		Value:    value,
		Data:     selector,
	})
}

func TestFundChannel_HappyPath(t *testing.T) {
	m := &mockContract{
		withdrawable: big.NewInt(100),
		locked:       big.NewInt(0),
		fundTx:       fakeTx([]byte{0x26, 0x0b, 0x60, 0x86}, big.NewInt(0)),
	}

	var captured *types.Transaction
	a := newTestAdapter(m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, func(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
		captured = tx
		return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
	})

	err := a.FundChannel(context.Background(), [32]byte{1}, big.NewInt(42))
	require.NoError(t, err)
	require.NotNil(t, captured)
	require.Equal(t, []byte{0x26, 0x0b, 0x60, 0x86}, captured.Data())
}

func TestFundChannel_AlreadyFunded(t *testing.T) {
	m := &mockContract{withdrawable: big.NewInt(100), locked: big.NewInt(1)}
	a := newTestAdapter(m, nil, nil)

	err := a.FundChannel(context.Background(), [32]byte{2}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestFundChannel_InsufficientFree(t *testing.T) {
	m := &mockContract{withdrawable: big.NewInt(9), locked: big.NewInt(0)}
	a := newTestAdapter(m, nil, nil)

	err := a.FundChannel(context.Background(), [32]byte{2}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestSettleChannel_HappyPath(t *testing.T) {
	m := &mockContract{
		locked:   big.NewInt(15),
		settleTx: fakeTx([]byte{0xfd, 0x3d, 0x31, 0x99}, big.NewInt(15)),
	}

	var captured *types.Transaction
	a := newTestAdapter(m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, func(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
		captured = tx
		return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
	})

	ret := big.NewInt(20)
	err := a.SettleChannel(context.Background(), [32]byte{3}, ret)
	require.NoError(t, err)
	require.Equal(t, ret, m.lastSettleValue)
	require.Equal(t, []byte{0xfd, 0x3d, 0x31, 0x99}, captured.Data())
}

func TestSettleChannel_BelowPrincipal(t *testing.T) {
	m := &mockContract{locked: big.NewInt(11)}
	a := newTestAdapter(m, nil, nil)

	err := a.SettleChannel(context.Background(), [32]byte{4}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestSettleChannel_BelowFeeFloor(t *testing.T) {
	// Principal covered but the on-chain fee floor is not.
	m := &mockContract{locked: big.NewInt(11), minSettle: big.NewInt(12)}
	a := newTestAdapter(m, nil, nil)

	err := a.SettleChannel(context.Background(), [32]byte{4}, big.NewInt(11))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestFundChannel_InsufficientBondCoverage(t *testing.T) {
	m := &mockContract{
		withdrawable: big.NewInt(100),
		locked:       big.NewInt(0),
		totalLocked:  big.NewInt(30),
		bond:         big.NewInt(40),
	}
	a := newTestAdapter(m, nil, nil)

	// 30 locked + 11 requested = 41 > 40 bonded.
	err := a.FundChannel(context.Background(), [32]byte{5}, big.NewInt(11))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
	require.Contains(t, err.Error(), "insufficient coverage")
}

func TestBondETH_HappyPath(t *testing.T) {
	m := &mockContract{
		bondTx: fakeTx([]byte{0x0e, 0xf1, 0x1e, 0x5c}, big.NewInt(50)),
	}
	a := newTestAdapter(m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, func(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
		return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
	})

	err := a.BondETH(context.Background(), big.NewInt(50))
	require.NoError(t, err)
	require.Equal(t, big.NewInt(50), m.lastBondValue)
}

func TestBondETH_RejectsNonPositive(t *testing.T) {
	a := newTestAdapter(&mockContract{}, nil, nil)

	err := a.BondETH(context.Background(), big.NewInt(0))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

// TestDepositFor_HappyPath also locks the wei range: a conversion credits the
// full countervalue, which for a 100-ETH-scale pool exceeds 2^64.
func TestDepositFor_HappyPath(t *testing.T) {
	amount, ok := new(big.Int).SetString("100000000000000000000", 10) // 100 ETH
	require.True(t, ok)
	require.Positive(t, amount.Cmp(new(big.Int).Lsh(big.NewInt(1), 64)))

	beneficiary := common.HexToAddress("0xA298Fc05bccff341f340a11FffA30567a00e651f")
	m := &mockContract{
		depositTx: fakeTx([]byte{0x2f, 0x4f, 0x21, 0xe2}, amount),
	}
	a := newTestAdapter(m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, func(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
		return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
	})

	err := a.DepositFor(context.Background(), beneficiary, amount)
	require.NoError(t, err)
	require.Equal(t, amount, m.lastDepositValue)
	require.Equal(t, beneficiary, m.lastDepositRecipient)
}

func TestDepositFor_RejectsNonPositive(t *testing.T) {
	a := newTestAdapter(&mockContract{}, nil, nil)

	err := a.DepositFor(context.Background(), common.HexToAddress("0x1"), big.NewInt(0))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestDepositFor_RejectsZeroBeneficiary(t *testing.T) {
	a := newTestAdapter(&mockContract{}, nil, nil)

	// The contract reverts on the zero beneficiary; fail before spending gas.
	err := a.DepositFor(context.Background(), common.Address{}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestDepositFor_RevertedTx(t *testing.T) {
	m := &mockContract{
		depositTx: fakeTx([]byte{0x2f, 0x4f, 0x21, 0xe2}, big.NewInt(10)),
	}
	a := newTestAdapter(m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, func(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
		return &types.Receipt{Status: types.ReceiptStatusFailed}, nil
	})

	err := a.DepositFor(context.Background(), common.HexToAddress("0x1"), big.NewInt(10))
	require.ErrorIs(t, err, ErrContractRevert)
}

func TestSettleChannel_ChannelNotFound(t *testing.T) {
	m := &mockContract{locked: big.NewInt(0)}
	a := newTestAdapter(m, nil, nil)

	err := a.SettleChannel(context.Background(), [32]byte{4}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestGetPoolState_HappyPath(t *testing.T) {
	m := &mockContract{totalAssets: big.NewInt(123), totalLocked: big.NewInt(45)}
	a := newTestAdapter(m, nil, nil)

	reserve, locked, err := a.GetPoolState(context.Background())
	require.NoError(t, err)
	require.Equal(t, big.NewInt(123), reserve)
	require.Equal(t, big.NewInt(45), locked)
}

func TestGetPoolState_BeyondUint64(t *testing.T) {
	// 100 ETH = 1e20 wei exceeds uint64 (~18.45 ETH); values must pass
	// through untruncated.
	total, ok := new(big.Int).SetString("100000000000000000000", 10)
	require.True(t, ok)
	locked, ok := new(big.Int).SetString("20000000000000000000", 10)
	require.True(t, ok)
	m := &mockContract{totalAssets: total, totalLocked: locked}
	a := newTestAdapter(m, nil, nil)

	gotReserve, gotLocked, err := a.GetPoolState(context.Background())
	require.NoError(t, err)
	require.Zero(t, gotReserve.Cmp(total))
	require.Zero(t, gotLocked.Cmp(locked))
}

func TestGetOperator_Delegated(t *testing.T) {
	want := common.HexToAddress("0x1000")
	m := &mockContract{operator: want}
	a := newTestAdapter(m, nil, nil)

	got, err := a.GetOperator(context.Background())
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestSharesOf_HappyPath(t *testing.T) {
	m := &mockContract{shares: big.NewInt(777)}
	a := newTestAdapter(m, nil, nil)

	got, err := a.SharesOf(context.Background(), common.HexToAddress("0xABCD"))
	require.NoError(t, err)
	require.Equal(t, big.NewInt(777), got)
}

func TestSharesOf_RetriableOnError(t *testing.T) {
	m := &mockContract{sharesErr: errors.New("node down")}
	a := newTestAdapter(m, nil, nil)

	_, err := a.SharesOf(context.Background(), common.HexToAddress("0xABCD"))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRetriable)
}

func TestWithdrawableETH_HappyPath(t *testing.T) {
	m := &mockContract{withdrawable: big.NewInt(555)}
	a := newTestAdapter(m, nil, nil)

	got, err := a.WithdrawableETH(context.Background())
	require.NoError(t, err)
	require.Equal(t, big.NewInt(555), got)
}

func TestWithdrawableETH_RetriableOnError(t *testing.T) {
	m := &mockContract{withdrawErr: errors.New("node down")}
	a := newTestAdapter(m, nil, nil)

	_, err := a.WithdrawableETH(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRetriable)
}

func TestTotalShares_HappyPath(t *testing.T) {
	m := &mockContract{totalShares: big.NewInt(1000)}
	a := newTestAdapter(m, nil, nil)

	got, err := a.TotalShares(context.Background())
	require.NoError(t, err)
	require.Equal(t, big.NewInt(1000), got)
}

func TestTotalShares_RetriableOnError(t *testing.T) {
	m := &mockContract{totalSharesErr: errors.New("node down")}
	a := newTestAdapter(m, nil, nil)

	_, err := a.TotalShares(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRetriable)
}

func TestPreviewWithdrawETH_HappyPath(t *testing.T) {
	m := &mockContract{previewWithdraw: big.NewInt(250)}
	a := newTestAdapter(m, nil, nil)

	got, err := a.PreviewWithdrawETH(context.Background(), big.NewInt(100))
	require.NoError(t, err)
	require.Equal(t, big.NewInt(250), got)
}

func TestPreviewWithdrawETH_RetriableOnError(t *testing.T) {
	m := &mockContract{previewErr: errors.New("node down")}
	a := newTestAdapter(m, nil, nil)

	_, err := a.PreviewWithdrawETH(context.Background(), big.NewInt(100))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRetriable)
}

func TestPoolMetadata(t *testing.T) {
	addr := common.HexToAddress("0xABCD")
	a := newTestAdapter(&mockContract{}, nil, nil)
	a.poolAddress = addr
	a.chainID = big.NewInt(1337)

	gotAddr, gotChain := a.PoolMetadata()
	require.Equal(t, addr, gotAddr)
	require.Equal(t, big.NewInt(1337), gotChain)

	// Returned chainID must be a copy; mutating it must not affect the adapter.
	gotChain.SetInt64(1)
	_, again := a.PoolMetadata()
	require.Equal(t, big.NewInt(1337), again)
}

func TestPoolMetadata_NilChainID(t *testing.T) {
	a := newTestAdapter(&mockContract{}, nil, nil)
	addr, chainID := a.PoolMetadata()
	require.Equal(t, common.Address{}, addr)
	require.Nil(t, chainID)
}

func TestRetriableOnRPCTimeout(t *testing.T) {
	err := classifyEthError(context.DeadlineExceeded)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRetriable)
}

func TestOperatorUpdatedEvent(t *testing.T) {
	parsed, err := abi.JSON(strings.NewReader(liquiditypool.LiquidityPoolABI))
	require.NoError(t, err)

	prev := common.HexToAddress("0x1111")
	next := common.HexToAddress("0x2222")
	log := types.Log{
		Topics: []common.Hash{
			parsed.Events["OperatorUpdated"].ID,
			common.BytesToHash(common.LeftPadBytes(prev.Bytes(), 32)),
			common.BytesToHash(common.LeftPadBytes(next.Bytes(), 32)),
		},
	}

	decoded := struct {
		PreviousOperator common.Address
		NewOperator      common.Address
	}{}
	err = abi.ParseTopics(&decoded, parsed.Events["OperatorUpdated"].Inputs, log.Topics[1:])
	require.NoError(t, err)

	event := OperatorUpdatedEvent{Previous: decoded.PreviousOperator, New: decoded.NewOperator}
	require.Equal(t, prev, event.Previous)
	require.Equal(t, next, event.New)
	require.Equal(t, liquiditypool.LiquidityPoolOperatorUpdatedTopic, log.Topics[0])
}

func TestSelectors_MatchABI(t *testing.T) {
	parsed, err := abi.JSON(strings.NewReader(liquiditypool.LiquidityPoolABI))
	require.NoError(t, err)
	require.Equal(t, []byte{0x26, 0x0b, 0x60, 0x86}, parsed.Methods["fundChannel"].ID)
	require.Equal(t, []byte{0xfd, 0x3d, 0x31, 0x99}, parsed.Methods["settleChannel"].ID)
}

func TestClassifyUnknownIsDeterministic(t *testing.T) {
	err := classifyEthError(errors.New("some deterministic issue"))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}
