var shelleyGenesisPath = "../../cardano-node-testnet/testnet-shelley-genesis.json";
var socketPath = "../../cardano-node-testnet/path/to/db/node.socket";
const exec =typeof window !== "undefined" || require("child_process").execSync;

var slot=exec(`cardano-cli query tip --testnet-magic 1097911063`)
slot = JSON.parse(slot)
console.log(slot.slot)