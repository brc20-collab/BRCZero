package keeper

import (
	"github.com/brc20-collab/brczero/x/vmbridge/types"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	"github.com/brc20-collab/brczero/libs/tendermint/libs/log"
)

type Keeper struct {
	cdc *codec.CodecProxy

	logger log.Logger

	evmKeeper     EVMKeeper
	wasmKeeper    WASMKeeper
	accountKeeper AccountKeeper
	bankKeeper    BankKeeper
}

func NewKeeper(cdc *codec.CodecProxy, logger log.Logger, evmKeeper EVMKeeper, wasmKeeper WASMKeeper, accountKeeper AccountKeeper, bk BankKeeper) *Keeper {
	logger = logger.With("module", types.ModuleName)
	return &Keeper{cdc: cdc, logger: logger, evmKeeper: evmKeeper, wasmKeeper: wasmKeeper, accountKeeper: accountKeeper, bankKeeper: bk}
}

func (k Keeper) Logger() log.Logger {
	return k.logger
}

func (k Keeper) getAminoCodec() *codec.Codec {
	return k.cdc.GetCdc()
}

func (k Keeper) GetProtoCodec() *codec.ProtoCodec {
	return k.cdc.GetProtocMarshal()
}
