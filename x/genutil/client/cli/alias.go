package cli

import (
	genutilcli "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/genutil/client/cli"
)

type (
	stakingMsgBuildingHelpers = genutilcli.StakingMsgBuildingHelpers
)

var (
	// nolint
	ValidateGenesisCmd = genutilcli.ValidateGenesisCmd
)
