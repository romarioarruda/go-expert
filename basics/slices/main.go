package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("slice value => %v\n", slice)
	fmt.Printf("slice cap => %v\n", cap(slice))

	slice = append(slice, 11)

	fmt.Printf("slice value => %v\n", slice)
	fmt.Printf("slice capacity => %v\n", cap(slice))

	for index, value := range slice {
		fmt.Printf("Index %v valor %v\n", index, value)
	}
}
