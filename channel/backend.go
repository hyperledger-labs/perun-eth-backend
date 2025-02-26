// Copyright 2025 - See NOTICE file for copyright holders.
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
	"log"
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
)

const (
	phaseDispute = iota
	phaseForceExec
	phaseConcluded
	ethereumAddressLength = 20
)

var (
	// compile time check that we implement the channel backend interface.
	_ channel.Backend = new(Backend)
	// Definition of ABI datatypes.
	abiAddress, _ = abi.NewType("address", "", nil)
	abiBytes32, _ = abi.NewType("bytes32", "", nil)
	abiParams     abi.Type
	abiState      abi.Type
	abiProgress   abi.Method
	abiRegister   abi.Method
	// MaxBalance is the maximum amount of funds per asset that a user can possess.
	// It is set to 2 ^ 256 - 1.
	MaxBalance = abi.MaxUint256
)

func init() {
	// The ABI type parser is unable to parse the correct params and state types,
	// therefore we fetch them from the function signatures.
	adj, err := abi.JSON(strings.NewReader(adjudicator.AdjudicatorABI))
	if err != nil {
		panic("decoding ABI json")
	}
	// Get the Params type.
	chID, ok := adj.Methods["channelID"]
	if !ok || len(chID.Inputs) != 1 {
		panic("channelID not found")
	}
	abiParams = chID.Inputs[0].Type
	// Get the State type.
	hashState, ok := adj.Methods["hashState"]
	if !ok || len(hashState.Inputs) != 1 {
		panic("hashState not found")
	}
	abiState = hashState.Inputs[0].Type

	if abiProgress, ok = adj.Methods["progress"]; !ok {
		panic("Could not find method progress in adjudicator contract.")
	}

	if abiRegister, ok = adj.Methods["register"]; !ok {
		panic("Could not find method register in adjudicator contract.")
	}
}

// Backend implements the interface defined in channel/Backend.go.
type Backend struct{}

// CalcID calculates the channelID as needed by the ethereum smart contracts.
func (*Backend) CalcID(p *channel.Params) (id channel.ID, err error) {
	return CalcID(p)
}

// NewAppID creates a new app identifier, using an empty ethereum struct.
func (b *Backend) NewAppID() (channel.AppID, error) {
	addr := &ethwallet.Address{}
	return &AppID{addr}, nil
}

// Sign signs the channel state as needed by the ethereum smart contracts.
func (*Backend) Sign(acc wallet.Account, s *channel.State) (wallet.Sig, error) {
	return Sign(acc, s)
}

// Verify verifies that a state was signed correctly.
func (*Backend) Verify(addr wallet.Address, s *channel.State, sig wallet.Sig) (bool, error) {
	return Verify(addr, s, sig)
}

// NewAsset returns a variable of type Asset, which can be used
// for unmarshalling an asset from its binary representation.
func (b *Backend) NewAsset() channel.Asset {
	return &Asset{}
}

// CalcID calculates the channelID as needed by the ethereum smart contracts.
func CalcID(p *channel.Params) (id channel.ID, err error) {
	params := ToEthParams(p)
	bytes, err := EncodeParams(&params)
	if err != nil {
		return id, errors.WithMessage(err, "could not encode parameters")
	}
	// Hash encoded params.
	return crypto.Keccak256Hash(bytes), nil
}

// HashState calculates the hash of a state as needed by the ethereum smart contracts.
func HashState(s *channel.State) (id channel.ID) {
	state := ToEthState(s)
	bytes, err := EncodeState(&state)
	if err != nil {
		log.Panicf("could not encode parameters: %v", err)
	}
	return crypto.Keccak256Hash(bytes)
}

// Sign signs the channel state as needed by the ethereum smart contracts.
func Sign(acc wallet.Account, s *channel.State) (wallet.Sig, error) {
	state := ToEthState(s)
	enc, err := EncodeState(&state)
	if err != nil {
		return nil, errors.WithMessage(err, "encoding state")
	}
	return acc.SignData(enc)
}

// Verify verifies that a state was signed correctly.
func Verify(addr wallet.Address, s *channel.State, sig wallet.Sig) (bool, error) {
	state := ToEthState(s)
	enc, err := EncodeState(&state)
	if err != nil {
		return false, errors.WithMessage(err, "encoding state")
	}
	return ethwallet.VerifySignature(enc, sig, addr)
}

