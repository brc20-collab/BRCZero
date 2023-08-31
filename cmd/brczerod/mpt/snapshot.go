package mpt

import (
	stdlog "log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/brc20-collab/brczero/cmd/brczerod/base"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/mpt"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/rootmulti"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	cfg "github.com/brc20-collab/brczero/libs/tendermint/config"
	tmflags "github.com/brc20-collab/brczero/libs/tendermint/libs/cli/flags"
	"github.com/brc20-collab/brczero/libs/tendermint/libs/log"
)

func genSnapCmd(ctx *server.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-snapshot",
		Short: "generate mpt store's snapshot",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			stdlog.Println("--------- generate snapshot start ---------")
			dataDir := filepath.Join(ctx.Config.RootDir, "data")
			GenSnapshot(dataDir)
			stdlog.Println("--------- generate snapshot end ---------")
		},
	}
	return cmd
}

func GenSnapshot(dataDir string) {
	db, err := sdk.NewDB(applicationDB, dataDir)
	if err != nil {
		panic("fail to open application db: " + err.Error())
	}
	defer db.Close()

	mpt.SetSnapshotRebuild(true)
	mpt.AccountStateRootRetriever = base.AccountStateRootRetriever{}
	rs := rootmulti.NewStore(db)
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	const logLevel = "main:info,iavl:info,*:error,state:info,provider:info,root-multi:info"
	logger, err = tmflags.ParseLogLevel(logLevel, logger, cfg.DefaultLogLevel())
	rs.SetLogger(logger)
	rs.MountStoreWithDB(sdk.NewKVStoreKey(mpt.StoreKey), sdk.StoreTypeMPT, nil)
	rs.LoadLatestVersion()
}
