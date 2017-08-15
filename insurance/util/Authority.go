// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package insurance

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

// AuthorityABI is the input ABI used to generate the binding from.
const AuthorityABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBuyInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"maxTime\",\"type\":\"uint256\"}],\"name\":\"setMaxTimToConfirm\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"blockSeller\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isPubKeyUsed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"isBuyerBlocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"maxTime\",\"type\":\"uint256\"}],\"name\":\"setMaxTimeToUpload\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pubKey\",\"type\":\"string\"},{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"markPubKeyAsUsed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTimeToConfirm\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTimeToUpload\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"unblockBuyer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyInfoAd\",\"type\":\"address\"}],\"name\":\"processPendingContract\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sha256Sign\",\"type\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isPubKeySignUsed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sha256Sign\",\"type\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"markPubKeySignAsUsed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"company\",\"type\":\"address\"},{\"name\":\"buyerPubKey\",\"type\":\"string\"},{\"name\":\"dataH\",\"type\":\"string\"},{\"name\":\"idH\",\"type\":\"string\"}],\"name\":\"createBuyContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyInfoAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"unblockSeller\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserve\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"blockBuyer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"isSellerBlocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"ownerAd\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"returnValue\",\"type\":\"address\"}],\"name\":\"ContractEvent\",\"type\":\"event\"}]"

// Authority is an auto generated Go binding around an Ethereum contract.
type Authority struct {
	AuthorityCaller     // Read-only binding to the contract
	AuthorityTransactor // Write-only binding to the contract
}

// AuthorityCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthoritySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthoritySession struct {
	Contract     *Authority        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorityCallerSession struct {
	Contract *AuthorityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AuthorityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorityTransactorSession struct {
	Contract     *AuthorityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AuthorityRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorityRaw struct {
	Contract *Authority // Generic contract binding to access the raw methods on
}

// AuthorityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorityCallerRaw struct {
	Contract *AuthorityCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorityTransactorRaw struct {
	Contract *AuthorityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthority creates a new instance of Authority, bound to a specific deployed contract.
func NewAuthority(address common.Address, backend bind.ContractBackend) (*Authority, error) {
	contract, err := bindAuthority(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Authority{AuthorityCaller: AuthorityCaller{contract: contract}, AuthorityTransactor: AuthorityTransactor{contract: contract}}, nil
}

// NewAuthorityCaller creates a new read-only instance of Authority, bound to a specific deployed contract.
func NewAuthorityCaller(address common.Address, caller bind.ContractCaller) (*AuthorityCaller, error) {
	contract, err := bindAuthority(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityCaller{contract: contract}, nil
}

// NewAuthorityTransactor creates a new write-only instance of Authority, bound to a specific deployed contract.
func NewAuthorityTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorityTransactor, error) {
	contract, err := bindAuthority(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AuthorityTransactor{contract: contract}, nil
}

// bindAuthority binds a generic wrapper to an already deployed contract.
func bindAuthority(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authority *AuthorityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authority.Contract.AuthorityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authority *AuthorityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authority.Contract.AuthorityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authority *AuthorityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authority.Contract.AuthorityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authority *AuthorityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authority.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authority *AuthorityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authority.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authority *AuthorityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authority.Contract.contract.Transact(opts, method, params...)
}

// BuyInfoAddr is a free data retrieval call binding the contract method 0xb5f94a1c.
//
// Solidity: function buyInfoAddr() constant returns(address)
func (_Authority *AuthorityCaller) BuyInfoAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "buyInfoAddr")
	return *ret0, err
}

// BuyInfoAddr is a free data retrieval call binding the contract method 0xb5f94a1c.
//
// Solidity: function buyInfoAddr() constant returns(address)
func (_Authority *AuthoritySession) BuyInfoAddr() (common.Address, error) {
	return _Authority.Contract.BuyInfoAddr(&_Authority.CallOpts)
}

// BuyInfoAddr is a free data retrieval call binding the contract method 0xb5f94a1c.
//
// Solidity: function buyInfoAddr() constant returns(address)
func (_Authority *AuthorityCallerSession) BuyInfoAddr() (common.Address, error) {
	return _Authority.Contract.BuyInfoAddr(&_Authority.CallOpts)
}

// ContractPrice is a free data retrieval call binding the contract method 0x97e9a0bf.
//
// Solidity: function contractPrice() constant returns(uint256)
func (_Authority *AuthorityCaller) ContractPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "contractPrice")
	return *ret0, err
}

// ContractPrice is a free data retrieval call binding the contract method 0x97e9a0bf.
//
// Solidity: function contractPrice() constant returns(uint256)
func (_Authority *AuthoritySession) ContractPrice() (*big.Int, error) {
	return _Authority.Contract.ContractPrice(&_Authority.CallOpts)
}

// ContractPrice is a free data retrieval call binding the contract method 0x97e9a0bf.
//
// Solidity: function contractPrice() constant returns(uint256)
func (_Authority *AuthorityCallerSession) ContractPrice() (*big.Int, error) {
	return _Authority.Contract.ContractPrice(&_Authority.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Authority *AuthorityCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Authority *AuthoritySession) Fee() (*big.Int, error) {
	return _Authority.Contract.Fee(&_Authority.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Authority *AuthorityCallerSession) Fee() (*big.Int, error) {
	return _Authority.Contract.Fee(&_Authority.CallOpts)
}

// GetBuyInfo is a free data retrieval call binding the contract method 0x1aa7c1ce.
//
// Solidity: function getBuyInfo() constant returns(address)
func (_Authority *AuthorityCaller) GetBuyInfo(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "getBuyInfo")
	return *ret0, err
}

// GetBuyInfo is a free data retrieval call binding the contract method 0x1aa7c1ce.
//
// Solidity: function getBuyInfo() constant returns(address)
func (_Authority *AuthoritySession) GetBuyInfo() (common.Address, error) {
	return _Authority.Contract.GetBuyInfo(&_Authority.CallOpts)
}

// GetBuyInfo is a free data retrieval call binding the contract method 0x1aa7c1ce.
//
// Solidity: function getBuyInfo() constant returns(address)
func (_Authority *AuthorityCallerSession) GetBuyInfo() (common.Address, error) {
	return _Authority.Contract.GetBuyInfo(&_Authority.CallOpts)
}

// IsBuyerBlocked is a free data retrieval call binding the contract method 0x4bb9951e.
//
// Solidity: function isBuyerBlocked(buyer address) constant returns(bool)
func (_Authority *AuthorityCaller) IsBuyerBlocked(opts *bind.CallOpts, buyer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "isBuyerBlocked", buyer)
	return *ret0, err
}

// IsBuyerBlocked is a free data retrieval call binding the contract method 0x4bb9951e.
//
// Solidity: function isBuyerBlocked(buyer address) constant returns(bool)
func (_Authority *AuthoritySession) IsBuyerBlocked(buyer common.Address) (bool, error) {
	return _Authority.Contract.IsBuyerBlocked(&_Authority.CallOpts, buyer)
}

// IsBuyerBlocked is a free data retrieval call binding the contract method 0x4bb9951e.
//
// Solidity: function isBuyerBlocked(buyer address) constant returns(bool)
func (_Authority *AuthorityCallerSession) IsBuyerBlocked(buyer common.Address) (bool, error) {
	return _Authority.Contract.IsBuyerBlocked(&_Authority.CallOpts, buyer)
}

// IsPubKeySignUsed is a free data retrieval call binding the contract method 0x7dd5eefd.
//
// Solidity: function isPubKeySignUsed(sha256Sign bytes32, sender address) constant returns(bool)
func (_Authority *AuthorityCaller) IsPubKeySignUsed(opts *bind.CallOpts, sha256Sign [32]byte, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "isPubKeySignUsed", sha256Sign, sender)
	return *ret0, err
}

// IsPubKeySignUsed is a free data retrieval call binding the contract method 0x7dd5eefd.
//
// Solidity: function isPubKeySignUsed(sha256Sign bytes32, sender address) constant returns(bool)
func (_Authority *AuthoritySession) IsPubKeySignUsed(sha256Sign [32]byte, sender common.Address) (bool, error) {
	return _Authority.Contract.IsPubKeySignUsed(&_Authority.CallOpts, sha256Sign, sender)
}

// IsPubKeySignUsed is a free data retrieval call binding the contract method 0x7dd5eefd.
//
// Solidity: function isPubKeySignUsed(sha256Sign bytes32, sender address) constant returns(bool)
func (_Authority *AuthorityCallerSession) IsPubKeySignUsed(sha256Sign [32]byte, sender common.Address) (bool, error) {
	return _Authority.Contract.IsPubKeySignUsed(&_Authority.CallOpts, sha256Sign, sender)
}

