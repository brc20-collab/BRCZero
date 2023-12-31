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
  ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2", "$8}'
  ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2}' | sh
  echo "All <$NAME> killed!"
}


run() {
    LOG_LEVEL=main:info,iavl:info,*:error,state:info,provider:info
#--mempool.enable_delete_min_gp_tx false \
#    brczerod start --pruning=nothing --rpc.unsafe \
    nohup brczerod start --rpc.unsafe \
      --local-rpc-port 26657 \
      --log_level $LOG_LEVEL \
      --log_file json \
      --dynamic-gp-mode=2 \
      --consensus.timeout_commit 2000ms \
      --enable-preruntx=1 \
      --tree-enable-async-commit=false \
      --enable-gid \
      --fast-query=false \
      --append-pid=true \
      --iavl-output-modules evm=0,acc=0 \
      --commit-gap-height 3 \
      --trie.dirty-disabled=true \
      --trace --home $HOME_SERVER --chain-id $CHAINID \
      --elapsed Round=1,CommitRound=1,Produce=1 \
      --rest.laddr "tcp://localhost:8545" > brc10.txt 2>&1 &

# --iavl-commit-interval-height \
# --iavl-enable-async-commit \
#      --iavl-cache-size int                              Max size of iavl cache (default 1000000)
#      --iavl-commit-interval-height int                  Max interval to commit node cache into leveldb (default 100)
#      --iavl-debug int                                   Enable iavl project debug
#      --iavl-enable-async-commit                         Enable async commit
#      --iavl-enable-pruning-history-state                Enable pruning history state
#      --iavl-height-orphans-cache-size int               Max orphan version to cache in memory (default 8)
#      --iavl-max-committed-height-num int                Max committed version to cache in memory (default 8)
#      --iavl-min-commit-item-count int                   Min nodes num to triggle node cache commit (default 500000)
#      --iavl-output-modules

}


killbyname brczerod
killbyname brczerocli

set -x # activate debugging

# run

# remove existing daemon and client
rm -rf ~/.brczero*
rm -rf $HOME_SERVER

(cd .. && make install)

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
brczerod init $MONIKER --chain-id $CHAINID --home $HOME_SERVER

# Change parameter token denominations to brc10
cat $HOME_SERVER/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="brc10"' > $HOME_SERVER/config/tmp_genesis.json && mv $HOME_SERVER/config/tmp_genesis.json $HOME_SERVER/config/genesis.json
cat $HOME_SERVER/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="brc10"' > $HOME_SERVER/config/tmp_genesis.json && mv $HOME_SERVER/config/tmp_genesis.json $HOME_SERVER/config/genesis.json
cat $HOME_SERVER/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="brc10"' > $HOME_SERVER/config/tmp_genesis.json && mv $HOME_SERVER/config/tmp_genesis.json $HOME_SERVER/config/genesis.json
cat $HOME_SERVER/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="brc10"' > $HOME_SERVER/config/tmp_genesis.json && mv $HOME_SERVER/config/tmp_genesis.json $HOME_SERVER/config/genesis.json

# Enable EVM

if [ "$(uname -s)" == "Darwin" ]; then
    sed -i "" 's/"enable_call": false/"enable_call": true/' $HOME_SERVER/config/genesis.json
    sed -i "" 's/"enable_create": false/"enable_create": true/' $HOME_SERVER/config/genesis.json
    sed -i "" 's/"enable_contract_blocked_list": false/"enable_contract_blocked_list": true/' $HOME_SERVER/config/genesis.json
else
    sed -i 's/"enable_call": false/"enable_call": true/' $HOME_SERVER/config/genesis.json
    sed -i 's/"enable_create": false/"enable_create": true/' $HOME_SERVER/config/genesis.json
    sed -i 's/"enable_contract_blocked_list": false/"enable_contract_blocked_list": true/' $HOME_SERVER/config/genesis.json
fi

# Allocate genesis accounts (cosmos formatted addresses)
brczerod add-genesis-account $(brczerocli keys show $KEY    -a) 100000000brc10 --home $HOME_SERVER
brczerod add-genesis-account $(brczerocli keys show admin16 -a) 900000000brc10 --home $HOME_SERVER
brczerod add-genesis-account $(brczerocli keys show admin17 -a) 900000000brc10 --home $HOME_SERVER
brczerod add-genesis-account $(brczerocli keys show admin18 -a) 900000000brc10 --home $HOME_SERVER

# Sign genesis transaction
brczerod gentx --name $KEY --keyring-backend test --home $HOME_SERVER

# Collect genesis tx
brczerod collect-gentxs --home $HOME_SERVER

# Run this to ensure everything worked and that the genesis file is setup correctly
brczerod validate-genesis --home $HOME_SERVER
brczerocli config keyring-backend test

run

sleep 4

# brczerocli tx send captain 0x83D83497431C2D3FEab296a9fba4e5FaDD2f7eD0 1brc10 --fees 1brc10 -b block -y
