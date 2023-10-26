package keeper

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func (k Keeper) GetValue(ctx sdk.Context, key string) string {
	store := ctx.KVStore(k.storeKey)

	v := store.Get([]byte(key))
	return string(v)
}

func (k Keeper) SetValue(ctx sdk.Context, key string, value string) {
	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(key), []byte(value))
}
