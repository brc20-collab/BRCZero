package store

import (
	dbm "github.com/brc20-collab/brczero/libs/tm-db"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/cache"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/rootmulti"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
