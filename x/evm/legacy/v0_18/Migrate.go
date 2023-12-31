package v018

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/x/evm/legacy/v0_16"
)

const (
	ModuleName = "evm"
)

// Migrate adds contract
func Migrate(oldGenState v0_16.GenesisState) GenesisState {
	params := Params{
		EnableCreate:                      false,
		EnableCall:                        false,
		ExtraEIPs:                         oldGenState.Params.ExtraEIPs,
		EnableContractDeploymentWhitelist: true,
		EnableContractBlockedList:         true,
		MaxGasLimitPerTx:                  30000000,
	}

	return GenesisState{
		Accounts:                    oldGenState.Accounts,
		TxsLogs:                     oldGenState.TxsLogs,
		ContractDeploymentWhitelist: []sdk.AccAddress{},
		ContractBlockedList:         []sdk.AccAddress{},
		ChainConfig:                 oldGenState.ChainConfig,
		Params:                      params,
	}
}
