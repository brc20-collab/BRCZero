package btc_protocol

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/x/brcx/types"
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
	registerArc220QueryRoutes(cliCtx, r, ethApi)
}

func PostProcessBasicXApiResponse(w http.ResponseWriter, cliCtx context.CLIContext, body interface{}) {
	response := types.NewOKApiResponse(body)
	resp, err := json.Marshal(response)
	if err != nil {
		WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}

func WriteBasicXApiErrorResponse(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// MustMarshalJson
	bytes, err := json.Marshal(types.NewApiError(1, msg))
	if err != nil {
		panic(err)
	}
	_, _ = w.Write(bytes)
}
