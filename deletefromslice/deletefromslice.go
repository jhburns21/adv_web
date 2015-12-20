package main

import "fmt"

func main() {
    ms := []int{1,2,3,4,5,6,7,8,9}
    ms = append(ms[:4], ms[5:]...)
    fmt.Println(ms)
}