package types

import (
	"fmt"
	"math/big"
)

type Src20EventContext struct {
	Op         string   `json:"op" yaml:"op"`
	TickOrigin string   `json:"tickOrigin" yaml:"tickOrigin"`
	Max        *big.Int `json:"max" yaml:"max"`
	Lim        *big.Int `json:"lim" yaml:"lim"`
	Amt        *big.Int `json:"amt" yaml:"amt"`
	Dec        uint8    `json:"dec" yaml:"dec"`
	From       string   `json:"from" yaml:"from"`
	To         string   `json:"to" yaml:"to"`
	Valid      bool     `json:"valid" yaml:"valid"`
	Msg        string   `json:"msg" yaml:"msg"`
}

type Src20EventResponse struct {
	Op         string `json:"op" yaml:"op"`
	TickOrigin string `json:"tick" yaml:"tick"`
	Max        string `json:"max" yaml:"max"`
	Lim        string `json:"lim" yaml:"lim"`
	Amt        string `json:"amt" yaml:"amt"`
	Dec        string `json:"dec" yaml:"dec"`
	From       string `json:"from" yaml:"from"`
	To         string `json:"to" yaml:"to"`
	Valid      bool   `json:"valid" yaml:"valid"`
	Msg        string `json:"msg" yaml:"msg"`
}

type Src20WrappedEvent struct {
	Src20EventContext `json:"events" yaml:"events"`
}

func (we Src20WrappedEvent) ToEventResponse() Src20EventResponse {
	return Src20EventResponse{
		Op:         we.Op,
		TickOrigin: we.TickOrigin,
		Max:        we.Max.String(),
		Lim:        we.Lim.String(),
		Amt:        we.Amt.String(),
		Dec:        fmt.Sprintf("%d", we.Dec),
		From:       we.From,
		To:         we.To,
		//todo: judge valid field
		Valid: true,
		Msg:   we.Msg,
	}
}

type QuerySrc20TxEventsResponse struct {
	Events []Src20EventResponse `json:"events" yaml:"events"`
	Txid   string               `json:"tx_hash" yaml:"tx_hash"`
}

func NewQuerySrc20TxEventsResponse(e []Src20EventResponse, txid string) QuerySrc20TxEventsResponse {
	return QuerySrc20TxEventsResponse{
		Events: e,
		Txid:   txid,
	}
}

type QuerySrc20TxEventsByBlockHashResponse struct {
	BlockEvents []QuerySrc20TxEventsResponse `json:"block" yaml:"block"`
}

func NewQuerySrc20TxEventsByBlockHashResponse(se []QuerySrc20TxEventsResponse) QuerySrc20TxEventsByBlockHashResponse {
	return QuerySrc20TxEventsByBlockHashResponse{BlockEvents: se}
}
