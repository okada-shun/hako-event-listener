// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hkfinance

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

// HkfinanceMetaData contains all meta data concerning the Hkfinance contract.
var HkfinanceMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"memberCheckOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"memberCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debtOfHako\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"changeUpperLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"lendRecordOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upperLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hakoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newHakoOwner\",\"type\":\"address\"}],\"name\":\"changeHakoOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"joinHako\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"registerBorrowing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_debtor\",\"type\":\"address\"}],\"name\":\"debtToMemberOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lendCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hakoOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"debtToHakoOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balanceOfHakoOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_creditor\",\"type\":\"address\"}],\"name\":\"creditToMemberOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"createCredit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"borrowValueDurationOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_creditor\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"returnDebtTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"lendCredit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balanceOfHako\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"creditToHakoOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"reduceDebt\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferCredit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creditOfHako\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_debtor\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"collectDebtFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"leaveHako\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"initialUpperLimit\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"CreateCredit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReduceDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RegisterBorrowing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LendCredit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"creditor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"debtor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CollectDebtFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"debtor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"creditor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ReturnDebtTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newMember\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"JoinHako\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LeaveHako\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DepositToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferCredit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldHakoOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newHakoOwner\",\"type\":\"address\"}],\"name\":\"ChangeHakoOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hakoOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newUpperLimit\",\"type\":\"uint256\"}],\"name\":\"ChangeUpperLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hakoOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"rewardValue\",\"type\":\"uint256\"}],\"name\":\"GetReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
}

// HkfinanceABI is the input ABI used to generate the binding from.
// Deprecated: Use HkfinanceMetaData.ABI instead.
var HkfinanceABI = HkfinanceMetaData.ABI

// Hkfinance is an auto generated Go binding around an Ethereum contract.
type Hkfinance struct {
	HkfinanceCaller     // Read-only binding to the contract
	HkfinanceTransactor // Write-only binding to the contract
	HkfinanceFilterer   // Log filterer for contract events
}

// HkfinanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type HkfinanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HkfinanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HkfinanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HkfinanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HkfinanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HkfinanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HkfinanceSession struct {
	Contract     *Hkfinance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HkfinanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HkfinanceCallerSession struct {
	Contract *HkfinanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HkfinanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HkfinanceTransactorSession struct {
	Contract     *HkfinanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HkfinanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type HkfinanceRaw struct {
	Contract *Hkfinance // Generic contract binding to access the raw methods on
}

// HkfinanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HkfinanceCallerRaw struct {
	Contract *HkfinanceCaller // Generic read-only contract binding to access the raw methods on
}

// HkfinanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HkfinanceTransactorRaw struct {
	Contract *HkfinanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHkfinance creates a new instance of Hkfinance, bound to a specific deployed contract.
func NewHkfinance(address common.Address, backend bind.ContractBackend) (*Hkfinance, error) {
	contract, err := bindHkfinance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hkfinance{HkfinanceCaller: HkfinanceCaller{contract: contract}, HkfinanceTransactor: HkfinanceTransactor{contract: contract}, HkfinanceFilterer: HkfinanceFilterer{contract: contract}}, nil
}

// NewHkfinanceCaller creates a new read-only instance of Hkfinance, bound to a specific deployed contract.
func NewHkfinanceCaller(address common.Address, caller bind.ContractCaller) (*HkfinanceCaller, error) {
	contract, err := bindHkfinance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HkfinanceCaller{contract: contract}, nil
}

// NewHkfinanceTransactor creates a new write-only instance of Hkfinance, bound to a specific deployed contract.
func NewHkfinanceTransactor(address common.Address, transactor bind.ContractTransactor) (*HkfinanceTransactor, error) {
	contract, err := bindHkfinance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HkfinanceTransactor{contract: contract}, nil
}

// NewHkfinanceFilterer creates a new log filterer instance of Hkfinance, bound to a specific deployed contract.
func NewHkfinanceFilterer(address common.Address, filterer bind.ContractFilterer) (*HkfinanceFilterer, error) {
	contract, err := bindHkfinance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HkfinanceFilterer{contract: contract}, nil
}

// bindHkfinance binds a generic wrapper to an already deployed contract.
func bindHkfinance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HkfinanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hkfinance *HkfinanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hkfinance.Contract.HkfinanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hkfinance *HkfinanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hkfinance.Contract.HkfinanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hkfinance *HkfinanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hkfinance.Contract.HkfinanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hkfinance *HkfinanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hkfinance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hkfinance *HkfinanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hkfinance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hkfinance *HkfinanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hkfinance.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Hkfinance *HkfinanceCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Hkfinance *HkfinanceSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.BalanceOf(&_Hkfinance.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.BalanceOf(&_Hkfinance.CallOpts, _owner)
}

// BalanceOfHako is a free data retrieval call binding the contract method 0xbf67020f.
//
// Solidity: function balanceOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) BalanceOfHako(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "balanceOfHako")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfHako is a free data retrieval call binding the contract method 0xbf67020f.
//
// Solidity: function balanceOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceSession) BalanceOfHako() (*big.Int, error) {
	return _Hkfinance.Contract.BalanceOfHako(&_Hkfinance.CallOpts)
}

// BalanceOfHako is a free data retrieval call binding the contract method 0xbf67020f.
//
// Solidity: function balanceOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) BalanceOfHako() (*big.Int, error) {
	return _Hkfinance.Contract.BalanceOfHako(&_Hkfinance.CallOpts)
}

// BalanceOfHakoOwner is a free data retrieval call binding the contract method 0x79f1f0e3.
//
// Solidity: function balanceOfHakoOwner() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) BalanceOfHakoOwner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "balanceOfHakoOwner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfHakoOwner is a free data retrieval call binding the contract method 0x79f1f0e3.
//
// Solidity: function balanceOfHakoOwner() view returns(uint256)
func (_Hkfinance *HkfinanceSession) BalanceOfHakoOwner() (*big.Int, error) {
	return _Hkfinance.Contract.BalanceOfHakoOwner(&_Hkfinance.CallOpts)
}

// BalanceOfHakoOwner is a free data retrieval call binding the contract method 0x79f1f0e3.
//
// Solidity: function balanceOfHakoOwner() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) BalanceOfHakoOwner() (*big.Int, error) {
	return _Hkfinance.Contract.BalanceOfHakoOwner(&_Hkfinance.CallOpts)
}

// BorrowValueDurationOf is a free data retrieval call binding the contract method 0x9635721c.
//
// Solidity: function borrowValueDurationOf(address _member) view returns(uint256[2])
func (_Hkfinance *HkfinanceCaller) BorrowValueDurationOf(opts *bind.CallOpts, _member common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "borrowValueDurationOf", _member)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// BorrowValueDurationOf is a free data retrieval call binding the contract method 0x9635721c.
//
// Solidity: function borrowValueDurationOf(address _member) view returns(uint256[2])
func (_Hkfinance *HkfinanceSession) BorrowValueDurationOf(_member common.Address) ([2]*big.Int, error) {
	return _Hkfinance.Contract.BorrowValueDurationOf(&_Hkfinance.CallOpts, _member)
}

// BorrowValueDurationOf is a free data retrieval call binding the contract method 0x9635721c.
//
// Solidity: function borrowValueDurationOf(address _member) view returns(uint256[2])
func (_Hkfinance *HkfinanceCallerSession) BorrowValueDurationOf(_member common.Address) ([2]*big.Int, error) {
	return _Hkfinance.Contract.BorrowValueDurationOf(&_Hkfinance.CallOpts, _member)
}

// CreditOfHako is a free data retrieval call binding the contract method 0xea413efd.
//
// Solidity: function creditOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) CreditOfHako(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "creditOfHako")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditOfHako is a free data retrieval call binding the contract method 0xea413efd.
//
// Solidity: function creditOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceSession) CreditOfHako() (*big.Int, error) {
	return _Hkfinance.Contract.CreditOfHako(&_Hkfinance.CallOpts)
}

// CreditOfHako is a free data retrieval call binding the contract method 0xea413efd.
//
// Solidity: function creditOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) CreditOfHako() (*big.Int, error) {
	return _Hkfinance.Contract.CreditOfHako(&_Hkfinance.CallOpts)
}

