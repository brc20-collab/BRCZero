package types

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const RuneEvnetName = "RuneEvent"

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
	err := runeAlphaABI.UnpackIntoInterface(&ec, RuneEvnetName, ret)
	if err != nil {
		return RuneAlphaWrappedEvent{}, err
	}
	return ec, nil
}
