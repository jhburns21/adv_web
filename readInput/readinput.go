package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int

	fmt.Print("Pick a favorite Number: ")
	fmt.Scan(&num)
	fmt.Println(strconv.Itoa(num) + " is my favorite too!")
}
