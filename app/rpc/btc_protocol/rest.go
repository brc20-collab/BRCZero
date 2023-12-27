package btc_protocol

import (
	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
)

func RegisterBtcProtocolRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	registerBrc20QueryRoutes(cliCtx, r, ethApi)
	registerRuneQueryRoutes(cliCtx, r, ethApi)
}
