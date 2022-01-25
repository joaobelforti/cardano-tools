package main

import (
    "fmt"
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
  
    size, _ := strconv.Atoi(os.Args[1])
    amountToSend,_:=strconv.Atoi(os.Args[2])
    amountToSend=amountToSend*1000000
    destinationAddress:=os.Args[3]

    dir:="priv/wallet/"

    //protocol
    protocol := exec.Command("cardano-cli","query","protocol-parameters","--mainnet","--out-file","protocol.json")
    protocol.Run()

    //slot
    cmd,_ := exec.Command("node","getSlot.js").Output()
    output := string(cmd[:])
    slot1:=strings.Trim(output,"\n")
    slot2,_ := strconv.Atoi(slot1)
    slot3:=slot2+10000

    wg.Add(size)
    for i := 0; i < size; i++ {
        go func(i int) {
            time.Sleep(time.Millisecond * time.Duration(100*i))

            //wallets config
            wallet:="wallet"+strconv.Itoa(i+1)
            walletDir:="../"+dir+wallet
            walletPaymentAddress:=walletDir+"/"+wallet+".payment.addr"
            walletAddressAux,_:=exec.Command("cat",walletPaymentAddress).Output()
            walletAddress := string(walletAddressAux[:])

            //tx_in
            cmd,_ = exec.Command("cardano-cli","query","utxo",
            "--address",walletAddress,
            "--mainnet").Output()
             output = string(cmd[:])
             a := regexp.MustCompile("[^\\s]+")
             array:=a.FindAllString(output,-1)
             //fmt.Println(len(array))
             //fmt.Println(array)
            tx_in:=array[4]+"#"+array[5]
            totalValue:=array[6]

            //buildRaw 1
            walletAddressRaw1:=walletAddress+"+0"
            destinationAddressRaw1:=destinationAddress+"+0"

            buildRaw := exec.Command("cardano-cli","transaction","build-raw",
            "--tx-in", tx_in,
            "--tx-out" ,walletAddressRaw1,
            "--tx-out" ,destinationAddressRaw1,
            "--invalid-hereafter",strconv.Itoa(slot3),
            "--fee" ,"0",
            "--out-file","tx.tmp")

            buildRaw.Run()
            
            //fee
            feeExec,_ := exec.Command("cardano-cli","transaction","calculate-min-fee",
            "--tx-body-file","tx.tmp",
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
            //fmt.Println(walletAddressRaw2)
            //fmt.Println(destinationAddressRaw2)
            buildRawFinal := exec.Command("cardano-cli","transaction","build-raw",
            "--tx-in",tx_in,
            "--tx-out",walletAddressRaw2,
            "--tx-out",destinationAddressRaw2,
            "--invalid-hereafter",strconv.Itoa(slot3),
            "--fee",fee,
            "--out-file","tx.raw")
            buildRawFinal.Run()

            //tx sign
            walletSkey:=walletDir+"/"+wallet+".payment.skey"
        
            txSign := exec.Command("cardano-cli","transaction","sign",
            "--tx-body-file","tx.raw",
            "--signing-key-file",walletSkey,
            "--mainnet",
            "--out-file","tx.signed")
            txSign.Run()

            //tx submit
            txSubmit:=exec.Command("cardano-cli","transaction","submit",
            "--tx-file","tx.signed",
            "--mainnet")
            txSubmit.Run();
            defer wg.Done()
        }(i)
    }
    wg.Wait()
   
    //execution time
    elapsed := time.Since(start)
    log.Printf("time took = %s", elapsed)

    //delete temp files
    delete := exec.Command("rm","balance.out","fullUtxo.out","protocol.json","tx.raw","tx.signed","tx.tmp")
    errDelete := delete.Run()
    if errDelete == nil {
        fmt.Println("end")
    }
}