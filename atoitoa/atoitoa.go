package main

import "fmt"
import "strconv"

func main() {
    c := "243"
    fmt.Println(c)
    i, _:= strconv.Atoi(c)
    i = i + 23
    fmt.Println(i)
    c = strconv.Itoa(i)
    fmt.Println(c[0:2])
}