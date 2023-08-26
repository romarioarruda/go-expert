package main

import ("fmt")

func main() {
	var arrayTest [3]int

	arrayTest[0] = 10
	arrayTest[1] = 20
	arrayTest[2] = 30

	for index, value := range arrayTest {
		fmt.Printf("id: %d, value: %d\n", index, value)
	}
}