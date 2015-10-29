package main

import "fmt"

func main() {
	var a *int
	var x int = 56

	a = &x
	fmt.Println(a)
	fmt.Println(*a)
}
