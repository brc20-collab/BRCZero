package types

import "fmt"

type Src20EventContext struct {
	Op         string `json:"op" yaml:"op"`
	TickOrigin string `json:"tickOrigin" yaml:"tickOrigin"`
	Max        uint64 `json:"max" yaml:"max"`
	Lim        uint64 `json:"lim" yaml:"lim"`
	Amt        uint64 `json:"amt" yaml:"amt"`
	Dec        uint8  `json:"dec" yaml:"dec"`
	From       string `json:"from" yaml:"from"`
	To         string `json:"to" yaml:"to"`
	Valid      bool   `json:"valid" yaml:"valid"`
	Msg        string `json:"msg" yaml:"msg"`
}

type Src20EventResponse struct {
	Op         string `json:"type" yaml:"type"`
	TickOrigin string `json:"tick" yaml:"tick"`
	Max        string `json:"max" yaml:"max"`
	Lim        string `json:"lim" yaml:"lim"`
	Amt        string `json:"amt" yaml:"amt"`
	Dec        uint8  `json:"dec" yaml:"dec"`
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
		Max:        fmt.Sprintf("%d", we.Max),
		Lim:        fmt.Sprintf("%d", we.Lim),
		Amt:        fmt.Sprintf("%d", we.Amt),
		Dec:        we.Dec,
		From:       we.From,
		To:         we.To,
		//todo: judge valid field
		Valid: true,
		Msg:   we.Msg,
	}
}

type QuerySrc20TxEventsResponse struct {
	Events []Src20EventResponse `json:"events" yaml:"events"`
	Txid   string               `json:"txid" yaml:"txid"`
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
