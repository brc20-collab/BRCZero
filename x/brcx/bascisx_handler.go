package brcx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/x/brcx/types"
)

func handleBasicXInscription(ctx sdk.Context, msg MsgBasicProtocolOp, k Keeper, info *ResultInfo) (*sdk.Result, error) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeBasicX,
			sdk.NewAttribute(AttributeProtocolName, msg.ProtocolName),
		),
	)
	switch msg.ProtocolName {
	case ManageContractProtocolName:
		result, err := handleBasicXManageContract(ctx, msg, k, info)
		if err != nil {
			return result, err
		}
		return result, nil
	default:
		return handleBasicXEntryPoint(ctx, msg, msg.ProtocolName, k, info)
	}
}

func handleBasicXManageContract(ctx sdk.Context, msg MsgBasicProtocolOp, k Keeper, info *ResultInfo) (*sdk.Result, error) {
	var context InscriptionContext
	if err := json.Unmarshal([]byte(msg.Context), &context); err != nil {
		return nil, ErrValidateInput(fmt.Sprintf("InscriptionContext json unmarshall is err: %s ", err))
	}
	from, err := ConvertBTCAddress(context.Sender)
	if err != nil {
		return nil, ErrValidateInput(fmt.Sprintf("InscriptionContext.Sender %s is not address: %s ", context.Sender, err))
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

func handleBasicXEntryPoint(ctx sdk.Context, msg MsgBasicProtocolOp, protocol string, k Keeper, info *ResultInfo) (*sdk.Result, error) {
	from := common.BytesToAddress(k.GetBRCXAddress().Bytes())
	info.EvmCaller = from.String()
	to, err := k.GetContractAddrByProtocol(ctx, protocol)
	if err != nil {
		return nil, ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}

	info.EvmTo = to.String()
	input, err := types.GetBasicXEntryPointInput(msg.Context, msg.Inscription)
	if err != nil {
		return nil, ErrPackInput(fmt.Sprintf("pack basicX entry point input failed: %s", err))
	}

	info.CallData = hex.EncodeToString(input)
	executionResult, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, info)
	if err != nil {
		return nil, ErrCallMethod(fmt.Sprintf("call basicX entryPoint failed: %s", err))
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeEntryPoint, sdk.NewAttribute(AttributeEvmOutput, hex.EncodeToString(contractResult.Ret))))

	return executionResult.Result, nil
}
