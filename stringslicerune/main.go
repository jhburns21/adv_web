package main

import "fmt"
import "unicode/utf8"

func main() {
    s := "This is a string"
    var r rune;
    r, _ = utf8.DecodeRuneInString(s) // get first rune in string
    fmt.Println(r)

    fmt.Println(s[2:6])
}