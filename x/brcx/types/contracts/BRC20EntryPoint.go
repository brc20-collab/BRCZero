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
	_ = abi.ConvertType
)

// IBrc20EntryPointBrc20Balance is an auto generated low-level Go binding around an user-defined struct.
type IBrc20EntryPointBrc20Balance struct {
	Tick                string
	TotalBalance        *big.Int
	AvailableBalance    *big.Int
	TransferableBalance *big.Int
}

// IBrc20EntryPointBrc20Information is an auto generated low-level Go binding around an user-defined struct.
type IBrc20EntryPointBrc20Information struct {
	Tick        string
	TickAddress common.Address
	MaxSupply   *big.Int
	NowSupply   *big.Int
	Decimals    *big.Int
	Lim         *big.Int
}

// IBrc20EntryPointInScriptionContext is an auto generated low-level Go binding around an user-defined struct.
type IBrc20EntryPointInScriptionContext struct {
	InscriptionId     string
	InscriptionNumber int64
	IsTransfer        bool
	Txid              string
	Sender            string
	Receiver          string
	CommitInput       string
	RevealOutput      string
	OldSatPoint       string
	NewSatPoint       string
	BlockHash         string
	BlockTime         uint32
	BlockHeight       uint64
}

// BRC20EntryPointMetaData contains all meta data concerning the BRC20EntryPoint contract.
var BRC20EntryPointMetaData = &bind.MetaData{
	ABI: "[{\"anonymofetus\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRCX_ADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_DEC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_DEPLOY_ITEM_SIZE_MAX_WITH_DEC_AND_LIM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_DEPLOY_ITEM_SIZE_MIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_DEPLOY_ITEM_SIZE_WITH_DEC_OR_LIM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_MINT_ITEM_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_TRANSFER_ITEM_SIZE_MAX_WITH_FEE_AND_TO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_TRANSFER_ITEM_SIZE_MIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_TRANSFER_ITEM_SIZE_WITH_FEE_OR_TO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"inscriptionId\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"inscriptionNumber\",\"type\":\"int64\"},{\"internalType\":\"bool\",\"name\":\"isTransfer\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sender\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"commitInput\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"revealOutput\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"oldSatPoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newSatPoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blockHash\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"blockTime\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"internalType\":\"structIBrc20EntryPoint.InScriptionContext\",\"name\":\"inScriptionContext\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"inscription\",\"type\":\"string\"}],\"name\":\"entryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"getAllBalance\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"tick\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferableBalance\",\"type\":\"uint256\"}],\"internalType\":\"structIBrc20EntryPoint.Brc20Balance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTickInformation\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"tick\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tickAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nowSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lim\",\"type\":\"uint256\"}],\"internalType\":\"structIBrc20EntryPoint.Brc20Information[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tick\",\"type\":\"string\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalBalances\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transferableBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexTo\",\"type\":\"uint256\"}],\"name\":\"getTickDeployedSlice\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tick\",\"type\":\"string\"}],\"name\":\"getTickInformation\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"tick\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tickAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nowSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lim\",\"type\":\"uint256\"}],\"internalType\":\"structIBrc20EntryPoint.Brc20Information\",\"name\":\"tickInformation\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalTickHolders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tickHodlers\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precomplieContarct\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tickDeployed\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BRC20EntryPointABI is the input ABI used to generate the binding from.
// Deprecated: Use BRC20EntryPointMetaData.ABI instead.
var BRC20EntryPointABI = BRC20EntryPointMetaData.ABI

// BRC20EntryPoint is an auto generated Go binding around an Ethereum contract.
type BRC20EntryPoint struct {
	BRC20EntryPointCaller     // Read-only binding to the contract
	BRC20EntryPointTransactor // Write-only binding to the contract
	BRC20EntryPointFilterer   // Log filterer for contract events
}

// BRC20EntryPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type BRC20EntryPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BRC20EntryPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BRC20EntryPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BRC20EntryPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BRC20EntryPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BRC20EntryPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BRC20EntryPointSession struct {
	Contract     *BRC20EntryPoint  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BRC20EntryPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BRC20EntryPointCallerSession struct {
	Contract *BRC20EntryPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BRC20EntryPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BRC20EntryPointTransactorSession struct {
	Contract     *BRC20EntryPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BRC20EntryPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type BRC20EntryPointRaw struct {
	Contract *BRC20EntryPoint // Generic contract binding to access the raw methods on
}

// BRC20EntryPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BRC20EntryPointCallerRaw struct {
	Contract *BRC20EntryPointCaller // Generic read-only contract binding to access the raw methods on
}

// BRC20EntryPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BRC20EntryPointTransactorRaw struct {
	Contract *BRC20EntryPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBRC20EntryPoint creates a new instance of BRC20EntryPoint, bound to a specific deployed contract.
func NewBRC20EntryPoint(address common.Address, backend bind.ContractBackend) (*BRC20EntryPoint, error) {
	contract, err := bindBRC20EntryPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BRC20EntryPoint{BRC20EntryPointCaller: BRC20EntryPointCaller{contract: contract}, BRC20EntryPointTransactor: BRC20EntryPointTransactor{contract: contract}, BRC20EntryPointFilterer: BRC20EntryPointFilterer{contract: contract}}, nil
}

// NewBRC20EntryPointCaller creates a new read-only instance of BRC20EntryPoint, bound to a specific deployed contract.
func NewBRC20EntryPointCaller(address common.Address, caller bind.ContractCaller) (*BRC20EntryPointCaller, error) {
	contract, err := bindBRC20EntryPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BRC20EntryPointCaller{contract: contract}, nil
}

// NewBRC20EntryPointTransactor creates a new write-only instance of BRC20EntryPoint, bound to a specific deployed contract.
func NewBRC20EntryPointTransactor(address common.Address, transactor bind.ContractTransactor) (*BRC20EntryPointTransactor, error) {
	contract, err := bindBRC20EntryPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BRC20EntryPointTransactor{contract: contract}, nil
}

// NewBRC20EntryPointFilterer creates a new log filterer instance of BRC20EntryPoint, bound to a specific deployed contract.
func NewBRC20EntryPointFilterer(address common.Address, filterer bind.ContractFilterer) (*BRC20EntryPointFilterer, error) {
	contract, err := bindBRC20EntryPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BRC20EntryPointFilterer{contract: contract}, nil
}

// bindBRC20EntryPoint binds a generic wrapper to an already deployed contract.
func bindBRC20EntryPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BRC20EntryPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BRC20EntryPoint *BRC20EntryPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BRC20EntryPoint.Contract.BRC20EntryPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BRC20EntryPoint *BRC20EntryPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.BRC20EntryPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BRC20EntryPoint *BRC20EntryPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.BRC20EntryPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BRC20EntryPoint *BRC20EntryPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BRC20EntryPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BRC20EntryPoint *BRC20EntryPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BRC20EntryPoint *BRC20EntryPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.contract.Transact(opts, method, params...)
}

// BRCXADMIN is a free data retrieval call binding the contract method 0xcde3608e.
//
// Solidity: function BRCX_ADMIN() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointCaller) BRCXADMIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "BRCX_ADMIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRCXADMIN is a free data retrieval call binding the contract method 0xcde3608e.
//
// Solidity: function BRCX_ADMIN() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointSession) BRCXADMIN() (common.Address, error) {
	return _BRC20EntryPoint.Contract.BRCXADMIN(&_BRC20EntryPoint.CallOpts)
}

