// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package whitelist

import (
	"math/big"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// VotingABI is the input ABI used to generate the binding from.
const VotingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"eligible\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totaleligible\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"administration\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tally\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reject\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votecast\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalvoted\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelVoting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"neededAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"countResult\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"vote\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_voters\",\"type\":\"address[]\"},{\"name\":\"consensus\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// VotingBin is the compiled bytecode used for deploying new contracts.
const VotingBin = `606060405234156200001057600080fd5b6040516200151138038062001511833981016040528080519060200190919080518201919060200180519060200190919050505b60008083511115156200005357fe5b600090505b825181101562000384576004600084838151811015156200007557fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151562000375576001600460008584815181101515620000e357fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060018054806001018281620001549190620004cd565b916000526020600020900160005b85848151811015156200017157fe5b90602001906020020151909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600654600260008584815181101515620001d257fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550604080519081016040528084838151811015156200023457fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff168152602001600060038111156200026757fe5b81525060036000600654815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690836003811115620002eb57fe5b021790555090505060006005600085848151811015156200030857fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016006600082825401925050819055505b5b808060010191505062000058565b33600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600860016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060006101000a81548160ff021916908360028111156200042557fe5b02179055506000600860006101000a81548160ff021916908360038111156200044a57fe5b0217905550600060018111156200045d57fe5b8260018111156200046a57fe5b14156200048057600654600a81905550620004c2565b6001808111156200048d57fe5b8260018111156200049a57fe5b1415620004c05760016002600654811515620004b257fe5b0401600a81905550620004c1565b5b5b5b5050505062000524565b815481835581811511620004f757818360005260206000209182019101620004f69190620004fc565b5b505050565b6200052191905b808211156200051d57600081600090555060010162000503565b5090565b90565b610fdd80620005346000396000f300606060405236156100e4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630713b139146100e957806314cf9c531461013a57806324b4bdfd146101875780632852b71c146101b05780633847cb59146101dd578063410673e5146102325780634dc415de14610269578063742d4d31146102965780637b6b698c146102e757806381f4b2321461031057806386cfc771146103255780638da5cb5b1461034e5780639e76e022146103a3578063c19d93fb146103da578063da58c7d914610411578063edf26d9b14610489575b600080fd5b34156100f457600080fd5b610120600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506104ec565b604051808215151515815260200191505060405180910390f35b341561014557600080fd5b610171600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061050c565b6040518082815260200191505060405180910390f35b341561019257600080fd5b61019a610524565b6040518082815260200191505060405180910390f35b34156101bb57600080fd5b6101c361052a565b604051808215151515815260200191505060405180910390f35b34156101e857600080fd5b6101f06106fa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561023d57600080fd5b610245610720565b6040518082600381111561025557fe5b60ff16815260200191505060405180910390f35b341561027457600080fd5b61027c610733565b604051808215151515815260200191505060405180910390f35b34156102a157600080fd5b6102cd600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610903565b604051808215151515815260200191505060405180910390f35b34156102f257600080fd5b6102fa610923565b6040518082815260200191505060405180910390f35b341561031b57600080fd5b610323610929565b005b341561033057600080fd5b610338610bf9565b6040518082815260200191505060405180910390f35b341561035957600080fd5b610361610bff565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103ae57600080fd5b6103b6610c25565b604051808260038111156103c657fe5b60ff16815260200191505060405180910390f35b34156103e557600080fd5b6103ed610f0e565b604051808260028111156103fd57fe5b60ff16815260200191505060405180910390f35b341561041c57600080fd5b6104326004808035906020019091905050610f20565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182600381111561047457fe5b60ff1681526020019250505060405180910390f35b341561049457600080fd5b6104aa6004808035906020019091905050610f71565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60046020528060005260406000206000915054906101000a900460ff1681565b60026020528060005260406000206000915090505481565b60065481565b60008080600281111561053957fe5b6000809054906101000a900460ff16600281111561055357fe5b14151561055f57600080fd5b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1680156106025750600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16155b156106f0576001600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160036000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054815260200190815260200160002060000160146101000a81548160ff021916908360038111156106d157fe5b02179055506001600760008282540192505081905550600191506106f5565b600091505b5b5090565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600860009054906101000a900460ff1681565b60008080600281111561074257fe5b6000809054906101000a900460ff16600281111561075c57fe5b14151561076857600080fd5b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16801561080b5750600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16155b156108f9576001600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600260036000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054815260200190815260200160002060000160146101000a81548160ff021916908360038111156108da57fe5b02179055506001600760008282540192505081905550600191506108fe565b600091505b5b5090565b60056020528060005260406000206000915054906101000a900460ff1681565b60075481565b600080600281111561093757fe5b6000809054906101000a900460ff16600281111561095157fe5b14151561095d57600080fd5b600860019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109b957600080fd5b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610a76576001600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016007600082825401925050819055505b6003806000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054815260200190815260200160002060000160146101000a81548160ff02191690836003811115610ae757fe5b021790555060026000806101000a81548160ff02191690836002811115610b0a57fe5b02179055506003600860006101000a81548160ff02191690836003811115610b2e57fe5b0217905550600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166321fe6e70600860009054906101000a900460ff166040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826003811115610bbc57fe5b60ff168152602001915050600060405180830381600087803b1515610be057600080fd5b6102c65a03f11515610bf157600080fd5b5050505b5b50565b600a5481565b600860019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000806000806002811115610c3957fe5b6000809054906101000a900460ff166002811115610c5357fe5b141515610c5f57600080fd5b600754600654141515610c7157600080fd5b6000935060009250600091505b600654821015610db657600560006003600085815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610d1657600080fd5b60016003811115610d2357fe5b6003600084815260200190815260200160002060000160149054906101000a900460ff166003811115610d5257fe5b1415610d5f576001840193505b60026003811115610d6c57fe5b6003600084815260200190815260200160002060000160149054906101000a900460ff166003811115610d9b57fe5b1415610da8576001830192505b5b8180600101925050610c7e565b60016000806101000a81548160ff02191690836002811115610dd457fe5b0217905550600a5484101515610e0d576001600860006101000a81548160ff02191690836003811115610e0357fe5b0217905550610e32565b6002600860006101000a81548160ff02191690836003811115610e2c57fe5b02179055505b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166321fe6e70600860009054906101000a900460ff166040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180826003811115610ebb57fe5b60ff168152602001915050600060405180830381600087803b1515610edf57600080fd5b6102c65a03f11515610ef057600080fd5b505050600860009054906101000a900460ff1694505b5b5050505090565b6000809054906101000a900460ff1681565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16905082565b600181815481101515610f8057fe5b906000526020600020900160005b915054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820d26459731cd6ea6f0e8b3259687ba88a79f2fa95970ab7f912513b22e2c8f4190029`

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _voters []common.Address, consensus uint8) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VotingBin), backend, _owner, _voters, consensus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses( uint256) constant returns(address)
func (_Voting *VotingCaller) Addresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "addresses", arg0)
	return *ret0, err
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses( uint256) constant returns(address)
func (_Voting *VotingSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _Voting.Contract.Addresses(&_Voting.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0xedf26d9b.
//
// Solidity: function addresses( uint256) constant returns(address)
func (_Voting *VotingCallerSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _Voting.Contract.Addresses(&_Voting.CallOpts, arg0)
}

// Addressid is a free data retrieval call binding the contract method 0x14cf9c53.
//
// Solidity: function addressid( address) constant returns(uint256)
func (_Voting *VotingCaller) Addressid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "addressid", arg0)
	return *ret0, err
}

