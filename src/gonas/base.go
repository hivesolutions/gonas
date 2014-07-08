package gonas

import "fmt"
import "net"

type Server interface {
	name() string
    count() int
    handle(conn net.Conn) error
}

func Serve(srv Server) error {
    fmt.Print("Starting gonas main loop\n")

    ln, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        return err
    }

    defer ln.Close()

    fmt.Printf("Accepting (%s) connections ...\n", srv.name())

    for {
        conn, err := ln.Accept()
        if err != nil {
            return err
        }
        fmt.Printf("Accepted new connection %d\n", srv.count())
        srv.handle(conn)
    }

    return nil
}
