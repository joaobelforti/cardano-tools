const CardanocliJs = require("cardanocli-js");

const shelleyGenesisPath = "../../cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../../cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath,socketPath });
const exec =typeof window !== "undefined" || require("child_process").execSync;


for (var n = 0; n < process.argv[2]; n++) {
        var wallet = ["wallet".concat((n+1).toString())];
        //var wallet = cardano.wallet(dir);
        const dir = "priv/wallet/"+wallet+"/"+wallet+".payment.addr"
        const address = exec(`cat ${dir}`).toString()
        console.log("wallet",n+1,"->",address)
}