// File: go/calc.go
package go_calc

// Add returns the sum of two float64 numbers.
func Add(a, b float64) float64 { return a + b }

// Sub returns the difference of two float64 numbers.
func Sub(a, b float64) float64 { return a - b } // <-- Typo 'а' is fixed to 'a'

// Mul returns the product of two float64 numbers.
func Mul(a, b float64) float64 { return a * b }

// Div returns the division of two float64 numbers.
// It will panic if the divisor is 0.
func Div(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
