package types

import (
	"fmt"
	"math/big"
)

type IssueEvent struct {
	Op           string   `json:"type"`
	Divisibility uint8    `json:"divisibility"`
	End          uint64   `json:"end"`
	Limit        *big.Int `json:"limit"`
	Number       *big.Int `json:"number"`
	Rune         string   `json:"rune"`
	Supply       *big.Int `json:"supply"`
	Symbol       string   `json:"symbol"`
	Timestamp    uint32   `json:"timestamp"`
	Id           *big.Int `json:"id"`
}

type BurnRuneEvent struct {
	Op           string   `json:"type"`
	Id           *big.Int `json:"id"`
	Rune         string   `json:"rune"`
	Divisibility uint8    `json:"divisibility"`
	Symbol       string   `json:"symbol"`
	Amount       *big.Int `json:"amount"`
}

type BurnInputEvent struct {
	Op string `json:"type"`
	//todo:input
	PreOutputId string `json:"preOutputId"`
	//todo:address
	Burn RuneBasicInfo `json:"burn"`
}

type MintRuneEvent struct {
	Op           string   `json:"type"`
	Id           *big.Int `json:"id"`
	Rune         string   `json:"rune"`
	Divisibility uint8    `json:"divisibility"`
	Symbol       string   `json:"symbol"`
	Amount       *big.Int `json:"amount"`
}

type MintOutputEvent struct {
	Op       string `json:"type"`
	OutputId string `json:"output"`
	//todo:address
	Mint RuneBasicInfo `json:"mint"`
}

type RuneBasicInfo struct {
	Id           *big.Int `json:"id"`
	Amount       *big.Int `json:"amount"`
	Rune         string   `json:"rune"`
	Divisibility uint8    `json:"divisibility"`
	Symbol       string   `json:"symbol"`
}

func (rbi *RuneBasicInfo) AddAmount(another RuneBasicInfo) error {
	if rbi.Id.Cmp(another.Id) != 0 ||
		rbi.Rune != another.Rune ||
		rbi.Divisibility != another.Divisibility ||
		rbi.Symbol != another.Symbol {
		return fmt.Errorf("RuneBasicInfo AddAmount wrong ")
	}
	rbi.Amount.Add(rbi.Amount, another.Amount)
	return nil
}

type MintRuneErrEvent struct {
	Op     string   `json:"type"`
	Id     *big.Int `json:"id"`
	Amount *big.Int `json:"amount"`
	Output *big.Int `json:"output"`
	Msg    string   `json:"msg"`
}

type RuneAlphaWrappedEvent[T interface{}] struct {
	Events T `json:"events" yaml:"events"`
}

type QueryRuneAlphaTxEventsResponse[T interface{}] struct {
	Events []T    `json:"events" yaml:"events"`
	Txid   string `json:"txid" yaml:"txid"`
}

func NewQueryRuneAlphaTxEventsResponse[T interface{}](e []T, txid string) QueryRuneAlphaTxEventsResponse[T] {
	return QueryRuneAlphaTxEventsResponse[T]{
		Events: e,
		Txid:   txid,
	}
}

type QueryRuneAlphaTxEventsByBlockHashResponse[T interface{}] struct {
	BlockEvents []T `json:"block" yaml:"block"`
}

func NewQueryRuneAlphaTxEventsByBlockHashResponse[T interface{}](se []T) QueryRuneAlphaTxEventsByBlockHashResponse[T] {
	return QueryRuneAlphaTxEventsByBlockHashResponse[T]{BlockEvents: se}
}

var limitNull *big.Int

func init() {
	limitNull = new(big.Int).SetUint64(0)
}

type IssueEventAdapter struct {
	Op           string   `json:"type"`
	Divisibility uint8    `json:"divisibility"`
	End          *big.Int `json:"end"`
	Limit        *big.Int `json:"limit"`
	Number       *big.Int `json:"number"`
	Rune         string   `json:"rune"`
	Supply       *big.Int `json:"supply"`
	Symbol       string   `json:"symbol"`
	Timestamp    uint32   `json:"timestamp"`
	Id           *big.Int `json:"id"`
}

func (ie IssueEvent) FormatResponse() IssueEventAdapter {
	var limit *big.Int
	if ie.Limit.CmpAbs(limitNull) > 0 {
		limit = ie.Limit
	}

	var end *big.Int
	if ie.End > 0 {
		end = new(big.Int).SetUint64(ie.End)
	}

	return IssueEventAdapter{
		Op:           ie.Op,
		Divisibility: ie.Divisibility,
		End:          end,
		Limit:        limit,
		Number:       ie.Number,
		Rune:         ie.Rune,
		Supply:       ie.Supply,
		Symbol:       ie.Symbol,
		Timestamp:    ie.Timestamp,
		Id:           ie.Id,
	}
}

type MintOutputKey struct {
	Txid     string
	Op       string
	OutputId string
}

func NewMintOutputKey(txid, op, output string) MintOutputKey {
	return MintOutputKey{
		Txid:     txid,
		Op:       op,
		OutputId: output,
	}
}

type MintOutputEventAdapter struct {
	Op       string `json:"type"`
	OutputId string `json:"output"`
	//todo:address
	Mint []RuneBasicInfo `json:"mint"`
}

func NewRawMintOutputEventAdapter(op string, outpointId string) MintOutputEventAdapter {
	return MintOutputEventAdapter{
		Op:       op,
		OutputId: outpointId,
		Mint:     make([]RuneBasicInfo, 0),
	}
}

type BurnInputKey struct {
	Txid        string
	Op          string
	PreOutputId string
}

func NewBurnInputKey(txid, op, preOutput string) BurnInputKey {
	return BurnInputKey{
		Txid:        txid,
		Op:          op,
		PreOutputId: preOutput,
	}
}

type BurnInputEventAdapter struct {
	Op string `json:"type"`
	//todo:input
	PreOutputId string `json:"preOutputId"`
	//todo:address
	Burn []RuneBasicInfo `json:"burn"`
}

func NewRawBurnInputEventAdapter(op string, preOutpointId string) BurnInputEventAdapter {
	return BurnInputEventAdapter{
		Op:          op,
		PreOutputId: preOutpointId,
		Burn:        make([]RuneBasicInfo, 0),
	}
}
