package main

import (
    "fmt"
    "math"
)

func main() {
    var x float64

    fmt.Print("Pick a favorite floating point Number: ")
    fmt.Scan(&x)
    fmt.Println(math.Ceil(x))
}