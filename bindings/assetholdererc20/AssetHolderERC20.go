// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package assetholdererc20

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

// AssetHolderWithdrawalAuth is an auto generated low-level Go binding around an user-defined struct.
type AssetHolderWithdrawalAuth struct {
	ChannelID   [32]byte
	Participant common.Address
	Receiver    common.Address
	Amount      *big.Int
}

// Assetholdererc20MetaData contains all meta data concerning the Assetholdererc20 contract.
var Assetholdererc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adjudicator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"}],\"name\":\"OutcomeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adjudicator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"holdings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBals\",\"type\":\"uint256[]\"}],\"name\":\"setOutcome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"settled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAssetHolder.WithdrawalAuth\",\"name\":\"authorization\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161118a38038061118a83398101604081905261002f91610065565b600280546001600160a01b0319166001600160a01b03939093169290921790915560601b6001600160601b0319166080526100b6565b60008060408385031215610077578182fd5b82516100828161009e565b60208401519092506100938161009e565b809150509250929050565b6001600160a01b03811681146100b357600080fd5b50565b60805160601c6110ad6100dd600039806103b6528061065b52806107c152506110ad6000f3fe6080604052600436106100705760003560e01c8063ae9ee18c1161004e578063ae9ee18c146100d5578063d945af1d14610102578063fc0c546a1461012f578063fc79a66d1461014457610070565b80631de26e16146100755780634ed4283c1461008a57806353c2ed8e146100aa575b600080fd5b610088610083366004610b24565b610164565b005b34801561009657600080fd5b506100886100a5366004610b45565b6101dc565b3480156100b657600080fd5b506100bf61037e565b6040516100cc9190610bfb565b60405180910390f35b3480156100e157600080fd5b506100f56100f0366004610a95565b61038d565b6040516100cc9190611056565b34801561010e57600080fd5b5061012261011d366004610a95565b61039f565b6040516100cc9190610c4c565b34801561013b57600080fd5b506100bf6103b4565b34801561015057600080fd5b5061008861015f366004610aad565b6103d8565b61016e82826105f6565b6000828152602081905260409020546101879082610618565b6000838152602081905260409020556101a08282610644565b817fcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9826040516101d09190611056565b60405180910390a25050565b823560009081526001602052604090205460ff166102155760405162461bcd60e51b815260040161020c90610f16565b60405180910390fd5b61027d83604051602001610229919061100a565b60408051601f198184030181526020601f860181900481028401810190925284835291908590859081908401838280828437600092019190915250610278925050506040870160208801610a59565b610702565b6102995760405162461bcd60e51b815260040161020c90610e9d565b60006102b584356102b06040870160208801610a59565b61073d565b600081815260208190526040902054909150606085013511156102ea5760405162461bcd60e51b815260040161020c90610f70565b6102f5848484610770565b600081815260208190526040902054610312906060860135610775565b60008281526020819052604090205561032c8484846107b7565b807fd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81606086018035906103629060408901610a59565b604051610370929190610c57565b60405180910390a250505050565b6002546001600160a01b031681565b60006020819052908152604090205481565b60016020526000908152604090205460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b6002546001600160a01b031633146104025760405162461bcd60e51b815260040161020c90610fc5565b8281146104215760405162461bcd60e51b815260040161020c90610dcd565b60008581526001602052604090205460ff16156104505760405162461bcd60e51b815260040161020c90610e16565b60008581526020819052604081208054908290559060608567ffffffffffffffff8111801561047e57600080fd5b506040519080825280602002602001820160405280156104a8578160200160208202803683370190505b50905060005b8681101561054f5760006104dd8a8a8a858181106104c857fe5b90506020020160208101906102b09190610a59565b9050808383815181106104ec57fe5b60200260200101818152505061051d600080838152602001908152602001600020548661061890919063ffffffff16565b945061054487878481811061052e57fe5b905060200201358561061890919063ffffffff16565b9350506001016104ae565b508183106105a95760005b868110156105a75785858281811061056e57fe5b9050602002013560008084848151811061058457fe5b60209081029190910181015182528101919091526040016000205560010161055a565b505b6000888152600160208190526040808320805460ff19169092179091555189917fef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b891a25050505050505050565b34156106145760405162461bcd60e51b815260040161020c90610d16565b5050565b60008282018381101561063d5760405162461bcd60e51b815260040161020c90610d96565b9392505050565b6040516323b872dd60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd9061069490339030908690600401610c0f565b602060405180830381600087803b1580156106ae57600080fd5b505af11580156106c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e69190610a75565b6106145760405162461bcd60e51b815260040161020c90610f43565b6000806107158580519060200120610886565b9050600061072382866108b6565b6001600160a01b0390811690851614925050509392505050565b60008282604051602001610752929190610c57565b60405160208183030381529060405280519060200120905092915050565b505050565b600061063d83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506109e4565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663a9059cbb6107f66060860160408701610a59565b85606001356040518363ffffffff1660e01b8152600401610818929190610c33565b602060405180830381600087803b15801561083257600080fd5b505af1158015610846573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086a9190610a75565b6107705760405162461bcd60e51b815260040161020c90610f9c565b6000816040516020016108999190610bca565b604051602081830303815290604052805190602001209050919050565b600081516041146108d95760405162461bcd60e51b815260040161020c90610d5f565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561092b5760405162461bcd60e51b815260040161020c90610e5b565b8060ff16601b1415801561094357508060ff16601c14155b156109605760405162461bcd60e51b815260040161020c90610ed4565b6000600187838686604051600081526020016040526040516109859493929190610c6e565b6020604051602081039080840390855afa1580156109a7573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166109da5760405162461bcd60e51b815260040161020c90610cdf565b9695505050505050565b60008184841115610a085760405162461bcd60e51b815260040161020c9190610c8c565b505050900390565b60008083601f840112610a21578182fd5b50813567ffffffffffffffff811115610a38578182fd5b6020830191508360208083028501011115610a5257600080fd5b9250929050565b600060208284031215610a6a578081fd5b813561063d8161105f565b600060208284031215610a86578081fd5b8151801515811461063d578182fd5b600060208284031215610aa6578081fd5b5035919050565b600080600080600060608688031215610ac4578081fd5b85359450602086013567ffffffffffffffff80821115610ae2578283fd5b610aee89838a01610a10565b90965094506040880135915080821115610b06578283fd5b50610b1388828901610a10565b969995985093965092949392505050565b60008060408385031215610b36578182fd5b50508035926020909101359150565b600080600083850360a0811215610b5a578384fd5b6080811215610b67578384fd5b50839250608084013567ffffffffffffffff80821115610b85578384fd5b818601915086601f830112610b98578384fd5b813581811115610ba6578485fd5b876020828501011115610bb7578485fd5b6020830194508093505050509250925092565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b9182526001600160a01b0316602082015260400190565b93845260ff9290921660208401526040830152606082015260800190565b6000602080835283518082850152825b81811015610cb857858101830151858201604001528201610c9c565b81811115610cc95783604083870101525b50601f01601f1916929092016040019392505050565b60208082526018908201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604082015260600190565b60208082526029908201527f6d6573736167652076616c7565206d757374206265203020666f7220746f6b656040820152681b8819195c1bdcda5d60ba1b606082015260800190565b6020808252601f908201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604082015260600190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b60208082526029908201527f7061727469636970616e7473206c656e6774682073686f756c6420657175616c6040820152682062616c616e63657360b81b606082015260800190565b60208082526025908201527f747279696e6720746f2073657420616c726561647920736574746c6564206368604082015264185b9b995b60da1b606082015260800190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604082015261756560f01b606082015260800190565b6020808252601d908201527f7369676e617475726520766572696669636174696f6e206661696c6564000000604082015260600190565b60208082526022908201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604082015261756560f01b606082015260800190565b60208082526013908201527218da185b9b995b081b9bdd081cd95d1d1b1959606a1b604082015260600190565b6020808252601390820152721d1c985b9cd9995c919c9bdb4819985a5b1959606a1b604082015260600190565b602080825260129082015271696e73756666696369656e742066756e647360701b604082015260600190565b6020808252600f908201526e1d1c985b9cd9995c8819985a5b1959608a1b604082015260600190565b60208082526025908201527f63616e206f6e6c792062652063616c6c6564206279207468652061646a75646960408201526431b0ba37b960d91b606082015260800190565b813581526080810160208301356110208161105f565b6001600160a01b03908116602084015260408401359061103f8261105f565b166040830152606092830135929091019190915290565b90815260200190565b6001600160a01b038116811461107457600080fd5b5056fea2646970667358221220a3ebba6f22bf093add0ee45389cbba766d78059df2df41fa7e923ad5264e9c7564736f6c63430007000033",
}

