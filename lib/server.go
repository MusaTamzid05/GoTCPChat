package lib

import (
    "net"
    "fmt"
    //"bufio"
    "encoding/gob"
)

type Server struct {
    addr string
    Listener net.Listener
    clients []net.Conn

}

func NewServer(addr string) (*Server, error) {
    server := Server{}

    server.addr = addr

    listener , err := net.Listen("tcp", addr)

    if err != nil {
        return nil, err
    }

    server.Listener = listener
    return &server, nil

}

func (s *Server) Start() {
    fmt.Println("Server is running at ", s.addr)
    defer s.Listener.Close()

    serverRunning := true


    for serverRunning {
        connection, err := s.Listener.Accept()

        if err != nil {
            fmt.Println(err)
            continue
        }

        s.clients = append(s.clients, connection)
        go s.handleClient(connection, len(s.clients) - 1)


    }
}

func (s *Server) handleClient(conn net.Conn, clientIndex int) {
    clientRunning := true
    decoder := gob.NewDecoder(conn)


    for clientRunning {

        var chatData ChatData
        err := decoder.Decode(&chatData)


        if err != nil {

            if err.Error() == "EOF" {
                fmt.Println("Connection close")

                if len(s.clients) == 1 {
                    s.clients = []net.Conn{}
                } else {
                    s.clients = append(s.clients[:clientIndex], s.clients[clientIndex+1:]...)
                }

            } else {
                fmt.Println(err)
            }

            clientRunning = false
            continue

        }

        fmt.Print(chatData.String())

        for index, client := range s.clients {

            if index == clientIndex {
                continue
            }

            encoder := gob.NewEncoder(client)
            err = encoder.Encode(chatData)




            if err != nil {
                fmt.Println("Broadcast error " , err)
            } else {
                fmt.Println("Broadcast success")

            }




        }

    }


}
