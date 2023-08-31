package types

import (
	codectypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/codec/types"
	"github.com/brc20-collab/brczero/libs/ibc-go/modules/core/exported"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}
