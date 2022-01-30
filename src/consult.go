package main

import (
    "fmt"
    "log"
    "sync"
    "os"
    "strconv"
    "os/exec"
    "time"
    "strings"
    "regexp"
)

func main() {
    start := time.Now()
    var wg sync.WaitGroup
    size, err := strconv.Atoi(os.Args[1])
	arrayAddress := make([]string, size)
	arrayConsult := make([]string, size)
    arrayAssets := make([]string, size)
    if(err!=nil){
        fmt.Println(err)
    }

    wg.Add(size)
    for i := 0; i < size; i++ {
        go func(i int) {
            txs:=""
            assets:=""
            ada:=0
            //configure wallets
			dir:="priv/wallet/"
            wallet:="wallet"+strconv.Itoa(i+1)
            walletDir:=dir+wallet
            walletPaymentAddress:=walletDir+"/"+wallet+".payment.addr"
            walletAddressAux,_:=exec.Command("cat",walletPaymentAddress).Output()
            walletAddress := string(walletAddressAux[:])

			cmd,_ := exec.Command("cardano-cli","query","utxo",
            "--address",walletAddress,
            "--mainnet").Output()
            output:=string(cmd[:])
             
            a := regexp.MustCompile("[^\\s]+")
            array:=a.FindAllString(output,-1)

            //cleaning data for printing
            for x := 4; x <= len(array)-1; x++{
                verify:=strings.Index(array[x],".")
                if array[x]!="+"{
                    if array[x]=="lovelace"{
                        adaAux,_:=strconv.Atoi(array[x-1])
                        ada=adaAux+ada
                    }
                    if array[x]=="TxOutDatumNone"{
                        txs =txs+"\n"
                    }
                    if array[x]!="lovelace" && array[x]!="TxOutDatumNone" && array[x+1]!="lovelace" && array[x+2]!="lovelace"&& array[x-1]!="+"&& verify==-1{
                        txs = array[x]+" "+txs
                    }
                    if(array[x-1]=="+" && array[x]!="TxOutDatumNone"){
                        assets=array[x+1]+" "+assets
                    }
                }
            }
            //calc total ada of each wallet
            adaStr := strconv.Itoa(ada)
            txs = adaStr+" "+txs
            
			arrayAddress[i]=walletAddress
			arrayConsult[i]=txs
            arrayAssets[i]=assets
			defer wg.Done()
        }(i)
    }
    wg.Wait()

	for i := 0; i < size; i++{
		//time.Sleep(time.Millisecond * time.Duration(100))
		fmt.Println("wallet",i+1,"->","["+arrayAddress[i]+"]","")
        txs := regexp.MustCompile("[^\\s]+")
        //divide each array index in other array 
        consultArray:=txs.FindAllString(arrayConsult[i],-1)
    
        //fmt.Println(consultArray)
        for x := 0; x < len(consultArray); x++{
            if(x==0){
                ada,_:=strconv.Atoi(consultArray[0])
                ada=ada/1000000
                fmt.Println("total ADA ->",ada)
            }else{
                fmt.Println("tx",x,"->","["+consultArray[x]+"]")
            }
        }
        
        fmt.Println("assets","->",arrayAssets[i])
    
        fmt.Println("\n----------------------------------------------------------------------------------------\n")
	}
    elapsed := time.Since(start)
    log.Printf("time took = %s", elapsed)
}