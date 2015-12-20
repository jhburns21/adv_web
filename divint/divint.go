package main

import "fmt"

func main() {
    num, e := divInt(5)
    fmt.Println("(", num, ", ", e, ")" )
}

func divInt(i int) (int, bool){
    isEven := i%2

    return (i/2), (isEven != 0)  
}