// Assetholdererc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Assetholdererc20MetaData.ABI instead.
var Assetholdererc20ABI = Assetholdererc20MetaData.ABI

// Assetholdererc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Assetholdererc20MetaData.Bin instead.
var Assetholdererc20Bin = Assetholdererc20MetaData.Bin

// DeployAssetholdererc20 deploys a new Ethereum contract, binding an instance of Assetholdererc20 to it.
func DeployAssetholdererc20(auth *bind.TransactOpts, backend bind.ContractBackend, _adjudicator common.Address, _token common.Address) (common.Address, *types.Transaction, *Assetholdererc20, error) {
	parsed, err := Assetholdererc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Assetholdererc20Bin), backend, _adjudicator, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Assetholdererc20{Assetholdererc20Caller: Assetholdererc20Caller{contract: contract}, Assetholdererc20Transactor: Assetholdererc20Transactor{contract: contract}, Assetholdererc20Filterer: Assetholdererc20Filterer{contract: contract}}, nil
}

// Assetholdererc20 is an auto generated Go binding around an Ethereum contract.
type Assetholdererc20 struct {
	Assetholdererc20Caller     // Read-only binding to the contract
	Assetholdererc20Transactor // Write-only binding to the contract
	Assetholdererc20Filterer   // Log filterer for contract events
}

