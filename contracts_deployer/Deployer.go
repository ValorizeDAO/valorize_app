// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deployer

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

// DeployerContractDeployParameters is an auto generated low-level Go binding around an user-defined struct.
type DeployerContractDeployParameters struct {
	ByteCodeHash [32]byte
	Price        *big.Int
}

// DeployerDeployedContractInfo is an auto generated low-level Go binding around an user-defined struct.
type DeployerDeployedContractInfo struct {
	DeploymentAddress common.Address
	ContractType      string
}

// DeployerMetaData contains all meta data concerning the Deployer contract.
var DeployerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"byteCodeHash\",\"type\":\"bytes32\"}],\"name\":\"ByteCodeUploaded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contractType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"ContractDiscontinued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"bytecode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"deployContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractKey\",\"type\":\"string\"}],\"name\":\"discontinueContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractKey\",\"type\":\"string\"}],\"name\":\"getContractByteCodeHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"byteCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structDeployer.ContractDeployParameters\",\"name\":\"contractParams\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"getDeployed\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"deploymentAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractType\",\"type\":\"string\"}],\"internalType\":\"structDeployer.DeployedContractInfo[]\",\"name\":\"contractsDeployed\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"byteCode\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"contractDeployPrice\",\"type\":\"uint256\"}],\"name\":\"setContractByteCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updateContractPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use DeployerMetaData.ABI instead.
var DeployerABI = DeployerMetaData.ABI

// Deployer is an auto generated Go binding around an Ethereum contract.
type Deployer struct {
	DeployerCaller     // Read-only binding to the contract
	DeployerTransactor // Write-only binding to the contract
	DeployerFilterer   // Log filterer for contract events
}

// DeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployerSession struct {
	Contract     *Deployer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployerCallerSession struct {
	Contract *DeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployerTransactorSession struct {
	Contract     *DeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployerRaw struct {
	Contract *Deployer // Generic contract binding to access the raw methods on
}

// DeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployerCallerRaw struct {
	Contract *DeployerCaller // Generic read-only contract binding to access the raw methods on
}

// DeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployerTransactorRaw struct {
	Contract *DeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployer creates a new instance of Deployer, bound to a specific deployed contract.
func NewDeployer(address common.Address, backend bind.ContractBackend) (*Deployer, error) {
	contract, err := bindDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deployer{DeployerCaller: DeployerCaller{contract: contract}, DeployerTransactor: DeployerTransactor{contract: contract}, DeployerFilterer: DeployerFilterer{contract: contract}}, nil
}

// NewDeployerCaller creates a new read-only instance of Deployer, bound to a specific deployed contract.
func NewDeployerCaller(address common.Address, caller bind.ContractCaller) (*DeployerCaller, error) {
	contract, err := bindDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployerCaller{contract: contract}, nil
}

// NewDeployerTransactor creates a new write-only instance of Deployer, bound to a specific deployed contract.
func NewDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployerTransactor, error) {
	contract, err := bindDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployerTransactor{contract: contract}, nil
}

// NewDeployerFilterer creates a new log filterer instance of Deployer, bound to a specific deployed contract.
func NewDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployerFilterer, error) {
	contract, err := bindDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployerFilterer{contract: contract}, nil
}

// bindDeployer binds a generic wrapper to an already deployed contract.
func bindDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployer *DeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployer.Contract.DeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployer *DeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.Contract.DeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployer *DeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployer.Contract.DeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployer *DeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployer *DeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployer *DeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployer.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deployer *DeployerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deployer *DeployerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Deployer.Contract.DEFAULTADMINROLE(&_Deployer.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deployer *DeployerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Deployer.Contract.DEFAULTADMINROLE(&_Deployer.CallOpts)
}

