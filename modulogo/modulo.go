package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Choose a number: ")
	fmt.Scan(&x)
	fmt.Print("Choose a Divisor: ")
	fmt.Scan(&y)

	fmt.Print("The Remainder is ")
	fmt.Println(x % y)
}
