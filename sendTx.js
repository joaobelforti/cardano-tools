const CardanocliJs = require("cardanocli-js");

const { sign } = require("crypto");

const shelleyGenesisPath = "../cardano-src/cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-src/cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath , socketPath});

const wallet = ["wallet".concat(process.argv[2].toString())];
const value = process.argv[3]
const receiver = process.argv[4]
    
    const sender = cardano.wallet(wallet);

    const txInfo = {
      txIn: cardano.queryUtxo(sender.paymentAddr),
      txOut: [
        {
          address: sender.paymentAddr,
          value: {
            lovelace:
              sender.balance().value.lovelace - cardano.toLovelace(value),
          },
        },
        {
          address: receiver,
          value: {
            lovelace: cardano.toLovelace(value),
            // send NFT "ad9c09fa0a62ee42fb9555ef7d7d58e782fa74687a23b62caf3a8025.BerrySpaceGreen": 1
          },
        },
      ],
    };

    const raw = cardano.transactionBuildRaw(txInfo);

    const fee = cardano.transactionCalculateMinFee({
      ...txInfo,
      txBody: raw,
      witnessCount: 1,
    });

    txInfo.txOut[0].value.lovelace -= fee;

    const tx = cardano.transactionBuildRaw({ ...txInfo, fee });

    const txSigned = cardano.transactionSign({
      txBody: tx,
      signingKeys: [sender.payment.skey],
    });

    const txHash = cardano.transactionSubmit(txSigned);

    console.log(txHash);

