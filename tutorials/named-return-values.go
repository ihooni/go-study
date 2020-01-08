package main

import "fmt"

func splitString(str string) (a, b string) {
	a = str[:5]
	b = str[5:]

	return	// naked return
}

func main() {
	fmt.Println(splitString("hello world"))
}
