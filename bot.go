package main

import (
    "log"
    "sync"
    "os"
    "strconv"
    "os/exec"
    "time"
    "regexp"
    "strings"
)

func main() {
    start := time.Now()
    var wg sync.WaitGroup
    errorCount:=0
    size, _ := strconv.Atoi(os.Args[1])
    amountToSend,_:=strconv.Atoi(os.Args[2])
    amountToSend=amountToSend*1000000
    destinationAddress:=os.Args[3]

    //protocol
    protocol := exec.Command("cardano-cli","query","protocol-parameters","--mainnet","--out-file","protocol.json")
    protocol.Run()

    //get slot
    slotCmd,_ := exec.Command("cardano-cli","query","tip",
    "--mainnet").Output()
    output := string(slotCmd[:])
    a := regexp.MustCompile("[^\\s]+")
    array:=a.FindAllString(output,-1)
    slotAux1:=strings.Trim(array[6],",")
    slot,_:=strconv.Atoi(slotAux1)
    

    //initialize multithread sending
    wg.Add(size)
    for i := 0; i < size; i++ {
        go func(i int) {
            //files
            tx_tmp:="tx"+strconv.Itoa(i+1)+".tmp"
            tx_raw:="tx"+strconv.Itoa(i+1)+".raw"
            tx_signed:="tx"+strconv.Itoa(i+1)+".signed"

            //calculating slot
            slot=slot+10000

            //some time between transactions, avoiding fails
            //time.Sleep(time.Millisecond * time.Duration(20*i))

            //wallets config
            dir:="priv/wallet/"
            wallet:="wallet"+strconv.Itoa(i+1)
            walletDir:=dir+wallet
            walletPaymentAddress:=walletDir+"/"+wallet+".payment.addr"
            walletAddressAux,_:=exec.Command("cat",walletPaymentAddress).Output()
            walletAddress := string(walletAddressAux[:])

            //tx_in
            cmd,_ := exec.Command("cardano-cli","query","utxo",
            "--address",walletAddress,
            "--mainnet").Output()
             output  = string(cmd[:])
             a = regexp.MustCompile("[^\\s]+")
             array=a.FindAllString(output,-1)
            tx_in:=array[4]+"#"+array[5]
            totalValue:=array[6]

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
 
            //build raw 2
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
            if(txSubmit.Run()!=nil){
                errorCount=errorCount+1
            }

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