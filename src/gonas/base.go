package gonas

import "fmt"
import "net"

type Server interface {
    // Retrieves the name as a string that represents
    // the server for the structure, this should be able
    // to address the comprehension of an end user.
    Name() string

    // Returns the number of connection that have been
    // already handled by the current server.
    Count() int

    // Handles a connection by the current server, this should
    // be called as a goroutine for parallel processing, proper
    // thread locking mechanisms should be used.
    Handle(conn net.Conn) error
}

type AbstractServer struct {
    counter int
}

func (srv *AbstractServer) Count() int {
    return srv.counter
}

func (srv *AbstractServer) Handle(conn net.Conn) error {
    srv.counter++
    return nil
}

func Serve(srv Server) error {
    fmt.Print("Starting gonas main loop\n")

    ln, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        return err
    }

    defer ln.Close()

    fmt.Printf("Accepting (%s) connections ...\n", srv.Name())

    for {
        conn, err := ln.Accept()
        if err != nil {
            return err
        }
        fmt.Printf("Accepted new connection %d\n", srv.Count())
        srv.Handle(conn)
    }

    return nil
}
