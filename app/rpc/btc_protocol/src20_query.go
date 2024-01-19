package btc_protocol

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	xtypes "github.com/brc20-collab/brczero/x/brcx/types"
)

func registerSrc20QueryRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	r.HandleFunc("/src20/events/block_hash/{btcBlockHash}", QuerySrc20TxsEventsByBtcHashHandlerFunc(cliCtx, ethApi)).Methods("POST")
	r.HandleFunc("/src20/events/tx/{txHash}", QuerySrc20TxEventsByTxHashHandlerFunc(cliCtx, ethApi)).Methods("POST")
	r.HandleFunc("/src20/events/block_index/{btcBlockIndex}", QuerySrc20TxsEventsByBtcIndexHandlerFunc(cliCtx, ethApi)).Methods("POST")

	r.HandleFunc("/src20/tick/info", QuerySrc20TokenInfoHandlerFunc(cliCtx)).Methods("POST")
	r.HandleFunc("/src20/height/latest", QuerySrc20HeightLastestHandlerFunc(cliCtx)).Methods("POST")
	r.HandleFunc("/src20/balance/tick/address/{address}", QuerySrc20BalanceHandlerFunc(cliCtx)).Methods("POST")
	r.HandleFunc("/src20/balance/address/{address}", QuerySrc20AllBalanceHandlerFunc(cliCtx)).Methods("POST")
	r.HandleFunc("/src20/tick/all", QuerySrc20AllTokenInfoHandlerFunc(cliCtx)).Methods("POST")
}

func QuerySrc20TxsEventsByBtcHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
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
		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, SRC20)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := map[string][]xtypes.Src20EventResponse{}
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := xtypes.UnpackSrc20EventContext(l.Data)
				if err != nil {
					WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]xtypes.Src20EventResponse, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext.ToEventResponse())
			}
		}

		txEventsResp := make([]xtypes.QuerySrc20TxEventsResponse, 0)
		for txid, events := range resMap {
			txEventsResp = append(txEventsResp, xtypes.NewQuerySrc20TxEventsResponse(events, txid))
		}

		resp := xtypes.NewQuerySrc20TxEventsByBlockHashResponse(txEventsResp)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QuerySrc20TxsEventsByBtcIndexHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		btcBlockIndexStr := vars["btcBlockIndex"]
		btcBlockIndex, err := strconv.ParseInt(btcBlockIndexStr, 10, 64)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		node, err := cliCtx.GetNode()
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		btcBlockHash, err := node.BtcBlockHashByBtcHeight(btcBlockIndex)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, SRC20)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		blockLogs, err := ethApi.GetLogsByBtcHash(btcBlockHash)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := map[string][]xtypes.Src20EventResponse{}
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := xtypes.UnpackSrc20EventContext(l.Data)
				if err != nil {
					WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]xtypes.Src20EventResponse, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext.ToEventResponse())
			}
		}

		txEventsResp := make([]xtypes.QuerySrc20TxEventsResponse, 0)
		for txid, events := range resMap {
			txEventsResp = append(txEventsResp, xtypes.NewQuerySrc20TxEventsResponse(events, txid))
		}

		resp := xtypes.NewQuerySrc20TxEventsByBlockHashResponse(txEventsResp)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QuerySrc20TxEventsByTxHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		targetTxid := vars["txHash"]

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

		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, SRC20)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := map[string][]xtypes.Src20EventResponse{}
		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				eventContext, err := xtypes.UnpackSrc20EventContext(l.Data)
				if err != nil {
					WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]xtypes.Src20EventResponse, 0, 1)
				}
				resMap[txid] = append(resMap[txid], eventContext.ToEventResponse())
			}
		}

		resp := xtypes.NewQuerySrc20TxEventsResponse(resMap[targetTxid], targetTxid)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QuerySrc20TokenInfoHandlerFunc(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req xtypes.Src20TokenInfoReq

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = cliCtx.Codec.UnmarshalJSON(body, &req)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		bz, err := cliCtx.Codec.MarshalJSON(req)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.Src20QueryTick), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var resp xtypes.QuerySrc20TickInfoResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &resp)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessBasicXApiResponse(w, cliCtx, resp)

	}
}

func QuerySrc20BalanceHandlerFunc(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		var req xtypes.Src20BalanceReq

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = cliCtx.Codec.UnmarshalJSON(body, &req)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		params := xtypes.NewSrc20BalanceParams(req.Tick, address)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.Src20QueryBalance), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var resp xtypes.QuerySrc20BalanceResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &resp)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessBasicXApiResponse(w, cliCtx, resp)

	}
}

func QuerySrc20AllBalanceHandlerFunc(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		params := xtypes.NewSrc20BalanceParams("", address)

		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.Src20QueryAllBalance), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var resp []xtypes.QuerySrc20BalanceResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &resp)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessBasicXApiResponse(w, cliCtx, resp)

	}
}

func QuerySrc20HeightLastestHandlerFunc(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hash, preHash string
		var blockIndex int64

		node, err := cliCtx.GetNode()
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		lastestHeight, err := node.LatestBlockNumber()
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		lastestBlock, err := node.Block(&lastestHeight)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		hash = lastestBlock.Block.BtcBlockHash
		blockIndex = lastestBlock.Block.BtcHeight

		prevHeight := lastestHeight - 1
		prevBlock, err := node.Block(&prevHeight)
		if err != nil {
			preHash = "null"
		} else {
			preHash = prevBlock.Block.BtcBlockHash
		}

		resp := xtypes.NewQuerySrc20LatestBlockIndexResponse(blockIndex, hash, preHash)

		PostProcessBasicXApiResponse(w, cliCtx, resp)

	}
}

func QuerySrc20AllTokenInfoHandlerFunc(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req xtypes.Src20AllTokenInfoReq

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = cliCtx.Codec.UnmarshalJSON(body, &req)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//default
		if req.Page == "" {
			req.Page = "1"
		}
		if req.Limit == "" {
			req.Limit = "30"
		}

		bz, err := cliCtx.Codec.MarshalJSON(req)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", xtypes.QuerierRoute, xtypes.Src20QueryAllTick), bz)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var resp xtypes.QuerySrc20AllTickInfoResponse
		err = cliCtx.Codec.UnmarshalJSON(res, &resp)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		PostProcessBasicXApiResponse(w, cliCtx, resp)

	}
}
