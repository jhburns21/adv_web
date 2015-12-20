package main

import "fmt"
import "reflect"

func main() {
    hw := func() {
        fmt.Println("Func Exp!")
    }
    hw()
    fmt.Println(reflect.TypeOf(hw))
}