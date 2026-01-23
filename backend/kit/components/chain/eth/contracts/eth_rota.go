// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ETHRotaMetaData contains all meta data concerning the ETHRota contract.
var ETHRotaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"contractIHasher\",\"name\":\"_hasher\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_denomination\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_merkleTreeHeight\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nullifierHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIELD_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"ROOT_HISTORY_SIZE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"currentRootIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"denomination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"filledSubtrees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getLastRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIHasher\",\"name\":\"_hasher\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"hasher\",\"outputs\":[{\"internalType\":\"contractIHasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"isKnownRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nullifierHash\",\"type\":\"bytes32\"}],\"name\":\"isSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_nullifierHashes\",\"type\":\"bytes32[]\"}],\"name\":\"isSpentArray\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"spent\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"nextIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nullifierHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_nullifierHash\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]",
}

// ETHRotaABI is the input ABI used to generate the binding from.
// Deprecated: Use ETHRotaMetaData.ABI instead.
var ETHRotaABI = ETHRotaMetaData.ABI

// ETHRota is an auto generated Go binding around an Ethereum contract.
type ETHRota struct {
	ETHRotaCaller     // Read-only binding to the contract
	ETHRotaTransactor // Write-only binding to the contract
	ETHRotaFilterer   // Log filterer for contract events
}

// ETHRotaCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHRotaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHRotaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHRotaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHRotaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHRotaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHRotaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHRotaSession struct {
	Contract     *ETHRota          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHRotaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHRotaCallerSession struct {
	Contract *ETHRotaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ETHRotaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHRotaTransactorSession struct {
	Contract     *ETHRotaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ETHRotaRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHRotaRaw struct {
	Contract *ETHRota // Generic contract binding to access the raw methods on
}

// ETHRotaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHRotaCallerRaw struct {
	Contract *ETHRotaCaller // Generic read-only contract binding to access the raw methods on
}

// ETHRotaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHRotaTransactorRaw struct {
	Contract *ETHRotaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETHRota creates a new instance of ETHRota, bound to a specific deployed contract.
func NewETHRota(address common.Address, backend bind.ContractBackend) (*ETHRota, error) {
	contract, err := bindETHRota(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHRota{ETHRotaCaller: ETHRotaCaller{contract: contract}, ETHRotaTransactor: ETHRotaTransactor{contract: contract}, ETHRotaFilterer: ETHRotaFilterer{contract: contract}}, nil
}

// NewETHRotaCaller creates a new read-only instance of ETHRota, bound to a specific deployed contract.
func NewETHRotaCaller(address common.Address, caller bind.ContractCaller) (*ETHRotaCaller, error) {
	contract, err := bindETHRota(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHRotaCaller{contract: contract}, nil
}

// NewETHRotaTransactor creates a new write-only instance of ETHRota, bound to a specific deployed contract.
func NewETHRotaTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHRotaTransactor, error) {
	contract, err := bindETHRota(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHRotaTransactor{contract: contract}, nil
}

// NewETHRotaFilterer creates a new log filterer instance of ETHRota, bound to a specific deployed contract.
func NewETHRotaFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHRotaFilterer, error) {
	contract, err := bindETHRota(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHRotaFilterer{contract: contract}, nil
}

// bindETHRota binds a generic wrapper to an already deployed contract.
func bindETHRota(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ETHRotaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHRota *ETHRotaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHRota.Contract.ETHRotaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHRota *ETHRotaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHRota.Contract.ETHRotaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHRota *ETHRotaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHRota.Contract.ETHRotaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHRota *ETHRotaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHRota.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHRota *ETHRotaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHRota.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHRota *ETHRotaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHRota.Contract.contract.Transact(opts, method, params...)
}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_ETHRota *ETHRotaCaller) FIELDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "FIELD_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_ETHRota *ETHRotaSession) FIELDSIZE() (*big.Int, error) {
	return _ETHRota.Contract.FIELDSIZE(&_ETHRota.CallOpts)
}

// FIELDSIZE is a free data retrieval call binding the contract method 0x414a37ba.
//
// Solidity: function FIELD_SIZE() view returns(uint256)
func (_ETHRota *ETHRotaCallerSession) FIELDSIZE() (*big.Int, error) {
	return _ETHRota.Contract.FIELDSIZE(&_ETHRota.CallOpts)
}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_ETHRota *ETHRotaCaller) ROOTHISTORYSIZE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "ROOT_HISTORY_SIZE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_ETHRota *ETHRotaSession) ROOTHISTORYSIZE() (uint32, error) {
	return _ETHRota.Contract.ROOTHISTORYSIZE(&_ETHRota.CallOpts)
}

