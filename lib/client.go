package lib

import (
    "net"
    "fmt"
    //"bufio"
    //"os"
    //"strings"
    "encoding/gob"

	"fyne.io/fyne/v2/data/binding"

)

type Client struct {
    serverConn net.Conn
    clientRunning bool
    name string

    messageData binding.UntypedList
}

func NewClient(serverAddr, name string) (*Client, error) {
    client := Client{clientRunning: false, name: name}

    connection, err := net.Dial("tcp", serverAddr)

    if err != nil {
        return nil, err
    }

    client.serverConn = connection
    return &client, nil
}

func (c *Client) SetMessageData(messageData binding.UntypedList) {
    c.messageData = messageData
}


/*
func (c *Client) Start() {
    fmt.Println("Client is running")
    go c.Listen()
    defer c.serverConn.Close()

    c.clientRunning = true

    encoder := gob.NewEncoder(c.serverConn)


    for c.clientRunning {
        //fmt.Print(c.name + ": ")
        newMessage, err := bufio.NewReader(os.Stdin).ReadString('\n')

        if err != nil {
            fmt.Println("Client io error ", newMessage)
            continue
        }

        trimedMessage := strings.Trim(newMessage, "\n")

        if trimedMessage == "exit" {
            os.Exit(3)

        }

        chatData := MakeChatData(c.name, newMessage)
        err = encoder.Encode(chatData)

        if err != nil {
            fmt.Println("client encoder error " , err)
        }





    }

}
*/


func (c* Client) Listen() {
    c.clientRunning = true

    for c.clientRunning {
        decoder := gob.NewDecoder(c.serverConn)

        var chatData ChatData
        err := decoder.Decode(&chatData)

        if err != nil {
            fmt.Println("Client listen error ", err)
            c.clientRunning = false

        }

        newMessage := chatData.String()
        fmt.Println(newMessage)
        c.messageData.Append(newMessage)

    }
}

func (c *Client) Send(message string) {
    encoder := gob.NewEncoder(c.serverConn)

    chatData := MakeChatData(c.name, message)
    err := encoder.Encode(&chatData)

    if err != nil {
        fmt.Println("client encoder error " , err)
    }

}

func (c *Client) Close() {
    c.serverConn.Close()
}