// CreditToHakoOf is a free data retrieval call binding the contract method 0xc9bb5709.
//
// Solidity: function creditToHakoOf(address _member) view returns(uint256)
func (_Hkfinance *HkfinanceCaller) CreditToHakoOf(opts *bind.CallOpts, _member common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "creditToHakoOf", _member)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditToHakoOf is a free data retrieval call binding the contract method 0xc9bb5709.
//
// Solidity: function creditToHakoOf(address _member) view returns(uint256)
func (_Hkfinance *HkfinanceSession) CreditToHakoOf(_member common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.CreditToHakoOf(&_Hkfinance.CallOpts, _member)
}

// CreditToHakoOf is a free data retrieval call binding the contract method 0xc9bb5709.
//
// Solidity: function creditToHakoOf(address _member) view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) CreditToHakoOf(_member common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.CreditToHakoOf(&_Hkfinance.CallOpts, _member)
}

// CreditToMemberOf is a free data retrieval call binding the contract method 0x80254a51.
//
// Solidity: function creditToMemberOf(address _creditor) view returns(uint256)
func (_Hkfinance *HkfinanceCaller) CreditToMemberOf(opts *bind.CallOpts, _creditor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "creditToMemberOf", _creditor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditToMemberOf is a free data retrieval call binding the contract method 0x80254a51.
//
// Solidity: function creditToMemberOf(address _creditor) view returns(uint256)
func (_Hkfinance *HkfinanceSession) CreditToMemberOf(_creditor common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.CreditToMemberOf(&_Hkfinance.CallOpts, _creditor)
}

// CreditToMemberOf is a free data retrieval call binding the contract method 0x80254a51.
//
// Solidity: function creditToMemberOf(address _creditor) view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) CreditToMemberOf(_creditor common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.CreditToMemberOf(&_Hkfinance.CallOpts, _creditor)
}

// DebtOfHako is a free data retrieval call binding the contract method 0x176db3d8.
//
// Solidity: function debtOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) DebtOfHako(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "debtOfHako")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOfHako is a free data retrieval call binding the contract method 0x176db3d8.
//
// Solidity: function debtOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceSession) DebtOfHako() (*big.Int, error) {
	return _Hkfinance.Contract.DebtOfHako(&_Hkfinance.CallOpts)
}

// DebtOfHako is a free data retrieval call binding the contract method 0x176db3d8.
//
// Solidity: function debtOfHako() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) DebtOfHako() (*big.Int, error) {
	return _Hkfinance.Contract.DebtOfHako(&_Hkfinance.CallOpts)
}

// DebtToHakoOf is a free data retrieval call binding the contract method 0x74c6333e.
//
// Solidity: function debtToHakoOf(address _member) view returns(uint256)
func (_Hkfinance *HkfinanceCaller) DebtToHakoOf(opts *bind.CallOpts, _member common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "debtToHakoOf", _member)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtToHakoOf is a free data retrieval call binding the contract method 0x74c6333e.
//
// Solidity: function debtToHakoOf(address _member) view returns(uint256)
func (_Hkfinance *HkfinanceSession) DebtToHakoOf(_member common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.DebtToHakoOf(&_Hkfinance.CallOpts, _member)
}

// DebtToHakoOf is a free data retrieval call binding the contract method 0x74c6333e.
//
// Solidity: function debtToHakoOf(address _member) view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) DebtToHakoOf(_member common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.DebtToHakoOf(&_Hkfinance.CallOpts, _member)
}

// DebtToMemberOf is a free data retrieval call binding the contract method 0x59acbf00.
//
// Solidity: function debtToMemberOf(address _debtor) view returns(uint256)
func (_Hkfinance *HkfinanceCaller) DebtToMemberOf(opts *bind.CallOpts, _debtor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "debtToMemberOf", _debtor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtToMemberOf is a free data retrieval call binding the contract method 0x59acbf00.
//
// Solidity: function debtToMemberOf(address _debtor) view returns(uint256)
func (_Hkfinance *HkfinanceSession) DebtToMemberOf(_debtor common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.DebtToMemberOf(&_Hkfinance.CallOpts, _debtor)
}

// DebtToMemberOf is a free data retrieval call binding the contract method 0x59acbf00.
//
// Solidity: function debtToMemberOf(address _debtor) view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) DebtToMemberOf(_debtor common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.DebtToMemberOf(&_Hkfinance.CallOpts, _debtor)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Hkfinance *HkfinanceSession) Decimals() (*big.Int, error) {
	return _Hkfinance.Contract.Decimals(&_Hkfinance.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) Decimals() (*big.Int, error) {
	return _Hkfinance.Contract.Decimals(&_Hkfinance.CallOpts)
}

// HakoAddress is a free data retrieval call binding the contract method 0x3e56c4c8.
//
// Solidity: function hakoAddress() view returns(address)
func (_Hkfinance *HkfinanceCaller) HakoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "hakoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HakoAddress is a free data retrieval call binding the contract method 0x3e56c4c8.
//
// Solidity: function hakoAddress() view returns(address)
func (_Hkfinance *HkfinanceSession) HakoAddress() (common.Address, error) {
	return _Hkfinance.Contract.HakoAddress(&_Hkfinance.CallOpts)
}

// HakoAddress is a free data retrieval call binding the contract method 0x3e56c4c8.
//
// Solidity: function hakoAddress() view returns(address)
func (_Hkfinance *HkfinanceCallerSession) HakoAddress() (common.Address, error) {
	return _Hkfinance.Contract.HakoAddress(&_Hkfinance.CallOpts)
}

// HakoOwner is a free data retrieval call binding the contract method 0x6f66068b.
//
// Solidity: function hakoOwner() view returns(address)
func (_Hkfinance *HkfinanceCaller) HakoOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "hakoOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HakoOwner is a free data retrieval call binding the contract method 0x6f66068b.
//
// Solidity: function hakoOwner() view returns(address)
func (_Hkfinance *HkfinanceSession) HakoOwner() (common.Address, error) {
	return _Hkfinance.Contract.HakoOwner(&_Hkfinance.CallOpts)
}

// HakoOwner is a free data retrieval call binding the contract method 0x6f66068b.
//
// Solidity: function hakoOwner() view returns(address)
func (_Hkfinance *HkfinanceCallerSession) HakoOwner() (common.Address, error) {
	return _Hkfinance.Contract.HakoOwner(&_Hkfinance.CallOpts)
}

// LendCount is a free data retrieval call binding the contract method 0x699e0f56.
//
// Solidity: function lendCount() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) LendCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "lendCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LendCount is a free data retrieval call binding the contract method 0x699e0f56.
//
// Solidity: function lendCount() view returns(uint256)
func (_Hkfinance *HkfinanceSession) LendCount() (*big.Int, error) {
	return _Hkfinance.Contract.LendCount(&_Hkfinance.CallOpts)
}

// LendCount is a free data retrieval call binding the contract method 0x699e0f56.
//
// Solidity: function lendCount() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) LendCount() (*big.Int, error) {
	return _Hkfinance.Contract.LendCount(&_Hkfinance.CallOpts)
}

// LendRecordOf is a free data retrieval call binding the contract method 0x341fc8ca.
//
// Solidity: function lendRecordOf(uint256 _id) view returns(uint256, address, address, uint256, uint256, uint256, uint256)
func (_Hkfinance *HkfinanceCaller) LendRecordOf(opts *bind.CallOpts, _id *big.Int) (*big.Int, common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "lendRecordOf", _id)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// LendRecordOf is a free data retrieval call binding the contract method 0x341fc8ca.
//
// Solidity: function lendRecordOf(uint256 _id) view returns(uint256, address, address, uint256, uint256, uint256, uint256)
func (_Hkfinance *HkfinanceSession) LendRecordOf(_id *big.Int) (*big.Int, common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Hkfinance.Contract.LendRecordOf(&_Hkfinance.CallOpts, _id)
}

// LendRecordOf is a free data retrieval call binding the contract method 0x341fc8ca.
//
// Solidity: function lendRecordOf(uint256 _id) view returns(uint256, address, address, uint256, uint256, uint256, uint256)
func (_Hkfinance *HkfinanceCallerSession) LendRecordOf(_id *big.Int) (*big.Int, common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Hkfinance.Contract.LendRecordOf(&_Hkfinance.CallOpts, _id)
}

