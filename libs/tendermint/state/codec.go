package state

import (
	amino "github.com/tendermint/go-amino"

	cryptoamino "github.com/brc20-collab/brczero/libs/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()
var ModuleCodec *amino.Codec

func init() {
	cryptoamino.RegisterAmino(cdc)
	ModuleCodec = cdc
}
