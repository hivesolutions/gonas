package main

import "fmt"
import "gonas"

func main() {
    err := gonas.Serve()
    if err != nil {
        fmt.Print(err)
    }
}