// MemberCheckOf is a free data retrieval call binding the contract method 0x06287b7d.
//
// Solidity: function memberCheckOf(address _who) view returns(uint256)
func (_Hkfinance *HkfinanceCaller) MemberCheckOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "memberCheckOf", _who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MemberCheckOf is a free data retrieval call binding the contract method 0x06287b7d.
//
// Solidity: function memberCheckOf(address _who) view returns(uint256)
func (_Hkfinance *HkfinanceSession) MemberCheckOf(_who common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.MemberCheckOf(&_Hkfinance.CallOpts, _who)
}

// MemberCheckOf is a free data retrieval call binding the contract method 0x06287b7d.
//
// Solidity: function memberCheckOf(address _who) view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) MemberCheckOf(_who common.Address) (*big.Int, error) {
	return _Hkfinance.Contract.MemberCheckOf(&_Hkfinance.CallOpts, _who)
}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) MemberCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "memberCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() view returns(uint256)
func (_Hkfinance *HkfinanceSession) MemberCount() (*big.Int, error) {
	return _Hkfinance.Contract.MemberCount(&_Hkfinance.CallOpts)
}

// MemberCount is a free data retrieval call binding the contract method 0x11aee380.
//
// Solidity: function memberCount() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) MemberCount() (*big.Int, error) {
	return _Hkfinance.Contract.MemberCount(&_Hkfinance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hkfinance *HkfinanceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hkfinance *HkfinanceSession) Name() (string, error) {
	return _Hkfinance.Contract.Name(&_Hkfinance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hkfinance *HkfinanceCallerSession) Name() (string, error) {
	return _Hkfinance.Contract.Name(&_Hkfinance.CallOpts)
}

// RewardTime is a free data retrieval call binding the contract method 0x91b66fd5.
//
// Solidity: function rewardTime() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) RewardTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "rewardTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTime is a free data retrieval call binding the contract method 0x91b66fd5.
//
// Solidity: function rewardTime() view returns(uint256)
func (_Hkfinance *HkfinanceSession) RewardTime() (*big.Int, error) {
	return _Hkfinance.Contract.RewardTime(&_Hkfinance.CallOpts)
}

// RewardTime is a free data retrieval call binding the contract method 0x91b66fd5.
//
// Solidity: function rewardTime() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) RewardTime() (*big.Int, error) {
	return _Hkfinance.Contract.RewardTime(&_Hkfinance.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hkfinance *HkfinanceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hkfinance *HkfinanceSession) Symbol() (string, error) {
	return _Hkfinance.Contract.Symbol(&_Hkfinance.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hkfinance *HkfinanceCallerSession) Symbol() (string, error) {
	return _Hkfinance.Contract.Symbol(&_Hkfinance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hkfinance *HkfinanceSession) TotalSupply() (*big.Int, error) {
	return _Hkfinance.Contract.TotalSupply(&_Hkfinance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) TotalSupply() (*big.Int, error) {
	return _Hkfinance.Contract.TotalSupply(&_Hkfinance.CallOpts)
}

// UpperLimit is a free data retrieval call binding the contract method 0x38392c40.
//
// Solidity: function upperLimit() view returns(uint256)
func (_Hkfinance *HkfinanceCaller) UpperLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hkfinance.contract.Call(opts, &out, "upperLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpperLimit is a free data retrieval call binding the contract method 0x38392c40.
//
// Solidity: function upperLimit() view returns(uint256)
func (_Hkfinance *HkfinanceSession) UpperLimit() (*big.Int, error) {
	return _Hkfinance.Contract.UpperLimit(&_Hkfinance.CallOpts)
}

// UpperLimit is a free data retrieval call binding the contract method 0x38392c40.
//
// Solidity: function upperLimit() view returns(uint256)
func (_Hkfinance *HkfinanceCallerSession) UpperLimit() (*big.Int, error) {
	return _Hkfinance.Contract.UpperLimit(&_Hkfinance.CallOpts)
}

// ChangeHakoOwner is a paid mutator transaction binding the contract method 0x3f04050c.
//
// Solidity: function changeHakoOwner(address _newHakoOwner) returns(bool)
func (_Hkfinance *HkfinanceTransactor) ChangeHakoOwner(opts *bind.TransactOpts, _newHakoOwner common.Address) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "changeHakoOwner", _newHakoOwner)
}

// ChangeHakoOwner is a paid mutator transaction binding the contract method 0x3f04050c.
//
// Solidity: function changeHakoOwner(address _newHakoOwner) returns(bool)
func (_Hkfinance *HkfinanceSession) ChangeHakoOwner(_newHakoOwner common.Address) (*types.Transaction, error) {
	return _Hkfinance.Contract.ChangeHakoOwner(&_Hkfinance.TransactOpts, _newHakoOwner)
}

// ChangeHakoOwner is a paid mutator transaction binding the contract method 0x3f04050c.
//
// Solidity: function changeHakoOwner(address _newHakoOwner) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) ChangeHakoOwner(_newHakoOwner common.Address) (*types.Transaction, error) {
	return _Hkfinance.Contract.ChangeHakoOwner(&_Hkfinance.TransactOpts, _newHakoOwner)
}

// ChangeUpperLimit is a paid mutator transaction binding the contract method 0x23a4a74c.
//
// Solidity: function changeUpperLimit(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactor) ChangeUpperLimit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "changeUpperLimit", _value)
}

// ChangeUpperLimit is a paid mutator transaction binding the contract method 0x23a4a74c.
//
// Solidity: function changeUpperLimit(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceSession) ChangeUpperLimit(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.ChangeUpperLimit(&_Hkfinance.TransactOpts, _value)
}

// ChangeUpperLimit is a paid mutator transaction binding the contract method 0x23a4a74c.
//
// Solidity: function changeUpperLimit(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) ChangeUpperLimit(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.ChangeUpperLimit(&_Hkfinance.TransactOpts, _value)
}

// CollectDebtFrom is a paid mutator transaction binding the contract method 0xebeea838.
//
// Solidity: function collectDebtFrom(address _debtor, uint256 _id) returns(bool)
func (_Hkfinance *HkfinanceTransactor) CollectDebtFrom(opts *bind.TransactOpts, _debtor common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "collectDebtFrom", _debtor, _id)
}

// CollectDebtFrom is a paid mutator transaction binding the contract method 0xebeea838.
//
// Solidity: function collectDebtFrom(address _debtor, uint256 _id) returns(bool)
func (_Hkfinance *HkfinanceSession) CollectDebtFrom(_debtor common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.CollectDebtFrom(&_Hkfinance.TransactOpts, _debtor, _id)
}

// CollectDebtFrom is a paid mutator transaction binding the contract method 0xebeea838.
//
// Solidity: function collectDebtFrom(address _debtor, uint256 _id) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) CollectDebtFrom(_debtor common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.CollectDebtFrom(&_Hkfinance.TransactOpts, _debtor, _id)
}

// CreateCredit is a paid mutator transaction binding the contract method 0x8727c78b.
//
// Solidity: function createCredit(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactor) CreateCredit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "createCredit", _value)
}

// CreateCredit is a paid mutator transaction binding the contract method 0x8727c78b.
//
// Solidity: function createCredit(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceSession) CreateCredit(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.CreateCredit(&_Hkfinance.TransactOpts, _value)
}

// CreateCredit is a paid mutator transaction binding the contract method 0x8727c78b.
//
// Solidity: function createCredit(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) CreateCredit(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.CreateCredit(&_Hkfinance.TransactOpts, _value)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactor) DepositToken(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "depositToken", _value)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceSession) DepositToken(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.DepositToken(&_Hkfinance.TransactOpts, _value)
}

// DepositToken is a paid mutator transaction binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) DepositToken(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.DepositToken(&_Hkfinance.TransactOpts, _value)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_Hkfinance *HkfinanceTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_Hkfinance *HkfinanceSession) GetReward() (*types.Transaction, error) {
	return _Hkfinance.Contract.GetReward(&_Hkfinance.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) GetReward() (*types.Transaction, error) {
	return _Hkfinance.Contract.GetReward(&_Hkfinance.TransactOpts)
}

// JoinHako is a paid mutator transaction binding the contract method 0x42f274df.
//
// Solidity: function joinHako(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactor) JoinHako(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "joinHako", _value)
}

// JoinHako is a paid mutator transaction binding the contract method 0x42f274df.
//
// Solidity: function joinHako(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceSession) JoinHako(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.JoinHako(&_Hkfinance.TransactOpts, _value)
}

