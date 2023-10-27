package hello

import (
	"github.com/brc20-collab/brczero/x/hello/internal/keeper"
	"github.com/brc20-collab/brczero/x/hello/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
)

type (
	GenesisState = types.GenesisState
	Keeper       = keeper.Keeper
)

var (
	RegisterCodec       = types.RegisterCodec
	ModuleCdc           = types.ModuleCdc
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewKeeper           = keeper.NewKeeper
)