// ToEthParams converts a channel.Params to a ChannelParams struct.
func ToEthParams(p *channel.Params) adjudicator.ChannelParams {
	var app common.Address
	if p.App != nil && !channel.IsNoApp(p.App) {
		appDef, ok := p.App.Def().(channel.AppID)
		ethAddress, err := ExtractEthereumAddress(appDef)
		if err != nil {
			log.Panicf("error extracting Ethereum address: %v", err)
		}
		if !ok {
			panic("appDef is not of type channel.AppID")
		}
		app = ethAddress
	}

	return adjudicator.ChannelParams{
		ChallengeDuration: new(big.Int).SetUint64(p.ChallengeDuration),
		Nonce:             p.Nonce,
		App:               app,
		Participants:      pwToCommonAddresses(p.Parts),
		LedgerChannel:     p.LedgerChannel,
		VirtualChannel:    p.VirtualChannel,
	}
}

// ExtractEthereumAddress extracts an Ethereum address from the given channel.AppID.
func ExtractEthereumAddress(appID channel.AppID) (common.Address, error) {
	// Marshal the AppID into bytes
	addressBytes, err := appID.MarshalBinary()
	if err != nil {
		return common.Address{}, errors.WithMessage(err, "failed to marshal AppID")
	}

	// Ensure the byte length is correct for an Ethereum address (20 bytes)
	if len(addressBytes) != ethereumAddressLength {
		return common.Address{}, errors.WithMessagef(err, "invalid length for Ethereum address: %d", len(addressBytes))
	}

	// Convert the byte slice to an Ethereum address
	ethAddress := common.BytesToAddress(addressBytes)
	return ethAddress, nil
}

// ToEthState converts a channel.State to a ChannelState struct.
func ToEthState(s *channel.State) adjudicator.ChannelState {
	backends := make([]*big.Int, len(s.Allocation.Assets))
	for i := range s.Allocation.Assets { // we assume that for each asset there is an element in backends corresponding to the backendID the asset belongs to.
		backends[i] = big.NewInt(int64(s.Allocation.Backends[i]))
	}
	locked := make([]adjudicator.ChannelSubAlloc, len(s.Locked))
	for i, sub := range s.Locked {
		// Create index map.
		indexMap := make([]uint16, s.NumParts())
		if len(sub.IndexMap) == 0 {
			for i := range indexMap {
				indexMap[i] = uint16(i)
			}
		} else {
			for i, x := range sub.IndexMap {
				indexMap[i] = uint16(x)
			}
		}

		locked[i] = adjudicator.ChannelSubAlloc{ID: sub.ID, Balances: sub.Bals, IndexMap: indexMap}
	}
	outcome := adjudicator.ChannelAllocation{
		Assets:   assetsToEthAssets(s.Allocation.Assets, s.Allocation.Backends),
		Backends: backends,
		Balances: s.Balances,
		Locked:   locked,
	}
	// Check allocation dimensions
	if len(outcome.Assets) != len(outcome.Balances) || len(s.Balances) != len(outcome.Balances) {
		log.Panic("invalid allocation dimensions")
	}
	appData, err := s.Data.MarshalBinary()
	if err != nil {
		log.Panicf("error encoding app data: %v", err)
	}
	return adjudicator.ChannelState{
		ChannelID: s.ID,
		Version:   s.Version,
		Outcome:   outcome,
		AppData:   appData,
		IsFinal:   s.IsFinal,
	}
}

// EncodeParams encodes the parameters as with abi.encode() in the smart contracts.
func EncodeParams(params *adjudicator.ChannelParams) ([]byte, error) {
	args := abi.Arguments{{Type: abiParams}}
	enc, err := args.Pack(*params)
	return enc, errors.WithStack(err)
}

// EncodeState encodes the state as with abi.encode() in the smart contracts.
func EncodeState(state *adjudicator.ChannelState) ([]byte, error) {
	args := abi.Arguments{{Type: abiState}}
	enc, err := args.Pack(*state)
	return enc, errors.WithStack(err)
}

