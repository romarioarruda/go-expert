package main

func main() {
	for i := 0; i < 5; i++ {
		println("i:", i+1)
	}

	println("=========================")

	people := []string{ "romario", "tati", "other" }

	for key, value := range people {
		println("key: ", key)
		println("value: ", value)
		println("get by key: ", people[key])
	}

	println("=========================")

	indexTest := 0
	for indexTest < 5 {
		println("Testing...", indexTest)
		indexTest++
	}
}