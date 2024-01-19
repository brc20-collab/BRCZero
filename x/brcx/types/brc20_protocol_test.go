package types

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpackGetBrc20TickInfoOutput(t *testing.T) {
	ret, err := hex.DecodeString("000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000069abc0ca57cf720e8d57cbf52032d895182bc1ee00000000000000000000000000000000000000000000000000000004e3b2920000000000000000000000000000000000000000000000000000000000000f4240000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000000046c66303600000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	_, err = UnpackGetBrc20TickInfoOutput(ret)
	require.NoError(t, err)
}