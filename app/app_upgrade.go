package app

import (
	cliContext "github.com/brc20-collab/brczero/libs/cosmos-sdk/client/context"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/client/utils"
)

func (app *BRCZeroApp) RegisterTxService(clientCtx cliContext.CLIContext) {
	utils.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.grpcSimulate, clientCtx.InterfaceRegistry)
}
func (app *BRCZeroApp) grpcSimulate(txBytes []byte) (sdk.GasInfo, *sdk.Result, error) {
	tx, err := app.GetTxDecoder()(txBytes)
	if err != nil {
		return sdk.GasInfo{}, nil, sdkerrors.Wrap(err, "failed to decode tx")
	}
	return app.Simulate(txBytes, tx, 0, nil)
}
