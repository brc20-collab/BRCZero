#!/bin/bash

./kill.sh brczerod
docker-compose -f bitcoin.yml down
./kill.sh target/debug/ord
