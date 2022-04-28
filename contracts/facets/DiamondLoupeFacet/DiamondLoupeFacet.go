// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DiamondLoupeFacet

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

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// DiamondLoupeFacetMetaData contains all meta data concerning the DiamondLoupeFacet contract.
var DiamondLoupeFacetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cdffacc6": "facetAddress(bytes4)",
		"52ef6b2c": "facetAddresses()",
		"adfca15e": "facetFunctionSelectors(address)",
		"7a0ed627": "facets()",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061068e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806301ffc9a71461005c57806352ef6b2c146100bd5780637a0ed627146100d2578063adfca15e146100e7578063cdffacc614610107575b600080fd5b6100a861006a366004610469565b6001600160e01b03191660009081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f602052604090205460ff1690565b60405190151581526020015b60405180910390f35b6100c561015f565b6040516100b4919061049a565b6100da6101d2565b6040516100b4919061052c565b6100fa6100f53660046105a9565b61039d565b6040516100b491906105d2565b610147610115366004610469565b6001600160e01b031916600090815260008051602061063983398151915260205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016100b4565b60606000600080516020610639833981519152600281018054604080516020808402820181019092528281529394508301828280156101c757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116101a9575b505050505091505090565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e54606090600080516020610639833981519152908067ffffffffffffffff811115610220576102206105e5565b60405190808252806020026020018201604052801561026657816020015b60408051808201909152600081526060602082015281526020019060019003908161023e5790505b50925060005b8181101561039757600083600201828154811061028b5761028b6105fb565b9060005260206000200160009054906101000a90046001600160a01b03169050808583815181106102be576102be6105fb565b6020908102919091018101516001600160a01b03928316905290821660009081526001860182526040908190208054825181850281018501909352808352919290919083018282801561035d57602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161031f5790505b5050505050858381518110610374576103746105fb565b60200260200101516020018190525050808061038f90610611565b91505061026c565b50505090565b6001600160a01b03811660009081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d60209081526040918290208054835181840281018401909452808452606093600080516020610639833981519152939092919083018282801561045c57602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161041e5790505b5050505050915050919050565b60006020828403121561047b57600080fd5b81356001600160e01b03198116811461049357600080fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156104db5783516001600160a01b0316835292840192918401916001016104b6565b50909695505050505050565b600081518084526020808501945080840160005b838110156105215781516001600160e01b031916875295820195908201906001016104fb565b509495945050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561059b57888303603f19018552815180516001600160a01b03168452870151878401879052610588878501826104e7565b9588019593505090860190600101610553565b509098975050505050505050565b6000602082840312156105bb57600080fd5b81356001600160a01b038116811461049357600080fd5b60208152600061049360208301846104e7565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60006001820161063157634e487b7160e01b600052601160045260246000fd5b506001019056fec8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131ca264697066735822122014d5faee0b69ec2040c775cd6cacdf4f7ff6b82a17a5e57467f00cc94789a26f64736f6c634300080d0033",
}

// DiamondLoupeFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondLoupeFacetMetaData.ABI instead.
var DiamondLoupeFacetABI = DiamondLoupeFacetMetaData.ABI

// Deprecated: Use DiamondLoupeFacetMetaData.Sigs instead.
// DiamondLoupeFacetFuncSigs maps the 4-byte function signature to its string representation.
var DiamondLoupeFacetFuncSigs = DiamondLoupeFacetMetaData.Sigs

// DiamondLoupeFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondLoupeFacetMetaData.Bin instead.
var DiamondLoupeFacetBin = DiamondLoupeFacetMetaData.Bin

// DeployDiamondLoupeFacet deploys a new Ethereum contract, binding an instance of DiamondLoupeFacet to it.
func DeployDiamondLoupeFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondLoupeFacet, error) {
	parsed, err := DiamondLoupeFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondLoupeFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondLoupeFacet{DiamondLoupeFacetCaller: DiamondLoupeFacetCaller{contract: contract}, DiamondLoupeFacetTransactor: DiamondLoupeFacetTransactor{contract: contract}, DiamondLoupeFacetFilterer: DiamondLoupeFacetFilterer{contract: contract}}, nil
}

