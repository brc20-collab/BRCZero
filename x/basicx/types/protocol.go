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
)

func init() {
	EvmABI = GetEVMABIConfig(abiJson)
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
