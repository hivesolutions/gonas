package gonas

import "net"

type HTTPHello struct {
	counter int
}

func (srv HTTPHello) count() int {
	return srv.counter
}

func (srv HTTPHello) handle(conn net.Conn) error {
	srv.counter++
    defer conn.Close()
    msg := make([]byte, 4096)
    conn.Read(msg)
    conn.Write([]byte("HTTP/1.1 200 OK\r\nServer: gonas\r\n\r\nHello World"))
    return nil
}
