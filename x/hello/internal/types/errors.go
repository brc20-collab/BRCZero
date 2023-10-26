package types

import (
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
)

var (
	ErrSetHello = sdkerrors.Register(ModuleName, 1, "set hello are disabled")
)
