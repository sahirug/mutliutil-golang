// Package numberutil contains numerical utility functions
package numberutil

// Factorial returns factorial of received integer
// another comment
func Factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * Factorial(num-1)
}
