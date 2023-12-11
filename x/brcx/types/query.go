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
	Tick        string         `json:"tick"`
	TickAddress common.Address `json:"tickAddress"`
	MaxSupply   *big.Int       `json:"maxSupply"`
	NowSupply   *big.Int       `json:"nowSupply"`
	Decimals    *big.Int       `json:"decimals"`
	Lim         *big.Int       `json:"lim"`
}

type WrappedBRC20Information struct {
	BRC20Information
}

type QueryTickInfoResponse struct {
	Tick        string `json:"tick" yaml:"tick"`
	TickAddress string `json:"tick_address" yaml:"tick_address"`
	MaxSupply   string `json:"max_supply" yaml:"max_supply"`
	NowSupply   string `json:"now_supply" yaml:"now_supply"`
	Decimals    string `json:"decimals" yaml:"decimals"`
	Lim         string `json:"lim" yaml:"lim"`
}

func NewQueryTickInfoResponse(info WrappedBRC20Information) QueryTickInfoResponse {
	return QueryTickInfoResponse{
		Tick:        info.Tick,
		TickAddress: info.TickAddress.String(),
		MaxSupply:   info.MaxSupply.String(),
		NowSupply:   info.NowSupply.String(),
		Decimals:    info.Decimals.String(),
		Lim:         info.Lim.String(),
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
