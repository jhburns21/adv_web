package main

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

func main() {
	fmt.Print("hello, world\n")
	ID, err := uuid.ParseHex("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
	fmt.Println(ID)
	if err != nil {
		fmt.Println(err)
	}

}
