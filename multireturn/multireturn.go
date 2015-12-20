package main

import "fmt"

func main() {
    namestring("john", 27)
    namestring("bob", 45)
}

func namestring(s string, i int) {
    fmt.Println(s, "is", i, "years old.")
}