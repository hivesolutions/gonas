package main

import "fmt"
import "github.com/hivesolutions/gonas/src/gonas"

func main() {
    srv := gonas.EchoServer{}
    err := gonas.Serve(&srv)
    if err != nil {
        fmt.Print(err)
    }
}
