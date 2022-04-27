// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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

// IDiamondCutFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// IDiamondCutMetaData contains all meta data concerning the IDiamondCut contract.
var IDiamondCutMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1f931c1c": "diamondCut((address,uint8,bytes4[])[],address,bytes)",
	},
}

// IDiamondCutABI is the input ABI used to generate the binding from.
// Deprecated: Use IDiamondCutMetaData.ABI instead.
var IDiamondCutABI = IDiamondCutMetaData.ABI

// Deprecated: Use IDiamondCutMetaData.Sigs instead.
// IDiamondCutFuncSigs maps the 4-byte function signature to its string representation.
var IDiamondCutFuncSigs = IDiamondCutMetaData.Sigs

// IDiamondCut is an auto generated Go binding around an Ethereum contract.
type IDiamondCut struct {
	IDiamondCutCaller     // Read-only binding to the contract
	IDiamondCutTransactor // Write-only binding to the contract
	IDiamondCutFilterer   // Log filterer for contract events
}

// IDiamondCutCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDiamondCutCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDiamondCutTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDiamondCutFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDiamondCutSession struct {
	Contract     *IDiamondCut      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDiamondCutCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDiamondCutCallerSession struct {
	Contract *IDiamondCutCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IDiamondCutTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDiamondCutTransactorSession struct {
	Contract     *IDiamondCutTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IDiamondCutRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDiamondCutRaw struct {
	Contract *IDiamondCut // Generic contract binding to access the raw methods on
}

// IDiamondCutCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDiamondCutCallerRaw struct {
	Contract *IDiamondCutCaller // Generic read-only contract binding to access the raw methods on
}

// IDiamondCutTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDiamondCutTransactorRaw struct {
	Contract *IDiamondCutTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDiamondCut creates a new instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCut(address common.Address, backend bind.ContractBackend) (*IDiamondCut, error) {
	contract, err := bindIDiamondCut(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDiamondCut{IDiamondCutCaller: IDiamondCutCaller{contract: contract}, IDiamondCutTransactor: IDiamondCutTransactor{contract: contract}, IDiamondCutFilterer: IDiamondCutFilterer{contract: contract}}, nil
}

// NewIDiamondCutCaller creates a new read-only instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutCaller(address common.Address, caller bind.ContractCaller) (*IDiamondCutCaller, error) {
	contract, err := bindIDiamondCut(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutCaller{contract: contract}, nil
}

// NewIDiamondCutTransactor creates a new write-only instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutTransactor(address common.Address, transactor bind.ContractTransactor) (*IDiamondCutTransactor, error) {
	contract, err := bindIDiamondCut(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutTransactor{contract: contract}, nil
}

// NewIDiamondCutFilterer creates a new log filterer instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutFilterer(address common.Address, filterer bind.ContractFilterer) (*IDiamondCutFilterer, error) {
	contract, err := bindIDiamondCut(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutFilterer{contract: contract}, nil
}

// bindIDiamondCut binds a generic wrapper to an already deployed contract.
func bindIDiamondCut(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDiamondCutABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondCut *IDiamondCutRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondCut.Contract.IDiamondCutCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondCut *IDiamondCutRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondCut.Contract.IDiamondCutTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondCut *IDiamondCutRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondCut.Contract.IDiamondCutTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondCut *IDiamondCutCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondCut.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondCut *IDiamondCutTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondCut.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondCut *IDiamondCutTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondCut.Contract.contract.Transact(opts, method, params...)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.Contract.DiamondCut(&_IDiamondCut.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutTransactorSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.Contract.DiamondCut(&_IDiamondCut.TransactOpts, _diamondCut, _init, _calldata)
}

// IDiamondCutDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the IDiamondCut contract.
type IDiamondCutDiamondCutIterator struct {
	Event *IDiamondCutDiamondCut // Event containing the contract specifics and raw log

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
func (it *IDiamondCutDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDiamondCutDiamondCut)
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
		it.Event = new(IDiamondCutDiamondCut)
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
func (it *IDiamondCutDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDiamondCutDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDiamondCutDiamondCut represents a DiamondCut event raised by the IDiamondCut contract.
type IDiamondCutDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_IDiamondCut *IDiamondCutFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*IDiamondCutDiamondCutIterator, error) {

	logs, sub, err := _IDiamondCut.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &IDiamondCutDiamondCutIterator{contract: _IDiamondCut.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_IDiamondCut *IDiamondCutFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *IDiamondCutDiamondCut) (event.Subscription, error) {

	logs, sub, err := _IDiamondCut.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDiamondCutDiamondCut)
				if err := _IDiamondCut.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_IDiamondCut *IDiamondCutFilterer) ParseDiamondCut(log types.Log) (*IDiamondCutDiamondCut, error) {
	event := new(IDiamondCutDiamondCut)
	if err := _IDiamondCut.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC173MetaData contains all meta data concerning the IERC173 contract.
var IERC173MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// IERC173ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC173MetaData.ABI instead.
var IERC173ABI = IERC173MetaData.ABI

// Deprecated: Use IERC173MetaData.Sigs instead.
// IERC173FuncSigs maps the 4-byte function signature to its string representation.
var IERC173FuncSigs = IERC173MetaData.Sigs

// IERC173 is an auto generated Go binding around an Ethereum contract.
type IERC173 struct {
	IERC173Caller     // Read-only binding to the contract
	IERC173Transactor // Write-only binding to the contract
	IERC173Filterer   // Log filterer for contract events
}

// IERC173Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC173Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC173Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC173Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC173Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC173Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC173Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC173Session struct {
	Contract     *IERC173          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC173CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC173CallerSession struct {
	Contract *IERC173Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC173TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC173TransactorSession struct {
	Contract     *IERC173Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC173Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC173Raw struct {
	Contract *IERC173 // Generic contract binding to access the raw methods on
}

// IERC173CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC173CallerRaw struct {
	Contract *IERC173Caller // Generic read-only contract binding to access the raw methods on
}

// IERC173TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC173TransactorRaw struct {
	Contract *IERC173Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC173 creates a new instance of IERC173, bound to a specific deployed contract.
func NewIERC173(address common.Address, backend bind.ContractBackend) (*IERC173, error) {
	contract, err := bindIERC173(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC173{IERC173Caller: IERC173Caller{contract: contract}, IERC173Transactor: IERC173Transactor{contract: contract}, IERC173Filterer: IERC173Filterer{contract: contract}}, nil
}

// NewIERC173Caller creates a new read-only instance of IERC173, bound to a specific deployed contract.
func NewIERC173Caller(address common.Address, caller bind.ContractCaller) (*IERC173Caller, error) {
	contract, err := bindIERC173(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC173Caller{contract: contract}, nil
}

// NewIERC173Transactor creates a new write-only instance of IERC173, bound to a specific deployed contract.
func NewIERC173Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC173Transactor, error) {
	contract, err := bindIERC173(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC173Transactor{contract: contract}, nil
}

// NewIERC173Filterer creates a new log filterer instance of IERC173, bound to a specific deployed contract.
func NewIERC173Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC173Filterer, error) {
	contract, err := bindIERC173(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC173Filterer{contract: contract}, nil
}

// bindIERC173 binds a generic wrapper to an already deployed contract.
func bindIERC173(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC173ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC173 *IERC173Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC173.Contract.IERC173Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC173 *IERC173Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC173.Contract.IERC173Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC173 *IERC173Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC173.Contract.IERC173Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC173 *IERC173CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC173.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC173 *IERC173TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC173.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC173 *IERC173TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC173.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_IERC173 *IERC173Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IERC173.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_IERC173 *IERC173Session) Owner() (common.Address, error) {
	return _IERC173.Contract.Owner(&_IERC173.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_IERC173 *IERC173CallerSession) Owner() (common.Address, error) {
	return _IERC173.Contract.Owner(&_IERC173.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_IERC173 *IERC173Transactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _IERC173.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_IERC173 *IERC173Session) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _IERC173.Contract.TransferOwnership(&_IERC173.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_IERC173 *IERC173TransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _IERC173.Contract.TransferOwnership(&_IERC173.TransactOpts, _newOwner)
}

// IERC173OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IERC173 contract.
type IERC173OwnershipTransferredIterator struct {
	Event *IERC173OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IERC173OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC173OwnershipTransferred)
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
		it.Event = new(IERC173OwnershipTransferred)
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
func (it *IERC173OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC173OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC173OwnershipTransferred represents a OwnershipTransferred event raised by the IERC173 contract.
type IERC173OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IERC173 *IERC173Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IERC173OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IERC173.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IERC173OwnershipTransferredIterator{contract: _IERC173.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IERC173 *IERC173Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IERC173OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IERC173.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC173OwnershipTransferred)
				if err := _IERC173.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IERC173 *IERC173Filterer) ParseOwnershipTransferred(log types.Log) (*IERC173OwnershipTransferred, error) {
	event := new(IERC173OwnershipTransferred)
	if err := _IERC173.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibDiamondMetaData contains all meta data concerning the LibDiamond contract.
var LibDiamondMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202f3455883fbb0e90186d7efa333813a455bef6a925e19a0ba98f11052f58cb5864736f6c634300080d0033",
}

// LibDiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use LibDiamondMetaData.ABI instead.
var LibDiamondABI = LibDiamondMetaData.ABI

// LibDiamondBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibDiamondMetaData.Bin instead.
var LibDiamondBin = LibDiamondMetaData.Bin

// DeployLibDiamond deploys a new Ethereum contract, binding an instance of LibDiamond to it.
func DeployLibDiamond(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibDiamond, error) {
	parsed, err := LibDiamondMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibDiamondBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibDiamond{LibDiamondCaller: LibDiamondCaller{contract: contract}, LibDiamondTransactor: LibDiamondTransactor{contract: contract}, LibDiamondFilterer: LibDiamondFilterer{contract: contract}}, nil
}

// LibDiamond is an auto generated Go binding around an Ethereum contract.
type LibDiamond struct {
	LibDiamondCaller     // Read-only binding to the contract
	LibDiamondTransactor // Write-only binding to the contract
	LibDiamondFilterer   // Log filterer for contract events
}

// LibDiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibDiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibDiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibDiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibDiamondSession struct {
	Contract     *LibDiamond       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibDiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibDiamondCallerSession struct {
	Contract *LibDiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LibDiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibDiamondTransactorSession struct {
	Contract     *LibDiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LibDiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibDiamondRaw struct {
	Contract *LibDiamond // Generic contract binding to access the raw methods on
}

// LibDiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibDiamondCallerRaw struct {
	Contract *LibDiamondCaller // Generic read-only contract binding to access the raw methods on
}

// LibDiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibDiamondTransactorRaw struct {
	Contract *LibDiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibDiamond creates a new instance of LibDiamond, bound to a specific deployed contract.
func NewLibDiamond(address common.Address, backend bind.ContractBackend) (*LibDiamond, error) {
	contract, err := bindLibDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibDiamond{LibDiamondCaller: LibDiamondCaller{contract: contract}, LibDiamondTransactor: LibDiamondTransactor{contract: contract}, LibDiamondFilterer: LibDiamondFilterer{contract: contract}}, nil
}

// NewLibDiamondCaller creates a new read-only instance of LibDiamond, bound to a specific deployed contract.
func NewLibDiamondCaller(address common.Address, caller bind.ContractCaller) (*LibDiamondCaller, error) {
	contract, err := bindLibDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibDiamondCaller{contract: contract}, nil
}

// NewLibDiamondTransactor creates a new write-only instance of LibDiamond, bound to a specific deployed contract.
func NewLibDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*LibDiamondTransactor, error) {
	contract, err := bindLibDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibDiamondTransactor{contract: contract}, nil
}

// NewLibDiamondFilterer creates a new log filterer instance of LibDiamond, bound to a specific deployed contract.
func NewLibDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*LibDiamondFilterer, error) {
	contract, err := bindLibDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibDiamondFilterer{contract: contract}, nil
}

// bindLibDiamond binds a generic wrapper to an already deployed contract.
func bindLibDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibDiamondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDiamond *LibDiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDiamond.Contract.LibDiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDiamond *LibDiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDiamond.Contract.LibDiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDiamond *LibDiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDiamond.Contract.LibDiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDiamond *LibDiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDiamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDiamond *LibDiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDiamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDiamond *LibDiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDiamond.Contract.contract.Transact(opts, method, params...)
}

// LibDiamondDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the LibDiamond contract.
type LibDiamondDiamondCutIterator struct {
	Event *LibDiamondDiamondCut // Event containing the contract specifics and raw log

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
func (it *LibDiamondDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibDiamondDiamondCut)
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
		it.Event = new(LibDiamondDiamondCut)
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
func (it *LibDiamondDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibDiamondDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibDiamondDiamondCut represents a DiamondCut event raised by the LibDiamond contract.
type LibDiamondDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_LibDiamond *LibDiamondFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*LibDiamondDiamondCutIterator, error) {

	logs, sub, err := _LibDiamond.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &LibDiamondDiamondCutIterator{contract: _LibDiamond.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_LibDiamond *LibDiamondFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *LibDiamondDiamondCut) (event.Subscription, error) {

	logs, sub, err := _LibDiamond.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibDiamondDiamondCut)
				if err := _LibDiamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_LibDiamond *LibDiamondFilterer) ParseDiamondCut(log types.Log) (*LibDiamondDiamondCut, error) {
	event := new(LibDiamondDiamondCut)
	if err := _LibDiamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibDiamondOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LibDiamond contract.
type LibDiamondOwnershipTransferredIterator struct {
	Event *LibDiamondOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LibDiamondOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibDiamondOwnershipTransferred)
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
		it.Event = new(LibDiamondOwnershipTransferred)
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
func (it *LibDiamondOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibDiamondOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibDiamondOwnershipTransferred represents a OwnershipTransferred event raised by the LibDiamond contract.
type LibDiamondOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LibDiamond *LibDiamondFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LibDiamondOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LibDiamond.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LibDiamondOwnershipTransferredIterator{contract: _LibDiamond.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LibDiamond *LibDiamondFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LibDiamondOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LibDiamond.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibDiamondOwnershipTransferred)
				if err := _LibDiamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LibDiamond *LibDiamondFilterer) ParseOwnershipTransferred(log types.Log) (*LibDiamondOwnershipTransferred, error) {
	event := new(LibDiamondOwnershipTransferred)
	if err := _LibDiamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnershipFacetMetaData contains all meta data concerning the OwnershipFacet contract.
var OwnershipFacetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610248806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b1461005f575b600080fd5b610043610074565b6040516001600160a01b03909116815260200160405180910390f35b61007261006d3660046101e2565b6100ac565b005b60006100a77fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c1320546001600160a01b031690565b905090565b6100b46100c0565b6100bd8161014d565b50565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c600401546001600160a01b0316331461014b5760405162461bcd60e51b815260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201526132b960f11b606482015260840160405180910390fd5b565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c132080546001600160a01b031981166001600160a01b038481169182179093556040517fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c939092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6000602082840312156101f457600080fd5b81356001600160a01b038116811461020b57600080fd5b939250505056fea2646970667358221220c45908f9a9a26e834923884dd007df491c52d4ee856cb105fbceadfdf19c601464736f6c634300080d0033",
}

// OwnershipFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnershipFacetMetaData.ABI instead.
var OwnershipFacetABI = OwnershipFacetMetaData.ABI

// Deprecated: Use OwnershipFacetMetaData.Sigs instead.
// OwnershipFacetFuncSigs maps the 4-byte function signature to its string representation.
var OwnershipFacetFuncSigs = OwnershipFacetMetaData.Sigs

// OwnershipFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnershipFacetMetaData.Bin instead.
var OwnershipFacetBin = OwnershipFacetMetaData.Bin

// DeployOwnershipFacet deploys a new Ethereum contract, binding an instance of OwnershipFacet to it.
func DeployOwnershipFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OwnershipFacet, error) {
	parsed, err := OwnershipFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnershipFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OwnershipFacet{OwnershipFacetCaller: OwnershipFacetCaller{contract: contract}, OwnershipFacetTransactor: OwnershipFacetTransactor{contract: contract}, OwnershipFacetFilterer: OwnershipFacetFilterer{contract: contract}}, nil
}

// OwnershipFacet is an auto generated Go binding around an Ethereum contract.
type OwnershipFacet struct {
	OwnershipFacetCaller     // Read-only binding to the contract
	OwnershipFacetTransactor // Write-only binding to the contract
	OwnershipFacetFilterer   // Log filterer for contract events
}

// OwnershipFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnershipFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnershipFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnershipFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnershipFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnershipFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnershipFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnershipFacetSession struct {
	Contract     *OwnershipFacet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnershipFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnershipFacetCallerSession struct {
	Contract *OwnershipFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OwnershipFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnershipFacetTransactorSession struct {
	Contract     *OwnershipFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OwnershipFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnershipFacetRaw struct {
	Contract *OwnershipFacet // Generic contract binding to access the raw methods on
}

// OwnershipFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnershipFacetCallerRaw struct {
	Contract *OwnershipFacetCaller // Generic read-only contract binding to access the raw methods on
}

// OwnershipFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnershipFacetTransactorRaw struct {
	Contract *OwnershipFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnershipFacet creates a new instance of OwnershipFacet, bound to a specific deployed contract.
func NewOwnershipFacet(address common.Address, backend bind.ContractBackend) (*OwnershipFacet, error) {
	contract, err := bindOwnershipFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacet{OwnershipFacetCaller: OwnershipFacetCaller{contract: contract}, OwnershipFacetTransactor: OwnershipFacetTransactor{contract: contract}, OwnershipFacetFilterer: OwnershipFacetFilterer{contract: contract}}, nil
}

// NewOwnershipFacetCaller creates a new read-only instance of OwnershipFacet, bound to a specific deployed contract.
func NewOwnershipFacetCaller(address common.Address, caller bind.ContractCaller) (*OwnershipFacetCaller, error) {
	contract, err := bindOwnershipFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetCaller{contract: contract}, nil
}

// NewOwnershipFacetTransactor creates a new write-only instance of OwnershipFacet, bound to a specific deployed contract.
func NewOwnershipFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnershipFacetTransactor, error) {
	contract, err := bindOwnershipFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetTransactor{contract: contract}, nil
}

// NewOwnershipFacetFilterer creates a new log filterer instance of OwnershipFacet, bound to a specific deployed contract.
func NewOwnershipFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnershipFacetFilterer, error) {
	contract, err := bindOwnershipFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetFilterer{contract: contract}, nil
}

// bindOwnershipFacet binds a generic wrapper to an already deployed contract.
func bindOwnershipFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnershipFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnershipFacet *OwnershipFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnershipFacet.Contract.OwnershipFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnershipFacet *OwnershipFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.OwnershipFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnershipFacet *OwnershipFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.OwnershipFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnershipFacet *OwnershipFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnershipFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnershipFacet *OwnershipFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnershipFacet *OwnershipFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_OwnershipFacet *OwnershipFacetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnershipFacet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_OwnershipFacet *OwnershipFacetSession) Owner() (common.Address, error) {
	return _OwnershipFacet.Contract.Owner(&_OwnershipFacet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address owner_)
func (_OwnershipFacet *OwnershipFacetCallerSession) Owner() (common.Address, error) {
	return _OwnershipFacet.Contract.Owner(&_OwnershipFacet.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_OwnershipFacet *OwnershipFacetTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _OwnershipFacet.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_OwnershipFacet *OwnershipFacetSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.TransferOwnership(&_OwnershipFacet.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_OwnershipFacet *OwnershipFacetTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _OwnershipFacet.Contract.TransferOwnership(&_OwnershipFacet.TransactOpts, _newOwner)
}

// OwnershipFacetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnershipFacet contract.
type OwnershipFacetOwnershipTransferredIterator struct {
	Event *OwnershipFacetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnershipFacetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnershipFacetOwnershipTransferred)
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
		it.Event = new(OwnershipFacetOwnershipTransferred)
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
func (it *OwnershipFacetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnershipFacetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnershipFacetOwnershipTransferred represents a OwnershipTransferred event raised by the OwnershipFacet contract.
type OwnershipFacetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnershipFacet *OwnershipFacetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnershipFacetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnershipFacet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnershipFacetOwnershipTransferredIterator{contract: _OwnershipFacet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnershipFacet *OwnershipFacetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnershipFacetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnershipFacet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnershipFacetOwnershipTransferred)
				if err := _OwnershipFacet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OwnershipFacet *OwnershipFacetFilterer) ParseOwnershipTransferred(log types.Log) (*OwnershipFacetOwnershipTransferred, error) {
	event := new(OwnershipFacetOwnershipTransferred)
	if err := _OwnershipFacet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
