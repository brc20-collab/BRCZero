package types

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	EntryPointMethodName          = "entryPoint"
	GetTickInfoMethodName         = "getTickInformation"
	GetAllTickInfoMethodName      = "getAllTickInformation"
	GetBalanceMethodName          = "getBalance"
	GetAllBalanceMethodName       = "getAllBalance"
	GetTotalTickHoldersMethodName = "getTotalTickHolders"
)

var (
	entryPointABI abi.ABI
	//go:embed abi/BRC20EntryPointABI.json
	entryPointABIJson []byte
)

func init() {
	entryPointABI = GetEVMABIConfig(entryPointABIJson)
}

func GetEVMABIConfig(data []byte) abi.ABI {
	ret, err := abi.JSON(bytes.NewReader(data))
	if err != nil {
		panic(fmt.Errorf("json decode failed: %s", err.Error()))
	}
	return ret
}

func GetEntryPointInput(context InscriptionContext, inscription string) ([]byte, error) {
	data, err := entryPointABI.Pack(EntryPointMethodName, context, inscription)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetTickInfoInput(tickName string) ([]byte, error) {
	data, err := entryPointABI.Pack(GetTickInfoMethodName, tickName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetAllTickInfoInput() ([]byte, error) {
	data, err := entryPointABI.Pack(GetAllTickInfoMethodName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetBalanceInput(addr string, tickName string) ([]byte, error) {
	data, err := entryPointABI.Pack(GetBalanceMethodName, addr, tickName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetBalanceOutput(ret []byte) (*big.Int, *big.Int, *big.Int, error) {
	res, err := entryPointABI.Methods[GetBalanceMethodName].Outputs.Unpack(ret)
	if len(res) != 3 || err != nil {
		return nil, nil, nil, err
	}

	totalBalances, ok := res[0].(*big.Int)
	if !ok {
		return nil, nil, nil, errors.New("decode totalBalances failed")
	}

	availableBalance, ok := res[1].(*big.Int)
	if !ok {
		return nil, nil, nil, errors.New("decode availableBalance failed")
	}

	transferableBalance, ok := res[2].(*big.Int)
	if !ok {
		return nil, nil, nil, errors.New("decode transferableBalance failed")
	}

	return totalBalances, availableBalance, transferableBalance, nil
}

func GetAllBalanceInput(addr string) ([]byte, error) {
	data, err := entryPointABI.Pack(GetAllBalanceMethodName, addr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnpackGetAllBalanceOutput(ret []byte) ([]BRC20Balance, error) {
	var output []BRC20Balance
	err := entryPointABI.UnpackIntoInterface(&output, GetAllBalanceMethodName, ret)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetTotalTickHoldersInput() ([]byte, error) {
	data, err := entryPointABI.Pack(GetTotalTickHoldersMethodName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
