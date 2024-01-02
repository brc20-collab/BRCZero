package types

import "math/big"

type RuneAlphaEventContext struct {
	Op           string
	Rune         string
	Symbol       string
	Id           *big.Int
	Divisibility uint8
	Limit        *big.Int
	Amount       *big.Int
	Supply       *big.Int
	End          uint64
	Number       *big.Int
	OutputId     string
	PreOutputId  string
	Timestamp    uint32
}

type RuneAlphaWrappedEvent struct {
	RuneAlphaEventContext `json:"events" yaml:"events"`
}

type RuneAlphaEventResponse struct {
	Op           string   `json:"type" yaml:"type"`
	Rune         string   `json:"rune,omitempty" yaml:"rune,omitempty"`
	Symbol       string   `json:"symbol,omitempty" yaml:"symbol,omitempty"`
	Id           *big.Int `json:"id" yaml:"id"`
	Divisibility uint8    `json:"divisibility,omitempty" yaml:"divisibility,omitempty"`
	Limit        *big.Int `json:"limit,omitempty" yaml:"limit,omitempty"`
	Amount       *big.Int `json:"amount,omitempty" yaml:"amount,omitempty"`
	Supply       *big.Int `json:"supply,omitempty" yaml:"supply,omitempty"`
	End          uint64   `json:"end,omitempty" yaml:"end,omitempty"`
	Number       *big.Int `json:"number,omitempty" yaml:"number,omitempty"`
	OutputId     string   `json:"output,omitempty" yaml:"output,omitempty"`
	PreOutputId  string   `json:"input,omitempty" yaml:"input,omitempty"`
	Timestamp    uint32   `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
}

func (re RuneAlphaWrappedEvent) ToEventResponse() RuneAlphaEventResponse {
	return RuneAlphaEventResponse{
		Op:           re.Op,
		Rune:         re.Rune,
		Symbol:       re.Symbol,
		Id:           re.Id,
		Divisibility: re.Divisibility,
		Limit:        re.Limit,
		Amount:       re.Amount,
		Supply:       re.Supply,
		End:          re.End,
		Number:       re.Number,
		OutputId:     re.OutputId,
		PreOutputId:  re.PreOutputId,
		Timestamp:    re.Timestamp,
	}
}

type QueryRuneAlphaTxEventsResponse struct {
	Events []RuneAlphaEventResponse `json:"events" yaml:"events"`
	Txid   string                   `json:"txid" yaml:"txid"`
}

func NewQueryRuneAlphaTxEventsResponse(e []RuneAlphaEventResponse, txid string) QueryRuneAlphaTxEventsResponse {
	return QueryRuneAlphaTxEventsResponse{
		Events: e,
		Txid:   txid,
	}
}

type QueryRuneAlphaTxEventsByBlockHashResponse struct {
	BlockEvents []QueryRuneAlphaTxEventsResponse `json:"block" yaml:"block"`
}

func NewQueryRuneAlphaTxEventsByBlockHashResponse(se []QueryRuneAlphaTxEventsResponse) QueryRuneAlphaTxEventsByBlockHashResponse {
	return QueryRuneAlphaTxEventsByBlockHashResponse{BlockEvents: se}
}
