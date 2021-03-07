// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a float64, b ...float64) float64 {
	r := a + b[0]

	// If b contains more than 1 element,
	// call ourself with the current result and those remaining elements.
	if len(b[1:]) > 0 {
		r = Add(r, b[1:]...)
	}
	return r
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a float64, b ...float64) float64 {
	r := a - b[0]

	// If b contains more than 1 element,
	// call ourself with the current result and those remaining elements.
	if len(b[1:]) > 0 {
		r = Subtract(r, b[1:]...)
	}
	return r
}

// Multiply takes two numbers and returns the result of multiplying them together.
func Multiply(a float64, b ...float64) float64 {
	r := a * b[0]

	// If b contains more than 1 element,
	// call ourself with the current result and those remaining elements.
	if len(b[1:]) > 0 {
		r = Multiply(r, b[1:]...)
	}
	return r
}

// Divide takes two numbers and returns the result of dividing the second
// from the first.
func Divide(a float64, b ...float64) (error, float64) {
	var err error // potential error from recursive call to Divide()

	if b[0] == 0 {
		return errors.New(fmt.Sprintf("divide-by-zero for Divide(%f, %f)", a, b[0])), 0.0
	}

	r := a / b[0]

	// If b contains more than 1 element,
	// call ourself with the current result and those remaining elements.
	if len(b[1:]) > 0 {
		err, r = Divide(r, b[1:]...)
		if err != nil {
			return err, r
		}
	}
	return nil, r
}

// Sqrt returns the square root of a number
func Sqrt(a float64) (error, float64) {
	if a == 0 {
		return errors.New(fmt.Sprintf("cannot-get-square-root-of-a-negative-number for Sqrt(%f)", a)), 0.0
	}

	return nil, math.Sqrt(a)
}
