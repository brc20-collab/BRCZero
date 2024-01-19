#!/bin/bash

set -e
set -o errexit
set -a
set -m

BIN_NAME=brczerod
CHAIN_TOP=../
CHAIN_ID="brczero-67"
NUM_NODE=4
NUM_RPC=0
BASE_PORT_PREFIX=26600
P2P_PORT_SUFFIX=56
RPC_PORT_SUFFIX=57
REST_PORT=8545

let BASE_PORT=${BASE_PORT_PREFIX}+${P2P_PORT_SUFFIX}
let seedp2pport=${BASE_PORT_PREFIX}+${P2P_PORT_SUFFIX}
let seedrpcport=${BASE_PORT_PREFIX}+${RPC_PORT_SUFFIX}
let seedrestport=${seedrpcport}+1

while getopts "r:isn:b:p:" opt; do
  case $opt in
  i)
    echo "CHAIN_INIT"
    CHAIN_INIT=1
    ;;
  r)
    echo "NUM_RPC=$OPTARG"
    NUM_RPC=$OPTARG
    ;;
  s)
    echo "CHAIN_START"
    CHAIN_START=1
    ;;
  n)
    echo "NUM_NODE=$OPTARG"
    NUM_NODE=$OPTARG
    ;;
  b)
    echo "BIN_NAME=$OPTARG"
    BIN_NAME=$OPTARG
    ;;
  p)
    echo "IP=$OPTARG"
    IP=$OPTARG
    ;;
  \?)
    echo "Invalid option: -$OPTARG"
    ;;
  esac
done

echorun() {
  echo "------------------------------------------------------------------------------------------------"
  echo "["$@"]"
  $@
  echo "------------------------------------------------------------------------------------------------"
}

killbyname() {
  NAME=$1
  ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2", "$8}'
  ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2}' | sh
  echo "All <$NAME> killed!"
}

init() {
  killbyname ${BIN_NAME}

  (cd ${CHAIN_TOP} && make install)

  rm -rf cache

  echo "===== Generate testnet configurations files... ===="
  echorun brczerod testnet --v $1 --r $2 -o cache -l \
    --chain-id ${CHAIN_ID} \
    --starting-ip-address ${IP} \
    --base-port ${BASE_PORT} \
    --keyring-backend test
}

run() {

  index=$1
  seed_mode=$2
  p2pport=$3
  rpcport=$4
  restport=$5
  p2p_seed_opt=$6
  p2p_seed_arg=$7

  brczerod add-genesis-account 0xbbE4733d85bc2b90682147779DA49caB38C0aA1F 900000000okt --home cache/node${index}/brczerod
  LOG_LEVEL=main:info,*:error

  echorun nohup brczerod start \
    --home cache/node${index}/brczerod \
    --p2p.seed_mode=$seed_mode \
    --p2p.allow_duplicate_ip \
    --p2p.pex=false \
    --p2p.addr_book_strict=false \
    $p2p_seed_opt $p2p_seed_arg \
    --p2p.laddr tcp://${IP}:${p2pport} \
    --rpc.laddr tcp://${IP}:${rpcport} \
    --log_level ${LOG_LEVEL} \
    --chain-id ${CHAIN_ID} \
    --consensus.timeout_commit 3s \
    --consensus.create_empty_blocks=false \
    --elapsed DeliverTxs=0,Round=1,CommitRound=1,Produce=1 \
    --rest.laddr tcp://localhost:$restport \
    --consensus-role=v$index \
    --active-view-change=false \
    --deliver-txs-mode=0 \
    --keyring-backend test >cache/val${index}.log 2>&1 &
}

function start() {
  killbyname ${BIN_NAME}
  index=0

  echo "=========== Startup seed node...============"
  ((restport = REST_PORT))
  run $index true ${seedp2pport} ${seedrpcport} $restport
  seed=$(brczerod tendermint show-node-id --home cache/node${index}/brczerod)

  echo "======== Startup validator nodes...========="
  for ((index = 1; index < ${1}; index++)); do
    ((p2pport = BASE_PORT_PREFIX + index * 100 + P2P_PORT_SUFFIX))
    ((rpcport = BASE_PORT_PREFIX + index * 100 + RPC_PORT_SUFFIX))
    ((restport = index * 100 + REST_PORT))
    run $index false ${p2pport} ${rpcport} $restport --p2p.seeds ${seed}@${IP}:${seedp2pport}
  done
  echo "start node done"
}

if [ -z ${IP} ]; then
  IP="127.0.0.1"
fi

if [ ! -z "${CHAIN_INIT}" ]; then
	((NUM_VAL=NUM_NODE-NUM_RPC))
  init ${NUM_VAL} ${NUM_RPC}
fi

if [ ! -z "${CHAIN_START}" ]; then
  start ${NUM_NODE}
fi
