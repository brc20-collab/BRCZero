package ante

import (
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	sdkerrors "github.com/brc20-collab/brczero/libs/cosmos-sdk/types/errors"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/types/innertx"
	types2 "github.com/brc20-collab/brczero/libs/tendermint/types"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"
)

// EVMKeeper defines the expected keeper interface used on the Eth AnteHandler
type EVMKeeper interface {
	innertx.InnerTxKeeper
	GetParams(ctx sdk.Context) evmtypes.Params
	IsAddressBlocked(ctx sdk.Context, addr sdk.AccAddress) bool
	IsMatchSysContractAddress(ctx sdk.Context, addr sdk.AccAddress) bool
}

// NewWasmGasLimitDecorator creates a new WasmGasLimitDecorator.
func NewWasmGasLimitDecorator(evm EVMKeeper) WasmGasLimitDecorator {
	return WasmGasLimitDecorator{
		GasLimitDecorator: NewGasLimitDecorator(evm),
	}
}

type WasmGasLimitDecorator struct {
	GasLimitDecorator
}

func (g WasmGasLimitDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// do another ante check for simulation
	if !types2.HigherThanEarth(ctx.BlockHeight()) {
		return next(ctx, tx, simulate)
	}
	return g.GasLimitDecorator.AnteHandle(ctx, tx, simulate, next)
}

// NewGasLimitDecorator creates a new GasLimitDecorator.
func NewGasLimitDecorator(evm EVMKeeper) GasLimitDecorator {
	return GasLimitDecorator{
		evm: evm,
	}
}

type GasLimitDecorator struct {
	evm EVMKeeper
}

func (g GasLimitDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	pinAnte(ctx.AnteTracer(), "GasLimitDecorator")

	currentGasMeter := ctx.GasMeter() // avoid race
	infGasMeter := sdk.GetReusableInfiniteGasMeter()
	ctx.SetGasMeter(infGasMeter)
	if tx.GetGas() > g.evm.GetParams(ctx).MaxGasLimitPerTx {
		ctx.SetGasMeter(currentGasMeter)
		sdk.ReturnInfiniteGasMeter(infGasMeter)
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrTxTooLarge, "too large gas limit, it must be less than %d", g.evm.GetParams(ctx).MaxGasLimitPerTx)
	}

	ctx.SetGasMeter(currentGasMeter)
	sdk.ReturnInfiniteGasMeter(infGasMeter)
	return next(ctx, tx, simulate)
}