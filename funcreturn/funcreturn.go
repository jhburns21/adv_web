package main

import "fmt"

func runFunc(ses []string, callback func(string)) {
    for _, s := range ses {
        callback(s)
    }
}

func main() {
    runFunc([]string{"Hello World", "Sup?", "Yo!"}, func(s string){
            fmt.Println(s)
        })
}