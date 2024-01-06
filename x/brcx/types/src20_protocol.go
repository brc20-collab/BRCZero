package types

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	src20EntryPointABI abi.ABI
	//go:embed abi/Src20Entrypoint.json
	src20EntryPointABIJson []byte

	src20ABI abi.ABI
	//go:embed abi/Src20.json
	src20ABIJson []byte
)

func init() {
	src20EntryPointABI = GetEVMABIConfig(src20EntryPointABIJson)
	src20ABI = GetEVMABIConfig(src20ABIJson)
}

func UnpackSrc20EventContext(ret []byte) (Src20WrappedEvent, error) {
	var ec Src20WrappedEvent
	err := src20ABI.UnpackIntoInterface(&ec, EventsName, ret)
	if err != nil {
		return Src20WrappedEvent{}, err
	}
	return ec, nil
}
