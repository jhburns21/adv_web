package main

import "fmt"

func main() {
    var b *bool = new(bool)
    fmt.Println(b)
    fmt.Println(*b) 
}