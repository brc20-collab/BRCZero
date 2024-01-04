package arc20

import (
	"bytes"
	_ "embed"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var (
	arc20EntryPointABI abi.ABI
	//go:embed abi/arc20_entrypoint.json
	arc20EntryPointABIJson []byte
)

const (
	AtomicalEvent = "AtomicalEvent"
)

func init() {
	var err error
	arc20EntryPointABI, err = abi.JSON(bytes.NewReader(arc20EntryPointABIJson))
	if err != nil {
		panic(fmt.Errorf("json decode failed: %s", err.Error()))
	}
}

func UnpackArc20EventContext(ret []byte, topic common.Hash) (AtomicalEventParam, error) {
	event, err := arc20EntryPointABI.EventByID(topic)
	if err != nil {
		return AtomicalEventParam{}, ErrNotExpectEvent
	}
	if event.Name != AtomicalEvent {
		return AtomicalEventParam{}, ErrNotExpectEvent
	}
	var ec Arc20WrappedEvent
	err = arc20EntryPointABI.UnpackIntoInterface(&ec, AtomicalEvent, ret)
	if err != nil {
		return AtomicalEventParam{}, err
	}
	return ec.AtomicalEventParam, nil
}

type Arc20WrappedEvent struct {
	AtomicalEventParam `json:"events" yaml:"events"`
}

type Hash common.Hash

type AtomicalEventParam struct {
	TxHash               Hash       `json:"tx_hash"`
	AtomicalsId          string     `json:"atomicals_id"`
	AtomicalsType        string     `json:"atomicals_type"`
	Subtype              string     `json:"subtype"`
	AtomicalsFromIndexes []*big.Int `json:"atomicals_from_indexs"`
	AtomicalsFromValues  []*big.Int `json:"atomicals_from_values"`
	ToAddress            string     `json:"to_address"`
	ToValueSats          *big.Int   `json:"to_value_sats"`
	MintAmount           *big.Int   `json:"mint_amount"`
	MintTicker           string     `json:"mint_ticker"`
	RequestTicker        string     `json:"request_ticker"`
}

func (hash *Hash) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", hex.EncodeToString(hash[:]))), nil
}

type ARC20Response struct {
	Success  bool             `json:"success"`
	Response AtomicalResponse `json:"response"`
}

type AtomicalResponse struct {
	Result []AtomicalEventParam `json:"result"`
}
