// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package assetholdereth

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

// AssetHolderWithdrawalAuth is an auto generated low-level Go binding around an user-defined struct.
type AssetHolderWithdrawalAuth struct {
	ChannelID   [32]byte
	Participant ChannelParticipant
	Receiver    common.Address
	Amount      *big.Int
}

// ChannelParticipant is an auto generated low-level Go binding around an user-defined struct.
type ChannelParticipant struct {
	EthAddress common.Address
	CcAddress  []byte
}

// AssetholderethMetaData contains all meta data concerning the Assetholdereth contract.
var AssetholderethMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adjudicator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"}],\"name\":\"OutcomeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adjudicator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"holdings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"parts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBals\",\"type\":\"uint256[]\"}],\"name\":\"setOutcome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"settled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant\",\"name\":\"participant\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAssetHolder.WithdrawalAuth\",\"name\":\"authorization\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161105b38038061105b83398101604081905261002f91610054565b600280546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610fc8806100936000396000f3fe6080604052600436106100555760003560e01c80631de26e161461005a578063295482ce1461006f57806353c2ed8e1461008f578063ae9ee18c146100cc578063d945af1d14610107578063fca0f77814610147575b600080fd5b61006d610068366004610c14565b610167565b005b34801561007b57600080fd5b5061006d61008a366004610c7b565b6101d7565b34801561009b57600080fd5b506002546100af906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100d857600080fd5b506100f96100e7366004610cf5565b60006020819052908152604090205481565b6040519081526020016100c3565b34801561011357600080fd5b50610137610122366004610cf5565b60016020526000908152604090205460ff1681565b60405190151581526020016100c3565b34801561015357600080fd5b5061006d610162366004610d0e565b6104fb565b6101718282610711565b60008281526020819052604090205461018a9082610764565b600083815260208190526040902055817fcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9826040516101cb91815260200190565b60405180910390a25050565b6002546001600160a01b031633146102445760405162461bcd60e51b815260206004820152602560248201527f63616e206f6e6c792062652063616c6c6564206279207468652061646a75646960448201526431b0ba37b960d91b60648201526084015b60405180910390fd5b8281146102a55760405162461bcd60e51b815260206004820152602960248201527f7061727469636970616e7473206c656e6774682073686f756c6420657175616c6044820152682062616c616e63657360b81b606482015260840161023b565b60008581526001602052604090205460ff16156103125760405162461bcd60e51b815260206004820152602560248201527f747279696e6720746f2073657420616c726561647920736574746c6564206368604482015264185b9b995b60da1b606482015260840161023b565b600085815260208190526040812080549082905590808567ffffffffffffffff81111561034157610341610dad565b60405190808252806020026020018201604052801561036a578160200160208202803683370190505b50905060005b8681101561043e5760006103b58a8a8a8581811061039057610390610dc3565b90506020028101906103a29190610dd9565b6103b0906020810190610e0e565b610777565b9050808383815181106103ca576103ca610dc3565b6020026020010181815250506103fb600080838152602001908152602001600020548661076490919063ffffffff16565b945061042887878481811061041257610412610dc3565b905060200201358561076490919063ffffffff16565b935050808061043690610e41565b915050610370565b508183106104ae5760005b868110156104ac5785858281811061046357610463610dc3565b9050602002013560008084848151811061047f5761047f610dc3565b602002602001015181526020019081526020016000208190555080806104a490610e41565b915050610449565b505b6000888152600160208190526040808320805460ff19169092179091555189917fef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b891a25050505050505050565b823560009081526001602052604090205460ff166105515760405162461bcd60e51b815260206004820152601360248201527218da185b9b995b081b9bdd081cd95d1d1b1959606a1b604482015260640161023b565b6105c4836040516020016105659190610e6a565b60408051601f198184030181526020601f8601819004810284018101909252848352919085908590819084018382808284376000920191909152506105b1925050506020870187610dd9565b6105bf906020810190610e0e565b6107bc565b6106105760405162461bcd60e51b815260206004820152601d60248201527f7369676e617475726520766572696669636174696f6e206661696c6564000000604482015260640161023b565b600061062484356103a26020870187610dd9565b6000818152602081905260409020549091506060850135111561067e5760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b604482015260640161023b565b60008181526020819052604090205461069b906060860135610845565b6000828152602081905260409020556106b5848484610851565b807fd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81606086018035906106eb9060408901610e0e565b604080519283526001600160a01b0390911660208301520160405180910390a250505050565b8034146107605760405162461bcd60e51b815260206004820152601f60248201527f77726f6e6720616d6f756e74206f662045544820666f72206465706f73697400604482015260640161023b565b5050565b60006107708284610f4d565b9392505050565b6000828260405160200161079e9291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120905092915050565b60008061081d85805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050600061082b82866108a3565b6001600160a01b0390811690851614925050509392505050565b60006107708284610f65565b6108616060840160408501610e0e565b6001600160a01b03166108fc84606001359081150290604051600060405180830381858888f1935050505015801561089d573d6000803e3d6000fd5b50505050565b60008060006108b285856108c7565b915091506108bf81610935565b509392505050565b60008082516041036108fd5760208301516040840151606085015160001a6108f187828585610aee565b9450945050505061092e565b8251604003610926576020830151604084015161091b868383610bdb565b93509350505061092e565b506000905060025b9250929050565b600081600481111561094957610949610f7c565b036109515750565b600181600481111561096557610965610f7c565b036109b25760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161023b565b60028160048111156109c6576109c6610f7c565b03610a135760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161023b565b6003816004811115610a2757610a27610f7c565b03610a7f5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161023b565b6004816004811115610a9357610a93610f7c565b03610aeb5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161023b565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610b255750600090506003610bd2565b8460ff16601b14158015610b3d57508460ff16601c14155b15610b4e5750600090506004610bd2565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610ba2573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610bcb57600060019250925050610bd2565b9150600090505b94509492505050565b6000806001600160ff1b03831681610bf860ff86901c601b610f4d565b9050610c0687828885610aee565b935093505050935093915050565b60008060408385031215610c2757600080fd5b50508035926020909101359150565b60008083601f840112610c4857600080fd5b50813567ffffffffffffffff811115610c6057600080fd5b6020830191508360208260051b850101111561092e57600080fd5b600080600080600060608688031215610c9357600080fd5b85359450602086013567ffffffffffffffff80821115610cb257600080fd5b610cbe89838a01610c36565b90965094506040880135915080821115610cd757600080fd5b50610ce488828901610c36565b969995985093965092949392505050565b600060208284031215610d0757600080fd5b5035919050565b600080600060408486031215610d2357600080fd5b833567ffffffffffffffff80821115610d3b57600080fd5b9085019060808288031215610d4f57600080fd5b90935060208501359080821115610d6557600080fd5b818601915086601f830112610d7957600080fd5b813581811115610d8857600080fd5b876020828501011115610d9a57600080fd5b6020830194508093505050509250925092565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60008235603e19833603018112610def57600080fd5b9190910192915050565b6001600160a01b0381168114610aeb57600080fd5b600060208284031215610e2057600080fd5b813561077081610df9565b634e487b7160e01b600052601160045260246000fd5b600060018201610e5357610e53610e2b565b5060010190565b8035610e6581610df9565b919050565b602081528135602082015260006020830135603e19843603018112610e8e57600080fd5b6080604084015283018035610ea281610df9565b6001600160a01b031660a0840152602081013536829003601e19018112610ec857600080fd5b0160208101903567ffffffffffffffff811115610ee457600080fd5b803603821315610ef357600080fd5b604060c08501528060e08501526101008183828701376000818387010152610f1d60408701610e5a565b6001600160a01b03811660608701529250606095909501356080850152601f01601f191690920190920192915050565b60008219821115610f6057610f60610e2b565b500190565b600082821015610f7757610f77610e2b565b500390565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220dea45cdf7eabfb0362894605f2b04ce28ac37ab4c50adf6c360fd904a7c797c364736f6c634300080f0033",
}

// AssetholderethABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetholderethMetaData.ABI instead.
var AssetholderethABI = AssetholderethMetaData.ABI

// AssetholderethBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssetholderethMetaData.Bin instead.
var AssetholderethBin = AssetholderethMetaData.Bin

// DeployAssetholdereth deploys a new Ethereum contract, binding an instance of Assetholdereth to it.
func DeployAssetholdereth(auth *bind.TransactOpts, backend bind.ContractBackend, _adjudicator common.Address) (common.Address, *types.Transaction, *Assetholdereth, error) {
	parsed, err := AssetholderethMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssetholderethBin), backend, _adjudicator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Assetholdereth{AssetholderethCaller: AssetholderethCaller{contract: contract}, AssetholderethTransactor: AssetholderethTransactor{contract: contract}, AssetholderethFilterer: AssetholderethFilterer{contract: contract}}, nil
}

// Assetholdereth is an auto generated Go binding around an Ethereum contract.
type Assetholdereth struct {
	AssetholderethCaller     // Read-only binding to the contract
	AssetholderethTransactor // Write-only binding to the contract
	AssetholderethFilterer   // Log filterer for contract events
}

// AssetholderethCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetholderethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetholderethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetholderethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetholderethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetholderethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetholderethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetholderethSession struct {
	Contract     *Assetholdereth   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetholderethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetholderethCallerSession struct {
	Contract *AssetholderethCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AssetholderethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetholderethTransactorSession struct {
	Contract     *AssetholderethTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AssetholderethRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetholderethRaw struct {
	Contract *Assetholdereth // Generic contract binding to access the raw methods on
}

// AssetholderethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetholderethCallerRaw struct {
	Contract *AssetholderethCaller // Generic read-only contract binding to access the raw methods on
}

// AssetholderethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetholderethTransactorRaw struct {
	Contract *AssetholderethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetholdereth creates a new instance of Assetholdereth, bound to a specific deployed contract.
func NewAssetholdereth(address common.Address, backend bind.ContractBackend) (*Assetholdereth, error) {
	contract, err := bindAssetholdereth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Assetholdereth{AssetholderethCaller: AssetholderethCaller{contract: contract}, AssetholderethTransactor: AssetholderethTransactor{contract: contract}, AssetholderethFilterer: AssetholderethFilterer{contract: contract}}, nil
}

// NewAssetholderethCaller creates a new read-only instance of Assetholdereth, bound to a specific deployed contract.
func NewAssetholderethCaller(address common.Address, caller bind.ContractCaller) (*AssetholderethCaller, error) {
	contract, err := bindAssetholdereth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetholderethCaller{contract: contract}, nil
}

// NewAssetholderethTransactor creates a new write-only instance of Assetholdereth, bound to a specific deployed contract.
func NewAssetholderethTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetholderethTransactor, error) {
	contract, err := bindAssetholdereth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetholderethTransactor{contract: contract}, nil
}

