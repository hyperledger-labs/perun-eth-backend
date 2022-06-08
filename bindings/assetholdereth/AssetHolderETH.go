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
)

// AssetHolderWithdrawalAuth is an auto generated low-level Go binding around an user-defined struct.
type AssetHolderWithdrawalAuth struct {
	ChannelID   [32]byte
	Participant common.Address
	Receiver    common.Address
	Amount      *big.Int
}

// AssetholderethMetaData contains all meta data concerning the Assetholdereth contract.
var AssetholderethMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adjudicator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"}],\"name\":\"OutcomeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adjudicator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"holdings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBals\",\"type\":\"uint256[]\"}],\"name\":\"setOutcome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"settled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAssetHolder.WithdrawalAuth\",\"name\":\"authorization\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610f6d380380610f6d83398101604081905261002f91610054565b600280546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610eda806100936000396000f3fe6080604052600436106100555760003560e01c80631de26e161461005a5780634ed4283c1461006f57806353c2ed8e1461008f578063ae9ee18c146100cc578063d945af1d14610107578063fc79a66d14610147575b600080fd5b61006d610068366004610c01565b610167565b005b34801561007b57600080fd5b5061006d61008a366004610c23565b6101d7565b34801561009b57600080fd5b506002546100af906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100d857600080fd5b506100f96100e7366004610cae565b60006020819052908152604090205481565b6040519081526020016100c3565b34801561011357600080fd5b50610137610122366004610cae565b60016020526000908152604090205460ff1681565b60405190151581526020016100c3565b34801561015357600080fd5b5061006d610162366004610d0c565b6103ef565b61017182826106fe565b60008281526020819052604090205461018a9082610751565b600083815260208190526040902055817fcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9826040516101cb91815260200190565b60405180910390a25050565b823560009081526001602052604090205460ff166102325760405162461bcd60e51b815260206004820152601360248201527218da185b9b995b081b9bdd081cd95d1d1b1959606a1b60448201526064015b60405180910390fd5b61029a836040516020016102469190610d9b565b60408051601f198184030181526020601f860181900481028401810190925284835291908590859081908401838280828437600092019190915250610295925050506040870160208801610de7565b610764565b6102e65760405162461bcd60e51b815260206004820152601d60248201527f7369676e617475726520766572696669636174696f6e206661696c65640000006044820152606401610229565b600061030284356102fd6040870160208801610de7565b6107ed565b6000818152602081905260409020549091506060850135111561035c5760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b6044820152606401610229565b600081815260208190526040902054610379906060860135610832565b60008281526020819052604090205561039384848461083e565b807fd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81606086018035906103c99060408901610de7565b604080519283526001600160a01b0390911660208301520160405180910390a250505050565b6002546001600160a01b031633146104575760405162461bcd60e51b815260206004820152602560248201527f63616e206f6e6c792062652063616c6c6564206279207468652061646a75646960448201526431b0ba37b960d91b6064820152608401610229565b8281146104b85760405162461bcd60e51b815260206004820152602960248201527f7061727469636970616e7473206c656e6774682073686f756c6420657175616c6044820152682062616c616e63657360b81b6064820152608401610229565b60008581526001602052604090205460ff16156105255760405162461bcd60e51b815260206004820152602560248201527f747279696e6720746f2073657420616c726561647920736574746c6564206368604482015264185b9b995b60da1b6064820152608401610229565b600085815260208190526040812080549082905590808567ffffffffffffffff81111561055457610554610e04565b60405190808252806020026020018201604052801561057d578160200160208202803683370190505b50905060005b868110156106415760006105b88a8a8a858181106105a3576105a3610e1a565b90506020020160208101906102fd9190610de7565b9050808383815181106105cd576105cd610e1a565b6020026020010181815250506105fe600080838152602001908152602001600020548661075190919063ffffffff16565b945061062b87878481811061061557610615610e1a565b905060200201358561075190919063ffffffff16565b935050808061063990610e46565b915050610583565b508183106106b15760005b868110156106af5785858281811061066657610666610e1a565b9050602002013560008084848151811061068257610682610e1a565b602002602001015181526020019081526020016000208190555080806106a790610e46565b91505061064c565b505b6000888152600160208190526040808320805460ff19169092179091555189917fef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b891a25050505050505050565b80341461074d5760405162461bcd60e51b815260206004820152601f60248201527f77726f6e6720616d6f756e74206f662045544820666f72206465706f736974006044820152606401610229565b5050565b600061075d8284610e5f565b9392505050565b6000806107c585805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905060006107d38286610890565b6001600160a01b0390811690851614925050509392505050565b600082826040516020016108149291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120905092915050565b600061075d8284610e77565b61084e6060840160408501610de7565b6001600160a01b03166108fc84606001359081150290604051600060405180830381858888f1935050505015801561088a573d6000803e3d6000fd5b50505050565b600080600061089f85856108b4565b915091506108ac81610922565b509392505050565b60008082516041036108ea5760208301516040840151606085015160001a6108de87828585610adb565b9450945050505061091b565b82516040036109135760208301516040840151610908868383610bc8565b93509350505061091b565b506000905060025b9250929050565b600081600481111561093657610936610e8e565b0361093e5750565b600181600481111561095257610952610e8e565b0361099f5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610229565b60028160048111156109b3576109b3610e8e565b03610a005760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610229565b6003816004811115610a1457610a14610e8e565b03610a6c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610229565b6004816004811115610a8057610a80610e8e565b03610ad85760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610229565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610b125750600090506003610bbf565b8460ff16601b14158015610b2a57508460ff16601c14155b15610b3b5750600090506004610bbf565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610b8f573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610bb857600060019250925050610bbf565b9150600090505b94509492505050565b6000806001600160ff1b03831681610be560ff86901c601b610e5f565b9050610bf387828885610adb565b935093505050935093915050565b60008060408385031215610c1457600080fd5b50508035926020909101359150565b600080600083850360a0811215610c3957600080fd5b6080811215610c4757600080fd5b50839250608084013567ffffffffffffffff80821115610c6657600080fd5b818601915086601f830112610c7a57600080fd5b813581811115610c8957600080fd5b876020828501011115610c9b57600080fd5b6020830194508093505050509250925092565b600060208284031215610cc057600080fd5b5035919050565b60008083601f840112610cd957600080fd5b50813567ffffffffffffffff811115610cf157600080fd5b6020830191508360208260051b850101111561091b57600080fd5b600080600080600060608688031215610d2457600080fd5b85359450602086013567ffffffffffffffff80821115610d4357600080fd5b610d4f89838a01610cc7565b90965094506040880135915080821115610d6857600080fd5b50610d7588828901610cc7565b969995985093965092949392505050565b6001600160a01b0381168114610ad857600080fd5b81358152608081016020830135610db181610d86565b6001600160a01b039081166020840152604084013590610dd082610d86565b166040830152606092830135929091019190915290565b600060208284031215610df957600080fd5b813561075d81610d86565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610e5857610e58610e30565b5060010190565b60008219821115610e7257610e72610e30565b500190565b600082821015610e8957610e89610e30565b500390565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220df38ffbef9b77f9300a169996ada023fc14ce7920d62f0fd409aafac2adfbe6c64736f6c634300080e0033",
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
	parsed, err := abi.JSON(strings.NewReader(AssetholderethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// SetOutcome is a paid mutator transaction binding the contract method 0xfc79a66d.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals) returns()
func (_Assetholdereth *AssetholderethTransactor) SetOutcome(opts *bind.TransactOpts, channelID [32]byte, parts []common.Address, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdereth.contract.Transact(opts, "setOutcome", channelID, parts, newBals)
}

// SetOutcome is a paid mutator transaction binding the contract method 0xfc79a66d.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals) returns()
func (_Assetholdereth *AssetholderethSession) SetOutcome(channelID [32]byte, parts []common.Address, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdereth.Contract.SetOutcome(&_Assetholdereth.TransactOpts, channelID, parts, newBals)
}

// SetOutcome is a paid mutator transaction binding the contract method 0xfc79a66d.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals) returns()
func (_Assetholdereth *AssetholderethTransactorSession) SetOutcome(channelID [32]byte, parts []common.Address, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdereth.Contract.SetOutcome(&_Assetholdereth.TransactOpts, channelID, parts, newBals)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
func (_Assetholdereth *AssetholderethTransactor) Withdraw(opts *bind.TransactOpts, authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdereth.contract.Transact(opts, "withdraw", authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
func (_Assetholdereth *AssetholderethSession) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdereth.Contract.Withdraw(&_Assetholdereth.TransactOpts, authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
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
