package keeper

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
)

const (
	QueryKV = "kv"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryKV:
			return queryKV(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryKV(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	ret := k.GetValue(ctx, string(req.Data))

	return []byte(ret), nil
}
