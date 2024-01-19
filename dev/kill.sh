#!/bin/bash

NAME=$1
ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2", "$8}'
ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2}' | sh
echo "All <$NAME> killed!"