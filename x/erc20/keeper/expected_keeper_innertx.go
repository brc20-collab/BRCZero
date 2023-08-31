package keeper

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"
)

type EvmKeeper interface {
	GetChainConfig(ctx sdk.Context) (evmtypes.ChainConfig, bool)
	GenerateCSDBParams() evmtypes.CommitStateDBParams
	GetParams(ctx sdk.Context) evmtypes.Params
	AddInnerTx(...interface{})
	AddContract(...interface{})
}
