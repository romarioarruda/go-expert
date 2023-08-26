package main

import ("fmt")

func main() {
	salary := make(map[string]int)

	salary["Romario"] = 10000

	for index, value := range salary {
		fmt.Printf("salary from %s is %d\n", index, value)
	}

	delete(salary, "Romario")

	fmt.Printf("salaries %v\n", salary)
}