// BRCXADMIN is a free data retrieval call binding the contract method 0xcde3608e.
//
// Solidity: function BRCX_ADMIN() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) BRCXADMIN() (common.Address, error) {
	return _BRC20EntryPoint.Contract.BRCXADMIN(&_BRC20EntryPoint.CallOpts)
}

// DEFAULTDEC is a free data retrieval call binding the contract method 0x58648cdf.
//
// Solidity: function DEFAULT_DEC() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCaller) DEFAULTDEC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "DEFAULT_DEC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTDEC is a free data retrieval call binding the contract method 0x58648cdf.
//
// Solidity: function DEFAULT_DEC() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointSession) DEFAULTDEC() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.DEFAULTDEC(&_BRC20EntryPoint.CallOpts)
}

// DEFAULTDEC is a free data retrieval call binding the contract method 0x58648cdf.
//
// Solidity: function DEFAULT_DEC() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) DEFAULTDEC() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.DEFAULTDEC(&_BRC20EntryPoint.CallOpts)
}

// OPDEPLOYITEMSIZEMAXWITHDECANDLIM is a free data retrieval call binding the contract method 0xcb60a0a2.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_MAX_WITH_DEC_AND_LIM() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCaller) OPDEPLOYITEMSIZEMAXWITHDECANDLIM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "OP_DEPLOY_ITEM_SIZE_MAX_WITH_DEC_AND_LIM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPDEPLOYITEMSIZEMAXWITHDECANDLIM is a free data retrieval call binding the contract method 0xcb60a0a2.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_MAX_WITH_DEC_AND_LIM() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointSession) OPDEPLOYITEMSIZEMAXWITHDECANDLIM() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPDEPLOYITEMSIZEMAXWITHDECANDLIM(&_BRC20EntryPoint.CallOpts)
}

