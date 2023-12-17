package types

type ZeroRequestTx struct {
	TxInfo string `json:"hex_rlp_encode_tx"`
	BTCFee uint64 `json:"btc_fee"`
}
