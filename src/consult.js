"use strict";
exports.__esModule = true;
var lookup = require("dns").lookup;
var shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";
var socketPath = "../cardano-node/path/to/db/node.socket";
var exec = typeof window !== "undefined" || require("child_process").execSync;
for (var n = 0; n < process.argv[2]; n++) {
    var wallet = ["wallet".concat((n).toString())];
    //var wallet = cardano.wallet(dir);
    var dir = "priv/wallet/" + wallet + "/" + wallet + ".payment.addr";
    var address = exec("cat " + dir).toString();
    var utxos = exec("cardano-cli query utxo             --mainnet             --address " + address + "             --cardano-mode\n            ").toString();
    console.log("wallet", n);
    console.log(address);
    console.log(utxos, "\n");
}
