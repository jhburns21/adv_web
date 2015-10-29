package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 15
	fmt.Println("Print a Number: " + strconv.Itoa(x))
	s := "99"
	fmt.Print("String math: ")
	x, _ = strconv.Atoi(s)
	x = x + 1
	fmt.Println(x)
}
