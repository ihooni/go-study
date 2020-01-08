package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("number", i)
	}

	sum := 1
	for sum < 1000 {	// go's while loop
		sum += sum
		fmt.Println("go's while ", sum)
	}
	
	my_map := map[string]string{
		"name": "ihooni",
		"university": "CAU",
	}

	for key, value := range my_map {
		fmt.Println(key, value)
	}

	for {	// infinite loop

	}
}
