const CardanocliJs = require("cardanocli-js");

const { sign } = require("crypto");

const shelleyGenesisPath = "../cardano-src/cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-src/cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath , socketPath});

for(let n = 1;n < process.argv[2]; n++){

	const dir = ["wallet".concat(n.toString())];
	
	const wallet = cardano.wallet(dir);
	
	console.log(cardano.queryUtxo(wallet.paymentAddr));
	
}
