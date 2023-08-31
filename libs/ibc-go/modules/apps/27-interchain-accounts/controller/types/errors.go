package types

import (
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
)

// ICA Controller sentinel errors
var (
	ErrControllerSubModuleDisabled = sdkerrors.Register(SubModuleName, 2, "controller submodule is disabled")
)
