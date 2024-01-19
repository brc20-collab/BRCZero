package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BRC20Balance struct {
	Tick                string
	TotalBalance        *big.Int
	AvailableBalance    *big.Int
	TransferableBalance *big.Int
}

func (b BRC20Balance) ToResponse() QueryBrc20BalanceResponse {
	return QueryBrc20BalanceResponse{
		Tick:                b.Tick,
		AvailableBalance:    b.AvailableBalance.String(),
		TransferableBalance: b.TransferableBalance.String(),
		OverallBalance:      b.TotalBalance.String(),
	}
}

type QueryBrc20BalanceResponse struct {
	Tick                string `json:"tick" yaml:"tick"`
	AvailableBalance    string `json:"available_balance" yaml:"available_balance"`
	TransferableBalance string `json:"transferable_balance" yaml:"transferable_balance"`
	OverallBalance      string `json:"overall_balance" yaml:"overall_balance"`
}

func NewQueryBrc20BalanceResponse(tick string, ab string, tb string, ob string) QueryBrc20BalanceResponse {
	return QueryBrc20BalanceResponse{
		Tick:                tick,
		AvailableBalance:    ab,
		TransferableBalance: tb,
		OverallBalance:      ob,
	}
}

type QueryBrc20AllBalanceResponse struct {
	Balance []QueryBrc20BalanceResponse `json:"balance" yaml:"balance"`
}

func NewQueryBrc20AllBalanceResponse(balance []QueryBrc20BalanceResponse) QueryBrc20AllBalanceResponse {
	return QueryBrc20AllBalanceResponse{
		Balance: balance,
	}
}

type BRC20Information struct {
	Tick              string
	TickAddress       common.Address
	MaxSupply         *big.Int
	NowSupply         *big.Int
	Decimals          *big.Int
	Lim               *big.Int
	InscriptionId     string
	InscriptionNumber int64
	Txid              string
	Sender            string
	BlockTime         uint32
	BlockHeight       uint64
}

func (info BRC20Information) ToResponse() QueryBrc20TickInfoResponse {
	return QueryBrc20TickInfoResponse{
		Tick:              info.Tick,
		InscriptionId:     info.InscriptionId,
		InscriptionNumber: info.InscriptionNumber,
		MaxSupply:         info.MaxSupply.String(),
		Lim:               info.Lim.String(),
		NowSupply:         info.NowSupply.String(),
		Decimals:          info.Decimals.String(),
		Sender:            info.Sender,
		Txid:              info.Txid,
		BlockHeight:       info.BlockHeight,
		BlockTime:         info.BlockTime,
	}
}

type WrappedBRC20Information struct {
	BRC20Information
}

type QueryBrc20TickInfoResponse struct {
	Tick              string `json:"tick" yaml:"tick"`
	InscriptionId     string `json:"inscription_id" yaml:"inscription_id"`
	InscriptionNumber int64  `json:"inscription_number" yaml:"inscription_number"`
	MaxSupply         string `json:"supply" yaml:"supply"`
	Lim               string `json:"limit_per_mint" yaml:"limit_per_mint"`
	NowSupply         string `json:"minted" yaml:"minted"`
	Decimals          string `json:"decimal" yaml:"decimal"`
	Sender            string `json:"deploy_by" yaml:"deploy_by"`
	Txid              string `json:"tx_id" yaml:"tx_id"`
	BlockHeight       uint64 `json:"deploy_height" yaml:"deploy_height"`
	BlockTime         uint32 `json:"deploy_blocktime" yaml:"deploy_blocktime"`
}

func NewQueryBrc20TickInfoResponse(info WrappedBRC20Information) QueryBrc20TickInfoResponse {
	return QueryBrc20TickInfoResponse{
		Tick:              info.Tick,
		InscriptionId:     info.InscriptionId,
		InscriptionNumber: info.InscriptionNumber,
		MaxSupply:         info.MaxSupply.String(),
		Lim:               info.Lim.String(),
		NowSupply:         info.NowSupply.String(),
		Decimals:          info.Decimals.String(),
		Sender:            info.Sender,
		Txid:              info.Txid,
		BlockHeight:       info.BlockHeight,
		BlockTime:         info.BlockTime,
	}
}

type QueryBrc20AllTickInfoResponse struct {
	Tokens []QueryBrc20TickInfoResponse `json:"tokens" yaml:"tokens"`
}

func NewQueryBrc20AllTickInfoResponse(infos []QueryBrc20TickInfoResponse) QueryBrc20AllTickInfoResponse {
	return QueryBrc20AllTickInfoResponse{
		Tokens: infos,
	}
}

type QueryBrc20TotalTickHoldersResponse struct {
	Holders string `json:"total_tick_holders" yaml:"total_tick_holders"`
}

func NewQueryBrc20TotalTickHoldersResponse(h string) QueryBrc20TotalTickHoldersResponse {
	return QueryBrc20TotalTickHoldersResponse{
		Holders: h,
	}
}

