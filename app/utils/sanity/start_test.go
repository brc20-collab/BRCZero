package sanity

import (
	"testing"

	"github.com/spf13/cobra"

	apptype "github.com/brc20-collab/brczero/app/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/types"
	"github.com/brc20-collab/brczero/libs/tendermint/state"
	sm "github.com/brc20-collab/brczero/libs/tendermint/state"
	"github.com/brc20-collab/brczero/x/evm/watcher"
)

func getCommandNodeModeRpcPruningNothing() *cobra.Command {
	return getCommand([]universeFlag{
		&stringFlag{
			Name:    apptype.FlagNodeMode,
			Default: "",
			Changed: true,
			Value:   string(apptype.RpcNode),
		},
		&stringFlag{
			Name:    server.FlagPruning,
			Default: types.PruningOptionDefault,
			Changed: true,
			Value:   types.PruningOptionNothing,
		},
	})
}

func getCommandFastQueryPruningNothing() *cobra.Command {
	return getCommand([]universeFlag{
		&boolFlag{
			Name:    watcher.FlagFastQuery,
			Default: false,
			Changed: true,
			Value:   true,
		},
		&stringFlag{
			Name:    server.FlagPruning,
			Default: "",
			Changed: true,
			Value:   types.PruningOptionNothing,
		},
	})
}

func getCommandDeliverTxsExecModeSerial(v int) *cobra.Command {
	return getCommand([]universeFlag{
		&intFlag{
			Name:    sm.FlagDeliverTxsExecMode,
			Default: 0,
			Changed: true,
			Value:   v,
		},
	})
}

func TestCheckStart(t *testing.T) {
	tests := []struct {
		name    string
		cmdFunc func()
		wantErr bool
	}{
		{name: "range-TxsExecModeSerial 0", cmdFunc: func() { getCommandDeliverTxsExecModeSerial(int(state.DeliverTxsExecModeSerial)) }, wantErr: false},
		{name: "range-TxsExecModeSerial 1", cmdFunc: func() { getCommandDeliverTxsExecModeSerial(1) }, wantErr: true},
		{name: "range-TxsExecModeSerial 2", cmdFunc: func() { getCommandDeliverTxsExecModeSerial(state.DeliverTxsExecModeParallel) }, wantErr: false},
		{name: "range-TxsExecModeSerial 3", cmdFunc: func() { getCommandDeliverTxsExecModeSerial(3) }, wantErr: true},
		{name: "1. conflicts --fast-query and --pruning=nothing", cmdFunc: func() { getCommandFastQueryPruningNothing() }, wantErr: true},
		{name: "3. conflicts --node-mod=rpc and --pruning=nothing", cmdFunc: func() { getCommandNodeModeRpcPruningNothing() }, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			tt.cmdFunc()
			if err = CheckStart(); (err != nil) != tt.wantErr {
				t.Errorf("CheckStart() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(err)
		})
	}
}

func getCommandNodeModeArchiveFastQuery() *cobra.Command {
	return getCommand([]universeFlag{
		&stringFlag{
			Name:    apptype.FlagNodeMode,
			Default: "",
			Changed: true,
			Value:   string(apptype.ArchiveNode),
		},
		&boolFlag{
			Name:    watcher.FlagFastQuery,
			Default: false,
			Changed: true,
			Value:   true,
		},
	})
}

func TestCheckStartArchive(t *testing.T) {
	type args struct {
		cmd *cobra.Command
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1. conflicts --node-mod=archive and --fast-query", args: args{cmd: getCommandNodeModeArchiveFastQuery()}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if err = CheckStart(); (err != nil) != tt.wantErr {
				t.Errorf("CheckStart() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(err)
		})
	}
}
