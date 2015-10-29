package main

import "fmt"

func main() {
	fmt.Println("Enter a String of at least 5 chars: ")
	var s string
	fmt.Scanln(&s)
	fmt.Println(s[5:])
	fmt.Println(s[:5])
}
