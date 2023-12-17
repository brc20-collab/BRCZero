package types

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	authtypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/types"
	"github.com/brc20-collab/brczero/libs/tendermint/types"
)

// verify interface at compile time
var _ sdk.Msg = &MsgProtocolOp{}

type MsgProtocolOp struct {
	// RawData represents the raw data of protocol operations on the chain.
	RawData string `json:"raw_data" yaml:"raw_data"`
	// Context represents the contextual data required for protocol operation execution.
	Context string `json:"inscription_context" yaml:"inscriptionContext"`
}

// NewMsgUnjail creates a new MsgUnjail instance
func NewMsgProtocolOp(data string, ctx string) MsgProtocolOp {
	return MsgProtocolOp{
		RawData: data,
		Context: ctx,
	}
}

// nolint
func (msg MsgProtocolOp) Route() string { return RouterKey }
func (msg MsgProtocolOp) Type() string  { return MsgProtocolOpType }
func (msg MsgProtocolOp) GetSigners() []sdk.AccAddress {
	return nil
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgProtocolOp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgProtocolOp) ValidateBasic() error {
	return nil
}

// Decoder Try to decode as MsgInscription by json
func Decoder(_ codec.CdcAbstraction, txBytes []byte) (tx sdk.Tx, err error) {
	var zeroTx types.ZeroRequestTx

	if err = rlp.DecodeBytes(txBytes, &zeroTx); err == nil {
		var msgOp MsgProtocolOp
		if err = json.Unmarshal([]byte(zeroTx.TxInfo), &msgOp); err == nil {
			// TODO 1000 is tmp
			fee := authtypes.NewStdFee(50000000, nil)
			return authtypes.NewStdTx([]sdk.Msg{msgOp}, fee, nil, ""), nil
		}
	}

	return nil, ErrValidateInput(fmt.Sprintf("basicx msg deocer failed: %s", err))
}
