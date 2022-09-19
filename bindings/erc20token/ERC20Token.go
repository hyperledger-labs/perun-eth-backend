// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20token

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

// Erc20tokenMetaData contains all meta data concerning the Erc20token contract.
var Erc20tokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"code\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e3038038062000e3083398101604081905262000034916200027c565b8383600362000044838262000427565b50600462000053828262000427565b50505060005b8251811015620000a657620000918382815181106200007c576200007c620004f3565b602002602001015183620000b160201b60201c565b806200009d816200051f565b91505062000059565b505050505062000556565b6001600160a01b0382166200010c5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200012091906200053b565b90915550506001600160a01b038216600090815260208190526040812080548392906200014f9084906200053b565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620001df57620001df6200019e565b604052919050565b600082601f830112620001f957600080fd5b81516001600160401b038111156200021557620002156200019e565b60206200022b601f8301601f19168201620001b4565b82815285828487010111156200024057600080fd5b60005b838110156200026057858101830151828201840152820162000243565b83811115620002725760008385840101525b5095945050505050565b600080600080608085870312156200029357600080fd5b84516001600160401b0380821115620002ab57600080fd5b620002b988838901620001e7565b9550602091508187015181811115620002d157600080fd5b620002df89828a01620001e7565b955050604087015181811115620002f557600080fd5b8701601f810189136200030757600080fd5b8051828111156200031c576200031c6200019e565b8060051b92506200032f848401620001b4565b818152928201840192848101908b8511156200034a57600080fd5b928501925b848410156200038457835192506001600160a01b0383168314620003735760008081fd5b82825292850192908501906200034f565b60609a909a0151989b979a5050505050505050565b600181811c90821680620003ae57607f821691505b602082108103620003cf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200019957600081815260208120601f850160051c81016020861015620003fe5750805b601f850160051c820191505b818110156200041f578281556001016200040a565b505050505050565b81516001600160401b038111156200044357620004436200019e565b6200045b8162000454845462000399565b84620003d5565b602080601f8311600181146200049357600084156200047a5750858301515b600019600386901b1c1916600185901b1785556200041f565b600085815260208120601f198616915b82811015620004c457888601518255948401946001909101908401620004a3565b5085821015620004e35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820162000534576200053462000509565b5060010190565b6000821982111562000551576200055162000509565b500190565b6108ca80620005666000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b411461015f578063a457c2d714610167578063a9059cbb1461017a578063dd62ed3e1461018d57600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101c6565b6040516100c39190610708565b60405180910390f35b6100df6100da366004610779565b610258565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f3660046107a3565b61026e565b604051601281526020016100c3565b6100df610131366004610779565b610324565b6100f36101443660046107df565b6001600160a01b031660009081526020819052604090205490565b6100b6610360565b6100df610175366004610779565b61036f565b6100df610188366004610779565b610408565b6100f361019b366004610801565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101d590610834565b80601f016020809104026020016040519081016040528092919081815260200182805461020190610834565b801561024e5780601f106102235761010080835404028352916020019161024e565b820191906000526020600020905b81548152906001019060200180831161023157829003601f168201915b5050505050905090565b6000610265338484610415565b50600192915050565b6001600160a01b0383166000908152600160209081526040808320338452909152812054600019811461030e57828110156103015760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b61030e8533858403610415565b610319858585610539565b506001949350505050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161026591859061035b90869061086e565b610415565b6060600480546101d590610834565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156103f15760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016102f8565b6103fe3385858403610415565b5060019392505050565b6000610265338484610539565b6001600160a01b0383166104775760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016102f8565b6001600160a01b0382166104d85760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016102f8565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b03831661059d5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016102f8565b6001600160a01b0382166105ff5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016102f8565b6001600160a01b038316600090815260208190526040902054818110156106775760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016102f8565b6001600160a01b038085166000908152602081905260408082208585039055918516815290812080548492906106ae90849061086e565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106fa91815260200190565b60405180910390a350505050565b600060208083528351808285015260005b8181101561073557858101830151858201604001528201610719565b81811115610747576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461077457600080fd5b919050565b6000806040838503121561078c57600080fd5b6107958361075d565b946020939093013593505050565b6000806000606084860312156107b857600080fd5b6107c18461075d565b92506107cf6020850161075d565b9150604084013590509250925092565b6000602082840312156107f157600080fd5b6107fa8261075d565b9392505050565b6000806040838503121561081457600080fd5b61081d8361075d565b915061082b6020840161075d565b90509250929050565b600181811c9082168061084857607f821691505b60208210810361086857634e487b7160e01b600052602260045260246000fd5b50919050565b6000821982111561088f57634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220bb923ec46558bfb5ef2dc156a8fe5a2e3a94bb98af38953a9b454925a083f0c264736f6c634300080f0033",
}

