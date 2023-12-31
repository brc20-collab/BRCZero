package types

import (
	sdkgenutiltypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/genutil/types"
)

// const
const (
	ModuleName = sdkgenutiltypes.ModuleName
)

type (
	// AccountKeeper is the type alias of the one in cmsdk
	AccountKeeper = sdkgenutiltypes.AccountKeeper
	// StakingKeeper is the type alias of the one in cmsdk
	StakingKeeper = sdkgenutiltypes.StakingKeeper
)

var (
	NewInitConfig = sdkgenutiltypes.NewInitConfig
)
