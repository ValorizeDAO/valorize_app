// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// BondingCurveABI is the input ABI used to generate the binding from.
const BondingCurveABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_connectorBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_connectorWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"calculatePurchaseReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_connectorBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_connectorWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_sellAmount\",\"type\":\"uint256\"}],\"name\":\"calculateSaleReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BondingCurveFuncSigs maps the 4-byte function signature to its string representation.
var BondingCurveFuncSigs = map[string]string{
	"29a00e7c": "calculatePurchaseReturn(uint256,uint256,uint32,uint256)",
	"49f9b0f7": "calculateSaleReturn(uint256,uint256,uint32,uint256)",
}

// BondingCurve is an auto generated Go binding around an Ethereum contract.
type BondingCurve struct {
	BondingCurveCaller     // Read-only binding to the contract
	BondingCurveTransactor // Write-only binding to the contract
	BondingCurveFilterer   // Log filterer for contract events
}

// BondingCurveCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondingCurveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingCurveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondingCurveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingCurveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondingCurveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingCurveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondingCurveSession struct {
	Contract     *BondingCurve     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondingCurveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondingCurveCallerSession struct {
	Contract *BondingCurveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BondingCurveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondingCurveTransactorSession struct {
	Contract     *BondingCurveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BondingCurveRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondingCurveRaw struct {
	Contract *BondingCurve // Generic contract binding to access the raw methods on
}

// BondingCurveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondingCurveCallerRaw struct {
	Contract *BondingCurveCaller // Generic read-only contract binding to access the raw methods on
}

// BondingCurveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondingCurveTransactorRaw struct {
	Contract *BondingCurveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondingCurve creates a new instance of BondingCurve, bound to a specific deployed contract.
func NewBondingCurve(address common.Address, backend bind.ContractBackend) (*BondingCurve, error) {
	contract, err := bindBondingCurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BondingCurve{BondingCurveCaller: BondingCurveCaller{contract: contract}, BondingCurveTransactor: BondingCurveTransactor{contract: contract}, BondingCurveFilterer: BondingCurveFilterer{contract: contract}}, nil
}

// NewBondingCurveCaller creates a new read-only instance of BondingCurve, bound to a specific deployed contract.
func NewBondingCurveCaller(address common.Address, caller bind.ContractCaller) (*BondingCurveCaller, error) {
	contract, err := bindBondingCurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondingCurveCaller{contract: contract}, nil
}

// NewBondingCurveTransactor creates a new write-only instance of BondingCurve, bound to a specific deployed contract.
func NewBondingCurveTransactor(address common.Address, transactor bind.ContractTransactor) (*BondingCurveTransactor, error) {
	contract, err := bindBondingCurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondingCurveTransactor{contract: contract}, nil
}

// NewBondingCurveFilterer creates a new log filterer instance of BondingCurve, bound to a specific deployed contract.
func NewBondingCurveFilterer(address common.Address, filterer bind.ContractFilterer) (*BondingCurveFilterer, error) {
	contract, err := bindBondingCurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondingCurveFilterer{contract: contract}, nil
}

// bindBondingCurve binds a generic wrapper to an already deployed contract.
func bindBondingCurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BondingCurveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondingCurve *BondingCurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondingCurve.Contract.BondingCurveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondingCurve *BondingCurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingCurve.Contract.BondingCurveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondingCurve *BondingCurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondingCurve.Contract.BondingCurveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondingCurve *BondingCurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondingCurve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondingCurve *BondingCurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingCurve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondingCurve *BondingCurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondingCurve.Contract.contract.Transact(opts, method, params...)
}

