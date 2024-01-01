package btc_protocol

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/rest"
	srcxtypes "github.com/brc20-collab/brczero/x/brcx/types"
)

func registerSrc20QueryRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	r.HandleFunc("/src20/block/{btcBlockHash}/events", QuerySrc20TxsEventsByBtcHashHandlerFunc(cliCtx, ethApi)).Methods("Get")

}

func QuerySrc20TxsEventsByBtcHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		btcBlockHash := vars["btcBlockHash"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		blockLogs, err := ethApi.GetLogsByBtcHash(btcBlockHash)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		node, err := cliCtx.GetNode()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, SRC20)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := map[string][]srcxtypes.Src20EventResponse{}
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := srcxtypes.UnpackSrc20EventContext(l.Data)
				if err != nil {
					rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]srcxtypes.Src20EventResponse, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext.ToEventResponse())
			}
		}

		txEventsResp := make([]srcxtypes.QuerySrc20TxEventsResponse, 0)
		for txid, events := range resMap {
			txEventsResp = append(txEventsResp, srcxtypes.NewQuerySrc20TxEventsResponse(events, txid))
		}

		blockEventsResp := srcxtypes.NewQuerySrc20TxEventsByBlockHashResponse(txEventsResp)

		response := srcxtypes.NewOKApiResponse(blockEventsResp)
		resp, err := cliCtx.Codec.MarshalJSON(response)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponseBare(w, cliCtx, resp)
	}
}
