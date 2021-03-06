// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CalletclotteryABI is the input ABI used to generate the binding from.
const CalletclotteryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getEndowmentBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"delegatecallByFun\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"testbuy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"callByFun\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// CalletclotteryBin is the compiled bytecode used for deploying new contracts.
const CalletclotteryBin = `0x608060405234801561001057600080fd5b506102da806100206000396000f3006080604052600436106100535763ffffffff60e060020a6000350416635acce36b811461005557806366d382031461007c5780638020d1461461009d578063934f203f146100d2578063d11f17ad146100e7575b005b34801561006157600080fd5b5061006a610108565b60408051918252519081900360200190f35b34801561008857600080fd5b50610053600160a060020a036004351661010d565b3480156100a957600080fd5b506100be600160a060020a036004351661013c565b604080519115158252519081900360200190f35b3480156100de57600080fd5b506100536101b7565b3480156100f357600080fd5b506100be600160a060020a0360043516610232565b303190565b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b604080517f6275792875696e743829000000000000000000000000000000000000000000008152815190819003600a01812063ffffffff60e060020a8083049182160283526001600484015292516000939192600160a060020a0386169291602480830192879291908290030181865af49695505050505050565b60008054604080517f14107f3c000000000000000000000000000000000000000000000000000000008152600160048201529051600160a060020a03909216926314107f3c9260248084019382900301818387803b15801561021857600080fd5b505af115801561022c573d6000803e3d6000fd5b50505050565b604080517f6275792875696e743829000000000000000000000000000000000000000000008152815190819003600a01812063ffffffff60e060020a8083049182160283526001600484015292516000939192600160a060020a038616929160248083019287929190829003018183875af196955050505050505600a165627a7a723058201205880715f1ba50086f035c9ae76b00243fb2f714f6889108ad4e7a0fcd8b400029`

// DeployCalletclottery deploys a new Ethereum contract, binding an instance of Calletclottery to it.
func DeployCalletclottery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Calletclottery, error) {
	parsed, err := abi.JSON(strings.NewReader(CalletclotteryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CalletclotteryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Calletclottery{CalletclotteryCaller: CalletclotteryCaller{contract: contract}, CalletclotteryTransactor: CalletclotteryTransactor{contract: contract}, CalletclotteryFilterer: CalletclotteryFilterer{contract: contract}}, nil
}

// Calletclottery is an auto generated Go binding around an Ethereum contract.
type Calletclottery struct {
	CalletclotteryCaller     // Read-only binding to the contract
	CalletclotteryTransactor // Write-only binding to the contract
	CalletclotteryFilterer   // Log filterer for contract events
}

// CalletclotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalletclotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalletclotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalletclotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalletclotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalletclotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalletclotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalletclotterySession struct {
	Contract     *Calletclottery   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalletclotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalletclotteryCallerSession struct {
	Contract *CalletclotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CalletclotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalletclotteryTransactorSession struct {
	Contract     *CalletclotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CalletclotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalletclotteryRaw struct {
	Contract *Calletclottery // Generic contract binding to access the raw methods on
}

// CalletclotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalletclotteryCallerRaw struct {
	Contract *CalletclotteryCaller // Generic read-only contract binding to access the raw methods on
}

// CalletclotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalletclotteryTransactorRaw struct {
	Contract *CalletclotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalletclottery creates a new instance of Calletclottery, bound to a specific deployed contract.
func NewCalletclottery(address common.Address, backend bind.ContractBackend) (*Calletclottery, error) {
	contract, err := bindCalletclottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Calletclottery{CalletclotteryCaller: CalletclotteryCaller{contract: contract}, CalletclotteryTransactor: CalletclotteryTransactor{contract: contract}, CalletclotteryFilterer: CalletclotteryFilterer{contract: contract}}, nil
}

// NewCalletclotteryCaller creates a new read-only instance of Calletclottery, bound to a specific deployed contract.
func NewCalletclotteryCaller(address common.Address, caller bind.ContractCaller) (*CalletclotteryCaller, error) {
	contract, err := bindCalletclottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalletclotteryCaller{contract: contract}, nil
}

// NewCalletclotteryTransactor creates a new write-only instance of Calletclottery, bound to a specific deployed contract.
func NewCalletclotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*CalletclotteryTransactor, error) {
	contract, err := bindCalletclottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalletclotteryTransactor{contract: contract}, nil
}

// NewCalletclotteryFilterer creates a new log filterer instance of Calletclottery, bound to a specific deployed contract.
func NewCalletclotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*CalletclotteryFilterer, error) {
	contract, err := bindCalletclottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalletclotteryFilterer{contract: contract}, nil
}

// bindCalletclottery binds a generic wrapper to an already deployed contract.
func bindCalletclottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalletclotteryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calletclottery *CalletclotteryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Calletclottery.Contract.CalletclotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calletclottery *CalletclotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calletclottery.Contract.CalletclotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calletclottery *CalletclotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calletclottery.Contract.CalletclotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calletclottery *CalletclotteryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Calletclottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calletclottery *CalletclotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calletclottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calletclottery *CalletclotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calletclottery.Contract.contract.Transact(opts, method, params...)
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_Calletclottery *CalletclotteryCaller) GetEndowmentBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Calletclottery.contract.Call(opts, out, "getEndowmentBalance")
	return *ret0, err
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_Calletclottery *CalletclotterySession) GetEndowmentBalance() (*big.Int, error) {
	return _Calletclottery.Contract.GetEndowmentBalance(&_Calletclottery.CallOpts)
}