// GetContractByteCodeHash is a free data retrieval call binding the contract method 0xcc188d75.
//
// Solidity: function getContractByteCodeHash(string contractKey) view returns(bool success, (bytes32,uint256) contractParams)
func (_Deployer *DeployerCaller) GetContractByteCodeHash(opts *bind.CallOpts, contractKey string) (struct {
	Success        bool
	ContractParams DeployerContractDeployParameters
}, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "getContractByteCodeHash", contractKey)

	outstruct := new(struct {
		Success        bool
		ContractParams DeployerContractDeployParameters
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ContractParams = *abi.ConvertType(out[1], new(DeployerContractDeployParameters)).(*DeployerContractDeployParameters)

	return *outstruct, err

}

// GetContractByteCodeHash is a free data retrieval call binding the contract method 0xcc188d75.
//
// Solidity: function getContractByteCodeHash(string contractKey) view returns(bool success, (bytes32,uint256) contractParams)
func (_Deployer *DeployerSession) GetContractByteCodeHash(contractKey string) (struct {
	Success        bool
	ContractParams DeployerContractDeployParameters
}, error) {
	return _Deployer.Contract.GetContractByteCodeHash(&_Deployer.CallOpts, contractKey)
}

// GetContractByteCodeHash is a free data retrieval call binding the contract method 0xcc188d75.
//
// Solidity: function getContractByteCodeHash(string contractKey) view returns(bool success, (bytes32,uint256) contractParams)
func (_Deployer *DeployerCallerSession) GetContractByteCodeHash(contractKey string) (struct {
	Success        bool
	ContractParams DeployerContractDeployParameters
}, error) {
	return _Deployer.Contract.GetContractByteCodeHash(&_Deployer.CallOpts, contractKey)
}

// GetDeployed is a free data retrieval call binding the contract method 0x43f03ab6.
//
// Solidity: function getDeployed(address deployer) view returns((address,string)[] contractsDeployed)
func (_Deployer *DeployerCaller) GetDeployed(opts *bind.CallOpts, deployer common.Address) ([]DeployerDeployedContractInfo, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "getDeployed", deployer)

	if err != nil {
		return *new([]DeployerDeployedContractInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DeployerDeployedContractInfo)).(*[]DeployerDeployedContractInfo)

	return out0, err

}

// GetDeployed is a free data retrieval call binding the contract method 0x43f03ab6.
//
// Solidity: function getDeployed(address deployer) view returns((address,string)[] contractsDeployed)
func (_Deployer *DeployerSession) GetDeployed(deployer common.Address) ([]DeployerDeployedContractInfo, error) {
	return _Deployer.Contract.GetDeployed(&_Deployer.CallOpts, deployer)
}

