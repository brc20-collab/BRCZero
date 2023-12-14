package rest

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	sdktypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/rest"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/client/utils"
	"github.com/brc20-collab/brczero/x/brcx/types"
	"github.com/brc20-collab/brczero/x/common"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/brc20/tick/{tickName}",
		QueryTickByNameHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/tick",
		QueryAllTickHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/tick/{tickName}/address/{address}/balance",
		QueryBalanceByNameAndAddrHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/address/{address}/balance",
		QueryAllBalanceByAddrHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/tick/{tickName}/address/{address}/transferable",
		QueryTransferableBalanceByNameAndAddrHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/address/{address}/transferable",
		QueryAllTransferableBalanceByAddrHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/holders",
		QueryTotalTickHoldersHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/tx/{txid}",
		QueryTxsByBtcTxIDRequestHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/tx/{txid}/events",
		QueryTxsEventsByBtcTxIDRequestHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/block/{btcBlockHash}/events",
		QueryTxsEventsByBtcHashRequestHandlerFn(cliCtx),
	).Methods("GET")

}

func QueryTickByNameHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickName := mux.Vars(r)["tickName"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryTickParams(tickName)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTick), bz)
		if err != nil {
			sdkErr := common.ParseSDKError(err.Error())
			common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func QueryAllTickHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//todo: pagination
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllTick), nil)
		if err != nil {
			sdkErr := common.ParseSDKError(err.Error())
			common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func QueryBalanceByNameAndAddrHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickName := mux.Vars(r)["tickName"]
		addr := mux.Vars(r)["address"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryDataParams(addr, tickName)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryBalance), bz)
		if err != nil {
			sdkErr := common.ParseSDKError(err.Error())
			common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func QueryAllBalanceByAddrHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addr := mux.Vars(r)["address"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryAllDataParams(addr)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllBalance), bz)
		if err != nil {
			sdkErr := common.ParseSDKError(err.Error())
			common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func QueryTransferableBalanceByNameAndAddrHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickName := mux.Vars(r)["tickName"]
		addr := mux.Vars(r)["address"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryDataParams(addr, tickName)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTransferableTick), bz)
		if err != nil {
			sdkErr := common.ParseSDKError(err.Error())
			common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func QueryAllTransferableBalanceByAddrHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addr := mux.Vars(r)["address"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := types.NewQueryAllDataParams(addr)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			common.HandleErrorMsg(w, cliCtx, common.CodeMarshalJSONFailed, err.Error())
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAllTransferableTick), bz)
		if err != nil {
			sdkErr := common.ParseSDKError(err.Error())
			common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func QueryTotalTickHoldersHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalTickHolders), nil)
		if err != nil {
			sdkErr := common.ParseSDKError(err.Error())
			common.HandleErrorMsg(w, cliCtx, sdkErr.Code, sdkErr.Message)
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func QueryTxsByBtcTxIDRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		btcTxID := vars["txid"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		events := make([]string, 0)
		tag := fmt.Sprintf("brcx.btc_txid='%s'", btcTxID)
		events = append(events, tag)

		searchResult, err := utils.QueryTxsByEvents(cliCtx, events, 1, 30)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponseBare(w, cliCtx, searchResult.Txs[0])
	}
}

func QueryTxsEventsByBtcTxIDRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		btcTxID := vars["txid"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		events := make([]string, 0)
		tag := fmt.Sprintf("brcx.btc_txid='%s'", btcTxID)
		events = append(events, tag)

		searchResult, err := utils.QueryTxsByEvents(cliCtx, events, 1, 30)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		eventResp, err := extractEventResponseFromTxResponse(searchResult.Txs)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		response := types.NewQueryTxEventsResponse(eventResp, btcTxID)
		rest.PostProcessResponseBare(w, cliCtx, response)
	}
}

func QueryTxsEventsByBtcHashRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		btcBlockHash := vars["btcBlockHash"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		events := make([]string, 0)
		tag := fmt.Sprintf("brcx.btc_block_hash='%s'", btcBlockHash)
		events = append(events, tag)

		searchResult, err := utils.QueryTxsByEvents(cliCtx, events, 1, 30)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		sortedTxRespMap, err := sortSearchedTxByBtcTxId(searchResult.Txs)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		responses := make([]types.QueryTxEventsResponse, 0)
		for txid, txs := range sortedTxRespMap {
			eventResp, err := extractEventResponseFromTxResponse(txs)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			response := types.NewQueryTxEventsResponse(eventResp, txid)
			responses = append(responses, response)
		}

		blockEvents := types.NewQueryTxEventsByBlockHashResponse(responses)

		rest.PostProcessResponseBare(w, cliCtx, blockEvents)
	}
}

func extractEventResponseFromTxResponse(txs []sdktypes.TxResponse) ([]types.EventResponse, error) {
	eventsJsonBytes := make([][]byte, 0)
	for _, tx := range txs {
		// A tx must contain only one msg and one log
		if len(tx.Logs) != 1 {
			return nil, errors.New(fmt.Sprintf("len(txResponse.Logs) is %d, which is expected to be 1", len(tx.Logs)))
		}

		for _, e := range tx.Logs[0].Events {
			if e.Type == "call_evm" {
				for _, a := range e.Attributes {
					if a.Key == "result" {
						//eventsJsonStrs = append(eventsJsonStrs, a.Value)
						eventsJsonStr := a.Value
						logs := gjson.Get(eventsJsonStr, "logs").Array()
						for _, l := range logs {
							event := l.Get("data").Str
							event = strings.TrimPrefix(event, "0x")

							eventBytes, err := hex.DecodeString(event)
							if err != nil {
								return nil, err
							}

							eventsJsonBytes = append(eventsJsonBytes, eventBytes)
						}
					}
				}
			}
		}
	}

	var resps []types.EventResponse
	for _, eb := range eventsJsonBytes {
		eventContext, err := types.UnpackEventContext(eb)
		if err != nil {
			return nil, err
		}

		resps = append(resps, eventContext.ToEventResponse())
	}

	return resps, nil
}

func sortSearchedTxByBtcTxId(txs []sdktypes.TxResponse) (map[string][]sdktypes.TxResponse, error) {
	sortedMap := make(map[string][]sdktypes.TxResponse)

	for _, tx := range txs {
		// A tx must contain only one msg and one log
		if len(tx.Logs) != 1 {
			return nil, errors.New(fmt.Sprintf("len(txResponse.Logs) is %d, which is expected to be 1", len(tx.Logs)))
		}

		for _, e := range tx.Logs[0].Events {
			if e.Type == types.ModuleName {
				for _, a := range e.Attributes {
					if a.Key == types.AttributeBTCTXID {
						txid := a.Value
						if _, ok := sortedMap[txid]; !ok {
							sortedMap[txid] = make([]sdktypes.TxResponse, 0)
						}

						sortedMap[txid] = append(sortedMap[txid], tx)
					}
				}
			}
		}

	}

	return sortedMap, nil
}
