const CardanocliJs = require("cardanocli-js");

const shelleyGenesisPath = "../cardano-src/cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-src/cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath , socketPath});

console.log(process.argv[2]+1)

for(let n = 0;n < process.argv[2]; n++){
	const dir = ["wallet".concat((n+1).toString())];

	console.log("wallet",n+1)

	const wallet = cardano.wallet(dir);
	
	console.log(cardano.queryUtxo(wallet.paymentAddr));
}
