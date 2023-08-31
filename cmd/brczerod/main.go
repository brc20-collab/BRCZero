package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/brc20-collab/brczero/libs/system"

	"github.com/brc20-collab/brczero/app/logevents"
	"github.com/brc20-collab/brczero/cmd/brczerod/fss"
	"github.com/brc20-collab/brczero/cmd/brczerod/mpt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/brc20-collab/brczero/app/rpc"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"

	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	tmamino "github.com/brc20-collab/brczero/libs/tendermint/crypto/encoding/amino"
	"github.com/brc20-collab/brczero/libs/tendermint/crypto/multisig"
	"github.com/brc20-collab/brczero/libs/tendermint/libs/cli"
	"github.com/brc20-collab/brczero/libs/tendermint/libs/log"
	tmtypes "github.com/brc20-collab/brczero/libs/tendermint/types"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/baseapp"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/flags"
	clientkeys "github.com/brc20-collab/brczero/libs/cosmos-sdk/client/keys"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/crypto/keys"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"

	"github.com/brc20-collab/brczero/app"
	"github.com/brc20-collab/brczero/app/codec"
	"github.com/brc20-collab/brczero/app/crypto/ethsecp256k1"
	chain "github.com/brc20-collab/brczero/app/types"
	"github.com/brc20-collab/brczero/cmd/client"
	"github.com/brc20-collab/brczero/x/genutil"
	genutilcli "github.com/brc20-collab/brczero/x/genutil/client/cli"
	genutiltypes "github.com/brc20-collab/brczero/x/genutil/types"
	"github.com/brc20-collab/brczero/x/staking"
)

const flagInvCheckPeriod = "inv-check-period"
const brczeroEnvPrefix = system.EnvPrefix

var invCheckPeriod uint

func main() {
	cobra.EnableCommandSorting = false

	codecProxy, registry := codec.MakeCodecSuit(app.ModuleBasics)

	tmamino.RegisterKeyType(ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName)
	tmamino.RegisterKeyType(ethsecp256k1.PrivKey{}, ethsecp256k1.PrivKeyName)
	multisig.RegisterKeyType(ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName)

	keys.CryptoCdc = codecProxy.GetCdc()
	genutil.ModuleCdc = codecProxy.GetCdc()
	genutiltypes.ModuleCdc = codecProxy.GetCdc()
	clientkeys.KeysCdc = codecProxy.GetCdc()

	config := sdk.GetConfig()
	chain.SetBech32Prefixes(config)
	chain.SetBip44CoinType(config)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               system.Server,
		Short:             "brczero App Daemon (server)",
		PersistentPreRunE: preRun(ctx),
	}
	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		client.ValidateChainID(
			genutilcli.InitCmd(ctx, codecProxy.GetCdc(), app.ModuleBasics, app.DefaultNodeHome),
		),
		genutilcli.CollectGenTxsCmd(ctx, codecProxy.GetCdc(), auth.GenesisAccountIterator{}, app.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(ctx, codecProxy.GetCdc()),
		genutilcli.GenTxCmd(
			ctx, codecProxy.GetCdc(), app.ModuleBasics, staking.AppModuleBasic{}, auth.GenesisAccountIterator{},
			app.DefaultNodeHome, app.DefaultCLIHome,
		),
		genutilcli.ValidateGenesisCmd(ctx, codecProxy.GetCdc(), app.ModuleBasics),
		client.TestnetCmd(ctx, codecProxy.GetCdc(), app.ModuleBasics, auth.GenesisAccountIterator{}),
		replayCmd(ctx, client.RegisterAppFlag, codecProxy, newApp, registry, registerRoutes),
		repairStateCmd(ctx),
		displayStateCmd(ctx),
		mpt.MptCmd(ctx),
		fss.Command(ctx),
		// AddGenesisAccountCmd allows users to add accounts to the genesis file
		AddGenesisAccountCmd(ctx, codecProxy.GetCdc(), app.DefaultNodeHome, app.DefaultCLIHome),
		flags.NewCompletionCmd(rootCmd, true),
		dataCmd(ctx),
		exportAppCmd(ctx),
		iaviewerCmd(ctx, codecProxy.GetCdc()),
		subscribeCmd(codecProxy.GetCdc()),
	)

	subFunc := func(logger log.Logger) log.Subscriber {
		return logevents.NewProvider(logger)
	}

	// Tendermint node base commands
	server.AddCommands(ctx, codecProxy, registry, rootCmd, newApp, closeApp, exportAppStateAndTMValidators,
		registerRoutes, client.RegisterAppFlag, app.PreRun, subFunc)

	// precheck flag syntax
	preCheckLongFlagSyntax()

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, brczeroEnvPrefix, app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	rootCmd.PersistentFlags().Bool(server.FlagGops, false, "Enable gops metrics collection")

	initEnv()
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func initEnv() {
	checkSetEnv("mempool_size", "200000")
	checkSetEnv("mempool_cache_size", "300000")
	checkSetEnv("mempool_force_recheck_gap", "2000")
	checkSetEnv("mempool_recheck", "false")
	checkSetEnv("consensus_timeout_commit", fmt.Sprintf("%dms", tmtypes.TimeoutCommit))
}

func checkSetEnv(envName string, value string) {
	realEnvName := brczeroEnvPrefix + "_" + strings.ToUpper(envName)
	_, ok := os.LookupEnv(realEnvName)
	if !ok {
		_ = os.Setenv(realEnvName, value)
	}
}

func closeApp(iApp abci.Application) {
	fmt.Println("Close App")
	app := iApp.(*app.BRCZeroApp)
	app.StopBaseApp()
	evmtypes.CloseIndexer()
	rpc.CloseEthBackend()
	app.EvmKeeper.Watcher.Stop()
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	pruningOpts, err := server.GetPruningOptionsFromFlags()
	if err != nil {
		panic(err)
	}

	return app.NewBRCZeroApp(
		logger,
		db,
		traceStore,
		true,
		map[int64]bool{},
		viper.GetUint(flagInvCheckPeriod),
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	var ethermintApp *app.BRCZeroApp
	if height != -1 {
		ethermintApp = app.NewBRCZeroApp(logger, db, traceStore, false, map[int64]bool{}, 0)

		if err := ethermintApp.LoadHeight(height); err != nil {
			return nil, nil, err
		}
	} else {
		ethermintApp = app.NewBRCZeroApp(logger, db, traceStore, true, map[int64]bool{}, 0)
	}

	return ethermintApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}

// All long flag must be in k=v format
func preCheckLongFlagSyntax() {
	params := os.Args[1:]
	for _, f := range params {
		tf := strings.TrimSpace(f)

		if strings.ToUpper(tf) == "TRUE" ||
			strings.ToUpper(tf) == "FALSE" {
			fmt.Fprintf(os.Stderr, "ERROR: Invalid parameter,"+
				" boolean flag should be --flag=true or --flag=false \n")
			os.Exit(1)
		}
	}
}

func preRun(ctx *server.Context) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		setReplayDefaultFlag()
		return server.PersistentPreRunEFn(ctx)(cmd, args)
	}
}