// Erc20tokenABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20tokenMetaData.ABI instead.
var Erc20tokenABI = Erc20tokenMetaData.ABI

// Erc20tokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20tokenMetaData.Bin instead.
var Erc20tokenBin = Erc20tokenMetaData.Bin

// DeployErc20token deploys a new Ethereum contract, binding an instance of Erc20token to it.
func DeployErc20token(auth *bind.TransactOpts, backend bind.ContractBackend, name string, code string, accounts []common.Address, initBalance *big.Int) (common.Address, *types.Transaction, *Erc20token, error) {
	parsed, err := Erc20tokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20tokenBin), backend, name, code, accounts, initBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20token{Erc20tokenCaller: Erc20tokenCaller{contract: contract}, Erc20tokenTransactor: Erc20tokenTransactor{contract: contract}, Erc20tokenFilterer: Erc20tokenFilterer{contract: contract}}, nil
}

// Erc20token is an auto generated Go binding around an Ethereum contract.
type Erc20token struct {
	Erc20tokenCaller     // Read-only binding to the contract
	Erc20tokenTransactor // Write-only binding to the contract
	Erc20tokenFilterer   // Log filterer for contract events
}

// Erc20tokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20tokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20tokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20tokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20tokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20tokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20tokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20tokenSession struct {
	Contract     *Erc20token       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20tokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20tokenCallerSession struct {
	Contract *Erc20tokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Erc20tokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20tokenTransactorSession struct {
	Contract     *Erc20tokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Erc20tokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20tokenRaw struct {
	Contract *Erc20token // Generic contract binding to access the raw methods on
}

// Erc20tokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20tokenCallerRaw struct {
	Contract *Erc20tokenCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20tokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20tokenTransactorRaw struct {
	Contract *Erc20tokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20token creates a new instance of Erc20token, bound to a specific deployed contract.
func NewErc20token(address common.Address, backend bind.ContractBackend) (*Erc20token, error) {
	contract, err := bindErc20token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20token{Erc20tokenCaller: Erc20tokenCaller{contract: contract}, Erc20tokenTransactor: Erc20tokenTransactor{contract: contract}, Erc20tokenFilterer: Erc20tokenFilterer{contract: contract}}, nil
}

// NewErc20tokenCaller creates a new read-only instance of Erc20token, bound to a specific deployed contract.
func NewErc20tokenCaller(address common.Address, caller bind.ContractCaller) (*Erc20tokenCaller, error) {
	contract, err := bindErc20token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20tokenCaller{contract: contract}, nil
}

// NewErc20tokenTransactor creates a new write-only instance of Erc20token, bound to a specific deployed contract.
func NewErc20tokenTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20tokenTransactor, error) {
	contract, err := bindErc20token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20tokenTransactor{contract: contract}, nil
}

// NewErc20tokenFilterer creates a new log filterer instance of Erc20token, bound to a specific deployed contract.
func NewErc20tokenFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20tokenFilterer, error) {
	contract, err := bindErc20token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20tokenFilterer{contract: contract}, nil
}

// bindErc20token binds a generic wrapper to an already deployed contract.
func bindErc20token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20tokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20token *Erc20tokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20token.Contract.Erc20tokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20token *Erc20tokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20token.Contract.Erc20tokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20token *Erc20tokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20token.Contract.Erc20tokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20token *Erc20tokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20token *Erc20tokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20token *Erc20tokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20token *Erc20tokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20token *Erc20tokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20token.Contract.Allowance(&_Erc20token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20token *Erc20tokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20token.Contract.Allowance(&_Erc20token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20token *Erc20tokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20token *Erc20tokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20token.Contract.BalanceOf(&_Erc20token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20token *Erc20tokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20token.Contract.BalanceOf(&_Erc20token.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20token *Erc20tokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20token *Erc20tokenSession) Decimals() (uint8, error) {
	return _Erc20token.Contract.Decimals(&_Erc20token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20token *Erc20tokenCallerSession) Decimals() (uint8, error) {
	return _Erc20token.Contract.Decimals(&_Erc20token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20token *Erc20tokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20token *Erc20tokenSession) Name() (string, error) {
	return _Erc20token.Contract.Name(&_Erc20token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20token *Erc20tokenCallerSession) Name() (string, error) {
	return _Erc20token.Contract.Name(&_Erc20token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20token *Erc20tokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20token *Erc20tokenSession) Symbol() (string, error) {
	return _Erc20token.Contract.Symbol(&_Erc20token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20token *Erc20tokenCallerSession) Symbol() (string, error) {
	return _Erc20token.Contract.Symbol(&_Erc20token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20token *Erc20tokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20token *Erc20tokenSession) TotalSupply() (*big.Int, error) {
	return _Erc20token.Contract.TotalSupply(&_Erc20token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20token *Erc20tokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20token.Contract.TotalSupply(&_Erc20token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.Approve(&_Erc20token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.Approve(&_Erc20token.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20token *Erc20tokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20token *Erc20tokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.DecreaseAllowance(&_Erc20token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Erc20token *Erc20tokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.DecreaseAllowance(&_Erc20token.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20token *Erc20tokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20token *Erc20tokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.IncreaseAllowance(&_Erc20token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Erc20token *Erc20tokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.IncreaseAllowance(&_Erc20token.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.Transfer(&_Erc20token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.Transfer(&_Erc20token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.TransferFrom(&_Erc20token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20token *Erc20tokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20token.Contract.TransferFrom(&_Erc20token.TransactOpts, sender, recipient, amount)
}

// Erc20tokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20token contract.
type Erc20tokenApprovalIterator struct {
	Event *Erc20tokenApproval // Event containing the contract specifics and raw log

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
func (it *Erc20tokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20tokenApproval)
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
		it.Event = new(Erc20tokenApproval)
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
func (it *Erc20tokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20tokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20tokenApproval represents a Approval event raised by the Erc20token contract.
type Erc20tokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20token *Erc20tokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20tokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20tokenApprovalIterator{contract: _Erc20token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20token *Erc20tokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20tokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20tokenApproval)
				if err := _Erc20token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc20token *Erc20tokenFilterer) ParseApproval(log types.Log) (*Erc20tokenApproval, error) {
	event := new(Erc20tokenApproval)
	if err := _Erc20token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20tokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20token contract.
type Erc20tokenTransferIterator struct {
	Event *Erc20tokenTransfer // Event containing the contract specifics and raw log

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
func (it *Erc20tokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20tokenTransfer)
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
		it.Event = new(Erc20tokenTransfer)
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
func (it *Erc20tokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20tokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20tokenTransfer represents a Transfer event raised by the Erc20token contract.
type Erc20tokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20token *Erc20tokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20tokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20tokenTransferIterator{contract: _Erc20token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20token *Erc20tokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20tokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20tokenTransfer)
				if err := _Erc20token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc20token *Erc20tokenFilterer) ParseTransfer(log types.Log) (*Erc20tokenTransfer, error) {
	event := new(Erc20tokenTransfer)
	if err := _Erc20token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
