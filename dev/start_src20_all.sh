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

echo "************* Start ord... *************"


echo "************* Start brczero node... *************"
#cd $self_path
#./start.sh
#rm -rf ./_cache2
#nohup ./target/debug/ord \
#  --log-level=DEBUG \
#  --data-dir=./_cache2 \
#  --rpc-url=http://localhost:18443 \
#  --regtest \
#  --bitcoin-rpc-user bitcoinrpc \
#  --bitcoin-rpc-pass bitcoinrpc \
#  --brczero-rpc-url=http://127.0.0.1:26757 \
#  --first-brczero-height=120 \
#  server --http-port=81 >/dev/null 2>&1 &
#
#rm -rf ./_cache3
#nohup ./target/debug/ord \
# --log-level=DEBUG \
# --data-dir=./_cache3 \
# --rpc-url=http://localhost:18443 \
# --regtest \
# --bitcoin-rpc-user bitcoinrpc \
# --bitcoin-rpc-pass bitcoinrpc \
# --brczero-rpc-url=http://127.0.0.1:26857 \
# --first-brczero-height=120 \
# server --http-port=82 >/dev/null 2>&1 &
#
#rm -rf ./_cache4
#nohup ./target/debug/ord \
#--log-level=DEBUG \
#--data-dir=./_cache4 \
#--rpc-url=http://localhost:18443 \
#--regtest \
#--bitcoin-rpc-user bitcoinrpc \
#--bitcoin-rpc-pass bitcoinrpc \
#--brczero-rpc-url=http://127.0.0.1:26957 \
#--first-brczero-height=120 \
#server --http-port=83 >/dev/null 2>&1 &
#
#rm -rf ./_cache5
#nohup ./target/debug/ord \
# --log-level=DEBUG \
# --data-dir=./_cache5 \
# --rpc-url=http://localhost:18443 \
# --regtest \
# --bitcoin-rpc-user bitcoinrpc \
# --bitcoin-rpc-pass bitcoinrpc \
# --brczero-rpc-url=http://127.0.0.1:27057 \
# --first-brczero-height=120 \
# server --http-port=84 >/dev/null 2>&1 &

