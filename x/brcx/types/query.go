package types

import (
	"math/big"
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
