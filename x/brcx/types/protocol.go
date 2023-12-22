package types

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	BrczeroCalledMethodName = "entryPoint"
)

var (
	EvmABI abi.ABI
	//go:embed abi.json
	abiJson []byte

	BascisXABI abi.ABI
	//go:embed bascis_abi.json
	bascisXABIJson []byte
)

func init() {
	EvmABI = GetEVMABIConfig(abiJson)
	BascisXABI = GetEVMABIConfig(bascisXABIJson)
}

func GetEntryPointInput(context InscriptionContext, inscription string) ([]byte, error) {
	data, err := EvmABI.Pack(BrczeroCalledMethodName, context, inscription)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetEVMABIConfig(data []byte) abi.ABI {
	ret, err := abi.JSON(bytes.NewReader(data))
	if err != nil {
		panic(fmt.Errorf("json decode failed: %s", err.Error()))
	}
	return ret
}

func GetBascisXEntryPointInput(context string, inscription string) ([]byte, error) {
	data, err := BascisXABI.Pack(BrczeroCalledMethodName, context, inscription)
	if err != nil {
		return nil, err
	}
	return data, nil
}