// pwToCommonAddresses converts an array of perun/ethwallet.Addresses to common.Addresses.
func pwToCommonAddresses(addr []map[wallet.BackendID]wallet.Address) []adjudicator.ChannelParticipant {
	cAddrs := make([]adjudicator.ChannelParticipant, len(addr))
	for i, part := range addr {
		ethAddr, ok := part[ethwallet.BackendID]
		if !ok {
			log.Panic("eth address not found")
		}
		cAddrs[i].EthAddress = ethwallet.AsEthAddr(ethAddr)
		for backendID, walletAddr := range part {
			if backendID == ethwallet.BackendID {
				continue // Skip Ethereum address
			}

			addBytes, err := walletAddr.MarshalBinary()
			if err != nil {
				log.Panicf("error encoding address for backend %d: %v", backendID, err)
			}
			cAddrs[i].CcAddress = addBytes
			break // Take only the first non-eth address
		}

		// If no other addresses exist, initialize CcAddress with 32 zero bytes
		if cAddrs[i].CcAddress == nil {
			cAddrs[i].CcAddress = make([]byte, 32) //nolint:gomnd
		}
	}
	return cAddrs
}

// FromEthState converts a ChannelState to a channel.State struct.
func FromEthState(app channel.App, s *adjudicator.ChannelState) channel.State {
	backends := make([]wallet.BackendID, len(s.Outcome.Backends))
	for i, b := range s.Outcome.Backends {
		backends[i] = wallet.BackendID(b.Int64())
	}
	locked := make([]channel.SubAlloc, len(s.Outcome.Locked))
	for i, sub := range s.Outcome.Locked {
		indexMap := makeIndexMap(sub.IndexMap)
		locked[i] = *channel.NewSubAlloc(sub.ID, sub.Balances, indexMap)
	}
	alloc := channel.Allocation{
		Assets:   fromEthAssets(s.Outcome.Assets, backends),
		Balances: s.Outcome.Balances,
		Backends: backends,
		Locked:   locked,
	}
	// Check allocation dimensions
	if len(alloc.Assets) != len(alloc.Balances) || len(s.Outcome.Balances) != len(alloc.Balances) || len(alloc.Backends) != len(alloc.Assets) {
		log.Panic("invalid allocation dimensions")
	}

	data := app.NewData()
	if err := data.UnmarshalBinary(s.AppData); err != nil {
		log.Panicf("decoding app data: %v", err)
	}

	return channel.State{
		ID:         s.ChannelID,
		Version:    s.Version,
		Allocation: alloc,
		App:        app,
		Data:       data,
		IsFinal:    s.IsFinal,
	}
}

func makeIndexMap(m []uint16) []channel.Index {
	_m := make([]channel.Index, len(m))
	for i, x := range m {
		_m[i] = channel.Index(x)
	}
	return _m
}

// assetsToEthAssets converts an array of Assets to adjudicator.ChannelAsset.
func assetsToEthAssets(assets []channel.Asset, bIDs []wallet.BackendID) []adjudicator.ChannelAsset {
	cAddrs := make([]adjudicator.ChannelAsset, len(assets))
	for i, a := range assets {
		// This means the Asset was defined in this backend.
		if bIDs[i] == ethwallet.BackendID {
			asset, ok := a.(*Asset)
			if !ok {
				log.Panicf("wrong address type: %T", a)
			}
			cAddrs[i] = adjudicator.ChannelAsset{
				ChainID:   asset.assetID.ChainID(),
				EthHolder: asset.EthAddress(),
				CcHolder:  make([]byte, 32), //nolint:gomnd
			}
		} else {
			asset, err := a.MarshalBinary()
			if err != nil {
				log.Panicf("error encoding asset: %v", err)
			}
			cAddrs[i] = adjudicator.ChannelAsset{
				ChainID:   big.NewInt(int64(bIDs[i])),
				EthHolder: common.HexToAddress("0x0000000000000000000000000000000000000000"),
				CcHolder:  asset,
			}
		}
	}
	return cAddrs
}

func fromEthAssets(assets []adjudicator.ChannelAsset, bIDs []wallet.BackendID) []channel.Asset {
	_assets := make([]channel.Asset, len(assets))
	for i, a := range assets {
		if bIDs[i] == ethwallet.BackendID {
			_assets[i] = NewAsset(a.ChainID, a.EthHolder)
		} else {
			_assets[i] = &Asset{}
			err := _assets[i].UnmarshalBinary(a.CcHolder)
			if err != nil {
				log.Panicf("error decoding asset: %v", err)
			}
		}
	}
	return _assets
}
