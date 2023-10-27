package types

import (
	"github.com/tendermint/go-amino"

	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
)

const (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

type MsgKV struct {
	FromAddress sdk.AccAddress `json:"from_address" yaml:"from_address"`
	Key         string         `json:"key" yaml:"key"`
	Value       string         `json:"value" yaml:"value"`
}

var _ sdk.Msg = MsgKV{}

func NewMsgHello(fromAddr sdk.AccAddress, key string, val string) MsgKV {
	return MsgKV{FromAddress: fromAddr, Key: key, Value: val}
}

func (msg MsgKV) Route() string { return RouterKey }

func (msg MsgKV) Type() string { return "hello" }

func (msg MsgKV) ValidateBasic() error {
	if msg.Value == "hello" {
		return sdkerrors.Wrapf(ErrSetHello, "what you set is %s", msg.Value)
	}
	return nil
}

func (msg MsgKV) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgKV) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

func (msg *MsgKV) UnmarshalFromAmino(cdc *amino.Codec, data []byte) error {
	return cdc.UnmarshalBinaryBare(data, &msg)
}
