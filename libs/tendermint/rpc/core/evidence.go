package core

import (
	"github.com/brc20-collab/brczero/libs/tendermint/evidence"
	ctypes "github.com/brc20-collab/brczero/libs/tendermint/rpc/core/types"
	rpctypes "github.com/brc20-collab/brczero/libs/tendermint/rpc/jsonrpc/types"
	"github.com/brc20-collab/brczero/libs/tendermint/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://docs.tendermint.com/master/rpc/#/Info/broadcast_evidence
func BroadcastEvidence(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	err := env.EvidencePool.AddEvidence(ev)
	if _, ok := err.(evidence.ErrEvidenceAlreadyStored); err == nil || ok {
		return &ctypes.ResultBroadcastEvidence{Hash: ev.Hash()}, nil
	}
	return nil, err
}