// Assetholdererc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Assetholdererc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Assetholdererc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Assetholdererc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Assetholdererc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Assetholdererc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Assetholdererc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Assetholdererc20Session struct {
	Contract     *Assetholdererc20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Assetholdererc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Assetholdererc20CallerSession struct {
	Contract *Assetholdererc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Assetholdererc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Assetholdererc20TransactorSession struct {
	Contract     *Assetholdererc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Assetholdererc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Assetholdererc20Raw struct {
	Contract *Assetholdererc20 // Generic contract binding to access the raw methods on
}

// Assetholdererc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Assetholdererc20CallerRaw struct {
	Contract *Assetholdererc20Caller // Generic read-only contract binding to access the raw methods on
}

// Assetholdererc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Assetholdererc20TransactorRaw struct {
	Contract *Assetholdererc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetholdererc20 creates a new instance of Assetholdererc20, bound to a specific deployed contract.
func NewAssetholdererc20(address common.Address, backend bind.ContractBackend) (*Assetholdererc20, error) {
	contract, err := bindAssetholdererc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20{Assetholdererc20Caller: Assetholdererc20Caller{contract: contract}, Assetholdererc20Transactor: Assetholdererc20Transactor{contract: contract}, Assetholdererc20Filterer: Assetholdererc20Filterer{contract: contract}}, nil
}

// NewAssetholdererc20Caller creates a new read-only instance of Assetholdererc20, bound to a specific deployed contract.
func NewAssetholdererc20Caller(address common.Address, caller bind.ContractCaller) (*Assetholdererc20Caller, error) {
	contract, err := bindAssetholdererc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20Caller{contract: contract}, nil
}

// NewAssetholdererc20Transactor creates a new write-only instance of Assetholdererc20, bound to a specific deployed contract.
func NewAssetholdererc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Assetholdererc20Transactor, error) {
	contract, err := bindAssetholdererc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20Transactor{contract: contract}, nil
}

// NewAssetholdererc20Filterer creates a new log filterer instance of Assetholdererc20, bound to a specific deployed contract.
func NewAssetholdererc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Assetholdererc20Filterer, error) {
	contract, err := bindAssetholdererc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20Filterer{contract: contract}, nil
}

