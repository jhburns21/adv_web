package main


import "fmt"
import "reflect"

func main() {
    i := 5
    s := "lsakdhfalksdh"

    fmt.Println(reflect.TypeOf(i))
    fmt.Printf("Type is %T \n", s)
}