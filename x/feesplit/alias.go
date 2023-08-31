package feesplit

import (
	"github.com/brc20-collab/brczero/x/feesplit/keeper"
	"github.com/brc20-collab/brczero/x/feesplit/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

var (
	NewKeeper           = keeper.NewKeeper
	SetParamsNeedUpdate = types.SetParamsNeedUpdate
)

type (
	Keeper = keeper.Keeper
)
