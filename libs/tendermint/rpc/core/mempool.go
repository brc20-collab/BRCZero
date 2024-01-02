package core

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/baseapp"
	ctypes "github.com/brc20-collab/brczero/libs/tendermint/rpc/core/types"
	rpctypes "github.com/brc20-collab/brczero/libs/tendermint/rpc/jsonrpc/types"
	"github.com/brc20-collab/brczero/libs/tendermint/types"
)

//-----------------------------------------------------------------------------
// NOTE: tx should be signed, but this is only checked at the app level (not by Tendermint!)

// BroadcastTxAsync returns right away, with no response. Does not wait for
// CheckTx nor DeliverTx results.
// More: https://docs.tendermint.com/master/rpc/#/Tx/broadcast_tx_async
func BroadcastTxAsync(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	return nil, fmt.Errorf("BroadcastTxAsync is not provided yet")
}

// BroadcastTxSync returns with the response from CheckTx. Does not wait for
// DeliverTx result.
// More: https://docs.tendermint.com/master/rpc/#/Tx/broadcast_tx_sync
func BroadcastTxSync(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	return nil, fmt.Errorf("BroadcastTxSync is not provided yet")
}

// BroadcastTxCommit returns with the responses from CheckTx and DeliverTx.
// More: https://docs.tendermint.com/master/rpc/#/Tx/broadcast_tx_commit
func BroadcastTxCommit(ctx *rpctypes.Context, tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	return nil, fmt.Errorf("BroadcastTxCommit is not provided yet")
}

func BroadcastBrczeroTxsAsync(ctx *rpctypes.Context, btcHeight int64, btcBlockHash string, isConfirmed bool, brczeroTxs []types.ZeroRequestTx) ([]*ctypes.ResultBroadcastTx, error) {
	//txs := make([]types.Tx, 0)
	res := make([]*ctypes.ResultBroadcastTx, 0)
	//for _, s := range brczeroTxs {
	//	tx, err := rlp.EncodeToBytes(s)
	//	if err != nil {
	//		return nil, err
	//	}
	//	txs = append(txs, tx)
	//	res = append(res, &ctypes.ResultBroadcastTx{Hash: types.Tx(tx).Hash()})
	//}
	//
	//err := env.Mempool.AddBrczeroData(btcHeight, btcBlockHash, isConfirmed, txs)
	//if err != nil {
	//	return nil, err
	//}
	return res, nil
}

// UnconfirmedTxs gets unconfirmed transactions (maximum ?limit entries)
// including their number.
// More: https://docs.tendermint.com/master/rpc/#/Info/unconfirmed_txs
func UnconfirmedTxs(ctx *rpctypes.Context, limit int) (*ctypes.ResultUnconfirmedTxs, error) {

	txs := env.Mempool.ReapMaxTxs(limit)
	return &ctypes.ResultUnconfirmedTxs{
		Count:      len(txs),
		Total:      env.Mempool.Size(),
		TotalBytes: env.Mempool.TxsBytes(),
		Txs:        txs}, nil
}

// NumUnconfirmedTxs gets number of unconfirmed transactions.
// More: https://docs.tendermint.com/master/rpc/#/Info/num_unconfirmed_txs
func NumUnconfirmedTxs(ctx *rpctypes.Context) (*ctypes.ResultUnconfirmedTxs, error) {
	return &ctypes.ResultUnconfirmedTxs{
		Count:      env.Mempool.Size(),
		Total:      env.Mempool.Size(),
		TotalBytes: env.Mempool.TxsBytes()}, nil
}

func TxSimulateGasCost(ctx *rpctypes.Context, hash string) (*ctypes.ResponseTxSimulateGas, error) {
	return &ctypes.ResponseTxSimulateGas{
		GasCost: env.Mempool.GetTxSimulateGas(hash),
	}, nil
}

func UserUnconfirmedTxs(address string, limit int) (*ctypes.ResultUserUnconfirmedTxs, error) {
	txs := env.Mempool.ReapUserTxs(address, limit)
	return &ctypes.ResultUserUnconfirmedTxs{
		Count: len(txs),
		Txs:   txs}, nil
}

func TmUserUnconfirmedTxs(ctx *rpctypes.Context, address string, limit int) (*ctypes.ResultUserUnconfirmedTxs, error) {
	return UserUnconfirmedTxs(address, limit)
}

func UserNumUnconfirmedTxs(address string) (*ctypes.ResultUserUnconfirmedTxs, error) {
	nums := env.Mempool.ReapUserTxsCnt(address)
	return &ctypes.ResultUserUnconfirmedTxs{
		Count: nums}, nil
}

func TmUserNumUnconfirmedTxs(ctx *rpctypes.Context, address string) (*ctypes.ResultUserUnconfirmedTxs, error) {
	return UserNumUnconfirmedTxs(address)
}

func GetUnconfirmedTxByHash(hash [sha256.Size]byte) (types.Tx, error) {
	return env.Mempool.GetTxByHash(hash)
}

func GetAddressList() (*ctypes.ResultUnconfirmedAddresses, error) {
	addressList := env.Mempool.GetAddressList()
	return &ctypes.ResultUnconfirmedAddresses{
		Addresses: addressList,
	}, nil
}

func TmGetAddressList(ctx *rpctypes.Context) (*ctypes.ResultUnconfirmedAddresses, error) {
	return GetAddressList()
}

func GetPendingNonce(address string) (*ctypes.ResultPendingNonce, bool) {
	nonce, ok := env.Mempool.GetPendingNonce(address)
	if !ok {
		return nil, false
	}
	return &ctypes.ResultPendingNonce{
		Nonce: nonce,
	}, true
}

func GetEnableDeleteMinGPTx(ctx *rpctypes.Context) (*ctypes.ResultEnableDeleteMinGPTx, error) {
	status := env.Mempool.GetEnableDeleteMinGPTx()
	return &ctypes.ResultEnableDeleteMinGPTx{Enable: status}, nil
}

func GetPendingTxs(ctx *rpctypes.Context) (*ctypes.ResultPendingTxs, error) {
	pendingTx := make(map[string]map[string]types.WrappedMempoolTx)
	if baseapp.IsMempoolEnablePendingPool() {
		pendingTx = env.Mempool.GetPendingPoolTxsBytes()
	}
	return &ctypes.ResultPendingTxs{Txs: pendingTx}, nil
}

func GetCurrentZeroData(ctx *rpctypes.Context) (*ctypes.ResultZeroData, error) {
	data := env.Mempool.GetCurrentZeroData()
	res := make(map[string]types.ZeroData, 0)
	for h, d := range data {
		res[strconv.FormatInt(h, 10)] = d
	}
	return &ctypes.ResultZeroData{Data: res}, nil
}
