package btc_protocol

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/brc20-collab/brczero/app/rpc/namespaces/eth"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/rest"
	xtypes "github.com/brc20-collab/brczero/x/brcx/types"
)

func registerRuneQueryRoutes(cliCtx context.CLIContext, r *mux.Router, ethApi *eth.PublicEthereumAPI) {
	r.HandleFunc("/runealpha/block/{btcBlockHash}/events", QueryRuneTxsEventsByBtcHashHandlerFunc(cliCtx, ethApi)).Methods("GET")
	r.HandleFunc("/runealpha/tx/{txHash}/events", QueryRuneTxEventsByTxHashHandlerFunc(cliCtx, ethApi)).Methods("GET")
}

func QueryRuneTxsEventsByBtcHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
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

		mintOutputMap := make(map[xtypes.MintOutputKey]map[string]xtypes.RuneBasicInfo)
		burnInputMap := make(map[xtypes.BurnInputKey]map[string]xtypes.RuneBasicInfo)

		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				if len(l.Topics) == 0 {
					continue
				}
				//todo: optimize
				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]interface{}, 0, 1)
				}

				var eventContext interface{}
				switch l.Topics[0] {
				case xtypes.MintRuneTopic0:
					eventContext, err = xtypes.UnpackMintRuneEvent(l.Data)
				case xtypes.BurnRuneTopic0:
					eventContext, err = xtypes.UnpackBurnRuneEvent(l.Data)
				case xtypes.IssueTopic0:
					eventContext, err = xtypes.UnpackIssueEvent(l.Data)
				case xtypes.MintErrTopic0:
					eventContext, err = xtypes.UnpackMintErrEvent(l.Data)
				case xtypes.MintOutputTopic0:
					err = aggregateMintOutputEvents(l.Data, txid, mintOutputMap)
					// do not process single MintOutput event
					continue
				case xtypes.BurnInputTopic0:
					err = aggregateBurnInputEvents(l.Data, txid, burnInputMap)
					// do not process single BurnInput event
					continue
				default:
					continue
				}
				if err != nil {
					rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				resMap[txid] = append(resMap[txid], eventContext)
			}
		}
		// add aggregated MintOutPut and BurnInput events to resMap
		for key, subMap := range mintOutputMap {
			if len(subMap) == 0 {
				continue
			}

			adapter := xtypes.NewRawMintOutputEventAdapter(key.Op, key.OutputId)
			for _, info := range subMap {
				adapter.Mint = append(adapter.Mint, info)
			}
			resMap[key.Txid] = append(resMap[key.Txid], adapter)
		}

		for key, subMap := range burnInputMap {
			if len(subMap) == 0 {
				continue
			}

			adapter := xtypes.NewRawBurnInputEventAdapter(key.Op, key.PreOutputId)
			for _, info := range subMap {
				adapter.Burn = append(adapter.Burn, info)
			}
			resMap[key.Txid] = append(resMap[key.Txid], adapter)
		}

		// format response
		txEventsResp := make([]interface{}, 0)
		for txid, events := range resMap {
			txEventsResp = append(txEventsResp, xtypes.NewQueryRuneAlphaTxEventsResponse(events, txid))
		}

		resp := xtypes.NewQueryRuneAlphaTxEventsByBlockHashResponse(txEventsResp)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func QueryRuneTxEventsByTxHashHandlerFunc(cliCtx context.CLIContext, ethApi *eth.PublicEthereumAPI) http.HandlerFunc {
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

		zeroTxHashBtcTxidMap, err := node.MapTxhashTxid(btcBlockHash, RUNEALPHA)
		if err != nil {
			WriteBasicXApiErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resMap := make(map[string][]interface{})

		mintOutputMap := make(map[xtypes.MintOutputKey]map[string]xtypes.RuneBasicInfo)
		burnInputMap := make(map[xtypes.BurnInputKey]map[string]xtypes.RuneBasicInfo)

		for _, txLogs := range blockLogs {
			for _, l := range txLogs {
				if len(l.Data) == 0 {
					// means this tx has no events
					continue
				}

				if len(l.Topics) == 0 {
					continue
				}
				//todo: optimize
				zeroTxhash := strings.TrimPrefix(l.TxHash.String(), "0x")
				txid := zeroTxHashBtcTxidMap[zeroTxhash]

				if _, ok := resMap[txid]; !ok {
					resMap[txid] = make([]interface{}, 0, 1)
				}

				var eventContext interface{}
				switch l.Topics[0] {
				case xtypes.MintRuneTopic0:
					eventContext, err = xtypes.UnpackMintRuneEvent(l.Data)
				case xtypes.BurnRuneTopic0:
					eventContext, err = xtypes.UnpackBurnRuneEvent(l.Data)
				case xtypes.IssueTopic0:
					eventContext, err = xtypes.UnpackIssueEvent(l.Data)
				case xtypes.MintErrTopic0:
					eventContext, err = xtypes.UnpackMintErrEvent(l.Data)
				case xtypes.MintOutputTopic0:
					err = aggregateMintOutputEvents(l.Data, txid, mintOutputMap)
					// do not process single MintOutput event
					continue
				case xtypes.BurnInputTopic0:
					err = aggregateBurnInputEvents(l.Data, txid, burnInputMap)
					// do not process single BurnInput event
					continue
				default:
					continue
				}
				if err != nil {
					rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}

				resMap[txid] = append(resMap[txid], eventContext)
			}
		}
		// add aggregated MintOutPut and BurnInput events to resMap
		for key, subMap := range mintOutputMap {
			if len(subMap) == 0 {
				continue
			}

			adapter := xtypes.NewRawMintOutputEventAdapter(key.Op, key.OutputId)
			for _, info := range subMap {
				adapter.Mint = append(adapter.Mint, info)
			}
			resMap[key.Txid] = append(resMap[key.Txid], adapter)
		}

		for key, subMap := range burnInputMap {
			if len(subMap) == 0 {
				continue
			}

			adapter := xtypes.NewRawBurnInputEventAdapter(key.Op, key.PreOutputId)
			for _, info := range subMap {
				adapter.Burn = append(adapter.Burn, info)
			}
			resMap[key.Txid] = append(resMap[key.Txid], adapter)
		}

		resp := xtypes.NewQueryRuneAlphaTxEventsResponse(resMap[targetTxid], targetTxid)

		PostProcessBasicXApiResponse(w, cliCtx, resp)
	}
}

