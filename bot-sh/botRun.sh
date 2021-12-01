#!/bin/bash

export CARDANO_NODE_SOCKET_PATH=../../cardano-src/cardano-node/path/to/db/node.socket


for (( i=1; i<=$1; i++ ))
do
      ./nftBot.sh $i $2 $3
done
