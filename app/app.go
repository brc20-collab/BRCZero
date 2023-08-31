package app

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/client/flags"

	paramstypes "github.com/brc20-collab/brczero/x/params/types"

	"github.com/brc20-collab/brczero/x/vmbridge"

	ica "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/27-interchain-accounts"
	icacontroller "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/27-interchain-accounts/controller"
	icahost "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/27-interchain-accounts/host"
	"github.com/brc20-collab/brczero/x/icamauth"

	ibccommon "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/common"

	icacontrollertypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/27-interchain-accounts/host/types"
	icamauthtypes "github.com/brc20-collab/brczero/x/icamauth/types"

	icacontrollerkeeper "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/27-interchain-accounts/host/keeper"
	icamauthkeeper "github.com/brc20-collab/brczero/x/icamauth/keeper"

	ibcfeekeeper "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/29-fee/keeper"

	icatypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/29-fee/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"

	ibcfee "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/29-fee"

	"github.com/brc20-collab/brczero/app/utils/appstatus"

	"github.com/brc20-collab/brczero/app/ante"
	chaincodec "github.com/brc20-collab/brczero/app/codec"
	appconfig "github.com/brc20-collab/brczero/app/config"
	"github.com/brc20-collab/brczero/app/refund"
	chain "github.com/brc20-collab/brczero/app/types"
	"github.com/brc20-collab/brczero/app/utils/sanity"
	bam "github.com/brc20-collab/brczero/libs/cosmos-sdk/baseapp"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/codec"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/simapp"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/mpt"
	stypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/store/types"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/module"
	upgradetypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/upgrade"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/version"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"
	authtypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/bank"
	capabilityModule "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/capability/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/crisis"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mint"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/supply"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/upgrade"
	"github.com/brc20-collab/brczero/libs/iavl"
	ibctransfer "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/transfer"
	ibctransferkeeper "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/apps/transfer/types"
	ibc "github.com/brc20-collab/brczero/libs/ibc-go/modules/core"
	ibcclient "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/02-client"
	ibcclienttypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/02-client/types"
	ibcporttypes "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/05-port/types"
	ibchost "github.com/brc20-collab/brczero/libs/ibc-go/modules/core/24-host"
	"github.com/brc20-collab/brczero/libs/system"
	"github.com/brc20-collab/brczero/libs/system/trace"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/libs/tendermint/libs/log"
	tmos "github.com/brc20-collab/brczero/libs/tendermint/libs/os"
	sm "github.com/brc20-collab/brczero/libs/tendermint/state"
	tmtypes "github.com/brc20-collab/brczero/libs/tendermint/types"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"
	commonversion "github.com/brc20-collab/brczero/x/common/version"
	distr "github.com/brc20-collab/brczero/x/distribution"
	"github.com/brc20-collab/brczero/x/erc20"
	"github.com/brc20-collab/brczero/x/evidence"
	"github.com/brc20-collab/brczero/x/evm"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"
	"github.com/brc20-collab/brczero/x/feesplit"
	"github.com/brc20-collab/brczero/x/genutil"
	"github.com/brc20-collab/brczero/x/gov"
	"github.com/brc20-collab/brczero/x/gov/keeper"
	"github.com/brc20-collab/brczero/x/infura"
	"github.com/brc20-collab/brczero/x/params"
	"github.com/brc20-collab/brczero/x/slashing"
	"github.com/brc20-collab/brczero/x/staking"
	"github.com/brc20-collab/brczero/x/token"
	"github.com/brc20-collab/brczero/x/wasm"
	wasmkeeper "github.com/brc20-collab/brczero/x/wasm/keeper"
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
		//mint.AppModuleBasic{},
		//distr.AppModuleBasic{},
		//gov.NewAppModuleBasic(
		//	paramsclient.ProposalHandler,
		//	paramsclient.UpgradeProposalHandler,
		//	distr.CommunityPoolSpendProposalHandler,
		//	distr.ChangeDistributionTypeProposalHandler,
		//	distr.WithdrawRewardEnabledProposalHandler,
		//	distr.RewardTruncatePrecisionProposalHandler,
		//	evmclient.ManageContractDeploymentWhitelistProposalHandler,
		//	evmclient.ManageContractBlockedListProposalHandler,
		//	evmclient.ManageContractMethodGuFactorProposalHandler,
		//	evmclient.ManageContractMethodBlockedListProposalHandler,
		//	evmclient.ManageSysContractAddressProposalHandler,
		//	evmclient.ManageContractByteCodeProposalHandler,
		//	govclient.ManageTreasuresProposalHandler,
		//	govclient.ExtraProposalHandler,
		//	erc20client.TokenMappingProposalHandler,
		//	erc20client.ProxyContractRedirectHandler,
		//	erc20client.ContractTemplateProposalHandler,
		//	client.UpdateClientProposalHandler,
		//	fsclient.FeeSplitSharesProposalHandler,
		//	wasmclient.MigrateContractProposalHandler,
		//	wasmclient.UpdateContractAdminProposalHandler,
		//	wasmclient.PinCodesProposalHandler,
		//	wasmclient.UnpinCodesProposalHandler,
		//	wasmclient.UpdateDeploymentWhitelistProposalHandler,
		//	wasmclient.UpdateWASMContractMethodBlockedListProposalHandler,
		//	wasmclient.GetCmdExtraProposal,
		//	stakingclient.ProposeValidatorProposalHandler,
		//),
		params.AppModuleBasic{},
		//crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		evidence.AppModuleBasic{},
		//upgrade.AppModuleBasic{},
		evm.AppModuleBasic{},
		//token.AppModuleBasic{},
		//infura.AppModuleBasic{},
		//capabilityModule.AppModuleBasic{},
		//ibc.AppModuleBasic{},
		//ibctransfer.AppModuleBasic{},
		//erc20.AppModuleBasic{},
		//wasm.AppModuleBasic{},
		//feesplit.AppModuleBasic{},
		//ica.AppModuleBasic{},
		//ibcfee.AppModuleBasic{},
		//icamauth.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		auth.FeeCollectorName:       nil,
		distr.ModuleName:            nil,
		mint.ModuleName:             {supply.Minter},
		staking.BondedPoolName:      {supply.Burner, supply.Staking},
		staking.NotBondedPoolName:   {supply.Burner, supply.Staking},
		gov.ModuleName:              nil,
		token.ModuleName:            {supply.Minter, supply.Burner},
		ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		erc20.ModuleName:            {authtypes.Minter, authtypes.Burner},
		wasm.ModuleName:             nil,
		feesplit.ModuleName:         nil,
		ibcfeetypes.ModuleName:      nil,
		icatypes.ModuleName:         nil,
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
	AccountKeeper        auth.AccountKeeper
	BankKeeper           bank.Keeper
	SupplyKeeper         supply.Keeper
	StakingKeeper        staking.Keeper
	SlashingKeeper       slashing.Keeper
	MintKeeper           mint.Keeper
	DistrKeeper          distr.Keeper
	GovKeeper            gov.Keeper
	CrisisKeeper         crisis.Keeper
	UpgradeKeeper        upgrade.Keeper
	ParamsKeeper         params.Keeper
	EvidenceKeeper       evidence.Keeper
	EvmKeeper            *evm.Keeper
	TokenKeeper          token.Keeper
	WasmKeeper           wasm.Keeper
	WasmPermissionKeeper wasm.ContractOpsKeeper
	InfuraKeeper         infura.Keeper
	FeeSplitKeeper       feesplit.Keeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	configurator module.Configurator
	// ibc
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	ScopedIBCMockKeeper  capabilitykeeper.ScopedKeeper
	TransferKeeper       ibctransferkeeper.Keeper
	CapabilityKeeper     *capabilitykeeper.Keeper
	IBCKeeper            *ibc.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCFeeKeeper         ibcfeekeeper.Keeper
	marshal              *codec.CodecProxy
	heightTasks          map[int64]*upgradetypes.HeightTasks
	Erc20Keeper          erc20.Keeper
	ICAMauthKeeper       icamauthkeeper.Keeper
	ICAControllerKeeper  icacontrollerkeeper.Keeper
	ICAHostKeeper        icahostkeeper.Keeper
	VMBridgeKeeper       *vmbridge.Keeper

	WasmHandler wasmkeeper.HandlerOption
}

