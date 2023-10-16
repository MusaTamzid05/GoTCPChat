package lib

import (
    "net"
    "fmt"
)

type Server struct {
    addr string
    Listener net.Listener

}

func NewServer(addr string) (*Server, error) {
    server := Server{}

    addr = ":" + addr
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

        go s.handleClient(connection)


    }
}

func (s *Server) handleClient(conn net.Conn) {

    fmt.Println("Handle client")


}
