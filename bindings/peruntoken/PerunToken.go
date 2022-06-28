// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package peruntoken

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
)

// PeruntokenMetaData contains all meta data concerning the Peruntoken contract.
var PeruntokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000d1138038062000d118339810160408190526200003491620002be565b6040518060400160405280600a8152602001692832b93ab72a37b5b2b760b11b8152506040518060400160405280600381526020016228292760e91b81525081600390805190602001906200008b929190620001e5565b508051620000a1906004906020840190620001e5565b50505060005b8251811015620000f457620000df838281518110620000ca57620000ca62000398565b602002602001015183620000fd60201b60201c565b80620000eb81620003c4565b915050620000a7565b50505062000437565b6001600160a01b038216620001585760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200016c9190620003e0565b90915550506001600160a01b038216600090815260208190526040812080548392906200019b908490620003e0565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b828054620001f390620003fb565b90600052602060002090601f01602090048101928262000217576000855562000262565b82601f106200023257805160ff191683800117855562000262565b8280016001018555821562000262579182015b828111156200026257825182559160200191906001019062000245565b506200027092915062000274565b5090565b5b8082111562000270576000815560010162000275565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b0381168114620002b957600080fd5b919050565b60008060408385031215620002d257600080fd5b82516001600160401b0380821115620002ea57600080fd5b818501915085601f830112620002ff57600080fd5b81516020828211156200031657620003166200028b565b8160051b604051601f19603f830116810181811086821117156200033e576200033e6200028b565b6040529283528183019350848101820192898411156200035d57600080fd5b948201945b8386101562000386576200037686620002a1565b8552948201949382019362000362565b97909101519698969750505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201620003d957620003d9620003ae565b5060010190565b60008219821115620003f657620003f6620003ae565b500190565b600181811c908216806200041057607f821691505b6020821081036200043157634e487b7160e01b600052602260045260246000fd5b50919050565b6108ca80620004476000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b411461015f578063a457c2d714610167578063a9059cbb1461017a578063dd62ed3e1461018d57600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101c6565b6040516100c39190610708565b60405180910390f35b6100df6100da366004610779565b610258565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f3660046107a3565b61026e565b604051601281526020016100c3565b6100df610131366004610779565b610324565b6100f36101443660046107df565b6001600160a01b031660009081526020819052604090205490565b6100b6610360565b6100df610175366004610779565b61036f565b6100df610188366004610779565b610408565b6100f361019b366004610801565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101d590610834565b80601f016020809104026020016040519081016040528092919081815260200182805461020190610834565b801561024e5780601f106102235761010080835404028352916020019161024e565b820191906000526020600020905b81548152906001019060200180831161023157829003601f168201915b5050505050905090565b6000610265338484610415565b50600192915050565b6001600160a01b0383166000908152600160209081526040808320338452909152812054600019811461030e57828110156103015760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b61030e8533858403610415565b610319858585610539565b506001949350505050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161026591859061035b90869061086e565b610415565b6060600480546101d590610834565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156103f15760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016102f8565b6103fe3385858403610415565b5060019392505050565b6000610265338484610539565b6001600160a01b0383166104775760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016102f8565b6001600160a01b0382166104d85760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016102f8565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b03831661059d5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016102f8565b6001600160a01b0382166105ff5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016102f8565b6001600160a01b038316600090815260208190526040902054818110156106775760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016102f8565b6001600160a01b038085166000908152602081905260408082208585039055918516815290812080548492906106ae90849061086e565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106fa91815260200190565b60405180910390a350505050565b600060208083528351808285015260005b8181101561073557858101830151858201604001528201610719565b81811115610747576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461077457600080fd5b919050565b6000806040838503121561078c57600080fd5b6107958361075d565b946020939093013593505050565b6000806000606084860312156107b857600080fd5b6107c18461075d565b92506107cf6020850161075d565b9150604084013590509250925092565b6000602082840312156107f157600080fd5b6107fa8261075d565b9392505050565b6000806040838503121561081457600080fd5b61081d8361075d565b915061082b6020840161075d565b90509250929050565b600181811c9082168061084857607f821691505b60208210810361086857634e487b7160e01b600052602260045260246000fd5b50919050565b6000821982111561088f57634e487b7160e01b600052601160045260246000fd5b50019056fea26469706673582212201797e6f3a2c42869c2547bb748953b2d48de1120e66a8667d9a1fc31694b9d8f64736f6c634300080e0033",
}

// PeruntokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PeruntokenMetaData.ABI instead.
var PeruntokenABI = PeruntokenMetaData.ABI

// PeruntokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PeruntokenMetaData.Bin instead.
var PeruntokenBin = PeruntokenMetaData.Bin

// DeployPeruntoken deploys a new Ethereum contract, binding an instance of Peruntoken to it.
func DeployPeruntoken(auth *bind.TransactOpts, backend bind.ContractBackend, accounts []common.Address, initBalance *big.Int) (common.Address, *types.Transaction, *Peruntoken, error) {
	parsed, err := PeruntokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PeruntokenBin), backend, accounts, initBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Peruntoken{PeruntokenCaller: PeruntokenCaller{contract: contract}, PeruntokenTransactor: PeruntokenTransactor{contract: contract}, PeruntokenFilterer: PeruntokenFilterer{contract: contract}}, nil
}

// Peruntoken is an auto generated Go binding around an Ethereum contract.
type Peruntoken struct {
	PeruntokenCaller     // Read-only binding to the contract
	PeruntokenTransactor // Write-only binding to the contract
	PeruntokenFilterer   // Log filterer for contract events
}

// PeruntokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeruntokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeruntokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeruntokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeruntokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeruntokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeruntokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeruntokenSession struct {
	Contract     *Peruntoken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeruntokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeruntokenCallerSession struct {
	Contract *PeruntokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PeruntokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeruntokenTransactorSession struct {
	Contract     *PeruntokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PeruntokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeruntokenRaw struct {
	Contract *Peruntoken // Generic contract binding to access the raw methods on
}

// PeruntokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeruntokenCallerRaw struct {
	Contract *PeruntokenCaller // Generic read-only contract binding to access the raw methods on
}

// PeruntokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeruntokenTransactorRaw struct {
	Contract *PeruntokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeruntoken creates a new instance of Peruntoken, bound to a specific deployed contract.
func NewPeruntoken(address common.Address, backend bind.ContractBackend) (*Peruntoken, error) {
	contract, err := bindPeruntoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Peruntoken{PeruntokenCaller: PeruntokenCaller{contract: contract}, PeruntokenTransactor: PeruntokenTransactor{contract: contract}, PeruntokenFilterer: PeruntokenFilterer{contract: contract}}, nil
}

// NewPeruntokenCaller creates a new read-only instance of Peruntoken, bound to a specific deployed contract.
func NewPeruntokenCaller(address common.Address, caller bind.ContractCaller) (*PeruntokenCaller, error) {
	contract, err := bindPeruntoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeruntokenCaller{contract: contract}, nil
}

// NewPeruntokenTransactor creates a new write-only instance of Peruntoken, bound to a specific deployed contract.
func NewPeruntokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PeruntokenTransactor, error) {
	contract, err := bindPeruntoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeruntokenTransactor{contract: contract}, nil
}

// NewPeruntokenFilterer creates a new log filterer instance of Peruntoken, bound to a specific deployed contract.
func NewPeruntokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PeruntokenFilterer, error) {
	contract, err := bindPeruntoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeruntokenFilterer{contract: contract}, nil
}