// CalculatePurchaseReturn is a free data retrieval call binding the contract method 0x29a00e7c.
//
// Solidity: function calculatePurchaseReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _depositAmount) view returns(uint256)
func (_BondingCurve *BondingCurveCaller) CalculatePurchaseReturn(opts *bind.CallOpts, _supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _depositAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BondingCurve.contract.Call(opts, &out, "calculatePurchaseReturn", _supply, _connectorBalance, _connectorWeight, _depositAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePurchaseReturn is a free data retrieval call binding the contract method 0x29a00e7c.
//
// Solidity: function calculatePurchaseReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _depositAmount) view returns(uint256)
func (_BondingCurve *BondingCurveSession) CalculatePurchaseReturn(_supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _depositAmount *big.Int) (*big.Int, error) {
	return _BondingCurve.Contract.CalculatePurchaseReturn(&_BondingCurve.CallOpts, _supply, _connectorBalance, _connectorWeight, _depositAmount)
}

// CalculatePurchaseReturn is a free data retrieval call binding the contract method 0x29a00e7c.
//
// Solidity: function calculatePurchaseReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _depositAmount) view returns(uint256)
func (_BondingCurve *BondingCurveCallerSession) CalculatePurchaseReturn(_supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _depositAmount *big.Int) (*big.Int, error) {
	return _BondingCurve.Contract.CalculatePurchaseReturn(&_BondingCurve.CallOpts, _supply, _connectorBalance, _connectorWeight, _depositAmount)
}

// CalculateSaleReturn is a free data retrieval call binding the contract method 0x49f9b0f7.
//
// Solidity: function calculateSaleReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _sellAmount) view returns(uint256)
func (_BondingCurve *BondingCurveCaller) CalculateSaleReturn(opts *bind.CallOpts, _supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _sellAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BondingCurve.contract.Call(opts, &out, "calculateSaleReturn", _supply, _connectorBalance, _connectorWeight, _sellAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSaleReturn is a free data retrieval call binding the contract method 0x49f9b0f7.
//
// Solidity: function calculateSaleReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _sellAmount) view returns(uint256)
func (_BondingCurve *BondingCurveSession) CalculateSaleReturn(_supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _sellAmount *big.Int) (*big.Int, error) {
	return _BondingCurve.Contract.CalculateSaleReturn(&_BondingCurve.CallOpts, _supply, _connectorBalance, _connectorWeight, _sellAmount)
}

// CalculateSaleReturn is a free data retrieval call binding the contract method 0x49f9b0f7.
//
// Solidity: function calculateSaleReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _sellAmount) view returns(uint256)
func (_BondingCurve *BondingCurveCallerSession) CalculateSaleReturn(_supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _sellAmount *big.Int) (*big.Int, error) {
	return _BondingCurve.Contract.CalculateSaleReturn(&_BondingCurve.CallOpts, _supply, _connectorBalance, _connectorWeight, _sellAmount)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// CreatorTokenABI is the input ABI used to generate the binding from.
const CreatorTokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveRatio\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_To\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountDeposited\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountDistributedToBuyer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountDistributedToOwner\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyNewTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_connectorBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_connectorWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"calculatePurchaseReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_connectorBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_connectorWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_sellAmount\",\"type\":\"uint256\"}],\"name\":\"calculateSaleReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"calculateTokenBuyReturns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_newPercentage\",\"type\":\"uint8\"}],\"name\":\"changeFounderPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToCheckAgainst\",\"type\":\"uint256\"}],\"name\":\"checkAndReturnInitialContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToUseAsTokenBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"founderPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sellTokensForEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CreatorTokenFuncSigs maps the 4-byte function signature to its string representation.
var CreatorTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"b593217c": "buyNewTokens()",
	"29a00e7c": "calculatePurchaseReturn(uint256,uint256,uint32,uint256)",
	"49f9b0f7": "calculateSaleReturn(uint256,uint256,uint32,uint256)",
	"a73d5647": "calculateTokenBuyReturns(uint256)",
	"0cc4e3a7": "changeFounderPercentage(uint8)",
	"ff50e196": "checkAndReturnInitialContractBalance(uint256,uint256)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"17cf6914": "founderPercentage()",
	"70ed0ada": "getEthBalance()",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"a10954fe": "reserveBalance()",
	"0c7d5cd8": "reserveRatio()",
	"b3e7a384": "sellTokensForEth(uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
}

// CreatorTokenBin is the compiled bytecode used for deploying new contracts.
var CreatorTokenBin = "0x60a0604052670de0b6b3a76400006086556088805461ff00191690553480156200002857600080fd5b506040516200386c3803806200386c8339810160408190526200004b9162000a9f565b6001641c35fedd1560601b0360209081556001646c3390ecc9605e1b036021556001640cf801476160611b0360225560016431bdb23e1d605f1b0360235560016502fb1d8fe083605b1b0360245560016505b771955b37605a1b036025556001650af67a93bb5160591b0360265560016515060c256cb360581b036027556001651428a2f98d7360581b036028556001654d515663970960561b03602955600165944620b0e70f60551b03602a55600166011c592761c66760541b03602b5560016602214d10d014eb60531b03602c55600166020ade36b7dbef60531b03602d5560016603eab73b3bbfe360521b03602e556001660782ee3593f6d760511b03602f556001661ccf4b44bb4821604f1b0360305560016606e7f88ad8a77760511b0360315560016669f3d1c921891d604d1b03603255600166cb2ff529eb71e5604c1b03603355600166c2d415c3db974b604c1b0360345560016702eb40f9f620fda7604a1b0360355560016705990681d961a1eb60491b03603655600167055e12902701414760491b0360375560016714962dee9dc9764160471b0360385560016704ef57b9b560fab560491b0360395560016712ed7b32a58f552b60471b03603a556001679131271922eaa60760441b03603b556001678b380f3558668c4760441b03603c556001680215f77c045fbe885760421b03603d556001600160831b03603e556001670f577eded5773a1160471b03603f556001680eb5ec597592befbf5603f1b036040556001681c35fedd14b861eb05603e1b036041556001683619c87664579bc94b603d1b0360425560016867c00a3b07ffc01fd7603c1b03604355600168c6f6c8f8739773a7a5603b1b03604455600168bec763f8209b7a72b1603b1b0360455560016902dbb8caad9b7097b91b60391b03604655600169057b3d49dda84556d6f760381b03604755600169054183095b2c8ececf3160381b036048556001690a14517cc6b9457111ef60371b0360495560016913545598e5c23276ccf160361b03604a556001692511882c39c3adea96ff60351b03604b55600169471649d87199aa99075760341b03604c557004429a21a029d4c1457cfbffffffffffff604d55700415bc6d6fb7dd71af2cb3ffffffffffff604e557003eab73b3bbfe282243ce1ffffffffffff604f557003c1771ac9fb6b4c18e229ffffffffffff605055700399e96897690418f785257fffffffffff605155700373fc456c53bb779bf0ea9fffffffffff60525570034f9e8e490c48e67e6ab8bfffffffffff60535570032cbfd4a7adc790560b3337ffffffffff60545570030b50570f6e5d2acca94613ffffffffff6055557002eb40f9f620fda6b56c2861ffffffffff6056557002cc8340ecb0d0f520a6af58ffffffffff6057557002af09481380a0a35cf1ba02ffffffffff605855700292c5bdd3b92ec810287b1b3fffffffff605955700277abdcdab07d5a77ac6d6b9fffffffff605a5570025daf6654b1eaa55fd64df5efffffffff605b55700244c49c648baa98192dce88b7ffffffff605c5570022ce03cd5619a311b2471268bffffffff605d55700215f77c045fbe885654a44a0fffffffff605e556001600160811b03605f557001eaefdbdaaee7421fc4d3ede5ffffffff6060557001d6bd8b2eb257df7e8ca57b09bfffffff6061557001c35fedd14b861eb0443f7f133fffffff6062557001b0ce43b322bcde4a56e8ada5afffffff60635570019f0028ec1fff007f5a195a39dfffffff60645570018ded91f0e72ee74f49b15ba527ffffff60655570017d8ec7f04136f4e5615fd41a63ffffff60665570016ddc6556cdb84bdc8d12d22e6fffffff60675570015ecf52776a1155b5bd8395814f7fffff60685570015060c256cb23b3b3cc3754cf40ffffff6069557001428a2f98d728ae223ddab715be3fffff606a5570013545598e5c23276ccf0ede68034fffff606b557001288c4161ce1d6f54b7f61081194fffff606c5570011c592761c666aa641d5a01a40f17ffff606d55700110a688680a7530515f3e6e6cfdcdffff606e557001056f1b5bedf75c6bcb2ce8aed428ffff606f556ffaadceceeff8a0890f3875f008277fff6070556ff05dc6b27edad306388a600f6ba0bfff6071556fe67a5a25da41063de1495d5b18cdbfff6072556fdcff115b14eedde6fc3aa5353f2e4fff6073556fd3e7a3924312399f9aae2e0f868f8fff6074556fcb2ff529eb71e41582cccd5a1ee26fff6075556fc2d415c3db974ab32a51840c0b67edff6076556fbad03e7d883f69ad5b0a186184e06bff6077556fb320d03b2c343d4829abd6075f0cc5ff6078556fabc25204e02828d73c6e80bcdb1a95bf6079556fa4b16f74ee4bb2040a1ec6c15fbbf2df607a556f9deaf736ac1f569deb1b5ae3f36c130f607b556f976bd9952c7aa957f5937d790ef65037607c556f9131271922eaa6064b73a22d0bd4f2bf607d556f8b380f3558668c46c91c49a2f8e967b9607e556f857ddf0117efa215952912839f6473e6607f558251839183916200078e916083919085019062000942565b508051620007a490608490602084019062000942565b505050620007c1620007bb6200080460201b60201c565b62000808565b8315620007e657620007d433856200085a565b83608654620007e4919062000b1d565b505b50506080919091526088805460ff1916600a17905560875562000b97565b3390565b608580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038216620008b55760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b8060826000828254620008c9919062000b1d565b90915550506001600160a01b03821660009081526080602052604081208054839290620008f890849062000b1d565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b828054620009509062000b44565b90600052602060002090601f016020900481019282620009745760008555620009bf565b82601f106200098f57805160ff1916838001178555620009bf565b82800160010185558215620009bf579182015b82811115620009bf578251825591602001919060010190620009a2565b50620009cd929150620009d1565b5090565b5b80821115620009cd5760008155600101620009d2565b600082601f830112620009fa57600080fd5b81516001600160401b038082111562000a175762000a1762000b81565b604051601f8301601f19908116603f0116810190828211818310171562000a425762000a4262000b81565b8160405283815260209250868385880101111562000a5f57600080fd5b600091505b8382101562000a83578582018301518183018401529082019062000a64565b8382111562000a955760008385830101525b9695505050505050565b6000806000806080858703121562000ab657600080fd5b84516020860151604087015191955093506001600160401b038082111562000add57600080fd5b62000aeb88838901620009e8565b9350606087015191508082111562000b0257600080fd5b5062000b1187828801620009e8565b91505092959194509250565b6000821982111562000b3f57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c9082168062000b5957607f821691505b6020821081141562000b7b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b608051612cbc62000bb060003960005050612cbc6000f3fe6080604052600436106101665760003560e01c806370ed0ada116100d1578063a73d56471161008a578063b593217c11610064578063b593217c14610427578063dd62ed3e1461042f578063f2fde38b14610475578063ff50e1961461049557600080fd5b8063a73d5647146103b2578063a9059cbb146103e7578063b3e7a3841461040757600080fd5b806370ed0ada14610317578063715018a61461032a5780638da5cb5b1461033f57806395d89b4114610367578063a10954fe1461037c578063a457c2d71461039257600080fd5b806323b872dd1161012357806323b872dd1461024d57806329a00e7c1461026d578063313ce5671461028d57806339509351146102a157806349f9b0f7146102c157806370a08231146102e157600080fd5b806306fdde031461016b578063095ea7b3146101965780630c7d5cd8146101c65780630cc4e3a7146101ea57806317cf69141461020c57806318160ddd14610238575b600080fd5b34801561017757600080fd5b506101806104b5565b60405161018d9190612a6c565b60405180910390f35b3480156101a257600080fd5b506101b66101b136600461299c565b610547565b604051901515815260200161018d565b3480156101d257600080fd5b506101dc60875481565b60405190815260200161018d565b3480156101f657600080fd5b5061020a610205366004612a49565b61055e565b005b34801561021857600080fd5b506088546102269060ff1681565b60405160ff909116815260200161018d565b34801561024457600080fd5b506082546101dc565b34801561025957600080fd5b506101b6610268366004612960565b61060c565b34801561027957600080fd5b506101dc610288366004612a01565b6106b6565b34801561029957600080fd5b506012610226565b3480156102ad57600080fd5b506101b66102bc36600461299c565b610782565b3480156102cd57600080fd5b506101dc6102dc366004612a01565b6107be565b3480156102ed57600080fd5b506101dc6102fc366004612914565b6001600160a01b031660009081526080602052604090205490565b34801561032357600080fd5b50476101dc565b34801561033657600080fd5b5061020a61089f565b34801561034b57600080fd5b506085546040516001600160a01b03909116815260200161018d565b34801561037357600080fd5b506101806108d5565b34801561038857600080fd5b506101dc60865481565b34801561039e57600080fd5b506101b66103ad36600461299c565b6108e4565b3480156103be57600080fd5b506103d26103cd3660046129c6565b61097d565b6040805192835260208301919091520161018d565b3480156103f357600080fd5b506101b661040236600461299c565b6109c3565b34801561041357600080fd5b5061020a6104223660046129c6565b6109d0565b61020a610b7f565b34801561043b57600080fd5b506101dc61044a366004612936565b6001600160a01b03918216600090815260816020908152604080832093909416825291909152205490565b34801561048157600080fd5b5061020a610490366004612914565b610be9565b3480156104a157600080fd5b506101dc6104b03660046129df565b610c81565b6060608380546104c490612bdf565b80601f01602080910402602001604051908101604052809291908181526020018280546104f090612bdf565b801561053d5780601f106105125761010080835404028352916020019161053d565b820191906000526020600020905b81548152906001019060200180831161052057829003601f168201915b5050505050905090565b6000610554338484610ca8565b5060015b92915050565b6085546001600160a01b031633146105915760405162461bcd60e51b815260040161058890612ac1565b60405180910390fd5b60648160ff1611156105f65760405162461bcd60e51b815260206004820152602860248201527f466f756e6465722070657263656e74616765206d757374206265206c6573732060448201526707468616e203130360c41b6064820152608401610588565b6088805460ff191660ff92909216919091179055565b6000610619848484610dcd565b6001600160a01b03841660009081526081602090815260408083203384529091529020548281101561069e5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152608401610588565b6106ab8533858403610ca8565b506001949350505050565b600080851180156106c75750600084115b80156106d9575060008363ffffffff16115b80156106ee5750620f424063ffffffff841611155b6106f757600080fd5b816107045750600061077a565b63ffffffff8316620f4240141561073157836107208387612b69565b61072a9190612b33565b905061077a565b6000808061073f8786612af6565b9050610750818888620f4240610f9c565b9093509150600060ff8316610765858b612b69565b901c90506107738982612b88565b9450505050505b949350505050565b3360008181526081602090815260408083206001600160a01b038716845290915281205490916105549185906107b9908690612af6565b610ca8565b600080851180156107cf5750600084115b80156107e1575060008363ffffffff16115b80156107f65750620f424063ffffffff841611155b80156108025750848211155b61080b57600080fd5b816108185750600061077a565b8482141561082757508261077a565b63ffffffff8316620f4240141561084357846107208386612b69565b600080806108518589612b88565b90506108628882620f424089610f9c565b909350915060006108738489612b69565b905060ff831688901b846108878284612b88565b6108919190612b33565b9a9950505050505050505050565b6085546001600160a01b031633146108c95760405162461bcd60e51b815260040161058890612ac1565b6108d3600061111c565b565b6060608480546104c490612bdf565b3360009081526081602090815260408083206001600160a01b0386168452909152812054828110156109665760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610588565b6109733385858403610ca8565b5060019392505050565b60008060006109a561098e60825490565b61099c866104b08147612af6565b608754876106b6565b6088549091506109b990829060ff1661116e565b9250925050915091565b6000610554338484610dcd565b60008111610a205760405162461bcd60e51b815260206004820152601860248201527f416d6f756e74206d757374206265206e6f6e2d7a65726f2e00000000000000006044820152606401610588565b33600090815260806020526040902054811115610a7f5760405162461bcd60e51b815260206004820152601960248201527f6e6f7420656e6f75676820746f6b656e7320746f2073656c6c000000000000006044820152606401610588565b608854610100900460ff1615610a9457600080fd5b6088805461ff0019166101001790556000610aae826111b8565b604051909150339082156108fc029083906000818181858888f1935050505015610b335780608654610ae09190612b88565b608655610aed33836111d0565b60408051338152602081018490529081018290527f23ff0e75edf108e3d0392d92e13e8c8a868ef19001bd49f9e94876dc46dff87f9060600160405180910390a1610b70565b60405162461bcd60e51b81526020600482015260126024820152711dda5d1a191c985dda5b99c819985a5b195960721b6044820152606401610588565b50506088805461ff0019169055565b60003411610bcf5760405162461bcd60e51b815260206004820152601b60248201527f4d7573742073656e642045544820746f2062757920746f6b656e7300000000006044820152606401610588565b6000610bda34611316565b9050610be68134611337565b50565b6085546001600160a01b03163314610c135760405162461bcd60e51b815260040161058890612ac1565b6001600160a01b038116610c785760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610588565b610be68161111c565b808281141561055857600080610c9b856001806002610f9c565b5060201c95945050505050565b6001600160a01b038316610d0a5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610588565b6001600160a01b038216610d6b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610588565b6001600160a01b0383811660008181526081602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b038316610e315760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610588565b6001600160a01b038216610e935760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610588565b6001600160a01b03831660009081526080602052604090205481811015610f0b5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610588565b6001600160a01b03808516600090815260806020526040808220858503905591851681529081208054849290610f42908490612af6565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f8e91815260200190565b60405180910390a350505050565b600080600160811b8610610ff25760405162461bcd60e51b815260206004820152601860248201527f626173654e2065786365656473206d61782076616c75652e00000000000000006044820152606401610588565b848610156110425760405162461bcd60e51b815260206004820152601c60248201527f4261736573203c203120617265206e6f7420737570706f727465642e000000006044820152606401610588565b600080866110546001607f1b8a612b69565b61105e9190612b33565b905070015bf0a8b1457695355fb8ac404e7a79e381101561108957611082816113eb565b9150611095565b61109281611964565b91505b60008563ffffffff168763ffffffff16846110b09190612b69565b6110ba9190612b33565b9050600160831b8110156110df576110d181611a53565b607f94509450505050611113565b60006110ea8261207e565b90506111066110fa82607f612b9f565b60ff1683901c82612143565b9550935061111392505050565b94509492505050565b608580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080606461117d8482612b9f565b61118a9060ff1686612b69565b6111949190612b33565b915060646111a560ff851686612b69565b6111af9190612b33565b90509250929050565b60006105586111c660825490565b47608754856107be565b6001600160a01b0382166112305760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610588565b6001600160a01b038216600090815260806020526040902054818110156112a45760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610588565b6001600160a01b03831660009081526080602052604081208383039055608280548492906112d3908490612b88565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610dc0565b600061055861132460825490565b61132e8447610c81565b608754856106b6565b608854600090819061134d90859060ff1661116e565b9150915061135b33836127aa565b6113766113706085546001600160a01b031690565b826127aa565b60006113828284612af6565b9050836086546113929190612af6565b608655604080513381526020810186905290810182905260608101849052608081018390527f19f5f791ee407773427bf7b970bbbc3375065c32edd1ab142e23a84f94b0719b9060a00160405180910390a15050505050565b6000808080806fd3094c70f034de4b96ff7d5b6f99fcd88610611444576114166001607e1b85612af6565b93506fd3094c70f034de4b96ff7d5b6f99fcd86114376001607f1b88612b69565b6114419190612b33565b95505b6fa45af1e1f40c333b3de1db4dd55f29a78610611497576114696001607d1b85612af6565b93506fa45af1e1f40c333b3de1db4dd55f29a761148a6001607f1b88612b69565b6114949190612b33565b95505b6f910b022db7ae67ce76b441c27035c6a186106114ea576114bc6001607c1b85612af6565b93506f910b022db7ae67ce76b441c27035c6a16114dd6001607f1b88612b69565b6114e79190612b33565b95505b6f88415abbe9a76bead8d00cf112e4d4a8861061153d5761150f6001607b1b85612af6565b93506f88415abbe9a76bead8d00cf112e4d4a86115306001607f1b88612b69565b61153a9190612b33565b95505b6f84102b00893f64c705e841d5d4064bd38610611590576115626001607a1b85612af6565b93506f84102b00893f64c705e841d5d4064bd36115836001607f1b88612b69565b61158d9190612b33565b95505b6f8204055aaef1c8bd5c3259f4822735a286106115e3576115b5600160791b85612af6565b93506f8204055aaef1c8bd5c3259f4822735a26115d66001607f1b88612b69565b6115e09190612b33565b95505b6f810100ab00222d861931c15e39b44e99861061163657611608600160781b85612af6565b93506f810100ab00222d861931c15e39b44e996116296001607f1b88612b69565b6116339190612b33565b95505b6f808040155aabbbe9451521693554f73386106116895761165b600160771b85612af6565b93506f808040155aabbbe9451521693554f73361167c6001607f1b88612b69565b6116869190612b33565b95505b6116976001607f1b87612b88565b92508291506001607f1b6116ab8380612b69565b6116b59190612b33565b9050600160801b6116c68482612b88565b6116d09084612b69565b6116da9190612b33565b6116e49085612af6565b93506001607f1b6116f58284612b69565b6116ff9190612b33565b9150600160811b611720846faaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa612b88565b61172a9084612b69565b6117349190612b33565b61173e9085612af6565b93506001607f1b61174f8284612b69565b6117599190612b33565b9150600360801b61177a846f99999999999999999999999999999999612b88565b6117849084612b69565b61178e9190612b33565b6117989085612af6565b93506001607f1b6117a98284612b69565b6117b39190612b33565b9150600160821b6117d4846f92492492492492492492492492492492612b88565b6117de9084612b69565b6117e89190612b33565b6117f29085612af6565b93506001607f1b6118038284612b69565b61180d9190612b33565b9150600560801b61182e846f8e38e38e38e38e38e38e38e38e38e38e612b88565b6118389084612b69565b6118429190612b33565b61184c9085612af6565b93506001607f1b61185d8284612b69565b6118679190612b33565b9150600360811b611888846f8ba2e8ba2e8ba2e8ba2e8ba2e8ba2e8b612b88565b6118929084612b69565b61189c9190612b33565b6118a69085612af6565b93506001607f1b6118b78284612b69565b6118c19190612b33565b9150600760801b6118e2846f89d89d89d89d89d89d89d89d89d89d89612b88565b6118ec9084612b69565b6118f69190612b33565b6119009085612af6565b93506001607f1b6119118284612b69565b61191b9190612b33565b9150600160831b61193c846f88888888888888888888888888888888612b88565b6119469084612b69565b6119509190612b33565b61195a9085612af6565b9695505050505050565b60008082600160801b81106119a957600061198b6119866001607f1b84612b33565b612889565b60ff811692831c929091506119a5906001607f1b90612b69565b9250505b6001607f1b811115611a1e57607f5b60ff811615611a1c576001607f1b6119d08380612b69565b6119da9190612b33565b9150600160801b8210611a0c57600191821c916119f79082612b9f565b60ff166001901b83611a099190612af6565b92505b611a1581612bc2565b90506119b8565b505b6f05b9de1d10bf4103d647b0955897ba80611a496f03f80fe03f80fe03f80fe03f80fe03f884612b69565b61077a9190612b33565b6000808080611a666001607c1b86612c1a565b91508190506001607f1b611a7a8280612b69565b611a849190612b33565b9050611a98816710e1b3be415a0000612b69565b611aa29084612af6565b92506001607f1b611ab38383612b69565b611abd9190612b33565b9050611ad1816705a0913f6b1e0000612b69565b611adb9084612af6565b92506001607f1b611aec8383612b69565b611af69190612b33565b9050611b0a81670168244fdac78000612b69565b611b149084612af6565b92506001607f1b611b258383612b69565b611b2f9190612b33565b9050611b4281664807432bc18000612b69565b611b4c9084612af6565b92506001607f1b611b5d8383612b69565b611b679190612b33565b9050611b7a81660c0135dca04000612b69565b611b849084612af6565b92506001607f1b611b958383612b69565b611b9f9190612b33565b9050611bb2816601b707b1cdc000612b69565b611bbc9084612af6565b92506001607f1b611bcd8383612b69565b611bd79190612b33565b9050611be9816536e0f639b800612b69565b611bf39084612af6565b92506001607f1b611c048383612b69565b611c0e9190612b33565b9050611c2081650618fee9f800612b69565b611c2a9084612af6565b92506001607f1b611c3b8383612b69565b611c459190612b33565b9050611c5681649c197dcc00612b69565b611c609084612af6565b92506001607f1b611c718383612b69565b611c7b9190612b33565b9050611c8c81640e30dce400612b69565b611c969084612af6565b92506001607f1b611ca78383612b69565b611cb19190612b33565b9050611cc28164012ebd1300612b69565b611ccc9084612af6565b92506001607f1b611cdd8383612b69565b611ce79190612b33565b9050611cf7816317499f00612b69565b611d019084612af6565b92506001607f1b611d128383612b69565b611d1c9190612b33565b9050611d2c816301a9d480612b69565b611d369084612af6565b92506001607f1b611d478383612b69565b611d519190612b33565b9050611d6081621c6380612b69565b611d6a9084612af6565b92506001607f1b611d7b8383612b69565b611d859190612b33565b9050611d94816201c638612b69565b611d9e9084612af6565b92506001607f1b611daf8383612b69565b611db99190612b33565b9050611dc781611ab8612b69565b611dd19084612af6565b92506001607f1b611de28383612b69565b611dec9190612b33565b9050611dfa8161017c612b69565b611e049084612af6565b92506001607f1b611e158383612b69565b611e1f9190612b33565b9050611e2c816014612b69565b611e369084612af6565b92506001607f1b611e478383612b69565b611e519190612b33565b9050611e5e816001612b69565b611e689084612af6565b92506001607f1b82611e826721c3677c82b4000086612b33565b611e8c9190612af6565b611e969190612af6565b92506001607c1b851615611ede5770018ebef9eac820ae8682b9793ac6d1e776611ed1847001c3d6a24ed82218787d624d3e5eba95f9612b69565b611edb9190612b33565b92505b6001607d1b851615611f24577001368b2fc6f9609fe7aceb46aa619baed4611f178470018ebef9eac820ae8682b9793ac6d1e778612b69565b611f219190612b33565b92505b6001607e1b851615611f69576fbc5ab1b16779be3575bd8f0520a9f21f611f5c847001368b2fc6f9609fe7aceb46aa619baed5612b69565b611f669190612b33565b92505b6001607f1b851615611fad576f454aaa8efe072e7f6ddbab84b40a55c9611fa0846fbc5ab1b16779be3575bd8f0520a9f21e612b69565b611faa9190612b33565b92505b600160801b851615611ff1576f0960aadc109e7a3bf4578099615711ea611fe4846f454aaa8efe072e7f6ddbab84b40a55c5612b69565b611fee9190612b33565b92505b600160811b851615612034576e2bf84208204f5977f9a8cf01fdce3d612027846f0960aadc109e7a3bf4578099615711d7612b69565b6120319190612b33565b92505b600160821b851615612075576d03c6ab775dd0b95b4cbee7e65d11612068846e2bf84208204f5977f9a8cf01fdc307612b69565b6120729190612b33565b92505b50909392505050565b60006020607f5b60ff8116612094836001612b0e565b60ff1610156120e757600060026120ab8385612b0e565b6120b59190612b47565b90508460008260ff16608081106120ce576120ce612c70565b0154106120dd578092506120e1565b8091505b50612085565b8360008260ff16608081106120fe576120fe612c70565b01541061210c579392505050565b8360008360ff166080811061212357612123612c70565b015410612131575092915050565b612139612c2e565b5060009392505050565b6000828160ff84166121558380612b69565b901c9150612173826f03442c4e6074a82f1797f72ac0000000612b69565b61217d9082612af6565b905060ff841661218d8684612b69565b901c91506121ab826f0116b96f757c380fb287fd0e40000000612b69565b6121b59082612af6565b905060ff84166121c58684612b69565b901c91506121e2826e45ae5bdd5f0e03eca1ff4390000000612b69565b6121ec9082612af6565b905060ff84166121fc8684612b69565b901c9150612219826e0defabf91302cd95b9ffda50000000612b69565b6122239082612af6565b905060ff84166122338684612b69565b901c9150612250826e02529ca9832b22439efff9b8000000612b69565b61225a9082612af6565b905060ff841661226a8684612b69565b901c9150612286826d54f1cf12bd04e516b6da88000000612b69565b6122909082612af6565b905060ff84166122a08684612b69565b901c91506122bc826d0a9e39e257a09ca2d6db51000000612b69565b6122c69082612af6565b905060ff84166122d68684612b69565b901c91506122f2826d012e066e7b839fa050c309000000612b69565b6122fc9082612af6565b905060ff841661230c8684612b69565b901c9150612327826c1e33d7d926c329a1ad1a800000612b69565b6123319082612af6565b905060ff84166123418684612b69565b901c915061235c826c02bee513bdb4a6b19b5f800000612b69565b6123669082612af6565b905060ff84166123768684612b69565b901c9150612390826b3a9316fa79b88eccf2a00000612b69565b61239a9082612af6565b905060ff84166123aa8684612b69565b901c91506123c4826b048177ebe1fa812375200000612b69565b6123ce9082612af6565b905060ff84166123de8684612b69565b901c91506123f7826a5263fe90242dcbacf00000612b69565b6124019082612af6565b905060ff84166124118684612b69565b901c915061242a826a057e22099c030d94100000612b69565b6124349082612af6565b905060ff84166124448684612b69565b901c915061245c826957e22099c030d9410000612b69565b6124669082612af6565b905060ff84166124768684612b69565b901c915061248e8269052b6b54569976310000612b69565b6124989082612af6565b905060ff84166124a88684612b69565b901c91506124bf82684985f67696bf748000612b69565b6124c99082612af6565b905060ff84166124d98684612b69565b901c91506124f0826803dea12ea99e498000612b69565b6124fa9082612af6565b905060ff841661250a8684612b69565b901c9150612520826731880f2214b6e000612b69565b61252a9082612af6565b905060ff841661253a8684612b69565b901c91506125508267025bcff56eb36000612b69565b61255a9082612af6565b905060ff841661256a8684612b69565b901c915061257f82661b722e10ab1000612b69565b6125899082612af6565b905060ff84166125998684612b69565b901c91506125ae826601317c70077000612b69565b6125b89082612af6565b905060ff84166125c88684612b69565b901c91506125dc82650cba84aafa00612b69565b6125e69082612af6565b905060ff84166125f68684612b69565b901c9150612609826482573a0a00612b69565b6126139082612af6565b905060ff84166126238684612b69565b901c9150612636826405035ad900612b69565b6126409082612af6565b905060ff84166126508684612b69565b901c915061266282632f881b00612b69565b61266c9082612af6565b905060ff841661267c8684612b69565b901c915061268e826301b29340612b69565b6126989082612af6565b905060ff84166126a88684612b69565b901c91506126b982620efc40612b69565b6126c39082612af6565b905060ff84166126d38684612b69565b901c91506126e382617fe0612b69565b6126ed9082612af6565b905060ff84166126fd8684612b69565b901c915061270d82610420612b69565b6127179082612af6565b905060ff84166127278684612b69565b901c9150612736826021612b69565b6127409082612af6565b905060ff84166127508684612b69565b901c915061275f826001612b69565b6127699082612af6565b9050600160ff85161b8561278d6f0688589cc0e9505e2f2fee558000000084612b33565b6127979190612af6565b6127a19190612af6565b95945050505050565b6001600160a01b0382166128005760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610588565b80608260008282546128129190612af6565b90915550506001600160a01b0382166000908152608060205260408120805483929061283f908490612af6565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b600080826101008110156128bd575b60018111156128b857600190811c906128b19083612b0e565b9150612898565b6128f1565b60805b60ff8116156128ef57600160ff82161b82106128e4579182179160ff81169190911c905b60011c607f166128c0565b505b5092915050565b80356001600160a01b038116811461290f57600080fd5b919050565b60006020828403121561292657600080fd5b61292f826128f8565b9392505050565b6000806040838503121561294957600080fd5b612952836128f8565b91506111af602084016128f8565b60008060006060848603121561297557600080fd5b61297e846128f8565b925061298c602085016128f8565b9150604084013590509250925092565b600080604083850312156129af57600080fd5b6129b8836128f8565b946020939093013593505050565b6000602082840312156129d857600080fd5b5035919050565b600080604083850312156129f257600080fd5b50508035926020909101359150565b60008060008060808587031215612a1757600080fd5b8435935060208501359250604085013563ffffffff81168114612a3957600080fd5b9396929550929360600135925050565b600060208284031215612a5b57600080fd5b813560ff8116811461292f57600080fd5b600060208083528351808285015260005b81811015612a9957858101830151858201604001528201612a7d565b81811115612aab576000604083870101525b50601f01601f1916929092016040019392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60008219821115612b0957612b09612c44565b500190565b600060ff821660ff84168060ff03821115612b2b57612b2b612c44565b019392505050565b600082612b4257612b42612c5a565b500490565b600060ff831680612b5a57612b5a612c5a565b8060ff84160491505092915050565b6000816000190483118215151615612b8357612b83612c44565b500290565b600082821015612b9a57612b9a612c44565b500390565b600060ff821660ff841680821015612bb957612bb9612c44565b90039392505050565b600060ff821680612bd557612bd5612c44565b6000190192915050565b600181811c90821680612bf357607f821691505b60208210811415612c1457634e487b7160e01b600052602260045260246000fd5b50919050565b600082612c2957612c29612c5a565b500690565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fdfea264697066735822122083649436cb6d0199cb1edd4444c207ef5c81b96547d4ca7d199d7aeec2fdc57864736f6c63430008060033"

// DeployCreatorToken deploys a new Ethereum contract, binding an instance of CreatorToken to it.
func DeployCreatorToken(auth *bind.TransactOpts, backend bind.ContractBackend, _initialSupply *big.Int, _reserveRatio *big.Int, name string, symbol string) (common.Address, *types.Transaction, *CreatorToken, error) {
	parsed, err := abi.JSON(strings.NewReader(CreatorTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CreatorTokenBin), backend, _initialSupply, _reserveRatio, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CreatorToken{CreatorTokenCaller: CreatorTokenCaller{contract: contract}, CreatorTokenTransactor: CreatorTokenTransactor{contract: contract}, CreatorTokenFilterer: CreatorTokenFilterer{contract: contract}}, nil
}

// CreatorToken is an auto generated Go binding around an Ethereum contract.
type CreatorToken struct {
	CreatorTokenCaller     // Read-only binding to the contract
	CreatorTokenTransactor // Write-only binding to the contract
	CreatorTokenFilterer   // Log filterer for contract events
}

// CreatorTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreatorTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreatorTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreatorTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreatorTokenSession struct {
	Contract     *CreatorToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreatorTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreatorTokenCallerSession struct {
	Contract *CreatorTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CreatorTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreatorTokenTransactorSession struct {
	Contract     *CreatorTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CreatorTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreatorTokenRaw struct {
	Contract *CreatorToken // Generic contract binding to access the raw methods on
}

// CreatorTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreatorTokenCallerRaw struct {
	Contract *CreatorTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CreatorTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreatorTokenTransactorRaw struct {
	Contract *CreatorTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreatorToken creates a new instance of CreatorToken, bound to a specific deployed contract.
func NewCreatorToken(address common.Address, backend bind.ContractBackend) (*CreatorToken, error) {
	contract, err := bindCreatorToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreatorToken{CreatorTokenCaller: CreatorTokenCaller{contract: contract}, CreatorTokenTransactor: CreatorTokenTransactor{contract: contract}, CreatorTokenFilterer: CreatorTokenFilterer{contract: contract}}, nil
}

// NewCreatorTokenCaller creates a new read-only instance of CreatorToken, bound to a specific deployed contract.
func NewCreatorTokenCaller(address common.Address, caller bind.ContractCaller) (*CreatorTokenCaller, error) {
	contract, err := bindCreatorToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorTokenCaller{contract: contract}, nil
}

// NewCreatorTokenTransactor creates a new write-only instance of CreatorToken, bound to a specific deployed contract.
func NewCreatorTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CreatorTokenTransactor, error) {
	contract, err := bindCreatorToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorTokenTransactor{contract: contract}, nil
}

// NewCreatorTokenFilterer creates a new log filterer instance of CreatorToken, bound to a specific deployed contract.
func NewCreatorTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CreatorTokenFilterer, error) {
	contract, err := bindCreatorToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreatorTokenFilterer{contract: contract}, nil
}

// bindCreatorToken binds a generic wrapper to an already deployed contract.
func bindCreatorToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreatorTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreatorToken *CreatorTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreatorToken.Contract.CreatorTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreatorToken *CreatorTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorToken.Contract.CreatorTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreatorToken *CreatorTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreatorToken.Contract.CreatorTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreatorToken *CreatorTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreatorToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreatorToken *CreatorTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreatorToken *CreatorTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreatorToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CreatorToken *CreatorTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CreatorToken *CreatorTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CreatorToken.Contract.Allowance(&_CreatorToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CreatorToken *CreatorTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CreatorToken.Contract.Allowance(&_CreatorToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CreatorToken *CreatorTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CreatorToken *CreatorTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CreatorToken.Contract.BalanceOf(&_CreatorToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CreatorToken *CreatorTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CreatorToken.Contract.BalanceOf(&_CreatorToken.CallOpts, account)
}

// CalculatePurchaseReturn is a free data retrieval call binding the contract method 0x29a00e7c.
//
// Solidity: function calculatePurchaseReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _depositAmount) view returns(uint256)
func (_CreatorToken *CreatorTokenCaller) CalculatePurchaseReturn(opts *bind.CallOpts, _supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _depositAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "calculatePurchaseReturn", _supply, _connectorBalance, _connectorWeight, _depositAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePurchaseReturn is a free data retrieval call binding the contract method 0x29a00e7c.
//
// Solidity: function calculatePurchaseReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _depositAmount) view returns(uint256)
func (_CreatorToken *CreatorTokenSession) CalculatePurchaseReturn(_supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _depositAmount *big.Int) (*big.Int, error) {
	return _CreatorToken.Contract.CalculatePurchaseReturn(&_CreatorToken.CallOpts, _supply, _connectorBalance, _connectorWeight, _depositAmount)
}

// CalculatePurchaseReturn is a free data retrieval call binding the contract method 0x29a00e7c.
//
// Solidity: function calculatePurchaseReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _depositAmount) view returns(uint256)
func (_CreatorToken *CreatorTokenCallerSession) CalculatePurchaseReturn(_supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _depositAmount *big.Int) (*big.Int, error) {
	return _CreatorToken.Contract.CalculatePurchaseReturn(&_CreatorToken.CallOpts, _supply, _connectorBalance, _connectorWeight, _depositAmount)
}

// CalculateSaleReturn is a free data retrieval call binding the contract method 0x49f9b0f7.
//
// Solidity: function calculateSaleReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _sellAmount) view returns(uint256)
func (_CreatorToken *CreatorTokenCaller) CalculateSaleReturn(opts *bind.CallOpts, _supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _sellAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "calculateSaleReturn", _supply, _connectorBalance, _connectorWeight, _sellAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSaleReturn is a free data retrieval call binding the contract method 0x49f9b0f7.
//
// Solidity: function calculateSaleReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _sellAmount) view returns(uint256)
func (_CreatorToken *CreatorTokenSession) CalculateSaleReturn(_supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _sellAmount *big.Int) (*big.Int, error) {
	return _CreatorToken.Contract.CalculateSaleReturn(&_CreatorToken.CallOpts, _supply, _connectorBalance, _connectorWeight, _sellAmount)
}

// CalculateSaleReturn is a free data retrieval call binding the contract method 0x49f9b0f7.
//
// Solidity: function calculateSaleReturn(uint256 _supply, uint256 _connectorBalance, uint32 _connectorWeight, uint256 _sellAmount) view returns(uint256)
func (_CreatorToken *CreatorTokenCallerSession) CalculateSaleReturn(_supply *big.Int, _connectorBalance *big.Int, _connectorWeight uint32, _sellAmount *big.Int) (*big.Int, error) {
	return _CreatorToken.Contract.CalculateSaleReturn(&_CreatorToken.CallOpts, _supply, _connectorBalance, _connectorWeight, _sellAmount)
}

// CalculateTokenBuyReturns is a free data retrieval call binding the contract method 0xa73d5647.
//
// Solidity: function calculateTokenBuyReturns(uint256 _amount) view returns(uint256, uint256)
func (_CreatorToken *CreatorTokenCaller) CalculateTokenBuyReturns(opts *bind.CallOpts, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "calculateTokenBuyReturns", _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CalculateTokenBuyReturns is a free data retrieval call binding the contract method 0xa73d5647.
//
// Solidity: function calculateTokenBuyReturns(uint256 _amount) view returns(uint256, uint256)
func (_CreatorToken *CreatorTokenSession) CalculateTokenBuyReturns(_amount *big.Int) (*big.Int, *big.Int, error) {
	return _CreatorToken.Contract.CalculateTokenBuyReturns(&_CreatorToken.CallOpts, _amount)
}

// CalculateTokenBuyReturns is a free data retrieval call binding the contract method 0xa73d5647.
//
// Solidity: function calculateTokenBuyReturns(uint256 _amount) view returns(uint256, uint256)
func (_CreatorToken *CreatorTokenCallerSession) CalculateTokenBuyReturns(_amount *big.Int) (*big.Int, *big.Int, error) {
	return _CreatorToken.Contract.CalculateTokenBuyReturns(&_CreatorToken.CallOpts, _amount)
}

// CheckAndReturnInitialContractBalance is a free data retrieval call binding the contract method 0xff50e196.
//
// Solidity: function checkAndReturnInitialContractBalance(uint256 _amountToDeposit, uint256 balanceToCheckAgainst) view returns(uint256 _amountToUseAsTokenBalance)
func (_CreatorToken *CreatorTokenCaller) CheckAndReturnInitialContractBalance(opts *bind.CallOpts, _amountToDeposit *big.Int, balanceToCheckAgainst *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "checkAndReturnInitialContractBalance", _amountToDeposit, balanceToCheckAgainst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAndReturnInitialContractBalance is a free data retrieval call binding the contract method 0xff50e196.
//
// Solidity: function checkAndReturnInitialContractBalance(uint256 _amountToDeposit, uint256 balanceToCheckAgainst) view returns(uint256 _amountToUseAsTokenBalance)
func (_CreatorToken *CreatorTokenSession) CheckAndReturnInitialContractBalance(_amountToDeposit *big.Int, balanceToCheckAgainst *big.Int) (*big.Int, error) {
	return _CreatorToken.Contract.CheckAndReturnInitialContractBalance(&_CreatorToken.CallOpts, _amountToDeposit, balanceToCheckAgainst)
}

// CheckAndReturnInitialContractBalance is a free data retrieval call binding the contract method 0xff50e196.
//
// Solidity: function checkAndReturnInitialContractBalance(uint256 _amountToDeposit, uint256 balanceToCheckAgainst) view returns(uint256 _amountToUseAsTokenBalance)
func (_CreatorToken *CreatorTokenCallerSession) CheckAndReturnInitialContractBalance(_amountToDeposit *big.Int, balanceToCheckAgainst *big.Int) (*big.Int, error) {
	return _CreatorToken.Contract.CheckAndReturnInitialContractBalance(&_CreatorToken.CallOpts, _amountToDeposit, balanceToCheckAgainst)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CreatorToken *CreatorTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CreatorToken *CreatorTokenSession) Decimals() (uint8, error) {
	return _CreatorToken.Contract.Decimals(&_CreatorToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CreatorToken *CreatorTokenCallerSession) Decimals() (uint8, error) {
	return _CreatorToken.Contract.Decimals(&_CreatorToken.CallOpts)
}

// FounderPercentage is a free data retrieval call binding the contract method 0x17cf6914.
//
// Solidity: function founderPercentage() view returns(uint8)
func (_CreatorToken *CreatorTokenCaller) FounderPercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "founderPercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FounderPercentage is a free data retrieval call binding the contract method 0x17cf6914.
//
// Solidity: function founderPercentage() view returns(uint8)
func (_CreatorToken *CreatorTokenSession) FounderPercentage() (uint8, error) {
	return _CreatorToken.Contract.FounderPercentage(&_CreatorToken.CallOpts)
}

// FounderPercentage is a free data retrieval call binding the contract method 0x17cf6914.
//
// Solidity: function founderPercentage() view returns(uint8)
func (_CreatorToken *CreatorTokenCallerSession) FounderPercentage() (uint8, error) {
	return _CreatorToken.Contract.FounderPercentage(&_CreatorToken.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_CreatorToken *CreatorTokenCaller) GetEthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "getEthBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_CreatorToken *CreatorTokenSession) GetEthBalance() (*big.Int, error) {
	return _CreatorToken.Contract.GetEthBalance(&_CreatorToken.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_CreatorToken *CreatorTokenCallerSession) GetEthBalance() (*big.Int, error) {
	return _CreatorToken.Contract.GetEthBalance(&_CreatorToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorToken *CreatorTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorToken *CreatorTokenSession) Name() (string, error) {
	return _CreatorToken.Contract.Name(&_CreatorToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreatorToken *CreatorTokenCallerSession) Name() (string, error) {
	return _CreatorToken.Contract.Name(&_CreatorToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorToken *CreatorTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorToken *CreatorTokenSession) Owner() (common.Address, error) {
	return _CreatorToken.Contract.Owner(&_CreatorToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CreatorToken *CreatorTokenCallerSession) Owner() (common.Address, error) {
	return _CreatorToken.Contract.Owner(&_CreatorToken.CallOpts)
}

// ReserveBalance is a free data retrieval call binding the contract method 0xa10954fe.
//
// Solidity: function reserveBalance() view returns(uint256)
func (_CreatorToken *CreatorTokenCaller) ReserveBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "reserveBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveBalance is a free data retrieval call binding the contract method 0xa10954fe.
//
// Solidity: function reserveBalance() view returns(uint256)
func (_CreatorToken *CreatorTokenSession) ReserveBalance() (*big.Int, error) {
	return _CreatorToken.Contract.ReserveBalance(&_CreatorToken.CallOpts)
}

// ReserveBalance is a free data retrieval call binding the contract method 0xa10954fe.
//
// Solidity: function reserveBalance() view returns(uint256)
func (_CreatorToken *CreatorTokenCallerSession) ReserveBalance() (*big.Int, error) {
	return _CreatorToken.Contract.ReserveBalance(&_CreatorToken.CallOpts)
}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint256)
func (_CreatorToken *CreatorTokenCaller) ReserveRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "reserveRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint256)
func (_CreatorToken *CreatorTokenSession) ReserveRatio() (*big.Int, error) {
	return _CreatorToken.Contract.ReserveRatio(&_CreatorToken.CallOpts)
}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint256)
func (_CreatorToken *CreatorTokenCallerSession) ReserveRatio() (*big.Int, error) {
	return _CreatorToken.Contract.ReserveRatio(&_CreatorToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorToken *CreatorTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorToken *CreatorTokenSession) Symbol() (string, error) {
	return _CreatorToken.Contract.Symbol(&_CreatorToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CreatorToken *CreatorTokenCallerSession) Symbol() (string, error) {
	return _CreatorToken.Contract.Symbol(&_CreatorToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CreatorToken *CreatorTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreatorToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CreatorToken *CreatorTokenSession) TotalSupply() (*big.Int, error) {
	return _CreatorToken.Contract.TotalSupply(&_CreatorToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CreatorToken *CreatorTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CreatorToken.Contract.TotalSupply(&_CreatorToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.Approve(&_CreatorToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.Approve(&_CreatorToken.TransactOpts, spender, amount)
}

// BuyNewTokens is a paid mutator transaction binding the contract method 0xb593217c.
//
// Solidity: function buyNewTokens() payable returns()
func (_CreatorToken *CreatorTokenTransactor) BuyNewTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "buyNewTokens")
}

// BuyNewTokens is a paid mutator transaction binding the contract method 0xb593217c.
//
// Solidity: function buyNewTokens() payable returns()
func (_CreatorToken *CreatorTokenSession) BuyNewTokens() (*types.Transaction, error) {
	return _CreatorToken.Contract.BuyNewTokens(&_CreatorToken.TransactOpts)
}

// BuyNewTokens is a paid mutator transaction binding the contract method 0xb593217c.
//
// Solidity: function buyNewTokens() payable returns()
func (_CreatorToken *CreatorTokenTransactorSession) BuyNewTokens() (*types.Transaction, error) {
	return _CreatorToken.Contract.BuyNewTokens(&_CreatorToken.TransactOpts)
}

// ChangeFounderPercentage is a paid mutator transaction binding the contract method 0x0cc4e3a7.
//
// Solidity: function changeFounderPercentage(uint8 _newPercentage) returns()
func (_CreatorToken *CreatorTokenTransactor) ChangeFounderPercentage(opts *bind.TransactOpts, _newPercentage uint8) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "changeFounderPercentage", _newPercentage)
}

// ChangeFounderPercentage is a paid mutator transaction binding the contract method 0x0cc4e3a7.
//
// Solidity: function changeFounderPercentage(uint8 _newPercentage) returns()
func (_CreatorToken *CreatorTokenSession) ChangeFounderPercentage(_newPercentage uint8) (*types.Transaction, error) {
	return _CreatorToken.Contract.ChangeFounderPercentage(&_CreatorToken.TransactOpts, _newPercentage)
}

// ChangeFounderPercentage is a paid mutator transaction binding the contract method 0x0cc4e3a7.
//
// Solidity: function changeFounderPercentage(uint8 _newPercentage) returns()
func (_CreatorToken *CreatorTokenTransactorSession) ChangeFounderPercentage(_newPercentage uint8) (*types.Transaction, error) {
	return _CreatorToken.Contract.ChangeFounderPercentage(&_CreatorToken.TransactOpts, _newPercentage)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CreatorToken *CreatorTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CreatorToken *CreatorTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.DecreaseAllowance(&_CreatorToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CreatorToken *CreatorTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.DecreaseAllowance(&_CreatorToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CreatorToken *CreatorTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CreatorToken *CreatorTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.IncreaseAllowance(&_CreatorToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CreatorToken *CreatorTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.IncreaseAllowance(&_CreatorToken.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorToken *CreatorTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorToken *CreatorTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _CreatorToken.Contract.RenounceOwnership(&_CreatorToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CreatorToken *CreatorTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CreatorToken.Contract.RenounceOwnership(&_CreatorToken.TransactOpts)
}

// SellTokensForEth is a paid mutator transaction binding the contract method 0xb3e7a384.
//
// Solidity: function sellTokensForEth(uint256 _amount) returns()
func (_CreatorToken *CreatorTokenTransactor) SellTokensForEth(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "sellTokensForEth", _amount)
}

// SellTokensForEth is a paid mutator transaction binding the contract method 0xb3e7a384.
//
// Solidity: function sellTokensForEth(uint256 _amount) returns()
func (_CreatorToken *CreatorTokenSession) SellTokensForEth(_amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.SellTokensForEth(&_CreatorToken.TransactOpts, _amount)
}

// SellTokensForEth is a paid mutator transaction binding the contract method 0xb3e7a384.
//
// Solidity: function sellTokensForEth(uint256 _amount) returns()
func (_CreatorToken *CreatorTokenTransactorSession) SellTokensForEth(_amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.SellTokensForEth(&_CreatorToken.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.Transfer(&_CreatorToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.Transfer(&_CreatorToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.TransferFrom(&_CreatorToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CreatorToken *CreatorTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreatorToken.Contract.TransferFrom(&_CreatorToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorToken *CreatorTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CreatorToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorToken *CreatorTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreatorToken.Contract.TransferOwnership(&_CreatorToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CreatorToken *CreatorTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CreatorToken.Contract.TransferOwnership(&_CreatorToken.TransactOpts, newOwner)
}

// CreatorTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CreatorToken contract.
type CreatorTokenApprovalIterator struct {
	Event *CreatorTokenApproval // Event containing the contract specifics and raw log

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
func (it *CreatorTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorTokenApproval)
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
		it.Event = new(CreatorTokenApproval)
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
func (it *CreatorTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorTokenApproval represents a Approval event raised by the CreatorToken contract.
type CreatorTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CreatorToken *CreatorTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CreatorTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CreatorToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CreatorTokenApprovalIterator{contract: _CreatorToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CreatorToken *CreatorTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CreatorTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CreatorToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorTokenApproval)
				if err := _CreatorToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CreatorToken *CreatorTokenFilterer) ParseApproval(log types.Log) (*CreatorTokenApproval, error) {
	event := new(CreatorTokenApproval)
	if err := _CreatorToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorTokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the CreatorToken contract.
type CreatorTokenBurnedIterator struct {
	Event *CreatorTokenBurned // Event containing the contract specifics and raw log

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
func (it *CreatorTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorTokenBurned)
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
		it.Event = new(CreatorTokenBurned)
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
func (it *CreatorTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorTokenBurned represents a Burned event raised by the CreatorToken contract.
type CreatorTokenBurned struct {
	To              common.Address
	AmountMinted    *big.Int
	AmountDeposited *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x23ff0e75edf108e3d0392d92e13e8c8a868ef19001bd49f9e94876dc46dff87f.
//
// Solidity: event Burned(address _To, uint256 _amountMinted, uint256 _amountDeposited)
func (_CreatorToken *CreatorTokenFilterer) FilterBurned(opts *bind.FilterOpts) (*CreatorTokenBurnedIterator, error) {

	logs, sub, err := _CreatorToken.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &CreatorTokenBurnedIterator{contract: _CreatorToken.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x23ff0e75edf108e3d0392d92e13e8c8a868ef19001bd49f9e94876dc46dff87f.
//
// Solidity: event Burned(address _To, uint256 _amountMinted, uint256 _amountDeposited)
func (_CreatorToken *CreatorTokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *CreatorTokenBurned) (event.Subscription, error) {

	logs, sub, err := _CreatorToken.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorTokenBurned)
				if err := _CreatorToken.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x23ff0e75edf108e3d0392d92e13e8c8a868ef19001bd49f9e94876dc46dff87f.
//
// Solidity: event Burned(address _To, uint256 _amountMinted, uint256 _amountDeposited)
func (_CreatorToken *CreatorTokenFilterer) ParseBurned(log types.Log) (*CreatorTokenBurned, error) {
	event := new(CreatorTokenBurned)
	if err := _CreatorToken.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorTokenMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the CreatorToken contract.
type CreatorTokenMintedIterator struct {
	Event *CreatorTokenMinted // Event containing the contract specifics and raw log

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
func (it *CreatorTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorTokenMinted)
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
		it.Event = new(CreatorTokenMinted)
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
func (it *CreatorTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorTokenMinted represents a Minted event raised by the CreatorToken contract.
type CreatorTokenMinted struct {
	Buyer                    common.Address
	Deposited                *big.Int
	AmountMinted             *big.Int
	AmountDistributedToBuyer *big.Int
	AmountDistributedToOwner *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x19f5f791ee407773427bf7b970bbbc3375065c32edd1ab142e23a84f94b0719b.
//
// Solidity: event Minted(address buyer, uint256 deposited, uint256 amountMinted, uint256 amountDistributedToBuyer, uint256 amountDistributedToOwner)
func (_CreatorToken *CreatorTokenFilterer) FilterMinted(opts *bind.FilterOpts) (*CreatorTokenMintedIterator, error) {

	logs, sub, err := _CreatorToken.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &CreatorTokenMintedIterator{contract: _CreatorToken.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x19f5f791ee407773427bf7b970bbbc3375065c32edd1ab142e23a84f94b0719b.
//
// Solidity: event Minted(address buyer, uint256 deposited, uint256 amountMinted, uint256 amountDistributedToBuyer, uint256 amountDistributedToOwner)
func (_CreatorToken *CreatorTokenFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *CreatorTokenMinted) (event.Subscription, error) {

	logs, sub, err := _CreatorToken.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorTokenMinted)
				if err := _CreatorToken.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x19f5f791ee407773427bf7b970bbbc3375065c32edd1ab142e23a84f94b0719b.
//
// Solidity: event Minted(address buyer, uint256 deposited, uint256 amountMinted, uint256 amountDistributedToBuyer, uint256 amountDistributedToOwner)
func (_CreatorToken *CreatorTokenFilterer) ParseMinted(log types.Log) (*CreatorTokenMinted, error) {
	event := new(CreatorTokenMinted)
	if err := _CreatorToken.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CreatorToken contract.
type CreatorTokenOwnershipTransferredIterator struct {
	Event *CreatorTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CreatorTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorTokenOwnershipTransferred)
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
		it.Event = new(CreatorTokenOwnershipTransferred)
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
func (it *CreatorTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorTokenOwnershipTransferred represents a OwnershipTransferred event raised by the CreatorToken contract.
type CreatorTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreatorToken *CreatorTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CreatorTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreatorToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreatorTokenOwnershipTransferredIterator{contract: _CreatorToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreatorToken *CreatorTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CreatorTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CreatorToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorTokenOwnershipTransferred)
				if err := _CreatorToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CreatorToken *CreatorTokenFilterer) ParseOwnershipTransferred(log types.Log) (*CreatorTokenOwnershipTransferred, error) {
	event := new(CreatorTokenOwnershipTransferred)
	if err := _CreatorToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatorTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CreatorToken contract.
type CreatorTokenTransferIterator struct {
	Event *CreatorTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CreatorTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatorTokenTransfer)
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
		it.Event = new(CreatorTokenTransfer)
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
func (it *CreatorTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatorTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatorTokenTransfer represents a Transfer event raised by the CreatorToken contract.
type CreatorTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CreatorToken *CreatorTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CreatorTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreatorToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreatorTokenTransferIterator{contract: _CreatorToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CreatorToken *CreatorTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CreatorTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CreatorToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatorTokenTransfer)
				if err := _CreatorToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CreatorToken *CreatorTokenFilterer) ParseTransfer(log types.Log) (*CreatorTokenTransfer, error) {
	event := new(CreatorTokenTransfer)
	if err := _CreatorToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x60806040523480156200001157600080fd5b5060405162000b5638038062000b568339810160408190526200003491620001c5565b81516200004990600390602085019062000068565b5080516200005f90600490602084019062000068565b50505062000282565b82805462000076906200022f565b90600052602060002090601f0160209004810192826200009a5760008555620000e5565b82601f10620000b557805160ff1916838001178555620000e5565b82800160010185558215620000e5579182015b82811115620000e5578251825591602001919060010190620000c8565b50620000f3929150620000f7565b5090565b5b80821115620000f35760008155600101620000f8565b600082601f8301126200012057600080fd5b81516001600160401b03808211156200013d576200013d6200026c565b604051601f8301601f19908116603f011681019082821181831017156200016857620001686200026c565b816040528381526020925086838588010111156200018557600080fd5b600091505b83821015620001a957858201830151818301840152908201906200018a565b83821115620001bb5760008385830101525b9695505050505050565b60008060408385031215620001d957600080fd5b82516001600160401b0380821115620001f157600080fd5b620001ff868387016200010e565b935060208501519150808211156200021657600080fd5b5062000225858286016200010e565b9150509250929050565b600181811c908216806200024457607f821691505b602082108114156200026657634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6108c480620002926000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b411461015f578063a457c2d714610167578063a9059cbb1461017a578063dd62ed3e1461018d57600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101c6565b6040516100c391906107d8565b60405180910390f35b6100df6100da3660046107ae565b610258565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f366004610772565b61026e565b604051601281526020016100c3565b6100df6101313660046107ae565b61031d565b6100f361014436600461071d565b6001600160a01b031660009081526020819052604090205490565b6100b6610359565b6100df6101753660046107ae565b610368565b6100df6101883660046107ae565b610401565b6100f361019b36600461073f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101d590610853565b80601f016020809104026020016040519081016040528092919081815260200182805461020190610853565b801561024e5780601f106102235761010080835404028352916020019161024e565b820191906000526020600020905b81548152906001019060200180831161023157829003601f168201915b5050505050905090565b600061026533848461040e565b50600192915050565b600061027b848484610532565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156103055760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b610312853385840361040e565b506001949350505050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161026591859061035490869061082d565b61040e565b6060600480546101d590610853565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156103ea5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016102fc565b6103f7338585840361040e565b5060019392505050565b6000610265338484610532565b6001600160a01b0383166104705760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016102fc565b6001600160a01b0382166104d15760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016102fc565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383166105965760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016102fc565b6001600160a01b0382166105f85760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016102fc565b6001600160a01b038316600090815260208190526040902054818110156106705760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016102fc565b6001600160a01b038085166000908152602081905260408082208585039055918516815290812080548492906106a790849061082d565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106f391815260200190565b60405180910390a350505050565b80356001600160a01b038116811461071857600080fd5b919050565b60006020828403121561072f57600080fd5b61073882610701565b9392505050565b6000806040838503121561075257600080fd5b61075b83610701565b915061076960208401610701565b90509250929050565b60008060006060848603121561078757600080fd5b61079084610701565b925061079e60208501610701565b9150604084013590509250925092565b600080604083850312156107c157600080fd5b6107ca83610701565b946020939093013593505050565b600060208083528351808285015260005b81811015610805578581018301518582016040015282016107e9565b81811115610817576000604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561084e57634e487b7160e01b600052601160045260246000fd5b500190565b600181811c9082168061086757607f821691505b6020821081141561088857634e487b7160e01b600052602260045260246000fd5b5091905056fea2646970667358221220794cb484502131fc5219af293d27b91f91f35c000661646e4b266544b7aa817064736f6c63430008060033"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataABI is the input ABI used to generate the binding from.
const IERC20MetadataABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20MetadataFuncSigs maps the 4-byte function signature to its string representation.
var IERC20MetadataFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	IERC20MetadataCaller     // Read-only binding to the contract
	IERC20MetadataTransactor // Write-only binding to the contract
	IERC20MetadataFilterer   // Log filterer for contract events
}

// IERC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataSession struct {
	Contract     *IERC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataCallerSession struct {
	Contract *IERC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataTransactorSession struct {
	Contract     *IERC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataRaw struct {
	Contract *IERC20Metadata // Generic contract binding to access the raw methods on
}

// IERC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataCallerRaw struct {
	Contract *IERC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactorRaw struct {
	Contract *IERC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Metadata creates a new instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20Metadata(address common.Address, backend bind.ContractBackend) (*IERC20Metadata, error) {
	contract, err := bindIERC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Metadata{IERC20MetadataCaller: IERC20MetadataCaller{contract: contract}, IERC20MetadataTransactor: IERC20MetadataTransactor{contract: contract}, IERC20MetadataFilterer: IERC20MetadataFilterer{contract: contract}}, nil
}

// NewIERC20MetadataCaller creates a new read-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataCaller, error) {
	contract, err := bindIERC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataCaller{contract: contract}, nil
}

// NewIERC20MetadataTransactor creates a new write-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataTransactor, error) {
	contract, err := bindIERC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransactor{contract: contract}, nil
}

// NewIERC20MetadataFilterer creates a new log filterer instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataFilterer, error) {
	contract, err := bindIERC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataFilterer{contract: contract}, nil
}

// bindIERC20Metadata binds a generic wrapper to an already deployed contract.
func bindIERC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.IERC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// IERC20MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Metadata contract.
type IERC20MetadataApprovalIterator struct {
	Event *IERC20MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataApproval)
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
		it.Event = new(IERC20MetadataApproval)
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
func (it *IERC20MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataApprovalIterator{contract: _IERC20Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataApproval)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseApproval(log types.Log) (*IERC20MetadataApproval, error) {
	event := new(IERC20MetadataApproval)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Metadata contract.
type IERC20MetadataTransferIterator struct {
	Event *IERC20MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataTransfer)
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
		it.Event = new(IERC20MetadataTransfer)
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
func (it *IERC20MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransferIterator{contract: _IERC20Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataTransfer)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseTransfer(log types.Log) (*IERC20MetadataTransfer, error) {
	event := new(IERC20MetadataTransfer)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PowerABI is the input ABI used to generate the binding from.
const PowerABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PowerBin is the compiled bytecode used for deploying new contracts.
var PowerBin = "0x608060405234801561001057600080fd5b506001641c35fedd1560601b036020556001646c3390ecc9605e1b036021556001640cf801476160611b0360225560016431bdb23e1d605f1b0360235560016502fb1d8fe083605b1b0360245560016505b771955b37605a1b036025556001650af67a93bb5160591b0360265560016515060c256cb360581b036027556001651428a2f98d7360581b036028556001654d515663970960561b03602955600165944620b0e70f60551b03602a55600166011c592761c66760541b03602b5560016602214d10d014eb60531b03602c55600166020ade36b7dbef60531b03602d5560016603eab73b3bbfe360521b03602e556001660782ee3593f6d760511b03602f556001661ccf4b44bb4821604f1b0360305560016606e7f88ad8a77760511b0360315560016669f3d1c921891d604d1b03603255600166cb2ff529eb71e5604c1b03603355600166c2d415c3db974b604c1b0360345560016702eb40f9f620fda7604a1b0360355560016705990681d961a1eb60491b03603655600167055e12902701414760491b0360375560016714962dee9dc9764160471b0360385560016704ef57b9b560fab560491b0360395560016712ed7b32a58f552b60471b03603a556001679131271922eaa60760441b03603b556001678b380f3558668c4760441b03603c556001680215f77c045fbe885760421b03603d556001600160831b03603e556001670f577eded5773a1160471b03603f556001680eb5ec597592befbf5603f1b036040556001681c35fedd14b861eb05603e1b036041556001683619c87664579bc94b603d1b0360425560016867c00a3b07ffc01fd7603c1b03604355600168c6f6c8f8739773a7a5603b1b03604455600168bec763f8209b7a72b1603b1b0360455560016902dbb8caad9b7097b91b60391b03604655600169057b3d49dda84556d6f760381b03604755600169054183095b2c8ececf3160381b036048556001690a14517cc6b9457111ef60371b0360495560016913545598e5c23276ccf160361b03604a556001692511882c39c3adea96ff60351b03604b55600169471649d87199aa99075760341b03604c557004429a21a029d4c1457cfbffffffffffff604d55700415bc6d6fb7dd71af2cb3ffffffffffff604e557003eab73b3bbfe282243ce1ffffffffffff604f557003c1771ac9fb6b4c18e229ffffffffffff605055700399e96897690418f785257fffffffffff605155700373fc456c53bb779bf0ea9fffffffffff60525570034f9e8e490c48e67e6ab8bfffffffffff60535570032cbfd4a7adc790560b3337ffffffffff60545570030b50570f6e5d2acca94613ffffffffff6055557002eb40f9f620fda6b56c2861ffffffffff6056557002cc8340ecb0d0f520a6af58ffffffffff6057557002af09481380a0a35cf1ba02ffffffffff605855700292c5bdd3b92ec810287b1b3fffffffff605955700277abdcdab07d5a77ac6d6b9fffffffff605a5570025daf6654b1eaa55fd64df5efffffffff605b55700244c49c648baa98192dce88b7ffffffff605c5570022ce03cd5619a311b2471268bffffffff605d55700215f77c045fbe885654a44a0fffffffff605e556001600160811b03605f557001eaefdbdaaee7421fc4d3ede5ffffffff6060557001d6bd8b2eb257df7e8ca57b09bfffffff6061557001c35fedd14b861eb0443f7f133fffffff6062557001b0ce43b322bcde4a56e8ada5afffffff60635570019f0028ec1fff007f5a195a39dfffffff60645570018ded91f0e72ee74f49b15ba527ffffff60655570017d8ec7f04136f4e5615fd41a63ffffff60665570016ddc6556cdb84bdc8d12d22e6fffffff60675570015ecf52776a1155b5bd8395814f7fffff60685570015060c256cb23b3b3cc3754cf40ffffff6069557001428a2f98d728ae223ddab715be3fffff606a5570013545598e5c23276ccf0ede68034fffff606b557001288c4161ce1d6f54b7f61081194fffff606c5570011c592761c666aa641d5a01a40f17ffff606d55700110a688680a7530515f3e6e6cfdcdffff606e557001056f1b5bedf75c6bcb2ce8aed428ffff606f556ffaadceceeff8a0890f3875f008277fff6070556ff05dc6b27edad306388a600f6ba0bfff6071556fe67a5a25da41063de1495d5b18cdbfff6072556fdcff115b14eedde6fc3aa5353f2e4fff6073556fd3e7a3924312399f9aae2e0f868f8fff6074556fcb2ff529eb71e41582cccd5a1ee26fff6075556fc2d415c3db974ab32a51840c0b67edff6076556fbad03e7d883f69ad5b0a186184e06bff6077556fb320d03b2c343d4829abd6075f0cc5ff6078556fabc25204e02828d73c6e80bcdb1a95bf6079556fa4b16f74ee4bb2040a1ec6c15fbbf2df607a556f9deaf736ac1f569deb1b5ae3f36c130f607b556f976bd9952c7aa957f5937d790ef65037607c556f9131271922eaa6064b73a22d0bd4f2bf607d556f8b380f3558668c46c91c49a2f8e967b9607e556f857ddf0117efa215952912839f6473e6607f55603f806107486000396000f3fe6080604052600080fdfea2646970667358221220e6fb00fbf520777156520085a5db77428b9183231001e841a7fb24fe9803bb5964736f6c63430008060033"

// DeployPower deploys a new Ethereum contract, binding an instance of Power to it.
func DeployPower(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Power, error) {
	parsed, err := abi.JSON(strings.NewReader(PowerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PowerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Power{PowerCaller: PowerCaller{contract: contract}, PowerTransactor: PowerTransactor{contract: contract}, PowerFilterer: PowerFilterer{contract: contract}}, nil
}

// Power is an auto generated Go binding around an Ethereum contract.
type Power struct {
	PowerCaller     // Read-only binding to the contract
	PowerTransactor // Write-only binding to the contract
	PowerFilterer   // Log filterer for contract events
}

// PowerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PowerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PowerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PowerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PowerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PowerSession struct {
	Contract     *Power            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PowerCallerSession struct {
	Contract *PowerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PowerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PowerTransactorSession struct {
	Contract     *PowerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PowerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PowerRaw struct {
	Contract *Power // Generic contract binding to access the raw methods on
}

// PowerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PowerCallerRaw struct {
	Contract *PowerCaller // Generic read-only contract binding to access the raw methods on
}

// PowerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PowerTransactorRaw struct {
	Contract *PowerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPower creates a new instance of Power, bound to a specific deployed contract.
func NewPower(address common.Address, backend bind.ContractBackend) (*Power, error) {
	contract, err := bindPower(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Power{PowerCaller: PowerCaller{contract: contract}, PowerTransactor: PowerTransactor{contract: contract}, PowerFilterer: PowerFilterer{contract: contract}}, nil
}

// NewPowerCaller creates a new read-only instance of Power, bound to a specific deployed contract.
func NewPowerCaller(address common.Address, caller bind.ContractCaller) (*PowerCaller, error) {
	contract, err := bindPower(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PowerCaller{contract: contract}, nil
}

// NewPowerTransactor creates a new write-only instance of Power, bound to a specific deployed contract.
func NewPowerTransactor(address common.Address, transactor bind.ContractTransactor) (*PowerTransactor, error) {
	contract, err := bindPower(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PowerTransactor{contract: contract}, nil
}

// NewPowerFilterer creates a new log filterer instance of Power, bound to a specific deployed contract.
func NewPowerFilterer(address common.Address, filterer bind.ContractFilterer) (*PowerFilterer, error) {
	contract, err := bindPower(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PowerFilterer{contract: contract}, nil
}

// bindPower binds a generic wrapper to an already deployed contract.
func bindPower(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PowerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Power *PowerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Power.Contract.PowerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Power *PowerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Power.Contract.PowerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Power *PowerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Power.Contract.PowerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Power *PowerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Power.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Power *PowerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Power.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Power *PowerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Power.Contract.contract.Transact(opts, method, params...)
}
