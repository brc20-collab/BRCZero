package keeper

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	govtypes "github.com/brc20-collab/brczero/x/gov/types"
)

// GovKeeper defines the expected gov Keeper
type GovKeeper interface {
	GetDepositParams(ctx sdk.Context) govtypes.DepositParams
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}
