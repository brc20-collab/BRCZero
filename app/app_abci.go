package app

import (
	"runtime"
	"time"

	appconfig "github.com/brc20-collab/brczero/app/config"
	"github.com/brc20-collab/brczero/libs/system/trace"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
)

// BeginBlock implements the Application interface
func (app *BRCZeroApp) BeginBlock(req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {
	trace.OnAppBeginBlockEnter(app.LastBlockHeight() + 1)
	app.EvmKeeper.Watcher.DelayEraseKey()
	return app.BaseApp.BeginBlock(req)
}

func (app *BRCZeroApp) DeliverTx(req abci.RequestDeliverTx) (res abci.ResponseDeliverTx) {

	trace.OnAppDeliverTxEnter()

	resp := app.BaseApp.DeliverTx(req)

	return resp
}

func (app *BRCZeroApp) PreDeliverRealTx(req []byte) (res abci.TxEssentials) {
	return app.BaseApp.PreDeliverRealTx(req)
}

func (app *BRCZeroApp) DeliverRealTx(req abci.TxEssentials) (res abci.ResponseDeliverTx) {
	trace.OnAppDeliverTxEnter()
	resp := app.BaseApp.DeliverRealTx(req)
	app.EvmKeeper.Watcher.RecordTxAndFailedReceipt(req, &resp, app.GetTxDecoder())

	return resp
}

// EndBlock implements the Application interface
func (app *BRCZeroApp) EndBlock(req abci.RequestEndBlock) (res abci.ResponseEndBlock) {
	return app.BaseApp.EndBlock(req)
}

// Commit implements the Application interface
func (app *BRCZeroApp) Commit(req abci.RequestCommit) abci.ResponseCommit {
	if gcInterval := appconfig.GetBRCZeroConfig().GetGcInterval(); gcInterval > 0 {
		if (app.BaseApp.LastBlockHeight()+1)%int64(gcInterval) == 0 {
			startTime := time.Now()
			runtime.GC()
			elapsed := time.Now().Sub(startTime).Milliseconds()
			app.Logger().Info("force gc for debug", "height", app.BaseApp.LastBlockHeight()+1,
				"elapsed(ms)", elapsed)
		}
	}
	//defer trace.GetTraceSummary().Dump()
	defer trace.OnCommitDone()

	res := app.BaseApp.Commit(req)

	// we call watch#Commit here ,because
	// 1. this round commit a valid block
	// 2. before commit the block,State#updateToState hasent not called yet,so the proposalBlockPart is not nil which means we wont
	// 	  call the prerun during commit step(edge case)
	app.EvmKeeper.Watcher.Commit()

	return res
}
