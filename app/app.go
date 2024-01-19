package app

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"

	"github.com/brc20-collab/brczero/x/brcx"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"

	"github.com/brc20-collab/brczero/app/ante"
	chaincodec "github.com/brc20-collab/brczero/app/codec"
	appconfig "github.com/brc20-collab/brczero/app/config"
	"github.com/brc20-collab/brczero/app/refund"
	chain "github.com/brc20-collab/brczero/app/types"
	"github.com/brc20-collab/brczero/app/utils/appstatus"
	"github.com/brc20-collab/brczero/app/utils/sanity"
	bam "github.com/brc20-collab/brczero/libs/cosmos-sdk/baseapp"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/flags"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/simapp"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/mpt"
	stypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/store/types"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/module"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/version"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"
	authtypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/bank"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/crisis"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/supply"
	"github.com/brc20-collab/brczero/libs/iavl"
	"github.com/brc20-collab/brczero/libs/system"
	"github.com/brc20-collab/brczero/libs/system/trace"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/libs/tendermint/libs/log"
	tmos "github.com/brc20-collab/brczero/libs/tendermint/libs/os"
	sm "github.com/brc20-collab/brczero/libs/tendermint/state"
	tmtypes "github.com/brc20-collab/brczero/libs/tendermint/types"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"
	"github.com/brc20-collab/brczero/x/evidence"
	"github.com/brc20-collab/brczero/x/evm"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"
	"github.com/brc20-collab/brczero/x/genutil"
	"github.com/brc20-collab/brczero/x/gov"
	"github.com/brc20-collab/brczero/x/gov/keeper"
	"github.com/brc20-collab/brczero/x/params"
	paramsclient "github.com/brc20-collab/brczero/x/params/client"
	paramstypes "github.com/brc20-collab/brczero/x/params/types"
	"github.com/brc20-collab/brczero/x/slashing"
	"github.com/brc20-collab/brczero/x/staking"
	stakingclient "github.com/brc20-collab/brczero/x/staking/client"
)

func init() {
	// set the address prefixes
	config := sdk.GetConfig()
	chain.SetBech32Prefixes(config)
	chain.SetBip44CoinType(config)
}

const (
	appName = system.AppName
)

var (
	// DefaultCLIHome sets the default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv(system.ClientHome)

	// DefaultNodeHome sets the folder where the applcation data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv(system.ServerHome)

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		supply.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler,
			stakingclient.ProposeValidatorProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		evidence.AppModuleBasic{},
		evm.AppModuleBasic{},
		brcx.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		auth.FeeCollectorName:     nil,
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		gov.ModuleName:            nil,
		brcx.ModuleName:           nil,
	}

	onceLog              sync.Once
	FlagGolangMaxThreads string = "golang-max-threads"
)

var _ simapp.App = (*BRCZeroApp)(nil)

// BRCZeroApp implements an extended ABCI application. It is an application
// that may process transactions through Ethereum's EVM running atop of
// Tendermint consensus.
type BRCZeroApp struct {
	*bam.BaseApp

	invCheckPeriod uint

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// subspaces
	subspaces map[string]params.Subspace

	// keepers
	AccountKeeper  auth.AccountKeeper
	BankKeeper     bank.Keeper
	SupplyKeeper   supply.Keeper
	StakingKeeper  staking.Keeper
	SlashingKeeper slashing.Keeper
	GovKeeper      gov.Keeper
	CrisisKeeper   crisis.Keeper
	ParamsKeeper   params.Keeper
	EvidenceKeeper evidence.Keeper
	EvmKeeper      *evm.Keeper
	BRCXKeeper     *brcx.Keeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	configurator module.Configurator
	marshal      *codec.CodecProxy
}

