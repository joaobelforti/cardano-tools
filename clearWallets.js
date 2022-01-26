const CardanocliJs = require("cardanocli-js");

const exec = typeof window !== "undefined" || require("child_process").execSync;

const shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";

const socketPath = "../cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath , socketPath});

const addr = "addr1q8dtzpf2yh9ms2e07pck0x9jl2ufmky9y526pjept2vscxhnt267zvmmkhtpt3kala3ewnehhvtf2t4kgd98gpqcrxzqqfx6au"

for(let n = 0;n < process.argv[2]; n++){
 
    const dir = ["wallet".concat((n+1).toString())];
    const wallet = cardano.wallet(dir);

    if(cardano.queryUtxo(wallet.paymentAddr).length!=0){
        console.log(exec(`node sendAll.js ${dir} ${addr}`))
    }

}