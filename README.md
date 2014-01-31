## A Simple Test micro-framework for Go

Pronounced "ghost", this is the smallest possible test framework for making short, useful assertions in your Go tests.

`Assert(t, actual, expected)` and `Refute(t, actual, expected)` abort a test immediately with `t.Fatal`.

`Expect(t, actual, expected)` and `Reject(t, actual, expected)` allow a test to continue, reporting failure at the end with `t.Error`.

### Usage

A contrived example for demonstration purposes:

```go
import (
	"testing"
	"github.com/nbio/st"
)

func TestExample(t *testing.T) {
	st.Expect(t, "a", "a")
	st.Expect(t, uint8(0), byte(0))
	st.Reject(t, "a", "A")
	st.Reject(t, []string{}, nil)

	// Table-based tests
	examples := []struct{ a, b string }{
		{"first", "first"},
		{"second", "second"},
	}

	for i, ex := range examples {
		st.Expect(t, ex, ex, i)
		st.Reject(t, ex, &ex, i)
	}

	st.Assert(t, "a", "a")
	st.Assert(t, uint8(0), byte(0))
	st.Refute(t, "a", "A")
	st.Refute(t, []string{}, nil)
}
```

See GoDoc at godoc.org/github.com/nbio/st for more detail.

Run benchmarks to see example output:

```
BenchmarkExpectationMessages	--- FAIL: BenchmarkExpectationMessages
	st.go:23:
		st_test.go:101: expected equality
		 	want (type int): 2
			have (type int): 1
	st.go:50:
		st_test.go:102: expected inequality
		 	want (type int): 1
			have (type int): 1
BenchmarkAssertMessage	--- FAIL: BenchmarkAssertMessage
	st.go:41:
		st_test.go:109: expected equality
		 	want (type int): 2
			have (type int): 1
BenchmarkRefuteMessage	--- FAIL: BenchmarkRefuteMessage
	st.go:50:
		st_test.go:116: expected inequality
		 	want (type int): 1
			have (type int): 1
ok  	github.com/nbio/st	0.010s

```