// JoinHako is a paid mutator transaction binding the contract method 0x42f274df.
//
// Solidity: function joinHako(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) JoinHako(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.JoinHako(&_Hkfinance.TransactOpts, _value)
}

// LeaveHako is a paid mutator transaction binding the contract method 0xfbccb503.
//
// Solidity: function leaveHako() returns(bool)
func (_Hkfinance *HkfinanceTransactor) LeaveHako(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "leaveHako")
}

// LeaveHako is a paid mutator transaction binding the contract method 0xfbccb503.
//
// Solidity: function leaveHako() returns(bool)
func (_Hkfinance *HkfinanceSession) LeaveHako() (*types.Transaction, error) {
	return _Hkfinance.Contract.LeaveHako(&_Hkfinance.TransactOpts)
}

// LeaveHako is a paid mutator transaction binding the contract method 0xfbccb503.
//
// Solidity: function leaveHako() returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) LeaveHako() (*types.Transaction, error) {
	return _Hkfinance.Contract.LeaveHako(&_Hkfinance.TransactOpts)
}

// LendCredit is a paid mutator transaction binding the contract method 0xb332d1ec.
//
// Solidity: function lendCredit(address _to, uint256 _value, uint256 _duration) returns(bool)
func (_Hkfinance *HkfinanceTransactor) LendCredit(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "lendCredit", _to, _value, _duration)
}

// LendCredit is a paid mutator transaction binding the contract method 0xb332d1ec.
//
// Solidity: function lendCredit(address _to, uint256 _value, uint256 _duration) returns(bool)
func (_Hkfinance *HkfinanceSession) LendCredit(_to common.Address, _value *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.LendCredit(&_Hkfinance.TransactOpts, _to, _value, _duration)
}

// LendCredit is a paid mutator transaction binding the contract method 0xb332d1ec.
//
// Solidity: function lendCredit(address _to, uint256 _value, uint256 _duration) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) LendCredit(_to common.Address, _value *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.LendCredit(&_Hkfinance.TransactOpts, _to, _value, _duration)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactor) ReduceDebt(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "reduceDebt", _value)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceSession) ReduceDebt(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.ReduceDebt(&_Hkfinance.TransactOpts, _value)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) ReduceDebt(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.ReduceDebt(&_Hkfinance.TransactOpts, _value)
}

// RegisterBorrowing is a paid mutator transaction binding the contract method 0x46fc9948.
//
// Solidity: function registerBorrowing(uint256 _value, uint256 _duration) returns(bool)
func (_Hkfinance *HkfinanceTransactor) RegisterBorrowing(opts *bind.TransactOpts, _value *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "registerBorrowing", _value, _duration)
}

// RegisterBorrowing is a paid mutator transaction binding the contract method 0x46fc9948.
//
// Solidity: function registerBorrowing(uint256 _value, uint256 _duration) returns(bool)
func (_Hkfinance *HkfinanceSession) RegisterBorrowing(_value *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.RegisterBorrowing(&_Hkfinance.TransactOpts, _value, _duration)
}

// RegisterBorrowing is a paid mutator transaction binding the contract method 0x46fc9948.
//
// Solidity: function registerBorrowing(uint256 _value, uint256 _duration) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) RegisterBorrowing(_value *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.RegisterBorrowing(&_Hkfinance.TransactOpts, _value, _duration)
}

// ReturnDebtTo is a paid mutator transaction binding the contract method 0x976643a6.
//
// Solidity: function returnDebtTo(address _creditor, uint256 _id) returns(bool)
func (_Hkfinance *HkfinanceTransactor) ReturnDebtTo(opts *bind.TransactOpts, _creditor common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "returnDebtTo", _creditor, _id)
}

// ReturnDebtTo is a paid mutator transaction binding the contract method 0x976643a6.
//
// Solidity: function returnDebtTo(address _creditor, uint256 _id) returns(bool)
func (_Hkfinance *HkfinanceSession) ReturnDebtTo(_creditor common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.ReturnDebtTo(&_Hkfinance.TransactOpts, _creditor, _id)
}

// ReturnDebtTo is a paid mutator transaction binding the contract method 0x976643a6.
//
// Solidity: function returnDebtTo(address _creditor, uint256 _id) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) ReturnDebtTo(_creditor common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.ReturnDebtTo(&_Hkfinance.TransactOpts, _creditor, _id)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.Transfer(&_Hkfinance.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.Transfer(&_Hkfinance.TransactOpts, _to, _value)
}

// TransferCredit is a paid mutator transaction binding the contract method 0xcc58a6bb.
//
// Solidity: function transferCredit(address _to, uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactor) TransferCredit(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "transferCredit", _to, _value)
}

// TransferCredit is a paid mutator transaction binding the contract method 0xcc58a6bb.
//
// Solidity: function transferCredit(address _to, uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceSession) TransferCredit(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.TransferCredit(&_Hkfinance.TransactOpts, _to, _value)
}

// TransferCredit is a paid mutator transaction binding the contract method 0xcc58a6bb.
//
// Solidity: function transferCredit(address _to, uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) TransferCredit(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.TransferCredit(&_Hkfinance.TransactOpts, _to, _value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactor) WithdrawToken(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.contract.Transact(opts, "withdrawToken", _value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceSession) WithdrawToken(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.WithdrawToken(&_Hkfinance.TransactOpts, _value)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x50baa622.
//
// Solidity: function withdrawToken(uint256 _value) returns(bool)
func (_Hkfinance *HkfinanceTransactorSession) WithdrawToken(_value *big.Int) (*types.Transaction, error) {
	return _Hkfinance.Contract.WithdrawToken(&_Hkfinance.TransactOpts, _value)
}

