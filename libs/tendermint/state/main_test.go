package state_test

import (
	"os"
	"testing"

	"github.com/brc20-collab/brczero/libs/tendermint/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}
