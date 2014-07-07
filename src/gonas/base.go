package gonas

import "fmt"
import "net"
import "sync"

func echo(conn net.Conn, wg sync.WaitGroup) {
	
	
	msg := make([]byte, 1000)
	
	conn.Read()
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
