package main

import "fmt"

func main() {
	fmt.Println("First: ")
	s := []int{1, 4, 7, 9, 0, 6, 4, 3, 4, 5, 6}
	printNumbers(s)
	fmt.Println("Second: ")
	t := []int{2, 5, 4, 3, 5}
	printNumbers(t)
}

func printNumbers(nums []int) {
	for _, x := range nums {
		fmt.Println(x)
	}
}
