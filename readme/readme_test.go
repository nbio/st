package readme

import (
	"testing"

	"github.com/nbio/st"
)

func TestExample(t *testing.T) {
	st.Expect(t, "a", "a")
	st.Reject(t, 42, int64(42))

	st.Assert(t, "t", "t")
	st.Refute(t, 99, int64(99))
}

func TestTableExample(t *testing.T) {
	examples := []struct{ a, t string }{
		{"first", "first"},
		{"second", "second"},
	}

	// Pass example index to improve the error message for table-based tests.
	for i, ex := range examples {
		st.Expect(t, ex, ex, i)
		st.Reject(t, ex, &ex, i)
	}

	// Cannot pass index into Assert or Refute, they fail fast.
	for _, ex := range examples {
		st.Assert(t, ex, ex)
		st.Refute(t, ex, &ex)
	}
}

// Prints failure output, including the correct line number.
func TestFailedExpectationMessages(t *testing.T) {
	st.Expect(t, 1, 2)
	st.Reject(t, "same", "same")
	var typedNil *string
	st.Expect(t, typedNil, nil) // in Go, a typed nil != nil
}

// Prints failure output, including the correct line number.
func TestFailedAssertMessage(t *testing.T) {
	type chicken struct{}
	type egg struct{}
	st.Assert(t, egg{}, chicken{})
}

// Prints failure output, including the correct line number.
func TestFailedRefuteMessage(t *testing.T) {
	st.Reject(t, 42, 7*6)
}

// Prints failure output, including the correct line number and example index.
func TestFailedTableMessages(t *testing.T) {
	table := []struct{ val int }{
		{0}, {1}, {2},
	}
	// Continues if expectation fails
	for i, example := range table {
		st.Expect(t, example.val, 1, i)
	}
	// Stops when first assertion fails
	for _, example := range table {
		st.Assert(t, example.val, 1)
	}
}
