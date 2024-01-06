package keeper

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	ethermint "github.com/brc20-collab/brczero/app/types"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	tmtypes "github.com/brc20-collab/brczero/libs/tendermint/types"
	"github.com/brc20-collab/brczero/x/brcx/types"
	evmtypes "github.com/brc20-collab/brczero/x/evm/types"
)

// CallEvm execute an evm message from native module
func (k Keeper) CallEvm(ctx sdk.Context, callerAddr common.Address, to *common.Address, value *big.Int, data []byte, info *types.ResultInfo) (*evmtypes.ExecutionResult, *evmtypes.ResultData, error) {

	config, found := k.evmKeeper.GetChainConfig(ctx)
	if !found {
		return nil, nil, types.ErrChainConfigNotFound
	}

	chainIDEpoch, err := ethermint.ParseChainID(ctx.ChainID())
	if err != nil {
		return nil, nil, err
	}

	acc := k.accountKeeper.GetAccount(ctx, callerAddr.Bytes())
	if acc == nil {
		acc = k.accountKeeper.NewAccountWithAddress(ctx, callerAddr.Bytes())
	}
	nonce := acc.GetSequence()
	info.Nonce = nonce
	txHash := tmtypes.Tx(ctx.TxBytes()).Hash()
	ethTxHash := common.BytesToHash(txHash)

	gasLimit := ctx.GasMeter().Limit()
	if gasLimit == sdk.NewInfiniteGasMeter().Limit() {
		gasLimit = k.evmKeeper.GetParams(ctx).MaxGasLimitPerTx
	}
	trace := ctx.IsTraceTx()
	st := evmtypes.StateTransition{
		AccountNonce: nonce,
		Price:        big.NewInt(0),
		GasLimit:     gasLimit,
		Recipient:    to,
		Amount:       value,
		Payload:      data,
		Csdb:         evmtypes.CreateEmptyCommitStateDB(k.evmKeeper.GenerateCSDBParams(), ctx),
		ChainID:      chainIDEpoch,
		TxHash:       &ethTxHash,
		Sender:       callerAddr,
		Simulate:     ctx.IsCheckTx(),
		TraceTx:      trace,
		TraceTxLog:   trace,
	}
	st.Csdb.Prepare(ethTxHash, k.evmKeeper.GetBlockHash(), 0)
	ctx.SetIsTraceTxLog(trace)
	ctx.SetIsTraceTx(trace)
	st.SetCallToCM(k.evmKeeper.GetCallToCM())
	executionResult, resultData, err, innertxs, contracts := st.TransitionDb(ctx, config)
	if !ctx.IsCheckTx() && !ctx.IsTraceTx() {
		if innertxs != nil {
			k.evmKeeper.AddInnerTx(ethTxHash.Hex(), innertxs)
		}
		if contracts != nil {
			k.evmKeeper.AddContract(contracts)
		}
	}
	attributes := make([]sdk.Attribute, 0)
	if err != nil {
		attribute := sdk.NewAttribute(types.AttributeResult, err.Error())
		attributes = append(attributes, attribute)
	} else {
		buff, err := json.Marshal(resultData)
		if err != nil {
			attribute := sdk.NewAttribute(types.AttributeResult, err.Error())
			attributes = append(attributes, attribute)
		} else {
			attribute := sdk.NewAttribute(types.AttributeResult, string(buff))
			attributes = append(attributes, attribute)
		}

	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCallEvm,
			attributes...,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	st.Csdb.Commit(false) // write code to db

	temp := k.accountKeeper.GetAccount(ctx, callerAddr.Bytes())
	if temp == nil {
		if err := acc.SetCoins(sdk.Coins{}); err != nil {
			return nil, nil, err
		}
		temp = acc
	}
	if err := temp.SetSequence(nonce + 1); err != nil {
		return nil, nil, err
	}
	k.accountKeeper.SetAccount(ctx, temp)

	return executionResult, resultData, err
}