// IsPubKeyUsed is a free data retrieval call binding the contract method 0x36f52388.
//
// Solidity: function isPubKeyUsed(pubKey string, sender address) constant returns(bool)
func (_Authority *AuthorityCaller) IsPubKeyUsed(opts *bind.CallOpts, pubKey string, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "isPubKeyUsed", pubKey, sender)
	return *ret0, err
}

// IsPubKeyUsed is a free data retrieval call binding the contract method 0x36f52388.
//
// Solidity: function isPubKeyUsed(pubKey string, sender address) constant returns(bool)
func (_Authority *AuthoritySession) IsPubKeyUsed(pubKey string, sender common.Address) (bool, error) {
	return _Authority.Contract.IsPubKeyUsed(&_Authority.CallOpts, pubKey, sender)
}

// IsPubKeyUsed is a free data retrieval call binding the contract method 0x36f52388.
//
// Solidity: function isPubKeyUsed(pubKey string, sender address) constant returns(bool)
func (_Authority *AuthorityCallerSession) IsPubKeyUsed(pubKey string, sender common.Address) (bool, error) {
	return _Authority.Contract.IsPubKeyUsed(&_Authority.CallOpts, pubKey, sender)
}

// IsSellerBlocked is a free data retrieval call binding the contract method 0xfd660d68.
//
// Solidity: function isSellerBlocked(seller address) constant returns(bool)
func (_Authority *AuthorityCaller) IsSellerBlocked(opts *bind.CallOpts, seller common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "isSellerBlocked", seller)
	return *ret0, err
}

// IsSellerBlocked is a free data retrieval call binding the contract method 0xfd660d68.
//
// Solidity: function isSellerBlocked(seller address) constant returns(bool)
func (_Authority *AuthoritySession) IsSellerBlocked(seller common.Address) (bool, error) {
	return _Authority.Contract.IsSellerBlocked(&_Authority.CallOpts, seller)
}

// IsSellerBlocked is a free data retrieval call binding the contract method 0xfd660d68.
//
// Solidity: function isSellerBlocked(seller address) constant returns(bool)
func (_Authority *AuthorityCallerSession) IsSellerBlocked(seller common.Address) (bool, error) {
	return _Authority.Contract.IsSellerBlocked(&_Authority.CallOpts, seller)
}

// MaxTimeToConfirm is a free data retrieval call binding the contract method 0x666dc604.
//
// Solidity: function maxTimeToConfirm() constant returns(uint256)
func (_Authority *AuthorityCaller) MaxTimeToConfirm(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "maxTimeToConfirm")
	return *ret0, err
}

// MaxTimeToConfirm is a free data retrieval call binding the contract method 0x666dc604.
//
// Solidity: function maxTimeToConfirm() constant returns(uint256)
func (_Authority *AuthoritySession) MaxTimeToConfirm() (*big.Int, error) {
	return _Authority.Contract.MaxTimeToConfirm(&_Authority.CallOpts)
}

// MaxTimeToConfirm is a free data retrieval call binding the contract method 0x666dc604.
//
// Solidity: function maxTimeToConfirm() constant returns(uint256)
func (_Authority *AuthorityCallerSession) MaxTimeToConfirm() (*big.Int, error) {
	return _Authority.Contract.MaxTimeToConfirm(&_Authority.CallOpts)
}

// MaxTimeToUpload is a free data retrieval call binding the contract method 0x6a6bb1a4.
//
// Solidity: function maxTimeToUpload() constant returns(uint256)
func (_Authority *AuthorityCaller) MaxTimeToUpload(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "maxTimeToUpload")
	return *ret0, err
}

// MaxTimeToUpload is a free data retrieval call binding the contract method 0x6a6bb1a4.
//
// Solidity: function maxTimeToUpload() constant returns(uint256)
func (_Authority *AuthoritySession) MaxTimeToUpload() (*big.Int, error) {
	return _Authority.Contract.MaxTimeToUpload(&_Authority.CallOpts)
}

