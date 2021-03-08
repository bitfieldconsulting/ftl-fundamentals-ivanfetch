// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64, v ...float64) float64 {
	r := a + b
	for _, x := range v {
		r += x
	}
	return r
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64, v ...float64) float64 {
	r := a - b
	for _, x := range v {
		r -= x
	}
	return r
}

// Multiply takes two numbers and returns the result of multiplying them together.
func Multiply(a, b float64, v ...float64) float64 {
	r := a * b
	for _, x := range v {
		r *= x
	}
	return r
}

// Divide takes two numbers and returns the result of dividing the second
// from the first.
func Divide(a, b float64, v ...float64) (error, float64) {
	if b == 0 {
		return fmt.Errorf("divide-by-zero for Divide(%f, %f)", a, b), 0.0
	}

	r := a / b
	for i, x := range v {
		if x == 0 {
			return fmt.Errorf("divide-by-zero for Divide(%f, %f), in iteration %v processing variadic float64", r, x, i), 0.0
		}
		r /= x
	}
	return nil, r
}

// Sqrt returns the square root of a number
func Sqrt(a float64) (error, float64) {
	if a == 0 {
		return fmt.Errorf("cannot-get-square-root-of-a-negative-number for Sqrt(%f)", a), 0.0
	}

	return nil, math.Sqrt(a)
}

func Expression(e string) (error, float64) {
	// This regular expression matches:
	//   (\d+\.?\d*) - one or more digits, optionally with a decimal and more digits
	//   \s* - 0 or more space
	//   ([\+\-\*\/]) - one of the operators: + - * or /
	//   \s* - 0 or more space
	//   (\d+\.?\d*) - one or more digits, optionally with a decimal and more digits
	re := regexp.MustCompile(`^(\d+\.?\d*)\s*([\+\-\*\/])\s*(\d+\.?\d*)`)
	fields := re.FindStringSubmatch(e)
	// THe fields include the original string,
	// so len(fields) is off-by-one.
	if len(fields) < 3 {
		return fmt.Errorf("nable to parse expression \"%s\"", e), 0.0
	}

	a, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return fmt.Errorf("unable to parse \"%s\" to a float64 in expression \"%s\"", fields[1], e), 0.0
	}

	b, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return fmt.Errorf("unable to parse \"%s\" to a float64 in expression \"%s\"", fields[3], e), 0.0
	}

	var operator string = fields[2]

	switch operator {
	case "+":
		return nil, Add(a, b)
	case "-":
		return nil, Subtract(a, b)
	case "*":
		return nil, Multiply(a, b)
	case "/":
		// Divide() inclludes its own error, although, should we wrap this?
		return Divide(a, b)
	// We should never get here because the regular expression would fail to parse with an invalid operator.
	default:
		return fmt.Errorf("unknown operator %s in expression \"%s\"", operator, e), 0.0

	}
}
