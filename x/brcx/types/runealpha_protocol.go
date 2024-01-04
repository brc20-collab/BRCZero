package types

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

const RuneEvnetName = "RuneEvent"

var (
	MintRuneEventSig   = []byte("MintRune((string,uint128,string,uint8,string,uint128))")
	MintOutputEventSig = []byte("MintOutput((string,string,(uint128,uint128,string,uint8,string)))")
	MintErrEventSig    = []byte("MintError((string,uint128,uint128,uint128,string))")
	IssueEventSig      = []byte("Issue((string,uint8,uint64,uint128,uint256,string,uint128,string,uint32,uint128))")
	BurnRuneEventSig   = []byte("BurnRune((string,uint128,string,uint8,string,uint128))")
	BurnInputEventSig  = []byte("BurnInput((string,string,(uint128,uint128,string,uint8,string))")

	MintRuneTopic0   = crypto.Keccak256Hash(MintRuneEventSig).Hex()
	MintOutputTopic0 = crypto.Keccak256Hash(MintOutputEventSig).Hex()
	MintErrTopic0    = crypto.Keccak256Hash(MintErrEventSig).Hex()
	IssueTopic0      = crypto.Keccak256Hash(IssueEventSig).Hex()
	BurnRuneTopic0   = crypto.Keccak256Hash(BurnRuneEventSig).Hex()
	BurnInputTopic0  = crypto.Keccak256Hash(BurnInputEventSig).Hex()

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
