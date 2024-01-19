tx=$(curl -s http://localhost:8545/v1/block_tx_hashes/$1 | jq -r '.[0]')
brczerocli query tx $tx | jq -r '.'

