package main

import "fmt"
import "gonas"

func main() {
    srv := gonas.HTTPHello{}
    err := gonas.Serve(&srv)
    if err != nil {
        fmt.Print(err)
    }
}
