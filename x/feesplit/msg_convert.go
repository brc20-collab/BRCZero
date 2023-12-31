package feesplit

import (
	"encoding/json"
	"errors"
	"github.com/brc20-collab/brczero/libs/system"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/baseapp"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/x/common"
	"github.com/brc20-collab/brczero/x/feesplit/types"
)

var (
	ErrCheckSignerFail = errors.New("check signer fail")
)

func init() {
	RegisterConvert()
}

func RegisterConvert() {
	baseapp.RegisterCmHandle(system.Chain+"/MsgRegisterFeeSplit", baseapp.NewCMHandle(ConvertRegisterFeeSplitMsg, 0))
	baseapp.RegisterCmHandle(system.Chain+"/MsgUpdateFeeSplit", baseapp.NewCMHandle(ConvertUpdateFeeSplitMsg, 0))
	baseapp.RegisterCmHandle(system.Chain+"/MsgCancelFeeSplit", baseapp.NewCMHandle(ConvertCancelFeeSplitMsg, 0))
}

func ConvertRegisterFeeSplitMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgRegisterFeeSplit{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}

func ConvertUpdateFeeSplitMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgUpdateFeeSplit{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}

func ConvertCancelFeeSplitMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgCancelFeeSplit{}
	err := json.Unmarshal(data, &newMsg)
	if err != nil {
		return nil, err
	}
	err = newMsg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	if ok := common.CheckSignerAddress(signers, newMsg.GetSigners()); !ok {
		return nil, ErrCheckSignerFail
	}
	return newMsg, nil
}
