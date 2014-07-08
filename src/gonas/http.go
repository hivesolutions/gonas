package gonas

import "net"

type HTTPServer struct {
    AbstractServer
}

func (srv *HTTPServer) Name() string {
    return "HTTP"
}

func (srv *HTTPServer) Handle(conn net.Conn) error {
    srv.counter++
    defer conn.Close()
    msg := make([]byte, 4096)
    conn.Read(msg)
    conn.Write([]byte("HTTP/1.1 200 OK\r\nServer: gonas\r\n\r\nHello World"))
    return nil
}
