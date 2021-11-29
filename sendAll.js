const CardanocliJs = require("cardanocli-js");

const { sign } = require("crypto");

const shelleyGenesisPath = "../cardano-src/cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-src/cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath , socketPath});

const value = 1;

const receiver = "addr1q8c3fk3hwqsras54gggv99gd84yqsvw66x76trn2k22g6rlnt267zvmmkhtpt3kala3ewnehhvtf2t4kgd98gpqcrxzqcupqml";
    
    let fee=0;
    
    //const wallet = ["wallet".concat(i.toString())];
    const wallet = ["wallet1"];
    const sender = cardano.wallet(wallet);

    const txInfo = {
      txIn: cardano.queryUtxo(sender.paymentAddr),
      txOut: [
        {
          address: receiver,
          value: {
            lovelace: sender.balance().value.lovelace-fee,
            // send NFT "ad9c09fa0a62ee42fb9555ef7d7d58e782fa74687a23b62caf3a8025.BerrySpaceGreen": 1
          },
        },
      ],
    };

    const raw = cardano.transactionBuildRaw(txInfo);

    fee = cardano.transactionCalculateMinFee({
      ...txInfo,
      txBody: raw,
      witnessCount: 1,
    });
    console.log(sender.balance().value.lovelace)
    console.log(txInfo.txOut[0].value.lovelace)
    console.log(txInfo.txOut[0])
    txInfo.txOut[0].value.lovelace -= fee;
    console.log(txInfo.txOut[0].value.lovelace)

    const tx = cardano.transactionBuildRaw({ ...txInfo, fee });

    const txSigned = cardano.transactionSign({
      txBody: tx,
      signingKeys: [sender.payment.skey],
    });

    const txHash = cardano.transactionSubmit(txSigned);

    console.log(txHash);

