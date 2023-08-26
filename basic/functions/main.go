package main

import (
	"fmt"
	"errors"
)

func main() {
	age, err := sumAge(11, 10)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Returned age:", age)

	totalSum := sumAlotOfNumbers([]int{1, 2, 3, 4}...)

	fmt.Println("Total sum:", totalSum)

	totalFromClosure := closure()

	fmt.Println("totalFromClosure: ", totalFromClosure)

	totalFromClosure2 := func() int {
		return sumAlotOfNumbers(1, 2, 3, 4, 5, 10) * 2
	}()

	fmt.Println("totalFromClosure2: ", totalFromClosure2)
}

func sumAge(a, b int) (int, error) {
	age := a + b
	if age > 120 {
		return 0, errors.New("This is a really old guy")
	}

	return age, nil
}

func sumAlotOfNumbers(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func closure() int {
	return sumAlotOfNumbers(1, 2, 3) * 2
}
