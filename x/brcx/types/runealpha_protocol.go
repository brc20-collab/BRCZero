package types

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	IssueEventName      = "Issue"
	BurnRuneEventName   = "BurnRune"
	BurnInputEventName  = "BurnInput"
	MintErrorEventName  = "MintError"
	MintRuneEventName   = "MintRune"
	MintOutputEventName = "MintOutput"
)

var (
	MintRuneEventSig   = []byte("MintRune((string,uint128,string,uint8,string,uint128))")
	MintOutputEventSig = []byte("MintOutput((string,string,(uint128,uint128,string,uint8,string)))")
	MintErrEventSig    = []byte("MintError((string,uint128,uint128,uint128,string))")
	IssueEventSig      = []byte("Issue((string,uint8,uint64,uint128,uint256,string,uint128,string,uint32,uint128))")
	BurnRuneEventSig   = []byte("BurnRune((string,uint128,string,uint8,string,uint128))")
	BurnInputEventSig  = []byte("BurnInput((string,string,(uint128,uint128,string,uint8,string))")

	MintRuneTopic0   = crypto.Keccak256Hash(MintRuneEventSig)
	MintOutputTopic0 = crypto.Keccak256Hash(MintOutputEventSig)
	MintErrTopic0    = crypto.Keccak256Hash(MintErrEventSig)
	IssueTopic0      = crypto.Keccak256Hash(IssueEventSig)
	BurnRuneTopic0   = crypto.Keccak256Hash(BurnRuneEventSig)
	BurnInputTopic0  = crypto.Keccak256Hash(BurnInputEventSig)

	// TopicEventMap store topic -> contract eventName
	TopicEventMap = make(map[common.Hash]func([]byte) (interface{}, error))

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

	TopicEventMap[MintRuneTopic0] = UnpackMintRuneEvent
	TopicEventMap[MintOutputTopic0] = UnpackMintOutputEvent
	TopicEventMap[MintErrTopic0] = UnpackMintErrEvent
	TopicEventMap[IssueTopic0] = UnpackIssueEvent
	TopicEventMap[BurnRuneTopic0] = UnpackBurnRuneEvent
	TopicEventMap[BurnInputTopic0] = UnpackBurnInputEvent

}

func UnpackMintRuneEvent(ret []byte) (interface{}, error) {
	var ec RuneAlphaWrappedEvent[MintRuneEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, MintRuneEventName, ret)
	if err != nil {
		return RuneAlphaWrappedEvent[MintRuneEvent]{}, err
	}
	return ec.Events, nil
}

func UnpackMintOutputEvent(ret []byte) (interface{}, error) {
	var ec RuneAlphaWrappedEvent[MintOutputEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, MintOutputEventName, ret)
	if err != nil {
		return RuneAlphaWrappedEvent[MintOutputEvent]{}, err
	}
	return ec.Events, nil
}

func UnpackMintErrEvent(ret []byte) (interface{}, error) {
	var ec RuneAlphaWrappedEvent[MintRuneErrEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, MintErrorEventName, ret)
	if err != nil {
		return RuneAlphaWrappedEvent[MintRuneErrEvent]{}, err
	}
	return ec.Events, nil
}

func UnpackIssueEvent(ret []byte) (interface{}, error) {
	var ec RuneAlphaWrappedEvent[IssueEvent]
	err := runeAlphaEntryPointABI.UnpackIntoInterface(&ec, IssueEventName, ret)
	if err != nil {
		return RuneAlphaWrappedEvent[IssueEvent]{}, err
	}
	return ec.Events, nil
}

func UnpackBurnRuneEvent(ret []byte) (interface{}, error) {
	var ec RuneAlphaWrappedEvent[BurnRuneEvent]
	err := runeAlphaEntryPointABI.UnpackIntoInterface(&ec, BurnRuneEventName, ret)
	if err != nil {
		return RuneAlphaWrappedEvent[BurnRuneEvent]{}, err
	}
	return ec.Events, nil
}

func UnpackBurnInputEvent(ret []byte) (interface{}, error) {
	var ec RuneAlphaWrappedEvent[BurnInputEvent]
	err := runeAlphaEntryPointABI.UnpackIntoInterface(&ec, BurnInputEventName, ret)
	if err != nil {
		return RuneAlphaWrappedEvent[BurnInputEvent]{}, err
	}
	return ec.Events, nil
}
