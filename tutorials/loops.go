package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("number", i)
	}
	
	my_map := map[string]string{
		"name": "ihooni",
		"university": "CAU",
	}

	for key, value := range my_map {
		fmt.Println(key, value)
	}
}
