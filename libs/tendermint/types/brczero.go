package types

const FlagZeroReorgHeight = "reorg-height"

type BRCZeroRequestTx struct {
	HexRlpEncodeTx string `json:"hex_rlp_encode_tx"`
	BTCFee         uint64 `json:"btc_fee"`
}