// GetDeployed is a free data retrieval call binding the contract method 0x43f03ab6.
//
// Solidity: function getDeployed(address deployer) view returns((address,string)[] contractsDeployed)
func (_Deployer *DeployerCallerSession) GetDeployed(deployer common.Address) ([]DeployerDeployedContractInfo, error) {
	return _Deployer.Contract.GetDeployed(&_Deployer.CallOpts, deployer)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deployer *DeployerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deployer *DeployerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Deployer.Contract.GetRoleAdmin(&_Deployer.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deployer *DeployerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Deployer.Contract.GetRoleAdmin(&_Deployer.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deployer *DeployerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deployer *DeployerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Deployer.Contract.HasRole(&_Deployer.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deployer *DeployerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Deployer.Contract.HasRole(&_Deployer.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deployer *DeployerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Deployer.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deployer *DeployerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Deployer.Contract.SupportsInterface(&_Deployer.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deployer *DeployerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Deployer.Contract.SupportsInterface(&_Deployer.CallOpts, interfaceId)
}

// DeployContract is a paid mutator transaction binding the contract method 0x169b102a.
//
// Solidity: function deployContract(string contractType, bytes bytecode, bytes params, bytes32 salt) payable returns()
func (_Deployer *DeployerTransactor) DeployContract(opts *bind.TransactOpts, contractType string, bytecode []byte, params []byte, salt [32]byte) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "deployContract", contractType, bytecode, params, salt)
}

// DeployContract is a paid mutator transaction binding the contract method 0x169b102a.
//
// Solidity: function deployContract(string contractType, bytes bytecode, bytes params, bytes32 salt) payable returns()
func (_Deployer *DeployerSession) DeployContract(contractType string, bytecode []byte, params []byte, salt [32]byte) (*types.Transaction, error) {
	return _Deployer.Contract.DeployContract(&_Deployer.TransactOpts, contractType, bytecode, params, salt)
}

// DeployContract is a paid mutator transaction binding the contract method 0x169b102a.
//
// Solidity: function deployContract(string contractType, bytes bytecode, bytes params, bytes32 salt) payable returns()
func (_Deployer *DeployerTransactorSession) DeployContract(contractType string, bytecode []byte, params []byte, salt [32]byte) (*types.Transaction, error) {
	return _Deployer.Contract.DeployContract(&_Deployer.TransactOpts, contractType, bytecode, params, salt)
}

// DiscontinueContract is a paid mutator transaction binding the contract method 0x4f519f57.
//
// Solidity: function discontinueContract(string contractKey) returns()
func (_Deployer *DeployerTransactor) DiscontinueContract(opts *bind.TransactOpts, contractKey string) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "discontinueContract", contractKey)
}

// DiscontinueContract is a paid mutator transaction binding the contract method 0x4f519f57.
//
// Solidity: function discontinueContract(string contractKey) returns()
func (_Deployer *DeployerSession) DiscontinueContract(contractKey string) (*types.Transaction, error) {
	return _Deployer.Contract.DiscontinueContract(&_Deployer.TransactOpts, contractKey)
}

// DiscontinueContract is a paid mutator transaction binding the contract method 0x4f519f57.
//
// Solidity: function discontinueContract(string contractKey) returns()
func (_Deployer *DeployerTransactorSession) DiscontinueContract(contractKey string) (*types.Transaction, error) {
	return _Deployer.Contract.DiscontinueContract(&_Deployer.TransactOpts, contractKey)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deployer *DeployerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deployer *DeployerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.GrantRole(&_Deployer.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deployer *DeployerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.GrantRole(&_Deployer.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Deployer *DeployerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Deployer *DeployerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.RenounceRole(&_Deployer.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Deployer *DeployerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.RenounceRole(&_Deployer.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deployer *DeployerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deployer *DeployerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.RevokeRole(&_Deployer.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deployer *DeployerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deployer.Contract.RevokeRole(&_Deployer.TransactOpts, role, account)
}

// SetContractByteCode is a paid mutator transaction binding the contract method 0x3fd213a4.
//
// Solidity: function setContractByteCode(string contractKey, bytes byteCode, uint256 contractDeployPrice) returns()
func (_Deployer *DeployerTransactor) SetContractByteCode(opts *bind.TransactOpts, contractKey string, byteCode []byte, contractDeployPrice *big.Int) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "setContractByteCode", contractKey, byteCode, contractDeployPrice)
}

// SetContractByteCode is a paid mutator transaction binding the contract method 0x3fd213a4.
//
// Solidity: function setContractByteCode(string contractKey, bytes byteCode, uint256 contractDeployPrice) returns()
func (_Deployer *DeployerSession) SetContractByteCode(contractKey string, byteCode []byte, contractDeployPrice *big.Int) (*types.Transaction, error) {
	return _Deployer.Contract.SetContractByteCode(&_Deployer.TransactOpts, contractKey, byteCode, contractDeployPrice)
}

// SetContractByteCode is a paid mutator transaction binding the contract method 0x3fd213a4.
//
// Solidity: function setContractByteCode(string contractKey, bytes byteCode, uint256 contractDeployPrice) returns()
func (_Deployer *DeployerTransactorSession) SetContractByteCode(contractKey string, byteCode []byte, contractDeployPrice *big.Int) (*types.Transaction, error) {
	return _Deployer.Contract.SetContractByteCode(&_Deployer.TransactOpts, contractKey, byteCode, contractDeployPrice)
}

// UpdateContractPrice is a paid mutator transaction binding the contract method 0xcae00fb6.
//
// Solidity: function updateContractPrice(string contractKey, uint256 newPrice) returns()
func (_Deployer *DeployerTransactor) UpdateContractPrice(opts *bind.TransactOpts, contractKey string, newPrice *big.Int) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "updateContractPrice", contractKey, newPrice)
}

// UpdateContractPrice is a paid mutator transaction binding the contract method 0xcae00fb6.
//
// Solidity: function updateContractPrice(string contractKey, uint256 newPrice) returns()
func (_Deployer *DeployerSession) UpdateContractPrice(contractKey string, newPrice *big.Int) (*types.Transaction, error) {
	return _Deployer.Contract.UpdateContractPrice(&_Deployer.TransactOpts, contractKey, newPrice)
}

// UpdateContractPrice is a paid mutator transaction binding the contract method 0xcae00fb6.
//
// Solidity: function updateContractPrice(string contractKey, uint256 newPrice) returns()
func (_Deployer *DeployerTransactorSession) UpdateContractPrice(contractKey string, newPrice *big.Int) (*types.Transaction, error) {
	return _Deployer.Contract.UpdateContractPrice(&_Deployer.TransactOpts, contractKey, newPrice)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Deployer *DeployerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Deployer *DeployerSession) Withdraw() (*types.Transaction, error) {
	return _Deployer.Contract.Withdraw(&_Deployer.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Deployer *DeployerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Deployer.Contract.Withdraw(&_Deployer.TransactOpts)
}

// DeployerByteCodeUploadedIterator is returned from FilterByteCodeUploaded and is used to iterate over the raw logs and unpacked data for ByteCodeUploaded events raised by the Deployer contract.
type DeployerByteCodeUploadedIterator struct {
	Event *DeployerByteCodeUploaded // Event containing the contract specifics and raw log

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
func (it *DeployerByteCodeUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployerByteCodeUploaded)
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
		it.Event = new(DeployerByteCodeUploaded)
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
func (it *DeployerByteCodeUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployerByteCodeUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployerByteCodeUploaded represents a ByteCodeUploaded event raised by the Deployer contract.
type DeployerByteCodeUploaded struct {
	Key          string
	Price        *big.Int
	ByteCodeHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterByteCodeUploaded is a free log retrieval operation binding the contract event 0x42eab6542519210ad29b79f0355a9f6e664d4b220592651a45a36d6c203dbf22.
//
// Solidity: event ByteCodeUploaded(string key, uint256 price, bytes32 byteCodeHash)
func (_Deployer *DeployerFilterer) FilterByteCodeUploaded(opts *bind.FilterOpts) (*DeployerByteCodeUploadedIterator, error) {

	logs, sub, err := _Deployer.contract.FilterLogs(opts, "ByteCodeUploaded")
	if err != nil {
		return nil, err
	}
	return &DeployerByteCodeUploadedIterator{contract: _Deployer.contract, event: "ByteCodeUploaded", logs: logs, sub: sub}, nil
}

// WatchByteCodeUploaded is a free log subscription operation binding the contract event 0x42eab6542519210ad29b79f0355a9f6e664d4b220592651a45a36d6c203dbf22.
//
// Solidity: event ByteCodeUploaded(string key, uint256 price, bytes32 byteCodeHash)
func (_Deployer *DeployerFilterer) WatchByteCodeUploaded(opts *bind.WatchOpts, sink chan<- *DeployerByteCodeUploaded) (event.Subscription, error) {

	logs, sub, err := _Deployer.contract.WatchLogs(opts, "ByteCodeUploaded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployerByteCodeUploaded)
				if err := _Deployer.contract.UnpackLog(event, "ByteCodeUploaded", log); err != nil {
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

// ParseByteCodeUploaded is a log parse operation binding the contract event 0x42eab6542519210ad29b79f0355a9f6e664d4b220592651a45a36d6c203dbf22.
//
// Solidity: event ByteCodeUploaded(string key, uint256 price, bytes32 byteCodeHash)
func (_Deployer *DeployerFilterer) ParseByteCodeUploaded(log types.Log) (*DeployerByteCodeUploaded, error) {
	event := new(DeployerByteCodeUploaded)
	if err := _Deployer.contract.UnpackLog(event, "ByteCodeUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeployerContractDeployedIterator is returned from FilterContractDeployed and is used to iterate over the raw logs and unpacked data for ContractDeployed events raised by the Deployer contract.
type DeployerContractDeployedIterator struct {
	Event *DeployerContractDeployed // Event containing the contract specifics and raw log

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
func (it *DeployerContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployerContractDeployed)
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
		it.Event = new(DeployerContractDeployed)
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
func (it *DeployerContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployerContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployerContractDeployed represents a ContractDeployed event raised by the Deployer contract.
type DeployerContractDeployed struct {
	ContractAddress common.Address
	ContractType    string
	Paid            *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractDeployed is a free log retrieval operation binding the contract event 0x1a8fbd6fef18ec2c02e373a114addb92628002c1ce8b9df1f5b177b3eb6e29a6.
//
// Solidity: event ContractDeployed(address contractAddress, string contractType, uint256 paid)
func (_Deployer *DeployerFilterer) FilterContractDeployed(opts *bind.FilterOpts) (*DeployerContractDeployedIterator, error) {

	logs, sub, err := _Deployer.contract.FilterLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return &DeployerContractDeployedIterator{contract: _Deployer.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

// WatchContractDeployed is a free log subscription operation binding the contract event 0x1a8fbd6fef18ec2c02e373a114addb92628002c1ce8b9df1f5b177b3eb6e29a6.
//
// Solidity: event ContractDeployed(address contractAddress, string contractType, uint256 paid)
func (_Deployer *DeployerFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *DeployerContractDeployed) (event.Subscription, error) {

	logs, sub, err := _Deployer.contract.WatchLogs(opts, "ContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployerContractDeployed)
				if err := _Deployer.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
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

// ParseContractDeployed is a log parse operation binding the contract event 0x1a8fbd6fef18ec2c02e373a114addb92628002c1ce8b9df1f5b177b3eb6e29a6.
//
// Solidity: event ContractDeployed(address contractAddress, string contractType, uint256 paid)
func (_Deployer *DeployerFilterer) ParseContractDeployed(log types.Log) (*DeployerContractDeployed, error) {
	event := new(DeployerContractDeployed)
	if err := _Deployer.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeployerContractDiscontinuedIterator is returned from FilterContractDiscontinued and is used to iterate over the raw logs and unpacked data for ContractDiscontinued events raised by the Deployer contract.
type DeployerContractDiscontinuedIterator struct {
	Event *DeployerContractDiscontinued // Event containing the contract specifics and raw log

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
func (it *DeployerContractDiscontinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployerContractDiscontinued)
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
		it.Event = new(DeployerContractDiscontinued)
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
func (it *DeployerContractDiscontinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployerContractDiscontinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployerContractDiscontinued represents a ContractDiscontinued event raised by the Deployer contract.
type DeployerContractDiscontinued struct {
	Key string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractDiscontinued is a free log retrieval operation binding the contract event 0xb2a0d3dbddeb41393a8a3665602a228ff0505b37f7f8185d1bb6666ddcbdfbff.
//
// Solidity: event ContractDiscontinued(string key)
func (_Deployer *DeployerFilterer) FilterContractDiscontinued(opts *bind.FilterOpts) (*DeployerContractDiscontinuedIterator, error) {

	logs, sub, err := _Deployer.contract.FilterLogs(opts, "ContractDiscontinued")
	if err != nil {
		return nil, err
	}
	return &DeployerContractDiscontinuedIterator{contract: _Deployer.contract, event: "ContractDiscontinued", logs: logs, sub: sub}, nil
}

// WatchContractDiscontinued is a free log subscription operation binding the contract event 0xb2a0d3dbddeb41393a8a3665602a228ff0505b37f7f8185d1bb6666ddcbdfbff.
//
// Solidity: event ContractDiscontinued(string key)
func (_Deployer *DeployerFilterer) WatchContractDiscontinued(opts *bind.WatchOpts, sink chan<- *DeployerContractDiscontinued) (event.Subscription, error) {

	logs, sub, err := _Deployer.contract.WatchLogs(opts, "ContractDiscontinued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployerContractDiscontinued)
				if err := _Deployer.contract.UnpackLog(event, "ContractDiscontinued", log); err != nil {
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

// ParseContractDiscontinued is a log parse operation binding the contract event 0xb2a0d3dbddeb41393a8a3665602a228ff0505b37f7f8185d1bb6666ddcbdfbff.
//
// Solidity: event ContractDiscontinued(string key)
func (_Deployer *DeployerFilterer) ParseContractDiscontinued(log types.Log) (*DeployerContractDiscontinued, error) {
	event := new(DeployerContractDiscontinued)
	if err := _Deployer.contract.UnpackLog(event, "ContractDiscontinued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeployerPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the Deployer contract.
type DeployerPriceUpdatedIterator struct {
	Event *DeployerPriceUpdated // Event containing the contract specifics and raw log

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
func (it *DeployerPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployerPriceUpdated)
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
		it.Event = new(DeployerPriceUpdated)
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
func (it *DeployerPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployerPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployerPriceUpdated represents a PriceUpdated event raised by the Deployer contract.
type DeployerPriceUpdated struct {
	Key      string
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string key, uint256 newPrice)
func (_Deployer *DeployerFilterer) FilterPriceUpdated(opts *bind.FilterOpts) (*DeployerPriceUpdatedIterator, error) {

	logs, sub, err := _Deployer.contract.FilterLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return &DeployerPriceUpdatedIterator{contract: _Deployer.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string key, uint256 newPrice)
func (_Deployer *DeployerFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *DeployerPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _Deployer.contract.WatchLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployerPriceUpdated)
				if err := _Deployer.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0x159e83f4712ba2552e68be9d848e49bf6dd35c24f19564ffd523b6549450a2f4.
//
// Solidity: event PriceUpdated(string key, uint256 newPrice)
func (_Deployer *DeployerFilterer) ParsePriceUpdated(log types.Log) (*DeployerPriceUpdated, error) {
	event := new(DeployerPriceUpdated)
	if err := _Deployer.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeployerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Deployer contract.
type DeployerRoleAdminChangedIterator struct {
	Event *DeployerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DeployerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployerRoleAdminChanged)
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
		it.Event = new(DeployerRoleAdminChanged)
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
func (it *DeployerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployerRoleAdminChanged represents a RoleAdminChanged event raised by the Deployer contract.
type DeployerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Deployer *DeployerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DeployerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Deployer.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DeployerRoleAdminChangedIterator{contract: _Deployer.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Deployer *DeployerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DeployerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Deployer.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployerRoleAdminChanged)
				if err := _Deployer.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Deployer *DeployerFilterer) ParseRoleAdminChanged(log types.Log) (*DeployerRoleAdminChanged, error) {
	event := new(DeployerRoleAdminChanged)
	if err := _Deployer.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeployerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Deployer contract.
type DeployerRoleGrantedIterator struct {
	Event *DeployerRoleGranted // Event containing the contract specifics and raw log

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
func (it *DeployerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployerRoleGranted)
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
		it.Event = new(DeployerRoleGranted)
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
func (it *DeployerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployerRoleGranted represents a RoleGranted event raised by the Deployer contract.
type DeployerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deployer *DeployerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DeployerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Deployer.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DeployerRoleGrantedIterator{contract: _Deployer.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deployer *DeployerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DeployerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Deployer.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployerRoleGranted)
				if err := _Deployer.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deployer *DeployerFilterer) ParseRoleGranted(log types.Log) (*DeployerRoleGranted, error) {
	event := new(DeployerRoleGranted)
	if err := _Deployer.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DeployerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Deployer contract.
type DeployerRoleRevokedIterator struct {
	Event *DeployerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DeployerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeployerRoleRevoked)
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
		it.Event = new(DeployerRoleRevoked)
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
func (it *DeployerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeployerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeployerRoleRevoked represents a RoleRevoked event raised by the Deployer contract.
type DeployerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deployer *DeployerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DeployerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Deployer.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DeployerRoleRevokedIterator{contract: _Deployer.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deployer *DeployerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DeployerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Deployer.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeployerRoleRevoked)
				if err := _Deployer.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deployer *DeployerFilterer) ParseRoleRevoked(log types.Log) (*DeployerRoleRevoked, error) {
	event := new(DeployerRoleRevoked)
	if err := _Deployer.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
