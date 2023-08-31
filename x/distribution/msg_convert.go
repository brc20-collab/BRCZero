package distribution

import (
	"encoding/json"
	"errors"
	"github.com/brc20-collab/brczero/libs/system"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/baseapp"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/x/common"
	"github.com/brc20-collab/brczero/x/distribution/types"
)

var (
	ErrCheckSignerFail = errors.New("check signer fail")
)

func init() {
	RegisterConvert()
}

func RegisterConvert() {
	baseapp.RegisterCmHandle(system.Chain+"/distribution/MsgWithdrawDelegatorAllRewards", baseapp.NewCMHandle(ConvertWithdrawDelegatorAllRewardsMsg, 0))
}

func ConvertWithdrawDelegatorAllRewardsMsg(data []byte, signers []sdk.AccAddress) (sdk.Msg, error) {
	newMsg := types.MsgWithdrawDelegatorAllRewards{}
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
