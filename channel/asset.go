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

package channel

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/perun-network/perun-eth-backend/bindings/assetholder"
	"github.com/perun-network/perun-eth-backend/bindings/assetholdererc20"
	"github.com/perun-network/perun-eth-backend/bindings/assetholdereth"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
	"github.com/perun-network/perun-eth-backend/wallet"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/channel/multi"
	"perun.network/go-perun/wire/perunio"
)

// ChainID identifies a specific Ethereum backend.
type ChainID struct {
	*big.Int
}

func MakeChainID(id *big.Int) ChainID {
	if id.Sign() < 0 {
		panic("must not be smaller than zero")
	}
	return ChainID{id}
}

func (id *ChainID) UnmarshalBinary(data []byte) error {
	id.Int = new(big.Int).SetBytes(data)
	return nil
}

func (id ChainID) MarshalBinary() (data []byte, err error) {
	return id.Bytes(), nil
}

func (id ChainID) MapKey() multi.LedgerIDMapKey {
	return multi.LedgerIDMapKey(id.Int.String())
}

type (

	// Asset is an Ethereum asset.
	Asset struct {
		ChainID     ChainID
		AssetHolder wallet.Address
	}

	AssetMapKey string
)

func (a Asset) MapKey() AssetMapKey {
	d, err := a.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return AssetMapKey(d)
}

func (a Asset) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := perunio.Encode(&buf, &a.AssetHolder, a.ChainID)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (a Asset) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)
	return perunio.Decode(buf, &a.ChainID, &a.AssetHolder)
}

func (a Asset) LedgerID() multi.LedgerID {
	return &a.ChainID
}

// NewAsset creates a new asset from an chainID and the AssetHolder address.
func NewAsset(chainID *big.Int, assetHolder common.Address) *Asset {
	id := MakeChainID(chainID)
	return &Asset{id, *wallet.AsWalletAddr(assetHolder)}
}

// EthAddress returns the Ethereum address of the asset.
func (a Asset) EthAddress() common.Address {
	return common.Address(a.AssetHolder)
}

// Equal returns true iff the asset equals the given asset.
func (a Asset) Equal(b channel.Asset) bool {
	ethAsset, ok := b.(*Asset)
	if !ok {
		return false
	}
	return a.EthAddress() == ethAsset.EthAddress()
}

// FilterAssets filters the assets for the given chainID.
func FilterAssets(assets []channel.Asset, chainID ChainID) []channel.Asset {
	var filtered []channel.Asset
	for _, asset := range assets {
		if a := asset.(*Asset); a.ChainID.MapKey() == chainID.MapKey() {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

// GetAssetIdx returns the index of asset in the assets array.
func GetAssetIdx(assets []channel.Asset, asset channel.Asset) channel.Index {
	for i, a := range assets {
		if a.Equal(asset) {
			return channel.Index(i)
		}
	}
	panic("asset not found")
}

var _ channel.Asset = new(Asset)

// ValidateAssetHolderETH checks if the bytecode at the given asset holder ETH
// address is correct and if the adjudicator address is correctly set in the
// asset holder contract. The contract code at the adjudicator address is not
// validated, it is the user's responsibility to provide a valid adjudicator
// address.
//
// Returns a ContractBytecodeError if the bytecode is invalid. This error can
// be checked with function IsErrInvalidContractCode.
func ValidateAssetHolderETH(ctx context.Context,
	backend bind.ContractBackend, assetHolderETH, adjudicator common.Address) error {
	return validateAssetHolder(ctx, backend, assetHolderETH, adjudicator,
		assetholdereth.AssetHolderETHBinRuntime)
}

// ValidateAssetHolderERC20 checks if the bytecode at the given asset holder
// ERC20 address is correct and if the adjudicator address is correctly set in
// the asset holder contract. The contract code at the adjudicator address is
// not validated, it is the user's responsibility to provide a valid
// adjudicator address.
//
// Returns a ContractBytecodeError if the bytecode is invalid. This error can
// be checked with function IsErrInvalidContractCode.
func ValidateAssetHolderERC20(ctx context.Context,
	backend bind.ContractBackend, assetHolderERC20, adjudicator, token common.Address) error {
	return validateAssetHolder(ctx, backend, assetHolderERC20, adjudicator,
		assetHolderERC20BinRuntimeFor(token))
}

func validateAssetHolder(ctx context.Context,
	backend bind.ContractBackend, assetHolderAddr, adjudicatorAddr common.Address, bytecode string) error {
	if err := validateContract(ctx, backend, assetHolderAddr, bytecode); err != nil {
		return errors.WithMessage(err, "validating asset holder")
	}

	assetHolder, err := assetholder.NewAssetholder(assetHolderAddr, backend)
	if err != nil {
		return errors.Wrap(err, "binding AssetHolder")
	}
	opts := bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	if addrSetInContract, err := assetHolder.Adjudicator(&opts); err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return errors.WithMessage(err, "fetching adjudicator address set in asset holder contract")
	} else if addrSetInContract != adjudicatorAddr {
		return errors.Wrap(ErrInvalidContractCode, "incorrect adjudicator code")
	}

	return nil
}

func validateContract(ctx context.Context,
	backend bind.ContractCaller, contract common.Address, bytecode string) error {
	code, err := backend.CodeAt(ctx, contract, nil)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return errors.WithMessage(err, "fetching contract code")
	}
	if hex.EncodeToString(code) != bytecode {
		return errors.Wrap(ErrInvalidContractCode, "incorrect contract code")
	}
	return nil
}

func assetHolderERC20BinRuntimeFor(token common.Address) string {
	// runtimePlaceholder indicates constructor variables in runtime binary code.
	const runtimePlaceholder = "7f0000000000000000000000000000000000000000000000000000000000000000"

	tokenHex := hex.EncodeToString(token[:])
	return strings.ReplaceAll(assetholdererc20.AssetHolderERC20BinRuntime,
		runtimePlaceholder,
		runtimePlaceholder[:len(runtimePlaceholder)-len(tokenHex)]+tokenHex)
}