// Addressid is a free data retrieval call binding the contract method 0x14cf9c53.
//
// Solidity: function addressid( address) constant returns(uint256)
func (_Voting *VotingSession) Addressid(arg0 common.Address) (*big.Int, error) {
	return _Voting.Contract.Addressid(&_Voting.CallOpts, arg0)
}

// Addressid is a free data retrieval call binding the contract method 0x14cf9c53.
//
// Solidity: function addressid( address) constant returns(uint256)
func (_Voting *VotingCallerSession) Addressid(arg0 common.Address) (*big.Int, error) {
	return _Voting.Contract.Addressid(&_Voting.CallOpts, arg0)
}

// Administration is a free data retrieval call binding the contract method 0x3847cb59.
//
// Solidity: function administration() constant returns(address)
func (_Voting *VotingCaller) Administration(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "administration")
	return *ret0, err
}

// Administration is a free data retrieval call binding the contract method 0x3847cb59.
//
// Solidity: function administration() constant returns(address)
func (_Voting *VotingSession) Administration() (common.Address, error) {
	return _Voting.Contract.Administration(&_Voting.CallOpts)
}

// Administration is a free data retrieval call binding the contract method 0x3847cb59.
//
// Solidity: function administration() constant returns(address)
func (_Voting *VotingCallerSession) Administration() (common.Address, error) {
	return _Voting.Contract.Administration(&_Voting.CallOpts)
}

