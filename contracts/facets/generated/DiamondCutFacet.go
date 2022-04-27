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

// DiamondCutFacetMetaData contains all meta data concerning the DiamondCutFacet contract.
var DiamondCutFacetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1f931c1c": "diamondCut((address,uint8,bytes4[])[],address,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b506115ab806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631f931c1c14610030575b600080fd5b61004361003e366004610fb1565b610045565b005b61004d61009e565b61009761005a85876110f7565b8484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061011a92505050565b5050505050565b60008051602061150a833981519152600401546001600160a01b031633146101185760405162461bcd60e51b815260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201526132b960f11b60648201526084015b60405180910390fd5b565b60005b83518110156102e057600084828151811061013a5761013a61123b565b60200260200101516020015190506000600281111561015b5761015b611251565b81600281111561016d5761016d611251565b036101bb576101b68583815181106101875761018761123b565b6020026020010151600001518684815181106101a5576101a561123b565b60200260200101516040015161032b565b6102cd565b60018160028111156101cf576101cf611251565b03610218576101b68583815181106101e9576101e961123b565b6020026020010151600001518684815181106102075761020761123b565b60200260200101516040015161058c565b600281600281111561022c5761022c611251565b03610275576101b68583815181106102465761024661123b565b6020026020010151600001518684815181106102645761026461123b565b602002602001015160400151610816565b60405162461bcd60e51b815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f727265637420466163657443756044820152663a20b1ba34b7b760c91b606482015260840161010f565b50806102d88161127d565b91505061011d565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673838383604051610314939291906112ee565b60405180910390a16103268282610933565b505050565b600081511161034c5760405162461bcd60e51b815260040161010f906113ee565b60008051602061150a8339815191526001600160a01b0383166103815760405162461bcd60e51b815260040161010f90611439565b6001600160a01b03831660009081526001820160205260408120549061ffff82169003610426576103ca8460405180606001604052806024815260200161155260249139610b40565b6002820180546001600160a01b038616600081815260018087016020908152604083208201805461ffff191661ffff90961695909517909455845490810185559381529190912090910180546001600160a01b03191690911790555b60005b83518110156100975760008482815181106104465761044661123b565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b031680156104e45760405162461bcd60e51b815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f6044820152746e207468617420616c72656164792065786973747360581b606482015260840161010f565b6001600160a01b03871660008181526001878101602090815260408084208054938401815584528184206008840401805463ffffffff60079095166004026101000a948502191660e089901c94909402939093179092556001600160e01b031986168352889052902080546001600160b01b031916909117600160a01b61ffff8716021790558361057481611485565b945050505080806105849061127d565b915050610429565b60008151116105ad5760405162461bcd60e51b815260040161010f906113ee565b60008051602061150a8339815191526001600160a01b0383166105e25760405162461bcd60e51b815260040161010f90611439565b6001600160a01b03831660009081526001820160205260408120549061ffff821690036106875761062b8460405180606001604052806024815260200161155260249139610b40565b6002820180546001600160a01b038616600081815260018087016020908152604083208201805461ffff191661ffff90961695909517909455845490810185559381529190912090910180546001600160a01b03191690911790555b60005b83518110156100975760008482815181106106a7576106a761123b565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b0390811690871681036107525760405162461bcd60e51b815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e20776974682073616d652066756e6374696f6e0000000000000000606482015260840161010f565b61075c8183610b61565b6001600160e01b03198216600081815260208781526040808320805461ffff60a01b1916600160a01b61ffff8b16021781556001600160a01b038c168085526001808c0185529285208054938401815585528385206008840401805463ffffffff60079095166004026101000a948502191660e08a901c94909402939093179092559390925287905281546001600160a01b031916179055836107fe81611485565b9450505050808061080e9061127d565b91505061068a565b60008151116108375760405162461bcd60e51b815260040161010f906113ee565b60008051602061150a8339815191526001600160a01b038316156108bc5760405162461bcd60e51b815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f76652066616365742061646472604482015275657373206d757374206265206164647265737328302960501b606482015260840161010f565b60005b825181101561092d5760008382815181106108dc576108dc61123b565b6020908102919091018101516001600160e01b031981166000908152918590526040909120549091506001600160a01b03166109188183610b61565b505080806109259061127d565b9150506108bf565b50505050565b6001600160a01b0382166109ba578051156109b65760405162461bcd60e51b815260206004820152603c60248201527f4c69624469616d6f6e644375743a205f696e697420697320616464726573732860448201527f3029206275745f63616c6c64617461206973206e6f7420656d70747900000000606482015260840161010f565b5050565b6000815111610a315760405162461bcd60e51b815260206004820152603d60248201527f4c69624469616d6f6e644375743a205f63616c6c6461746120697320656d707460448201527f7920627574205f696e6974206973206e6f742061646472657373283029000000606482015260840161010f565b6001600160a01b0382163014610a6357610a638260405180606001604052806028815260200161152a60289139610b40565b600080836001600160a01b031683604051610a7e91906114a6565b600060405180830381855af49150503d8060008114610ab9576040519150601f19603f3d011682016040523d82523d6000602084013e610abe565b606091505b50915091508161092d57805115610ae9578060405162461bcd60e51b815260040161010f91906114c2565b60405162461bcd60e51b815260206004820152602660248201527f4c69624469616d6f6e644375743a205f696e69742066756e6374696f6e2072656044820152651d995c9d195960d21b606482015260840161010f565b813b818161092d5760405162461bcd60e51b815260040161010f91906114c2565b60008051602061150a8339815191526001600160a01b038316610bec5760405162461bcd60e51b815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e2774206578697374000000000000000000606482015260840161010f565b306001600160a01b03841603610c5b5760405162461bcd60e51b815260206004820152602e60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201526d3a30b1363290333ab731ba34b7b760911b606482015260840161010f565b6001600160e01b03198216600090815260208281526040808320546001600160a01b0387168452600180860190935290832054600160a01b90910461ffff169291610ca5916114dc565b9050808214610d91576001600160a01b03851660009081526001840160205260408120805483908110610cda57610cda61123b565b600091825260208083206008830401546001600160a01b038a168452600188019091526040909220805460079092166004026101000a90920460e01b925082919085908110610d2b57610d2b61123b565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c929092029390931790556001600160e01b031992909216825284905260409020805461ffff60a01b1916600160a01b61ffff8516021790555b6001600160a01b03851660009081526001840160205260409020805480610dba57610dba6114f3565b60008281526020808220600860001990940193840401805463ffffffff600460078716026101000a0219169055919092556001600160e01b0319861682528490526040812080546001600160b01b0319169055819003610097576002830154600090610e28906001906114dc565b6001600160a01b038716600090815260018087016020526040909120015490915061ffff16808214610ee7576000856002018381548110610e6b57610e6b61123b565b6000918252602090912001546002870180546001600160a01b039092169250829184908110610e9c57610e9c61123b565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394851617905592909116815260018781019092526040902001805461ffff191661ffff83161790555b84600201805480610efa57610efa6114f3565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03891682526001878101909152604090912001805461ffff1916905550505050505050565b80356001600160a01b0381168114610f6357600080fd5b919050565b60008083601f840112610f7a57600080fd5b50813567ffffffffffffffff811115610f9257600080fd5b602083019150836020828501011115610faa57600080fd5b9250929050565b600080600080600060608688031215610fc957600080fd5b853567ffffffffffffffff80821115610fe157600080fd5b818801915088601f830112610ff557600080fd5b81358181111561100457600080fd5b8960208260051b850101111561101957600080fd5b6020830197508096505061102f60208901610f4c565b9450604088013591508082111561104557600080fd5b5061105288828901610f68565b969995985093965092949392505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561109c5761109c611063565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156110cb576110cb611063565b604052919050565b600067ffffffffffffffff8211156110ed576110ed611063565b5060051b60200190565b600061110a611105846110d3565b6110a2565b83815260208082019190600586811b86013681111561112857600080fd5b865b8181101561122e57803567ffffffffffffffff8082111561114b5760008081fd5b818a019150606082360312156111615760008081fd5b611169611079565b61117283610f4c565b815286830135600381106111865760008081fd5b818801526040838101358381111561119e5760008081fd5b939093019236601f8501126111b557600092508283fd5b833592506111c5611105846110d3565b83815292871b840188019288810190368511156111e25760008081fd5b948901945b848610156112175785356001600160e01b0319811681146112085760008081fd5b825294890194908901906111e7565b91830191909152508852505094830194830161112a565b5092979650505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161128f5761128f611267565b5060010190565b60005b838110156112b1578181015183820152602001611299565b8381111561092d5750506000910152565b600081518084526112da816020860160208601611296565b601f01601f19169290920160200192915050565b60006060808301818452808751808352608092508286019150828160051b8701016020808b0160005b848110156113be57898403607f19018652815180516001600160a01b0316855283810151898601906003811061135d57634e487b7160e01b600052602160045260246000fd5b868601526040918201519186018a905281519081905290840190600090898701905b808310156113a95783516001600160e01b031916825292860192600192909201919086019061137f565b50978501979550505090820190600101611317565b50506001600160a01b038a169088015286810360408801526113e081896112c2565b9a9950505050505050505050565b6020808252602b908201527f4c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e206660408201526a1858d95d081d1bc818dd5d60aa1b606082015260800190565b6020808252602c908201527f4c69624469616d6f6e644375743a204164642066616365742063616e2774206260408201526b65206164647265737328302960a01b606082015260800190565b600061ffff80831681810361149c5761149c611267565b6001019392505050565b600082516114b8818460208701611296565b9190910192915050565b6020815260006114d560208301846112c2565b9392505050565b6000828210156114ee576114ee611267565b500390565b634e487b7160e01b600052603160045260246000fdfec8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c4c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f64654c69624469616d6f6e644375743a204e657720666163657420686173206e6f20636f6465a2646970667358221220306876229d623010f89fba9cc96a3f95e860290d5d3bf146f990328f939b2bd464736f6c634300080d0033",
}

// DiamondCutFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondCutFacetMetaData.ABI instead.
var DiamondCutFacetABI = DiamondCutFacetMetaData.ABI

// Deprecated: Use DiamondCutFacetMetaData.Sigs instead.
// DiamondCutFacetFuncSigs maps the 4-byte function signature to its string representation.
var DiamondCutFacetFuncSigs = DiamondCutFacetMetaData.Sigs

// DiamondCutFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondCutFacetMetaData.Bin instead.
var DiamondCutFacetBin = DiamondCutFacetMetaData.Bin

// DeployDiamondCutFacet deploys a new Ethereum contract, binding an instance of DiamondCutFacet to it.
func DeployDiamondCutFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondCutFacet, error) {
	parsed, err := DiamondCutFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondCutFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondCutFacet{DiamondCutFacetCaller: DiamondCutFacetCaller{contract: contract}, DiamondCutFacetTransactor: DiamondCutFacetTransactor{contract: contract}, DiamondCutFacetFilterer: DiamondCutFacetFilterer{contract: contract}}, nil
}

// DiamondCutFacet is an auto generated Go binding around an Ethereum contract.
type DiamondCutFacet struct {
	DiamondCutFacetCaller     // Read-only binding to the contract
	DiamondCutFacetTransactor // Write-only binding to the contract
	DiamondCutFacetFilterer   // Log filterer for contract events
}

// DiamondCutFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCutFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondCutFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondCutFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondCutFacetSession struct {
	Contract     *DiamondCutFacet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCutFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCutFacetCallerSession struct {
	Contract *DiamondCutFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DiamondCutFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondCutFacetTransactorSession struct {
	Contract     *DiamondCutFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DiamondCutFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondCutFacetRaw struct {
	Contract *DiamondCutFacet // Generic contract binding to access the raw methods on
}

// DiamondCutFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCutFacetCallerRaw struct {
	Contract *DiamondCutFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondCutFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondCutFacetTransactorRaw struct {
	Contract *DiamondCutFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondCutFacet creates a new instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacet(address common.Address, backend bind.ContractBackend) (*DiamondCutFacet, error) {
	contract, err := bindDiamondCutFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacet{DiamondCutFacetCaller: DiamondCutFacetCaller{contract: contract}, DiamondCutFacetTransactor: DiamondCutFacetTransactor{contract: contract}, DiamondCutFacetFilterer: DiamondCutFacetFilterer{contract: contract}}, nil
}

// NewDiamondCutFacetCaller creates a new read-only instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetCaller(address common.Address, caller bind.ContractCaller) (*DiamondCutFacetCaller, error) {
	contract, err := bindDiamondCutFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetCaller{contract: contract}, nil
}

// NewDiamondCutFacetTransactor creates a new write-only instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondCutFacetTransactor, error) {
	contract, err := bindDiamondCutFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetTransactor{contract: contract}, nil
}

// NewDiamondCutFacetFilterer creates a new log filterer instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondCutFacetFilterer, error) {
	contract, err := bindDiamondCutFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetFilterer{contract: contract}, nil
}

// bindDiamondCutFacet binds a generic wrapper to an already deployed contract.
func bindDiamondCutFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondCutFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCutFacet *DiamondCutFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCutFacet.Contract.DiamondCutFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCutFacet *DiamondCutFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCutFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCutFacet *DiamondCutFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCutFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCutFacet *DiamondCutFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCutFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCutFacet *DiamondCutFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCutFacet *DiamondCutFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.contract.Transact(opts, method, params...)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCut(&_DiamondCutFacet.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetTransactorSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCut(&_DiamondCutFacet.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCutFacetDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCutIterator struct {
	Event *DiamondCutFacetDiamondCut // Event containing the contract specifics and raw log

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
func (it *DiamondCutFacetDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondCutFacetDiamondCut)
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
		it.Event = new(DiamondCutFacetDiamondCut)
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
func (it *DiamondCutFacetDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondCutFacetDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondCutFacetDiamondCut represents a DiamondCut event raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*DiamondCutFacetDiamondCutIterator, error) {

	logs, sub, err := _DiamondCutFacet.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetDiamondCutIterator{contract: _DiamondCutFacet.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *DiamondCutFacetDiamondCut) (event.Subscription, error) {

	logs, sub, err := _DiamondCutFacet.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondCutFacetDiamondCut)
				if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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
func (_DiamondCutFacet *DiamondCutFacetFilterer) ParseDiamondCut(log types.Log) (*DiamondCutFacetDiamondCut, error) {
	event := new(DiamondCutFacetDiamondCut)
	if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
