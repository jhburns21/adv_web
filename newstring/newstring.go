package main

import "fmt"

func main() {
    var s *string = new(string)
    fmt.Println(s)
    fmt.Println(*s) 
}