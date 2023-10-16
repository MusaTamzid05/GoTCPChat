package lib

import (
    "net"
    "fmt"
    "bufio"
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

    for clientRunning {

        clientMessage , err := bufio.NewReader(conn).ReadString('\n')

        if err != nil {

            if err.Error() == "EOF" {
                fmt.Println("Connection close")
            } else {
                fmt.Println(err)
            }

            clientRunning = false
            continue

        }

        fmt.Print(clientMessage)

        for index, client := range s.clients {

            if index == clientIndex {
                continue
            }


            broadcastMessage := fmt.Sprintf("client-%d - %s", clientIndex, clientMessage)
            _, err = client.Write([]byte(broadcastMessage))

            if err != nil {
                fmt.Println("Broadcast error " , err)
            } else {
                fmt.Println("Broadcast success")

            }




        }

    }


}
