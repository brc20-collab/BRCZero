package mpt

import (
	"fmt"
	"path/filepath"

	ethcmn "github.com/ethereum/go-ethereum/common"
	ethstate "github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/brc20-collab/brczero/app"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/server"
	iavlstore "github.com/brc20-collab/brczero/libs/cosmos-sdk/store/iavl"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/mpt"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/rootmulti"
	sdk "github.com/brc20-collab/brczero/libs/cosmos-sdk/types"
	authtypes "github.com/brc20-collab/brczero/libs/cosmos-sdk/x/auth/types"
	"github.com/brc20-collab/brczero/libs/iavl"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"
	tmdb "github.com/brc20-collab/brczero/libs/tm-db"
)

const (
	applicationDB = "application"

	accStoreKey    = authtypes.StoreKey
	evmStoreKey    = "evm"
	legacyStoreKey = "evmlegacy"

	iavlAccKey       = "s/k:acc/"
	iavlEvmKey       = "s/k:evm/"
	iavlEvmLegacyKey = "s/k:evmlegacy/"
	KeyParams        = "s/k:params/"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

// checkValidKey checks if the key is equal to authtypes.StoreKey or evmtypes.StoreKey
func checkValidKey(key string) error {
	if key != accStoreKey && key != evmStoreKey && key != legacyStoreKey {
		return fmt.Errorf("invalid key %s", key)
	}
	return nil
}

/*
 * Common functions about cosmos-sdk
 */
// newMigrationApp generates a new app with the given key and application.db
func newMigrationApp(ctx *server.Context) *app.BRCZeroApp {
	appDb := openApplicationDb(ctx.Config.RootDir)
	return app.NewBRCZeroApp(
		ctx.Logger,
		appDb,
		nil,
		true,
		0,
	)
}

func openApplicationDb(rootdir string) tmdb.DB {
	dataDir := filepath.Join(rootdir, "data")
	appDb, err := sdk.NewDB(applicationDB, dataDir)
	if err != nil {
		panic("fail to open application db: " + err.Error())
	}
	return appDb
}

/*
 * Common functions about mpt
 */
// getStorageTrie returns the trie of the given address and stateRoot
func getStorageTrie(db ethstate.Database, addrHash, stateRoot ethcmn.Hash) ethstate.Trie {
	tr, err := db.OpenStorageTrie(addrHash, stateRoot)
	panicError(err)
	return tr
}

// pushData2Database commit the data to the database
func pushData2Database(db ethstate.Database, tree ethstate.Trie, height int64, isEvm bool, nodes *trie.MergedNodeSet) {

	root, set, err := tree.Commit(true)
	panicError(err)

	err = nodes.Merge(set)
	panicError(err)

	err = db.TrieDB().UpdateForOK(nodes, mpt.AccountStateRootRetriever.RetrieveStateRoot)
	panicError(err)
	err = db.TrieDB().Commit(root, false, nil)
	panicError(err)

	setMptRootHash(db, uint64(height), root, isEvm)
}

// setMptRootHash sets the mapping from block height to root mpt hash
func setMptRootHash(db ethstate.Database, height uint64, hash ethcmn.Hash, isEvm bool) {
	heightBytes := sdk.Uint64ToBigEndian(height)
	db.TrieDB().DiskDB().Put(mpt.KeyPrefixAccLatestStoredHeight, heightBytes)
	db.TrieDB().DiskDB().Put(append(mpt.KeyPrefixAccRootMptHash, heightBytes...), hash.Bytes())
}

func writeDataToRawdb(batch ethdb.Batch) {
	if err := batch.Write(); err != nil {
		panic(err)
	}
	batch.Reset()
}

func getUpgradedTree(db dbm.DB, prefix []byte, usePreLatest bool) *iavl.MutableTree {
	rs := rootmulti.NewStore(db)
	latestVersion := rs.GetLatestVersion()
	if latestVersion == 0 {
		return nil
	}

	db = dbm.NewPrefixDB(db, prefix)

	tree, err := iavl.NewMutableTree(db, iavlstore.IavlCacheSize)
	if err != nil {
		panic("Fail to get tree: " + err.Error())
	}

	if usePreLatest {
		latestVersion -= 1
	}

	if latestVersion <= 0 {
		panic(fmt.Sprintf("invalid version to load: %d", latestVersion))
	}

	_, err = tree.LoadVersion(latestVersion)
	if err != nil {
		panic("fail to load target version tree: " + err.Error())
	}

	return tree
}
