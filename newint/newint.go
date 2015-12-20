package main

import "fmt"

func main() {
    var num *int = new(int)
    fmt.Println(num)
    fmt.Println(*num) 
}