// MaxTimeToUpload is a free data retrieval call binding the contract method 0x6a6bb1a4.
//
// Solidity: function maxTimeToUpload() constant returns(uint256)
func (_Authority *AuthorityCallerSession) MaxTimeToUpload() (*big.Int, error) {
	return _Authority.Contract.MaxTimeToUpload(&_Authority.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Authority *AuthorityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Authority *AuthoritySession) Owner() (common.Address, error) {
	return _Authority.Contract.Owner(&_Authority.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Authority *AuthorityCallerSession) Owner() (common.Address, error) {
	return _Authority.Contract.Owner(&_Authority.CallOpts)
}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() constant returns(uint256)
func (_Authority *AuthorityCaller) Reserve(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "reserve")
	return *ret0, err
}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() constant returns(uint256)
func (_Authority *AuthoritySession) Reserve() (*big.Int, error) {
	return _Authority.Contract.Reserve(&_Authority.CallOpts)
}

// Reserve is a free data retrieval call binding the contract method 0xcd3293de.
//
// Solidity: function reserve() constant returns(uint256)
func (_Authority *AuthorityCallerSession) Reserve() (*big.Int, error) {
	return _Authority.Contract.Reserve(&_Authority.CallOpts)
}

// BlockBuyer is a paid mutator transaction binding the contract method 0xfc6c135c.
//
// Solidity: function blockBuyer(buyer address) returns()
func (_Authority *AuthorityTransactor) BlockBuyer(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "blockBuyer", buyer)
}

// BlockBuyer is a paid mutator transaction binding the contract method 0xfc6c135c.
//
// Solidity: function blockBuyer(buyer address) returns()
func (_Authority *AuthoritySession) BlockBuyer(buyer common.Address) (*types.Transaction, error) {
	return _Authority.Contract.BlockBuyer(&_Authority.TransactOpts, buyer)
}

// BlockBuyer is a paid mutator transaction binding the contract method 0xfc6c135c.
//
// Solidity: function blockBuyer(buyer address) returns()
func (_Authority *AuthorityTransactorSession) BlockBuyer(buyer common.Address) (*types.Transaction, error) {
	return _Authority.Contract.BlockBuyer(&_Authority.TransactOpts, buyer)
}

// BlockSeller is a paid mutator transaction binding the contract method 0x29e48379.
//
// Solidity: function blockSeller(seller address) returns()
func (_Authority *AuthorityTransactor) BlockSeller(opts *bind.TransactOpts, seller common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "blockSeller", seller)
}

// BlockSeller is a paid mutator transaction binding the contract method 0x29e48379.
//
// Solidity: function blockSeller(seller address) returns()
func (_Authority *AuthoritySession) BlockSeller(seller common.Address) (*types.Transaction, error) {
	return _Authority.Contract.BlockSeller(&_Authority.TransactOpts, seller)
}

// BlockSeller is a paid mutator transaction binding the contract method 0x29e48379.
//
// Solidity: function blockSeller(seller address) returns()
func (_Authority *AuthorityTransactorSession) BlockSeller(seller common.Address) (*types.Transaction, error) {
	return _Authority.Contract.BlockSeller(&_Authority.TransactOpts, seller)
}

// CreateBuyContract is a paid mutator transaction binding the contract method 0x8e67810c.
//
// Solidity: function createBuyContract(company address, buyerPubKey string, dataH string, idH string) returns(address)
func (_Authority *AuthorityTransactor) CreateBuyContract(opts *bind.TransactOpts, company common.Address, buyerPubKey string, dataH string, idH string) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "createBuyContract", company, buyerPubKey, dataH, idH)
}

// CreateBuyContract is a paid mutator transaction binding the contract method 0x8e67810c.
//
// Solidity: function createBuyContract(company address, buyerPubKey string, dataH string, idH string) returns(address)
func (_Authority *AuthoritySession) CreateBuyContract(company common.Address, buyerPubKey string, dataH string, idH string) (*types.Transaction, error) {
	return _Authority.Contract.CreateBuyContract(&_Authority.TransactOpts, company, buyerPubKey, dataH, idH)
}

// CreateBuyContract is a paid mutator transaction binding the contract method 0x8e67810c.
//
// Solidity: function createBuyContract(company address, buyerPubKey string, dataH string, idH string) returns(address)
func (_Authority *AuthorityTransactorSession) CreateBuyContract(company common.Address, buyerPubKey string, dataH string, idH string) (*types.Transaction, error) {
	return _Authority.Contract.CreateBuyContract(&_Authority.TransactOpts, company, buyerPubKey, dataH, idH)
}

// MarkPubKeyAsUsed is a paid mutator transaction binding the contract method 0x5e71fe25.
//
// Solidity: function markPubKeyAsUsed(pubKey string, sender address) returns()
func (_Authority *AuthorityTransactor) MarkPubKeyAsUsed(opts *bind.TransactOpts, pubKey string, sender common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "markPubKeyAsUsed", pubKey, sender)
}

