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

func (b BRC20Balance) ToResponse() QueryBalanceResponse {
	return QueryBalanceResponse{
		Tick:                b.Tick,
		AvailableBalance:    b.AvailableBalance.String(),
		TransferableBalance: b.TransferableBalance.String(),
		OverallBalance:      b.TotalBalance.String(),
	}
}

type QueryBalanceResponse struct {
	Tick                string `json:"tick" yaml:"tick"`
	AvailableBalance    string `json:"available_balance" yaml:"available_balance"`
	TransferableBalance string `json:"transferable_balance" yaml:"transferable_balance"`
	OverallBalance      string `json:"overall_balance" yaml:"overall_balance"`
}

func NewQueryBalanceResponse(tick string, ab string, tb string, ob string) QueryBalanceResponse {
	return QueryBalanceResponse{
		Tick:                tick,
		AvailableBalance:    ab,
		TransferableBalance: tb,
		OverallBalance:      ob,
	}
}

type QueryAllBalanceResponse struct {
	Balance []QueryBalanceResponse `json:"balance" yaml:"balance"`
}

func NewQueryAllBalanceResponse(balance []QueryBalanceResponse) QueryAllBalanceResponse {
	return QueryAllBalanceResponse{
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

func (info BRC20Information) ToResponse() QueryTickInfoResponse {
	return QueryTickInfoResponse{
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

type QueryTickInfoResponse struct {
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

func NewQueryTickInfoResponse(info WrappedBRC20Information) QueryTickInfoResponse {
	return QueryTickInfoResponse{
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

type QueryAllTickInfoResponse struct {
	Tokens []QueryTickInfoResponse `json:"tokens" yaml:"tokens"`
}

func NewQueryAllTickInfoResponse(infos []QueryTickInfoResponse) QueryAllTickInfoResponse {
	return QueryAllTickInfoResponse{
		Tokens: infos,
	}
}

type QueryTotalTickHoldersResponse struct {
	Holders string `json:"total_tick_holders" yaml:"total_tick_holders"`
}

func NewQueryTotalTickHoldersResponse(h string) QueryTotalTickHoldersResponse {
	return QueryTotalTickHoldersResponse{
		Holders: h,
	}
}

type EventContext struct {
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
	Msg               string   `json:"msg" yaml:"msg"`
	Txid              string   `json:"txid" yaml:"txid"`
}

type WrappedEvent struct {
	EventContext `json:"events" yaml:"events"`
}

type TransferableInscription struct {
	InscriptionId     string `json:"inscription_id" yaml:"inscription_id"`
	InscriptionNumber int64  `json:"inscription_number" yaml:"inscription_number"`
	Amount            string `json:"amount" yaml:"amount"`
	Tick              string `json:"tick" yaml:"tick"`
	Owner             string `json:"owner" yaml:"owner"`
}

type QueryTransferableInscriptionResponse struct {
	Inscriptions []TransferableInscription `json:"inscriptions" yaml:"inscriptions"`
}

func NewQueryTransferableInscriptionResponse(tis []TransferableInscription) QueryTransferableInscriptionResponse {
	return QueryTransferableInscriptionResponse{Inscriptions: tis}
}
