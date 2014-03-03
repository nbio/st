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
	t.Log("Tests purposely fail to demonstrate output")
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

// Allows comparing non-comparable types to prevent panics when comparing slices
// or maps.
func TestDeeperEquality(t *testing.T) {
	type testStr string
	slice1 := []interface{}{"A", 1, []byte("steak sauce")}
	slice2 := []interface{}{"R", 2, 'd', int64(2)}
	map1 := map[string]string{"clever": "crafty", "modest": "prim"}
	map2 := map[string]string{"silk": "scarf", "wool": "sweater"}
	str1 := "same"
	str2 := testStr("same")

	st.Expect(t, slice1, slice2)
	st.Reject(t, slice1, slice1)
	st.Expect(t, map1, map2)
	st.Reject(t, map1, map1)
	st.Expect(t, str1, str2)
	st.Reject(t, str1, str1)
}
