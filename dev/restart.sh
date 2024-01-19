#!/bin/bash

KEY="captain"
CHAINID="brczero-67"
MONIKER="brczero"
CURDIR=`dirname $0`
HOME_SERVER=$CURDIR/"_cache_evm"

set -e
set -o errexit
set -a
set -m


killbyname() {
  NAME=$1
  ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -2 "$2", "$8}'
  ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -2 "$2}' | sh
  echo "All <$NAME> killed!"
}


run() {
    LOG_LEVEL=main:info,iavl:info,*:error,state:info,provider:info
    nohup brczerod start --rpc.unsafe \
      --local-rpc-port 26657 \
      --log_level $LOG_LEVEL \
      --log_file json \
      --dynamic-gp-mode=2 \
      --consensus.timeout_commit 4000ms \
      --consensus.create_empty_blocks=false \
      --consensus.start_btc_height=120 \
      --zero-data-url="http://0.0.0.0/api/v1/crawler/zeroindexer/" \
      --tree-enable-async-commit=false \
      --enable-gid \
      --fast-query=false \
      --append-pid=true \
      --iavl-output-modules evm=0,acc=0 \
      --commit-gap-height 3 \
      --pruning=nothing \
      --trie.dirty-disabled=true \
      --trace --home $HOME_SERVER --chain-id $CHAINID \
      --elapsed Round=1,CommitRound=1,Produce=1 \
      --rest.laddr "tcp://0.0.0.0:8545" >> brc10.txt 2>&1 &

}


killbyname brczerod
killbyname brczerocli

set -x # activate debugging

# run



# Set up config for CLI
brczerocli config chain-id $CHAINID
brczerocli config output json
brczerocli config indent true
brczerocli config trust-node true
brczerocli config keyring-backend test

# if $KEY exists it should be deleted
#
#    "eth_address": "0xbbE4733d85bc2b90682147779DA49caB38C0aA1F",
#     prikey: 8ff3ca2d9985c3a52b459e2f6e7822b23e1af845961e22128d5f372fb9aa5f17
brczerocli keys add --recover captain -m "puzzle glide follow cruel say burst deliver wild tragic galaxy lumber offer" -y

#    "eth_address": "0x83D83497431C2D3FEab296a9fba4e5FaDD2f7eD0",
brczerocli keys add --recover admin16 -m "palace cube bitter light woman side pave cereal donor bronze twice work" -y

brczerocli keys add --recover admin17 -m "antique onion adult slot sad dizzy sure among cement demise submit scare" -y

brczerocli keys add --recover admin18 -m "lazy cause kite fence gravity regret visa fuel tone clerk motor rent" -y

# Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
brczerocli config keyring-backend test

run

sleep 1

