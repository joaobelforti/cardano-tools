const { lookup } = require("dns");
import { argv } from 'process';
var shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";
var socketPath = "../cardano-node/path/to/db/node.socket";
const exec =typeof window !== "undefined" || require("child_process").execSync;


for (var n = 0; n < process.argv[2]; n++) {
        var wallet = ["wallet".concat((n ).toString())];
        //var wallet = cardano.wallet(dir);
        const dir = "priv/wallet/"+wallet+"/"+wallet+".payment.addr"
        const address = exec(`cat ${dir}`).toString()
        const utxos=exec(`cardano-cli query utxo \
            --mainnet \
            --address ${address} \
            --cardano-mode
            `).toString()
        console.log("wallet", n);
        console.log(address)
        
        console.log(utxos, "\n");
    }