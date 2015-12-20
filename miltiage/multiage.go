package main

import "fmt"

func main() {
    fn := "john"
    dy, old := namefunc(fn, 27)
    if old {
        fmt.Println(fn, "is", dy, "in dog years and is old")
    } else {
        fmt.Println(fn, "is", dy, "in dog years and is not old")
    }
}

func namefunc(n string, a int) (int, bool) {
    dy := a*7
    var isold bool
    if a > 25 {
        isold = true
    } else {
        isold = false
    }
    return dy, isold
}