package keeper

import (
	"fmt"
	"math/big"

	ethcm "github.com/ethereum/go-ethereum/common"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/x/brcx/types"
	"github.com/brc20-collab/brczero/x/common"
)

// NewQuerier creates a new querier for slashing clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryProtocol:
			return queryProtocol(ctx, req, k)
		case types.QueryTick:
			return queryTick(ctx, req, k)
		case types.QueryAllTick:
			return queryAllTick(ctx, req, k)
		case types.QueryBalance:
			return queryBalance(ctx, req, k)
		case types.QueryAllBalance:
			return queryAllBalance(ctx, req, k)
		case types.QueryTotalTickHolders:
			return queryTotalTickHolders(ctx, req, k)
		case types.QueryTransferableTick:
			return queryTransferableTick(ctx, req, k)
		case types.QueryAllTransferableTick:
			return queryAllTransferableTick(ctx, req, k)
		case types.Src20QueryTick:
			return src20QueryTick(ctx, req, k)
		case types.Src20QueryBalance:
			return src20QueryBalance(ctx, req, k)
		case types.Src20QueryAllBalance:
			return src20QueryAllBalance(ctx, req, k)
		case types.Src20QueryAllTick:
			return src20QueryAllTick(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

func queryProtocol(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	name := string(req.Data)
	addresses := k.GetContractAddressByName(ctx, name)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, addresses)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryTick(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryTickParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}
	input, err := types.GetBrc20TickInfoInput(params.Name)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getTickInformation failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getTickInformation failed: %s", err))
	}

	tickInfo, err := types.UnpackGetBrc20TickInfoOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getTickInformation failed: %s", err))
	}

	response := types.NewQueryBrc20TickInfoResponse(tickInfo)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}

func queryAllTick(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}
	input, err := types.GetBrc20AllTickInfoInput()
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getAllTickInformation failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getAllTickInformation failed: %s", err))
	}

	tickInfos, err := types.UnpackGetBrc20AllTickInfoOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getAllTickInformation failed: %s", err))
	}

	tickInfosResp := make([]types.QueryBrc20TickInfoResponse, 0)
	for _, info := range tickInfos {
		tickInfosResp = append(tickInfosResp, info.ToResponse())
	}

	response := types.NewQueryBrc20AllTickInfoResponse(tickInfosResp)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}
	return res, nil
}

func queryBalance(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryDataParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}
	input, err := types.GetBrc20BalanceInput(params.Addr, params.Name)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getBalance failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getBalance failed: %s", err))
	}

	totalBalances, availableBalance, transferableBalance, err := types.UnpackGetBrc20BalanceOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getBalance failed: %s", err))
	}

	response := types.NewQueryBrc20BalanceResponse(params.Name, availableBalance.String(), transferableBalance.String(), totalBalances.String())
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}

func queryAllBalance(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryAllDataParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}

	input, err := types.GetBrc20AllBalanceInput(params.Addr)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getAllBalance failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getBalance failed: %s", err))
	}

	allBalance, err := types.UnpackGetBrc20AllBalanceOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getAllBalance failed: %s", err))
	}

	balanceResp := make([]types.QueryBrc20BalanceResponse, 0)
	for _, b := range allBalance {
		balanceResp = append(balanceResp, b.ToResponse())
	}

	response := types.NewQueryBrc20AllBalanceResponse(balanceResp)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}

func queryTotalTickHolders(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}

	input, err := types.GetBrc20TotalTickHoldersInput()
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getTotalTickHolders failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getTotalTickHolders failed: %s", err))
	}

	holders, err := types.UnpackGetBrc20TotalTickHoldersOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getTotalTickHolders failed: %s", err))
	}

	response := types.NewQueryBrc20TotalTickHoldersResponse(holders.String())
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}
	return res, nil
}

func queryTransferableTick(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryDataParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}
	input, err := types.GetBrc20TransferableTickInput(params.Addr, params.Name)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getUserTransferableTickInformation failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getUserTransferableTickInformation failed: %s", err))
	}

	tis, err := types.UnpackGetBrc20TransferableTickOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getUserTransferableTickInformation failed: %s", err))
	}

	response := types.NewQueryBrc20TransferableInscriptionResponse(tis)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}

func queryAllTransferableTick(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryAllDataParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}
	input, err := types.GetBrc20AllTransferableTickInput(params.Addr)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getUserAllTransferableTickInformation failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getUserAllTransferableTickInformation failed: %s", err))
	}

	tis, err := types.UnpackGetBrc20AllTransferableTickOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getUserAllTransferableTickInformation failed: %s", err))
	}

	response := types.NewQueryBrc20TransferableInscriptionResponse(tis)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}
