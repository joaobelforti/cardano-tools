const CardanocliJs = require("cardanocli-js");

const shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath,socketPath });

for(let n = 0;n < process.argv[2]; n++){
	const dir = ["wallet".concat((n+1).toString())];

	const wallet = cardano.wallet(dir);

	console.log("wallet",n+1," -> ",wallet.balance().value,"ADA")

	console.log(wallet.paymentAddr)

	console.log(cardano.queryUtxo(wallet.paymentAddr),"\n");
}