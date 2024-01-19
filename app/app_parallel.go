package app

import (
	"strings"

	appante "github.com/brc20-collab/brczero/app/ante"
	ethermint "github.com/brc20-collab/brczero/app/types"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth"
	authante "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/ante"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/bank"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/x/supply"
	abci "github.com/brc20-collab/brczero/libs/tendermint/abci/types"
	"github.com/brc20-collab/brczero/libs/tendermint/types"
	"github.com/brc20-collab/brczero/x/evm"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"
)

// feeCollectorHandler set or get the value of feeCollectorAcc
func updateFeeCollectorHandler(bk bank.Keeper, sk supply.Keeper) sdk.UpdateFeeCollectorAccHandler {
	return func(ctx sdk.Context, balance sdk.Coins) error {
		err := bk.SetCoins(ctx, sk.GetModuleAccount(ctx, auth.FeeCollectorName).GetAddress(), balance)
		if err != nil {
			return err
		}
		return nil
	}
}

// fixLogForParallelTxHandler fix log for parallel tx
func fixLogForParallelTxHandler(ek *evm.Keeper) sdk.LogFix {
	return func(tx []sdk.Tx, logIndex []int, hasEnterEvmTx []bool, anteErrs []error, resp []abci.ResponseDeliverTx) (logs [][]byte) {
		return ek.FixLog(tx, logIndex, hasEnterEvmTx, anteErrs, resp)
	}
}

func preDeliverTxHandler(ak auth.AccountKeeper) sdk.PreDeliverTxHandler {
	return func(ctx sdk.Context, tx sdk.Tx, onlyVerifySig bool) {
		if evmTx, ok := tx.(*evmtypes.MsgEthereumTx); ok {
			if evmTx.BaseTx.From == "" {
				_ = evmTxVerifySigHandler(ctx.ChainID(), ctx.BlockHeight(), evmTx)
			}
		}
	}
}

func evmTxVerifySigHandler(chainID string, blockHeight int64, evmTx *evmtypes.MsgEthereumTx) error {
	chainIDEpoch, err := ethermint.ParseChainID(chainID)
	if err != nil {
		return err
	}
	err = evmTx.VerifySig(chainIDEpoch, blockHeight)
	if err != nil {
		return err
	}
	return nil
}

func getTxFeeHandler() sdk.GetTxFeeHandler {
	return func(tx sdk.Tx) (fee sdk.Coins) {
		if feeTx, ok := tx.(authante.FeeTx); ok {
			fee = feeTx.GetFee()
		}

		return
	}
}

// getTxFeeAndFromHandler get tx fee and from
func getTxFeeAndFromHandler(ek appante.EVMKeeper) sdk.GetTxFeeAndFromHandler {
	return func(ctx sdk.Context, tx sdk.Tx) (fee sdk.Coins, isEvm bool, needUpdateTXCounter bool, from string, to string, err error, supportPara bool) {
		if evmTx, ok := tx.(*evmtypes.MsgEthereumTx); ok {
			isEvm = true
			supportPara = true
			if appante.IsE2CTx(ek, &ctx, evmTx) {
				needUpdateTXCounter = true
				// E2C will include cosmos Msg in the Payload.
				// Sometimes, this Msg do not support parallel execution.
				if !types.HigherThanMercury(ctx.BlockHeight()) {
					supportPara = false
				}
			}
			err = evmTxVerifySigHandler(ctx.ChainID(), ctx.BlockHeight(), evmTx)
			if err != nil {
				return
			}
			fee = evmTx.GetFee()
			from = evmTx.BaseTx.From
			if len(from) > 2 {
				from = strings.ToLower(from[2:])
			}
			if evmTx.To() != nil {
				to = strings.ToLower(evmTx.To().String()[2:])
			}
		} else if feeTx, ok := tx.(authante.FeeTx); ok {
			fee = feeTx.GetFee()
			if tx.GetType() == sdk.StdTxType {
				if types.HigherThanEarth(ctx.BlockHeight()) {
					needUpdateTXCounter = true
				}
				txMsgs := tx.GetMsgs()
				// only support one message
				if len(txMsgs) == 1 {
					if msg, ok := txMsgs[0].(interface{ CalFromAndToForPara() (string, string) }); ok {
						from, to = msg.CalFromAndToForPara()
						if types.HigherThanMercury(ctx.BlockHeight()) {
							supportPara = true
						}
					}
				}
			}
		}

		return
	}
}
