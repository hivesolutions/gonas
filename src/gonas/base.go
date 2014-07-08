package gonas

import "fmt"
import "net"

type Server interface {
	Name() string
    Count() int
    Handle(conn net.Conn) error
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
