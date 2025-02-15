package main

import "fmt"

const num = 3

func main() {
	constant()
	variables()
	arr()
	substring()
}

func constant() {
	fmt.Println("constant: ", num)
}

func variables() {
	var pi float64 = 3.14
	var name string = "ihooni"
	var string, number = "good", 35	// declaration with initializer
	a, b := true, false	// shorthand syntax (automatically assigned the correct type)

	fmt.Println("variables: ", pi, name, string, number, a, b)
}

func arr() {
	names := []string{ "chihoon", "ihooni" }

	fmt.Println("array: ", names)
}

func substring() {
	sentence := "Hello my name is ihooni!"

	fmt.Println(sentence[:10])
}
