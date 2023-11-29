package types

import (
	"bytes"
	_ "embed"
	"fmt"

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

func GetAllBalanceInput(addr string) ([]byte, error) {
	data, err := entryPointABI.Pack(GetAllBalanceMethodName, addr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetTotalTickHoldersInput() ([]byte, error) {
	data, err := entryPointABI.Pack(GetTotalTickHoldersMethodName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
