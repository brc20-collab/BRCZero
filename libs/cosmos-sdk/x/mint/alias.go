package mint

// nolint

import (
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mint/internal/keeper"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/mint/internal/types"
	govtypes "github.com/brc20-collab/brczero/x/gov/types"
)

const (
	ModuleName            = types.ModuleName
	DefaultParamspace     = types.DefaultParamspace
	StoreKey              = types.StoreKey
	QuerierRoute          = types.QuerierRoute
	QueryParameters       = types.QueryParameters
	QueryInflation        = types.QueryInflation
	QueryAnnualProvisions = types.QueryAnnualProvisions
	RouterKey             = types.RouterKey
)

var (
	// functions aliases
	NewKeeper                  = keeper.NewKeeper
	NewQuerier                 = keeper.NewQuerier
	NewGenesisState            = types.NewGenesisState
	DefaultGenesisState        = types.DefaultGenesisState
	ValidateGenesis            = types.ValidateGenesis
	NewMinter                  = types.NewMinter
	InitialMinter              = types.InitialMinter
	DefaultInitialMinter       = types.DefaultInitialMinter
	ValidateMinter             = types.ValidateMinter
	ParamKeyTable              = types.ParamKeyTable
	NewParams                  = types.NewParams
	DefaultParams              = types.DefaultParams
	ErrProposerMustBeValidator = types.ErrProposerMustBeValidator

	// variable aliases
	ModuleCdc    = types.ModuleCdc
	MinterKey    = types.MinterKey
	KeyMintDenom = types.KeyMintDenom
	//KeyInflationRateChange = types.KeyInflationRateChange
	//KeyInflationMax        = types.KeyInflationMax
	//KeyInflationMin        = types.KeyInflationMin
	//KeyGoalBonded          = types.KeyGoalBonded
	KeyBlocksPerYear = types.KeyBlocksPerYear
	
	_ govtypes.Content = (*ExtraProposal)(nil)
)

type (
	Keeper        = keeper.Keeper
	GenesisState  = types.GenesisState
	Minter        = types.Minter
	Params        = types.Params
	ExtraProposal = types.ExtraProposal
)