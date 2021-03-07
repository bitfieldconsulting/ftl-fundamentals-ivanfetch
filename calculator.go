// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
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
		return errors.New(fmt.Sprintf("nable to parse expression \"%s\"", e)), 0.0
	}

	a, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to parse \"%s\" to a float64 in expression \"%s\"", fields[1], e)), 0.0
	}

	b, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to parse \"%s\" to a float64 in expression \"%s\"", fields[3], e)), 0.0
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
		return errors.New(fmt.Sprintf("unknown operator %s in expression \"%s\"", operator, e)), 0.0

	}
}
