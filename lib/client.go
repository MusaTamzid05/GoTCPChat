package lib

import (
    "net"
    "fmt"
    "bufio"
    "os"

)

type Client struct {
    serverConn net.Conn
}

func NewClient(serverAddr string) (*Client, error) {
    client := Client{}

    connection, err := net.Dial("tcp", serverAddr)

    if err != nil {
        return nil, err
    }

    client.serverConn = connection
    return &client, nil
}

func (c *Client) Start() {
    fmt.Println("Client is running")

    clientRunning := true

    for clientRunning {
        newMessage, err := bufio.NewReader(os.Stdin).ReadString('\n')

        if err != nil {
            fmt.Println("Client io error ", newMessage)
            continue
        }

        c.serverConn.Write([]byte(newMessage))




    }

}