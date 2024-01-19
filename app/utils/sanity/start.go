package sanity

import (
	"github.com/spf13/viper"

	"github.com/brc20-collab/brczero/app/config"
	apptype "github.com/brc20-collab/brczero/app/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	cosmost "github.com/brc20-collab/brczero/libs/cosmos-sdk/store/types"
	"github.com/brc20-collab/brczero/libs/tendermint/state"
	"github.com/brc20-collab/brczero/x/evm/watcher"
)

// CheckStart check start command's flags. if user set conflict flags return error.
// the conflicts flags are:
// --fast-query      conflict with --pruning=nothing
// --enable-preruntx conflict with --download-delta
//
// based the conflicts above and node-mode below
// --node-mode=rpc manage the following flags:
//     --disable-checktx-mutex=true
//     --disable-query-mutex=true
//     --enable-bloom-filter=true
//     --fast-lru=10000
//     --fast-query=true
//     --iavl-enable-async-commit=true
//     --max-open=20000
//     --mempool.enable_pending_pool=true
//     --cors=*
//
// --node-mode=validator manage the following flags:
//     --disable-checktx-mutex=true
//     --disable-query-mutex=true
//     --dynamic-gp-mode=2
//     --iavl-enable-async-commit=true
//     --iavl-cache-size=10000000
//     --pruning=everything
//
// --node-mode=archive manage the following flags:
//    --pruning=nothing
//    --disable-checktx-mutex=true
//    --disable-query-mutex=true
//    --enable-bloom-filter=true
//    --iavl-enable-async-commit=true
//    --max-open=20000
//    --cors=*
//
// then
// --node-mode=archive(--pruning=nothing) conflicts with --fast-query

var (
	startDependentElems = []dependentPair{}
	// conflicts flags
	startConflictElems = []conflictPair{
		// --fast-query      conflict with --pruning=nothing
		{
			configA: boolItem{name: watcher.FlagFastQuery, expect: true},
			configB: stringItem{name: server.FlagPruning, expect: cosmost.PruningOptionNothing},
		},
		{
			configA: stringItem{name: apptype.FlagNodeMode, expect: string(apptype.RpcNode)},
			configB: stringItem{name: server.FlagPruning, expect: cosmost.PruningOptionNothing},
		},
		// --node-mode=archive(--pruning=nothing) conflicts with --fast-query
		{
			configA: stringItem{name: apptype.FlagNodeMode, expect: string(apptype.ArchiveNode)},
			configB: boolItem{name: watcher.FlagFastQuery, expect: true},
		},
		{
			configA: stringItem{name: apptype.FlagNodeMode, expect: string(apptype.RpcNode)},
			configB: boolItem{name: config.FlagEnablePGU, expect: true},
		},
		{
			configA: stringItem{name: apptype.FlagNodeMode, expect: string(apptype.ArchiveNode)},
			configB: boolItem{name: config.FlagEnablePGU, expect: true},
		},
		{
			configA: stringItem{name: apptype.FlagNodeMode, expect: string(apptype.InnertxNode)},
			configB: boolItem{name: config.FlagEnablePGU, expect: true},
		},
	}

	checkRangeItems = []rangeItem{
		{
			enumRange: []int{int(state.DeliverTxsExecModeSerial), state.DeliverTxsExecModeParallel},
			name:      state.FlagDeliverTxsExecMode,
		},
	}
)

// CheckStart check start command.If it has conflict pair above. then return the conflict error
func CheckStart() error {
	if viper.GetBool(FlagDisableSanity) {
		return nil
	}
	for _, v := range startDependentElems {
		if err := v.check(); err != nil {
			return err
		}
	}
	for _, v := range startConflictElems {
		if err := v.check(); err != nil {
			return err
		}
	}

	for _, v := range checkRangeItems {
		if err := v.checkRange(); err != nil {
			return err
		}
	}

	return nil
}
