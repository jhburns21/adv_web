package main

import "fmt"

func main() {
    myMap := map[int]string {
        0: "eggs",
        1: "bacon",
        2: "toast",
    }

    myMap[3] = "waffles"
    myMap[0] = "eggs benedict"
    delete(myMap, 2)
    fmt.Println(len(myMap))

    for key, v := range myMap {
        fmt.Println(key, ": ", v)
    }

    if _, ex := myMap[3]; ex {
        fmt.Println(myMap[3])
    } else {
        fmt.Println("Does Not Exist!")
    }
}