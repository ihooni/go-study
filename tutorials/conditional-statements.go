package main

import "fmt"

func main() {
	age := 18
	if age > 18 {
		fmt.Println("you are old")
	} else {
		fmt.Println("you are young")
	}

	grade := "A"
	switch grade {
	case "A":
		fmt.Println("very good")
	case "B":
		fmt.Println("good")
	case "C":
		fmt.Println("hmm...")
	default:
		//
	}
}