// MarkPubKeyAsUsed is a paid mutator transaction binding the contract method 0x5e71fe25.
//
// Solidity: function markPubKeyAsUsed(pubKey string, sender address) returns()
func (_Authority *AuthoritySession) MarkPubKeyAsUsed(pubKey string, sender common.Address) (*types.Transaction, error) {
	return _Authority.Contract.MarkPubKeyAsUsed(&_Authority.TransactOpts, pubKey, sender)
}

// MarkPubKeyAsUsed is a paid mutator transaction binding the contract method 0x5e71fe25.
//
// Solidity: function markPubKeyAsUsed(pubKey string, sender address) returns()
func (_Authority *AuthorityTransactorSession) MarkPubKeyAsUsed(pubKey string, sender common.Address) (*types.Transaction, error) {
	return _Authority.Contract.MarkPubKeyAsUsed(&_Authority.TransactOpts, pubKey, sender)
}

// MarkPubKeySignAsUsed is a paid mutator transaction binding the contract method 0x7e7c5888.
//
// Solidity: function markPubKeySignAsUsed(sha256Sign bytes32, sender address) returns()
func (_Authority *AuthorityTransactor) MarkPubKeySignAsUsed(opts *bind.TransactOpts, sha256Sign [32]byte, sender common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "markPubKeySignAsUsed", sha256Sign, sender)
}

// MarkPubKeySignAsUsed is a paid mutator transaction binding the contract method 0x7e7c5888.
//
// Solidity: function markPubKeySignAsUsed(sha256Sign bytes32, sender address) returns()
func (_Authority *AuthoritySession) MarkPubKeySignAsUsed(sha256Sign [32]byte, sender common.Address) (*types.Transaction, error) {
	return _Authority.Contract.MarkPubKeySignAsUsed(&_Authority.TransactOpts, sha256Sign, sender)
}

// MarkPubKeySignAsUsed is a paid mutator transaction binding the contract method 0x7e7c5888.
//
// Solidity: function markPubKeySignAsUsed(sha256Sign bytes32, sender address) returns()
func (_Authority *AuthorityTransactorSession) MarkPubKeySignAsUsed(sha256Sign [32]byte, sender common.Address) (*types.Transaction, error) {
	return _Authority.Contract.MarkPubKeySignAsUsed(&_Authority.TransactOpts, sha256Sign, sender)
}

// ProcessPendingContract is a paid mutator transaction binding the contract method 0x7866af46.
//
// Solidity: function processPendingContract(buyInfoAd address) returns()
func (_Authority *AuthorityTransactor) ProcessPendingContract(opts *bind.TransactOpts, buyInfoAd common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "processPendingContract", buyInfoAd)
}

// ProcessPendingContract is a paid mutator transaction binding the contract method 0x7866af46.
//
// Solidity: function processPendingContract(buyInfoAd address) returns()
func (_Authority *AuthoritySession) ProcessPendingContract(buyInfoAd common.Address) (*types.Transaction, error) {
	return _Authority.Contract.ProcessPendingContract(&_Authority.TransactOpts, buyInfoAd)
}

// ProcessPendingContract is a paid mutator transaction binding the contract method 0x7866af46.
//
// Solidity: function processPendingContract(buyInfoAd address) returns()
func (_Authority *AuthorityTransactorSession) ProcessPendingContract(buyInfoAd common.Address) (*types.Transaction, error) {
	return _Authority.Contract.ProcessPendingContract(&_Authority.TransactOpts, buyInfoAd)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(newFee uint256) returns()
func (_Authority *AuthorityTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(newFee uint256) returns()
func (_Authority *AuthoritySession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Authority.Contract.SetFee(&_Authority.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(newFee uint256) returns()
func (_Authority *AuthorityTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Authority.Contract.SetFee(&_Authority.TransactOpts, newFee)
}

// SetMaxTimToConfirm is a paid mutator transaction binding the contract method 0x23cdf219.
//
// Solidity: function setMaxTimToConfirm(maxTime uint256) returns()
func (_Authority *AuthorityTransactor) SetMaxTimToConfirm(opts *bind.TransactOpts, maxTime *big.Int) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "setMaxTimToConfirm", maxTime)
}

// SetMaxTimToConfirm is a paid mutator transaction binding the contract method 0x23cdf219.
//
// Solidity: function setMaxTimToConfirm(maxTime uint256) returns()
func (_Authority *AuthoritySession) SetMaxTimToConfirm(maxTime *big.Int) (*types.Transaction, error) {
	return _Authority.Contract.SetMaxTimToConfirm(&_Authority.TransactOpts, maxTime)
}

// SetMaxTimToConfirm is a paid mutator transaction binding the contract method 0x23cdf219.
//
// Solidity: function setMaxTimToConfirm(maxTime uint256) returns()
func (_Authority *AuthorityTransactorSession) SetMaxTimToConfirm(maxTime *big.Int) (*types.Transaction, error) {
	return _Authority.Contract.SetMaxTimToConfirm(&_Authority.TransactOpts, maxTime)
}

// SetMaxTimeToUpload is a paid mutator transaction binding the contract method 0x53e7c13b.
//
// Solidity: function setMaxTimeToUpload(maxTime uint256) returns()
func (_Authority *AuthorityTransactor) SetMaxTimeToUpload(opts *bind.TransactOpts, maxTime *big.Int) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "setMaxTimeToUpload", maxTime)
}

