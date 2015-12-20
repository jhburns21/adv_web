package main

import "fmt"

func main() {
    s := "this is a string"
    f := 3.2178
    c := 'a'

    fmt.Println(s)
    b := []byte(s)
    fmt.Println(b)
    fmt.Println(string(b))

    fmt.Println()
    fmt.Println(f)
    i := int(f)
    fmt.Println(i)
    fmt.Println(float64(i))

    fmt.Println()
    fmt.Println(c)
    fmt.Println(string(c))
}