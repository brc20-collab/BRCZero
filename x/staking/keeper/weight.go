package keeper

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/x/staking/types"
)

func calculateWeight(tokens sdk.Dec) types.Shares {
	return tokens
}

func SimulateWeight(tokens sdk.Dec) types.Shares {
	return calculateWeight(tokens)
}
