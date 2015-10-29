package main

import "fmt"

func main() {
	m := map[string]string{"a": "A", "b": "B", "c": "C"}
	for k := range m {
		fmt.Println(m[k])
	}
}
