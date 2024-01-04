package types

import "math/big"

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
