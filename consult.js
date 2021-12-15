var CardanocliJs = require("cardanocli-js");
var shelleyGenesisPath = "../cardano-node/mainnet-shelley-genesis.json";
var socketPath = "../cardano-node/path/to/db/node.socket";
var cardano = new CardanocliJs({ shelleyGenesisPath: shelleyGenesisPath, socketPath: socketPath });
for (var n = 0; n < process.argv[2]; n++) {
    var dir = ["wallet".concat((n + 1).toString())];
    var wallet = cardano.wallet(dir);
    console.log("wallet", n + 1, " -> ", wallet.balance().value, "ADA");
    console.log(wallet.paymentAddr);
    console.log(cardano.queryUtxo(wallet.paymentAddr), "\n");
}
