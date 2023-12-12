package types

type BRCZeroRequestTx struct {
	HexRlpEncodeTx string `json:"hex_rlp_encode_tx"`
	BTCFee         uint64 `json:"btc_fee"`
}

type BRCZeroRequestData struct {
	BTCHeight    int64  `json:"btc_height"`
	BTCBlockHash string `json:"btc_block_hash"`
	BRCTxs       []BRCZeroRequestTx
}