// HkfinanceChangeHakoOwnerIterator is returned from FilterChangeHakoOwner and is used to iterate over the raw logs and unpacked data for ChangeHakoOwner events raised by the Hkfinance contract.
type HkfinanceChangeHakoOwnerIterator struct {
	Event *HkfinanceChangeHakoOwner // Event containing the contract specifics and raw log

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
func (it *HkfinanceChangeHakoOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceChangeHakoOwner)
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
		it.Event = new(HkfinanceChangeHakoOwner)
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
func (it *HkfinanceChangeHakoOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceChangeHakoOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceChangeHakoOwner represents a ChangeHakoOwner event raised by the Hkfinance contract.
type HkfinanceChangeHakoOwner struct {
	OldHakoOwner common.Address
	NewHakoOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChangeHakoOwner is a free log retrieval operation binding the contract event 0xd4a5b0e93b2ae5e48114c9d0fc0be8eed177183ca40f06b7ad7088d189360757.
//
// Solidity: event ChangeHakoOwner(address indexed oldHakoOwner, address indexed newHakoOwner)
func (_Hkfinance *HkfinanceFilterer) FilterChangeHakoOwner(opts *bind.FilterOpts, oldHakoOwner []common.Address, newHakoOwner []common.Address) (*HkfinanceChangeHakoOwnerIterator, error) {

	var oldHakoOwnerRule []interface{}
	for _, oldHakoOwnerItem := range oldHakoOwner {
		oldHakoOwnerRule = append(oldHakoOwnerRule, oldHakoOwnerItem)
	}
	var newHakoOwnerRule []interface{}
	for _, newHakoOwnerItem := range newHakoOwner {
		newHakoOwnerRule = append(newHakoOwnerRule, newHakoOwnerItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "ChangeHakoOwner", oldHakoOwnerRule, newHakoOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceChangeHakoOwnerIterator{contract: _Hkfinance.contract, event: "ChangeHakoOwner", logs: logs, sub: sub}, nil
}

// WatchChangeHakoOwner is a free log subscription operation binding the contract event 0xd4a5b0e93b2ae5e48114c9d0fc0be8eed177183ca40f06b7ad7088d189360757.
//
// Solidity: event ChangeHakoOwner(address indexed oldHakoOwner, address indexed newHakoOwner)
func (_Hkfinance *HkfinanceFilterer) WatchChangeHakoOwner(opts *bind.WatchOpts, sink chan<- *HkfinanceChangeHakoOwner, oldHakoOwner []common.Address, newHakoOwner []common.Address) (event.Subscription, error) {

	var oldHakoOwnerRule []interface{}
	for _, oldHakoOwnerItem := range oldHakoOwner {
		oldHakoOwnerRule = append(oldHakoOwnerRule, oldHakoOwnerItem)
	}
	var newHakoOwnerRule []interface{}
	for _, newHakoOwnerItem := range newHakoOwner {
		newHakoOwnerRule = append(newHakoOwnerRule, newHakoOwnerItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "ChangeHakoOwner", oldHakoOwnerRule, newHakoOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceChangeHakoOwner)
				if err := _Hkfinance.contract.UnpackLog(event, "ChangeHakoOwner", log); err != nil {
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

// ParseChangeHakoOwner is a log parse operation binding the contract event 0xd4a5b0e93b2ae5e48114c9d0fc0be8eed177183ca40f06b7ad7088d189360757.
//
// Solidity: event ChangeHakoOwner(address indexed oldHakoOwner, address indexed newHakoOwner)
func (_Hkfinance *HkfinanceFilterer) ParseChangeHakoOwner(log types.Log) (*HkfinanceChangeHakoOwner, error) {
	event := new(HkfinanceChangeHakoOwner)
	if err := _Hkfinance.contract.UnpackLog(event, "ChangeHakoOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceChangeUpperLimitIterator is returned from FilterChangeUpperLimit and is used to iterate over the raw logs and unpacked data for ChangeUpperLimit events raised by the Hkfinance contract.
type HkfinanceChangeUpperLimitIterator struct {
	Event *HkfinanceChangeUpperLimit // Event containing the contract specifics and raw log

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
func (it *HkfinanceChangeUpperLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceChangeUpperLimit)
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
		it.Event = new(HkfinanceChangeUpperLimit)
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
func (it *HkfinanceChangeUpperLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceChangeUpperLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceChangeUpperLimit represents a ChangeUpperLimit event raised by the Hkfinance contract.
type HkfinanceChangeUpperLimit struct {
	HakoOwner     common.Address
	NewUpperLimit *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChangeUpperLimit is a free log retrieval operation binding the contract event 0x308fc51ce91b76c8ecf14fc26fbb826b3ff4ba6d92fbe0383692ecd7a533a12f.
//
// Solidity: event ChangeUpperLimit(address indexed hakoOwner, uint256 newUpperLimit)
func (_Hkfinance *HkfinanceFilterer) FilterChangeUpperLimit(opts *bind.FilterOpts, hakoOwner []common.Address) (*HkfinanceChangeUpperLimitIterator, error) {

	var hakoOwnerRule []interface{}
	for _, hakoOwnerItem := range hakoOwner {
		hakoOwnerRule = append(hakoOwnerRule, hakoOwnerItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "ChangeUpperLimit", hakoOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceChangeUpperLimitIterator{contract: _Hkfinance.contract, event: "ChangeUpperLimit", logs: logs, sub: sub}, nil
}

// WatchChangeUpperLimit is a free log subscription operation binding the contract event 0x308fc51ce91b76c8ecf14fc26fbb826b3ff4ba6d92fbe0383692ecd7a533a12f.
//
// Solidity: event ChangeUpperLimit(address indexed hakoOwner, uint256 newUpperLimit)
func (_Hkfinance *HkfinanceFilterer) WatchChangeUpperLimit(opts *bind.WatchOpts, sink chan<- *HkfinanceChangeUpperLimit, hakoOwner []common.Address) (event.Subscription, error) {

	var hakoOwnerRule []interface{}
	for _, hakoOwnerItem := range hakoOwner {
		hakoOwnerRule = append(hakoOwnerRule, hakoOwnerItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "ChangeUpperLimit", hakoOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceChangeUpperLimit)
				if err := _Hkfinance.contract.UnpackLog(event, "ChangeUpperLimit", log); err != nil {
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

// ParseChangeUpperLimit is a log parse operation binding the contract event 0x308fc51ce91b76c8ecf14fc26fbb826b3ff4ba6d92fbe0383692ecd7a533a12f.
//
// Solidity: event ChangeUpperLimit(address indexed hakoOwner, uint256 newUpperLimit)
func (_Hkfinance *HkfinanceFilterer) ParseChangeUpperLimit(log types.Log) (*HkfinanceChangeUpperLimit, error) {
	event := new(HkfinanceChangeUpperLimit)
	if err := _Hkfinance.contract.UnpackLog(event, "ChangeUpperLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceCollectDebtFromIterator is returned from FilterCollectDebtFrom and is used to iterate over the raw logs and unpacked data for CollectDebtFrom events raised by the Hkfinance contract.
type HkfinanceCollectDebtFromIterator struct {
	Event *HkfinanceCollectDebtFrom // Event containing the contract specifics and raw log

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
func (it *HkfinanceCollectDebtFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceCollectDebtFrom)
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
		it.Event = new(HkfinanceCollectDebtFrom)
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
func (it *HkfinanceCollectDebtFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceCollectDebtFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceCollectDebtFrom represents a CollectDebtFrom event raised by the Hkfinance contract.
type HkfinanceCollectDebtFrom struct {
	Creditor common.Address
	Debtor   common.Address
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCollectDebtFrom is a free log retrieval operation binding the contract event 0x59ff38dffc47484b9cf99ac7cc9c37a9acd4f0d8d8dfb9615724865c7d72c1ab.
//
// Solidity: event CollectDebtFrom(address indexed creditor, address indexed debtor, uint256 id)
func (_Hkfinance *HkfinanceFilterer) FilterCollectDebtFrom(opts *bind.FilterOpts, creditor []common.Address, debtor []common.Address) (*HkfinanceCollectDebtFromIterator, error) {

	var creditorRule []interface{}
	for _, creditorItem := range creditor {
		creditorRule = append(creditorRule, creditorItem)
	}
	var debtorRule []interface{}
	for _, debtorItem := range debtor {
		debtorRule = append(debtorRule, debtorItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "CollectDebtFrom", creditorRule, debtorRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceCollectDebtFromIterator{contract: _Hkfinance.contract, event: "CollectDebtFrom", logs: logs, sub: sub}, nil
}

// WatchCollectDebtFrom is a free log subscription operation binding the contract event 0x59ff38dffc47484b9cf99ac7cc9c37a9acd4f0d8d8dfb9615724865c7d72c1ab.
//
// Solidity: event CollectDebtFrom(address indexed creditor, address indexed debtor, uint256 id)
func (_Hkfinance *HkfinanceFilterer) WatchCollectDebtFrom(opts *bind.WatchOpts, sink chan<- *HkfinanceCollectDebtFrom, creditor []common.Address, debtor []common.Address) (event.Subscription, error) {

	var creditorRule []interface{}
	for _, creditorItem := range creditor {
		creditorRule = append(creditorRule, creditorItem)
	}
	var debtorRule []interface{}
	for _, debtorItem := range debtor {
		debtorRule = append(debtorRule, debtorItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "CollectDebtFrom", creditorRule, debtorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceCollectDebtFrom)
				if err := _Hkfinance.contract.UnpackLog(event, "CollectDebtFrom", log); err != nil {
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

// ParseCollectDebtFrom is a log parse operation binding the contract event 0x59ff38dffc47484b9cf99ac7cc9c37a9acd4f0d8d8dfb9615724865c7d72c1ab.
//
// Solidity: event CollectDebtFrom(address indexed creditor, address indexed debtor, uint256 id)
func (_Hkfinance *HkfinanceFilterer) ParseCollectDebtFrom(log types.Log) (*HkfinanceCollectDebtFrom, error) {
	event := new(HkfinanceCollectDebtFrom)
	if err := _Hkfinance.contract.UnpackLog(event, "CollectDebtFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceCreateCreditIterator is returned from FilterCreateCredit and is used to iterate over the raw logs and unpacked data for CreateCredit events raised by the Hkfinance contract.
type HkfinanceCreateCreditIterator struct {
	Event *HkfinanceCreateCredit // Event containing the contract specifics and raw log

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
func (it *HkfinanceCreateCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceCreateCredit)
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
		it.Event = new(HkfinanceCreateCredit)
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
func (it *HkfinanceCreateCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceCreateCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceCreateCredit represents a CreateCredit event raised by the Hkfinance contract.
type HkfinanceCreateCredit struct {
	Member common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreateCredit is a free log retrieval operation binding the contract event 0x99f9c157b4eb3f1cb2ee7fa9ecb9c27777c8d432759321cb54588f1500900ba3.
//
// Solidity: event CreateCredit(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) FilterCreateCredit(opts *bind.FilterOpts, member []common.Address) (*HkfinanceCreateCreditIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "CreateCredit", memberRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceCreateCreditIterator{contract: _Hkfinance.contract, event: "CreateCredit", logs: logs, sub: sub}, nil
}

// WatchCreateCredit is a free log subscription operation binding the contract event 0x99f9c157b4eb3f1cb2ee7fa9ecb9c27777c8d432759321cb54588f1500900ba3.
//
// Solidity: event CreateCredit(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) WatchCreateCredit(opts *bind.WatchOpts, sink chan<- *HkfinanceCreateCredit, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "CreateCredit", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceCreateCredit)
				if err := _Hkfinance.contract.UnpackLog(event, "CreateCredit", log); err != nil {
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

// ParseCreateCredit is a log parse operation binding the contract event 0x99f9c157b4eb3f1cb2ee7fa9ecb9c27777c8d432759321cb54588f1500900ba3.
//
// Solidity: event CreateCredit(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) ParseCreateCredit(log types.Log) (*HkfinanceCreateCredit, error) {
	event := new(HkfinanceCreateCredit)
	if err := _Hkfinance.contract.UnpackLog(event, "CreateCredit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceDepositTokenIterator is returned from FilterDepositToken and is used to iterate over the raw logs and unpacked data for DepositToken events raised by the Hkfinance contract.
type HkfinanceDepositTokenIterator struct {
	Event *HkfinanceDepositToken // Event containing the contract specifics and raw log

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
func (it *HkfinanceDepositTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceDepositToken)
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
		it.Event = new(HkfinanceDepositToken)
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
func (it *HkfinanceDepositTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceDepositTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceDepositToken represents a DepositToken event raised by the Hkfinance contract.
type HkfinanceDepositToken struct {
	Member common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositToken is a free log retrieval operation binding the contract event 0x2d0c0a8842b9944ece1495eb61121621b5e36bd6af3bba0318c695f525aef79f.
//
// Solidity: event DepositToken(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) FilterDepositToken(opts *bind.FilterOpts, member []common.Address) (*HkfinanceDepositTokenIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "DepositToken", memberRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceDepositTokenIterator{contract: _Hkfinance.contract, event: "DepositToken", logs: logs, sub: sub}, nil
}

// WatchDepositToken is a free log subscription operation binding the contract event 0x2d0c0a8842b9944ece1495eb61121621b5e36bd6af3bba0318c695f525aef79f.
//
// Solidity: event DepositToken(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) WatchDepositToken(opts *bind.WatchOpts, sink chan<- *HkfinanceDepositToken, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "DepositToken", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceDepositToken)
				if err := _Hkfinance.contract.UnpackLog(event, "DepositToken", log); err != nil {
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

// ParseDepositToken is a log parse operation binding the contract event 0x2d0c0a8842b9944ece1495eb61121621b5e36bd6af3bba0318c695f525aef79f.
//
// Solidity: event DepositToken(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) ParseDepositToken(log types.Log) (*HkfinanceDepositToken, error) {
	event := new(HkfinanceDepositToken)
	if err := _Hkfinance.contract.UnpackLog(event, "DepositToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceGetRewardIterator is returned from FilterGetReward and is used to iterate over the raw logs and unpacked data for GetReward events raised by the Hkfinance contract.
type HkfinanceGetRewardIterator struct {
	Event *HkfinanceGetReward // Event containing the contract specifics and raw log

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
func (it *HkfinanceGetRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceGetReward)
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
		it.Event = new(HkfinanceGetReward)
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
func (it *HkfinanceGetRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceGetRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceGetReward represents a GetReward event raised by the Hkfinance contract.
type HkfinanceGetReward struct {
	HakoOwner   common.Address
	RewardValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGetReward is a free log retrieval operation binding the contract event 0x25c30c62c42b51e4f667b70ef60f1f683c376f6ace28312ed45a40665e01af37.
//
// Solidity: event GetReward(address indexed hakoOwner, uint256 rewardValue)
func (_Hkfinance *HkfinanceFilterer) FilterGetReward(opts *bind.FilterOpts, hakoOwner []common.Address) (*HkfinanceGetRewardIterator, error) {

	var hakoOwnerRule []interface{}
	for _, hakoOwnerItem := range hakoOwner {
		hakoOwnerRule = append(hakoOwnerRule, hakoOwnerItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "GetReward", hakoOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceGetRewardIterator{contract: _Hkfinance.contract, event: "GetReward", logs: logs, sub: sub}, nil
}

// WatchGetReward is a free log subscription operation binding the contract event 0x25c30c62c42b51e4f667b70ef60f1f683c376f6ace28312ed45a40665e01af37.
//
// Solidity: event GetReward(address indexed hakoOwner, uint256 rewardValue)
func (_Hkfinance *HkfinanceFilterer) WatchGetReward(opts *bind.WatchOpts, sink chan<- *HkfinanceGetReward, hakoOwner []common.Address) (event.Subscription, error) {

	var hakoOwnerRule []interface{}
	for _, hakoOwnerItem := range hakoOwner {
		hakoOwnerRule = append(hakoOwnerRule, hakoOwnerItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "GetReward", hakoOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceGetReward)
				if err := _Hkfinance.contract.UnpackLog(event, "GetReward", log); err != nil {
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

// ParseGetReward is a log parse operation binding the contract event 0x25c30c62c42b51e4f667b70ef60f1f683c376f6ace28312ed45a40665e01af37.
//
// Solidity: event GetReward(address indexed hakoOwner, uint256 rewardValue)
func (_Hkfinance *HkfinanceFilterer) ParseGetReward(log types.Log) (*HkfinanceGetReward, error) {
	event := new(HkfinanceGetReward)
	if err := _Hkfinance.contract.UnpackLog(event, "GetReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceJoinHakoIterator is returned from FilterJoinHako and is used to iterate over the raw logs and unpacked data for JoinHako events raised by the Hkfinance contract.
type HkfinanceJoinHakoIterator struct {
	Event *HkfinanceJoinHako // Event containing the contract specifics and raw log

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
func (it *HkfinanceJoinHakoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceJoinHako)
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
		it.Event = new(HkfinanceJoinHako)
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
func (it *HkfinanceJoinHakoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceJoinHakoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceJoinHako represents a JoinHako event raised by the Hkfinance contract.
type HkfinanceJoinHako struct {
	NewMember common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJoinHako is a free log retrieval operation binding the contract event 0xb05efb51c3d5ae80070863c9742ead34baa1bb22d07005358c5eacc29dd99c77.
//
// Solidity: event JoinHako(address indexed newMember, uint256 value)
func (_Hkfinance *HkfinanceFilterer) FilterJoinHako(opts *bind.FilterOpts, newMember []common.Address) (*HkfinanceJoinHakoIterator, error) {

	var newMemberRule []interface{}
	for _, newMemberItem := range newMember {
		newMemberRule = append(newMemberRule, newMemberItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "JoinHako", newMemberRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceJoinHakoIterator{contract: _Hkfinance.contract, event: "JoinHako", logs: logs, sub: sub}, nil
}

// WatchJoinHako is a free log subscription operation binding the contract event 0xb05efb51c3d5ae80070863c9742ead34baa1bb22d07005358c5eacc29dd99c77.
//
// Solidity: event JoinHako(address indexed newMember, uint256 value)
func (_Hkfinance *HkfinanceFilterer) WatchJoinHako(opts *bind.WatchOpts, sink chan<- *HkfinanceJoinHako, newMember []common.Address) (event.Subscription, error) {

	var newMemberRule []interface{}
	for _, newMemberItem := range newMember {
		newMemberRule = append(newMemberRule, newMemberItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "JoinHako", newMemberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceJoinHako)
				if err := _Hkfinance.contract.UnpackLog(event, "JoinHako", log); err != nil {
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

// ParseJoinHako is a log parse operation binding the contract event 0xb05efb51c3d5ae80070863c9742ead34baa1bb22d07005358c5eacc29dd99c77.
//
// Solidity: event JoinHako(address indexed newMember, uint256 value)
func (_Hkfinance *HkfinanceFilterer) ParseJoinHako(log types.Log) (*HkfinanceJoinHako, error) {
	event := new(HkfinanceJoinHako)
	if err := _Hkfinance.contract.UnpackLog(event, "JoinHako", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceLeaveHakoIterator is returned from FilterLeaveHako and is used to iterate over the raw logs and unpacked data for LeaveHako events raised by the Hkfinance contract.
type HkfinanceLeaveHakoIterator struct {
	Event *HkfinanceLeaveHako // Event containing the contract specifics and raw log

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
func (it *HkfinanceLeaveHakoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceLeaveHako)
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
		it.Event = new(HkfinanceLeaveHako)
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
func (it *HkfinanceLeaveHakoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceLeaveHakoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceLeaveHako represents a LeaveHako event raised by the Hkfinance contract.
type HkfinanceLeaveHako struct {
	Member common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLeaveHako is a free log retrieval operation binding the contract event 0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0.
//
// Solidity: event LeaveHako(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) FilterLeaveHako(opts *bind.FilterOpts, member []common.Address) (*HkfinanceLeaveHakoIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "LeaveHako", memberRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceLeaveHakoIterator{contract: _Hkfinance.contract, event: "LeaveHako", logs: logs, sub: sub}, nil
}

// WatchLeaveHako is a free log subscription operation binding the contract event 0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0.
//
// Solidity: event LeaveHako(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) WatchLeaveHako(opts *bind.WatchOpts, sink chan<- *HkfinanceLeaveHako, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "LeaveHako", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceLeaveHako)
				if err := _Hkfinance.contract.UnpackLog(event, "LeaveHako", log); err != nil {
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

// ParseLeaveHako is a log parse operation binding the contract event 0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0.
//
// Solidity: event LeaveHako(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) ParseLeaveHako(log types.Log) (*HkfinanceLeaveHako, error) {
	event := new(HkfinanceLeaveHako)
	if err := _Hkfinance.contract.UnpackLog(event, "LeaveHako", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceLendCreditIterator is returned from FilterLendCredit and is used to iterate over the raw logs and unpacked data for LendCredit events raised by the Hkfinance contract.
type HkfinanceLendCreditIterator struct {
	Event *HkfinanceLendCredit // Event containing the contract specifics and raw log

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
func (it *HkfinanceLendCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceLendCredit)
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
		it.Event = new(HkfinanceLendCredit)
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
func (it *HkfinanceLendCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceLendCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceLendCredit represents a LendCredit event raised by the Hkfinance contract.
type HkfinanceLendCredit struct {
	From     common.Address
	To       common.Address
	Value    *big.Int
	Duration *big.Int
	Id       *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLendCredit is a free log retrieval operation binding the contract event 0x80dd2834aa0492a634508d0dd7f22dedff3f002260f33405882b54a4a6f28a41.
//
// Solidity: event LendCredit(address indexed from, address indexed to, uint256 value, uint256 duration, uint256 id, uint256 time)
func (_Hkfinance *HkfinanceFilterer) FilterLendCredit(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HkfinanceLendCreditIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "LendCredit", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceLendCreditIterator{contract: _Hkfinance.contract, event: "LendCredit", logs: logs, sub: sub}, nil
}

// WatchLendCredit is a free log subscription operation binding the contract event 0x80dd2834aa0492a634508d0dd7f22dedff3f002260f33405882b54a4a6f28a41.
//
// Solidity: event LendCredit(address indexed from, address indexed to, uint256 value, uint256 duration, uint256 id, uint256 time)
func (_Hkfinance *HkfinanceFilterer) WatchLendCredit(opts *bind.WatchOpts, sink chan<- *HkfinanceLendCredit, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "LendCredit", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceLendCredit)
				if err := _Hkfinance.contract.UnpackLog(event, "LendCredit", log); err != nil {
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

// ParseLendCredit is a log parse operation binding the contract event 0x80dd2834aa0492a634508d0dd7f22dedff3f002260f33405882b54a4a6f28a41.
//
// Solidity: event LendCredit(address indexed from, address indexed to, uint256 value, uint256 duration, uint256 id, uint256 time)
func (_Hkfinance *HkfinanceFilterer) ParseLendCredit(log types.Log) (*HkfinanceLendCredit, error) {
	event := new(HkfinanceLendCredit)
	if err := _Hkfinance.contract.UnpackLog(event, "LendCredit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceReduceDebtIterator is returned from FilterReduceDebt and is used to iterate over the raw logs and unpacked data for ReduceDebt events raised by the Hkfinance contract.
type HkfinanceReduceDebtIterator struct {
	Event *HkfinanceReduceDebt // Event containing the contract specifics and raw log

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
func (it *HkfinanceReduceDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceReduceDebt)
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
		it.Event = new(HkfinanceReduceDebt)
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
func (it *HkfinanceReduceDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceReduceDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceReduceDebt represents a ReduceDebt event raised by the Hkfinance contract.
type HkfinanceReduceDebt struct {
	Member common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReduceDebt is a free log retrieval operation binding the contract event 0x243d60a280579accabafa82109d7f67f7ce230d575b0ff7a301e8dabf3a3f505.
//
// Solidity: event ReduceDebt(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) FilterReduceDebt(opts *bind.FilterOpts, member []common.Address) (*HkfinanceReduceDebtIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "ReduceDebt", memberRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceReduceDebtIterator{contract: _Hkfinance.contract, event: "ReduceDebt", logs: logs, sub: sub}, nil
}

// WatchReduceDebt is a free log subscription operation binding the contract event 0x243d60a280579accabafa82109d7f67f7ce230d575b0ff7a301e8dabf3a3f505.
//
// Solidity: event ReduceDebt(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) WatchReduceDebt(opts *bind.WatchOpts, sink chan<- *HkfinanceReduceDebt, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "ReduceDebt", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceReduceDebt)
				if err := _Hkfinance.contract.UnpackLog(event, "ReduceDebt", log); err != nil {
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

// ParseReduceDebt is a log parse operation binding the contract event 0x243d60a280579accabafa82109d7f67f7ce230d575b0ff7a301e8dabf3a3f505.
//
// Solidity: event ReduceDebt(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) ParseReduceDebt(log types.Log) (*HkfinanceReduceDebt, error) {
	event := new(HkfinanceReduceDebt)
	if err := _Hkfinance.contract.UnpackLog(event, "ReduceDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceRegisterBorrowingIterator is returned from FilterRegisterBorrowing and is used to iterate over the raw logs and unpacked data for RegisterBorrowing events raised by the Hkfinance contract.
type HkfinanceRegisterBorrowingIterator struct {
	Event *HkfinanceRegisterBorrowing // Event containing the contract specifics and raw log

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
func (it *HkfinanceRegisterBorrowingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceRegisterBorrowing)
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
		it.Event = new(HkfinanceRegisterBorrowing)
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
func (it *HkfinanceRegisterBorrowingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceRegisterBorrowingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceRegisterBorrowing represents a RegisterBorrowing event raised by the Hkfinance contract.
type HkfinanceRegisterBorrowing struct {
	Member   common.Address
	Value    *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegisterBorrowing is a free log retrieval operation binding the contract event 0xb2b4456f5c99698ad5f7f3712f9236e3b2936700e2f1759ee08111b2c00927a1.
//
// Solidity: event RegisterBorrowing(address indexed member, uint256 value, uint256 duration)
func (_Hkfinance *HkfinanceFilterer) FilterRegisterBorrowing(opts *bind.FilterOpts, member []common.Address) (*HkfinanceRegisterBorrowingIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "RegisterBorrowing", memberRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceRegisterBorrowingIterator{contract: _Hkfinance.contract, event: "RegisterBorrowing", logs: logs, sub: sub}, nil
}

// WatchRegisterBorrowing is a free log subscription operation binding the contract event 0xb2b4456f5c99698ad5f7f3712f9236e3b2936700e2f1759ee08111b2c00927a1.
//
// Solidity: event RegisterBorrowing(address indexed member, uint256 value, uint256 duration)
func (_Hkfinance *HkfinanceFilterer) WatchRegisterBorrowing(opts *bind.WatchOpts, sink chan<- *HkfinanceRegisterBorrowing, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "RegisterBorrowing", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceRegisterBorrowing)
				if err := _Hkfinance.contract.UnpackLog(event, "RegisterBorrowing", log); err != nil {
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

// ParseRegisterBorrowing is a log parse operation binding the contract event 0xb2b4456f5c99698ad5f7f3712f9236e3b2936700e2f1759ee08111b2c00927a1.
//
// Solidity: event RegisterBorrowing(address indexed member, uint256 value, uint256 duration)
func (_Hkfinance *HkfinanceFilterer) ParseRegisterBorrowing(log types.Log) (*HkfinanceRegisterBorrowing, error) {
	event := new(HkfinanceRegisterBorrowing)
	if err := _Hkfinance.contract.UnpackLog(event, "RegisterBorrowing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceReturnDebtToIterator is returned from FilterReturnDebtTo and is used to iterate over the raw logs and unpacked data for ReturnDebtTo events raised by the Hkfinance contract.
type HkfinanceReturnDebtToIterator struct {
	Event *HkfinanceReturnDebtTo // Event containing the contract specifics and raw log

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
func (it *HkfinanceReturnDebtToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceReturnDebtTo)
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
		it.Event = new(HkfinanceReturnDebtTo)
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
func (it *HkfinanceReturnDebtToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceReturnDebtToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceReturnDebtTo represents a ReturnDebtTo event raised by the Hkfinance contract.
type HkfinanceReturnDebtTo struct {
	Debtor   common.Address
	Creditor common.Address
	Id       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReturnDebtTo is a free log retrieval operation binding the contract event 0x63fd0be6f61dcd8a83255d9e205e4c444b689733aa83dfc4fcf346f94540b327.
//
// Solidity: event ReturnDebtTo(address indexed debtor, address indexed creditor, uint256 id)
func (_Hkfinance *HkfinanceFilterer) FilterReturnDebtTo(opts *bind.FilterOpts, debtor []common.Address, creditor []common.Address) (*HkfinanceReturnDebtToIterator, error) {

	var debtorRule []interface{}
	for _, debtorItem := range debtor {
		debtorRule = append(debtorRule, debtorItem)
	}
	var creditorRule []interface{}
	for _, creditorItem := range creditor {
		creditorRule = append(creditorRule, creditorItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "ReturnDebtTo", debtorRule, creditorRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceReturnDebtToIterator{contract: _Hkfinance.contract, event: "ReturnDebtTo", logs: logs, sub: sub}, nil
}

// WatchReturnDebtTo is a free log subscription operation binding the contract event 0x63fd0be6f61dcd8a83255d9e205e4c444b689733aa83dfc4fcf346f94540b327.
//
// Solidity: event ReturnDebtTo(address indexed debtor, address indexed creditor, uint256 id)
func (_Hkfinance *HkfinanceFilterer) WatchReturnDebtTo(opts *bind.WatchOpts, sink chan<- *HkfinanceReturnDebtTo, debtor []common.Address, creditor []common.Address) (event.Subscription, error) {

	var debtorRule []interface{}
	for _, debtorItem := range debtor {
		debtorRule = append(debtorRule, debtorItem)
	}
	var creditorRule []interface{}
	for _, creditorItem := range creditor {
		creditorRule = append(creditorRule, creditorItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "ReturnDebtTo", debtorRule, creditorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceReturnDebtTo)
				if err := _Hkfinance.contract.UnpackLog(event, "ReturnDebtTo", log); err != nil {
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

// ParseReturnDebtTo is a log parse operation binding the contract event 0x63fd0be6f61dcd8a83255d9e205e4c444b689733aa83dfc4fcf346f94540b327.
//
// Solidity: event ReturnDebtTo(address indexed debtor, address indexed creditor, uint256 id)
func (_Hkfinance *HkfinanceFilterer) ParseReturnDebtTo(log types.Log) (*HkfinanceReturnDebtTo, error) {
	event := new(HkfinanceReturnDebtTo)
	if err := _Hkfinance.contract.UnpackLog(event, "ReturnDebtTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hkfinance contract.
type HkfinanceTransferIterator struct {
	Event *HkfinanceTransfer // Event containing the contract specifics and raw log

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
func (it *HkfinanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceTransfer)
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
		it.Event = new(HkfinanceTransfer)
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
func (it *HkfinanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceTransfer represents a Transfer event raised by the Hkfinance contract.
type HkfinanceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hkfinance *HkfinanceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HkfinanceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceTransferIterator{contract: _Hkfinance.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Hkfinance *HkfinanceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HkfinanceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceTransfer)
				if err := _Hkfinance.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Hkfinance *HkfinanceFilterer) ParseTransfer(log types.Log) (*HkfinanceTransfer, error) {
	event := new(HkfinanceTransfer)
	if err := _Hkfinance.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceTransferCreditIterator is returned from FilterTransferCredit and is used to iterate over the raw logs and unpacked data for TransferCredit events raised by the Hkfinance contract.
type HkfinanceTransferCreditIterator struct {
	Event *HkfinanceTransferCredit // Event containing the contract specifics and raw log

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
func (it *HkfinanceTransferCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceTransferCredit)
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
		it.Event = new(HkfinanceTransferCredit)
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
func (it *HkfinanceTransferCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceTransferCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceTransferCredit represents a TransferCredit event raised by the Hkfinance contract.
type HkfinanceTransferCredit struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferCredit is a free log retrieval operation binding the contract event 0xbfa5afa9ccfabe79b0b21f2a9d778cf773b963e00307dd68c96abc3a30c5d3be.
//
// Solidity: event TransferCredit(address indexed from, address indexed to, uint256 value)
func (_Hkfinance *HkfinanceFilterer) FilterTransferCredit(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HkfinanceTransferCreditIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "TransferCredit", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceTransferCreditIterator{contract: _Hkfinance.contract, event: "TransferCredit", logs: logs, sub: sub}, nil
}

// WatchTransferCredit is a free log subscription operation binding the contract event 0xbfa5afa9ccfabe79b0b21f2a9d778cf773b963e00307dd68c96abc3a30c5d3be.
//
// Solidity: event TransferCredit(address indexed from, address indexed to, uint256 value)
func (_Hkfinance *HkfinanceFilterer) WatchTransferCredit(opts *bind.WatchOpts, sink chan<- *HkfinanceTransferCredit, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "TransferCredit", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceTransferCredit)
				if err := _Hkfinance.contract.UnpackLog(event, "TransferCredit", log); err != nil {
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

// ParseTransferCredit is a log parse operation binding the contract event 0xbfa5afa9ccfabe79b0b21f2a9d778cf773b963e00307dd68c96abc3a30c5d3be.
//
// Solidity: event TransferCredit(address indexed from, address indexed to, uint256 value)
func (_Hkfinance *HkfinanceFilterer) ParseTransferCredit(log types.Log) (*HkfinanceTransferCredit, error) {
	event := new(HkfinanceTransferCredit)
	if err := _Hkfinance.contract.UnpackLog(event, "TransferCredit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HkfinanceWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the Hkfinance contract.
type HkfinanceWithdrawTokenIterator struct {
	Event *HkfinanceWithdrawToken // Event containing the contract specifics and raw log

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
func (it *HkfinanceWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HkfinanceWithdrawToken)
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
		it.Event = new(HkfinanceWithdrawToken)
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
func (it *HkfinanceWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HkfinanceWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HkfinanceWithdrawToken represents a WithdrawToken event raised by the Hkfinance contract.
type HkfinanceWithdrawToken struct {
	Member common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a9.
//
// Solidity: event WithdrawToken(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) FilterWithdrawToken(opts *bind.FilterOpts, member []common.Address) (*HkfinanceWithdrawTokenIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.FilterLogs(opts, "WithdrawToken", memberRule)
	if err != nil {
		return nil, err
	}
	return &HkfinanceWithdrawTokenIterator{contract: _Hkfinance.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a9.
//
// Solidity: event WithdrawToken(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *HkfinanceWithdrawToken, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _Hkfinance.contract.WatchLogs(opts, "WithdrawToken", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HkfinanceWithdrawToken)
				if err := _Hkfinance.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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

// ParseWithdrawToken is a log parse operation binding the contract event 0x992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a9.
//
// Solidity: event WithdrawToken(address indexed member, uint256 value)
func (_Hkfinance *HkfinanceFilterer) ParseWithdrawToken(log types.Log) (*HkfinanceWithdrawToken, error) {
	event := new(HkfinanceWithdrawToken)
	if err := _Hkfinance.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
