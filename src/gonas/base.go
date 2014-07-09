package gonas

import "os"
import "net"

import "io/ioutil"

type Server interface {
    Logger

    // Retrieves the name as a string that represents
    // the server for the structure, this should be able
    // to address the comprehension of an end user.
    Name() string

    // Runs the intialization process for the server
    // starting its own infra-strcture correctly.
    Init() error

    // Returns the number of connection that have been
    // already handled by the current server.
    Count() int

    // Handles a connection by the current server, this should
    // be called as a goroutine for parallel processing, proper
    // thread locking mechanisms should be used.
    Handle(conn net.Conn) error
}

type AbstractServer struct {
    AbstractLogger
    counter int
}

func (srv *AbstractServer) Count() int {
    return srv.counter
}

func (srv *AbstractServer) Handle(conn net.Conn) error {
    srv.counter++
    return nil
}

func (srv *AbstractServer) Cleanup(conn net.Conn) error {
    err := conn.Close()
    if err != nil {
        return err
    }
    srv.Trace("Connection closed")
    return nil
}

func (srv *AbstractServer) Init() error {
    srv.InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
    return nil
}

func Serve(srv Server) error {
    srv.Init()

    // prints an information message about the starting
    // of gonas and the name of the server to be used
    srv.Infof("Starting gonas %s main loop ...\n", srv.Name())

    // creates the listener on the usual host and port
    // and using a tcp listening on all protocols
    ln, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        return err
    }

    // defers the closing of the listener to the end of the
    // function so that it does not remaing open (garbage)
    defer ln.Close()

    // prints a message about the starting of the connections
    // acceptance for the current server
    srv.Tracef("Accepting %s connections ...\n", srv.Name())

    // runs the service main loop accepting new connections and
    // redirecting its handling to the proper server
    for {
        conn, err := ln.Accept()
        if err != nil {
            return err
        }
        srv.Tracef("Accepted new connection %d\n", srv.Count())
        srv.Handle(conn)
    }

    // returns an invalid value as no error has occurred for the
    // current service handling infra-structure (everything ok)
    return nil
}
