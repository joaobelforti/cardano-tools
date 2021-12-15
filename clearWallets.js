const CardanocliJs = require("cardanocli-js");

const exec = typeof window !== "undefined" || require("child_process").execSync;

const shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath , socketPath});

const addr = "addr1q8c3fk3hwqsras54gggv99gd84yqsvw66x76trn2k22g6rlnt267zvmmkhtpt3kala3ewnehhvtf2t4kgd98gpqcrxzqcupqml"

for(let n = 0;n < process.argv[2]; n++){

    const dir = ["wallet".concat((n+1).toString())];

	const wallet = cardano.wallet(dir);

    console.log(exec(`node sendAll.js ${dir} ${addr}`))

}