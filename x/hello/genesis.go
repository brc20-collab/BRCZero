package hello

import sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {

}

func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	return NewGenesisState()
}
