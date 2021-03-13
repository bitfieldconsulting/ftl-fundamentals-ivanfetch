package calculator_test

import (
	"calculator"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

// Note Fatal|FatalF are also useful methods,
// to stop further execution within the same test.

// Return the name of a function.
func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// Return a slice of N random numbers from 0-500
func randFloat64Slice(n int) []float64 {
	rand.Seed(time.Now().UnixNano())
	r := make([]float64, n)

	// This `range` serves to loop the right number of times for the size of r
	for i := range r {
		x := rand.Float64() * float64(rand.Intn(500))
		r[i] = x
	}
	return r
}

// Test Add(), Subtract(), and Multiply()
func TestAddSubtractMultiply(t *testing.T) {
	// Define test cases
	// This anonymous struct has no name, as its only used here.
	testCases := []struct {
		function    func(float64, float64, ...float64) float64
		a           float64
		b           float64
		want        float64
		description string
	}{
		{
			function:    calculator.Add,
			description: "two positive numbers which sum to a positive",
			a:           2,
			b:           2,
			want:        4,
		},
		{
			function:    calculator.Add,
			description: "a positive and negative number which sum to a positive",
			a:           7,
			b:           -2,
			want:        5,
		},
		{
			function:    calculator.Add,
			description: "a positive and negative number which sum to a negative",
			a:           3,
			b:           -5,
			want:        -2,
		},
		{
			function:    calculator.Subtract,
			description: "two positive numbers whos difference is negative",
			a:           2,
			b:           9,
			want:        -7,
		},
		{
			function:    calculator.Subtract,
			description: "two positive numbers whos difference is positive",
			a:           7,
			b:           2,
			want:        5,
		},
		{
			function:    calculator.Subtract,
			description: "one positive and one negative decimal number whos difference is a positive decimal",
			a:           3,
			b:           -2.5,
			want:        5.5,
		},
		{
			function:    calculator.Multiply,
			description: "two positive numbers whos product is positive",
			a:           2,
			b:           20,
			want:        40,
		},
		{
			function:    calculator.Multiply,
			description: " a positive and negative number whos product is negative",
			a:           7,
			b:           -2,
			want:        -14,
		},
		{
			function:    calculator.Multiply,
			description: "a positive decimal and negative decimal whos product is a negative decimal",
			a:           8.4,
			b:           -2.5,
			want:        -21,
		},
	}

	t.Parallel()

	for _, c := range testCases {
		got := c.function(c.a, c.b)
		if c.want != got {
			t.Errorf("want %v, got %v, while testing %s. The function call was: %s(%v, %v)", c.want, got, c.description, getFuncName(c.function), c.a, c.b)
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
		if want != got {
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
	b := rand.Float64() * float64(rand.Intn(500))
	// v will be assigned multiple random float64 in a slice
	v := randFloat64Slice(5)

	// Get our own sum to compare to Add()
	var want float64 = a + b
	for _, x := range v {
		want += x
	}

	got := calculator.Add(a, b, v...)
	t.Logf("Random variadic test: Add(%v, %v, %v), wants %v, got %v", a, b, v, want, got)
	if want != got {
		t.Errorf("want %v, got %v, while testing a random variadic case. The function call was: Add(%v, %v, %v)", want, got, a, b, v)
	}
}

// Randomly test the variadic call to Subtract()
func TestSubtractVariadicRandomly(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	// rand.Float64() returns a number in 0.0-1.0
	// Use another randomly-generated int to vary the whole number.
	a := rand.Float64() * float64(rand.Intn(500))
	b := rand.Float64() * float64(rand.Intn(500))
	// v will be assigned multiple random float64 in a slice
	v := randFloat64Slice(5)

	// Get our own difference to compare to Subtract()
	var want float64 = a - b
	for _, x := range v {
		want -= x
	}

	got := calculator.Subtract(a, b, v...)
	t.Logf("Random variadic test: Subtract(%v, %v, %v), wants %v, got %v", a, b, v, want, got)
	if want != got {
		t.Errorf("want %v, got %v, while testing a random variadic case. The function call was: Subtract(%v, %v, %v)", want, got, a, b, v)
	}
}

// Randomly test the variadic call to Multiply()
func TestMultiplyVariadicRandomly(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())

	// rand.Float64() returns a number in 0.0-1.0
	// Use another randomly-generated int to vary the whole number.
	a := rand.Float64() * float64(rand.Intn(500))
	b := rand.Float64() * float64(rand.Intn(500))
	// v will be assigned multiple random float64 in a slice
	v := randFloat64Slice(5)

	// Get our own product to compare to Multiply()
	var want float64 = a * b
	for _, x := range v {
		want = want * x
	}

	got := calculator.Multiply(a, b, v...)
	t.Logf("Random variadic test: Multiply(%v, %v, %v), wants %v, got %v", a, b, v, want, got)
	if want != got {
		t.Errorf("want %v, got %v, while testing a random variadic case. The function call was: Multiply(%v, %v, %v)", want, got, a, b, v)
	}
}

func TestDivide(t *testing.T) {
	// Define test cases
	testCases := []struct {
		a           float64
		b           float64
		want        float64
		description string
		errExpected bool
	}{
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
		got, err := calculator.Divide(c.a, c.b)
		if err != nil && !c.errExpected {
			t.Errorf("error received while testing %s. The function call was: Divide(%v, %v), and the error was: %v", c.description, c.a, c.b, err)
		}

		// Only fail on want != got if an error was not expected
		if !c.errExpected && c.want != got {
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
	b := rand.Float64() * float64(rand.Intn(500))

	// Avoid dividing by 0
	if b == 0 {
		b += 0.1
	}

	// v will be assigned multiple random float64 in a slice
	v := randFloat64Slice(5)

	// Get our own quotient to compare to Divide()
	// Also change any randomly-generated 0.0s, to avoid dividing-by-zero!
	var want float64 = a / b
	for i, x := range v {
		if x == 0 {
			x += 0.1
			v[i] = x // change the slice
		}
		want = want / x
	}

	got, err := calculator.Divide(a, b, v...)
	if err != nil {
		t.Errorf("error received while testing a random variadic case. The function call was: Divide(%v, %v, %v), and the error was: %v", a, b, v, err)
	}

	t.Logf("Random variadic test: Divide(%v, %v, %v), wants %v, got %v", a, b, v, want, got)
	if want != got {
		t.Errorf("want %v, got %v, while testing a random variadic case. The function call was: Divide(%v, %v, %v)", want, got, a, b, v)
	}
}

func TestSqrt(t *testing.T) {
	// Define test cases
	testCases := []struct {
		a           float64
		b           float64
		want        float64
		description string
		errExpected bool
	}{
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
		got, err := calculator.Sqrt(c.a)
		if err != nil && !c.errExpected {
			t.Errorf("error received while testing %s. The function call was: Sqrt(%v), and the error was: %v", c.description, c.a, err)
		}

		// Only fail on want != got if an error was not expected
		if !c.errExpected && c.want != got {
			t.Errorf("want %v, got %v, while testing %s. The function call was: Sqrt(%v)", c.want, got, c.description, c.a)
		}
	}
}

// Test evaluating an expression,
// a string representation of an operation like 1 + 8
func TestEvaluateExpression(t *testing.T) {
	// Define test cases
	testCases := []struct {
		e, description string
		want           float64
		errExpected    bool
	}{
		{
			description: "an expression with two positive numbers which sum to a positive",
			e:           "2+2",
			want:        4,
		},
		{
			description: "an expression with two positive numbers whos difference is negative",
			e:           "2 - 9",
			want:        -7,
		},
		{
			description: "an expression with two positive numbers whos product is positive",
			e:           "2 * 20",
			want:        40,
		},
		{
			description: "an expression that divides by zero",
			e:           "2 / 0",
			want:        123456789,
			errExpected: true,
		},
		{
			description: "an expression with two positive numbers whos quotient is positive",
			e:           "20 / 2",
			want:        10,
		},
		{
			description: "an expression with an invalid operator",
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
		got, err := calculator.EvaluateExpression(c.e)
		if err != nil && !c.errExpected {
			t.Errorf("error received while testing %s. The function call was: Expression(%v), and the error was: %v", c.description, c.e, err)
		}

		// Only fail on want != got if an error was not expected
		if !c.errExpected && c.want != got {
			t.Errorf("want %v, got %v, while testing %s. The function call was: Expression(%v)", c.want, got, c.description, c.e)
		}
	}
}
