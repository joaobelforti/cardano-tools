const CardanocliJs = require("cardanocli-js");

const shelleyGenesisPath = "../../cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../../cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath,socketPath });

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
        lovelace: sender.balance().value.lovelace - cardano.toLovelace(value),
      },
    },
  {
    address: receiver,
    value: {
        lovelace: cardano.toLovelace(value),
        "dac355946b4317530d9ec0cb142c63a4b624610786c2a32137d78e25.adapeSalvadorKing":1,
        "dac355946b4317530d9ec0cb142c63a4b624610786c2a32137d78e25.adapeTannerWhite":1,
        "dac355946b4317530d9ec0cb142c63a4b624610786c2a32137d78e25.adapeCharlemagneMaddock":1,
        "dac355946b4317530d9ec0cb142c63a4b624610786c2a32137d78e25.adapeRansleydelaCruz":1
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