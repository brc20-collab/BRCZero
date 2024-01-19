package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/brc20-collab/brczero/app"
	"github.com/brc20-collab/brczero/app/utils/appstatus"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/flatkv"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/mpt"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	tmiavl "github.com/brc20-collab/brczero/libs/iavl"
	"github.com/brc20-collab/brczero/libs/system/trace"
	sm "github.com/brc20-collab/brczero/libs/tendermint/state"
	tmtypes "github.com/brc20-collab/brczero/libs/tendermint/types"
)

const (
	flagChangeCodeHash = "change_codehash"
	flagChangeCode     = "change_code"
)

func repairStateCmd(ctx *server.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repair-state",
		Short: "Repair the SMB(state machine broken) data of node",
		PreRun: func(_ *cobra.Command, _ []string) {
			setExternalPackageValue()
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("--------- repair data start ---------")

			go func() {
				pprofAddress := viper.GetString(pprofAddrFlag)
				err := http.ListenAndServe(pprofAddress, nil)
				if err != nil {
					fmt.Println(err)
				}
			}()
			changeCodeHash := viper.GetString(flagChangeCodeHash)
			changeCode := viper.GetString(flagChangeCode)
			if len(changeCodeHash) != 0 && len(changeCode) != 0 {
				addr, err := hex.DecodeString(changeCodeHash)
				if err != nil {
					fmt.Println("change contract address failed", err)
				} else {
					if code, err := hex.DecodeString(changeCode); err != nil {
						fmt.Println("change code failed: ", err)
					} else {
						app.RepairStateChangeContract(ctx, false, addr, code)
					}
				}

			} else {
				app.RepairState(ctx, false)
			}

			log.Println("--------- repair data success ---------")
		},
	}
	cmd.Flags().Int64(app.FlagStartHeight, 0, "Set the start block height for repair")
	cmd.Flags().Bool(flatkv.FlagEnable, false, "Enable flat kv storage for read performance")
	cmd.Flags().String(app.Elapsed, app.DefaultElapsedSchemas, "schemaName=1|0,,,")
	cmd.Flags().Bool(trace.FlagEnableAnalyzer, false, "Enable auto open log analyzer")
	cmd.Flags().Int(sm.FlagDeliverTxsExecMode, 0, "execution mode for deliver txs, (0:serial[default], 1:deprecated, 2:parallel)")
	cmd.Flags().String(sdk.FlagDBBackend, tmtypes.DBBackend, "Database backend: goleveldb | rocksdb")
	cmd.Flags().StringP(pprofAddrFlag, "p", "0.0.0.0:6060", "Address and port of pprof HTTP server listening")
	cmd.Flags().Bool(tmiavl.FlagIavlDiscardFastStorage, false, "Discard fast storage")
	cmd.Flags().MarkHidden(tmiavl.FlagIavlDiscardFastStorage)
	cmd.Flags().String(flagChangeCodeHash, "", "change contract address")
	cmd.Flags().String(flagChangeCode, "", "change contract address")

	return cmd
}

func setExternalPackageValue() {
	tmiavl.SetForceReadIavl(true)
	isFastStorage := appstatus.IsFastStorageStrategy()
	tmiavl.SetEnableFastStorage(isFastStorage)
	if !isFastStorage &&
		!viper.GetBool(tmiavl.FlagIavlDiscardFastStorage) &&
		appstatus.GetFastStorageVersion() >= viper.GetInt64(app.FlagStartHeight) {
		tmiavl.SetEnableFastStorage(true)
		tmiavl.SetIgnoreAutoUpgrade(true)
	}
	if viper.GetBool(tmiavl.FlagIavlDiscardFastStorage) {
		mpt.DisableSnapshot()
	}
}
