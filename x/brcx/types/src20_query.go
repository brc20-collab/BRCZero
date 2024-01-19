package types

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

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
		Max:        fmt.Sprintf("%d", we.Max),
		Lim:        fmt.Sprintf("%d", we.Lim),
		Amt:        fmt.Sprintf("%d", we.Amt),
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

type Src20TokenInfoReq struct {
	Tick string `json:"tick" yaml:"tick"`
}

type QuerySrc20TickInfoResponse struct {
	Tick string `json:"tick" yaml:"tick"`
	//todo: stamp url is currently empty
	StampUrl   string `json:"stamp_url" yaml:"stamp_url"`
	Max        string `json:"max" yaml:"max"`
	Lim        string `json:"lim" yaml:"lim"`
	Amt        string `json:"amt" yaml:"amt"`
	Dec        string `json:"dec" yaml:"dec"`
	Creator    string `json:"creator" yaml:"creator"`
	TxHash     string `json:"tx_hash" yaml:"tx_hash"`
	BlockIndex string `json:"block_index" yaml:"block_index"`
}

type Src20TokenInfo struct {
	TickOrigin string
	Max        *big.Int
	Lim        *big.Int
	Amt        *big.Int
	Dec        uint8
	Creator    string
	TxHash     common.Hash
	BlockIndex *big.Int
	BlockTime  *big.Int
	Address    common.Address
}

type WrappedSrc20TokenInfo struct {
	Src20TokenInfo
}

func (info Src20TokenInfo) ToResponse() QuerySrc20TickInfoResponse {
	return QuerySrc20TickInfoResponse{
		Tick:       info.TickOrigin,
		StampUrl:   "",
		Max:        info.Max.String(),
		Lim:        info.Lim.String(),
		Amt:        info.Amt.String(),
		Dec:        strconv.FormatUint(uint64(info.Dec), 10),
		Creator:    info.Creator,
		TxHash:     info.TxHash.String(),
		BlockIndex: info.BlockIndex.String(),
	}
}

type Src20BalanceReq struct {
	Tick string `json:"tick" yaml:"tick"`
}

type Src20BalanceParams struct {
	Tick    string `json:"tick" yaml:"tick"`
	Address string `json:"address" yaml:"address"`
}

func NewSrc20BalanceParams(tick, address string) Src20BalanceParams {
	return Src20BalanceParams{
		Tick:    tick,
		Address: address,
	}
}

type Src20Balance struct {
	TickOrigin string
	Amt        *big.Int
	BlockIndex *big.Int
}

type QuerySrc20BalanceResponse struct {
	TickOrigin string `json:"tick" yaml:"tick"`
	Amt        string `json:"amt" yaml:"amt"`
}

type WrappedSrc20Balance struct {
	Src20Balance
}

func (b Src20Balance) ToResponse() QuerySrc20BalanceResponse {
	return QuerySrc20BalanceResponse{
		TickOrigin: b.TickOrigin,
		Amt:        b.Amt.String(),
	}
}

type QuerySrc20LatestBlockIndexResponse struct {
	BlockIndex int64  `json:"block_index" yaml:"block_index"`
	BlockTime  uint32 `json:"block_time" yaml:"block_time"`
	BlockHash  string `json:"block_hash" yaml:"block_hash"`
	PreHash    string `json:"pre_hash" yaml:"pre_hash"`
}

func NewQuerySrc20LatestBlockIndexResponse(index int64, hash string, preHash string) QuerySrc20LatestBlockIndexResponse {
	return QuerySrc20LatestBlockIndexResponse{
		BlockIndex: index,
		BlockTime:  0,
		BlockHash:  hash,
		PreHash:    preHash,
	}
}

type Src20AllTokenInfoReq struct {
	Page  string `json:"page" yaml:"page"`
	Limit string `json:"limit" yaml:"limit"`
}

type Src20AllTokenInfo struct {
	AllTokenInfo []Src20TokenInfo
	CurPage      *big.Int
	PageSize     *big.Int
	TotalRecords *big.Int
}

type WrappedSrc20AllTokenInfo struct {
	Src20AllTokenInfo
}

func (info Src20AllTokenInfo) ToResponse() QuerySrc20AllTickInfoResponse {

	allInfo := make([]QuerySrc20TickInfoResponse, 0)
	for _, i := range info.AllTokenInfo {
		allInfo = append(allInfo, i.ToResponse())
	}

	return QuerySrc20AllTickInfoResponse{
		AllTokenInfo: allInfo,
		CurPage:      info.CurPage,
		PageSize:     info.PageSize,
		TotalRecords: info.TotalRecords,
	}
}

type QuerySrc20AllTickInfoResponse struct {
	AllTokenInfo []QuerySrc20TickInfoResponse
	CurPage      *big.Int
	PageSize     *big.Int
	TotalRecords *big.Int
}
