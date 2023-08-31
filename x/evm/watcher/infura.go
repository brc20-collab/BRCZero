package watcher

import "github.com/brc20-collab/brczero/x/evm/types"

type InfuraKeeper interface {
	OnSaveTransactionReceipt(*TransactionReceipt)
	OnSaveBlock(types.Block)
	OnSaveTransaction(Transaction)
	OnSaveContractCode(address string, code []byte)
}
