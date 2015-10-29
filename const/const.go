package main

import (
	"fmt"
	"strconv"
)

const age int = 24
const (
	A = iota
	B
	C
	D
)

func main() {
	fmt.Println(A + B + C + D)
	fmt.Println("My age is " + strconv.Itoa(age))
}
