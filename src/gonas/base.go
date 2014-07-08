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
    srv.Infof("Starting gonas %s main loop ...\n", srv.Name())

    ln, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        return err
    }

    defer ln.Close()

    srv.Tracef("Accepting %s connections ...\n", srv.Name())

    for {
        conn, err := ln.Accept()
        if err != nil {
            return err
        }
        srv.Tracef("Accepted new connection %d\n", srv.Count())
        srv.Handle(conn)
    }

    return nil
}
