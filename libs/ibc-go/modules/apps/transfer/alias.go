package transfer

import (
	"github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/transfer/keeper"
	"github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/transfer/types"
)

var (
	NewKeeper  = keeper.NewKeeper
	ModuleCdc  = types.ModuleCdc
	SetMarshal = types.SetMarshal
)
