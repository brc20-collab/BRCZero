package types

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/common"
)

func ConvertBTCAddress(str string) (common.Address, error) {
	from := make([]byte, 0)
	addr, err := btcutil.DecodeAddress(str, &chaincfg.RegressionNetParams)
	if err != nil {
		return common.Address{}, fmt.Errorf("convert BTC address is error:%v", err)
	} else {
		if addr == nil {
			return common.Address{}, fmt.Errorf("the address of converted is empty")
		}
		switch addr := addr.(type) {
		case *btcutil.AddressPubKeyHash:
			from = addr.ScriptAddress()
		case *btcutil.AddressPubKey:
			from = btcutil.Hash160(addr.ScriptAddress())
		case *btcutil.AddressWitnessPubKeyHash:
			from = addr.ScriptAddress()
		default:
			fmt.Printf("it's dangerous: %s\n", fmt.Errorf("%s is not support type of address, only support p2pkh p2pk p2wpk", addr.String()))
			from = btcutil.Hash160(addr.ScriptAddress())
		}
	}

	return common.BytesToAddress(from), nil
}
