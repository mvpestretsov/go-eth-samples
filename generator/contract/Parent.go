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

// ParentABI is the input ABI used to generate the binding from.
const ParentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"t\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"generateChild\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ParentBin is the compiled bytecode used for deploying new contracts.
const ParentBin = `6060604052341561000f57600080fd5b5b61022a8061001f6000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806392d0d15314610049578063b3c7fc2414610072575b600080fd5b341561005457600080fd5b61005c6100ca565b6040518082815260200191505060405180910390f35b61008860048080359060200190919050506100d0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60005481565b600080826100dc610112565b80828152602001915050604051809103906000f08015156100fc57600080fd5b9050600283026000819055508091505b50919050565b60405160dd806101228339019056006060604052341561000f57600080fd5b6040516020806100dd833981016040528080519060200190919050505b806000819055505b505b6099806100446000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063136eb43c14603d575b600080fd5b3415604757600080fd5b604d6063565b6040518082815260200191505060405180910390f35b6000805490505b905600a165627a7a7230582068a2d1f6b14577f6f0f7e40787aa91bc35a3785161197a71381e2619fdb1012f0029a165627a7a72305820a11c2620f1a21293f0bfe31047a29d5c6a01edfabc3686fadd33ae069b8200360029`

// DeployParent deploys a new Ethereum contract, binding an instance of Parent to it.
func DeployParent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Parent, error) {
	parsed, err := abi.JSON(strings.NewReader(ParentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Parent{ParentCaller: ParentCaller{contract: contract}, ParentTransactor: ParentTransactor{contract: contract}}, nil
}

// Parent is an auto generated Go binding around an Ethereum contract.
type Parent struct {
	ParentCaller     // Read-only binding to the contract
	ParentTransactor // Write-only binding to the contract
}

// ParentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParentSession struct {
	Contract     *Parent           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParentCallerSession struct {
	Contract *ParentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ParentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParentTransactorSession struct {
	Contract     *ParentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParentRaw struct {
	Contract *Parent // Generic contract binding to access the raw methods on
}

// ParentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParentCallerRaw struct {
	Contract *ParentCaller // Generic read-only contract binding to access the raw methods on
}

// ParentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParentTransactorRaw struct {
	Contract *ParentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParent creates a new instance of Parent, bound to a specific deployed contract.
func NewParent(address common.Address, backend bind.ContractBackend) (*Parent, error) {
	contract, err := bindParent(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parent{ParentCaller: ParentCaller{contract: contract}, ParentTransactor: ParentTransactor{contract: contract}}, nil
}

// NewParentCaller creates a new read-only instance of Parent, bound to a specific deployed contract.
func NewParentCaller(address common.Address, caller bind.ContractCaller) (*ParentCaller, error) {
	contract, err := bindParent(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ParentCaller{contract: contract}, nil
}

// NewParentTransactor creates a new write-only instance of Parent, bound to a specific deployed contract.
func NewParentTransactor(address common.Address, transactor bind.ContractTransactor) (*ParentTransactor, error) {
	contract, err := bindParent(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ParentTransactor{contract: contract}, nil
}

// bindParent binds a generic wrapper to an already deployed contract.
func bindParent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parent *ParentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parent.Contract.ParentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parent *ParentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parent.Contract.ParentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parent *ParentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parent.Contract.ParentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parent *ParentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parent *ParentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parent *ParentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parent.Contract.contract.Transact(opts, method, params...)
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(uint256)
func (_Parent *ParentCaller) T(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parent.contract.Call(opts, out, "t")
	return *ret0, err
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(uint256)
func (_Parent *ParentSession) T() (*big.Int, error) {
	return _Parent.Contract.T(&_Parent.CallOpts)
}

// T is a free data retrieval call binding the contract method 0x92d0d153.
//
// Solidity: function t() constant returns(uint256)
func (_Parent *ParentCallerSession) T() (*big.Int, error) {
	return _Parent.Contract.T(&_Parent.CallOpts)
}

// GenerateChild is a paid mutator transaction binding the contract method 0xb3c7fc24.
//
// Solidity: function generateChild(i uint256) returns(address)
func (_Parent *ParentTransactor) GenerateChild(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Parent.contract.Transact(opts, "generateChild", i)
}

// GenerateChild is a paid mutator transaction binding the contract method 0xb3c7fc24.
//
// Solidity: function generateChild(i uint256) returns(address)
func (_Parent *ParentSession) GenerateChild(i *big.Int) (*types.Transaction, error) {
	return _Parent.Contract.GenerateChild(&_Parent.TransactOpts, i)
}

// GenerateChild is a paid mutator transaction binding the contract method 0xb3c7fc24.
//
// Solidity: function generateChild(i uint256) returns(address)
func (_Parent *ParentTransactorSession) GenerateChild(i *big.Int) (*types.Transaction, error) {
	return _Parent.Contract.GenerateChild(&_Parent.TransactOpts, i)
}
