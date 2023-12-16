package main

import "fmt"

func main() {
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Printf("Array value => %v\n", array)
	fmt.Printf("Array tipo => %T\n", array)
	fmt.Printf("Array length => %v\n", len(array))

	for index, value := range array {
		fmt.Printf("Index %d => value %v\n", index, value)
	}
}
