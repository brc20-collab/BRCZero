package types

const OK_CODE = 0

type ZeroRequestTx struct {
	ProtocolName       string `json:"protocol_name"`
	Inscription        string `json:"inscription"`
	InscriptionContext string `json:"inscription_context"`
	BTCTxid            string `json:"btc_txid"`
	BTCFee             string `json:"btc_fee"`
}

type ZeroAPIResponse[T interface{}] struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type ZeroResponseData struct {
	Page             uint            `json:"page"`
	Count            uint            `json:"count"`
	Sum              uint            `json:"sum"`
	BTCHeight        uint64          `json:"block_height"`
	BTCBlockHash     string          `json:"block_hash"`
	BTCPrevBlockHash string          `json:"prev_block_hash"`
	BTCBlockTime     uint32          `json:"block_time"`
	ZeroTxs          []ZeroRequestTx `json:"txs"`
}

type CrawlerHeightData struct {
	CrawlerHeight uint64 `json:"crawler_height"`
}
