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
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e0038038062000e008339810160408190526200003491620002cc565b6040518060400160405280600a8152602001692832b93ab72a37b5b2b760b11b8152506040518060400160405280600381526020016228292760e91b81525081600390805190602001906200008b92919062000218565b508051620000a190600490602084019062000218565b50506005805460ff191660121790555060005b8251811015620000ed57620000e4838281518110620000cf57fe5b602002602001015183620000f660201b60201c565b600101620000b4565b50505062000427565b6001600160a01b038216620001285760405162461bcd60e51b81526004016200011f90620003c0565b60405180910390fd5b6200013660008383620001e2565b6200015281600254620001e760201b6200044b1790919060201c565b6002556001600160a01b03821660009081526020818152604090912054620001859183906200044b620001e7821b17901c565b6001600160a01b0383166000818152602081905260408082209390935591519091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90620001d6908590620003f7565b60405180910390a35050565b505050565b6000828201838110156200020f5760405162461bcd60e51b81526004016200011f9062000389565b90505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200025b57805160ff19168380011785556200028b565b828001600101855582156200028b579182015b828111156200028b5782518255916020019190600101906200026e565b50620002999291506200029d565b5090565b5b808211156200029957600081556001016200029e565b80516001600160a01b03811681146200021257600080fd5b60008060408385031215620002df578182fd5b82516001600160401b0380821115620002f6578384fd5b818501915085601f8301126200030a578384fd5b81518181111562000319578485fd5b602091508181026200032d83820162000400565b8281528381019085850183870186018b101562000348578889fd5b8896505b848710156200037657620003618b82620002b4565b8352600196909601959185019185016200034c565b5097909301519698969750505050505050565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252601f908201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604082015260600190565b90815260200190565b6040518181016001600160401b03811182821017156200041f57600080fd5b604052919050565b6109c980620004376000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012957806370a082311461013c57806395d89b411461014f578063a457c2d714610157578063a9059cbb1461016a578063dd62ed3e1461017d576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ec57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b6610190565b6040516100c39190610759565b60405180910390f35b6100df6100da366004610724565b610226565b6040516100c3919061074e565b6100f4610244565b6040516100c391906108f1565b6100df61010f3660046106e4565b61024a565b61011c6102d1565b6040516100c391906108fa565b6100df610137366004610724565b6102da565b6100f461014a366004610695565b610328565b6100b6610343565b6100df610165366004610724565b6103a4565b6100df610178366004610724565b61040c565b6100f461018b3660046106b0565b610420565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561021c5780601f106101f15761010080835404028352916020019161021c565b820191906000526020600020905b8154815290600101906020018083116101ff57829003601f168201915b5050505050905090565b600061023a610233610480565b8484610484565b5060015b92915050565b60025490565b6000610257848484610538565b6102c784610263610480565b6102c285604051806060016040528060288152602001610947602891396001600160a01b038a166000908152600160205260408120906102a1610480565b6001600160a01b03168152602081019190915260400160002054919061064d565b610484565b5060019392505050565b60055460ff1690565b600061023a6102e7610480565b846102c285600160006102f8610480565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549061044b565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561021c5780601f106101f15761010080835404028352916020019161021c565b600061023a6103b1610480565b846102c28560405180606001604052806025815260200161096f60259139600160006103db610480565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919061064d565b600061023a610419610480565b8484610538565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6000828201838110156104795760405162461bcd60e51b815260040161047090610831565b60405180910390fd5b9392505050565b3390565b6001600160a01b0383166104aa5760405162461bcd60e51b8152600401610470906108ad565b6001600160a01b0382166104d05760405162461bcd60e51b8152600401610470906107ef565b6001600160a01b0380841660008181526001602090815260408083209487168084529490915290819020849055517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259061052b9085906108f1565b60405180910390a3505050565b6001600160a01b03831661055e5760405162461bcd60e51b815260040161047090610868565b6001600160a01b0382166105845760405162461bcd60e51b8152600401610470906107ac565b61058f838383610679565b6105cc81604051806060016040528060268152602001610921602691396001600160a01b038616600090815260208190526040902054919061064d565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546105fb908261044b565b6001600160a01b0380841660008181526020819052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061052b9085906108f1565b600081848411156106715760405162461bcd60e51b81526004016104709190610759565b505050900390565b505050565b80356001600160a01b038116811461023e57600080fd5b6000602082840312156106a6578081fd5b610479838361067e565b600080604083850312156106c2578081fd5b6106cc848461067e565b91506106db846020850161067e565b90509250929050565b6000806000606084860312156106f8578081fd5b833561070381610908565b9250602084013561071381610908565b929592945050506040919091013590565b60008060408385031215610736578182fd5b610740848461067e565b946020939093013593505050565b901515815260200190565b6000602080835283518082850152825b8181101561078557858101830151858201604001528201610769565b818111156107965783604083870101525b50601f01601f1916929092016040019392505050565b60208082526023908201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526022908201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604082015261737360f01b606082015260800190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b60208082526025908201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526024908201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646040820152637265737360e01b606082015260800190565b90815260200190565b60ff91909116815260200190565b6001600160a01b038116811461091d57600080fd5b5056fe45524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122096a65db3b71c1d5b775f4216129ef7e5120fbacbf536805c3f48919682c8c9b364736f6c63430007000033",
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
