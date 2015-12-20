package main

import "fmt"

func main() {
    s := "this is a string"
    b := []byte(s)
    
    fmt.Println(b)

    s2 := string(b)
    fmt.Println(s2)
}