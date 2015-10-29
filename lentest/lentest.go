package main

import "fmt"

func main() {
	s := "this is a string"
	a := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 7, 65, 4, 4}
	fmt.Println("String Length: ", len(s))
	fmt.Println("Array Length: ", len(a))
}