// OPDEPLOYITEMSIZEMAXWITHDECANDLIM is a free data retrieval call binding the contract method 0xcb60a0a2.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_MAX_WITH_DEC_AND_LIM() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) OPDEPLOYITEMSIZEMAXWITHDECANDLIM() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPDEPLOYITEMSIZEMAXWITHDECANDLIM(&_BRC20EntryPoint.CallOpts)
}

// OPDEPLOYITEMSIZEMIN is a free data retrieval call binding the contract method 0x360b5d8e.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_MIN() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCaller) OPDEPLOYITEMSIZEMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "OP_DEPLOY_ITEM_SIZE_MIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPDEPLOYITEMSIZEMIN is a free data retrieval call binding the contract method 0x360b5d8e.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_MIN() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointSession) OPDEPLOYITEMSIZEMIN() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPDEPLOYITEMSIZEMIN(&_BRC20EntryPoint.CallOpts)
}

// OPDEPLOYITEMSIZEMIN is a free data retrieval call binding the contract method 0x360b5d8e.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_MIN() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) OPDEPLOYITEMSIZEMIN() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPDEPLOYITEMSIZEMIN(&_BRC20EntryPoint.CallOpts)
}

// OPDEPLOYITEMSIZEWITHDECORLIM is a free data retrieval call binding the contract method 0xded3384c.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_WITH_DEC_OR_LIM() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCaller) OPDEPLOYITEMSIZEWITHDECORLIM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "OP_DEPLOY_ITEM_SIZE_WITH_DEC_OR_LIM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPDEPLOYITEMSIZEWITHDECORLIM is a free data retrieval call binding the contract method 0xded3384c.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_WITH_DEC_OR_LIM() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointSession) OPDEPLOYITEMSIZEWITHDECORLIM() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPDEPLOYITEMSIZEWITHDECORLIM(&_BRC20EntryPoint.CallOpts)
}

// OPDEPLOYITEMSIZEWITHDECORLIM is a free data retrieval call binding the contract method 0xded3384c.
//
// Solidity: function OP_DEPLOY_ITEM_SIZE_WITH_DEC_OR_LIM() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) OPDEPLOYITEMSIZEWITHDECORLIM() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPDEPLOYITEMSIZEWITHDECORLIM(&_BRC20EntryPoint.CallOpts)
}

// OPMINTITEMSIZE is a free data retrieval call binding the contract method 0xe7ab22e0.
//
// Solidity: function OP_MINT_ITEM_SIZE() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCaller) OPMINTITEMSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "OP_MINT_ITEM_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPMINTITEMSIZE is a free data retrieval call binding the contract method 0xe7ab22e0.
//
// Solidity: function OP_MINT_ITEM_SIZE() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointSession) OPMINTITEMSIZE() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPMINTITEMSIZE(&_BRC20EntryPoint.CallOpts)
}

// OPMINTITEMSIZE is a free data retrieval call binding the contract method 0xe7ab22e0.
//
// Solidity: function OP_MINT_ITEM_SIZE() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) OPMINTITEMSIZE() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPMINTITEMSIZE(&_BRC20EntryPoint.CallOpts)
}

// OPTRANSFERITEMSIZEMAXWITHFEEANDTO is a free data retrieval call binding the contract method 0x85410acd.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_MAX_WITH_FEE_AND_TO() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCaller) OPTRANSFERITEMSIZEMAXWITHFEEANDTO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "OP_TRANSFER_ITEM_SIZE_MAX_WITH_FEE_AND_TO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPTRANSFERITEMSIZEMAXWITHFEEANDTO is a free data retrieval call binding the contract method 0x85410acd.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_MAX_WITH_FEE_AND_TO() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointSession) OPTRANSFERITEMSIZEMAXWITHFEEANDTO() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPTRANSFERITEMSIZEMAXWITHFEEANDTO(&_BRC20EntryPoint.CallOpts)
}

