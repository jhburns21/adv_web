package main

import "fmt"

func main() {
	s := "This is a string"
	bytestring := []byte(s)
	var r rune
	r = rune(bytestring[3])
	fmt.Println(r)
}