// NewBRCZeroApp returns a reference to a new initialized BRCZero application.
func NewBRCZeroApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	invCheckPeriod uint,
	baseAppOptions ...func(*bam.BaseApp),
) *BRCZeroApp {
	logger.Info("Starting " + system.ChainName)
	onceLog.Do(func() {
		iavl.SetLogger(logger.With("module", "iavl"))
		logStartingFlags(logger)
	})

	codecProxy, interfaceReg := chaincodec.MakeCodecSuit(ModuleBasics)
	vmbridge.RegisterInterface(interfaceReg)
	// NOTE we use custom BRCZero transaction decoder that supports the sdk.Tx interface instead of sdk.StdTx
	bApp := bam.NewBaseApp(appName, logger, db, evm.TxDecoder(codecProxy), baseAppOptions...)

	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.SetStartLogHandler(trace.StartTxLog)
	bApp.SetEndLogHandler(trace.StopTxLog)

	bApp.SetInterfaceRegistry(interfaceReg)

	keys := sdk.NewKVStoreKeys(
		bam.MainStoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distr.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, upgrade.StoreKey, evidence.StoreKey,
		token.StoreKey, token.KeyLock,
		ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		ibchost.StoreKey,
		erc20.StoreKey,
		mpt.StoreKey,
		wasm.StoreKey,
		feesplit.StoreKey,
		icacontrollertypes.StoreKey, icahosttypes.StoreKey, ibcfeetypes.StoreKey,
		icamauthtypes.StoreKey,
	)

	tkeys := sdk.NewTransientStoreKeys(params.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &BRCZeroApp{
		BaseApp:        bApp,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tkeys:          tkeys,
		subspaces:      make(map[string]params.Subspace),
		heightTasks:    make(map[int64]*upgradetypes.HeightTasks),
	}
	bApp.SetInterceptors(makeInterceptors())

	// init params keeper and subspaces
	app.ParamsKeeper = params.NewKeeper(codecProxy.GetCdc(), keys[params.StoreKey], tkeys[params.TStoreKey], logger)
	app.subspaces[auth.ModuleName] = app.ParamsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.ParamsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.ParamsKeeper.Subspace(staking.DefaultParamspace)
	app.subspaces[mint.ModuleName] = app.ParamsKeeper.Subspace(mint.DefaultParamspace)
	app.subspaces[distr.ModuleName] = app.ParamsKeeper.Subspace(distr.DefaultParamspace)
	app.subspaces[slashing.ModuleName] = app.ParamsKeeper.Subspace(slashing.DefaultParamspace)
	app.subspaces[gov.ModuleName] = app.ParamsKeeper.Subspace(gov.DefaultParamspace)
	app.subspaces[crisis.ModuleName] = app.ParamsKeeper.Subspace(crisis.DefaultParamspace)
	app.subspaces[evidence.ModuleName] = app.ParamsKeeper.Subspace(evidence.DefaultParamspace)
	app.subspaces[evm.ModuleName] = app.ParamsKeeper.Subspace(evm.DefaultParamspace)
	app.subspaces[token.ModuleName] = app.ParamsKeeper.Subspace(token.DefaultParamspace)
	app.subspaces[ibchost.ModuleName] = app.ParamsKeeper.Subspace(ibchost.ModuleName)
	app.subspaces[ibctransfertypes.ModuleName] = app.ParamsKeeper.Subspace(ibctransfertypes.ModuleName)
	app.subspaces[erc20.ModuleName] = app.ParamsKeeper.Subspace(erc20.DefaultParamspace)
	app.subspaces[wasm.ModuleName] = app.ParamsKeeper.Subspace(wasm.ModuleName)
	app.subspaces[feesplit.ModuleName] = app.ParamsKeeper.Subspace(feesplit.ModuleName)
	app.subspaces[icacontrollertypes.SubModuleName] = app.ParamsKeeper.Subspace(icacontrollertypes.SubModuleName)
	app.subspaces[icahosttypes.SubModuleName] = app.ParamsKeeper.Subspace(icahosttypes.SubModuleName)

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
	app.MintKeeper = mint.NewKeeper(
		codecProxy.GetCdc(), keys[mint.StoreKey], app.subspaces[mint.ModuleName], &stakingKeeper,
		app.SupplyKeeper, auth.FeeCollectorName,
	)
	app.DistrKeeper = distr.NewKeeper(
		codecProxy.GetCdc(), keys[distr.StoreKey], app.subspaces[distr.ModuleName], &stakingKeeper,
		app.SupplyKeeper, auth.FeeCollectorName, app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashing.NewKeeper(
		codecProxy.GetCdc(), keys[slashing.StoreKey], &stakingKeeper, app.subspaces[slashing.ModuleName],
	)
	app.CrisisKeeper = crisis.NewKeeper(
		app.subspaces[crisis.ModuleName], invCheckPeriod, app.SupplyKeeper, auth.FeeCollectorName,
	)
	app.UpgradeKeeper = upgrade.NewKeeper(skipUpgradeHeights, keys[upgrade.StoreKey], app.marshal.GetCdc())
	app.ParamsKeeper.RegisterSignal(evmtypes.SetEvmParamsNeedUpdate)
	app.EvmKeeper = evm.NewKeeper(
		app.marshal.GetCdc(), keys[mpt.StoreKey], app.subspaces[evm.ModuleName], &app.AccountKeeper, app.SupplyKeeper, app.BankKeeper, &stakingKeeper, logger)
	(&bankKeeper).SetInnerTxKeeper(app.EvmKeeper)

	app.TokenKeeper = token.NewKeeper(app.BankKeeper, app.subspaces[token.ModuleName], auth.FeeCollectorName, app.SupplyKeeper,
		keys[token.StoreKey], keys[token.KeyLock], app.marshal.GetCdc(), false, &app.AccountKeeper)

	app.InfuraKeeper = infura.NewKeeper(app.EvmKeeper, logger, streamMetrics)
	// create evidence keeper with router
	evidenceKeeper := evidence.NewKeeper(
		codecProxy.GetCdc(), keys[evidence.StoreKey], app.subspaces[evidence.ModuleName], &app.StakingKeeper, app.SlashingKeeper,
	)
	evidenceRouter := evidence.NewRouter()
	evidenceKeeper.SetRouter(evidenceRouter)
	app.EvidenceKeeper = *evidenceKeeper

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(codecProxy, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	// NOTE: the IBC mock keeper and application module is used only for testing core IBC. Do
	// note replicate if you do not need to test core IBC or light clients.
	scopedIBCMockKeeper := app.CapabilityKeeper.ScopeToModule("mock")
	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	scopedICAMauthKeeper := app.CapabilityKeeper.ScopeToModule(icamauthtypes.ModuleName)

	v2keeper := ibc.NewKeeper(
		codecProxy, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), &stakingKeeper, app.UpgradeKeeper, &scopedIBCKeeper, interfaceReg,
	)
	v4Keeper := ibc.NewV4Keeper(v2keeper, app.ParamsKeeper)
	facadedKeeper := ibc.NewFacadedKeeper(v2keeper)
	facadedKeeper.RegisterKeeper(ibccommon.DefaultFactory(tmtypes.HigherThanVenus4, ibc.IBCV4, v4Keeper))
	app.IBCKeeper = facadedKeeper
	supplyKeeperAdapter := supply.NewSupplyKeeperAdapter(app.SupplyKeeper)
	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		codecProxy, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
		v2keeper.ChannelKeeper, &v2keeper.PortKeeper,
		app.SupplyKeeper, supplyKeeperAdapter, scopedTransferKeeper, interfaceReg,
	)
	ibctransfertypes.SetMarshal(codecProxy)
	app.IBCFeeKeeper = ibcfeekeeper.NewKeeper(codecProxy, keys[ibcfeetypes.StoreKey], app.GetSubspace(ibcfeetypes.ModuleName),
		v2keeper.ChannelKeeper, // may be replaced with IBC middleware
		v2keeper.ChannelKeeper,
		&v2keeper.PortKeeper, app.SupplyKeeper, supplyKeeperAdapter,
	)

	// ICA Controller keeper
	app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		codecProxy, keys[icacontrollertypes.StoreKey], app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack
		app.IBCKeeper.V2Keeper.ChannelKeeper, &app.IBCKeeper.V2Keeper.PortKeeper,
		scopedICAControllerKeeper, app.MsgServiceRouter(),
	)

	// ICA Host keeper
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		codecProxy, keys[icahosttypes.StoreKey], app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.V2Keeper.ChannelKeeper, &app.IBCKeeper.V2Keeper.PortKeeper,
		supplyKeeperAdapter, scopedICAHostKeeper, app.MsgServiceRouter(),
	)

	app.ICAMauthKeeper = icamauthkeeper.NewKeeper(
		codecProxy,
		keys[icamauthtypes.StoreKey],
		app.ICAControllerKeeper,
		scopedICAMauthKeeper,
	)

	app.Erc20Keeper = erc20.NewKeeper(app.marshal.GetCdc(), app.keys[erc20.ModuleName], app.subspaces[erc20.ModuleName],
		app.AccountKeeper, app.SupplyKeeper, app.BankKeeper, app.EvmKeeper, app.TransferKeeper)

	app.FeeSplitKeeper = feesplit.NewKeeper(
		app.keys[feesplit.StoreKey], app.marshal.GetCdc(), app.subspaces[feesplit.ModuleName],
		app.EvmKeeper, app.SupplyKeeper, app.AccountKeeper)
	app.ParamsKeeper.RegisterSignal(feesplit.SetParamsNeedUpdate)

	//wasm keeper
	wasmDir := wasm.WasmDir()
	wasmConfig := wasm.WasmConfig()

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	supportedFeatures := wasm.SupportedFeatures
	app.WasmKeeper = wasm.NewKeeper(
		app.marshal,
		keys[wasm.StoreKey],
		keys[mpt.StoreKey],
		app.subspaces[wasm.ModuleName],
		&app.AccountKeeper,
		bank.NewBankKeeperAdapter(app.BankKeeper),
		v2keeper.ChannelKeeper,
		&v2keeper.PortKeeper,
		nil,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		supportedFeatures,
		vmbridge.GetWasmOpts(app.marshal.GetProtocMarshal()),
	)
	(&app.WasmKeeper).SetInnerTxKeeper(app.EvmKeeper)

	app.ParamsKeeper.RegisterSignal(wasm.SetNeedParamsUpdate)

	// register the proposal types
	// 3.register the proposal types
	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(&app.ParamsKeeper)).
		AddRoute(distr.RouterKey, distr.NewDistributionProposalHandler(app.DistrKeeper)).
		AddRoute(evm.RouterKey, evm.NewManageContractDeploymentWhitelistProposalHandler(app.EvmKeeper)).
		AddRoute(mint.RouterKey, mint.NewManageTreasuresProposalHandler(&app.MintKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientUpdateProposalHandler(app.IBCKeeper.V2Keeper.ClientKeeper)).
		AddRoute(erc20.RouterKey, erc20.NewProposalHandler(&app.Erc20Keeper)).
		AddRoute(feesplit.RouterKey, feesplit.NewProposalHandler(&app.FeeSplitKeeper)).
		AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(&app.WasmKeeper, wasm.NecessaryProposals)).
		AddRoute(params.UpgradeRouterKey, params.NewUpgradeProposalHandler(&app.ParamsKeeper)).
		AddRoute(staking.RouterKey, staking.NewProposalHandler(&app.StakingKeeper))

	govProposalHandlerRouter := keeper.NewProposalHandlerRouter()
	govProposalHandlerRouter.AddRoute(params.RouterKey, &app.ParamsKeeper).
		AddRoute(evm.RouterKey, app.EvmKeeper).
		AddRoute(mint.RouterKey, &app.MintKeeper).
		AddRoute(erc20.RouterKey, &app.Erc20Keeper).
		AddRoute(feesplit.RouterKey, &app.FeeSplitKeeper).
		AddRoute(distr.RouterKey, &app.DistrKeeper).
		AddRoute(params.UpgradeRouterKey, &app.ParamsKeeper)

	app.GovKeeper = gov.NewKeeper(
		app.marshal.GetCdc(), app.keys[gov.StoreKey], app.ParamsKeeper, app.subspaces[gov.DefaultParamspace],
		app.SupplyKeeper, &stakingKeeper, gov.DefaultParamspace, govRouter,
		app.BankKeeper, govProposalHandlerRouter, auth.FeeCollectorName,
	)
	app.ParamsKeeper.SetGovKeeper(app.GovKeeper)
	app.EvmKeeper.SetGovKeeper(app.GovKeeper)
	app.MintKeeper.SetGovKeeper(app.GovKeeper)
	app.Erc20Keeper.SetGovKeeper(app.GovKeeper)
	app.FeeSplitKeeper.SetGovKeeper(app.GovKeeper)
	app.DistrKeeper.SetGovKeeper(app.GovKeeper)

	// Set IBC hooks
	app.TransferKeeper = *app.TransferKeeper.SetHooks(erc20.NewIBCTransferHooks(app.Erc20Keeper))
	transferModule := ibctransfer.NewAppModule(app.TransferKeeper, codecProxy)

	middle := ibctransfer.NewIBCModule(app.TransferKeeper, transferModule)
	right := ibcfee.NewIBCMiddleware(middle, app.IBCFeeKeeper)
	transferStack := ibcporttypes.NewFacadedMiddleware(middle,
		ibccommon.DefaultFactory(tmtypes.HigherThanVenus4, ibc.IBCV4, right),
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()

	var icaControllerStack ibcporttypes.IBCModule
	icaMauthIBCModule := icamauth.NewIBCModule(app.ICAMauthKeeper)
	icaControllerStack = icaMauthIBCModule
	icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, app.ICAControllerKeeper)
	icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerStack, app.IBCFeeKeeper)
	var icaHostStack ibcporttypes.IBCModule
	icaHostStack = icahost.NewIBCModule(app.ICAHostKeeper)
	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, app.IBCFeeKeeper)
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack)
	ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerStack)
	ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostStack)
	ibcRouter.AddRoute(icamauthtypes.ModuleName, icaControllerStack)

	//ibcRouter.AddRoute(ibcmock.ModuleName, mockModule)
	v2keeper.SetRouter(ibcRouter)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	wasmModule := wasm.NewAppModule(*app.marshal, &app.WasmKeeper)
	app.WasmPermissionKeeper = wasmModule.GetPermissionKeeper()
	app.VMBridgeKeeper = vmbridge.NewKeeper(app.marshal, app.Logger(), app.EvmKeeper, app.WasmPermissionKeeper, app.AccountKeeper, app.BankKeeper)
	app.EvmKeeper.SetCallToCM(vmbridge.PrecompileHooks(app.VMBridgeKeeper))
	// Set EVM hooks
	app.EvmKeeper.SetHooks(
		evm.NewMultiEvmHooks(
			evm.NewLogProcessEvmHook(
				erc20.NewSendToIbcEventHandler(app.Erc20Keeper),
				erc20.NewSendNative20ToIbcEventHandler(app.Erc20Keeper),
				vmbridge.NewSendToWasmEventHandler(*app.VMBridgeKeeper),
				vmbridge.NewCallToWasmEventHandler(*app.VMBridgeKeeper),
			),
			app.FeeSplitKeeper.Hooks(),
		),
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
		mint.NewAppModule(app.MintKeeper),
		slashing.NewAppModule(app.SlashingKeeper, app.AccountKeeper, app.StakingKeeper),
		distr.NewAppModule(app.DistrKeeper, app.SupplyKeeper),
		staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.SupplyKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		evm.NewAppModule(app.EvmKeeper, &app.AccountKeeper),
		token.NewAppModule(commonversion.ProtocolVersionV0, app.TokenKeeper, app.SupplyKeeper),
		infura.NewAppModule(app.InfuraKeeper),
		params.NewAppModule(app.ParamsKeeper),
		// ibc
		ibc.NewAppModule(app.IBCKeeper),
		capabilityModule.NewAppModule(codecProxy, *app.CapabilityKeeper),
		transferModule,
		erc20.NewAppModule(app.Erc20Keeper),
		wasmModule,
		feesplit.NewAppModule(app.FeeSplitKeeper),
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		ica.NewAppModule(codecProxy, &app.ICAControllerKeeper, &app.ICAHostKeeper),
		icamauth.NewAppModule(codecProxy, app.ICAMauthKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	app.mm.SetOrderBeginBlockers(
		infura.ModuleName,
		bank.ModuleName, // we must sure bank.beginblocker must be first beginblocker for innerTx. infura can not gengerate tx, so infura can be first in the list.
		capabilitytypes.ModuleName,
		token.ModuleName,
		mint.ModuleName,
		distr.ModuleName,
		slashing.ModuleName,
		staking.ModuleName,
		evidence.ModuleName,
		evm.ModuleName,
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		wasm.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		crisis.ModuleName,
		gov.ModuleName,
		staking.ModuleName,
		wasm.ModuleName,
		evm.ModuleName, // we must sure evm.endblocker must be last endblocker for innerTx.infura can not gengerate tx, so infura can be last in the list.
		infura.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		auth.ModuleName, distr.ModuleName, staking.ModuleName, bank.ModuleName,
		slashing.ModuleName, gov.ModuleName, mint.ModuleName, supply.ModuleName,
		token.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		evm.ModuleName, crisis.ModuleName, genutil.ModuleName, params.ModuleName, evidence.ModuleName,
		erc20.ModuleName,
		wasm.ModuleName,
		feesplit.ModuleName,
		icatypes.ModuleName, ibcfeetypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())
	app.configurator = module.NewConfigurator(app.Codec(), app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)
	app.setupUpgradeModules(false)

	vmbridge.RegisterServices(app.configurator, *app.VMBridgeKeeper)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(app.AccountKeeper),
		bank.NewAppModule(app.BankKeeper, app.AccountKeeper, app.SupplyKeeper),
		supply.NewAppModule(app.SupplyKeeper, app.AccountKeeper),
		gov.NewAppModule(app.GovKeeper, app.SupplyKeeper),
		mint.NewAppModule(app.MintKeeper),
		staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.SupplyKeeper),
		distr.NewAppModule(app.DistrKeeper, app.SupplyKeeper),
		slashing.NewAppModule(app.SlashingKeeper, app.AccountKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper), // NOTE: only used for simulation to generate randomized param change proposals
		ibc.NewAppModule(app.IBCKeeper),
		wasm.NewAppModule(*app.marshal, &app.WasmKeeper),
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.WasmHandler = wasmkeeper.HandlerOption{
		WasmConfig:        &wasmConfig,
		TXCounterStoreKey: keys[wasm.StoreKey],
	}
	app.SetAnteHandler(ante.NewAnteHandler(app.AccountKeeper, app.EvmKeeper, app.SupplyKeeper, validateMsgHook(), app.WasmHandler, app.IBCKeeper, app.StakingKeeper, app.ParamsKeeper))
	app.SetEndBlocker(app.EndBlocker)
	app.SetGasRefundHandler(refund.NewGasRefundHandler(app.AccountKeeper, app.SupplyKeeper, app.EvmKeeper))
	app.SetAccNonceHandler(NewAccNonceHandler(app.AccountKeeper))

	app.SetUpdateWasmTxCount(fixCosmosTxCountInWasmForParallelTx(app.WasmHandler.TXCounterStoreKey))
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
		// Initialize pinned codes in wasmvm as they are not persisted there
		if err := app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit(fmt.Sprintf("failed initialize pinned codes %s", err))
		}
		app.InitUpgrade(ctx)
		app.WasmKeeper.UpdateGasRegister(ctx)
		app.WasmKeeper.UpdateCurBlockNum(ctx)
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	// NOTE: the IBC mock keeper and application module is used only for testing core IBC. Do
	// note replicate if you do not need to test core IBC or light clients.
	app.ScopedIBCMockKeeper = scopedIBCMockKeeper

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
		app.WasmKeeper.UpdateMilestone(ctx, "wasm_v1", info.EffectiveHeight)
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