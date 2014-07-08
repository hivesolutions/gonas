package gonas

import "io"
import "fmt"
import "net"

type Server interface {
    count() int
    handle(conn net.Conn) error
}

func echo(conn net.Conn) {
    // defers the closing of the current connection
    // to the end of this handling function, so that
    // no connection are left pending in the "wild"
    defer conn.Close()

    // allocates space fot the buffer that will
    // hold the data comming from the connection
    msg := make([]byte, 4096)

    for {
        // reads some data from the connection until the
        // eon of file (eof) indicator is found
        count, err := conn.Read(msg)
        if err == io.EOF {
            fmt.Printf("SERVER: received EOF (%d bytes ignored)\n", count)
            return
        } else  if err != nil {
            fmt.Printf("ERROR: read\n")
            fmt.Print(err)
            return
        }
        fmt.Printf("SERVER: received %v bytes\n", count)

        count, err = conn.Write(msg[:count])
        if err != nil {
            fmt.Printf("ERROR: write\n")
            fmt.Print(err)
            return
        }
        fmt.Printf("SERVER: sent %v bytes\n", count)
    }
}

func Serve(srv Server) error {
    fmt.Print("Starting gonas main loop\n")

    ln, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        return err
    }

    defer ln.Close()

    fmt.Print("Accepting new connections ...\n")

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
