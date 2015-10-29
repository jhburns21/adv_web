package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter your favorite number: ")
	fmt.Scanln(&x)
	x = x * 2
	fmt.Println("My favorite number is twice as cool. It is number ", x)
}
