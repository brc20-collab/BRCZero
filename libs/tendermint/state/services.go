package state

import (
	"github.com/brc20-collab/brczero/libs/tendermint/types"
)

//------------------------------------------------------
// blockchain services types
// NOTE: Interfaces used by RPC must be thread safe!
//------------------------------------------------------

//------------------------------------------------------
// blockstore

// BlockStore defines the interface used by the ConsensusState.
type BlockStore interface {
	Base() int64
	Height() int64
	Size() int64

	LoadBlockMeta(height int64) *types.BlockMeta
	LoadBlock(height int64) *types.Block
	LoadBTCMeta(height int64) (*types.BTCBlockMeta, error)

	SaveBlock(block *types.Block, blockParts *types.PartSet, seenCommit *types.Commit)

	PruneBlocks(height int64) (uint64, error)
	DeleteBlocksFromTop(height int64) (uint64, error)

	LoadBlockByHash(hash []byte) *types.Block
	LoadZeroHeightByBtcHash(hash string) (int64, error)
	LoadBtcBlockHashByBtcTxid(btcTxid string) (string, error)
	LoadBtcBlockHashByBtcHeight(btcHeight int64) (string, error)
	LoadMapTxhashTxidByBtcHash(btcHash, protocolName string) (map[string]string, error)
	LoadBlockPart(height int64, index int) *types.Part

	LoadBlockCommit(height int64) *types.Commit
	LoadSeenCommit(height int64) *types.Commit
}

//-----------------------------------------------------------------------------
// evidence pool

// EvidencePool defines the EvidencePool interface used by the ConsensusState.
// Get/Set/Commit
type EvidencePool interface {
	PendingEvidence(int64) []types.Evidence
	AddEvidence(types.Evidence) error
	Update(*types.Block, State)
	// IsCommitted indicates if this evidence was already marked committed in another block.
	IsCommitted(types.Evidence) bool
}

// MockEvidencePool is an empty implementation of EvidencePool, useful for testing.
type MockEvidencePool struct{}

func (m MockEvidencePool) PendingEvidence(int64) []types.Evidence { return nil }
func (m MockEvidencePool) AddEvidence(types.Evidence) error       { return nil }
func (m MockEvidencePool) Update(*types.Block, State)             {}
func (m MockEvidencePool) IsCommitted(types.Evidence) bool        { return false }