// bindAssetholdererc20 binds a generic wrapper to an already deployed contract.
func bindAssetholdererc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Assetholdererc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assetholdererc20 *Assetholdererc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assetholdererc20.Contract.Assetholdererc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assetholdererc20 *Assetholdererc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Assetholdererc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assetholdererc20 *Assetholdererc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Assetholdererc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assetholdererc20 *Assetholdererc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assetholdererc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assetholdererc20 *Assetholdererc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assetholdererc20 *Assetholdererc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.contract.Transact(opts, method, params...)
}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdererc20 *Assetholdererc20Caller) Adjudicator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Assetholdererc20.contract.Call(opts, &out, "adjudicator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdererc20 *Assetholdererc20Session) Adjudicator() (common.Address, error) {
	return _Assetholdererc20.Contract.Adjudicator(&_Assetholdererc20.CallOpts)
}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdererc20 *Assetholdererc20CallerSession) Adjudicator() (common.Address, error) {
	return _Assetholdererc20.Contract.Adjudicator(&_Assetholdererc20.CallOpts)
}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdererc20 *Assetholdererc20Caller) Holdings(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Assetholdererc20.contract.Call(opts, &out, "holdings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdererc20 *Assetholdererc20Session) Holdings(arg0 [32]byte) (*big.Int, error) {
	return _Assetholdererc20.Contract.Holdings(&_Assetholdererc20.CallOpts, arg0)
}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdererc20 *Assetholdererc20CallerSession) Holdings(arg0 [32]byte) (*big.Int, error) {
	return _Assetholdererc20.Contract.Holdings(&_Assetholdererc20.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdererc20 *Assetholdererc20Caller) Settled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Assetholdererc20.contract.Call(opts, &out, "settled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdererc20 *Assetholdererc20Session) Settled(arg0 [32]byte) (bool, error) {
	return _Assetholdererc20.Contract.Settled(&_Assetholdererc20.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdererc20 *Assetholdererc20CallerSession) Settled(arg0 [32]byte) (bool, error) {
	return _Assetholdererc20.Contract.Settled(&_Assetholdererc20.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Assetholdererc20 *Assetholdererc20Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Assetholdererc20.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Assetholdererc20 *Assetholdererc20Session) Token() (common.Address, error) {
	return _Assetholdererc20.Contract.Token(&_Assetholdererc20.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Assetholdererc20 *Assetholdererc20CallerSession) Token() (common.Address, error) {
	return _Assetholdererc20.Contract.Token(&_Assetholdererc20.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdererc20 *Assetholdererc20Transactor) Deposit(opts *bind.TransactOpts, fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.contract.Transact(opts, "deposit", fundingID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdererc20 *Assetholdererc20Session) Deposit(fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Deposit(&_Assetholdererc20.TransactOpts, fundingID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdererc20 *Assetholdererc20TransactorSession) Deposit(fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Deposit(&_Assetholdererc20.TransactOpts, fundingID, amount)
}

// SetOutcome is a paid mutator transaction binding the contract method 0xfc79a66d.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals) returns()
func (_Assetholdererc20 *Assetholdererc20Transactor) SetOutcome(opts *bind.TransactOpts, channelID [32]byte, parts []common.Address, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.contract.Transact(opts, "setOutcome", channelID, parts, newBals)
}

// SetOutcome is a paid mutator transaction binding the contract method 0xfc79a66d.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals) returns()
func (_Assetholdererc20 *Assetholdererc20Session) SetOutcome(channelID [32]byte, parts []common.Address, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.SetOutcome(&_Assetholdererc20.TransactOpts, channelID, parts, newBals)
}

// SetOutcome is a paid mutator transaction binding the contract method 0xfc79a66d.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals) returns()
func (_Assetholdererc20 *Assetholdererc20TransactorSession) SetOutcome(channelID [32]byte, parts []common.Address, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.SetOutcome(&_Assetholdererc20.TransactOpts, channelID, parts, newBals)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
func (_Assetholdererc20 *Assetholdererc20Transactor) Withdraw(opts *bind.TransactOpts, authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdererc20.contract.Transact(opts, "withdraw", authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
func (_Assetholdererc20 *Assetholdererc20Session) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Withdraw(&_Assetholdererc20.TransactOpts, authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
func (_Assetholdererc20 *Assetholdererc20TransactorSession) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Withdraw(&_Assetholdererc20.TransactOpts, authorization, signature)
}

// Assetholdererc20DepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Assetholdererc20 contract.
type Assetholdererc20DepositedIterator struct {
	Event *Assetholdererc20Deposited // Event containing the contract specifics and raw log

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
func (it *Assetholdererc20DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Assetholdererc20Deposited)
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
		it.Event = new(Assetholdererc20Deposited)
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
func (it *Assetholdererc20DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Assetholdererc20DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Assetholdererc20Deposited represents a Deposited event raised by the Assetholdererc20 contract.
type Assetholdererc20Deposited struct {
	FundingID [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assetholdererc20 *Assetholdererc20Filterer) FilterDeposited(opts *bind.FilterOpts, fundingID [][32]byte) (*Assetholdererc20DepositedIterator, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.FilterLogs(opts, "Deposited", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20DepositedIterator{contract: _Assetholdererc20.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assetholdererc20 *Assetholdererc20Filterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *Assetholdererc20Deposited, fundingID [][32]byte) (event.Subscription, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.WatchLogs(opts, "Deposited", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Assetholdererc20Deposited)
				if err := _Assetholdererc20.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_Assetholdererc20 *Assetholdererc20Filterer) ParseDeposited(log types.Log) (*Assetholdererc20Deposited, error) {
	event := new(Assetholdererc20Deposited)
	if err := _Assetholdererc20.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Assetholdererc20OutcomeSetIterator is returned from FilterOutcomeSet and is used to iterate over the raw logs and unpacked data for OutcomeSet events raised by the Assetholdererc20 contract.
type Assetholdererc20OutcomeSetIterator struct {
	Event *Assetholdererc20OutcomeSet // Event containing the contract specifics and raw log

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
func (it *Assetholdererc20OutcomeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Assetholdererc20OutcomeSet)
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
		it.Event = new(Assetholdererc20OutcomeSet)
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
func (it *Assetholdererc20OutcomeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Assetholdererc20OutcomeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Assetholdererc20OutcomeSet represents a OutcomeSet event raised by the Assetholdererc20 contract.
type Assetholdererc20OutcomeSet struct {
	ChannelID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutcomeSet is a free log retrieval operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assetholdererc20 *Assetholdererc20Filterer) FilterOutcomeSet(opts *bind.FilterOpts, channelID [][32]byte) (*Assetholdererc20OutcomeSetIterator, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.FilterLogs(opts, "OutcomeSet", channelIDRule)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20OutcomeSetIterator{contract: _Assetholdererc20.contract, event: "OutcomeSet", logs: logs, sub: sub}, nil
}

// WatchOutcomeSet is a free log subscription operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assetholdererc20 *Assetholdererc20Filterer) WatchOutcomeSet(opts *bind.WatchOpts, sink chan<- *Assetholdererc20OutcomeSet, channelID [][32]byte) (event.Subscription, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.WatchLogs(opts, "OutcomeSet", channelIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Assetholdererc20OutcomeSet)
				if err := _Assetholdererc20.contract.UnpackLog(event, "OutcomeSet", log); err != nil {
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
func (_Assetholdererc20 *Assetholdererc20Filterer) ParseOutcomeSet(log types.Log) (*Assetholdererc20OutcomeSet, error) {
	event := new(Assetholdererc20OutcomeSet)
	if err := _Assetholdererc20.contract.UnpackLog(event, "OutcomeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Assetholdererc20WithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Assetholdererc20 contract.
type Assetholdererc20WithdrawnIterator struct {
	Event *Assetholdererc20Withdrawn // Event containing the contract specifics and raw log

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
func (it *Assetholdererc20WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Assetholdererc20Withdrawn)
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
		it.Event = new(Assetholdererc20Withdrawn)
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
func (it *Assetholdererc20WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Assetholdererc20WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Assetholdererc20Withdrawn represents a Withdrawn event raised by the Assetholdererc20 contract.
type Assetholdererc20Withdrawn struct {
	FundingID [32]byte
	Amount    *big.Int
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assetholdererc20 *Assetholdererc20Filterer) FilterWithdrawn(opts *bind.FilterOpts, fundingID [][32]byte) (*Assetholdererc20WithdrawnIterator, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.FilterLogs(opts, "Withdrawn", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20WithdrawnIterator{contract: _Assetholdererc20.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assetholdererc20 *Assetholdererc20Filterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *Assetholdererc20Withdrawn, fundingID [][32]byte) (event.Subscription, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.WatchLogs(opts, "Withdrawn", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Assetholdererc20Withdrawn)
				if err := _Assetholdererc20.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_Assetholdererc20 *Assetholdererc20Filterer) ParseWithdrawn(log types.Log) (*Assetholdererc20Withdrawn, error) {
	event := new(Assetholdererc20Withdrawn)
	if err := _Assetholdererc20.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
