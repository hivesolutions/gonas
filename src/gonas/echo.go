package gonas

import "net"

type EchoServer struct {
    AbstractServer
}

func (srv *EchoServer) Name() string {
    return "Echo"
}

func (srv *EchoServer) Handle(conn net.Conn) error {
    // defers the closing of the current connection
    // to the end of this handling function, so that
    // no connection are left pending in the "wild"
    defer conn.Close()

    // allocates space fot the buffer that will
    // hold the data comming from the connection
    msg := make([]byte, 4096)

    for {
        // reads some data from the connection until the
        // and end of stream is found, in case there's an
        // error the function returns immeidately
        count, err := conn.Read(msg)
        if err != nil {
            return err
        }

		// writes the received data back to the client as
		// an echo response based and then in case there's
		// an error returns it to the caller function
        count, err = conn.Write(msg[:count])
        if err != nil {
            return err
        }
    }
}
