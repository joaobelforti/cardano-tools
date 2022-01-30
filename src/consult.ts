const { lookup } = require("dns");
import { argv } from 'process';
var shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";
var socketPath = "../cardano-node/path/to/db/node.socket";
const exec =typeof window !== "undefined" || require("child_process").execSync;

async function asyncUtxos(address){
    //return new Promise(resolve => {
    console.log(address)
    return await new Promise(function(execute) {
        execute(exec(`cardano-cli query utxo \
            --mainnet \
            --address ${address} \
            --cardano-mode
            `).toString())
    });
}

async function loop(){
    for (var n = 0; n < process.argv[2]; n++) {
        var wallet = ["wallet".concat((n + 1).toString())];
        //var wallet = cardano.wallet(dir);
        const dir = "priv/wallet/"+wallet+"/"+wallet+".payment.addr"
        const address = exec(`cat ${dir}`)
        console.log("wallet", n + 1);
        console.log(address)
        var utxos= await Promise.all([asyncUtxos(address)])
        console.log(utxos, "\n");
    }
}
(async() => {
    await loop();
})();