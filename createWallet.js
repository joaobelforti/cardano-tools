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
 
  const wallet = createWallet("wallet5");
  //const wallet = cardano.wallet("teste4");
  //console.log(cardano.queryTip(wallet));
  //console.log(cardano.queryUtxo("addr1vxjx5l4xwm9f04e6mhh0mp88z6u77v9v4r84p7qzx9usv8gn77mz4"));
  //console.log("A");