// NewAssetholderethFilterer creates a new log filterer instance of Assetholdereth, bound to a specific deployed contract.
func NewAssetholderethFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetholderethFilterer, error) {
	contract, err := bindAssetholdereth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetholderethFilterer{contract: contract}, nil
}

// bindAssetholdereth binds a generic wrapper to an already deployed contract.
func bindAssetholdereth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetholderethMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assetholdereth *AssetholderethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assetholdereth.Contract.AssetholderethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assetholdereth *AssetholderethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assetholdereth.Contract.AssetholderethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assetholdereth *AssetholderethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assetholdereth.Contract.AssetholderethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assetholdereth *AssetholderethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assetholdereth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assetholdereth *AssetholderethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assetholdereth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assetholdereth *AssetholderethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assetholdereth.Contract.contract.Transact(opts, method, params...)
}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdereth *AssetholderethCaller) Adjudicator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Assetholdereth.contract.Call(opts, &out, "adjudicator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdereth *AssetholderethSession) Adjudicator() (common.Address, error) {
	return _Assetholdereth.Contract.Adjudicator(&_Assetholdereth.CallOpts)
}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdereth *AssetholderethCallerSession) Adjudicator() (common.Address, error) {
	return _Assetholdereth.Contract.Adjudicator(&_Assetholdereth.CallOpts)
}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdereth *AssetholderethCaller) Holdings(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Assetholdereth.contract.Call(opts, &out, "holdings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdereth *AssetholderethSession) Holdings(arg0 [32]byte) (*big.Int, error) {
	return _Assetholdereth.Contract.Holdings(&_Assetholdereth.CallOpts, arg0)
}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdereth *AssetholderethCallerSession) Holdings(arg0 [32]byte) (*big.Int, error) {
	return _Assetholdereth.Contract.Holdings(&_Assetholdereth.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdereth *AssetholderethCaller) Settled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Assetholdereth.contract.Call(opts, &out, "settled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdereth *AssetholderethSession) Settled(arg0 [32]byte) (bool, error) {
	return _Assetholdereth.Contract.Settled(&_Assetholdereth.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdereth *AssetholderethCallerSession) Settled(arg0 [32]byte) (bool, error) {
	return _Assetholdereth.Contract.Settled(&_Assetholdereth.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdereth *AssetholderethTransactor) Deposit(opts *bind.TransactOpts, fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdereth.contract.Transact(opts, "deposit", fundingID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdereth *AssetholderethSession) Deposit(fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdereth.Contract.Deposit(&_Assetholdereth.TransactOpts, fundingID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdereth *AssetholderethTransactorSession) Deposit(fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdereth.Contract.Deposit(&_Assetholdereth.TransactOpts, fundingID, amount)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x295482ce.
//
// Solidity: function setOutcome(bytes32 channelID, (address,bytes)[] parts, uint256[] newBals) returns()
func (_Assetholdereth *AssetholderethTransactor) SetOutcome(opts *bind.TransactOpts, channelID [32]byte, parts []ChannelParticipant, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdereth.contract.Transact(opts, "setOutcome", channelID, parts, newBals)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x295482ce.
//
// Solidity: function setOutcome(bytes32 channelID, (address,bytes)[] parts, uint256[] newBals) returns()
func (_Assetholdereth *AssetholderethSession) SetOutcome(channelID [32]byte, parts []ChannelParticipant, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdereth.Contract.SetOutcome(&_Assetholdereth.TransactOpts, channelID, parts, newBals)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x295482ce.
//
// Solidity: function setOutcome(bytes32 channelID, (address,bytes)[] parts, uint256[] newBals) returns()
func (_Assetholdereth *AssetholderethTransactorSession) SetOutcome(channelID [32]byte, parts []ChannelParticipant, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdereth.Contract.SetOutcome(&_Assetholdereth.TransactOpts, channelID, parts, newBals)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfca0f778.
//
// Solidity: function withdraw((bytes32,(address,bytes),address,uint256) authorization, bytes signature) returns()
func (_Assetholdereth *AssetholderethTransactor) Withdraw(opts *bind.TransactOpts, authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdereth.contract.Transact(opts, "withdraw", authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfca0f778.
//
// Solidity: function withdraw((bytes32,(address,bytes),address,uint256) authorization, bytes signature) returns()
func (_Assetholdereth *AssetholderethSession) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdereth.Contract.Withdraw(&_Assetholdereth.TransactOpts, authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfca0f778.
//
// Solidity: function withdraw((bytes32,(address,bytes),address,uint256) authorization, bytes signature) returns()
func (_Assetholdereth *AssetholderethTransactorSession) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdereth.Contract.Withdraw(&_Assetholdereth.TransactOpts, authorization, signature)
}

// AssetholderethDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Assetholdereth contract.
type AssetholderethDepositedIterator struct {
	Event *AssetholderethDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetholderethDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetholderethDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetholderethDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetholderethDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetholderethDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetholderethDeposited represents a Deposited event raised by the Assetholdereth contract.
type AssetholderethDeposited struct {
	FundingID [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assetholdereth *AssetholderethFilterer) FilterDeposited(opts *bind.FilterOpts, fundingID [][32]byte) (*AssetholderethDepositedIterator, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdereth.contract.FilterLogs(opts, "Deposited", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return &AssetholderethDepositedIterator{contract: _Assetholdereth.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assetholdereth *AssetholderethFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *AssetholderethDeposited, fundingID [][32]byte) (event.Subscription, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdereth.contract.WatchLogs(opts, "Deposited", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetholderethDeposited)
				if err := _Assetholdereth.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assetholdereth *AssetholderethFilterer) ParseDeposited(log types.Log) (*AssetholderethDeposited, error) {
	event := new(AssetholderethDeposited)
	if err := _Assetholdereth.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetholderethOutcomeSetIterator is returned from FilterOutcomeSet and is used to iterate over the raw logs and unpacked data for OutcomeSet events raised by the Assetholdereth contract.
type AssetholderethOutcomeSetIterator struct {
	Event *AssetholderethOutcomeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetholderethOutcomeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetholderethOutcomeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetholderethOutcomeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetholderethOutcomeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetholderethOutcomeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetholderethOutcomeSet represents a OutcomeSet event raised by the Assetholdereth contract.
type AssetholderethOutcomeSet struct {
	ChannelID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutcomeSet is a free log retrieval operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assetholdereth *AssetholderethFilterer) FilterOutcomeSet(opts *bind.FilterOpts, channelID [][32]byte) (*AssetholderethOutcomeSetIterator, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Assetholdereth.contract.FilterLogs(opts, "OutcomeSet", channelIDRule)
	if err != nil {
		return nil, err
	}
	return &AssetholderethOutcomeSetIterator{contract: _Assetholdereth.contract, event: "OutcomeSet", logs: logs, sub: sub}, nil
}

// WatchOutcomeSet is a free log subscription operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assetholdereth *AssetholderethFilterer) WatchOutcomeSet(opts *bind.WatchOpts, sink chan<- *AssetholderethOutcomeSet, channelID [][32]byte) (event.Subscription, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Assetholdereth.contract.WatchLogs(opts, "OutcomeSet", channelIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetholderethOutcomeSet)
				if err := _Assetholdereth.contract.UnpackLog(event, "OutcomeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutcomeSet is a log parse operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assetholdereth *AssetholderethFilterer) ParseOutcomeSet(log types.Log) (*AssetholderethOutcomeSet, error) {
	event := new(AssetholderethOutcomeSet)
	if err := _Assetholdereth.contract.UnpackLog(event, "OutcomeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetholderethWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Assetholdereth contract.
type AssetholderethWithdrawnIterator struct {
	Event *AssetholderethWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssetholderethWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetholderethWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssetholderethWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssetholderethWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetholderethWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetholderethWithdrawn represents a Withdrawn event raised by the Assetholdereth contract.
type AssetholderethWithdrawn struct {
	FundingID [32]byte
	Amount    *big.Int
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assetholdereth *AssetholderethFilterer) FilterWithdrawn(opts *bind.FilterOpts, fundingID [][32]byte) (*AssetholderethWithdrawnIterator, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdereth.contract.FilterLogs(opts, "Withdrawn", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return &AssetholderethWithdrawnIterator{contract: _Assetholdereth.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assetholdereth *AssetholderethFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *AssetholderethWithdrawn, fundingID [][32]byte) (event.Subscription, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdereth.contract.WatchLogs(opts, "Withdrawn", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetholderethWithdrawn)
				if err := _Assetholdereth.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assetholdereth *AssetholderethFilterer) ParseWithdrawn(log types.Log) (*AssetholderethWithdrawn, error) {
	event := new(AssetholderethWithdrawn)
	if err := _Assetholdereth.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
