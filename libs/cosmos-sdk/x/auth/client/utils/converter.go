package utils

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	cm40types "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/ibc-adapter"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
)

func ConvCM39SimulateResultTCM40(cm39 *types.Result) *cm40types.Result {
	ret := &cm40types.Result{
		Data:   cm39.Data,
		Log:    cm39.Log,
		Events: ConvCM39EventsTCM40(cm39.Events),
	}
	return ret
}
func ConvCM39EventsTCM40(es []types.Event) []abci.Event {
	ret := make([]abci.Event, 0)
	for _, v := range es {
		eve := abci.Event{
			Type:       v.Type,
			Attributes: v.Attributes,
		}
		ret = append(ret, eve)
	}

	return ret
}