// NewBRCZeroApp returns a reference to a new initialized BRCZero application.
func NewBRCZeroApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	invCheckPeriod uint,
	baseAppOptions ...func(*bam.BaseApp),
) *BRCZeroApp {
	logger.Info("Starting " + system.ChainName)
	onceLog.Do(func() {
		iavl.SetLogger(logger.With("module", "iavl"))
		logStartingFlags(logger)
	})

	codecProxy, interfaceReg := chaincodec.MakeCodecSuit(ModuleBasics)
	// NOTE we use custom BRCZero transaction decoder that supports the sdk.Tx interface instead of sdk.StdTx
	bApp := bam.NewBaseApp(appName, logger, db, evm.TxDecoder(codecProxy), baseAppOptions...)

	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.SetStartLogHandler(trace.StartTxLog)
	bApp.SetEndLogHandler(trace.StopTxLog)

	bApp.SetInterfaceRegistry(interfaceReg)

	keys := sdk.NewKVStoreKeys(
		bam.MainStoreKey,
		staking.StoreKey,
		supply.StoreKey,
		slashing.StoreKey,
		gov.StoreKey,
		params.StoreKey,
		evidence.StoreKey,
		mpt.StoreKey,
		brcx.StoreKey,
	)

	tkeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	app := &BRCZeroApp{
		BaseApp:        bApp,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tkeys:          tkeys,
		subspaces:      make(map[string]params.Subspace),
	}
	bApp.SetInterceptors(makeInterceptors())

	// init params keeper and subspaces
	app.ParamsKeeper = params.NewKeeper(codecProxy.GetCdc(), keys[params.StoreKey], tkeys[params.TStoreKey], logger)
	app.subspaces[auth.ModuleName] = app.ParamsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.ParamsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.ParamsKeeper.Subspace(staking.DefaultParamspace)
	app.subspaces[slashing.ModuleName] = app.ParamsKeeper.Subspace(slashing.DefaultParamspace)
	app.subspaces[gov.ModuleName] = app.ParamsKeeper.Subspace(gov.DefaultParamspace)
	app.subspaces[crisis.ModuleName] = app.ParamsKeeper.Subspace(crisis.DefaultParamspace)
	app.subspaces[evidence.ModuleName] = app.ParamsKeeper.Subspace(evidence.DefaultParamspace)
	app.subspaces[evm.ModuleName] = app.ParamsKeeper.Subspace(evm.DefaultParamspace)

	//proxy := codec.NewMarshalProxy(cc, cdc)
	app.marshal = codecProxy
	// use custom BRCZero account for contracts
	app.AccountKeeper = auth.NewAccountKeeper(
		codecProxy.GetCdc(), keys[mpt.StoreKey], app.subspaces[auth.ModuleName], chain.ProtoAccount,
	)

	bankKeeper := bank.NewBaseKeeperWithMarshal(
		&app.AccountKeeper, codecProxy, app.subspaces[bank.ModuleName], app.ModuleAccountAddrs(),
	)
	app.BankKeeper = &bankKeeper
	app.ParamsKeeper.SetBankKeeper(app.BankKeeper)
	app.SupplyKeeper = supply.NewKeeper(
		codecProxy.GetCdc(), keys[supply.StoreKey], &app.AccountKeeper, bank.NewBankKeeperAdapter(app.BankKeeper), maccPerms,
	)

	stakingKeeper := staking.NewKeeper(
		codecProxy, keys[staking.StoreKey], app.SupplyKeeper, app.subspaces[staking.ModuleName],
	)
	app.ParamsKeeper.SetStakingKeeper(stakingKeeper)
	app.SlashingKeeper = slashing.NewKeeper(
		codecProxy.GetCdc(), keys[slashing.StoreKey], &stakingKeeper, app.subspaces[slashing.ModuleName],
	)
	app.CrisisKeeper = crisis.NewKeeper(
		app.subspaces[crisis.ModuleName], invCheckPeriod, app.SupplyKeeper, auth.FeeCollectorName,
	)
	app.ParamsKeeper.RegisterSignal(evmtypes.SetEvmParamsNeedUpdate)
	app.EvmKeeper = evm.NewKeeper(
		app.marshal.GetCdc(), keys[mpt.StoreKey], app.subspaces[evm.ModuleName], &app.AccountKeeper, app.SupplyKeeper, app.BankKeeper, &stakingKeeper, logger)
	(&bankKeeper).SetInnerTxKeeper(app.EvmKeeper)
	// create evidence keeper with router
	evidenceKeeper := evidence.NewKeeper(
		codecProxy.GetCdc(), keys[evidence.StoreKey], app.subspaces[evidence.ModuleName], &app.StakingKeeper, app.SlashingKeeper,
	)
	evidenceRouter := evidence.NewRouter()
	evidenceKeeper.SetRouter(evidenceRouter)
	app.EvidenceKeeper = *evidenceKeeper

	app.BRCXKeeper = brcx.NewKeeper(codecProxy, keys[brcx.StoreKey], logger, app.EvmKeeper, app.AccountKeeper, app.BankKeeper, app.SupplyKeeper)

	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(&app.ParamsKeeper)).
		AddRoute(staking.RouterKey, staking.NewProposalHandler(&app.StakingKeeper))

	govProposalHandlerRouter := keeper.NewProposalHandlerRouter()
	govProposalHandlerRouter.AddRoute(params.RouterKey, &app.ParamsKeeper)

	app.GovKeeper = gov.NewKeeper(
		app.marshal.GetCdc(), app.keys[gov.StoreKey], app.ParamsKeeper, app.subspaces[gov.DefaultParamspace],
		app.SupplyKeeper, &stakingKeeper, gov.DefaultParamspace, govRouter,
		app.BankKeeper, govProposalHandlerRouter, auth.FeeCollectorName,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(app.SlashingKeeper.Hooks()),
	)

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.AccountKeeper),
		bank.NewAppModule(app.BankKeeper, app.AccountKeeper, app.SupplyKeeper),
		crisis.NewAppModule(&app.CrisisKeeper),
		supply.NewAppModule(app.SupplyKeeper, app.AccountKeeper),
		gov.NewAppModule(app.GovKeeper, app.SupplyKeeper),
		slashing.NewAppModule(app.SlashingKeeper, app.AccountKeeper, app.StakingKeeper),
		staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.SupplyKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		evm.NewAppModule(app.EvmKeeper, &app.AccountKeeper),
		brcx.NewAppModule(*app.BRCXKeeper),
		params.NewAppModule(app.ParamsKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	app.mm.SetOrderBeginBlockers(
		bank.ModuleName,
		slashing.ModuleName,
		staking.ModuleName,
		evidence.ModuleName,
		evm.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		crisis.ModuleName,
		gov.ModuleName,
		staking.ModuleName,
		evm.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		auth.ModuleName,
		staking.ModuleName,
		bank.ModuleName,
		slashing.ModuleName,
		gov.ModuleName,
		supply.ModuleName,
		evm.ModuleName,
		crisis.ModuleName,
		genutil.ModuleName,
		params.ModuleName,
		evidence.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())
	app.configurator = module.NewConfigurator(app.Codec(), app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(app.AccountKeeper),
		bank.NewAppModule(app.BankKeeper, app.AccountKeeper, app.SupplyKeeper),
		supply.NewAppModule(app.SupplyKeeper, app.AccountKeeper),
		gov.NewAppModule(app.GovKeeper, app.SupplyKeeper),
		staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.SupplyKeeper),
		slashing.NewAppModule(app.SlashingKeeper, app.AccountKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper), // NOTE: only used for simulation to generate randomized param change proposals
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(ante.NewAnteHandler(app.AccountKeeper, app.EvmKeeper, app.SupplyKeeper, validateMsgHook()))
	app.SetEndBlocker(app.EndBlocker)
	app.SetGasRefundHandler(refund.NewGasRefundHandler(app.AccountKeeper, app.SupplyKeeper, app.EvmKeeper))
	app.SetAccNonceHandler(NewAccNonceHandler(app.AccountKeeper))

	//todo: delete useless Handler
	app.SetUpdateFeeCollectorAccHandler(updateFeeCollectorHandler(app.BankKeeper, app.SupplyKeeper))
	app.SetParallelTxLogHandlers(fixLogForParallelTxHandler(app.EvmKeeper))
	app.SetPreDeliverTxHandler(preDeliverTxHandler(app.AccountKeeper))
	app.SetPartialConcurrentHandlers(getTxFeeAndFromHandler(app.EvmKeeper))
	app.SetGetTxFeeHandler(getTxFeeHandler())
	app.SetEvmSysContractAddressHandler(NewEvmSysContractAddressHandler(app.EvmKeeper))
	app.SetEvmWatcherCollector(app.EvmKeeper.Watcher.Collect)
	app.SetUpdateCMTxNonceHandler(NewUpdateCMTxNonceHandler())
	app.SetGetGasConfigHandler(NewGetGasConfigHandler(app.ParamsKeeper))
	app.SetGetBlockConfigHandler(NewGetBlockConfigHandler(app.ParamsKeeper))
	mpt.AccountStateRootRetriever = app.AccountKeeper
	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			tmos.Exit(err.Error())
		}
		ctx := app.BaseApp.NewContext(true, abci.Header{})
		app.InitUpgrade(ctx)
	}

	enableAnalyzer := sm.DeliverTxsExecMode(viper.GetInt(sm.FlagDeliverTxsExecMode)) == sm.DeliverTxsExecModeSerial
	trace.EnableAnalyzer(enableAnalyzer)

	return app
}

