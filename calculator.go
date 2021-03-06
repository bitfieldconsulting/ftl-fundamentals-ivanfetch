// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them together.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the second
// from the first.
func Divide(a, b float64) (error, float64) {
	if b == 0 {
		return errors.New(fmt.Sprintf("divide-by-zero for Divide(%f, %f)", a, b)), 0.0
	}

	return nil, a / b
}
