package vmbridge

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/module"
	"github.com/brc20-collab/brczero/x/vmbridge/keeper"
	"github.com/brc20-collab/brczero/x/wasm"
)

func RegisterServices(cfg module.Configurator, keeper keeper.Keeper) {
	RegisterMsgServer(cfg.MsgServer(), NewMsgServerImpl(keeper))
}

func GetWasmOpts(cdc *codec.ProtoCodec) wasm.Option {
	return wasm.WithMessageEncoders(RegisterSendToEvmEncoder(cdc))
}
