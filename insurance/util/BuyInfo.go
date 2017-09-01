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

// BuyInfoABI is the input ABI used to generate the binding from.
const BuyInfoABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"created\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWaitingForConfirmation\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"encData\",\"type\":\"string\"}],\"name\":\"uploadData\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"freeze\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uploaded\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDeclined\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPubKeySign\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"idHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"decline\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pubKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freezed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFreezed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"encryptedData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWaitingForUpload\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"company\",\"type\":\"address\"},{\"name\":\"buyerPubKey\",\"type\":\"string\"},{\"name\":\"dataH\",\"type\":\"string\"},{\"name\":\"idH\",\"type\":\"string\"},{\"name\":\"authorityH\",\"type\":\"address\"},{\"name\":\"priceV\",\"type\":\"uint256\"},{\"name\":\"feeV\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"returnValue\",\"type\":\"string\"}],\"name\":\"DataUploaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"Closed\",\"type\":\"event\"}]"

// BuyInfoBin is the compiled bytecode used for deploying new contracts.
const BuyInfoBin = `606060405234156200001057600080fd5b6040516200159538038062001595833981016040528080519060200190919080518201919060200180518201919060200180518201919060200180519060200190919080519060200190919080519060200190919050505b326000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555086600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550856005908051906020019062000142929190620001b6565b506000600760006101000a81548160ff021916908360058111156200016357fe5b0217905550846003908051906020019062000180929190620001b6565b50836004908051906020019062000199929190620001b6565b5081600881905550806009819055505b5050505050505062000265565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001f957805160ff19168380011785556200022a565b828001600101855582156200022a579182015b82811115620002295782518255916020019190600101906200020c565b5b5090506200023991906200023d565b5090565b6200026291905b808211156200025e57600081600090555060010162000244565b5090565b90565b61132080620002756000396000f30060606040523615610147576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806308551a531461014c5780631b3012a3146101a1578063325a19f11461023057806347d42dd51461025957806350969f4414610286578063597e1fb5146102e357806362a5af3b1461030c5780637022b58e146103215780637150d8ae1461033657806389832d651461038b5780639b4e88aa146103b45780639f1cf530146103e1578063a035b1fe14610412578063a79a3cee1461043b578063aa41073214610468578063ab040107146104f7578063ac2a5dfd1461050c578063b7540d9f1461059b578063b9469e1a146105c4578063bf7e214f146105f1578063c19d93fb14610646578063c2b6b58c1461067d578063ddca3f43146106aa578063e9fc715a146106d3578063f0bde54914610762575b600080fd5b341561015757600080fd5b61015f61078f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156101ac57600080fd5b6101b46107b5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101f55780820151818401525b6020810190506101d9565b50505050905090810190601f1680156102225780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561023b57600080fd5b610243610853565b6040518082815260200191505060405180910390f35b341561026457600080fd5b61026c610859565b604051808215151515815260200191505060405180910390f35b341561029157600080fd5b6102e1600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190505061088a565b005b34156102ee57600080fd5b6102f6610aaa565b6040518082815260200191505060405180910390f35b341561031757600080fd5b61031f610ab0565b005b341561032c57600080fd5b610334610b97565b005b341561034157600080fd5b610349610d04565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561039657600080fd5b61039e610d29565b6040518082815260200191505060405180910390f35b34156103bf57600080fd5b6103c7610d2f565b604051808215151515815260200191505060405180910390f35b34156103ec57600080fd5b6103f4610d60565b60405180826000191660001916815260200191505060405180910390f35b341561041d57600080fd5b610425610dfb565b6040518082815260200191505060405180910390f35b341561044657600080fd5b61044e610e01565b604051808215151515815260200191505060405180910390f35b341561047357600080fd5b61047b610e62565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104bc5780820151818401525b6020810190506104a0565b50505050905090810190601f1680156104e95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561050257600080fd5b61050a610f00565b005b341561051757600080fd5b61051f61103b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105605780820151818401525b602081019050610544565b50505050905090810190601f16801561058d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156105a657600080fd5b6105ae6110d9565b6040518082815260200191505060405180910390f35b34156105cf57600080fd5b6105d76110df565b604051808215151515815260200191505060405180910390f35b34156105fc57600080fd5b61060461110f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561065157600080fd5b610659611135565b6040518082600581111561066957fe5b60ff16815260200191505060405180910390f35b341561068857600080fd5b610690611148565b604051808215151515815260200191505060405180910390f35b34156106b557600080fd5b6106bd611167565b6040518082815260200191505060405180910390f35b34156106de57600080fd5b6106e661116d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107275780820151818401525b60208101905061070b565b50505050905090810190601f1680156107545780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561076d57600080fd5b61077561120b565b604051808215151515815260200191505060405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561084b5780601f106108205761010080835404028352916020019161084b565b820191906000526020600020905b81548152906001019060200180831161082e57829003601f168201915b505050505081565b600a5481565b60006001600581111561086857fe5b600760009054906101000a900460ff16600581111561088357fe5b1490505b90565b61089261123b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108eb57fe5b600060058111156108f857fe5b600760009054906101000a900460ff16600581111561091357fe5b14151561091c57fe5b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109b25780601f10610987576101008083540402835291602001916109b2565b820191906000526020600020905b81548152906001019060200180831161099557829003601f168201915b50505050509050600081511415156109c657fe5b81600690805190602001906109dc92919061124f565b506001600760006101000a81548160ff021916908360058111156109fc57fe5b021790555042600b819055507f465cebfab7b351fe49711ea387ca88123bad559fcb0cccfbc51cad946962ca4f826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a6b5780820151818401525b602081019050610a4f565b50505050905090810190601f168015610a985780820380516001836020036101000a031916815260200191505b509250505060405180910390a15b5050565b600c5481565b60016005811115610abd57fe5b600760009054906101000a900460ff166005811115610ad857fe5b1480610b085750600580811115610aeb57fe5b600760009054906101000a900460ff166005811115610b0657fe5b145b1515610b1057fe5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b6957fe5b42600d819055506005600760006101000a81548160ff02191690836005811115610b8f57fe5b02179055505b565b60008090506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610c1e576002600760006101000a81548160ff02191690836005811115610c1057fe5b021790555060019050610c9e565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610c9d576003600760006101000a81548160ff02191690836005811115610c9357fe5b0217905550600190505b5b8015610d005742600c819055507fe0449c44a78638ebb7df243f5f79ca5e160e2a1ad65fd2c39fcbfcb866aa23cc600760009054906101000a900460ff1660405180826005811115610cec57fe5b60ff16815260200191505060405180910390a15b5b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b5481565b600060046005811115610d3e57fe5b600760009054906101000a900460ff166005811115610d5957fe5b1490505b90565b6000600260056000604051602001526040518082805460018160011615610100020316600290048015610dca5780601f10610da8576101008083540402835291820191610dca565b820191906000526020600020905b815481529060010190602001808311610db6575b505091505060206040518083038160008661646e5a03f11515610dec57600080fd5b50506040518051905090505b90565b60085481565b600060026005811115610e1057fe5b600760009054906101000a900460ff166005811115610e2b57fe5b1480610e5c575060036005811115610e3f57fe5b600760009054906101000a900460ff166005811115610e5a57fe5b145b90505b90565b60048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ef85780601f10610ecd57610100808354040283529160200191610ef8565b820191906000526020600020905b815481529060010190602001808311610edb57829003601f168201915b505050505081565b60016005811115610f0d57fe5b600760009054906101000a900460ff166005811115610f2857fe5b1480610f585750600580811115610f3b57fe5b600760009054906101000a900460ff166005811115610f5657fe5b145b1515610f6057fe5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610fb957fe5b6004600760006101000a81548160ff02191690836005811115610fd857fe5b021790555042600c819055507fe0449c44a78638ebb7df243f5f79ca5e160e2a1ad65fd2c39fcbfcb866aa23cc600760009054906101000a900460ff166040518082600581111561102557fe5b60ff16815260200191505060405180910390a15b565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110d15780601f106110a6576101008083540402835291602001916110d1565b820191906000526020600020905b8154815290600101906020018083116110b457829003601f168201915b505050505081565b600d5481565b60006005808111156110ed57fe5b600760009054906101000a900460ff16600581111561110857fe5b1490505b90565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900460ff1681565b6000611152610d2f565b806111615750611160610e01565b5b90505b90565b60095481565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112035780601f106111d857610100808354040283529160200191611203565b820191906000526020600020905b8154815290600101906020018083116111e657829003601f168201915b505050505081565b600080600581111561121957fe5b600760009054906101000a900460ff16600581111561123457fe5b1490505b90565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061129057805160ff19168380011785556112be565b828001600101855582156112be579182015b828111156112bd5782518255916020019190600101906112a2565b5b5090506112cb91906112cf565b5090565b6112f191905b808211156112ed5760008160009055506001016112d5565b5090565b905600a165627a7a72305820f0f17e40ef8354dd7a225b17c302e9933658995a67f27ef026391f154fa164e70029`