// ROOTHISTORYSIZE is a free data retrieval call binding the contract method 0xcd87a3b4.
//
// Solidity: function ROOT_HISTORY_SIZE() view returns(uint32)
func (_ETHRota *ETHRotaCallerSession) ROOTHISTORYSIZE() (uint32, error) {
	return _ETHRota.Contract.ROOTHISTORYSIZE(&_ETHRota.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_ETHRota *ETHRotaCaller) ZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_ETHRota *ETHRotaSession) ZEROVALUE() (*big.Int, error) {
	return _ETHRota.Contract.ZEROVALUE(&_ETHRota.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_ETHRota *ETHRotaCallerSession) ZEROVALUE() (*big.Int, error) {
	return _ETHRota.Contract.ZEROVALUE(&_ETHRota.CallOpts)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_ETHRota *ETHRotaCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_ETHRota *ETHRotaSession) Commitments(arg0 [32]byte) (bool, error) {
	return _ETHRota.Contract.Commitments(&_ETHRota.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_ETHRota *ETHRotaCallerSession) Commitments(arg0 [32]byte) (bool, error) {
	return _ETHRota.Contract.Commitments(&_ETHRota.CallOpts, arg0)
}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_ETHRota *ETHRotaCaller) CurrentRootIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "currentRootIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_ETHRota *ETHRotaSession) CurrentRootIndex() (uint32, error) {
	return _ETHRota.Contract.CurrentRootIndex(&_ETHRota.CallOpts)
}

// CurrentRootIndex is a free data retrieval call binding the contract method 0x90eeb02b.
//
// Solidity: function currentRootIndex() view returns(uint32)
func (_ETHRota *ETHRotaCallerSession) CurrentRootIndex() (uint32, error) {
	return _ETHRota.Contract.CurrentRootIndex(&_ETHRota.CallOpts)
}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_ETHRota *ETHRotaCaller) Denomination(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "denomination")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_ETHRota *ETHRotaSession) Denomination() (*big.Int, error) {
	return _ETHRota.Contract.Denomination(&_ETHRota.CallOpts)
}

// Denomination is a free data retrieval call binding the contract method 0x8bca6d16.
//
// Solidity: function denomination() view returns(uint256)
func (_ETHRota *ETHRotaCallerSession) Denomination() (*big.Int, error) {
	return _ETHRota.Contract.Denomination(&_ETHRota.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ETHRota *ETHRotaCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ETHRota *ETHRotaSession) Fee() (*big.Int, error) {
	return _ETHRota.Contract.Fee(&_ETHRota.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_ETHRota *ETHRotaCallerSession) Fee() (*big.Int, error) {
	return _ETHRota.Contract.Fee(&_ETHRota.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ETHRota *ETHRotaCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ETHRota *ETHRotaSession) FeeTo() (common.Address, error) {
	return _ETHRota.Contract.FeeTo(&_ETHRota.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_ETHRota *ETHRotaCallerSession) FeeTo() (common.Address, error) {
	return _ETHRota.Contract.FeeTo(&_ETHRota.CallOpts)
}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_ETHRota *ETHRotaCaller) FilledSubtrees(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "filledSubtrees", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_ETHRota *ETHRotaSession) FilledSubtrees(arg0 *big.Int) ([32]byte, error) {
	return _ETHRota.Contract.FilledSubtrees(&_ETHRota.CallOpts, arg0)
}

// FilledSubtrees is a free data retrieval call binding the contract method 0xf178e47c.
//
// Solidity: function filledSubtrees(uint256 ) view returns(bytes32)
func (_ETHRota *ETHRotaCallerSession) FilledSubtrees(arg0 *big.Int) ([32]byte, error) {
	return _ETHRota.Contract.FilledSubtrees(&_ETHRota.CallOpts, arg0)
}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_ETHRota *ETHRotaCaller) GetLastRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "getLastRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_ETHRota *ETHRotaSession) GetLastRoot() ([32]byte, error) {
	return _ETHRota.Contract.GetLastRoot(&_ETHRota.CallOpts)
}

// GetLastRoot is a free data retrieval call binding the contract method 0xba70f757.
//
// Solidity: function getLastRoot() view returns(bytes32)
func (_ETHRota *ETHRotaCallerSession) GetLastRoot() ([32]byte, error) {
	return _ETHRota.Contract.GetLastRoot(&_ETHRota.CallOpts)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ETHRota *ETHRotaCaller) HashLeftRight(opts *bind.CallOpts, _hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "hashLeftRight", _hasher, _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ETHRota *ETHRotaSession) HashLeftRight(_hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	return _ETHRota.Contract.HashLeftRight(&_ETHRota.CallOpts, _hasher, _left, _right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x8ea3099e.
//
// Solidity: function hashLeftRight(address _hasher, bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_ETHRota *ETHRotaCallerSession) HashLeftRight(_hasher common.Address, _left [32]byte, _right [32]byte) ([32]byte, error) {
	return _ETHRota.Contract.HashLeftRight(&_ETHRota.CallOpts, _hasher, _left, _right)
}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_ETHRota *ETHRotaCaller) Hasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "hasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_ETHRota *ETHRotaSession) Hasher() (common.Address, error) {
	return _ETHRota.Contract.Hasher(&_ETHRota.CallOpts)
}