func (app *BRCZeroApp) InitUpgrade(ctx sdk.Context) {
	// Claim before ApplyEffectiveUpgrade
	app.ParamsKeeper.ClaimReadyForUpgrade(tmtypes.MILESTONE_EARTH, func(info paramstypes.UpgradeInfo) {
		tmtypes.InitMilestoneEarthHeight(int64(info.EffectiveHeight))
	})
	app.ParamsKeeper.ClaimReadyForUpgrade(tmtypes.MILESTONE_MERCURY, func(info paramstypes.UpgradeInfo) {
		tmtypes.InitMilestoneMercuryHeight(int64(info.EffectiveHeight))
	})
	app.ParamsKeeper.ClaimReadyForUpgrade(tmtypes.MILESTONE_VENUS7_NAME, func(info paramstypes.UpgradeInfo) {
		tmtypes.InitMilestoneVenus7Height(int64(info.EffectiveHeight))
	})
	if err := app.ParamsKeeper.ApplyEffectiveUpgrade(ctx); err != nil {
		tmos.Exit(fmt.Sprintf("failed apply effective upgrade height info: %s", err))
	}
}

func (app *BRCZeroApp) SetOption(req abci.RequestSetOption) (res abci.ResponseSetOption) {
	if req.Key == "CheckChainID" {
		if err := chain.IsValidateChainIdWithGenesisHeight(req.Value); err != nil {
			app.Logger().Error(err.Error())
			panic(err)
		}
		err := chain.SetChainId(req.Value)
		if err != nil {
			app.Logger().Error(err.Error())
			panic(err)
		}
	}
	return app.BaseApp.SetOption(req)
}

