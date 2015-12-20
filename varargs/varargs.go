package main

import "fmt"

func main() {
    ls := []int{1,6,3,5,8,9,23,52,79,88,99,34}
    l := findmax(ls...)
    fmt.Println(l)
}

func findmax(ns ...int) int {
    l := 0
    for _, n := range(ns) {
        if n > l {
            l = n
        }
    }
    return l
}