package main

import "fmt"

func main() {
    fmt.Println(dogyears(7))
}

func dogyears(y int) (dogYears int) {
    return y*7
}