package main

import (
	"fmt"
)

type Student struct {
	id int
	name string
	length float32
	isGraduate bool
}

func main() {
	studentMap := map[string]interface{}{
		"id": 39,
		"name": "ihooni",
		"length": 186.5,
		"isGraduate": false,
	}

	fmt.Println(studentMap)
	fmt.Println(studentMap["id"])
	fmt.Println(studentMap["name"])
	fmt.Println(studentMap["length"])
	fmt.Println(studentMap["isGraduate"])

	studentStruct := Student{
		id: 39,
		name: "ihooni",
		length: 186.5,
		isGraduate: false,
	}

	fmt.Println(studentStruct)
	fmt.Println(studentStruct.id)
	fmt.Println(studentStruct.name)
	fmt.Println(studentStruct.length)
	fmt.Println(studentStruct.isGraduate)
}
