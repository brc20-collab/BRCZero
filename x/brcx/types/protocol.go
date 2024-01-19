package types

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	EntryPointMethodName = "entryPoint"

	EventsName = "Events"
)

var (
	basicXEntryPointABI abi.ABI
	//go:embed abi/basicXEntryPoint.json
	basicXEntryPointABIJson []byte
)

func init() {
	basicXEntryPointABI = GetEVMABIConfig(basicXEntryPointABIJson)
}

func GetEVMABIConfig(data []byte) abi.ABI {
	ret, err := abi.JSON(bytes.NewReader(data))
	if err != nil {
		panic(fmt.Errorf("json decode failed: %s", err.Error()))
	}
	return ret
}

func GetBasicXEntryPointInput(context string, inscription string) ([]byte, error) {
	data, err := basicXEntryPointABI.Pack(EntryPointMethodName, context, inscription)
	if err != nil {
		return nil, err
	}
	return data, nil
}
