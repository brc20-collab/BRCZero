package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
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

		params := types.NewQueryBalanceParams(addr, tickName)
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

		params := types.NewQueryAllBalanceParams(addr)
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

	}
}

func QueryAllTransferableBalanceByAddrHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		rest.PostProcessResponseBare(w, cliCtx, searchResult.Txs[0].Logs[0].Events)
	}
}
