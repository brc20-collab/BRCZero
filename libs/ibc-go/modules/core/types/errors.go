package types

import (
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
	host "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/24-host"
)

var ErrIbcDisabled = sdkerrors.Register(host.ModuleName, 1, "IBC are disabled")

var ErrMisSpecificKeeper = sdkerrors.Register(host.ModuleName, 2, "asd")

var ErrInternalConfigError = sdkerrors.Register(host.ModuleName, 3, "asd")
