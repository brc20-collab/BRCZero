package node

import (
	"fmt"
	cmstore "github.com/brc20-collab/brczero/libs/cosmos-sdk/store"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/iavl"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/rootmulti"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/types"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	cfg "github.com/brc20-collab/brczero/libs/tendermint/config"
	"github.com/brc20-collab/brczero/libs/tendermint/store"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"
	"log"
	"time"
)

var pruneH int64 = 50

func handleReorg(blockStore *store.BlockStore, appDB dbm.DB, config *cfg.Config) {
	log.Printf("start height [%d,%d)...", blockStore.Base(), blockStore.Height())
	if blockStore.Height() < pruneH {
		return
	}
	pruneBlocksFromTop(blockStore, pruneH)
	log.Println("===========finish prune block!!!===========")
	pruneApp(pruneH, appDB, config)
}

func pruneBlocksFromTop(blockStore *store.BlockStore, retainHeight int64) {
	baseHeight := blockStore.Base()
	log.Printf("Prune blocks [%d,%d)...", baseHeight, retainHeight)
	if retainHeight <= baseHeight {
		return
	}

	baseHeightBefore, sizeBefore := getBlockInfo(blockStore)
	start := time.Now()
	_, err := blockStore.DeleteBlocksFromTop(retainHeight)
	if err != nil {
		panic(fmt.Errorf("failed to prune block store: %w", err))
	}

	baseHeightAfter, sizeAfter := getBlockInfo(blockStore)
	log.Printf("Block db info [baseHeight,size]: [%d,%d] --> [%d,%d]\n", baseHeightBefore, sizeBefore, baseHeightAfter, sizeAfter)
	log.Printf("Prune blocks done in %v \n", time.Since(start))
}

func getBlockInfo(blockStore *store.BlockStore) (baseHeight, size int64) {
	baseHeight = blockStore.Base()
	size = blockStore.Size()
	return
}

// pruneApp deletes app states between the given heights (including from, excluding to).
func pruneApp(from int64, appDB dbm.DB, config *cfg.Config) {
	log.Printf("Prune applcation [%d:)...", from)

	rs := initAppStore(appDB)
	versions := rs.GetVersions()
	if len(versions) == 0 {
		return
	}
	pruneHeights := rs.GetPruningHeights()

	newVersions := make([]int64, 0)
	newPruneHeights := make([]int64, 0)
	deleteVersions := make([]int64, 0)

	for _, v := range pruneHeights {
		if v < from {
			newPruneHeights = append(newPruneHeights, v)
			continue
		}
		deleteVersions = append(deleteVersions, v)
	}

	for _, v := range versions {
		if v < from {
			newVersions = append(newVersions, v)
			continue
		}
		deleteVersions = append(deleteVersions, v)
	}
	log.Printf("Prune application: Versions=%v, PruneVersions=%v", len(versions)+len(pruneHeights), len(deleteVersions))

	keysNumBefore, kvSizeBefore := calcKeysNum(appDB)
	start := time.Now()
	for key, store := range rs.GetStores() {
		if store.GetStoreType() == types.StoreTypeIAVL {
			// If the store is wrapped with an inter-block cache, we must first unwrap
			// it to get the underlying IAVL store.
			store = rs.GetCommitKVStore(key)

			if err := store.(*iavl.Store).DeleteVersions(deleteVersions...); err != nil {
				log.Printf("failed to delete version: %s", err)
			}
		}
	}

	rs.FlushPruneHeights(newPruneHeights, newVersions)

	keysNumAfter, kvSizeAfter := calcKeysNum(appDB)
	log.Printf("Application db key info [keysNum,kvSize]: [%d,%d] --> [%d,%d]\n", keysNumBefore, kvSizeBefore, keysNumAfter, kvSizeAfter)
	log.Printf("Prune application done in %v \n", time.Since(start))
}

func initAppStore(appDB dbm.DB) *rootmulti.Store {
	cms := cmstore.NewCommitMultiStore(appDB)

	keys := sdk.NewKVStoreKeys(
		"main", "mpt", "staking",
		"supply", "mint", "distribution", "slashing",
		"gov", "params", "upgrade", "evidence",
		"token", "lock",
	)
	tkeys := sdk.NewTransientStoreKeys("transient_params")

	for _, key := range keys {
		cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, nil)

	}
	for _, key := range tkeys {
		cms.MountStoreWithDB(key, sdk.StoreTypeTransient, nil)
	}

	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	rs, ok := cms.(*rootmulti.Store)
	if !ok {
		panic("cms of from app is not rootmulti store")
	}

	return rs
}

func calcKeysNum(db dbm.DB) (keys, kvSize uint64) {
	iter, err := db.Iterator(nil, nil)
	if err != nil {
		panic(err)
	}
	for ; iter.Valid(); iter.Next() {
		keys++
		kvSize += uint64(len(iter.Key())) + uint64(len(iter.Value()))
	}
	iter.Close()
	return
}
