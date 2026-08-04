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
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Subscription aliases the go-ethereum subscription contract.
type Subscription = ethereum.Subscription

// ETHFunder is called by Hub/WebsocketBackend to lock ETH in the pool.
type ETHFunder interface {
	FundChannel(ctx context.Context, channelID [32]byte, amount *big.Int) error
}

// ETHSettler is called by Hub/WebsocketBackend to return ETH to the pool.
type ETHSettler interface {
	SettleChannel(ctx context.Context, channelID [32]byte, returnAmount *big.Int) error
}

// ETHPoolReader exposes read-only pool queries to orchestration layers.
type ETHPoolReader interface {
	GetOperator(ctx context.Context) (common.Address, error)
	GetPoolState(ctx context.Context) (reserve uint64, locked uint64, err error)
}

// ETHLPReader exposes read-only LP-owner queries for the liquidity dashboard.
// Deposit/withdraw stay non-custodial (the LP signs in their own wallet), so this
// interface intentionally exposes reads only.
type ETHLPReader interface {
	SharesOf(ctx context.Context, owner common.Address) (*big.Int, error)
	TotalShares(ctx context.Context) (*big.Int, error)
	WithdrawableETH(ctx context.Context) (*big.Int, error)
	PreviewWithdrawETH(ctx context.Context, shares *big.Int) (*big.Int, error)
	PoolMetadata() (addr common.Address, chainID *big.Int)
}

// ETHEventSubscriber exposes typed event subscriptions for reconciliation.
type ETHEventSubscriber interface {
	SubscribeChannelFunded(ctx context.Context, ch chan<- ChannelFundedEvent) (Subscription, error)
	SubscribeChannelSettled(ctx context.Context, ch chan<- ChannelSettledEvent) (Subscription, error)
	SubscribeOperatorUpdated(ctx context.Context, ch chan<- OperatorUpdatedEvent) (Subscription, error)
}

// ChannelFundedEvent mirrors LiquidityPool.ChannelFunded.
type ChannelFundedEvent struct {
	ChannelID [32]byte
	Principal *big.Int
	Operator  common.Address
	BlockNum  uint64
}

// ChannelSettledEvent mirrors LiquidityPool.ChannelSettled.
type ChannelSettledEvent struct {
	ChannelID     [32]byte
	Principal     *big.Int
	TotalReturned *big.Int
	FeeGain       *big.Int
	BlockNum      uint64
}

// OperatorUpdatedEvent mirrors LiquidityPool.OperatorUpdated.
type OperatorUpdatedEvent struct {
	Previous common.Address
	New      common.Address
}
