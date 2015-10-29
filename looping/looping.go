package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}
	j := 10

	for j > 1 {
		fmt.Println(j)
		j--
	}

	for {
		fmt.Println("FOREVER")
		if j > 10 {
			break
		}
		j++
	}

	fmt.Println("dang")
}
