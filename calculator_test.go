package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
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

// This repeats the above struct,
// but b is a slice-of-float64 instead.
// Re-using the same struct for variadic and non-variadic would be nice,
// but separate structs makes test more readable.
type variadicTestCase struct {
	a           float64
	b           []float64
	want        float64
	description string
	errExpected bool
}

// Test case for expressions (string parsed into numbers and an operator)
type expressionTestCase struct {
	e           string
	want        float64
	description string
	errExpected bool
}

// Return a slice of N random numbers from 0-500
func randFloat64Slice(n int) []float64 {
	rand.Seed(time.Now().UnixNano())
	// TODO: Should this be pre-allocated instead of using append()?
	var r []float64

	for i := 0; i < n; i++ {
		x := rand.Float64() * float64(rand.Intn(500))
		r = append(r, x)
	}
	return r
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

// Generate 100 random test-cases for Add()
func TestAddRandomly(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	// logs are viewable using `go test -v`
	t.Log("Beginning random test-cases for Add(). . .")

	for i := 0; i < 100; i++ {
		// rand.Float64() returns a number in 0.0-1.0
		// Use another randomly-generated int to vary the whole number.
		a := rand.Float64() * float64(rand.Intn(500))
		b := rand.Float64() * float64(rand.Intn(500))
		want := a + b
		t.Logf("Random test %d: Add(%v, %v), wants %v", i, a, b, want)
		got := calculator.Add(a, b)
		if got != want {
			t.Errorf("want %v, got %v, while testing randomly-generated cases. The function call was: Add(%v, %v)", want, got, a, b)
		}
	}

	t.Log("Completed random test-cases for Add(). . .")
}

// Randomly test the variadic call to Add()
func TestAddVariadicRandomly(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	// rand.Float64() returns a number in 0.0-1.0
	// Use another randomly-generated int to vary the whole number.
	a := rand.Float64() * float64(rand.Intn(500))
	// b will be assigned multiple random float64 in a slice
	b := randFloat64Slice(5)

	// Get our own sum to compare to Add()
	var want float64 = a
	for _, x := range b {
		want += x
	}

	got := calculator.Add(a, b...)
	t.Logf("Random variadic test: Add(%v, %v), wants %v, got %v", a, b, want, got)
	if got != want {
		t.Errorf("want %v, got %v, while testing a random variadic case. The function call was: Add(%v, %v)", want, got, a, b)
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

// Randomly test the variadic call to Subtract()
func TestSubtractVariadicRandomly(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	// rand.Float64() returns a number in 0.0-1.0
	// Use another randomly-generated int to vary the whole number.
	a := rand.Float64() * float64(rand.Intn(500))
	// b will be assigned multiple random float64 in a slice
	b := randFloat64Slice(5)

	// Get our own difference to compare to Subtract()
	var want float64 = a
	for _, x := range b {
		want -= x
	}

	got := calculator.Subtract(a, b...)
	t.Logf("Random variadic test: Subtract(%v, %v), wants %v, got %v", a, b, want, got)
	if got != want {
		t.Errorf("want %v, got %v, while testing a random variadic case. The function call was: Subtract(%v, %v)", want, got, a, b)
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

// Randomly test the variadic call to Multiply()
func TestMultiplyVariadicRandomly(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	// rand.Float64() returns a number in 0.0-1.0
	// Use another randomly-generated int to vary the whole number.
	a := rand.Float64() * float64(rand.Intn(500))
	// b will be assigned multiple random float64 in a slice
	b := randFloat64Slice(5)

	// Get our own product to compare to Multiply()
	var want float64 = a
	for _, x := range b {
		want = want * x
	}

	got := calculator.Multiply(a, b...)
	t.Logf("Random variadic test: Multiply(%v, %v), wants %v, got %v", a, b, want, got)
	if got != want {
		t.Errorf("want %v, got %v, while testing a random variadic case. The function call was: Multiply(%v, %v)", want, got, a, b)
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

// Randomly test the variadic call to Divide()
func TestDivideVariadicRandomly(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	// rand.Float64() returns a number in 0.0-1.0
	// Use another randomly-generated int to vary the whole number.
	a := rand.Float64() * float64(rand.Intn(500))
	// b will be assigned multiple random float64 in a slice
	b := randFloat64Slice(5)

	// Get our own quotient to compare to Divide()
	// Also change any randomly-generated 0.0s, to avoid dividing-by-zero!
	var want float64 = a
	for i, x := range b {
		if x == 0 {
			x += 0.1
			b[i] = x // change the slice
		}
		want = want / x
	}

	err, got := calculator.Divide(a, b...)
	if err != nil {
		t.Errorf("error received while testing a random variadic case. The function call was: Divide(%v, %v), and the error was: %v", a, b, err)
	}

	t.Logf("Random variadic test: Divide(%v, %v), wants %v, got %v", a, b, want, got)
	if got != want {
		t.Errorf("want %v, got %v, while testing a random variadic case. The function call was: Divide(%v, %v)", want, got, a, b)
	}
}

func TestSqrt(t *testing.T) {
	// Define test cases
	testCases := []testCase{
		{
			description: "negative input",
			a:           -64,
			want:        123456789,
			errExpected: true,
		},
		{
			description: "64",
			a:           64,
			want:        8,
		},
	}

	t.Parallel()
	for _, c := range testCases {
		err, got := calculator.Sqrt(c.a)
		if err != nil && c.errExpected == false {
			t.Errorf("error received while testing %s. The function call was: Sqrt(%v), and the error was: %v", c.description, c.a, err)
		}

		// Only fail on got != want if an error was not expected
		if c.errExpected == false && got != c.want {
			t.Errorf("want %v, got %v, while testing %s. The function call was: Sqrt(%v)", c.want, got, c.description, c.a)
		}
	}
}

func TestExpression(t *testing.T) {
	// Define test cases
	testCases := []expressionTestCase{
		{
			description: "two positive numbers which sum to a positive",
			e:           "2 + 2",
			want:        4,
		},
		{
			description: "an invalid operator",
			e:           "2 X 2",
			want:        123456789,
			errExpected: true,
		},
		{
			description: "an invalid multi-expression",
			e:           "2 + 2 * 2",
			want:        123456789,
			errExpected: true,
		},
	}

	t.Parallel()

	for _, c := range testCases {
		err, got := calculator.Expression("c.e")
		if err != nil && c.errExpected == false {
			t.Errorf("error received while testing %s. The function call was: Expression(%v), and the error was: %v", c.description, c.a, err)
		}

		// Only fail on got != want if an error was not expected
		if c.errExpected == false && got != c.want {
			t.Errorf("want %v, got %v, while testing %s. The function call was: Expression(%v)", c.want, got, c.description, c.a)
		}
	}
}
