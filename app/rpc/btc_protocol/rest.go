package btc_protocol

import (
	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
)

const (
	BRC20     = "brc-20"
	SRC20     = "src-20"
	RUNEALPHA = "runealpha"
)

func RegisterBtcProtocolRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	registerBrc20QueryRoutes(cliCtx, r, ethApi)
	registerRuneQueryRoutes(cliCtx, r, ethApi)
	registerSrc20QueryRoutes(cliCtx, r, ethApi)
}
