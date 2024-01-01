package types

import ethtypes "github.com/ethereum/go-ethereum/core/types"

type EthLogWithTxid struct {
	Logs []*ethtypes.Log `json:"logs"`
	Txid string          `json:"txid"`
}
