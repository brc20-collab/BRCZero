package capability

import (
	"encoding/json"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	capabilityModule "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/capability"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/capability/keeper"
	types2 "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/capability/types"
	"github.com/brc20-collab/brczero/libs/ibc-go/testing/simapp/adapter"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
)

type CapabilityModuleAdapter struct {
	capabilityModule.AppModule

	tkeeper keeper.Keeper
}

func TNewCapabilityModuleAdapter(cdc *codec.CodecProxy, keeper keeper.Keeper) *CapabilityModuleAdapter {
	ret := &CapabilityModuleAdapter{}
	ret.AppModule = capabilityModule.NewAppModule(cdc, keeper)
	ret.tkeeper = keeper
	return ret
}

func (a CapabilityModuleAdapter) DefaultGenesis() json.RawMessage {
	return adapter.ModuleCdc.MustMarshalJSON(types2.DefaultGenesis())
}
func (am CapabilityModuleAdapter) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return am.initGenesis(ctx, data)
}

func (am CapabilityModuleAdapter) initGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genState types2.GenesisState
	// Initialize global index to index in genesis state

	adapter.ModuleCdc.MustUnmarshalJSON(data, &genState)

	capabilityModule.InitGenesis(ctx, am.tkeeper, genState)

	return []abci.ValidatorUpdate{}
}
