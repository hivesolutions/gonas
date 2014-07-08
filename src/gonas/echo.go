package gonas

import "net"

type EchoServer struct {
    counter int
}

func (srv *EchoServer) name() string {
    return "Echo"
}

func (srv *EchoServer) count() int {
    return srv.counter
}

func (srv *EchoServer) handle(conn net.Conn) error {
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
        if err != nil {
            return err
        }

        count, err = conn.Write(msg[:count])
        if err != nil {
            return err
        }
    }
}
