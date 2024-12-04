// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trivialapp

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ChannelAllocation is an auto generated low-level Go binding around an user-defined struct.
type ChannelAllocation struct {
	Assets   []ChannelAsset
	Backends []*big.Int
	Balances [][]*big.Int
	Locked   []ChannelSubAlloc
}

// ChannelAsset is an auto generated low-level Go binding around an user-defined struct.
type ChannelAsset struct {
	ChainID   *big.Int
	EthHolder common.Address
	CcHolder  []byte
}

// ChannelParams is an auto generated low-level Go binding around an user-defined struct.
type ChannelParams struct {
	ChallengeDuration *big.Int
	Nonce             *big.Int
	Participants      []ChannelParticipant
	App               common.Address
	LedgerChannel     bool
	VirtualChannel    bool
}

// ChannelParticipant is an auto generated low-level Go binding around an user-defined struct.
type ChannelParticipant struct {
	EthAddress common.Address
	CcAddress  []byte
}

// ChannelState is an auto generated low-level Go binding around an user-defined struct.
type ChannelState struct {
	ChannelID [][32]byte
	Version   uint64
	Outcome   ChannelAllocation
	AppData   []byte
	IsFinal   bool
}

// ChannelSubAlloc is an auto generated low-level Go binding around an user-defined struct.
type ChannelSubAlloc struct {
	ID       [][32]byte
	Balances []*big.Int
	IndexMap []uint16
}

// TrivialappMetaData contains all meta data concerning the Trivialapp contract.
var TrivialappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"channelID\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"ID\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"channelID\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"ID\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"actorIdx\",\"type\":\"uint256\"}],\"name\":\"validTransition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061011c806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806318dc6c0014602d575b600080fd5b603e60383660046057565b50505050565b005b600060a08284031215605157600080fd5b50919050565b60008060008060808587031215606c57600080fd5b843567ffffffffffffffff80821115608357600080fd5b9086019060c08289031215609657600080fd5b9094506020860135908082111560ab57600080fd5b60b5888389016040565b9450604087013591508082111560ca57600080fd5b5060d5878288016040565b94979396509394606001359350505056fea26469706673582212205081e797202cb035c07d4df0f6988f47b842ec24f720e8dc1336efef5813300764736f6c634300080f0033",
}

// TrivialappABI is the input ABI used to generate the binding from.
// Deprecated: Use TrivialappMetaData.ABI instead.
var TrivialappABI = TrivialappMetaData.ABI

// TrivialappBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TrivialappMetaData.Bin instead.
var TrivialappBin = TrivialappMetaData.Bin

// DeployTrivialapp deploys a new Ethereum contract, binding an instance of Trivialapp to it.
func DeployTrivialapp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trivialapp, error) {
	parsed, err := TrivialappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TrivialappBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trivialapp{TrivialappCaller: TrivialappCaller{contract: contract}, TrivialappTransactor: TrivialappTransactor{contract: contract}, TrivialappFilterer: TrivialappFilterer{contract: contract}}, nil
}

// Trivialapp is an auto generated Go binding around an Ethereum contract.
type Trivialapp struct {
	TrivialappCaller     // Read-only binding to the contract
	TrivialappTransactor // Write-only binding to the contract
	TrivialappFilterer   // Log filterer for contract events
}

// TrivialappCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrivialappCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrivialappTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrivialappTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrivialappFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrivialappFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrivialappSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrivialappSession struct {
	Contract     *Trivialapp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrivialappCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrivialappCallerSession struct {
	Contract *TrivialappCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TrivialappTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrivialappTransactorSession struct {
	Contract     *TrivialappTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TrivialappRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrivialappRaw struct {
	Contract *Trivialapp // Generic contract binding to access the raw methods on
}

// TrivialappCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrivialappCallerRaw struct {
	Contract *TrivialappCaller // Generic read-only contract binding to access the raw methods on
}

// TrivialappTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrivialappTransactorRaw struct {
	Contract *TrivialappTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrivialapp creates a new instance of Trivialapp, bound to a specific deployed contract.
func NewTrivialapp(address common.Address, backend bind.ContractBackend) (*Trivialapp, error) {
	contract, err := bindTrivialapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trivialapp{TrivialappCaller: TrivialappCaller{contract: contract}, TrivialappTransactor: TrivialappTransactor{contract: contract}, TrivialappFilterer: TrivialappFilterer{contract: contract}}, nil
}

// NewTrivialappCaller creates a new read-only instance of Trivialapp, bound to a specific deployed contract.
func NewTrivialappCaller(address common.Address, caller bind.ContractCaller) (*TrivialappCaller, error) {
	contract, err := bindTrivialapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrivialappCaller{contract: contract}, nil
}

// NewTrivialappTransactor creates a new write-only instance of Trivialapp, bound to a specific deployed contract.
func NewTrivialappTransactor(address common.Address, transactor bind.ContractTransactor) (*TrivialappTransactor, error) {
	contract, err := bindTrivialapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrivialappTransactor{contract: contract}, nil
}

// NewTrivialappFilterer creates a new log filterer instance of Trivialapp, bound to a specific deployed contract.
func NewTrivialappFilterer(address common.Address, filterer bind.ContractFilterer) (*TrivialappFilterer, error) {
	contract, err := bindTrivialapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrivialappFilterer{contract: contract}, nil
}

// bindTrivialapp binds a generic wrapper to an already deployed contract.
func bindTrivialapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrivialappMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trivialapp *TrivialappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trivialapp.Contract.TrivialappCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trivialapp *TrivialappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trivialapp.Contract.TrivialappTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trivialapp *TrivialappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trivialapp.Contract.TrivialappTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trivialapp *TrivialappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trivialapp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trivialapp *TrivialappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trivialapp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trivialapp *TrivialappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trivialapp.Contract.contract.Transact(opts, method, params...)
}

// ValidTransition is a free data retrieval call binding the contract method 0x18dc6c00.
//
// Solidity: function validTransition((uint256,uint256,(address,bytes)[],address,bool,bool) params, (bytes32[],uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32[],uint256[],uint16[])[]),bytes,bool) from, (bytes32[],uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32[],uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_Trivialapp *TrivialappCaller) ValidTransition(opts *bind.CallOpts, params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	var out []interface{}
	err := _Trivialapp.contract.Call(opts, &out, "validTransition", params, from, to, actorIdx)

	if err != nil {
		return err
	}

	return err

}

// ValidTransition is a free data retrieval call binding the contract method 0x18dc6c00.
//
// Solidity: function validTransition((uint256,uint256,(address,bytes)[],address,bool,bool) params, (bytes32[],uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32[],uint256[],uint16[])[]),bytes,bool) from, (bytes32[],uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32[],uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_Trivialapp *TrivialappSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	return _Trivialapp.Contract.ValidTransition(&_Trivialapp.CallOpts, params, from, to, actorIdx)
}

// ValidTransition is a free data retrieval call binding the contract method 0x18dc6c00.
//
// Solidity: function validTransition((uint256,uint256,(address,bytes)[],address,bool,bool) params, (bytes32[],uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32[],uint256[],uint16[])[]),bytes,bool) from, (bytes32[],uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32[],uint256[],uint16[])[]),bytes,bool) to, uint256 actorIdx) pure returns()
func (_Trivialapp *TrivialappCallerSession) ValidTransition(params ChannelParams, from ChannelState, to ChannelState, actorIdx *big.Int) error {
	return _Trivialapp.Contract.ValidTransition(&_Trivialapp.CallOpts, params, from, to, actorIdx)
}
