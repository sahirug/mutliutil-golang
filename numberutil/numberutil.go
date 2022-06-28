// Package numberutil contains numerical utility functions
package numberutil

// Factorial returns factorial of received integer
func Factorial(num int) int{
   if num == 1 || num == 0{
      return num
   }
   return num * factorial(num-1)
}
