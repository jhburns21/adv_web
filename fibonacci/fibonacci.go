package main

import "fmt"

func main() {
    i := fib(7)
    fmt.Println(i)
}

func fib(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fib(n-1) + fib(n-2)
    }
}