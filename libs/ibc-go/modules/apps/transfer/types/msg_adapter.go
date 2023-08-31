package types

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
	"github.com/brc20-collab/brczero/libs/system"
)

// for denom convert wei to brc10 and reject brc10 direct
func (m *MsgTransfer) RulesFilter() (sdk.Msg, error) {
	if m.Token.Denom == sdk.DefaultBondDenom {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "ibc MsgTransfer not support "+system.Currency+" denom")
	}
	return m, nil
}