// bindPeruntoken binds a generic wrapper to an already deployed contract.
func bindPeruntoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeruntokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peruntoken *PeruntokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Peruntoken.Contract.PeruntokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peruntoken *PeruntokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peruntoken.Contract.PeruntokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peruntoken *PeruntokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peruntoken.Contract.PeruntokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peruntoken *PeruntokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Peruntoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peruntoken *PeruntokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peruntoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peruntoken *PeruntokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peruntoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Peruntoken *PeruntokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Peruntoken *PeruntokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Peruntoken.Contract.Allowance(&_Peruntoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Peruntoken *PeruntokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Peruntoken.Contract.Allowance(&_Peruntoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Peruntoken *PeruntokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Peruntoken *PeruntokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Peruntoken.Contract.BalanceOf(&_Peruntoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Peruntoken *PeruntokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Peruntoken.Contract.BalanceOf(&_Peruntoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Peruntoken *PeruntokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Peruntoken *PeruntokenSession) Decimals() (uint8, error) {
	return _Peruntoken.Contract.Decimals(&_Peruntoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Peruntoken *PeruntokenCallerSession) Decimals() (uint8, error) {
	return _Peruntoken.Contract.Decimals(&_Peruntoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peruntoken *PeruntokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peruntoken *PeruntokenSession) Name() (string, error) {
	return _Peruntoken.Contract.Name(&_Peruntoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peruntoken *PeruntokenCallerSession) Name() (string, error) {
	return _Peruntoken.Contract.Name(&_Peruntoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peruntoken *PeruntokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peruntoken *PeruntokenSession) Symbol() (string, error) {
	return _Peruntoken.Contract.Symbol(&_Peruntoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peruntoken *PeruntokenCallerSession) Symbol() (string, error) {
	return _Peruntoken.Contract.Symbol(&_Peruntoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peruntoken *PeruntokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peruntoken *PeruntokenSession) TotalSupply() (*big.Int, error) {
	return _Peruntoken.Contract.TotalSupply(&_Peruntoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peruntoken *PeruntokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Peruntoken.Contract.TotalSupply(&_Peruntoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.Approve(&_Peruntoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.Approve(&_Peruntoken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Peruntoken *PeruntokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Peruntoken *PeruntokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.DecreaseAllowance(&_Peruntoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.DecreaseAllowance(&_Peruntoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Peruntoken *PeruntokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Peruntoken *PeruntokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.IncreaseAllowance(&_Peruntoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.IncreaseAllowance(&_Peruntoken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.Transfer(&_Peruntoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.Transfer(&_Peruntoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.TransferFrom(&_Peruntoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.TransferFrom(&_Peruntoken.TransactOpts, sender, recipient, amount)
}

// PeruntokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Peruntoken contract.
type PeruntokenApprovalIterator struct {
	Event *PeruntokenApproval // Event containing the contract specifics and raw log

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
func (it *PeruntokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeruntokenApproval)
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
		it.Event = new(PeruntokenApproval)
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
func (it *PeruntokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeruntokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeruntokenApproval represents a Approval event raised by the Peruntoken contract.
type PeruntokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Peruntoken *PeruntokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PeruntokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Peruntoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PeruntokenApprovalIterator{contract: _Peruntoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Peruntoken *PeruntokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PeruntokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Peruntoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeruntokenApproval)
				if err := _Peruntoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Peruntoken *PeruntokenFilterer) ParseApproval(log types.Log) (*PeruntokenApproval, error) {
	event := new(PeruntokenApproval)
	if err := _Peruntoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeruntokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Peruntoken contract.
type PeruntokenTransferIterator struct {
	Event *PeruntokenTransfer // Event containing the contract specifics and raw log

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
func (it *PeruntokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeruntokenTransfer)
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
		it.Event = new(PeruntokenTransfer)
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
func (it *PeruntokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeruntokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeruntokenTransfer represents a Transfer event raised by the Peruntoken contract.
type PeruntokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Peruntoken *PeruntokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PeruntokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Peruntoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PeruntokenTransferIterator{contract: _Peruntoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Peruntoken *PeruntokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PeruntokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Peruntoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeruntokenTransfer)
				if err := _Peruntoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Peruntoken *PeruntokenFilterer) ParseTransfer(log types.Log) (*PeruntokenTransfer, error) {
	event := new(PeruntokenTransfer)
	if err := _Peruntoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