// SetMaxTimeToUpload is a paid mutator transaction binding the contract method 0x53e7c13b.
//
// Solidity: function setMaxTimeToUpload(maxTime uint256) returns()
func (_Authority *AuthoritySession) SetMaxTimeToUpload(maxTime *big.Int) (*types.Transaction, error) {
	return _Authority.Contract.SetMaxTimeToUpload(&_Authority.TransactOpts, maxTime)
}

// SetMaxTimeToUpload is a paid mutator transaction binding the contract method 0x53e7c13b.
//
// Solidity: function setMaxTimeToUpload(maxTime uint256) returns()
func (_Authority *AuthorityTransactorSession) SetMaxTimeToUpload(maxTime *big.Int) (*types.Transaction, error) {
	return _Authority.Contract.SetMaxTimeToUpload(&_Authority.TransactOpts, maxTime)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_Authority *AuthorityTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_Authority *AuthoritySession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Authority.Contract.SetOwner(&_Authority.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(newOwner address) returns()
func (_Authority *AuthorityTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Authority.Contract.SetOwner(&_Authority.TransactOpts, newOwner)
}

// UnblockBuyer is a paid mutator transaction binding the contract method 0x6f8f075c.
//
// Solidity: function unblockBuyer(buyer address) returns()
func (_Authority *AuthorityTransactor) UnblockBuyer(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "unblockBuyer", buyer)
}

// UnblockBuyer is a paid mutator transaction binding the contract method 0x6f8f075c.
//
// Solidity: function unblockBuyer(buyer address) returns()
func (_Authority *AuthoritySession) UnblockBuyer(buyer common.Address) (*types.Transaction, error) {
	return _Authority.Contract.UnblockBuyer(&_Authority.TransactOpts, buyer)
}

// UnblockBuyer is a paid mutator transaction binding the contract method 0x6f8f075c.
//
// Solidity: function unblockBuyer(buyer address) returns()
func (_Authority *AuthorityTransactorSession) UnblockBuyer(buyer common.Address) (*types.Transaction, error) {
	return _Authority.Contract.UnblockBuyer(&_Authority.TransactOpts, buyer)
}

// UnblockSeller is a paid mutator transaction binding the contract method 0xc9041ac9.
//
// Solidity: function unblockSeller(seller address) returns()
func (_Authority *AuthorityTransactor) UnblockSeller(opts *bind.TransactOpts, seller common.Address) (*types.Transaction, error) {
	return _Authority.contract.Transact(opts, "unblockSeller", seller)
}

// UnblockSeller is a paid mutator transaction binding the contract method 0xc9041ac9.
//
// Solidity: function unblockSeller(seller address) returns()
func (_Authority *AuthoritySession) UnblockSeller(seller common.Address) (*types.Transaction, error) {
	return _Authority.Contract.UnblockSeller(&_Authority.TransactOpts, seller)
}

// UnblockSeller is a paid mutator transaction binding the contract method 0xc9041ac9.
//
// Solidity: function unblockSeller(seller address) returns()
func (_Authority *AuthorityTransactorSession) UnblockSeller(seller common.Address) (*types.Transaction, error) {
	return _Authority.Contract.UnblockSeller(&_Authority.TransactOpts, seller)
}