func (app *BRCZeroApp) LoadStartVersion(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

// Name returns the name of the App
func (app *BRCZeroApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker updates every begin block
func (app *BRCZeroApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker updates every end block
func (app *BRCZeroApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer updates at chain initialization
func (app *BRCZeroApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {

	var genesisState simapp.GenesisState
	app.marshal.GetCdc().MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return app.mm.InitGenesis(ctx, genesisState)
}

// LoadHeight loads state at a particular height
func (app *BRCZeroApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *BRCZeroApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// SimulationManager implements the SimulationApp interface
func (app *BRCZeroApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *BRCZeroApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// Codec returns BRCZero's codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *BRCZeroApp) Codec() *codec.Codec {
	return app.marshal.GetCdc()
}

func (app *BRCZeroApp) Marshal() *codec.CodecProxy {
	return app.marshal
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *BRCZeroApp) GetSubspace(moduleName string) params.Subspace {
	return app.subspaces[moduleName]
}

var protoCodec = encoding.GetCodec(proto.Name)

func makeInterceptors() map[string]bam.Interceptor {
	m := make(map[string]bam.Interceptor)
	m["/cosmos.tx.v1beta1.Service/Simulate"] = bam.NewRedirectInterceptor("app/simulate")
	m["/cosmos.bank.v1beta1.Query/AllBalances"] = bam.NewRedirectInterceptor("custom/bank/grpc_balances")
	m["/cosmos.staking.v1beta1.Query/Params"] = bam.NewRedirectInterceptor("custom/staking/params4ibc")
	return m
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}

	return dupMaccPerms
}

func validateMsgHook() ante.ValidateMsgHandler {
	return func(newCtx sdk.Context, msgs []sdk.Msg) error {

		wrongMsgErr := sdk.ErrUnknownRequest(
			"It is not allowed that a transaction with more than one message contains order or evm message")
		var err error

		for _, msg := range msgs {
			switch msg.(type) {
			case *evmtypes.MsgEthereumTx:
				if len(msgs) > 1 {
					return wrongMsgErr
				}
			}

			if err != nil {
				return err
			}
		}
		return nil
	}
}

func NewAccNonceHandler(ak auth.AccountKeeper) sdk.AccNonceHandler {
	return func(
		ctx sdk.Context, addr sdk.AccAddress,
	) uint64 {
		if acc := ak.GetAccount(ctx, addr); acc != nil {
			return acc.GetSequence()
		}
		return 0
	}
}

func PreRun(ctx *server.Context, cmd *cobra.Command) error {

	prepareSnapshotDataIfNeed(viper.GetString(server.FlagStartFromSnapshot), viper.GetString(flags.FlagHome), ctx.Logger)

	// check start flag conflicts
	err := sanity.CheckStart()
	if err != nil {
		return err
	}

	if maxThreads := viper.GetInt(FlagGolangMaxThreads); maxThreads != 0 {
		debug.SetMaxThreads(maxThreads)
	}
	// set config by node mode
	err = setNodeConfig(ctx)
	if err != nil {
		return err
	}

	//download pprof
	appconfig.PprofDownload(ctx)

	// pruning options
	_, err = server.GetPruningOptionsFromFlags()
	if err != nil {
		return err
	}
	// repair state on start
	if viper.GetBool(FlagEnableRepairState) {
		repairStateOnStart(ctx)
	}

	// init tx signature cache
	tmtypes.InitSignatureCache()

	isFastStorage := appstatus.IsFastStorageStrategy()
	iavl.SetEnableFastStorage(isFastStorage)
	// set external package flags
	server.SetExternalPackageValue(cmd)

	ctx.Logger.Info("The database storage strategy", "fast-storage", iavl.GetEnableFastStorage())
	// set the dynamic config
	appconfig.RegisterDynamicConfig(ctx.Logger.With("module", "config"))

	return nil
}

func NewEvmSysContractAddressHandler(ak *evm.Keeper) sdk.EvmSysContractAddressHandler {
	if ak == nil {
		panic("NewEvmSysContractAddressHandler ak is nil")
	}
	return func(
		ctx sdk.Context, addr sdk.AccAddress,
	) bool {
		if addr.Empty() {
			return false
		}
		return ak.IsMatchSysContractAddress(ctx, addr)
	}
}

func NewUpdateCMTxNonceHandler() sdk.UpdateCMTxNonceHandler {
	return func(tx sdk.Tx, nonce uint64) {
		if nonce != 0 {
			switch v := tx.(type) {
			case *authtypes.StdTx:
				v.Nonce = nonce
			case *authtypes.IbcTx:
				v.Nonce = nonce
			}
		}
	}
}

func NewGetGasConfigHandler(pk params.Keeper) sdk.GetGasConfigHandler {
	return func(ctx sdk.Context) *stypes.GasConfig {
		return pk.GetGasConfig(ctx)
	}
}

func NewGetBlockConfigHandler(pk params.Keeper) sdk.GetBlockConfigHandler {
	return func(ctx sdk.Context) *sdk.BlockConfig {
		return pk.GetBlockConfig(ctx)
	}
}
