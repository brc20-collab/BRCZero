## Starting a single-node brczero network locally
```
./start.sh
```

## Starting a multi-node brczero network locally
```
./testnet.sh -s -i -n 4
```
* `-i` delete data and reinitialize
* `-s` start node
* `-n` total number of nodes
* `-r` number of rpc nodes in the total number of nodes

## Starting a all-node network locally
```
./start_all.sh
```
* start 4 validator nodes
* start 1 btc node
* start 1 ord server

## Stop all node and service
```
./kill_all.sh
```