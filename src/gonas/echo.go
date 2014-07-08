package gonas

import "net"

type EchoServer struct {
    AbstractServer
}

func (srv *EchoServer) Name() string {
    return "Echo"
}

func (srv *EchoServer) Handle(conn net.Conn) error {
    // calls the abstract handle operation that is going
    // to start the structures properly and check for error
    err := srv.AbstractServer.Handle(conn)
    if err != nil {
        return err
    }

    // defers the cleanup operation for the current connection
    // that is being handled, these will run the various tasks
    // associated with the cleaning of the connection/handling
    defer srv.AbstractServer.Cleanup(conn)

    // allocates space for the buffer that will
    // hold the data comming from the connection
    buff := make([]byte, 4096)

    for {
        // reads some data from the connection until the
        // and end of stream is found, in case there's an
        // error the function returns immediately
        count, err := conn.Read(buff)
        if err != nil {
            return err
        }

        // writes the received data back to the client as
        // an echo response based and then in case there's
        // an error returns it to the caller function
        count, err = conn.Write(buff[:count])
        if err != nil {
            return err
        }
    }
}
