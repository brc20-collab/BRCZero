request=$(curl --location "http://127.0.0.1:80/api/v1/brc0/rpc_request/$1" | jq -r '.data')
echo "{\"jsonrpc\":\"2.0\",\"id\":3,\"method\":\"broadcast_brczero_txs_async\",\"params\":$request}"
curl --location 'http://localhost:26657' \
--header 'Content-Type: application/json' \
--data "{\"jsonrpc\":\"2.0\",\"id\":3,\"method\":\"broadcast_brczero_txs_async\",\"params\":$request}"