// OPTRANSFERITEMSIZEMAXWITHFEEANDTO is a free data retrieval call binding the contract method 0x85410acd.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_MAX_WITH_FEE_AND_TO() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) OPTRANSFERITEMSIZEMAXWITHFEEANDTO() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPTRANSFERITEMSIZEMAXWITHFEEANDTO(&_BRC20EntryPoint.CallOpts)
}

// OPTRANSFERITEMSIZEMIN is a free data retrieval call binding the contract method 0x447c6380.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_MIN() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCaller) OPTRANSFERITEMSIZEMIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "OP_TRANSFER_ITEM_SIZE_MIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPTRANSFERITEMSIZEMIN is a free data retrieval call binding the contract method 0x447c6380.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_MIN() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointSession) OPTRANSFERITEMSIZEMIN() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPTRANSFERITEMSIZEMIN(&_BRC20EntryPoint.CallOpts)
}

// OPTRANSFERITEMSIZEMIN is a free data retrieval call binding the contract method 0x447c6380.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_MIN() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) OPTRANSFERITEMSIZEMIN() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPTRANSFERITEMSIZEMIN(&_BRC20EntryPoint.CallOpts)
}

// OPTRANSFERITEMSIZEWITHFEEORTO is a free data retrieval call binding the contract method 0x088946d4.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_WITH_FEE_OR_TO() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCaller) OPTRANSFERITEMSIZEWITHFEEORTO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "OP_TRANSFER_ITEM_SIZE_WITH_FEE_OR_TO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPTRANSFERITEMSIZEWITHFEEORTO is a free data retrieval call binding the contract method 0x088946d4.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_WITH_FEE_OR_TO() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointSession) OPTRANSFERITEMSIZEWITHFEEORTO() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPTRANSFERITEMSIZEWITHFEEORTO(&_BRC20EntryPoint.CallOpts)
}

// OPTRANSFERITEMSIZEWITHFEEORTO is a free data retrieval call binding the contract method 0x088946d4.
//
// Solidity: function OP_TRANSFER_ITEM_SIZE_WITH_FEE_OR_TO() view returns(uint256)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) OPTRANSFERITEMSIZEWITHFEEORTO() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.OPTRANSFERITEMSIZEWITHFEEORTO(&_BRC20EntryPoint.CallOpts)
}

// GetAllBalance is a free data retrieval call binding the contract method 0x017ae8ee.
//
// Solidity: function getAllBalance(string addr) view returns((string,uint256,uint256,uint256)[])
func (_BRC20EntryPoint *BRC20EntryPointCaller) GetAllBalance(opts *bind.CallOpts, addr string) ([]IBrc20EntryPointBrc20Balance, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "getAllBalance", addr)

	if err != nil {
		return *new([]IBrc20EntryPointBrc20Balance), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBrc20EntryPointBrc20Balance)).(*[]IBrc20EntryPointBrc20Balance)

	return out0, err

}

// GetAllBalance is a free data retrieval call binding the contract method 0x017ae8ee.
//
// Solidity: function getAllBalance(string addr) view returns((string,uint256,uint256,uint256)[])
func (_BRC20EntryPoint *BRC20EntryPointSession) GetAllBalance(addr string) ([]IBrc20EntryPointBrc20Balance, error) {
	return _BRC20EntryPoint.Contract.GetAllBalance(&_BRC20EntryPoint.CallOpts, addr)
}

// GetAllBalance is a free data retrieval call binding the contract method 0x017ae8ee.
//
// Solidity: function getAllBalance(string addr) view returns((string,uint256,uint256,uint256)[])
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) GetAllBalance(addr string) ([]IBrc20EntryPointBrc20Balance, error) {
	return _BRC20EntryPoint.Contract.GetAllBalance(&_BRC20EntryPoint.CallOpts, addr)
}

