package types

import (
	"github.com/tendermint/go-amino"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgKV{}, "hello/MsgKV", nil)

	cdc.RegisterConcreteUnmarshaller("hello/MsgKV", func(c *amino.Codec, bytes []byte) (interface{}, int, error) {
		var msg MsgKV
		err := msg.UnmarshalFromAmino(c, bytes)
		if err != nil {
			return nil, 0, err
		}
		return msg, len(bytes), nil
	})
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
