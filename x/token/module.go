package token

import (
	"encoding/json"

	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/module"
	authTypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/types"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"

	"github.com/brc20-collab/brczero/x/common/version"
	tokenTypes "github.com/brc20-collab/brczero/x/token/types"
)

var (
	_ module.AppModule = AppModule{}
)

// AppModule app module
type AppModule struct {
	AppModuleBasic
	keeper       Keeper
	supplyKeeper authTypes.SupplyKeeper
	version      version.ProtocolVersionType
}

// NewAppModule creates a new AppModule object
func NewAppModule(v version.ProtocolVersionType, keeper Keeper, supplyKeeper authTypes.SupplyKeeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
		supplyKeeper:   supplyKeeper,
		version:        v,
	}
}

// nolint
func (AppModule) Name() string {
	return tokenTypes.ModuleName
}

// nolint
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
}

// Route module message route name
func (AppModule) Route() string {
	return tokenTypes.RouterKey
}

// nolint
func (am AppModule) NewHandler() sdk.Handler {
	return NewTokenHandler(am.keeper, am.version)
}

// nolint
func (AppModule) QuerierRoute() string {
	return tokenTypes.QuerierRoute
}

// nolint
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// nolint
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	tokenTypes.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	initGenesis(ctx, am.keeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// nolint
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return tokenTypes.ModuleCdc.MustMarshalJSON(gs)
}

// nolint
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	beginBlocker(ctx, am.keeper)
}

// nolint
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