// GetAllTickInformation is a free data retrieval call binding the contract method 0xd836fd03.
//
// Solidity: function getAllTickInformation() view returns((string,address,uint256,uint256,uint256,uint256)[])
func (_BRC20EntryPoint *BRC20EntryPointCaller) GetAllTickInformation(opts *bind.CallOpts) ([]IBrc20EntryPointBrc20Information, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "getAllTickInformation")

	if err != nil {
		return *new([]IBrc20EntryPointBrc20Information), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBrc20EntryPointBrc20Information)).(*[]IBrc20EntryPointBrc20Information)

	return out0, err

}

// GetAllTickInformation is a free data retrieval call binding the contract method 0xd836fd03.
//
// Solidity: function getAllTickInformation() view returns((string,address,uint256,uint256,uint256,uint256)[])
func (_BRC20EntryPoint *BRC20EntryPointSession) GetAllTickInformation() ([]IBrc20EntryPointBrc20Information, error) {
	return _BRC20EntryPoint.Contract.GetAllTickInformation(&_BRC20EntryPoint.CallOpts)
}

// GetAllTickInformation is a free data retrieval call binding the contract method 0xd836fd03.
//
// Solidity: function getAllTickInformation() view returns((string,address,uint256,uint256,uint256,uint256)[])
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) GetAllTickInformation() ([]IBrc20EntryPointBrc20Information, error) {
	return _BRC20EntryPoint.Contract.GetAllTickInformation(&_BRC20EntryPoint.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x6ac3d07b.
//
// Solidity: function getBalance(string addr, string tick) view returns(uint256 totalBalances, uint256 availableBalance, uint256 transferableBalance)
func (_BRC20EntryPoint *BRC20EntryPointCaller) GetBalance(opts *bind.CallOpts, addr string, tick string) (struct {
	TotalBalances       *big.Int
	AvailableBalance    *big.Int
	TransferableBalance *big.Int
}, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "getBalance", addr, tick)

	outstruct := new(struct {
		TotalBalances       *big.Int
		AvailableBalance    *big.Int
		TransferableBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalBalances = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AvailableBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TransferableBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBalance is a free data retrieval call binding the contract method 0x6ac3d07b.
//
// Solidity: function getBalance(string addr, string tick) view returns(uint256 totalBalances, uint256 availableBalance, uint256 transferableBalance)
func (_BRC20EntryPoint *BRC20EntryPointSession) GetBalance(addr string, tick string) (struct {
	TotalBalances       *big.Int
	AvailableBalance    *big.Int
	TransferableBalance *big.Int
}, error) {
	return _BRC20EntryPoint.Contract.GetBalance(&_BRC20EntryPoint.CallOpts, addr, tick)
}

// GetBalance is a free data retrieval call binding the contract method 0x6ac3d07b.
//
// Solidity: function getBalance(string addr, string tick) view returns(uint256 totalBalances, uint256 availableBalance, uint256 transferableBalance)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) GetBalance(addr string, tick string) (struct {
	TotalBalances       *big.Int
	AvailableBalance    *big.Int
	TransferableBalance *big.Int
}, error) {
	return _BRC20EntryPoint.Contract.GetBalance(&_BRC20EntryPoint.CallOpts, addr, tick)
}

// GetTickDeployedSlice is a free data retrieval call binding the contract method 0xf719f6aa.
//
// Solidity: function getTickDeployedSlice(uint256 indexFrom, uint256 indexTo) view returns(string[])
func (_BRC20EntryPoint *BRC20EntryPointCaller) GetTickDeployedSlice(opts *bind.CallOpts, indexFrom *big.Int, indexTo *big.Int) ([]string, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "getTickDeployedSlice", indexFrom, indexTo)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTickDeployedSlice is a free data retrieval call binding the contract method 0xf719f6aa.
//
// Solidity: function getTickDeployedSlice(uint256 indexFrom, uint256 indexTo) view returns(string[])
func (_BRC20EntryPoint *BRC20EntryPointSession) GetTickDeployedSlice(indexFrom *big.Int, indexTo *big.Int) ([]string, error) {
	return _BRC20EntryPoint.Contract.GetTickDeployedSlice(&_BRC20EntryPoint.CallOpts, indexFrom, indexTo)
}

// GetTickDeployedSlice is a free data retrieval call binding the contract method 0xf719f6aa.
//
// Solidity: function getTickDeployedSlice(uint256 indexFrom, uint256 indexTo) view returns(string[])
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) GetTickDeployedSlice(indexFrom *big.Int, indexTo *big.Int) ([]string, error) {
	return _BRC20EntryPoint.Contract.GetTickDeployedSlice(&_BRC20EntryPoint.CallOpts, indexFrom, indexTo)
}

// GetTickInformation is a free data retrieval call binding the contract method 0x9234b733.
//
// Solidity: function getTickInformation(string tick) view returns((string,address,uint256,uint256,uint256,uint256) tickInformation)
func (_BRC20EntryPoint *BRC20EntryPointCaller) GetTickInformation(opts *bind.CallOpts, tick string) (IBrc20EntryPointBrc20Information, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "getTickInformation", tick)

	if err != nil {
		return *new(IBrc20EntryPointBrc20Information), err
	}

	out0 := *abi.ConvertType(out[0], new(IBrc20EntryPointBrc20Information)).(*IBrc20EntryPointBrc20Information)

	return out0, err

}

