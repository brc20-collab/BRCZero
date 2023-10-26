package rest

import (
	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	//r.HandleFunc("/hello/say", SayHandlerFn(cliCtx)).Methods("POST")
}
