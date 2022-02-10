// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SimplestABI is the input ABI used to generate the binding from.
const SimplestABI = "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setVersion\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"newVersion\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"version\",\"inputs\":[]}]"

// Simplest is an auto generated Go binding around an Ethereum contract.
type Simplest struct {
	SimplestCaller     // Read-only binding to the contract
	SimplestTransactor // Write-only binding to the contract
	SimplestFilterer   // Log filterer for contract events
}

// SimplestCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplestSession struct {
	Contract     *Simplest         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimplestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplestCallerSession struct {
	Contract *SimplestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SimplestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplestTransactorSession struct {
	Contract     *SimplestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SimplestRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplestRaw struct {
	Contract *Simplest // Generic contract binding to access the raw methods on
}

// SimplestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplestCallerRaw struct {
	Contract *SimplestCaller // Generic read-only contract binding to access the raw methods on
}

// SimplestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplestTransactorRaw struct {
	Contract *SimplestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplest creates a new instance of Simplest, bound to a specific deployed contract.
func NewSimplest(address common.Address, backend bind.ContractBackend) (*Simplest, error) {
	contract, err := bindSimplest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplest{SimplestCaller: SimplestCaller{contract: contract}, SimplestTransactor: SimplestTransactor{contract: contract}, SimplestFilterer: SimplestFilterer{contract: contract}}, nil
}

// NewSimplestCaller creates a new read-only instance of Simplest, bound to a specific deployed contract.
func NewSimplestCaller(address common.Address, caller bind.ContractCaller) (*SimplestCaller, error) {
	contract, err := bindSimplest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplestCaller{contract: contract}, nil
}

// NewSimplestTransactor creates a new write-only instance of Simplest, bound to a specific deployed contract.
func NewSimplestTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplestTransactor, error) {
	contract, err := bindSimplest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplestTransactor{contract: contract}, nil
}

// NewSimplestFilterer creates a new log filterer instance of Simplest, bound to a specific deployed contract.
func NewSimplestFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplestFilterer, error) {
	contract, err := bindSimplest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplestFilterer{contract: contract}, nil
}

// bindSimplest binds a generic wrapper to an already deployed contract.
func bindSimplest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimplestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplest *SimplestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplest.Contract.SimplestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplest *SimplestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplest.Contract.SimplestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplest *SimplestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplest.Contract.SimplestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplest *SimplestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplest *SimplestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplest *SimplestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplest.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Simplest *SimplestCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simplest.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Simplest *SimplestSession) Version() (*big.Int, error) {
	return _Simplest.Contract.Version(&_Simplest.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Simplest *SimplestCallerSession) Version() (*big.Int, error) {
	return _Simplest.Contract.Version(&_Simplest.CallOpts)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 newVersion) returns()
func (_Simplest *SimplestTransactor) SetVersion(opts *bind.TransactOpts, newVersion *big.Int) (*types.Transaction, error) {
	return _Simplest.contract.Transact(opts, "setVersion", newVersion)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 newVersion) returns()
func (_Simplest *SimplestSession) SetVersion(newVersion *big.Int) (*types.Transaction, error) {
	return _Simplest.Contract.SetVersion(&_Simplest.TransactOpts, newVersion)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 newVersion) returns()
func (_Simplest *SimplestTransactorSession) SetVersion(newVersion *big.Int) (*types.Transaction, error) {
	return _Simplest.Contract.SetVersion(&_Simplest.TransactOpts, newVersion)
}
