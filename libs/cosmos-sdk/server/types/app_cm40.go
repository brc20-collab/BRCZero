package app

import (
	cliContext "github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	gogogrpc "github.com/gogo/protobuf/grpc"
)

type ApplicationAdapter interface {
	RegisterGRPCServer(gogogrpc.Server)
	RegisterTxService(clientCtx cliContext.CLIContext)
}
