package types

import "github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"

var ModuleCdc *codec.Codec

func init(){
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
