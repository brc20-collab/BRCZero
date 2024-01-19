package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSrc20TickInfoInput(t *testing.T) {
	_, err := GetSrc20TickInfoInput("kevin")
	require.NoError(t, err)
}
