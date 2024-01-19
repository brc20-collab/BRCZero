package brcx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
	"github.com/brc20-collab/brczero/x/brcx/types"
)

// NewHandler creates a sdk.Handler for all the slashing type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx.SetEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgInscription:
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					EventTypeBasicX,
					sdk.NewAttribute(AttributeBTCTXID, msg.InscriptionContext.Txid),
				),
			)
			info := types.ResultInfo{BTCTxid: msg.InscriptionContext.Txid}
			result, err := handleInscription(ctx, msg, k, &info)
			// json.Marshal can not be error. even if error it hash a few influence with execute of transaction.
			buff, _ := json.Marshal(info)
			if err != nil {
				return &sdk.Result{Events: ctx.EventManager().Events(), Info: buff}, err
			}
			result.Events = append(result.Events, ctx.EventManager().Events()...)
			result.Info = buff
			return result, err
		case types.MsgBasicProtocolOp:
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					EventTypeBasicX,
					sdk.NewAttribute(AttributeBTCTXID, msg.BTCTxid),
				),
			)
			info := types.ResultInfo{BTCTxid: msg.BTCTxid}
			result, err := handleBasicXInscription(ctx, msg, k, &info)
			// json.Marshal can not be error. even if error it hash a few influence with execute of transaction.
			buff, _ := json.Marshal(info)
			if err != nil {
				return &sdk.Result{Events: ctx.EventManager().Events(), Info: buff}, err
			}
			result.Events = append(result.Events, ctx.EventManager().Events()...)
			result.Info = buff
			return result, err
		default:
			return &sdk.Result{Events: ctx.EventManager().Events()}, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleInscription(ctx sdk.Context, msg MsgInscription, k Keeper, info *ResultInfo) (*sdk.Result, error) {
	inscription := make(map[string]interface{})
	err := json.Unmarshal([]byte(msg.Inscription), &inscription)
	if err != nil {
		return &sdk.Result{}, ErrValidateInput("msg Inscription json marshal failed")
	}
	p, ok := inscription["p"]
	if !ok {
		return &sdk.Result{}, ErrValidateInput("can not analyze protocol")
	}
	protocol, ok := p.(string)
	if !ok {
		return &sdk.Result{}, ErrValidateInput("the type of protocol must be string")
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeBasicX,
			sdk.NewAttribute(AttributeProtocolName, protocol),
		),
	)
	switch protocol {
	case ManageContractProtocolName:
		result, err := handleManageContract(ctx, msg, k, info)
		if err != nil {
			return result, err
		}
		return result, nil
	default:
		return handleEntryPoint(ctx, msg, protocol, k, info)
	}
}

func handleManageContract(ctx sdk.Context, msg MsgInscription, k Keeper, info *ResultInfo) (*sdk.Result, error) {
	if msg.InscriptionContext.IsTransfer {
		return nil, ErrValidateInput("manageContract can't deal inscription of transfer")
	}
	from, err := ConvertBTCAddress(msg.InscriptionContext.Sender)
	if err != nil {
		return nil, ErrValidateInput(fmt.Sprintf("InscriptionContext.Sender %s is not address: %s ", msg.InscriptionContext.Sender, err))
	}
	info.EvmCaller = from.String()

	var manageContract ManageContract
	if err := json.Unmarshal([]byte(msg.Inscription), &manageContract); err != nil {
		return nil, ErrValidateInput(fmt.Sprintf("Inscription json unmarshal failed: %s ", err))
	}

	if err := manageContract.ValidateBasic(); err != nil {
		return nil, err
	}
	calldata, err := manageContract.GetCallData()
	if err != nil {
		return nil, ErrValidateInput(fmt.Sprintf("Inscription data is not hex: %s ", err))
	}
	info.CallData = manageContract.CallData
	manageContractEvent := sdk.NewEvent(EventTypeManageContract, sdk.NewAttribute(AttributeManageContractOperation, manageContract.Operation))
	var result sdk.Result
	switch manageContract.Operation {
	case ManageCreateContract:
		executeResult, contractResult, err := k.CallEvm(ctx, common.BytesToAddress(from[:]), nil, common.Big0, calldata, info)
		if err != nil {
			return nil, ErrExecute(fmt.Sprintf("create contract failed: %s", err))
		}
		result = *executeResult.Result
		k.InsertContractAddressWithName(ctx, manageContract.Name, contractResult.ContractAddress.Bytes())
		manageContractEvent = manageContractEvent.AppendAttributes(
			sdk.NewAttribute(AttributeManageContractAddress, contractResult.ContractAddress.Hex()),
			sdk.NewAttribute(AttributeEvmOutput, hex.EncodeToString(contractResult.Ret)))
	case ManageCallContract:
		to := common.HexToAddress(manageContract.Contract)
		info.EvmTo = to.String()
		executeResult, contractResult, err := k.CallEvm(ctx, common.BytesToAddress(from[:]), &to, common.Big0, calldata, info)
		if err != nil {
			return nil, fmt.Errorf("create contract failed: %v", err)
		}
		manageContractEvent = manageContractEvent.AppendAttributes(
			sdk.NewAttribute(AttributeEvmOutput, hex.EncodeToString(contractResult.Ret)),
		)
		result = *executeResult.Result
	default:
		return nil, ErrUnknownOperationOfManageContract(manageContract.Operation)
	}

	ctx.EventManager().EmitEvent(manageContractEvent)
	return &result, nil
}

func handleEntryPoint(ctx sdk.Context, msg MsgInscription, protocol string, k Keeper, info *ResultInfo) (*sdk.Result, error) {
	from := common.BytesToAddress(k.GetBRCXAddress().Bytes())
	info.EvmCaller = from.String()
	to, err := k.GetContractAddrByProtocol(ctx, protocol)
	if err != nil {
		return nil, ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}
	info.EvmTo = to.String()
	input, err := types.GetBrc20EntryPointInput(msg.InscriptionContext, msg.Inscription)
	if err != nil {
		return nil, ErrPackInput(fmt.Sprintf("pack entry point input failed: %s", err))
	}

	info.CallData = hex.EncodeToString(input)
	executionResult, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, info)
	if err != nil {
		return nil, ErrCallMethod(fmt.Sprintf("call entryPoint failed: %s", err))
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeEntryPoint, sdk.NewAttribute(AttributeEvmOutput, hex.EncodeToString(contractResult.Ret))))

	return executionResult.Result, nil
}
