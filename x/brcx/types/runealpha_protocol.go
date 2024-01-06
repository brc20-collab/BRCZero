package types

import (
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

func UnpackMintRuneEvent(ret []byte) (MintRuneEvent, error) {
	var ec RuneAlphaWrappedEvent[MintRuneEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, MintRuneEventName, ret)
	if err != nil {
		return MintRuneEvent{}, err
	}
	return ec.Events, nil
}

func UnpackMintOutputEvent(ret []byte) (MintOutputEvent, error) {
	var ec RuneAlphaWrappedEvent[MintOutputEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, MintOutputEventName, ret)
	if err != nil {
		return MintOutputEvent{}, err
	}
	return ec.Events, nil
}

func UnpackMintErrEvent(ret []byte) (MintRuneErrEvent, error) {
	var ec RuneAlphaWrappedEvent[MintRuneErrEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, MintErrorEventName, ret)
	if err != nil {
		return MintRuneErrEvent{}, err
	}
	return ec.Events, nil
}

func UnpackIssueEvent(ret []byte) (IssueEventAdapter, error) {
	var ec RuneAlphaWrappedEvent[IssueEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, IssueEventName, ret)
	if err != nil {
		return IssueEventAdapter{}, err
	}
	return ec.Events.FormatResponse(), nil
}

func UnpackBurnRuneEvent(ret []byte) (BurnRuneEvent, error) {
	var ec RuneAlphaWrappedEvent[BurnRuneEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, BurnRuneEventName, ret)
	if err != nil {
		return BurnRuneEvent{}, err
	}
	return ec.Events, nil
}

func UnpackBurnInputEvent(ret []byte) (BurnInputEvent, error) {
	var ec RuneAlphaWrappedEvent[BurnInputEvent]
	err := runeAlphaABI.UnpackIntoInterface(&ec, BurnInputEventName, ret)
	if err != nil {
		return BurnInputEvent{}, err
	}
	return ec.Events, nil
}