// DeployBuyInfo deploys a new Ethereum contract, binding an instance of BuyInfo to it.
func DeployBuyInfo(auth *bind.TransactOpts, backend bind.ContractBackend, company common.Address, buyerPubKey string, dataH string, idH string, authorityH common.Address, priceV *big.Int, feeV *big.Int) (common.Address, *types.Transaction, *BuyInfo, error) {
	parsed, err := abi.JSON(strings.NewReader(BuyInfoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BuyInfoBin), backend, company, buyerPubKey, dataH, idH, authorityH, priceV, feeV)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BuyInfo{BuyInfoCaller: BuyInfoCaller{contract: contract}, BuyInfoTransactor: BuyInfoTransactor{contract: contract}}, nil
}

// BuyInfo is an auto generated Go binding around an Ethereum contract.
type BuyInfo struct {
	BuyInfoCaller     // Read-only binding to the contract
	BuyInfoTransactor // Write-only binding to the contract
}

// BuyInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuyInfoSession struct {
	Contract     *BuyInfo          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyInfoCallerSession struct {
	Contract *BuyInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BuyInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyInfoTransactorSession struct {
	Contract     *BuyInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BuyInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyInfoRaw struct {
	Contract *BuyInfo // Generic contract binding to access the raw methods on
}

// BuyInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyInfoCallerRaw struct {
	Contract *BuyInfoCaller // Generic read-only contract binding to access the raw methods on
}

// BuyInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyInfoTransactorRaw struct {
	Contract *BuyInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuyInfo creates a new instance of BuyInfo, bound to a specific deployed contract.
func NewBuyInfo(address common.Address, backend bind.ContractBackend) (*BuyInfo, error) {
	contract, err := bindBuyInfo(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuyInfo{BuyInfoCaller: BuyInfoCaller{contract: contract}, BuyInfoTransactor: BuyInfoTransactor{contract: contract}}, nil
}

// NewBuyInfoCaller creates a new read-only instance of BuyInfo, bound to a specific deployed contract.
func NewBuyInfoCaller(address common.Address, caller bind.ContractCaller) (*BuyInfoCaller, error) {
	contract, err := bindBuyInfo(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BuyInfoCaller{contract: contract}, nil
}

// NewBuyInfoTransactor creates a new write-only instance of BuyInfo, bound to a specific deployed contract.
func NewBuyInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyInfoTransactor, error) {
	contract, err := bindBuyInfo(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BuyInfoTransactor{contract: contract}, nil
}

// bindBuyInfo binds a generic wrapper to an already deployed contract.
func bindBuyInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuyInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyInfo *BuyInfoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BuyInfo.Contract.BuyInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyInfo *BuyInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyInfo.Contract.BuyInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyInfo *BuyInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyInfo.Contract.BuyInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuyInfo *BuyInfoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BuyInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuyInfo *BuyInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuyInfo *BuyInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuyInfo.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BuyInfo *BuyInfoCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BuyInfo *BuyInfoSession) Authority() (common.Address, error) {
	return _BuyInfo.Contract.Authority(&_BuyInfo.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BuyInfo *BuyInfoCallerSession) Authority() (common.Address, error) {
	return _BuyInfo.Contract.Authority(&_BuyInfo.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() constant returns(address)
func (_BuyInfo *BuyInfoCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "buyer")
	return *ret0, err
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() constant returns(address)
func (_BuyInfo *BuyInfoSession) Buyer() (common.Address, error) {
	return _BuyInfo.Contract.Buyer(&_BuyInfo.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() constant returns(address)
func (_BuyInfo *BuyInfoCallerSession) Buyer() (common.Address, error) {
	return _BuyInfo.Contract.Buyer(&_BuyInfo.CallOpts)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() constant returns(uint256)
func (_BuyInfo *BuyInfoCaller) Closed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "closed")
	return *ret0, err
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() constant returns(uint256)
func (_BuyInfo *BuyInfoSession) Closed() (*big.Int, error) {
	return _BuyInfo.Contract.Closed(&_BuyInfo.CallOpts)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() constant returns(uint256)
func (_BuyInfo *BuyInfoCallerSession) Closed() (*big.Int, error) {
	return _BuyInfo.Contract.Closed(&_BuyInfo.CallOpts)
}

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_BuyInfo *BuyInfoCaller) Created(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "created")
	return *ret0, err
}

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_BuyInfo *BuyInfoSession) Created() (*big.Int, error) {
	return _BuyInfo.Contract.Created(&_BuyInfo.CallOpts)
}

// Created is a free data retrieval call binding the contract method 0x325a19f1.
//
// Solidity: function created() constant returns(uint256)
func (_BuyInfo *BuyInfoCallerSession) Created() (*big.Int, error) {
	return _BuyInfo.Contract.Created(&_BuyInfo.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(string)
func (_BuyInfo *BuyInfoCaller) DataHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "dataHash")
	return *ret0, err
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(string)
func (_BuyInfo *BuyInfoSession) DataHash() (string, error) {
	return _BuyInfo.Contract.DataHash(&_BuyInfo.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() constant returns(string)
func (_BuyInfo *BuyInfoCallerSession) DataHash() (string, error) {
	return _BuyInfo.Contract.DataHash(&_BuyInfo.CallOpts)
}

// EncryptedData is a free data retrieval call binding the contract method 0xe9fc715a.
//
// Solidity: function encryptedData() constant returns(string)
func (_BuyInfo *BuyInfoCaller) EncryptedData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "encryptedData")
	return *ret0, err
}

// EncryptedData is a free data retrieval call binding the contract method 0xe9fc715a.
//
// Solidity: function encryptedData() constant returns(string)
func (_BuyInfo *BuyInfoSession) EncryptedData() (string, error) {
	return _BuyInfo.Contract.EncryptedData(&_BuyInfo.CallOpts)
}

// EncryptedData is a free data retrieval call binding the contract method 0xe9fc715a.
//
// Solidity: function encryptedData() constant returns(string)
func (_BuyInfo *BuyInfoCallerSession) EncryptedData() (string, error) {
	return _BuyInfo.Contract.EncryptedData(&_BuyInfo.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_BuyInfo *BuyInfoCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_BuyInfo *BuyInfoSession) Fee() (*big.Int, error) {
	return _BuyInfo.Contract.Fee(&_BuyInfo.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_BuyInfo *BuyInfoCallerSession) Fee() (*big.Int, error) {
	return _BuyInfo.Contract.Fee(&_BuyInfo.CallOpts)
}

// Freezed is a free data retrieval call binding the contract method 0xb7540d9f.
//
// Solidity: function freezed() constant returns(uint256)
func (_BuyInfo *BuyInfoCaller) Freezed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "freezed")
	return *ret0, err
}

// Freezed is a free data retrieval call binding the contract method 0xb7540d9f.
//
// Solidity: function freezed() constant returns(uint256)
func (_BuyInfo *BuyInfoSession) Freezed() (*big.Int, error) {
	return _BuyInfo.Contract.Freezed(&_BuyInfo.CallOpts)
}

// Freezed is a free data retrieval call binding the contract method 0xb7540d9f.
//
// Solidity: function freezed() constant returns(uint256)
func (_BuyInfo *BuyInfoCallerSession) Freezed() (*big.Int, error) {
	return _BuyInfo.Contract.Freezed(&_BuyInfo.CallOpts)
}

// GetPubKeySign is a free data retrieval call binding the contract method 0x9f1cf530.
//
// Solidity: function getPubKeySign() constant returns(bytes32)
func (_BuyInfo *BuyInfoCaller) GetPubKeySign(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "getPubKeySign")
	return *ret0, err
}

// GetPubKeySign is a free data retrieval call binding the contract method 0x9f1cf530.
//
// Solidity: function getPubKeySign() constant returns(bytes32)
func (_BuyInfo *BuyInfoSession) GetPubKeySign() ([32]byte, error) {
	return _BuyInfo.Contract.GetPubKeySign(&_BuyInfo.CallOpts)
}

// GetPubKeySign is a free data retrieval call binding the contract method 0x9f1cf530.
//
// Solidity: function getPubKeySign() constant returns(bytes32)
func (_BuyInfo *BuyInfoCallerSession) GetPubKeySign() ([32]byte, error) {
	return _BuyInfo.Contract.GetPubKeySign(&_BuyInfo.CallOpts)
}

// IdHash is a free data retrieval call binding the contract method 0xaa410732.
//
// Solidity: function idHash() constant returns(string)
func (_BuyInfo *BuyInfoCaller) IdHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "idHash")
	return *ret0, err
}

// IdHash is a free data retrieval call binding the contract method 0xaa410732.
//
// Solidity: function idHash() constant returns(string)
func (_BuyInfo *BuyInfoSession) IdHash() (string, error) {
	return _BuyInfo.Contract.IdHash(&_BuyInfo.CallOpts)
}

// IdHash is a free data retrieval call binding the contract method 0xaa410732.
//
// Solidity: function idHash() constant returns(string)
func (_BuyInfo *BuyInfoCallerSession) IdHash() (string, error) {
	return _BuyInfo.Contract.IdHash(&_BuyInfo.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BuyInfo *BuyInfoCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "isClosed")
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BuyInfo *BuyInfoSession) IsClosed() (bool, error) {
	return _BuyInfo.Contract.IsClosed(&_BuyInfo.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() constant returns(bool)
func (_BuyInfo *BuyInfoCallerSession) IsClosed() (bool, error) {
	return _BuyInfo.Contract.IsClosed(&_BuyInfo.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0xa79a3cee.
//
// Solidity: function isConfirmed() constant returns(bool)
func (_BuyInfo *BuyInfoCaller) IsConfirmed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "isConfirmed")
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0xa79a3cee.
//
// Solidity: function isConfirmed() constant returns(bool)
func (_BuyInfo *BuyInfoSession) IsConfirmed() (bool, error) {
	return _BuyInfo.Contract.IsConfirmed(&_BuyInfo.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0xa79a3cee.
//
// Solidity: function isConfirmed() constant returns(bool)
func (_BuyInfo *BuyInfoCallerSession) IsConfirmed() (bool, error) {
	return _BuyInfo.Contract.IsConfirmed(&_BuyInfo.CallOpts)
}

// IsDeclined is a free data retrieval call binding the contract method 0x9b4e88aa.
//
// Solidity: function isDeclined() constant returns(bool)
func (_BuyInfo *BuyInfoCaller) IsDeclined(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "isDeclined")
	return *ret0, err
}

// IsDeclined is a free data retrieval call binding the contract method 0x9b4e88aa.
//
// Solidity: function isDeclined() constant returns(bool)
func (_BuyInfo *BuyInfoSession) IsDeclined() (bool, error) {
	return _BuyInfo.Contract.IsDeclined(&_BuyInfo.CallOpts)
}

// IsDeclined is a free data retrieval call binding the contract method 0x9b4e88aa.
//
// Solidity: function isDeclined() constant returns(bool)
func (_BuyInfo *BuyInfoCallerSession) IsDeclined() (bool, error) {
	return _BuyInfo.Contract.IsDeclined(&_BuyInfo.CallOpts)
}

// IsFreezed is a free data retrieval call binding the contract method 0xb9469e1a.
//
// Solidity: function isFreezed() constant returns(bool)
func (_BuyInfo *BuyInfoCaller) IsFreezed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "isFreezed")
	return *ret0, err
}

// IsFreezed is a free data retrieval call binding the contract method 0xb9469e1a.
//
// Solidity: function isFreezed() constant returns(bool)
func (_BuyInfo *BuyInfoSession) IsFreezed() (bool, error) {
	return _BuyInfo.Contract.IsFreezed(&_BuyInfo.CallOpts)
}

// IsFreezed is a free data retrieval call binding the contract method 0xb9469e1a.
//
// Solidity: function isFreezed() constant returns(bool)
func (_BuyInfo *BuyInfoCallerSession) IsFreezed() (bool, error) {
	return _BuyInfo.Contract.IsFreezed(&_BuyInfo.CallOpts)
}

// IsWaitingForConfirmation is a free data retrieval call binding the contract method 0x47d42dd5.
//
// Solidity: function isWaitingForConfirmation() constant returns(bool)
func (_BuyInfo *BuyInfoCaller) IsWaitingForConfirmation(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "isWaitingForConfirmation")
	return *ret0, err
}

// IsWaitingForConfirmation is a free data retrieval call binding the contract method 0x47d42dd5.
//
// Solidity: function isWaitingForConfirmation() constant returns(bool)
func (_BuyInfo *BuyInfoSession) IsWaitingForConfirmation() (bool, error) {
	return _BuyInfo.Contract.IsWaitingForConfirmation(&_BuyInfo.CallOpts)
}

// IsWaitingForConfirmation is a free data retrieval call binding the contract method 0x47d42dd5.
//
// Solidity: function isWaitingForConfirmation() constant returns(bool)
func (_BuyInfo *BuyInfoCallerSession) IsWaitingForConfirmation() (bool, error) {
	return _BuyInfo.Contract.IsWaitingForConfirmation(&_BuyInfo.CallOpts)
}

// IsWaitingForUpload is a free data retrieval call binding the contract method 0xf0bde549.
//
// Solidity: function isWaitingForUpload() constant returns(bool)
func (_BuyInfo *BuyInfoCaller) IsWaitingForUpload(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "isWaitingForUpload")
	return *ret0, err
}

// IsWaitingForUpload is a free data retrieval call binding the contract method 0xf0bde549.
//
// Solidity: function isWaitingForUpload() constant returns(bool)
func (_BuyInfo *BuyInfoSession) IsWaitingForUpload() (bool, error) {
	return _BuyInfo.Contract.IsWaitingForUpload(&_BuyInfo.CallOpts)
}

// IsWaitingForUpload is a free data retrieval call binding the contract method 0xf0bde549.
//
// Solidity: function isWaitingForUpload() constant returns(bool)
func (_BuyInfo *BuyInfoCallerSession) IsWaitingForUpload() (bool, error) {
	return _BuyInfo.Contract.IsWaitingForUpload(&_BuyInfo.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BuyInfo *BuyInfoCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BuyInfo *BuyInfoSession) Price() (*big.Int, error) {
	return _BuyInfo.Contract.Price(&_BuyInfo.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_BuyInfo *BuyInfoCallerSession) Price() (*big.Int, error) {
	return _BuyInfo.Contract.Price(&_BuyInfo.CallOpts)
}

// PubKey is a free data retrieval call binding the contract method 0xac2a5dfd.
//
// Solidity: function pubKey() constant returns(string)
func (_BuyInfo *BuyInfoCaller) PubKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "pubKey")
	return *ret0, err
}

// PubKey is a free data retrieval call binding the contract method 0xac2a5dfd.
//
// Solidity: function pubKey() constant returns(string)
func (_BuyInfo *BuyInfoSession) PubKey() (string, error) {
	return _BuyInfo.Contract.PubKey(&_BuyInfo.CallOpts)
}

// PubKey is a free data retrieval call binding the contract method 0xac2a5dfd.
//
// Solidity: function pubKey() constant returns(string)
func (_BuyInfo *BuyInfoCallerSession) PubKey() (string, error) {
	return _BuyInfo.Contract.PubKey(&_BuyInfo.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_BuyInfo *BuyInfoCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "seller")
	return *ret0, err
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_BuyInfo *BuyInfoSession) Seller() (common.Address, error) {
	return _BuyInfo.Contract.Seller(&_BuyInfo.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() constant returns(address)
func (_BuyInfo *BuyInfoCallerSession) Seller() (common.Address, error) {
	return _BuyInfo.Contract.Seller(&_BuyInfo.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BuyInfo *BuyInfoCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BuyInfo *BuyInfoSession) State() (uint8, error) {
	return _BuyInfo.Contract.State(&_BuyInfo.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BuyInfo *BuyInfoCallerSession) State() (uint8, error) {
	return _BuyInfo.Contract.State(&_BuyInfo.CallOpts)
}

// Uploaded is a free data retrieval call binding the contract method 0x89832d65.
//
// Solidity: function uploaded() constant returns(uint256)
func (_BuyInfo *BuyInfoCaller) Uploaded(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BuyInfo.contract.Call(opts, out, "uploaded")
	return *ret0, err
}

// Uploaded is a free data retrieval call binding the contract method 0x89832d65.
//
// Solidity: function uploaded() constant returns(uint256)
func (_BuyInfo *BuyInfoSession) Uploaded() (*big.Int, error) {
	return _BuyInfo.Contract.Uploaded(&_BuyInfo.CallOpts)
}

// Uploaded is a free data retrieval call binding the contract method 0x89832d65.
//
// Solidity: function uploaded() constant returns(uint256)
func (_BuyInfo *BuyInfoCallerSession) Uploaded() (*big.Int, error) {
	return _BuyInfo.Contract.Uploaded(&_BuyInfo.CallOpts)
}

// Confirm is a paid mutator transaction binding the contract method 0x7022b58e.
//
// Solidity: function confirm() returns()
func (_BuyInfo *BuyInfoTransactor) Confirm(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyInfo.contract.Transact(opts, "confirm")
}

// Confirm is a paid mutator transaction binding the contract method 0x7022b58e.
//
// Solidity: function confirm() returns()
func (_BuyInfo *BuyInfoSession) Confirm() (*types.Transaction, error) {
	return _BuyInfo.Contract.Confirm(&_BuyInfo.TransactOpts)
}

// Confirm is a paid mutator transaction binding the contract method 0x7022b58e.
//
// Solidity: function confirm() returns()
func (_BuyInfo *BuyInfoTransactorSession) Confirm() (*types.Transaction, error) {
	return _BuyInfo.Contract.Confirm(&_BuyInfo.TransactOpts)
}

// Decline is a paid mutator transaction binding the contract method 0xab040107.
//
// Solidity: function decline() returns()
func (_BuyInfo *BuyInfoTransactor) Decline(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyInfo.contract.Transact(opts, "decline")
}

// Decline is a paid mutator transaction binding the contract method 0xab040107.
//
// Solidity: function decline() returns()
func (_BuyInfo *BuyInfoSession) Decline() (*types.Transaction, error) {
	return _BuyInfo.Contract.Decline(&_BuyInfo.TransactOpts)
}

// Decline is a paid mutator transaction binding the contract method 0xab040107.
//
// Solidity: function decline() returns()
func (_BuyInfo *BuyInfoTransactorSession) Decline() (*types.Transaction, error) {
	return _BuyInfo.Contract.Decline(&_BuyInfo.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_BuyInfo *BuyInfoTransactor) Freeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuyInfo.contract.Transact(opts, "freeze")
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_BuyInfo *BuyInfoSession) Freeze() (*types.Transaction, error) {
	return _BuyInfo.Contract.Freeze(&_BuyInfo.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_BuyInfo *BuyInfoTransactorSession) Freeze() (*types.Transaction, error) {
	return _BuyInfo.Contract.Freeze(&_BuyInfo.TransactOpts)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(encData string) returns()
func (_BuyInfo *BuyInfoTransactor) UploadData(opts *bind.TransactOpts, encData string) (*types.Transaction, error) {
	return _BuyInfo.contract.Transact(opts, "uploadData", encData)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(encData string) returns()
func (_BuyInfo *BuyInfoSession) UploadData(encData string) (*types.Transaction, error) {
	return _BuyInfo.Contract.UploadData(&_BuyInfo.TransactOpts, encData)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(encData string) returns()
func (_BuyInfo *BuyInfoTransactorSession) UploadData(encData string) (*types.Transaction, error) {
	return _BuyInfo.Contract.UploadData(&_BuyInfo.TransactOpts, encData)
}
