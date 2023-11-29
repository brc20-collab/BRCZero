package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/rest"
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

	//r.HandleFunc(
	//	"/brc20/holders",
	//	QueryAllTransferableBalanceByAddrHandlerFn(cliCtx),
	//).Methods("GET")

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
