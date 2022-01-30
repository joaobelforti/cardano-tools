const exec = typeof window !== "undefined" || require("child_process").execSync;
async function create(n){
    return new Promise(function(execute) {
        execute(exec(`node createWallet.js wallet${n}`))
    });
}
for(let n = process.argv[2];n <= process.argv[3]; n++){
    create(n)
}