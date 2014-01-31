// Copyright 2014 nb.io, LLC
// Author: Cameron Walters <cameron@nb.io>

package st

import (
	"fmt"
	"testing"
)

func Example_caller() {
	f := func() {
		file, line := caller()
		fmt.Printf("%s:%d", file, line)
	}
	f() // the output will contain this line's number
	// Output: st_test.go:16
}

type stTest struct{}

func TestExpectReject(t *testing.T) {
	// Standard expectations
	Expect(t, "a", "a")
	Expect(t, 42, 42)
	Expect(t, nil, nil)

	// Standard rejections
	Reject(t, "a", "A")
	Reject(t, 42, int64(42.0))
	Reject(t, 42, 42.0)
	Reject(t, 42, "42")
	Reject(t, []string{}, nil)
	Reject(t, []stTest{}, nil)

	// Table-based test
	examples := []struct{ a, b string }{
		{"first", "first"},
		{"second", "second"},
	}

	for i, ex := range examples {
		Expect(t, ex, ex, i)
		Expect(t, &ex, &ex, i)

		Reject(t, ex, &ex, i)
		Reject(t, ex, 0, i)
		Reject(t, ex, "", i)
		Reject(t, ex, byte('a'), i)
		Reject(t, ex, float64(5.9), i)
	}
}

func TestAssertRefute(t *testing.T) {
	// Standard assertions
	Assert(t, "a", "a")
	Assert(t, 42, 42)
	Assert(t, nil, nil)

	// Standard refutations
	Refute(t, "a", "A")
	Refute(t, 42, int64(42.0))
	Refute(t, 42, 42.0)
	Refute(t, 42, "42")
	Refute(t, []string{}, nil)
	Refute(t, []stTest{}, nil)

	// Table-based test
	examples := []struct{ a, b string }{
		{"first", "first"},
		{"second", "second"},
	}

	// Note: there's no argument to pass the index to assertions.
	for _, ex := range examples {
		Assert(t, ex, ex)
		Assert(t, &ex, &ex)

		Refute(t, ex, &ex)
		Refute(t, ex, 0)
		Refute(t, ex, "")
		Refute(t, ex, byte('a'))
		Refute(t, ex, float64(5.9))
	}
}

func TestExampleNum(t *testing.T) {
	expectationFunc := func(t *testing.T, n ...int) []int {
		return n
	}
	Expect(t, exampleNum(expectationFunc(t)), "")
	Expect(t, exampleNum(expectationFunc(t, 0)), "1.")
	Expect(t, exampleNum(expectationFunc(t, 1)), "2.")
	Expect(t, exampleNum(expectationFunc(t, 2)), "3.")
}

// NOT A REAL BENCHMARK
// Intended to demonstrate the output, including the correct line number, when a
// test fails one or more expectations.
func BenchmarkExpectationMessages(b *testing.B) {
	Expect(b, 1, 2)
	Refute(b, 1, 1)
}

// NOT A REAL BENCHMARK
// Intended to demonstrate the output, including the correct line number, when a
// test fails an assertion.
func BenchmarkAssertMessage(b *testing.B) {
	Assert(b, 1, 2)
}

// NOT A REAL BENCHMARK
// Intended to demonstrate the output, including the correct line number, when a
// test fails a refutation.
func BenchmarkRefuteMessage(b *testing.B) {
	Refute(b, 1, 1)
}
