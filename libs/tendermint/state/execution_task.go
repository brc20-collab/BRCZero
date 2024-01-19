package state

import (
	"time"

	"github.com/brc20-collab/brczero/libs/tendermint/libs/log"
	"github.com/brc20-collab/brczero/libs/tendermint/proxy"
	"github.com/brc20-collab/brczero/libs/tendermint/types"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"
)

type executionResult struct {
	res      *ABCIResponses
	duration time.Duration
	err      error
}

type executionTask struct {
	height         int64
	index          int64
	block          *types.Block
	stopped        bool
	taskResultChan chan *executionTask
	result         *executionResult
	proxyApp       proxy.AppConnConsensus
	db             dbm.DB
	logger         log.Logger
	blockHash      string
}
