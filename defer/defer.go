package main

import "fmt"

func main() {
    defer atEnd()
    fmt.Println("This will print first")
}

func atEnd() {
    fmt.Println("This prints last")
}