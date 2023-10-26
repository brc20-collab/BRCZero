package keeper

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) {
	//do something
}

func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	//do something
	return []abci.ValidatorUpdate{}
}
