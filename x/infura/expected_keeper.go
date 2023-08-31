package infura

import evm "github.com/brc20-collab/brczero/x/evm/watcher"

type EvmKeeper interface {
	SetObserverKeeper(keeper evm.InfuraKeeper)
}
