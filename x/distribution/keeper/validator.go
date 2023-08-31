package keeper

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"

	"github.com/brc20-collab/brczero/x/staking/exported"
)

// initialize rewards for a new validator
func (k Keeper) initializeValidator(ctx sdk.Context, val exported.ValidatorI) {
	k.initializeValidatorDistrProposal(ctx, val)
	return
}