// GetTickInformation is a free data retrieval call binding the contract method 0x9234b733.
//
// Solidity: function getTickInformation(string tick) view returns((string,address,uint256,uint256,uint256,uint256) tickInformation)
func (_BRC20EntryPoint *BRC20EntryPointSession) GetTickInformation(tick string) (IBrc20EntryPointBrc20Information, error) {
	return _BRC20EntryPoint.Contract.GetTickInformation(&_BRC20EntryPoint.CallOpts, tick)
}

// GetTickInformation is a free data retrieval call binding the contract method 0x9234b733.
//
// Solidity: function getTickInformation(string tick) view returns((string,address,uint256,uint256,uint256,uint256) tickInformation)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) GetTickInformation(tick string) (IBrc20EntryPointBrc20Information, error) {
	return _BRC20EntryPoint.Contract.GetTickInformation(&_BRC20EntryPoint.CallOpts, tick)
}

// GetTotalTickHolders is a free data retrieval call binding the contract method 0x67ecf69b.
//
// Solidity: function getTotalTickHolders() view returns(uint256 tickHodlers)
func (_BRC20EntryPoint *BRC20EntryPointCaller) GetTotalTickHolders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "getTotalTickHolders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalTickHolders is a free data retrieval call binding the contract method 0x67ecf69b.
//
// Solidity: function getTotalTickHolders() view returns(uint256 tickHodlers)
func (_BRC20EntryPoint *BRC20EntryPointSession) GetTotalTickHolders() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.GetTotalTickHolders(&_BRC20EntryPoint.CallOpts)
}

// GetTotalTickHolders is a free data retrieval call binding the contract method 0x67ecf69b.
//
// Solidity: function getTotalTickHolders() view returns(uint256 tickHodlers)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) GetTotalTickHolders() (*big.Int, error) {
	return _BRC20EntryPoint.Contract.GetTotalTickHolders(&_BRC20EntryPoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointSession) Owner() (common.Address, error) {
	return _BRC20EntryPoint.Contract.Owner(&_BRC20EntryPoint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) Owner() (common.Address, error) {
	return _BRC20EntryPoint.Contract.Owner(&_BRC20EntryPoint.CallOpts)
}

// PrecomplieContarct is a free data retrieval call binding the contract method 0xd57dc253.
//
// Solidity: function precomplieContarct() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointCaller) PrecomplieContarct(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "precomplieContarct")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecomplieContarct is a free data retrieval call binding the contract method 0xd57dc253.
//
// Solidity: function precomplieContarct() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointSession) PrecomplieContarct() (common.Address, error) {
	return _BRC20EntryPoint.Contract.PrecomplieContarct(&_BRC20EntryPoint.CallOpts)
}

// PrecomplieContarct is a free data retrieval call binding the contract method 0xd57dc253.
//
// Solidity: function precomplieContarct() view returns(address)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) PrecomplieContarct() (common.Address, error) {
	return _BRC20EntryPoint.Contract.PrecomplieContarct(&_BRC20EntryPoint.CallOpts)
}

