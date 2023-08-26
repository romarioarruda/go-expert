package main

import ("fmt")

func main() {
	mySlice := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }

	fmt.Printf("length=%d, capacity=%d, %v\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("length=%d, capacity=%d, %v\n", len(mySlice[2:]), cap(mySlice), mySlice[2:])
	fmt.Printf("length=%d, capacity=%d, %v\n", len(mySlice[:2]), cap(mySlice), mySlice[:2])

	mySlice = append(mySlice, 11)
	fmt.Printf("length=%d, capacity=%d, %v\n", len(mySlice), cap(mySlice), mySlice)
}