// GetEndowmentBalance is a free data retrieval call binding the contract method 0x5acce36b.
//
// Solidity: function getEndowmentBalance() constant returns(uint256)
func (_Calletclottery *CalletclotteryCallerSession) GetEndowmentBalance() (*big.Int, error) {
	return _Calletclottery.Contract.GetEndowmentBalance(&_Calletclottery.CallOpts)
}

// CallByFun is a paid mutator transaction binding the contract method 0xd11f17ad.
//
// Solidity: function callByFun(addr address) returns(bool)
func (_Calletclottery *CalletclotteryTransactor) CallByFun(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.contract.Transact(opts, "callByFun", addr)
}

// CallByFun is a paid mutator transaction binding the contract method 0xd11f17ad.
//
// Solidity: function callByFun(addr address) returns(bool)
func (_Calletclottery *CalletclotterySession) CallByFun(addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.Contract.CallByFun(&_Calletclottery.TransactOpts, addr)
}

// CallByFun is a paid mutator transaction binding the contract method 0xd11f17ad.
//
// Solidity: function callByFun(addr address) returns(bool)
func (_Calletclottery *CalletclotteryTransactorSession) CallByFun(addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.Contract.CallByFun(&_Calletclottery.TransactOpts, addr)
}

// DelegatecallByFun is a paid mutator transaction binding the contract method 0x8020d146.
//
// Solidity: function delegatecallByFun(addr address) returns(bool)
func (_Calletclottery *CalletclotteryTransactor) DelegatecallByFun(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.contract.Transact(opts, "delegatecallByFun", addr)
}

// DelegatecallByFun is a paid mutator transaction binding the contract method 0x8020d146.
//
// Solidity: function delegatecallByFun(addr address) returns(bool)
func (_Calletclottery *CalletclotterySession) DelegatecallByFun(addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.Contract.DelegatecallByFun(&_Calletclottery.TransactOpts, addr)
}