// DiamondLoupeFacet is an auto generated Go binding around an Ethereum contract.
type DiamondLoupeFacet struct {
	DiamondLoupeFacetCaller     // Read-only binding to the contract
	DiamondLoupeFacetTransactor // Write-only binding to the contract
	DiamondLoupeFacetFilterer   // Log filterer for contract events
}

// DiamondLoupeFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondLoupeFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondLoupeFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondLoupeFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondLoupeFacetSession struct {
	Contract     *DiamondLoupeFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondLoupeFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondLoupeFacetCallerSession struct {
	Contract *DiamondLoupeFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DiamondLoupeFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondLoupeFacetTransactorSession struct {
	Contract     *DiamondLoupeFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DiamondLoupeFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondLoupeFacetRaw struct {
	Contract *DiamondLoupeFacet // Generic contract binding to access the raw methods on
}

// DiamondLoupeFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondLoupeFacetCallerRaw struct {
	Contract *DiamondLoupeFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondLoupeFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondLoupeFacetTransactorRaw struct {
	Contract *DiamondLoupeFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondLoupeFacet creates a new instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacet(address common.Address, backend bind.ContractBackend) (*DiamondLoupeFacet, error) {
	contract, err := bindDiamondLoupeFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacet{DiamondLoupeFacetCaller: DiamondLoupeFacetCaller{contract: contract}, DiamondLoupeFacetTransactor: DiamondLoupeFacetTransactor{contract: contract}, DiamondLoupeFacetFilterer: DiamondLoupeFacetFilterer{contract: contract}}, nil
}

// NewDiamondLoupeFacetCaller creates a new read-only instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetCaller(address common.Address, caller bind.ContractCaller) (*DiamondLoupeFacetCaller, error) {
	contract, err := bindDiamondLoupeFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetCaller{contract: contract}, nil
}

// NewDiamondLoupeFacetTransactor creates a new write-only instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondLoupeFacetTransactor, error) {
	contract, err := bindDiamondLoupeFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetTransactor{contract: contract}, nil
}

// NewDiamondLoupeFacetFilterer creates a new log filterer instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondLoupeFacetFilterer, error) {
	contract, err := bindDiamondLoupeFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetFilterer{contract: contract}, nil
}

