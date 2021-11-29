#!/bin/bash
export CARDANO_NODE_SOCKET_PATH=../../cardano-src/cardano-node/path/to/db/node.socket
amountToSend=$(($2*1000000))
walletDir=../priv/wallet/wallet${1}
walletPayment=${walletDir}/wallet${1}.payment.addr
walletSkey=${walletDir}/wallet${1}.payment.skey
destinationAddress=$3

cardano-cli query protocol-parameters \
                  --mainnet \
                              --out-file protocol.json


currentSlot=$(cardano-cli query tip --mainnet | jq -r '.slot')

cardano-cli query utxo \
        --address $(cat $walletPayment) \
        --mainnet > fullUtxo.out

tail -n +3 fullUtxo.out | sort -k3 -nr > balance.out
txOut=0
txcnt=$(cat balance.out | wc -l)
total_balance=0
i=0
zero=0
subtract=0
while read -r utxo; do
                    in_addr=$(awk '{ print $1 }' <<< "${utxo}")
                idx=$(awk '{ print $2 }' <<< "${utxo}")
                ada=$(awk '{ print $3 }' <<< "${utxo}")
                    utxo_balance=$(awk '{ print $3 }' <<< "${utxo}")
                        total_balance=$((${total_balance}+${utxo_balance}))
                        if [ $i -eq $zero ];
                                then
                                        tx_in="${tx_in} --tx-in ${in_addr}#${idx}"
                                else
                                        subtract=$((${ada}+${subtract}))
                        fi
                        i=$((${i}+1))
                    done < balance.out
                    txcnt=$(cat balance.out | wc -l)

echo ${tx_in}
cat fullUtxo.out

cardano-cli transaction build-raw \
    ${tx_in} \
    --tx-out $(cat $walletPayment)+0 \
    --tx-out ${destinationAddress}+0 \
    --invalid-hereafter $(( ${currentSlot} + 10000)) \
    --fee 0 \
    --out-file tx.tmp

fee=$(cardano-cli transaction calculate-min-fee \
    --tx-body-file tx.tmp \
    --tx-in-count ${txcnt} \
    --tx-out-count 2 \
    --mainnet \
    --witness-count 1 \
    --byron-witness-count 0 \
    --protocol-params-file protocol.json | awk '{ print $1 }')

txOut=$((${total_balance}-${fee}-${amountToSend}-${subtract}))

cardano-cli transaction build-raw \
    ${tx_in} \
    --tx-out $(cat $walletPayment)+${txOut} \
    --tx-out ${destinationAddress}+${amountToSend} \
    --invalid-hereafter $(( ${currentSlot} + 10000)) \
    --fee ${fee} \
    --out-file tx.raw

cardano-cli transaction sign \
    --tx-body-file tx.raw \
    --signing-key-file $walletSkey \
    --mainnet \
    --out-file tx.signed

cardano-cli transaction submit \
    --tx-file tx.signed \
    --mainnet
