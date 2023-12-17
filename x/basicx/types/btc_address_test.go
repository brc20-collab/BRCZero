package types

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertBTCAddress(t *testing.T) {
	addrstr := "bcrt1pql65pjwt4ckvwwujmwuvdvd6rxa4ef64ras59fg46x38mygpu66svj2l9y"
	_, err := btcutil.DecodeAddress(addrstr, &chaincfg.RegressionNetParams)
	require.NoError(t, err)
}