// bindDiamondLoupeFacet binds a generic wrapper to an already deployed contract.
func bindDiamondLoupeFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondLoupeFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupeFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupeFacet *DiamondLoupeFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupeFacet *DiamondLoupeFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddress(&_DiamondLoupeFacet.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddress(&_DiamondLoupeFacet.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddresses(&_DiamondLoupeFacet.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddresses(&_DiamondLoupeFacet.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupeFacet.Contract.FacetFunctionSelectors(&_DiamondLoupeFacet.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupeFacet.Contract.FacetFunctionSelectors(&_DiamondLoupeFacet.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupeFacet.Contract.Facets(&_DiamondLoupeFacet.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupeFacet.Contract.Facets(&_DiamondLoupeFacet.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupeFacet.Contract.SupportsInterface(&_DiamondLoupeFacet.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupeFacet.Contract.SupportsInterface(&_DiamondLoupeFacet.CallOpts, _interfaceId)
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

// IDiamondLoupeMetaData contains all meta data concerning the IDiamondLoupe contract.
var IDiamondLoupeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cdffacc6": "facetAddress(bytes4)",
		"52ef6b2c": "facetAddresses()",
		"adfca15e": "facetFunctionSelectors(address)",
		"7a0ed627": "facets()",
	},
}

// IDiamondLoupeABI is the input ABI used to generate the binding from.
// Deprecated: Use IDiamondLoupeMetaData.ABI instead.
var IDiamondLoupeABI = IDiamondLoupeMetaData.ABI

// Deprecated: Use IDiamondLoupeMetaData.Sigs instead.
// IDiamondLoupeFuncSigs maps the 4-byte function signature to its string representation.
var IDiamondLoupeFuncSigs = IDiamondLoupeMetaData.Sigs

// IDiamondLoupe is an auto generated Go binding around an Ethereum contract.
type IDiamondLoupe struct {
	IDiamondLoupeCaller     // Read-only binding to the contract
	IDiamondLoupeTransactor // Write-only binding to the contract
	IDiamondLoupeFilterer   // Log filterer for contract events
}

// IDiamondLoupeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDiamondLoupeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDiamondLoupeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDiamondLoupeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDiamondLoupeSession struct {
	Contract     *IDiamondLoupe    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDiamondLoupeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDiamondLoupeCallerSession struct {
	Contract *IDiamondLoupeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IDiamondLoupeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDiamondLoupeTransactorSession struct {
	Contract     *IDiamondLoupeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IDiamondLoupeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDiamondLoupeRaw struct {
	Contract *IDiamondLoupe // Generic contract binding to access the raw methods on
}

// IDiamondLoupeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDiamondLoupeCallerRaw struct {
	Contract *IDiamondLoupeCaller // Generic read-only contract binding to access the raw methods on
}

// IDiamondLoupeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDiamondLoupeTransactorRaw struct {
	Contract *IDiamondLoupeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDiamondLoupe creates a new instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupe(address common.Address, backend bind.ContractBackend) (*IDiamondLoupe, error) {
	contract, err := bindIDiamondLoupe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupe{IDiamondLoupeCaller: IDiamondLoupeCaller{contract: contract}, IDiamondLoupeTransactor: IDiamondLoupeTransactor{contract: contract}, IDiamondLoupeFilterer: IDiamondLoupeFilterer{contract: contract}}, nil
}

// NewIDiamondLoupeCaller creates a new read-only instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeCaller(address common.Address, caller bind.ContractCaller) (*IDiamondLoupeCaller, error) {
	contract, err := bindIDiamondLoupe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeCaller{contract: contract}, nil
}

// NewIDiamondLoupeTransactor creates a new write-only instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeTransactor(address common.Address, transactor bind.ContractTransactor) (*IDiamondLoupeTransactor, error) {
	contract, err := bindIDiamondLoupe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeTransactor{contract: contract}, nil
}

// NewIDiamondLoupeFilterer creates a new log filterer instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeFilterer(address common.Address, filterer bind.ContractFilterer) (*IDiamondLoupeFilterer, error) {
	contract, err := bindIDiamondLoupe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeFilterer{contract: contract}, nil
}

// bindIDiamondLoupe binds a generic wrapper to an already deployed contract.
func bindIDiamondLoupe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDiamondLoupeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondLoupe *IDiamondLoupeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondLoupe.Contract.IDiamondLoupeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondLoupe *IDiamondLoupeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.IDiamondLoupeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondLoupe *IDiamondLoupeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.IDiamondLoupeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondLoupe *IDiamondLoupeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondLoupe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondLoupe *IDiamondLoupeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondLoupe *IDiamondLoupeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddress(&_IDiamondLoupe.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddress(&_IDiamondLoupe.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetAddresses() ([]common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddresses(&_IDiamondLoupe.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetAddresses() ([]common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddresses(&_IDiamondLoupe.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IDiamondLoupe.Contract.FacetFunctionSelectors(&_IDiamondLoupe.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IDiamondLoupe.Contract.FacetFunctionSelectors(&_IDiamondLoupe.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _IDiamondLoupe.Contract.Facets(&_IDiamondLoupe.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _IDiamondLoupe.Contract.Facets(&_IDiamondLoupe.CallOpts)
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// Deprecated: Use IERC165MetaData.Sigs instead.
// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = IERC165MetaData.Sigs

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
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
