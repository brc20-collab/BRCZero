package mpt

import (
	"encoding/hex"
	"fmt"
	"io"
	"sync"

	tmtypes "github.com/brc20-collab/brczero/libs/tendermint/types"

	mpttype "github.com/brc20-collab/brczero/libs/cosmos-sdk/store/mpt/types"

	ethcmn "github.com/ethereum/go-ethereum/common"
	ethstate "github.com/ethereum/go-ethereum/core/state"

	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/cachekv"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/tracekv"
	"github.com/brc20-collab/brczero/libs/cosmos-sdk/store/types"
)

type ImmutableMptStore struct {
	trie ethstate.Trie
	db   ethstate.Database
	root ethcmn.Hash
	mtx  sync.Mutex

	brcRpcStateCache map[string][]byte
}

func NewImmutableMptStore(db ethstate.Database, root ethcmn.Hash, brcRpcStateCache map[string][]byte) (*ImmutableMptStore, error) {
	ms := &ImmutableMptStore{
		db:               db,
		root:             root,
		brcRpcStateCache: brcRpcStateCache,
	}
	trie, err := ms.db.OpenTrie(root)
	if err != nil {
		return nil, err
	}
	ms.trie = trie
	return ms, nil
}

func NewImmutableMptStoreFromTrie(db ethstate.Database, trie ethstate.Trie) *ImmutableMptStore {
	ms := &ImmutableMptStore{
		trie: db.CopyTrie(trie),
		db:   db,
		root: trie.Hash(),
	}
	return ms
}

func (ms *ImmutableMptStore) GetBrcRpcState(key []byte) []byte {
	if value, ok := ms.brcRpcStateCache[hex.EncodeToString(key)]; ok {
		return value
	}
	return nil
}

func (ms *ImmutableMptStore) Get(key []byte) []byte {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	if tmtypes.RpcFlag != tmtypes.RpcApplyBlockMode {
		if value := ms.GetBrcRpcState(key); value != nil {
			return value
		}
	}
	switch mptKeyType(key) {
	case storageType:
		_, stateRoot, realKey := decodeAddressStorageInfo(key)

		t, err := ms.db.OpenStorageTrie(ethcmn.Hash{}, stateRoot)
		if err != nil {
			return nil
		}

		value, err := t.TryGet(realKey)
		if err != nil {
			return nil
		}
		return value
	case addressType:
		value, err := ms.trie.TryGet(key)
		if err != nil {
			return nil
		}
		return value
	case putType:
		value, err := ms.db.TrieDB().DiskDB().Get(key[1:])
		if err != nil {
			return nil
		}
		return value
	default:
		panic(fmt.Errorf("not support key %s for immutable mpt get", hex.EncodeToString(key)))
	}
}

func (ms *ImmutableMptStore) Has(key []byte) bool {
	return ms.Get(key) != nil
}

func (ms *ImmutableMptStore) Set(key []byte, value []byte) {
	panic("immutable store cannot set")
}

func (ms *ImmutableMptStore) Delete(key []byte) {
	panic("immutable store cannot delete")
}

func (ms *ImmutableMptStore) CleanBrcRpcState() {
	ms.brcRpcStateCache = make(map[string][]byte, 0)
}

func (ms *ImmutableMptStore) getStorageTrie(addr ethcmn.Address, stateRoot ethcmn.Hash) ethstate.Trie {
	addrHash := mpttype.Keccak256HashWithSyncPool(addr[:])
	var t ethstate.Trie
	var err error
	t, err = ms.db.OpenStorageTrie(addrHash, stateRoot)
	if err != nil {
		t, err = ms.db.OpenStorageTrie(addrHash, ethcmn.Hash{})
		if err != nil {
			panic("unexpected err")
		}
	}

	return t
}

func (ms *ImmutableMptStore) Iterator(start, end []byte) types.Iterator {
	if IsStorageKey(start) {
		addr, stateRoot, _ := decodeAddressStorageInfo(start)
		t := ms.getStorageTrie(addr, stateRoot)

		return newMptIterator(t, start, end, true)
	}
	return newMptIterator(ms.db.CopyTrie(ms.trie), start, end, true)
}

func (ms *ImmutableMptStore) ReverseIterator(start, end []byte) types.Iterator {
	if IsStorageKey(start) {
		addr, stateRoot, _ := decodeAddressStorageInfo(start)
		t := ms.getStorageTrie(addr, stateRoot)

		return newMptIterator(t, start, end, false)
	}
	return newMptIterator(ms.db.CopyTrie(ms.trie), start, end, false)
}

func (ms *ImmutableMptStore) GetStoreType() types.StoreType {
	return StoreTypeMPT
}
func (ms *ImmutableMptStore) GetStoreName() string {
	return "ImmutableMptStore"
}

func (ms *ImmutableMptStore) CacheWrap() types.CacheWrap {
	//TODO implement me
	return cachekv.NewStore(ms)
}

func (ms *ImmutableMptStore) CacheWrapWithTrace(w io.Writer, tc types.TraceContext) types.CacheWrap {
	//TODO implement me
	return cachekv.NewStore(tracekv.NewStore(ms, w, tc))
}

var _ types.KVStore = (*ImmutableMptStore)(nil)
