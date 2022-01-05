const CardanocliJs = require("cardanocli-js");
const { networkInterfaces } = require("os");

const shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath,socketPath });

var sender = cardano.wallet("wallet1");
value = process.argv[3]
const txInfo = {
    txIn: cardano.queryUtxo(sender.paymentAddr),
    txOut: [
  {
    address: sender.paymentAddr,
    value: {
        lovelace: sender.balance().value.lovelace - cardano.toLovelace(value*(process.argv[2]-1)),
      },
    },
  ],
};

for (var n = 1; n < process.argv[2]; n++) {

    var dir = ["wallet".concat((n + 1).toString())];
    var wallet = cardano.wallet(dir);
    receiver=wallet.paymentAddr
    const newAddr={
    address: receiver,
    value: {
        lovelace: cardano.toLovelace(value)
      },
    }
    txInfo.txOut.push(newAddr)
}

const raw = cardano.transactionBuildRaw(txInfo);

const fee = cardano.transactionCalculateMinFee({
    ...txInfo,
    txBody: raw,
    witnessCount: parseInt(process.argv[2]-1),
});

txInfo.txOut[0].value.lovelace -= fee;

console.log(txInfo.txOut)

const tx = cardano.transactionBuildRaw({ ...txInfo, fee });

const txSigned = cardano.transactionSign({
    txBody: tx,
    signingKeys: [sender.payment.skey],
});

const txHash = cardano.transactionSubmit(txSigned);

console.log(txHash);