# cardano-tools

## Credits to <a href="https://github.com/Berry-Pool/">Berry-Pool</a>.

## Overview
This repository was made to help you at managing many wallets, by consulting their addresses, balances, sending multiple transaction and other features.
For simple management, all wallets are named like "wallet1", "wallet2" etc, its simple to change it on code if you want.
##

### Prerequisites 
### Install <a href="https://github.com/Berry-Pool/cardanocli-js">cardanocli-js</a>

##

### sendTx.js -> Send a single transaction, in this case from wallet 3, value 5 ADA to addr1 below.
```
node sendTx.js 3 5 addr1q8c3fk3hwqsras54gggv99gd84yqsvw66x76trn2k22g6rlnt267zvmmkhtpt3kala3ewnehhvtf2t4kgd98gpqcrxzqcupqml
```
##
### sendAll.js -> Send all ADA, in this case from wallet4, to addr1 below.
```
node sendAll.js wallet4 addr1q8c3fk3hwqsras54gggv99gd84yqsvw66x76trn2k22g6rlnt267zvmmkhtpt3kala3ewnehhvtf2t4kgd98gpqcrxzqcupqml
```
##
#### ./botRun runs an algorithm that send single transactions from many wallets, in this case, it sends 1 ADA (first number), from 3 wallets (second number) to the addr1 below ###
```
./botRun 3 1 addr1q8c3fk3hwqsras54gggv99gd84yqsvw66x76trn2k22g6rlnt267zvmmkhtpt3kala3ewnehhvtf2t4kgd98gpqcrxzqcupqml
```
##
### List all wallets funds and addresses, in this case from wallet 1 to wallet 10.
``` 
node consult.js 10 
```
##
### Create new wallet on priv dir.
```
node createWallet.js walletName
```

## sendTx.js 
```
node sendTx.js 3 5 addr1q8c3fk3hwqsras54gggv99gd84yqsvw66x76trn2k22g6rlnt267zvmmkhtpt3kala3ewnehhvtf2t4kgd98gpqcrxzqcupqml
```
