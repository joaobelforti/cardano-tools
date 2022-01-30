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
    var wg sync.WaitGroup
    size, err := strconv.Atoi(os.Args[1])
    errorCount :=0
    
    if(err!=nil){
        fmt.Println(err)
    }

    wg.Add(size)
    for i := 1; i <= size; i++ {
        go func(i int) {
            time.Sleep(time.Millisecond * time.Duration(115*i*3))
            defer wg.Done()
            var wallet = strconv.Itoa(i)
            cmd := exec.Command("node","sendAll.js", wallet, os.Args[2])
            err := cmd.Run()
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