// Eligible is a free data retrieval call binding the contract method 0x0713b139.
//
// Solidity: function eligible( address) constant returns(bool)
func (_Voting *VotingCaller) Eligible(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "eligible", arg0)
	return *ret0, err
}

// Eligible is a free data retrieval call binding the contract method 0x0713b139.
//
// Solidity: function eligible( address) constant returns(bool)
func (_Voting *VotingSession) Eligible(arg0 common.Address) (bool, error) {
	return _Voting.Contract.Eligible(&_Voting.CallOpts, arg0)
}

// Eligible is a free data retrieval call binding the contract method 0x0713b139.
//
// Solidity: function eligible( address) constant returns(bool)
func (_Voting *VotingCallerSession) Eligible(arg0 common.Address) (bool, error) {
	return _Voting.Contract.Eligible(&_Voting.CallOpts, arg0)
}

// NeededAmount is a free data retrieval call binding the contract method 0x86cfc771.
//
// Solidity: function neededAmount() constant returns(uint256)
func (_Voting *VotingCaller) NeededAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "neededAmount")
	return *ret0, err
}

// NeededAmount is a free data retrieval call binding the contract method 0x86cfc771.
//
// Solidity: function neededAmount() constant returns(uint256)
func (_Voting *VotingSession) NeededAmount() (*big.Int, error) {
	return _Voting.Contract.NeededAmount(&_Voting.CallOpts)
}

