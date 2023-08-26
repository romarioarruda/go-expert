package main

const a = "Hello world"

type User int

var (
	b int = 1
	c string = "Test"
	d float64 = 1.23
	f User = 23
)

func main() {
	e := "Short hand"
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}