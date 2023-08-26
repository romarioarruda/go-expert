package main

import (
	"fmt"
)

func main() {
	salary := make(map[string]int)

	salary["Romario"] = 10000
	salary["Other"] = 1000

	if salary["Romario"] >= 10000 {
		fmt.Printf("Realy good compensation: %v\n", salary["Romario"])
	}

	switch salary["Other"] {
		case 10000:
			println("Realy good compensation")
		case 5000:
			println("Good compensation")
		case 1000:
			println("Bad compensation")
		
	}	
}