// Hasher is a free data retrieval call binding the contract method 0xed33639f.
//
// Solidity: function hasher() view returns(address)
func (_ETHRota *ETHRotaCallerSession) Hasher() (common.Address, error) {
	return _ETHRota.Contract.Hasher(&_ETHRota.CallOpts)
}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_ETHRota *ETHRotaCaller) IsKnownRoot(opts *bind.CallOpts, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "isKnownRoot", _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_ETHRota *ETHRotaSession) IsKnownRoot(_root [32]byte) (bool, error) {
	return _ETHRota.Contract.IsKnownRoot(&_ETHRota.CallOpts, _root)
}

// IsKnownRoot is a free data retrieval call binding the contract method 0x6d9833e3.
//
// Solidity: function isKnownRoot(bytes32 _root) view returns(bool)
func (_ETHRota *ETHRotaCallerSession) IsKnownRoot(_root [32]byte) (bool, error) {
	return _ETHRota.Contract.IsKnownRoot(&_ETHRota.CallOpts, _root)
}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_ETHRota *ETHRotaCaller) IsSpent(opts *bind.CallOpts, _nullifierHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "isSpent", _nullifierHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_ETHRota *ETHRotaSession) IsSpent(_nullifierHash [32]byte) (bool, error) {
	return _ETHRota.Contract.IsSpent(&_ETHRota.CallOpts, _nullifierHash)
}

// IsSpent is a free data retrieval call binding the contract method 0xe5285dcc.
//
// Solidity: function isSpent(bytes32 _nullifierHash) view returns(bool)
func (_ETHRota *ETHRotaCallerSession) IsSpent(_nullifierHash [32]byte) (bool, error) {
	return _ETHRota.Contract.IsSpent(&_ETHRota.CallOpts, _nullifierHash)
}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_ETHRota *ETHRotaCaller) IsSpentArray(opts *bind.CallOpts, _nullifierHashes [][32]byte) ([]bool, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "isSpentArray", _nullifierHashes)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_ETHRota *ETHRotaSession) IsSpentArray(_nullifierHashes [][32]byte) ([]bool, error) {
	return _ETHRota.Contract.IsSpentArray(&_ETHRota.CallOpts, _nullifierHashes)
}

// IsSpentArray is a free data retrieval call binding the contract method 0x9fa12d0b.
//
// Solidity: function isSpentArray(bytes32[] _nullifierHashes) view returns(bool[] spent)
func (_ETHRota *ETHRotaCallerSession) IsSpentArray(_nullifierHashes [][32]byte) ([]bool, error) {
	return _ETHRota.Contract.IsSpentArray(&_ETHRota.CallOpts, _nullifierHashes)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_ETHRota *ETHRotaCaller) Levels(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "levels")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_ETHRota *ETHRotaSession) Levels() (uint32, error) {
	return _ETHRota.Contract.Levels(&_ETHRota.CallOpts)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint32)
