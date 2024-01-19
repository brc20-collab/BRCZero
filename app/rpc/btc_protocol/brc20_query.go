package btc_protocol

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/rest"
	xtypes "github.com/brc20-collab/brczero/x/brcx/types"
)

func registerBrc20QueryRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	r.HandleFunc(
		"/brc20/block/{btcBlockHash}/events",
		QueryBrc20TxsEventsByBtcHashHandlerFunc(cliCtx, ethApi),
	).Methods("Get")

	r.HandleFunc(
		"/brc20/tx/{txid}/events",
		QueryBrc20TxsEventsByBtcTxidHandlerFunc(cliCtx, ethApi),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/tick/{tickName}",
		QueryBrc20TickByNameHandlerFn(cliCtx),
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
		QueryBrc20AllBalanceByAddrHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/tick/{tickName}/address/{address}/transferable",
		QueryBrc20TransferableBalanceByNameAndAddrHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/address/{address}/transferable",
		QueryBrc20AllTransferableBalanceByAddrHandlerFn(cliCtx),
	).Methods("GET")

	r.HandleFunc(
		"/brc20/holders",
		QueryBrc20TotalTickHoldersHandlerFn(cliCtx),
	).Methods("GET")
}

func QueryBrc20TxsEventsByBtcHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
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
		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, BRC20)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := map[string][]xtypes.Brc20EventResponse{}
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := xtypes.UnpackBrc20EventContext(l.Data)
				if err != nil {
					WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]xtypes.Brc20EventResponse, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext.ToEventResponse())
			}
		}

		txEventsResp := make([]xtypes.QueryBrc20TxEventsResponse, 0)
		for txid, events := range resMap {
			txEventsResp = append(txEventsResp, xtypes.NewQueryBrc20TxEventsResponse(events, txid))
		}

		resp := xtypes.NewQueryBrc20TxEventsByBlockHashResponse(txEventsResp)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QueryBrc20TxsEventsByBtcTxidHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		targetTxid := vars["txid"]

		node, err := cliCtx.GetNode()
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		btcBlockHash, err := node.BtcBlockHashByBtcTxid(targetTxid)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		blockLogs, err := ethApi.GetLogsByBtcHash(btcBlockHash)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, BRC20)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := map[string][]xtypes.Brc20EventResponse{}
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := xtypes.UnpackBrc20EventContext(l.Data)
				if err != nil {
					WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]xtypes.Brc20EventResponse, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext.ToEventResponse())
			}
		}

		resp := xtypes.NewQueryBrc20TxEventsResponse(resMap[targetTxid], targetTxid)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QueryBrc20TickByNameHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickName := mux.Vars(r)["tickName"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := xtypes.NewQueryTickParams(tickName)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.QueryTick), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var data xtypes.QueryBrc20TickInfoResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &data)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		resp := xtypes.NewOKApiResponse(data)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QueryAllTickHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//todo: pagination
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.QueryAllTick), nil)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var data xtypes.QueryBrc20AllTickInfoResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &data)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		resp := xtypes.NewOKApiResponse(data)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
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

		params := xtypes.NewQueryDataParams(addr, tickName)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.QueryBalance), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var data xtypes.QueryBrc20BalanceResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &data)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		resp := xtypes.NewOKApiResponse(data)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QueryBrc20AllBalanceByAddrHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addr := mux.Vars(r)["address"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := xtypes.NewQueryAllDataParams(addr)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.QueryAllBalance), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		var data xtypes.QueryBrc20AllBalanceResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &data)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		resp := xtypes.NewOKApiResponse(data)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QueryBrc20TransferableBalanceByNameAndAddrHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickName := mux.Vars(r)["tickName"]
		addr := mux.Vars(r)["address"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := xtypes.NewQueryDataParams(addr, tickName)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.QueryTransferableTick), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var data xtypes.QueryBrc20TransferableInscriptionResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &data)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		resp := xtypes.NewOKApiResponse(data)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QueryBrc20AllTransferableBalanceByAddrHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addr := mux.Vars(r)["address"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		params := xtypes.NewQueryAllDataParams(addr)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.QueryAllTransferableTick), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var data xtypes.QueryBrc20TransferableInscriptionResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &data)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp := xtypes.NewOKApiResponse(data)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QueryBrc20TotalTickHoldersHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.QueryTotalTickHolders), nil)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var data xtypes.QueryBrc20TotalTickHoldersResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &data)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp := xtypes.NewOKApiResponse(data)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}
