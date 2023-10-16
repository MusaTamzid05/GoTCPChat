package main

import (
    "fmt"
    "net1_recording/lib"
    "os"
)

func main() {
    addr := "8080"

    server  , err := lib.NewServer(addr)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    server.Start()

}

