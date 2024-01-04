package btc_protocol

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/rest"
	basicxtypes "github.com/brc20-collab/brczero/x/brcx/types"
)

func registerRuneQueryRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	r.HandleFunc("/runealpha/block/{btcBlockHash}/events", QueryRuneAlphaTxsEventsByBtcHashHandlerFunc(cliCtx, ethApi)).Methods("Get")

}

func QueryRuneAlphaTxsEventsByBtcHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		btcBlockHash := vars["btcBlockHash"]

		blockLogs, err := ethApi.GetLogsByBtcHash(btcBlockHash)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		node, err := cliCtx.GetNode()
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, RUNEALPHA)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := make(map[string][]interface{})
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				if len(l.Topics) == 0 {
					continue
				}

				handler, ok := basicxtypes.TopicEventMap[l.Topics[0]]
				if !ok {
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := handler(l.Data)
				if err != nil {
					rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]interface{}, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext)
			}
		}

		txEventsResp := make([]interface{}, 0)
		for txid, events := range resMap {
			txEventsResp = append(txEventsResp, basicxtypes.NewQueryRuneAlphaTxEventsResponse(events, txid))
		}

		resp := basicxtypes.NewQueryRuneAlphaTxEventsByBlockHashResponse(txEventsResp)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}