// DelegatecallByFun is a paid mutator transaction binding the contract method 0x8020d146.
//
// Solidity: function delegatecallByFun(addr address) returns(bool)
func (_Calletclottery *CalletclotteryTransactorSession) DelegatecallByFun(addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.Contract.DelegatecallByFun(&_Calletclottery.TransactOpts, addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(addr address) returns()
func (_Calletclottery *CalletclotteryTransactor) Setup(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.contract.Transact(opts, "setup", addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(addr address) returns()
func (_Calletclottery *CalletclotterySession) Setup(addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.Contract.Setup(&_Calletclottery.TransactOpts, addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(addr address) returns()
func (_Calletclottery *CalletclotteryTransactorSession) Setup(addr common.Address) (*types.Transaction, error) {
	return _Calletclottery.Contract.Setup(&_Calletclottery.TransactOpts, addr)
}

// Testbuy is a paid mutator transaction binding the contract method 0x934f203f.
//
// Solidity: function testbuy() returns()
func (_Calletclottery *CalletclotteryTransactor) Testbuy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calletclottery.contract.Transact(opts, "testbuy")
}

// Testbuy is a paid mutator transaction binding the contract method 0x934f203f.
//
// Solidity: function testbuy() returns()
func (_Calletclottery *CalletclotterySession) Testbuy() (*types.Transaction, error) {
	return _Calletclottery.Contract.Testbuy(&_Calletclottery.TransactOpts)
}

// Testbuy is a paid mutator transaction binding the contract method 0x934f203f.
//
// Solidity: function testbuy() returns()
func (_Calletclottery *CalletclotteryTransactorSession) Testbuy() (*types.Transaction, error) {
	return _Calletclottery.Contract.Testbuy(&_Calletclottery.TransactOpts)
}

// EtclotteryInterfaceABI is the input ABI used to generate the binding from.
const EtclotteryInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"playerWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"_team\",\"type\":\"uint8\"}],\"name\":\"reinvest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EtclotteryInterfaceBin is the compiled bytecode used for deploying new contracts.
const EtclotteryInterfaceBin = `0x`

// DeployEtclotteryInterface deploys a new Ethereum contract, binding an instance of EtclotteryInterface to it.
func DeployEtclotteryInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EtclotteryInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(EtclotteryInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtclotteryInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtclotteryInterface{EtclotteryInterfaceCaller: EtclotteryInterfaceCaller{contract: contract}, EtclotteryInterfaceTransactor: EtclotteryInterfaceTransactor{contract: contract}, EtclotteryInterfaceFilterer: EtclotteryInterfaceFilterer{contract: contract}}, nil
}

// EtclotteryInterface is an auto generated Go binding around an Ethereum contract.
type EtclotteryInterface struct {
	EtclotteryInterfaceCaller     // Read-only binding to the contract
	EtclotteryInterfaceTransactor // Write-only binding to the contract
	EtclotteryInterfaceFilterer   // Log filterer for contract events
}

// EtclotteryInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtclotteryInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtclotteryInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtclotteryInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtclotteryInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtclotteryInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtclotteryInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtclotteryInterfaceSession struct {
	Contract     *EtclotteryInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EtclotteryInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtclotteryInterfaceCallerSession struct {
	Contract *EtclotteryInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// EtclotteryInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtclotteryInterfaceTransactorSession struct {
	Contract     *EtclotteryInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// EtclotteryInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtclotteryInterfaceRaw struct {
	Contract *EtclotteryInterface // Generic contract binding to access the raw methods on
}

// EtclotteryInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtclotteryInterfaceCallerRaw struct {
	Contract *EtclotteryInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// EtclotteryInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtclotteryInterfaceTransactorRaw struct {
	Contract *EtclotteryInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtclotteryInterface creates a new instance of EtclotteryInterface, bound to a specific deployed contract.
func NewEtclotteryInterface(address common.Address, backend bind.ContractBackend) (*EtclotteryInterface, error) {
	contract, err := bindEtclotteryInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtclotteryInterface{EtclotteryInterfaceCaller: EtclotteryInterfaceCaller{contract: contract}, EtclotteryInterfaceTransactor: EtclotteryInterfaceTransactor{contract: contract}, EtclotteryInterfaceFilterer: EtclotteryInterfaceFilterer{contract: contract}}, nil
}

// NewEtclotteryInterfaceCaller creates a new read-only instance of EtclotteryInterface, bound to a specific deployed contract.
func NewEtclotteryInterfaceCaller(address common.Address, caller bind.ContractCaller) (*EtclotteryInterfaceCaller, error) {
	contract, err := bindEtclotteryInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtclotteryInterfaceCaller{contract: contract}, nil
}

// NewEtclotteryInterfaceTransactor creates a new write-only instance of EtclotteryInterface, bound to a specific deployed contract.
func NewEtclotteryInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*EtclotteryInterfaceTransactor, error) {
	contract, err := bindEtclotteryInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtclotteryInterfaceTransactor{contract: contract}, nil
}

// NewEtclotteryInterfaceFilterer creates a new log filterer instance of EtclotteryInterface, bound to a specific deployed contract.
func NewEtclotteryInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*EtclotteryInterfaceFilterer, error) {
	contract, err := bindEtclotteryInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtclotteryInterfaceFilterer{contract: contract}, nil
}

// bindEtclotteryInterface binds a generic wrapper to an already deployed contract.
func bindEtclotteryInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtclotteryInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtclotteryInterface *EtclotteryInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtclotteryInterface.Contract.EtclotteryInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtclotteryInterface *EtclotteryInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.EtclotteryInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtclotteryInterface *EtclotteryInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.EtclotteryInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtclotteryInterface *EtclotteryInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtclotteryInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtclotteryInterface *EtclotteryInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtclotteryInterface *EtclotteryInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.contract.Transact(opts, method, params...)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_EtclotteryInterface *EtclotteryInterfaceCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtclotteryInterface.contract.Call(opts, out, "getFee")
	return *ret0, err
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_EtclotteryInterface *EtclotteryInterfaceSession) GetFee() (*big.Int, error) {
	return _EtclotteryInterface.Contract.GetFee(&_EtclotteryInterface.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() constant returns(uint256)
func (_EtclotteryInterface *EtclotteryInterfaceCallerSession) GetFee() (*big.Int, error) {
	return _EtclotteryInterface.Contract.GetFee(&_EtclotteryInterface.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactor) Buy(opts *bind.TransactOpts, _team uint8) (*types.Transaction, error) {
	return _EtclotteryInterface.contract.Transact(opts, "buy", _team)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_EtclotteryInterface *EtclotteryInterfaceSession) Buy(_team uint8) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.Buy(&_EtclotteryInterface.TransactOpts, _team)
}

// Buy is a paid mutator transaction binding the contract method 0x14107f3c.
//
// Solidity: function buy(_team uint8) returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactorSession) Buy(_team uint8) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.Buy(&_EtclotteryInterface.TransactOpts, _team)
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactor) GetBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtclotteryInterface.contract.Transact(opts, "getBlock")
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_EtclotteryInterface *EtclotteryInterfaceSession) GetBlock() (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.GetBlock(&_EtclotteryInterface.TransactOpts)
}

// GetBlock is a paid mutator transaction binding the contract method 0x2e97766d.
//
// Solidity: function getBlock() returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactorSession) GetBlock() (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.GetBlock(&_EtclotteryInterface.TransactOpts)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactor) PlayerWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EtclotteryInterface.contract.Transact(opts, "playerWithdraw", amount)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_EtclotteryInterface *EtclotteryInterfaceSession) PlayerWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.PlayerWithdraw(&_EtclotteryInterface.TransactOpts, amount)
}

// PlayerWithdraw is a paid mutator transaction binding the contract method 0x5629365b.
//
// Solidity: function playerWithdraw(amount uint256) returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactorSession) PlayerWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.PlayerWithdraw(&_EtclotteryInterface.TransactOpts, amount)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactor) Reinvest(opts *bind.TransactOpts, amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _EtclotteryInterface.contract.Transact(opts, "reinvest", amount, _team)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_EtclotteryInterface *EtclotteryInterfaceSession) Reinvest(amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.Reinvest(&_EtclotteryInterface.TransactOpts, amount, _team)
}

// Reinvest is a paid mutator transaction binding the contract method 0xf0a3843d.
//
// Solidity: function reinvest(amount uint256, _team uint8) returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactorSession) Reinvest(amount *big.Int, _team uint8) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.Reinvest(&_EtclotteryInterface.TransactOpts, amount, _team)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactor) WithdrawFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EtclotteryInterface.contract.Transact(opts, "withdrawFee", amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_EtclotteryInterface *EtclotteryInterfaceSession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.WithdrawFee(&_EtclotteryInterface.TransactOpts, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xbe357616.
//
// Solidity: function withdrawFee(amount uint256) returns()
func (_EtclotteryInterface *EtclotteryInterfaceTransactorSession) WithdrawFee(amount *big.Int) (*types.Transaction, error) {
	return _EtclotteryInterface.Contract.WithdrawFee(&_EtclotteryInterface.TransactOpts, amount)
}