// TickDeployed is a free data retrieval call binding the contract method 0x705bff65.
//
// Solidity: function tickDeployed(uint256 ) view returns(string)
func (_BRC20EntryPoint *BRC20EntryPointCaller) TickDeployed(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _BRC20EntryPoint.contract.Call(opts, &out, "tickDeployed", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TickDeployed is a free data retrieval call binding the contract method 0x705bff65.
//
// Solidity: function tickDeployed(uint256 ) view returns(string)
func (_BRC20EntryPoint *BRC20EntryPointSession) TickDeployed(arg0 *big.Int) (string, error) {
	return _BRC20EntryPoint.Contract.TickDeployed(&_BRC20EntryPoint.CallOpts, arg0)
}

// TickDeployed is a free data retrieval call binding the contract method 0x705bff65.
//
// Solidity: function tickDeployed(uint256 ) view returns(string)
func (_BRC20EntryPoint *BRC20EntryPointCallerSession) TickDeployed(arg0 *big.Int) (string, error) {
	return _BRC20EntryPoint.Contract.TickDeployed(&_BRC20EntryPoint.CallOpts, arg0)
}

// EntryPoint is a paid mutator transaction binding the contract method 0xe8d5e69a.
//
// Solidity: function entryPoint((string,int64,bool,string,string,string,string,string,string,string,string,uint32,uint64) inScriptionContext, string inscription) returns()
func (_BRC20EntryPoint *BRC20EntryPointTransactor) EntryPoint(opts *bind.TransactOpts, inScriptionContext IBrc20EntryPointInScriptionContext, inscription string) (*types.Transaction, error) {
	return _BRC20EntryPoint.contract.Transact(opts, "entryPoint", inScriptionContext, inscription)
}

// EntryPoint is a paid mutator transaction binding the contract method 0xe8d5e69a.
//
// Solidity: function entryPoint((string,int64,bool,string,string,string,string,string,string,string,string,uint32,uint64) inScriptionContext, string inscription) returns()
func (_BRC20EntryPoint *BRC20EntryPointSession) EntryPoint(inScriptionContext IBrc20EntryPointInScriptionContext, inscription string) (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.EntryPoint(&_BRC20EntryPoint.TransactOpts, inScriptionContext, inscription)
}

// EntryPoint is a paid mutator transaction binding the contract method 0xe8d5e69a.
//
// Solidity: function entryPoint((string,int64,bool,string,string,string,string,string,string,string,string,uint32,uint64) inScriptionContext, string inscription) returns()
func (_BRC20EntryPoint *BRC20EntryPointTransactorSession) EntryPoint(inScriptionContext IBrc20EntryPointInScriptionContext, inscription string) (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.EntryPoint(&_BRC20EntryPoint.TransactOpts, inScriptionContext, inscription)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BRC20EntryPoint *BRC20EntryPointTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BRC20EntryPoint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BRC20EntryPoint *BRC20EntryPointSession) RenounceOwnership() (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.RenounceOwnership(&_BRC20EntryPoint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BRC20EntryPoint *BRC20EntryPointTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.RenounceOwnership(&_BRC20EntryPoint.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BRC20EntryPoint *BRC20EntryPointTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BRC20EntryPoint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BRC20EntryPoint *BRC20EntryPointSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.TransferOwnership(&_BRC20EntryPoint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BRC20EntryPoint *BRC20EntryPointTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BRC20EntryPoint.Contract.TransferOwnership(&_BRC20EntryPoint.TransactOpts, newOwner)
}

// BRC20EntryPointOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BRC20EntryPoint contract.
type BRC20EntryPointOwnershipTransferredIterator struct {
	Event *BRC20EntryPointOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BRC20EntryPointOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BRC20EntryPointOwnershipTransferred)
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
		it.Event = new(BRC20EntryPointOwnershipTransferred)
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
func (it *BRC20EntryPointOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BRC20EntryPointOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BRC20EntryPointOwnershipTransferred represents a OwnershipTransferred event raised by the BRC20EntryPoint contract.
type BRC20EntryPointOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BRC20EntryPoint *BRC20EntryPointFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BRC20EntryPointOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BRC20EntryPoint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BRC20EntryPointOwnershipTransferredIterator{contract: _BRC20EntryPoint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BRC20EntryPoint *BRC20EntryPointFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BRC20EntryPointOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BRC20EntryPoint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BRC20EntryPointOwnershipTransferred)
				if err := _BRC20EntryPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BRC20EntryPoint *BRC20EntryPointFilterer) ParseOwnershipTransferred(log types.Log) (*BRC20EntryPointOwnershipTransferred, error) {
	event := new(BRC20EntryPointOwnershipTransferred)
	if err := _BRC20EntryPoint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
