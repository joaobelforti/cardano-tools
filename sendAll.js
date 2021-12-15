const CardanocliJs = require("cardanocli-js");

const shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath,socketPath });

const receiver = process.argv[3];

let fee=0;

const wallet = [process.argv[2]];

const sender = cardano.wallet(wallet);

const txInfo = {
    txIn: [...cardano.queryUtxo(sender.paymentAddr)],
    txOut: [
	{
	address: receiver,
      value: {
        ...sender.balance().value
      },
    },
  ],
};

delete txInfo.txOut[0].value.undefined

const raw = cardano.transactionBuildRaw({ ...txInfo, fee });

fee = cardano.transactionCalculateMinFee({
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