package main


import "fmt"

func main() {
    f := 5.23478
    var i interface{} = 8.6457
    _, ok := i.(float64)
    if ok {
        fmt.Println(f+ i.(float64))
    } else {
        fmt.Println("not an int")
    }
}