package gonas

import "io"
import "fmt"
import "net"
import "sync"

func echo(conn net.Conn, wg sync.WaitGroup) {
	// defers both the closing of the connection
	// and the marking of the wait group as done
	// (unblocks the other side of the channel)
    defer conn.Close()
    defer wg.Done()

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

func Serve() error {
    fmt.Print("Starting gonas main loop")

    var wg sync.WaitGroup

    ln, err := net.Listen("tcp", "0.0.0.0")
    if err != nil {
        return err
    }

    defer ln.Close()

    conn, err := ln.Accept()
    if err != nil {
        return err
    }

    wg.Add(1)
    go echo(conn, wg)

    wg.Wait()

    return nil
}
