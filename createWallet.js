const CardanocliJs = require("cardanocli-js");

const shelleyGenesisPath = "../cardano-src/cardano-node/mainnet-shelley-genesis.json";
const socketPath = "../cardano-src/cardano-node/path/to/db/node.socket";

const cardano = new CardanocliJs({ shelleyGenesisPath,socketPath });

const createWallet = (account) => {
    const addr = cardano.addressKeyGen(account);
    cardano.stakeAddressKeyGen(account);
    cardano.stakeAddressBuild(account);
    cardano.addressBuild(account, {
        paymentVkey: addr.vkey,
      });
    return cardano.wallet(account);
  };
 
  const wallet = createWallet(process.argv[2]);