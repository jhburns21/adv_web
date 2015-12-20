package main

import "fmt"

func main() {
    s := "Print 1"
    {
        s := "Print 2"
        fmt.Println(s)
    }
    fmt.Println(s)
}