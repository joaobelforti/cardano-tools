package main

import (
    "log"
    "sync"
    "os"
    "strconv"
    "os/exec"
    "time"
    "fmt"
    "regexp"
    "strings"
    "bufio"
)

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex
    errorCount:=0
    size, _ := strconv.Atoi(os.Args[1])
    amountToSend,_:=strconv.Atoi(os.Args[2])
    amountToSend=amountToSend*1000000 
    txsIn := make([]string, size)
    totalValues := make([]string, size)
    dir:="priv/wallet/"

    //load txs in before sending starts and stores it at an array
    fmt.Println("Loading...")
    wg.Add(size)
    for i := 0; i < size; i++ {
        go func(i int) {
            totalValue:=0
            //wallets config
            wallet:="wallet"+strconv.Itoa(i+1)
            walletDir:=dir+wallet
            walletPaymentAddress:=walletDir+"/"+wallet+".payment.addr"
            walletAddressAux,_:=exec.Command("cat",walletPaymentAddress).Output()   
            walletAddress := string(walletAddressAux[:])

            //search on cardano blockchain each address transcations
            cmd,_ := exec.Command("cardano-cli","query","utxo",
            "--address",walletAddress,
            "--mainnet").Output()
            //output takes return of cmd variable and turns into a string.
            output := string(cmd[:])
            //turn string into array splitting all empty spaces.
            a := regexp.MustCompile("[^\\s]+")
            array:=a.FindAllString(output,-1)
            //return an error if some wallet is not filled
            if(len(array)==4){
                log.Fatalf("\nERROR! WALLET NOT FILLED.")
            }

            //it selects the tx with higher ada volume as a tx-in.
            for x := 0; x < len(array); x++ {
                if(array[x]=="lovelace"){
                    lovelace,_:=strconv.Atoi(array[x-1])
                    if(totalValue < lovelace){
                        txsIn[i]=array[x-3]+"#"+array[x-2]
                        totalValues[i]=array[x-1]
                        totalValue,_=strconv.Atoi(array[x-1])
                    }
                }   
            }
        defer wg.Done()
        }(i)
    }
    wg.Wait()

    //reveice minting address to spam transactions.
    fmt.Println("All txs are loaded - ready to START.")
    fmt.Println("Transactions to be sent:",size)
    fmt.Println("Amount to send:",amountToSend/1000000)
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("To address: ")
    destinationAddress, _ := reader.ReadString('\n')

    //start a timer until program finish.
    start := time.Now()

    //get protocol.json, file required for this operation
    protocol := exec.Command("cardano-cli","query","protocol-parameters","--mainnet","--out-file","protocol.json")
    protocol.Run()

    //calculating slot
    slotCmd,_ := exec.Command("cardano-cli","query","tip",
    "--mainnet").Output()
    output := string(slotCmd[:])
    a := regexp.MustCompile("[^\\s]+")
    array:=a.FindAllString(output,-1)
    slotAux1:=strings.Trim(array[6],",")
    slot,_:=strconv.Atoi(slotAux1)
    slot=slot+10000

    //initialize multithread sending
    wg.Add(size)
    for i := 0; i < size; i++ {
        go func(i int) {

            tx_in:=txsIn[i]
            totalValue:=totalValues[i]

            //creating name files
            tx_tmp:="tx"+strconv.Itoa(i+1)+".tmp"
            tx_raw:="tx"+strconv.Itoa(i+1)+".raw"
            tx_signed:="tx"+strconv.Itoa(i+1)+".signed"

            //wallets config
            wallet:="wallet"+strconv.Itoa(i+1)
            walletDir:=dir+wallet
            walletPaymentAddress:=walletDir+"/"+wallet+".payment.addr"
            walletAddressAux,_:=exec.Command("cat",walletPaymentAddress).Output()
            walletAddress := string(walletAddressAux[:])

            //buildRaw 1
            walletAddressRaw1:=walletAddress+"+0"
            destinationAddressRaw1:=destinationAddress+"+0"
            buildRaw := exec.Command("cardano-cli","transaction","build-raw",
            "--tx-in", tx_in,
            "--tx-out" ,walletAddressRaw1,
            "--tx-out" ,destinationAddressRaw1,
            "--invalid-hereafter",strconv.Itoa(slot),
            "--fee" ,"0",
            "--out-file",tx_tmp)
            buildRaw.Run()
            
            //fee
            feeExec,_ := exec.Command("cardano-cli","transaction","calculate-min-fee",
            "--tx-body-file",tx_tmp,
            "--tx-in-count","1",
            "--tx-out-count","2",
            "--mainnet",
            "--witness-count","1",
            "--byron-witness-count","0",
            "--protocol-params-file","protocol.json").Output()
            feeAux:=string(feeExec[:])
            array=strings.Split(feeAux," ")
            fee:=array[0]

            //txOut
            totalValueInt,_:=strconv.Atoi(totalValue)
            feeInt,_:=strconv.Atoi(fee)
            txOut:=totalValueInt-feeInt-amountToSend
 
            //buildRaw 2
            walletAddressRaw2:=walletAddress+"+"+strconv.Itoa(txOut)//concat sender address + txout
            destinationAddressRaw2:=destinationAddress+"+"+strconv.Itoa(amountToSend)//concat receiver address + amountToSend
            buildRawFinal := exec.Command("cardano-cli","transaction","build-raw",
            "--tx-in",tx_in,
            "--tx-out",walletAddressRaw2,
            "--tx-out",destinationAddressRaw2,
            "--invalid-hereafter",strconv.Itoa(slot),
            "--fee",fee,
            "--out-file",tx_raw)
            buildRawFinal.Run()

            //tx sign
            walletSkey:=walletDir+"/"+wallet+".payment.skey"
            txSign := exec.Command("cardano-cli","transaction","sign",
            "--tx-body-file",tx_raw,
            "--signing-key-file",walletSkey,
            "--mainnet",
            "--out-file",tx_signed)
            txSign.Run()

            //tx submit
            txSubmit:=exec.Command("cardano-cli","transaction","submit",
            "--tx-file",tx_signed,
            "--mainnet")
            mu.Lock()
            if(txSubmit.Run()!=nil){
                errorCount=errorCount+1
            }
            mu.Unlock()

            //delete tmp files
            delete := exec.Command("rm",tx_raw,tx_signed,tx_tmp)
            delete.Run()

            defer wg.Done()
        }(i)
    }
    wg.Wait()
   
    //execution time
    elapsed := time.Since(start)
    log.Printf("total submitted = %v", size-errorCount)
    log.Printf("error count = %v", errorCount)
    log.Printf("time took = %s", elapsed)

    //delete tmp files
    delete := exec.Command("rm","protocol.json")
    delete.Run()

}