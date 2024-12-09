// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
	_ = abi.ConvertType
)

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isRegistered\",\"type\":\"bool\"}],\"name\":\"IsRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610614380380610614833981810160405281019061003191906100d4565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f80fd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b6105088061010c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c8063421b2d8b1461004e578063985751881461006a578063c3c5a54714610086578063f851a440146100b6575b5f80fd5b610068600480360381019061006391906103ad565b6100d4565b005b610084600480360381019061007f91906103ad565b6101f2565b005b6100a0600480360381019061009b91906103ad565b61030f565b6040516100ad91906103f2565b60405180910390f35b6100be61032c565b6040516100cb919061041a565b60405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610161576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101589061048d565b60405180910390fd5b6001805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507fd5c534ef2c1fdfacb8a8e6b8b54d5c54e81f58fc33cdd6317244764aee5970348160016040516101e79291906104ab565b60405180910390a150565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461027f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102769061048d565b60405180910390fd5b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507fd5c534ef2c1fdfacb8a8e6b8b54d5c54e81f58fc33cdd6317244764aee597034815f6040516103049291906104ab565b60405180910390a150565b6001602052805f5260405f205f915054906101000a900460ff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61037c82610353565b9050919050565b61038c81610372565b8114610396575f80fd5b50565b5f813590506103a781610383565b92915050565b5f602082840312156103c2576103c161034f565b5b5f6103cf84828501610399565b91505092915050565b5f8115159050919050565b6103ec816103d8565b82525050565b5f6020820190506104055f8301846103e3565b92915050565b61041481610372565b82525050565b5f60208201905061042d5f83018461040b565b92915050565b5f82825260208201905092915050565b7f496e76616c69642061646d696e000000000000000000000000000000000000005f82015250565b5f610477600d83610433565b915061048282610443565b602082019050919050565b5f6020820190508181035f8301526104a48161046b565b9050919050565b5f6040820190506104be5f83018561040b565b6104cb60208301846103e3565b939250505056fea2646970667358221220243bb2555bbfa3895a589d39ba355e980b7544225c228703ca4f424ef52166e364736f6c634300081a0033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Store *StoreCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Store *StoreSession) Admin() (common.Address, error) {
	return _Store.Contract.Admin(&_Store.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Store *StoreCallerSession) Admin() (common.Address, error) {
	return _Store.Contract.Admin(&_Store.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_Store *StoreCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_Store *StoreSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _Store.Contract.IsRegistered(&_Store.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_Store *StoreCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _Store.Contract.IsRegistered(&_Store.CallOpts, arg0)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address _user) returns()
func (_Store *StoreTransactor) AddUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addUser", _user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address _user) returns()
func (_Store *StoreSession) AddUser(_user common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddUser(&_Store.TransactOpts, _user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address _user) returns()
func (_Store *StoreTransactorSession) AddUser(_user common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddUser(&_Store.TransactOpts, _user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address _user) returns()
func (_Store *StoreTransactor) RemoveUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "removeUser", _user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address _user) returns()
func (_Store *StoreSession) RemoveUser(_user common.Address) (*types.Transaction, error) {
	return _Store.Contract.RemoveUser(&_Store.TransactOpts, _user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address _user) returns()
func (_Store *StoreTransactorSession) RemoveUser(_user common.Address) (*types.Transaction, error) {
	return _Store.Contract.RemoveUser(&_Store.TransactOpts, _user)
}

// StoreIsRegisteredIterator is returned from FilterIsRegistered and is used to iterate over the raw logs and unpacked data for IsRegistered events raised by the Store contract.
type StoreIsRegisteredIterator struct {
	Event *StoreIsRegistered // Event containing the contract specifics and raw log

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
func (it *StoreIsRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreIsRegistered)
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
		it.Event = new(StoreIsRegistered)
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
func (it *StoreIsRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreIsRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreIsRegistered represents a IsRegistered event raised by the Store contract.
type StoreIsRegistered struct {
	NewUser      common.Address
	IsRegistered bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIsRegistered is a free log retrieval operation binding the contract event 0xd5c534ef2c1fdfacb8a8e6b8b54d5c54e81f58fc33cdd6317244764aee597034.
//
// Solidity: event IsRegistered(address _newUser, bool _isRegistered)
func (_Store *StoreFilterer) FilterIsRegistered(opts *bind.FilterOpts) (*StoreIsRegisteredIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "IsRegistered")
	if err != nil {
		return nil, err
	}
	return &StoreIsRegisteredIterator{contract: _Store.contract, event: "IsRegistered", logs: logs, sub: sub}, nil
}

// WatchIsRegistered is a free log subscription operation binding the contract event 0xd5c534ef2c1fdfacb8a8e6b8b54d5c54e81f58fc33cdd6317244764aee597034.
//
// Solidity: event IsRegistered(address _newUser, bool _isRegistered)
func (_Store *StoreFilterer) WatchIsRegistered(opts *bind.WatchOpts, sink chan<- *StoreIsRegistered) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "IsRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreIsRegistered)
				if err := _Store.contract.UnpackLog(event, "IsRegistered", log); err != nil {
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

// ParseIsRegistered is a log parse operation binding the contract event 0xd5c534ef2c1fdfacb8a8e6b8b54d5c54e81f58fc33cdd6317244764aee597034.
//
// Solidity: event IsRegistered(address _newUser, bool _isRegistered)
func (_Store *StoreFilterer) ParseIsRegistered(log types.Log) (*StoreIsRegistered, error) {
	event := new(StoreIsRegistered)
	if err := _Store.contract.UnpackLog(event, "IsRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
