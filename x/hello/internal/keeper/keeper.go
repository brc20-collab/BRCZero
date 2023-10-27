package keeper

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(
	cdc *codec.Codec, storeKey sdk.StoreKey,
) Keeper {

	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

func (k Keeper) GetValue(ctx sdk.Context, key []byte) []byte {
	store := ctx.KVStore(k.storeKey)

	v := store.Get(key)
	return v
}

func (k Keeper) SetValue(ctx sdk.Context, key []byte, value []byte) {
	store := ctx.KVStore(k.storeKey)

	store.Set(key, value)
}
