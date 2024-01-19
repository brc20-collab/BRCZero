package keeper

import (
	"fmt"
	"math/big"

	ethcm "github.com/ethereum/go-ethereum/common"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/x/brcx/types"
	"github.com/brc20-collab/brczero/x/common"
)

func src20QueryTick(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.Src20TokenInfoReq

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolSRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}

	input, err := types.GetSrc20TickInfoInput(params.Tick)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getSRC20TokenInfo failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getSRC20TokenInfo failed: %s", err))
	}

	tickInfo, err := types.UnpackGetSrc20TickInfoOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getTickInformation failed: %s", err))
	}

	response := tickInfo.ToResponse()
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}

func src20QueryBalance(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.Src20BalanceParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolSRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}

	input, err := types.GetSrc20BalanceInput(params.Address, params.Tick)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getSRC20Balance failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getSRC20Balance failed: %s", err))
	}

	balance, err := types.UnpackGetSrc20BalanceOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getSRC20Balance failed: %s", err))
	}

	response := balance.ToResponse()
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}

func src20QueryAllBalance(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.Src20BalanceParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolSRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}

	input, err := types.GetSrc20AllBalanceInput(params.Address)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getAllSRC20Balance failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getAllSRC20Balance failed: %s", err))
	}

	balances, err := types.UnpackGetSrc20AllBalanceOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getSRC20Balance failed: %s", err))
	}
	response := make([]types.QuerySrc20BalanceResponse, 0)
	for _, b := range balances {
		response = append(response, b.ToResponse())
	}

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}

func src20QueryAllTick(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.Src20AllTokenInfoReq
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	page, ok := new(big.Int).SetString(params.Page, 10)
	if !ok {
		return nil, types.ErrInternal("parse page of getSRC20AllTokenInfo failed")
	}

	pageSize, ok := new(big.Int).SetString(params.Limit, 10)
	if !ok {
		return nil, types.ErrInternal("parse pageSize of getSRC20AllTokenInfo failed")
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolSRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}

	input, err := types.GetSrc20AllTickInfoInput(page, pageSize)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getSRC20AllTokenInfo failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getSRC20AllTokenInfo failed: %s", err))
	}

	allTickInfo, err := types.UnpackGetSrc20AllTickInfoOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getSRC20AllTokenInfo failed: %s", err))
	}

	response := allTickInfo.ToResponse()
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}
