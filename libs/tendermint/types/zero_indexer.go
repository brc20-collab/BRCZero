package types

type ZeroRequestTx struct {
	ProtocolName       string `json:"protocol_name"`
	Inscription        string `json:"inscription"`
	InscriptionContext string `json:"inscription_context"`
	BTCTxid            string `json:"btc_txid"`
	BTCFee             string `json:"btc_fee"`
}

type ZeroResponseData struct {
	BTCHeight        uint64          `json:"block_height"`
	BTCBlockHash     string          `json:"block_hash"`
	BTCPrevBlockHash string          `json:"prev_block_hash"`
	BTCBlockTime     uint32          `json:"block_time"`
	ZeroTxs          []ZeroRequestTx `json:"txs"`
}

type ZeroAPIResponse struct {
	Code int32            `json:"code"`
	Msg  string           `json:"msg"`
	Data ZeroResponseData `json:"data"`
}
