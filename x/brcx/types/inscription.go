package types

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type InscriptionContext struct {
	InscriptionId     string `json:"inscription_id" yaml:"inscription_id"`
	InscriptionNumber int64  `json:"inscription_number" yaml:"inscription_number"`
	IsTransfer        bool   `json:"is_transfer" yaml:"is_transfer"`
	Txid              string `json:"txid" yaml:"txid"`
	Sender            string `json:"sender" yaml:"sender"`
	Receiver          string `json:"receiver" yaml:"receiver"`
	CommitInput       string `json:"commit_input" yaml:"commit_input"`
	RevealOutput      string `json:"reveal_output" yaml:"reveal_output"`
	OldSatPoint       string `json:"old_sat_point" yaml:"old_sat_point"`
	NewSatPoint       string `json:"new_sat_point" yaml:"new_sat_point"`

	BlockHash   string `json:"block_hash" yaml:"block_hash"`     // btc block hash
	BlockTime   uint32 `json:"block_time" yaml:"block_time"`     // btc block time
	BlockHeight uint64 `json:"block_height" yaml:"block_height"` // btc block height
}

// Sat
type SatPoint struct {
	Outpoint Outpoint `json:"outpoint" yaml:"outpoint"`
	Offset   uint64   `json:"offset" yaml:"offset"`
}

// UTXO
type Outpoint struct {
	Txid Hash   `json:"txid" yaml:"txid"`
	Vout uint32 `json:"vout" yaml:"vout"`
}

// 交易输出
type TxOut struct {
	Txid     Hash   `json:"txid" yaml:"txid"`
	Vout     uint32 `json:"vout" yaml:"vout"`
	Value    uint64 `json:"value" yaml:"value"`
	PkScript []byte `json:"pk_script" json:"pkScript"`
}

type TxIn struct {
	PrevTxOut   TxOut  `json:"prev_tx_out" yaml:"prevTxOut"`
	CurrentTxid Hash   `json:"current_txid" yaml:"currentTxid"`
	IndexIn     uint32 `json:"index_in" yaml:"indexIn"`
}

type Hash []byte

func (h Hash) ValidateBasic() error {
	if len(h) != common.HashLength {
		return fmt.Errorf("%s error length 32", hex.EncodeToString(h))
	}
	return nil
}
