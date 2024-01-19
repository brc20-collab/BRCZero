package btc_protocol

import (
	"encoding/json"
	"fmt"
	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/rest"
	"github.com/brc20-collab/brczero/x/brcx/types/arc20"
	"github.com/gorilla/mux"
	"net/http"
)

func registerArc220QueryRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	r.HandleFunc("/proxy/blockchain.atomicals.get_block_ok_events", QueryArc20TxsEventsByBtcHashHandlerFunc(cliCtx, ethApi)).Methods("Get")

}

type ARC20EventRequestParams []uint64

func QueryArc20TxsEventsByBtcHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.FormValue("params")
		var heights ARC20EventRequestParams
		err := json.Unmarshal([]byte(params), &heights)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		if len(heights) != 1 {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("zero only support 1 height query events").Error())
			return
		}

		if heights[0]-cliCtx.StartBtcHeight < 1 {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("zero only support %d height events query", cliCtx.StartBtcHeight+1).Error())
			return
		}
		height := heights[0] - cliCtx.StartBtcHeight
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		blockLogs, _, err := ethApi.GetLogsOptimizeForEvent(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resEvent := make([]arc20.AtomicalEventParam, 0)
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}
				if len(l.Topics) == 0 {
					continue
				}
				eventContext, err := arc20.UnpackArc20EventContext(l.Data, l.Topics[0])
				if err != nil {
					if err == arc20.ErrNotExpectEvent {
						continue
					}
					rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}
				resEvent = append(resEvent, eventContext)
			}
		}

		response := arc20.ARC20Response{Success: true, Response: arc20.AtomicalResponse{Result: resEvent}}
		resp, err := cliCtx.Codec.MarshalJSON(response)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponseBare(w, cliCtx, resp)
	}
}