func (_ETHRota *ETHRotaCallerSession) Levels() (uint32, error) {
	return _ETHRota.Contract.Levels(&_ETHRota.CallOpts)
}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_ETHRota *ETHRotaCaller) NextIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "nextIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_ETHRota *ETHRotaSession) NextIndex() (uint32, error) {
	return _ETHRota.Contract.NextIndex(&_ETHRota.CallOpts)
}

// NextIndex is a free data retrieval call binding the contract method 0xfc7e9c6f.
//
// Solidity: function nextIndex() view returns(uint32)
func (_ETHRota *ETHRotaCallerSession) NextIndex() (uint32, error) {
	return _ETHRota.Contract.NextIndex(&_ETHRota.CallOpts)
}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_ETHRota *ETHRotaCaller) NullifierHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "nullifierHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_ETHRota *ETHRotaSession) NullifierHashes(arg0 [32]byte) (bool, error) {
	return _ETHRota.Contract.NullifierHashes(&_ETHRota.CallOpts, arg0)
}

// NullifierHashes is a free data retrieval call binding the contract method 0x17cc915c.
//
// Solidity: function nullifierHashes(bytes32 ) view returns(bool)
func (_ETHRota *ETHRotaCallerSession) NullifierHashes(arg0 [32]byte) (bool, error) {
	return _ETHRota.Contract.NullifierHashes(&_ETHRota.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ETHRota *ETHRotaCaller) Roots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ETHRota *ETHRotaSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _ETHRota.Contract.Roots(&_ETHRota.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ETHRota *ETHRotaCallerSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _ETHRota.Contract.Roots(&_ETHRota.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ETHRota *ETHRotaCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ETHRota *ETHRotaSession) Verifier() (common.Address, error) {
	return _ETHRota.Contract.Verifier(&_ETHRota.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ETHRota *ETHRotaCallerSession) Verifier() (common.Address, error) {
	return _ETHRota.Contract.Verifier(&_ETHRota.CallOpts)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_ETHRota *ETHRotaCaller) Zeros(opts *bind.CallOpts, i *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ETHRota.contract.Call(opts, &out, "zeros", i)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_ETHRota *ETHRotaSession) Zeros(i *big.Int) ([32]byte, error) {
	return _ETHRota.Contract.Zeros(&_ETHRota.CallOpts, i)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(bytes32)
func (_ETHRota *ETHRotaCallerSession) Zeros(i *big.Int) ([32]byte, error) {
	return _ETHRota.Contract.Zeros(&_ETHRota.CallOpts, i)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_ETHRota *ETHRotaTransactor) Deposit(opts *bind.TransactOpts, _commitment [32]byte) (*types.Transaction, error) {
	return _ETHRota.contract.Transact(opts, "deposit", _commitment)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_ETHRota *ETHRotaSession) Deposit(_commitment [32]byte) (*types.Transaction, error) {
	return _ETHRota.Contract.Deposit(&_ETHRota.TransactOpts, _commitment)
}

// Deposit is a paid mutator transaction binding the contract method 0xb214faa5.
//
// Solidity: function deposit(bytes32 _commitment) payable returns()
func (_ETHRota *ETHRotaTransactorSession) Deposit(_commitment [32]byte) (*types.Transaction, error) {
	return _ETHRota.Contract.Deposit(&_ETHRota.TransactOpts, _commitment)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ETHRota *ETHRotaTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _ETHRota.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ETHRota *ETHRotaSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _ETHRota.Contract.SetFee(&_ETHRota.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_ETHRota *ETHRotaTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _ETHRota.Contract.SetFee(&_ETHRota.TransactOpts, _fee)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_ETHRota *ETHRotaTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _ETHRota.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_ETHRota *ETHRotaSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _ETHRota.Contract.SetFeeTo(&_ETHRota.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_ETHRota *ETHRotaTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _ETHRota.Contract.SetFeeTo(&_ETHRota.TransactOpts, _feeTo)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb130eaa6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient) payable returns()
func (_ETHRota *ETHRotaTransactor) Withdraw(opts *bind.TransactOpts, _proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address) (*types.Transaction, error) {
	return _ETHRota.contract.Transact(opts, "withdraw", _proof, _root, _nullifierHash, _recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb130eaa6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient) payable returns()
func (_ETHRota *ETHRotaSession) Withdraw(_proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address) (*types.Transaction, error) {
	return _ETHRota.Contract.Withdraw(&_ETHRota.TransactOpts, _proof, _root, _nullifierHash, _recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb130eaa6.
//
// Solidity: function withdraw(bytes _proof, bytes32 _root, bytes32 _nullifierHash, address _recipient) payable returns()
func (_ETHRota *ETHRotaTransactorSession) Withdraw(_proof []byte, _root [32]byte, _nullifierHash [32]byte, _recipient common.Address) (*types.Transaction, error) {
	return _ETHRota.Contract.Withdraw(&_ETHRota.TransactOpts, _proof, _root, _nullifierHash, _recipient)
}

// ETHRotaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ETHRota contract.
type ETHRotaDepositIterator struct {
	Event *ETHRotaDeposit // Event containing the contract specifics and raw log

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
func (it *ETHRotaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHRotaDeposit)
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
		it.Event = new(ETHRotaDeposit)
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
func (it *ETHRotaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHRotaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHRotaDeposit represents a Deposit event raised by the ETHRota contract.
type ETHRotaDeposit struct {
	Commitment [32]byte
	LeafIndex  uint32
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xa945e51eec50ab98c161376f0db4cf2aeba3ec92755fe2fcd388bdbbb80ff196.
//
// Solidity: event Deposit(bytes32 indexed commitment, uint32 leafIndex, uint256 timestamp)
func (_ETHRota *ETHRotaFilterer) FilterDeposit(opts *bind.FilterOpts, commitment [][32]byte) (*ETHRotaDepositIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ETHRota.contract.FilterLogs(opts, "Deposit", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &ETHRotaDepositIterator{contract: _ETHRota.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xa945e51eec50ab98c161376f0db4cf2aeba3ec92755fe2fcd388bdbbb80ff196.
//
// Solidity: event Deposit(bytes32 indexed commitment, uint32 leafIndex, uint256 timestamp)
func (_ETHRota *ETHRotaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ETHRotaDeposit, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ETHRota.contract.WatchLogs(opts, "Deposit", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHRotaDeposit)
				if err := _ETHRota.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xa945e51eec50ab98c161376f0db4cf2aeba3ec92755fe2fcd388bdbbb80ff196.
//
// Solidity: event Deposit(bytes32 indexed commitment, uint32 leafIndex, uint256 timestamp)
func (_ETHRota *ETHRotaFilterer) ParseDeposit(log types.Log) (*ETHRotaDeposit, error) {
	event := new(ETHRotaDeposit)
	if err := _ETHRota.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHRotaWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ETHRota contract.
type ETHRotaWithdrawalIterator struct {
	Event *ETHRotaWithdrawal // Event containing the contract specifics and raw log

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
func (it *ETHRotaWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHRotaWithdrawal)
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
		it.Event = new(ETHRotaWithdrawal)
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
func (it *ETHRotaWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHRotaWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHRotaWithdrawal represents a Withdrawal event raised by the ETHRota contract.
type ETHRotaWithdrawal struct {
	To            common.Address
	NullifierHash [32]byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x4206db6775563d1043abfcf27cd0ecd19fcc464be574a1487fc95b24957a671a.
//
// Solidity: event Withdrawal(address indexed to, bytes32 nullifierHash, uint256 fee)
func (_ETHRota *ETHRotaFilterer) FilterWithdrawal(opts *bind.FilterOpts, to []common.Address) (*ETHRotaWithdrawalIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ETHRota.contract.FilterLogs(opts, "Withdrawal", toRule)
	if err != nil {
		return nil, err
	}
	return &ETHRotaWithdrawalIterator{contract: _ETHRota.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x4206db6775563d1043abfcf27cd0ecd19fcc464be574a1487fc95b24957a671a.
//
// Solidity: event Withdrawal(address indexed to, bytes32 nullifierHash, uint256 fee)
func (_ETHRota *ETHRotaFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ETHRotaWithdrawal, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ETHRota.contract.WatchLogs(opts, "Withdrawal", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHRotaWithdrawal)
				if err := _ETHRota.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x4206db6775563d1043abfcf27cd0ecd19fcc464be574a1487fc95b24957a671a.
//
// Solidity: event Withdrawal(address indexed to, bytes32 nullifierHash, uint256 fee)
func (_ETHRota *ETHRotaFilterer) ParseWithdrawal(log types.Log) (*ETHRotaWithdrawal, error) {
	event := new(ETHRotaWithdrawal)
	if err := _ETHRota.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
