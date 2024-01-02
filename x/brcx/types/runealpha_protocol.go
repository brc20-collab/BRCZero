package types

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	runeAlphaEntryPointABI abi.ABI
	//go:embed abi/RuneAlphaEntryPoint.json
	runeAlphaEntryPointABIJson []byte

	runeAlphaABI abi.ABI
	//go:embed abi/RuneAlpha.json
	runeAlphaABIJson []byte
)

func init() {
	runeAlphaEntryPointABI = GetEVMABIConfig(runeAlphaEntryPointABIJson)
	runeAlphaABI = GetEVMABIConfig(runeAlphaABIJson)
}

func UnpackRuneAlphaEventContext(ret []byte) (RuneAlphaWrappedEvent, error) {
	var ec RuneAlphaWrappedEvent
	err := runeAlphaABI.UnpackIntoInterface(&ec, EventsName, ret)
	if err != nil {
		return RuneAlphaWrappedEvent{}, err
	}
	return ec, nil
}
