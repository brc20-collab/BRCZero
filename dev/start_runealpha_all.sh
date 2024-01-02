#!/bin/bash

./kill_all.sh

set -e
set -o errexit
set -a
set -m

self_path=$(pwd)
# val0 26656  26657  8545
# val1 26756  26757  8645
# val2 26856  26857  8745
# val3 26956  26957  8845
# rpc  27056  27057  8945

echo "************* Start btc node... *************"
rm -rf ./bitcoin-data/regtest
docker-compose -f bitcoin.yml up -d
cp -r ./bitcoin-data/wallets ./bitcoin-data/regtest/
sleep 5
docker exec -it local_bitcoin_node bitcoin-cli loadwallet testwallet_01
docker exec -it local_bitcoin_node bitcoin-cli -rpcwallet=testwallet_01 getwalletinfo
docker exec -it local_bitcoin_node bitcoin-cli generatetoaddress 120 bcrt1qd28jewrz9y9hpl328em5fpljvgarucgcxf7fjt
docker exec -it local_bitcoin_node bitcoin-cli -rpcwallet=testwallet_01 getwalletinfo

echo "************* Start rune ord... *************"
cd /Users/oker/go/src/github.com/okex/runealpha-ord
cargo build
rm -rf ./_cache_runealpha
nohup ./target/debug/ord --regtest --rpc-url=http://localhost:18443 --bitcoin-rpc-user bitcoinrpc --bitcoin-rpc-pass bitcoinrpc --index-runes-pre-alpha-i-agree-to-get-rekt --data-dir ./_cache_runealpha server --enable-json-api --http-port 83 >./rune_ord.log 2>&1 &
sleep 5

echo "************* Start brczero node... *************"
cd $self_path
./start_runealpha.sh