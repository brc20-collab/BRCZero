package types

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	"github.com/brc20-collab/brczero/libs/system"
)

const (
	ManageTreasuresProposalName = system.Chain + "/mint/ManageTreasuresProposal"
	ExtraProposalName           = system.Chain + "/mint/ExtraProposal"
)

// ModuleCdc is a generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	codec.RegisterCrypto(ModuleCdc)
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(ManageTreasuresProposal{}, ManageTreasuresProposalName, nil)
	cdc.RegisterConcrete(ExtraProposal{}, ExtraProposalName, nil)
}
