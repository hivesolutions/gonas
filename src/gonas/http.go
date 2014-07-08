package gonas

import "net"

type HTTPServer struct {
    AbstractServer
}

func (srv *HTTPServer) Name() string {
    return "HTTP"
}

func (srv *HTTPServer) Handle(conn net.Conn) error {
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

    // allocates the proper space for the buffer usage, this is
    // going to be used allong the function for storage
    buff := make([]byte, 4096)

    // reads some data from the current connection and verifies
    // if there's an error if that's the case returns the error
    // to the caller function (as expected)
    _, err = conn.Read(buff)
    if err != nil {
        return err
    }

    // writes the hello message back to the client so that it may
    // be able to receive it correctly and display it to the end user
    // then verifies if an error was created by such write operation
    _, err = conn.Write([]byte("HTTP/1.1 200 OK\r\nServer: gonas\r\n\r\nHello World"))
    if err != nil {
        return err
    }

    // returns an invalid value as the error result for the current
    // function that handles a connection
    return nil
}
