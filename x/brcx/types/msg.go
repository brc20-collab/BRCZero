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
var _ sdk.Msg = &MsgInscription{}

// MsgInscription - struct for create contract
type MsgInscription struct {
	Inscription        string             `json:"inscription" yaml:"inscription"`
	InscriptionContext InscriptionContext `json:"inscription_context" yaml:"inscriptionContext"`
}

// NewMsgUnjail creates a new MsgUnjail instance
func NewMsgInscription(Inscription string, ctx InscriptionContext) MsgInscription {
	return MsgInscription{
		Inscription:        Inscription,
		InscriptionContext: ctx,
	}
}

// nolint
func (msg MsgInscription) Route() string { return RouterKey }
func (msg MsgInscription) Type() string  { return MsgInscriptionType }
func (msg MsgInscription) GetSigners() []sdk.AccAddress {
	return nil
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgInscription) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgInscription) ValidateBasic() error {

	return nil
}

// verify interface at compile time
var _ sdk.Msg = &MsgBascisX{}

type MsgBascisX struct {
	ProtocolName string `json:"protocol_name" yaml:"protocol_name"`
	// Inscription represents the inscription data of protocol operations on the chain.
	Inscription string `json:"inscription" yaml:"inscription"`
	BTCTxid     string `json:"btc_txid" yaml:"btc_txid"`
	BTCFee      string `json:"btc_fee" yaml:"btc_fee"`
	// Context represents the contextual data required for protocol operation execution.
	Context string `json:"inscription_context" yaml:"inscription_context"`
}

// NewMsgBascisX creates a new MsgBascisX instance
func NewMsgBascisX(protocolName string, inscription string, btcTxid string, btcFee string, ctx string) MsgBascisX {
	return MsgBascisX{
		ProtocolName: protocolName,
		Inscription:  inscription,
		BTCTxid:      btcTxid,
		BTCFee:       btcFee,
		Context:      ctx,
	}
}

// nolint
func (msg MsgBascisX) Route() string { return RouterKey }
func (msg MsgBascisX) Type() string  { return MsgBascisXType }
func (msg MsgBascisX) GetSigners() []sdk.AccAddress {
	return nil
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgBascisX) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgBascisX) ValidateBasic() error {
	if len(msg.ProtocolName) == 0 {
		return ErrValidateBasic("msg.ProtocolName is empty")
	}
	if len(msg.BTCTxid) == 0 {
		return ErrValidateBasic("msg.BTCTxid is empty")
	}
	return nil
}

// Decoder Try to decode as MsgInscription by json
func Decoder(_ codec.CdcAbstraction, txBytes []byte) (tx sdk.Tx, err error) {
	var zeroTx types.ZeroRequestTx

	if err = rlp.DecodeBytes(txBytes, &zeroTx); err == nil {
		var msgInscription MsgInscription
		var context InscriptionContext
		// TODO, It may not be common here，This may not be universal because the inscription_context of non-ordinals protocol will also be decoded by json.
		if err = json.Unmarshal([]byte(zeroTx.InscriptionContext), &context); err == nil {
			msgInscription.Inscription = zeroTx.Inscription
			// TODO 50000000 is tmp
			fee := authtypes.NewStdFee(50000000, nil)
			return authtypes.NewStdTx([]sdk.Msg{msgInscription}, fee, nil, ""), nil
		} else {
			msg := NewMsgBascisX(zeroTx.ProtocolName, zeroTx.Inscription, zeroTx.BTCTxid, zeroTx.BTCFee, zeroTx.InscriptionContext)
			// TODO 50000000 is tmp
			fee := authtypes.NewStdFee(50000000, nil)
			return authtypes.NewStdTx([]sdk.Msg{msg}, fee, nil, ""), nil
		}
	}

	return nil, ErrValidateInput(fmt.Sprintf("inscription msg deocer failed: %s", err))
}