// NeededAmount is a free data retrieval call binding the contract method 0x86cfc771.
//
// Solidity: function neededAmount() constant returns(uint256)
func (_Voting *VotingCallerSession) NeededAmount() (*big.Int, error) {
	return _Voting.Contract.NeededAmount(&_Voting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Voting *VotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Voting *VotingSession) Owner() (common.Address, error) {
	return _Voting.Contract.Owner(&_Voting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Voting *VotingCallerSession) Owner() (common.Address, error) {
	return _Voting.Contract.Owner(&_Voting.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Voting *VotingCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Voting *VotingSession) State() (uint8, error) {
	return _Voting.Contract.State(&_Voting.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Voting *VotingCallerSession) State() (uint8, error) {
	return _Voting.Contract.State(&_Voting.CallOpts)
}

// Tally is a free data retrieval call binding the contract method 0x410673e5.
//
// Solidity: function tally() constant returns(uint8)
func (_Voting *VotingCaller) Tally(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "tally")
	return *ret0, err
}

// Tally is a free data retrieval call binding the contract method 0x410673e5.
//
// Solidity: function tally() constant returns(uint8)
func (_Voting *VotingSession) Tally() (uint8, error) {
	return _Voting.Contract.Tally(&_Voting.CallOpts)
}

// Tally is a free data retrieval call binding the contract method 0x410673e5.
//
// Solidity: function tally() constant returns(uint8)
func (_Voting *VotingCallerSession) Tally() (uint8, error) {
	return _Voting.Contract.Tally(&_Voting.CallOpts)
}

// Totaleligible is a free data retrieval call binding the contract method 0x24b4bdfd.
//
// Solidity: function totaleligible() constant returns(uint256)
func (_Voting *VotingCaller) Totaleligible(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "totaleligible")
	return *ret0, err
}

// Totaleligible is a free data retrieval call binding the contract method 0x24b4bdfd.
//
// Solidity: function totaleligible() constant returns(uint256)
func (_Voting *VotingSession) Totaleligible() (*big.Int, error) {
	return _Voting.Contract.Totaleligible(&_Voting.CallOpts)
}

// Totaleligible is a free data retrieval call binding the contract method 0x24b4bdfd.
//
// Solidity: function totaleligible() constant returns(uint256)
func (_Voting *VotingCallerSession) Totaleligible() (*big.Int, error) {
	return _Voting.Contract.Totaleligible(&_Voting.CallOpts)
}

// Totalvoted is a free data retrieval call binding the contract method 0x7b6b698c.
//
// Solidity: function totalvoted() constant returns(uint256)
func (_Voting *VotingCaller) Totalvoted(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "totalvoted")
	return *ret0, err
}

// Totalvoted is a free data retrieval call binding the contract method 0x7b6b698c.
//
// Solidity: function totalvoted() constant returns(uint256)
func (_Voting *VotingSession) Totalvoted() (*big.Int, error) {
	return _Voting.Contract.Totalvoted(&_Voting.CallOpts)
}

// Totalvoted is a free data retrieval call binding the contract method 0x7b6b698c.
//
// Solidity: function totalvoted() constant returns(uint256)
func (_Voting *VotingCallerSession) Totalvoted() (*big.Int, error) {
	return _Voting.Contract.Totalvoted(&_Voting.CallOpts)
}

// Votecast is a free data retrieval call binding the contract method 0x742d4d31.
//
// Solidity: function votecast( address) constant returns(bool)
func (_Voting *VotingCaller) Votecast(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Voting.contract.Call(opts, out, "votecast", arg0)
	return *ret0, err
}

// Votecast is a free data retrieval call binding the contract method 0x742d4d31.
//
// Solidity: function votecast( address) constant returns(bool)
func (_Voting *VotingSession) Votecast(arg0 common.Address) (bool, error) {
	return _Voting.Contract.Votecast(&_Voting.CallOpts, arg0)
}

// Votecast is a free data retrieval call binding the contract method 0x742d4d31.
//
// Solidity: function votecast( address) constant returns(bool)
func (_Voting *VotingCallerSession) Votecast(arg0 common.Address) (bool, error) {
	return _Voting.Contract.Votecast(&_Voting.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters( uint256) constant returns(addr address, vote uint8)
func (_Voting *VotingCaller) Voters(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr common.Address
	Vote uint8
}, error) {
	ret := new(struct {
		Addr common.Address
		Vote uint8
	})
	out := ret
	err := _Voting.contract.Call(opts, out, "voters", arg0)
	return *ret, err
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters( uint256) constant returns(addr address, vote uint8)
func (_Voting *VotingSession) Voters(arg0 *big.Int) (struct {
	Addr common.Address
	Vote uint8
}, error) {
	return _Voting.Contract.Voters(&_Voting.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters( uint256) constant returns(addr address, vote uint8)
func (_Voting *VotingCallerSession) Voters(arg0 *big.Int) (struct {
	Addr common.Address
	Vote uint8
}, error) {
	return _Voting.Contract.Voters(&_Voting.CallOpts, arg0)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns(bool)
func (_Voting *VotingTransactor) Accept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "accept")
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns(bool)
func (_Voting *VotingSession) Accept() (*types.Transaction, error) {
	return _Voting.Contract.Accept(&_Voting.TransactOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns(bool)
func (_Voting *VotingTransactorSession) Accept() (*types.Transaction, error) {
	return _Voting.Contract.Accept(&_Voting.TransactOpts)
}

// CancelVoting is a paid mutator transaction binding the contract method 0x81f4b232.
//
// Solidity: function cancelVoting() returns()
func (_Voting *VotingTransactor) CancelVoting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "cancelVoting")
}

// CancelVoting is a paid mutator transaction binding the contract method 0x81f4b232.
//
// Solidity: function cancelVoting() returns()
func (_Voting *VotingSession) CancelVoting() (*types.Transaction, error) {
	return _Voting.Contract.CancelVoting(&_Voting.TransactOpts)
}

// CancelVoting is a paid mutator transaction binding the contract method 0x81f4b232.
//
// Solidity: function cancelVoting() returns()
func (_Voting *VotingTransactorSession) CancelVoting() (*types.Transaction, error) {
	return _Voting.Contract.CancelVoting(&_Voting.TransactOpts)
}

// CountResult is a paid mutator transaction binding the contract method 0x9e76e022.
//
// Solidity: function countResult() returns(uint8)
func (_Voting *VotingTransactor) CountResult(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "countResult")
}

// CountResult is a paid mutator transaction binding the contract method 0x9e76e022.
//
// Solidity: function countResult() returns(uint8)
func (_Voting *VotingSession) CountResult() (*types.Transaction, error) {
	return _Voting.Contract.CountResult(&_Voting.TransactOpts)
}

// CountResult is a paid mutator transaction binding the contract method 0x9e76e022.
//
// Solidity: function countResult() returns(uint8)
func (_Voting *VotingTransactorSession) CountResult() (*types.Transaction, error) {
	return _Voting.Contract.CountResult(&_Voting.TransactOpts)
}

// Reject is a paid mutator transaction binding the contract method 0x4dc415de.
//
// Solidity: function reject() returns(bool)
func (_Voting *VotingTransactor) Reject(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "reject")
}

// Reject is a paid mutator transaction binding the contract method 0x4dc415de.
//
// Solidity: function reject() returns(bool)
func (_Voting *VotingSession) Reject() (*types.Transaction, error) {
	return _Voting.Contract.Reject(&_Voting.TransactOpts)
}

// Reject is a paid mutator transaction binding the contract method 0x4dc415de.
//
// Solidity: function reject() returns(bool)
func (_Voting *VotingTransactorSession) Reject() (*types.Transaction, error) {
	return _Voting.Contract.Reject(&_Voting.TransactOpts)
}
