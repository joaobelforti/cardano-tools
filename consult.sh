#!/bin/bash

export CARDANO_NODE_SOCKET_PATH=../cardano-src/cardano-node/path/to/db/node.socket

for (( i=1; i<=$1; i++ ))
do
	walletDir=priv/wallet
	walletPayment=${walletDir}/wallet${i}/wallet${i}.payment.addr

	cardano-cli query utxo \
        --address $(cat $walletPayment) \
        --mainnet
done
