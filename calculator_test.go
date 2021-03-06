package calculator_test

import (
	"calculator"
	"testing"
)

// Note Fatal|FatalF are also useful methods,
// to stop further execution within the same test.

// Express test-case inputs and desired outputs
type testCase struct {
	a           float64
	b           float64
	want        float64
	description string
	errExpected bool
}

func TestAdd(t *testing.T) {
	// Define test cases
	//
	// Reference RE: `cases := []testCases` vs. `var cases ...`
	// ways to define this variable: https://blog.golang.org/slices-intro
	testCases := []testCase{
		{
			description: "two positive numbers which sum to a positive",
			a:           2,
			b:           2,
			want:        4,
		},
		{
			description: "a positive and negative number which sum to a positive",
			a:           7,
			b:           -2,
			want:        5,
		},
		{
			description: "a positive and negative number which sum to a negative",
			a:           3,
			b:           -5,
			want:        -2,
		},
	}

	t.Parallel()

	// Fun fact: I initially called this variable `case` which is a keyword.
	// The error did not make that obvious. Fun rabbit-hole.
	for _, c := range testCases {
		got := calculator.Add(c.a, c.b)
		if got != c.want {
			t.Errorf("want %v, got %v, while testing %s. The function call was: Add(%v, %v)", c.want, got, c.description, c.a, c.b)
		}
	}
}

func TestSubtract(t *testing.T) {
	// Define test cases
	testCases := []testCase{
		{
			description: "two positive numbers whos difference is negative",
			a:           2,
			b:           9,
			want:        -7,
		},
		{
			description: "two positive numbers whos difference is positive",
			a:           7,
			b:           2,
			want:        5,
		},
		{
			description: "one positive and one negative decimal number whos difference is a positive decimal",
			a:           3,
			b:           -2.5,
			want:        5.5,
		},
	}

	t.Parallel()

	for _, c := range testCases {
		got := calculator.Subtract(c.a, c.b)
		if got != c.want {
			t.Errorf("want %v, got %v, while testing %s. The function call was: Subtract(%v, %v)", c.want, got, c.description, c.a, c.b)
		}
	}
}
func TestMultiply(t *testing.T) {
	// Define test cases
	testCases := []testCase{
		{
			description: "two positive numbers whos product is positive",
			a:           2,
			b:           20,
			want:        40,
		},
		{
			description: " a positive and negative number whos product is negative",
			a:           7,
			b:           -2,
			want:        -14,
		},
		{
			description: "a positive decimal and negative decimal whos product is a negative decimal",
			a:           8.4,
			b:           -2.5,
			want:        -21,
		},
	}

	t.Parallel()
	for _, c := range testCases {
		got := calculator.Multiply(c.a, c.b)
		if got != c.want {
			t.Errorf("want %v, got %v, while testing %s. The function call was: Multiply(%v, %v)", c.want, got, c.description, c.a, c.b)
		}
	}
}

func TestDivide(t *testing.T) {
	// Define test cases
	testCases := []testCase{
		{
			description: "dividing by zero",
			a:           2,
			b:           0,
			want:        123456789,
			errExpected: true,
		},
		{
			description: "two positive numbers whos quotient is positive",
			a:           20,
			b:           2,
			want:        10,
		},
		{
			description: " a positive and negative number whos quotient is negative",
			a:           10,
			b:           -2,
			want:        -5,
		},
		{
			description: "a positive decimal and negative decimal whos quotient is a negative decimal",
			a:           8.4,
			b:           -2.5,
			want:        -3.3600000000000003,
		},
	}

	t.Parallel()
	for _, c := range testCases {
		err, got := calculator.Divide(c.a, c.b)
		if err != nil && c.errExpected == false {
			t.Errorf("error received while testing %s. The function call was: Divide(%v, %v), and the error was: %v", c.description, c.a, c.b, err)
		}

		// Only fail on got != want if an error was not expected
		if c.errExpected == false && got != c.want {
			t.Errorf("want %v, got %v, while testing %s. The function call was: Divide(%v, %v)", c.want, got, c.description, c.a, c.b)
		}
	}
}
