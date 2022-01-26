package main

import (
    "fmt"
    "log"
    "sync"
    "os"
    "strconv"
    "os/exec"
    "time"
)

func main() {
    start := time.Now()
    var mu sync.Mutex
    var wg sync.WaitGroup
    size, err := strconv.Atoi(os.Args[1])
    errorCount :=0
    
    if(err!=nil){
        fmt.Println(err)
    }

    wg.Add(size)
    for i := 1; i <= size; i++ {
        go func(i int) {
            //time.Sleep(time.Millisecond * time.Duration(115*i*3))
            defer wg.Done()
            var wallet = strconv.Itoa(i)
            cmd := exec.Command("node","sendAll.js", wallet, "addr1qxwc8cxurrktaf7k8y50mm062dx2vm8rpsw409m694kc8k8nt267zvmmkhtpt3kala3ewnehhvtf2t4kgd98gpqcrxzq6432kh")
            mu.Lock()
            err := cmd.Run()
            mu.Unlock()
            if err != nil {
                errorCount = errorCount+1
            }
        }(i)
    }
    wg.Wait()

    elapsed := time.Since(start)
    log.Printf("error count = %v", errorCount)
    log.Printf("time took = %s", elapsed)
}