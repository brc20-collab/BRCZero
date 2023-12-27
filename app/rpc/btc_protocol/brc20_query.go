package btc_protocol

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/rest"
	brcxtypes "github.com/brc20-collab/brczero/x/brcx/types"
)

func registerBrc20QueryRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	r.HandleFunc("/brc20/block/{btcBlockHash}/events", QueryBrc20TxsEventsByBtcHashHandlerFunc(cliCtx, ethApi)).Methods("Get")
	r.HandleFunc("/brc20/block/{txid}/events", QueryBrc20TxsEventsByBtcTxidHandlerFunc(cliCtx, ethApi)).Methods("GET")
}

func QueryBrc20TxsEventsByBtcHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
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
		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, BRC20)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := map[string][]brcxtypes.Brc20EventResponse{}
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := brcxtypes.UnpackBrc20EventContext(l.Data)
				if err != nil {
					rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]brcxtypes.Brc20EventResponse, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext.ToEventResponse())
			}
		}

		txEventsResp := make([]brcxtypes.QueryBrc20TxEventsResponse, 0)
		for txid, events := range resMap {
			txEventsResp = append(txEventsResp, brcxtypes.NewQueryBrc20TxEventsResponse(events, txid))
		}

		blockEventsResp := brcxtypes.NewQueryBrc20TxEventsByBlockHashResponse(txEventsResp)

		response := brcxtypes.NewOKApiResponse(blockEventsResp)
		resp, err := cliCtx.Codec.MarshalJSON(response)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponseBare(w, cliCtx, resp)
	}
}

func QueryBrc20TxsEventsByBtcTxidHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		targetTxid := vars["txid"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		node, err := cliCtx.GetNode()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		btcBlockHash, err := node.BtcBlockHashByBtcTxid(targetTxid)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		blockLogs, err := ethApi.GetLogsByBtcHash(btcBlockHash)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, BRC20)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := map[string][]brcxtypes.Brc20EventResponse{}
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := brcxtypes.UnpackBrc20EventContext(l.Data)
				if err != nil {
					rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]brcxtypes.Brc20EventResponse, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext.ToEventResponse())
			}
		}

		txEventsResp := brcxtypes.NewQueryBrc20TxEventsResponse(resMap[targetTxid], targetTxid)

		response := brcxtypes.NewOKApiResponse(txEventsResp)
		resp, err := cliCtx.Codec.MarshalJSON(response)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponseBare(w, cliCtx, resp)
	}
}
