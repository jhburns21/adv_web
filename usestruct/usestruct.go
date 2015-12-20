package main

import "fmt"

func main() {
    type customer struct {
        name string
        age int
    }
    c1 := customer{"John", 25}
    c2 := customer{"Somedude", 99}

    fmt.Println(c1.name)
    fmt.Println(c2.age)

    c2.age = 29
    fmt.Println(c2.age)
}