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
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
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
	input, err := types.GetTickInfoInput(params.Name)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getTickInformation failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getTickInformation failed: %s", err))
	}

	tickInfo, err := types.UnpackGetTickInfoOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getTickInformation failed: %s", err))
	}

	response := types.NewQueryTickInfoResponse(tickInfo)
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
	input, err := types.GetAllTickInfoInput()
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getAllTickInformation failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getAllTickInformation failed: %s", err))
	}

	tickInfos, err := types.UnpackGetAllTickInfoOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getAllTickInformation failed: %s", err))
	}

	tickInfosResp := make([]types.QueryTickInfoResponse, 0)
	for _, info := range tickInfos {
		tickInfosResp = append(tickInfosResp, info.ToResponse())
	}

	response := types.NewQueryAllTickInfoResponse(tickInfosResp)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}
	return res, nil
}

func queryBalance(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryBalanceParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}
	input, err := types.GetBalanceInput(params.Addr, params.Name)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getBalance failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getBalance failed: %s", err))
	}

	totalBalances, availableBalance, transferableBalance, err := types.UnpackGetBalanceOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getBalance failed: %s", err))
	}

	response := types.NewQueryBalanceResponse(params.Name, availableBalance.String(), transferableBalance.String(), totalBalances.String())
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}

	return res, nil
}

func queryAllBalance(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryAllBalanceParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, common.ErrUnMarshalJSONFailed(err.Error())
	}

	from := ethcm.BytesToAddress(k.GetBRCXAddress().Bytes())
	to, err := k.GetContractAddrByProtocol(ctx, types.ProtocolBRC20)
	if err != nil {
		return nil, types.ErrGetContractAddress(fmt.Sprintf("get contract address by protocol failed: %s", err))
	}

	input, err := types.GetAllBalanceInput(params.Addr)
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getAllBalance failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getBalance failed: %s", err))
	}

	allBalance, err := types.UnpackGetAllBalanceOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getAllBalance failed: %s", err))
	}

	balanceResp := make([]types.QueryBalanceResponse, 0)
	for _, b := range allBalance {
		balanceResp = append(balanceResp, b.ToResponse())
	}

	response := types.NewQueryAllBalanceResponse(balanceResp)
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

	input, err := types.GetTotalTickHoldersInput()
	if err != nil {
		return nil, types.ErrPackInput(fmt.Sprintf("pack input of getTotalTickHolders failed: %s", err))
	}

	_, contractResult, err := k.CallEvm(ctx, from, &to, big.NewInt(0), input, &types.ResultInfo{})
	if err != nil {
		return nil, types.ErrCallMethod(fmt.Sprintf("call getTotalTickHolders failed: %s", err))
	}

	holders, err := types.UnpackGetTotalTickHoldersOutput(contractResult.Ret)
	if err != nil {
		return nil, types.ErrUnpackOutput(fmt.Sprintf("unpack output of getTotalTickHolders failed: %s", err))
	}

	response := types.NewQueryTotalTickHoldersResponse(holders.String())
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, response)
	if err != nil {
		return nil, common.ErrMarshalJSONFailed(err.Error())
	}
	return res, nil
}