func aggregateMintOutputEvents(eventData []byte, btcTxid string, mintOutputMap map[xtypes.MintOutputKey]map[string]xtypes.RuneBasicInfo) error {
	moe, err := xtypes.UnpackMintOutputEvent(eventData)
	if err != nil {
		return err
	}

	id := moe.Mint.Id.String()
	key := xtypes.NewMintOutputKey(btcTxid, moe.Op, moe.OutputId)

	if subMap, ok := mintOutputMap[key]; !ok {
		auxMap := make(map[string]xtypes.RuneBasicInfo)
		auxMap[id] = moe.Mint
		mintOutputMap[key] = auxMap
	} else {
		// This means that some ID of BasicInfo already exists in the subMap,
		// but it's uncertain whether the ID for this event exists.
		if info, ok := subMap[id]; !ok {
			// If the ID for this event not exists, set it.
			subMap[id] = moe.Mint
		} else {
			err = info.AddAmount(moe.Mint)
			if err != nil {
				return err
			}
			subMap[id] = info
		}
	}
	return nil
}

func aggregateBurnInputEvents(eventData []byte, btcTxid string, burnInputMap map[xtypes.BurnInputKey]map[string]xtypes.RuneBasicInfo) error {
	bie, err := xtypes.UnpackBurnInputEvent(eventData)
	if err != nil {
		return err
	}

	id := bie.Burn.Id.String()
	key := xtypes.NewBurnInputKey(btcTxid, bie.Op, bie.PreOutputId)

	if subMap, ok := burnInputMap[key]; !ok {
		auxMap := make(map[string]xtypes.RuneBasicInfo)
		auxMap[id] = bie.Burn
		burnInputMap[key] = auxMap
	} else {
		// This means that some ID of BasicInfo already exists in the subMap,
		// but it's uncertain whether the ID for this event exists.
		if info, ok := subMap[id]; !ok {
			// If the ID for this event not exists, set it.
			subMap[id] = bie.Burn
		} else {
			err = info.AddAmount(bie.Burn)
			if err != nil {
				return err
			}
			subMap[id] = info
		}
	}
	return nil
}
