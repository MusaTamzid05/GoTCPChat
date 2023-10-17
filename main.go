package main

import (
    "fmt"
    "net1_recording/lib"
    "os"
    "flag"
)

func main() {
    addr := ":8080"

    serverFlagPtr := flag.Bool("server", true, "Flag for server")
    clientNameFlagPtr := flag.String("client", "", "The client user name")
    flag.Parse()

    serverFlag := *serverFlagPtr

    if serverFlag {

        server  , err := lib.NewServer(addr)

        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        server.Start()
        return

    } 

    if *clientNameFlagPtr == "" {
        fmt.Println("Client needs a client name.")
        os.Exit(4)
    }



    client, err := lib.NewClient(addr, *clientNameFlagPtr)


    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    client.Start()

}

