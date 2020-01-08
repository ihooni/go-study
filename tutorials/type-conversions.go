package main

import "fmt"

func main() {
	var x float64 = -43.7
	var y int = int(x)
	var z uint = uint(y)

	fmt.Println(x, y, z)
}
