package types

type ZeroRequestTx struct {
	HexRlpEncodeTx string `json:"hex_rlp_encode_tx"`
	BTCFee         string `json:"btc_fee"`
}

type ZeroResponseData struct {
	BTCHeight        string          `json:"block_height"`
	BTCBlockHash     string          `json:"block_hash"`
	BTCPrevBlockHash string          `json:"prev_block_hash"`
	IsConfirmed      bool            `json:"is_confirmed"`
	ZeroTxs          []ZeroRequestTx `json:"txs"`
}

type ZeroAPIResponse struct {
	Code int32            `json:"code"`
	Msg  string           `json:"msg"`
	Data ZeroResponseData `json:"data"`
}
