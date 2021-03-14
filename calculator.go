// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strings"
)

// Add returns the sum of two or more numbers.
func Add(a, b float64, v ...float64) float64 {
	r := a + b
	for _, x := range v {
		r += x
	}
	return r
}

// Subtract returns the difference of two or more numbers.
func Subtract(a, b float64, v ...float64) float64 {
	r := a - b
	for _, x := range v {
		r -= x
	}
	return r
}

// Multiply returns the product of two or more numbers.
func Multiply(a, b float64, v ...float64) float64 {
	r := a * b
	for _, x := range v {
		r *= x
	}
	return r
}

// Divide returns the quotient of two or more numbers.
func Divide(a, b float64, v ...float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("divide-by-zero for Divide(%f, %f)", a, b)
	}

	r := a / b
	for i, x := range v {
		if x == 0 {
			return 0.0, fmt.Errorf("divide-by-zero for Divide(%f, %f), in iteration %v processing variadic float64", r, x, i)
		}
		r /= x
	}
	return r, nil
}

// Sqrt returns the square root of a number
func Sqrt(a float64) (float64, error) {
	if a == 0 {
		return 0.0, fmt.Errorf("cannot-get-square-root-of-a-negative-number for Sqrt(%f)", a)
	}

	return math.Sqrt(a), nil
}

func EvaluateExpression(e string) (float64, error) {
	var a, b float64
	var operator string

	strings.ReplaceAll(e, " ", "") // remove all spaces

	numFields, err := fmt.Sscanf(e, "%f%1s%f", &a, &operator, &b)
	if err != nil {
		return 0.0, fmt.Errorf("Expression() error while parsing %q: %v", e, err)
	}

	if numFields != 3 {
		return 0.0, fmt.Errorf("nable to parse expression %q", e)
	}

	switch operator {
	case "+":
		return Add(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		// Divide() inclludes its own error
		return Divide(a, b)
	// We should never get here because the function input would fail to parse.
	default:
		return 0.0, fmt.Errorf("unknown operator %s in expression %q", operator, e)
	}
}
