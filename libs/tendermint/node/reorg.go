package node

import (
	"fmt"
	"github.com/brc20-collab/brczero/libs/iavl"
	cfg "github.com/brc20-collab/brczero/libs/tendermint/config"
	sm "github.com/brc20-collab/brczero/libs/tendermint/state"
	"github.com/brc20-collab/brczero/libs/tendermint/store"
	dbm "github.com/brc20-collab/brczero/libs/tm-db"
	"github.com/spf13/viper"
	"log"
	"time"
)

func handleReorgBlock(blockStore *store.BlockStore, stateDB dbm.DB, state sm.State, config *cfg.Config) sm.State {
	pruneH := viper.GetInt64("start-height")
	log.Printf("start height [%d,%d)...", blockStore.Base(), blockStore.Height())
	if pruneH == 0 || blockStore.Height() < pruneH {
		return state
	}
	sm.SetIgnoreSmbCheck(true)
	iavl.SetIgnoreVersionCheck(true)
	pruneBlocksFromTop(blockStore, pruneH)
	log.Println("===========finish prune block!!!===========")

	return constructStartState(state, stateDB, pruneH-1)
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

func constructStartState(state sm.State, stateStoreDB dbm.DB, startHeight int64) sm.State {
	stateCopy := state.Copy()
	validators, lastStoredHeight, err := sm.LoadValidatorsWithStoredHeight(stateStoreDB, startHeight+1)
	lastValidators, err := sm.LoadValidators(stateStoreDB, startHeight)
	if err != nil {
		return stateCopy
	}
	nextValidators, err := sm.LoadValidators(stateStoreDB, startHeight+2)
	if err != nil {
		return stateCopy
	}
	consensusParams, err := sm.LoadConsensusParams(stateStoreDB, startHeight+1)
	if err != nil {
		return stateCopy
	}
	stateCopy.Validators = validators
	stateCopy.LastValidators = lastValidators
	stateCopy.NextValidators = nextValidators
	stateCopy.ConsensusParams = consensusParams
	stateCopy.LastBlockHeight = startHeight
	stateCopy.LastHeightValidatorsChanged = lastStoredHeight
	return stateCopy
}
