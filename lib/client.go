package lib

import (
    "net"
    "fmt"
    "bufio"
    "os"
    "strings"

)

type Client struct {
    serverConn net.Conn
    clientRunning bool
}

func NewClient(serverAddr string) (*Client, error) {
    client := Client{clientRunning: false}

    connection, err := net.Dial("tcp", serverAddr)

    if err != nil {
        return nil, err
    }

    client.serverConn = connection
    return &client, nil
}


func (c *Client) Start() {
    fmt.Println("Client is running")
    go c.Listen()
    defer c.serverConn.Close()

    c.clientRunning = true

    for c.clientRunning {
        newMessage, err := bufio.NewReader(os.Stdin).ReadString('\n')

        if err != nil {
            fmt.Println("Client io error ", newMessage)
            continue
        }

        trimedMessage := strings.Trim(newMessage, "\n")

        if trimedMessage == "exit" {
            os.Exit(3)

        }

        c.serverConn.Write([]byte(newMessage))




    }

}


func (c* Client) Listen() {

    for c.clientRunning {
        newMessage, err := bufio.NewReader(c.serverConn).ReadString('\n')

        if err != nil {
            fmt.Println(err)
            c.clientRunning = false

        }

        fmt.Println(newMessage)

    }

}
