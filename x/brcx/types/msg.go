package types

import (
	"encoding/hex"
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

// Decoder Try to decode as MsgInscription by json
func Decoder(_ codec.CdcAbstraction, txBytes []byte) (tx sdk.Tx, err error) {
	var brczeroTx types.BRCZeroRequestTx

	fmt.Printf("txBytes:%s\n", hex.EncodeToString(txBytes))
	if err = rlp.DecodeBytes(txBytes, &brczeroTx); err == nil {
		var msgInscription MsgInscription
		if err = json.Unmarshal([]byte(brczeroTx.HexRlpEncodeTx), &msgInscription); err == nil {
			// TODO 1000 is tmp
			fee := authtypes.NewStdFee(50000000, nil)
			fmt.Printf("========fsc-test: msgInscription:%v\n", msgInscription)
			return authtypes.NewStdTx([]sdk.Msg{msgInscription}, fee, nil, ""), nil
		} else {
			fmt.Println("xxxxxxxxx-json-unmarshal err.")
		}
	}
	fmt.Println("xxxxxxxxx-decode tx err.")

	return nil, ErrValidateInput(fmt.Sprintf("brcx msg deocer failed: %s", err))
}
