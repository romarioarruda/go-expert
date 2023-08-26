package math

//Starting with capital letter means that Sum is automatically exported (same thing as module.export or export default in Node.js)
func Sum[T int | float64](a, b T) T {
	return a + b
}