type Brc20EventContext struct {
	EventType         string   `json:"event" yaml:"event"`
	Tick              string   `json:"tick" yaml:"tick"`
	InscriptionId     string   `json:"inscription_id" yaml:"inscription_id"`
	InscriptionNumber int64    `json:"inscription_number" yaml:"inscription_number"`
	OldSatPoint       string   `json:"old_satpoint" yaml:"old_satpoint"`
	NewSatPoint       string   `json:"new_satpoint" yaml:"new_satpoint"`
	Supply            *big.Int `json:"supply" yaml:"supply"`
	Lim               *big.Int `json:"lim_per_mint" yaml:"lim_per_mint"`
	Dec               *big.Int `json:"decimals" yaml:"decimals"`
	Amount            *big.Int `json:"amount" yaml:"amount"`
	Sender            string   `json:"from" yaml:"from"`
	Receiver          string   `json:"to" yaml:"to"`
	Msg               string   `json:"msg,omitempty" yaml:"msg,omitempty"`
	Txid              string   `json:"txid" yaml:"txid"`
}

type Brc20EventResponse struct {
	EventType         string   `json:"event" yaml:"event"`
	Tick              string   `json:"tick" yaml:"tick"`
	InscriptionId     string   `json:"inscription_id" yaml:"inscription_id"`
	InscriptionNumber int64    `json:"inscription_number" yaml:"inscription_number"`
	OldSatPoint       string   `json:"old_satpoint" yaml:"old_satpoint"`
	NewSatPoint       string   `json:"new_satpoint" yaml:"new_satpoint"`
	Supply            string   `json:"supply,omitempty" yaml:"supply,omitempty"`
	Lim               string   `json:"lim_per_mint,omitempty" yaml:"lim_per_mint,omitempty"`
	Dec               uint64   `json:"decimals,omitempty" yaml:"decimals,omitempty"`
	Amount            *big.Int `json:"amount" yaml:"amount"`
	Sender            string   `json:"from" yaml:"from"`
	Receiver          string   `json:"to" yaml:"to"`
	Valid             bool     `json:"valid" yaml:"valid"`
	Msg               string   `json:"msg" yaml:"msg"`
}

type Brc20WrappedEvent struct {
	Brc20EventContext `json:"events" yaml:"events"`
}

func (we Brc20WrappedEvent) ToEventResponse() Brc20EventResponse {
	return Brc20EventResponse{
		EventType:         we.EventType,
		Tick:              we.Tick,
		InscriptionId:     we.InscriptionId,
		InscriptionNumber: we.InscriptionNumber,
		OldSatPoint:       we.OldSatPoint,
		NewSatPoint:       we.NewSatPoint,
		Supply:            we.Supply.String(),
		Lim:               we.Lim.String(),
		Dec:               we.Dec.Uint64(),
		Amount:            we.Amount,
		Sender:            we.Sender,
		Receiver:          we.Receiver,
		//todo: judge valid field
		Valid: true,
		Msg:   we.Msg,
	}
}

type QueryBrc20TxEventsResponse struct {
	Events []Brc20EventResponse `json:"events" yaml:"events"`
	Txid   string               `json:"txid" yaml:"txid"`
}

func NewQueryBrc20TxEventsResponse(e []Brc20EventResponse, txid string) QueryBrc20TxEventsResponse {
	return QueryBrc20TxEventsResponse{
		Events: e,
		Txid:   txid,
	}
}

type QueryBrc20TxEventsByBlockHashResponse struct {
	BlockEvents []QueryBrc20TxEventsResponse `json:"block" yaml:"block"`
}

func NewQueryBrc20TxEventsByBlockHashResponse(be []QueryBrc20TxEventsResponse) QueryBrc20TxEventsByBlockHashResponse {
	return QueryBrc20TxEventsByBlockHashResponse{BlockEvents: be}
}

type Brc20TransferableInscription struct {
	InscriptionId     string `json:"inscription_id" yaml:"inscription_id"`
	InscriptionNumber int64  `json:"inscription_number" yaml:"inscription_number"`
	Amount            string `json:"amount" yaml:"amount"`
	Tick              string `json:"tick" yaml:"tick"`
	Owner             string `json:"owner" yaml:"owner"`
}

type QueryBrc20TransferableInscriptionResponse struct {
	Inscriptions []Brc20TransferableInscription `json:"inscriptions" yaml:"inscriptions"`
}

func NewQueryBrc20TransferableInscriptionResponse(tis []Brc20TransferableInscription) QueryBrc20TransferableInscriptionResponse {
	return QueryBrc20TransferableInscriptionResponse{Inscriptions: tis}
}

type Brc20InscriptionInfo struct {
	Action            string         `json:"action,omitempty" yaml:"action,omitempty"`
	InscriptionNumber int64          `json:"inscription_number,omitempty" yaml:"inscription_number,omitempty"`
	InscriptionId     string         `json:"inscription_id" yaml:"inscription_id"`
	From              string         `json:"from" yaml:"from"`
	To                string         `json:"to,omitempty" yaml:"to,omitempty"`
	OldSatPoint       string         `json:"old_satpoint,omitempty" yaml:"old_satpoint,omitempty"`
	NewSatPoint       string         `json:"new_satpoint,omitempty" yaml:"new_satpoint,omitempty"`
	Operation         Brc20Operation `json:"operation,omitempty" yaml:"operation,omitempty"`
}

type Brc20Operation struct {
	Tick string `json:"tick" yaml:"tick"`
	Amt  string `json:"amt,omitempty" yaml:"amt,omitempty"`
	Max  string `json:"max,omitempty" yaml:"max,omitempty"`
	Lim  string `json:"lim,omitempty" yaml:"lim,omitempty"`
	Dec  string `json:"dec,omitempty" yaml:"dec,omitempty"`
}
