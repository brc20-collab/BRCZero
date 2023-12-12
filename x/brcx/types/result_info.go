package types

type ResultInfo struct {
	BTCTxid      string `json:"btc_txid"`
	BTCBlockHash string `json:"btc_block_hash"`
	EvmCaller    string `json:"evm_caller"`
	EvmTo        string `json:"evm_to"`
	Nonce        uint64 `json:"nonce"`
	CallData     string `json:"call_data"`
}
