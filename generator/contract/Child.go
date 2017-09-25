// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generator

import (
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ChildABI is the input ABI used to generate the binding from.
const ChildABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getStrorage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ChildBin is the compiled bytecode used for deploying new contracts.
const ChildBin = `6060604052341561000f57600080fd5b6040516020806100dd833981016040528080519060200190919050505b806000819055505b505b6099806100446000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063136eb43c14603d575b600080fd5b3415604757600080fd5b604d6063565b6040518082815260200191505060405180910390f35b6000805490505b905600a165627a7a7230582068a2d1f6b14577f6f0f7e40787aa91bc35a3785161197a71381e2619fdb1012f0029`

// DeployChild deploys a new Ethereum contract, binding an instance of Child to it.
func DeployChild(auth *bind.TransactOpts, backend bind.ContractBackend, i *big.Int) (common.Address, *types.Transaction, *Child, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChildBin), backend, i)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Child{ChildCaller: ChildCaller{contract: contract}, ChildTransactor: ChildTransactor{contract: contract}}, nil
}

// Child is an auto generated Go binding around an Ethereum contract.
type Child struct {
	ChildCaller     // Read-only binding to the contract
	ChildTransactor // Write-only binding to the contract
}

// ChildCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildSession struct {
	Contract     *Child            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChildCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildCallerSession struct {
	Contract *ChildCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ChildTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildTransactorSession struct {
	Contract     *ChildTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChildRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildRaw struct {
	Contract *Child // Generic contract binding to access the raw methods on
}

// ChildCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildCallerRaw struct {
	Contract *ChildCaller // Generic read-only contract binding to access the raw methods on
}

// ChildTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildTransactorRaw struct {
	Contract *ChildTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChild creates a new instance of Child, bound to a specific deployed contract.
func NewChild(address common.Address, backend bind.ContractBackend) (*Child, error) {
	contract, err := bindChild(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Child{ChildCaller: ChildCaller{contract: contract}, ChildTransactor: ChildTransactor{contract: contract}}, nil
}

// NewChildCaller creates a new read-only instance of Child, bound to a specific deployed contract.
func NewChildCaller(address common.Address, caller bind.ContractCaller) (*ChildCaller, error) {
	contract, err := bindChild(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ChildCaller{contract: contract}, nil
}

// NewChildTransactor creates a new write-only instance of Child, bound to a specific deployed contract.
func NewChildTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildTransactor, error) {
	contract, err := bindChild(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ChildTransactor{contract: contract}, nil
}

// bindChild binds a generic wrapper to an already deployed contract.
func bindChild(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Child *ChildRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Child.Contract.ChildCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Child *ChildRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Child.Contract.ChildTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Child *ChildRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Child.Contract.ChildTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Child *ChildCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Child.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Child *ChildTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Child.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Child *ChildTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Child.Contract.contract.Transact(opts, method, params...)
}

// GetStrorage is a free data retrieval call binding the contract method 0x136eb43c.
//
// Solidity: function getStrorage() constant returns(uint256)
func (_Child *ChildCaller) GetStrorage(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Child.contract.Call(opts, out, "getStrorage")
	return *ret0, err
}

// GetStrorage is a free data retrieval call binding the contract method 0x136eb43c.
//
// Solidity: function getStrorage() constant returns(uint256)
func (_Child *ChildSession) GetStrorage() (*big.Int, error) {
	return _Child.Contract.GetStrorage(&_Child.CallOpts)
}

// GetStrorage is a free data retrieval call binding the contract method 0x136eb43c.
//
// Solidity: function getStrorage() constant returns(uint256)
func (_Child *ChildCallerSession) GetStrorage() (*big.Int, error) {
	return _Child.Contract.GetStrorage(&_Child.CallOpts)
}
