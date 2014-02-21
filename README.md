## A Simple Test micro-framework for Go

[![GoDoc](https://godoc.org/github.com/nbio/st?status.png)](https://godoc.org/github.com/nbio/st)

A tiny test framework for making short, useful assertions in your Go tests.

`Assert(t, actual, expected)` and `Refute(t, actual, expected)` abort a test immediately with `t.Fatal`.

`Expect(t, actual, expected)` and `Reject(t, actual, expected)` allow a test to continue, reporting failure at the end with `t.Error`.

They print nice error messages, preserving the order of actual == expected to minimize confusion.

### Usage

Examples of passing tests from `readme_test.go`:

```go
func TestExample(t *testing.T) {
	st.Expect(t, "a", "a")
	st.Reject(t, 42, int64(42))

	st.Assert(t, "b", "b")
	st.Refute(t, 99, int64(99))
}

func TestTableExample(t *testing.T) {
	examples := []struct{ a, b string }{
		{"first", "first"},
		{"second", "second"},
	}

	// Pass the index to improve the error message for table-based tests.
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
```

```console
=== RUN TestExample
--- PASS: TestExample (0.00 seconds)
=== RUN TestTableExample
--- PASS: TestTableExample (0.00 seconds)
PASS
ok  	github.com/nbio/st	0.010s
```

Failing tests produce nice output:

```go
func TestFailedExpectationMessages(t *testing.T) {
	st.Expect(t, 1, 2)
	st.Reject(t, "same", "same")
}

func TestFailedAssertMessage(t *testing.T) {
	type chicken struct{}
	type egg struct{}
	st.Assert(t, egg{}, chicken{})
}

func TestFailedRefuteMessage(t *testing.T) {
	st.Reject(t, 42, 7*6)
}

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

```

```console
--- FAIL: TestFailedExpectationMessages (0.00 seconds)
	st.go:31:
		readme_test.go:38: actual should == expected
		 	have (int): 2
			want (int): 1
	st.go:40:
		readme_test.go:39: actual should != expected
		 	have (string): same
			and  (string): same
--- FAIL: TestFailedAssertMessage (0.00 seconds)
	st.go:49:
		readme_test.go:46: actual should == expected
		 	have (readme.chicken): {}
			want (readme.egg): {}
--- FAIL: TestFailedRefuteMessage (0.00 seconds)
	st.go:40:
		readme_test.go:51: actual should != expected
		 	have (int): 42
			and  (int): 42
--- FAIL: TestFailedTableMessages (0.00 seconds)
	st.go:31:
		readme_test.go:61: actual should == expected
		0. 	have (int): 1
			want (int): 0
	st.go:31:
		readme_test.go:61: actual should == expected
		2. 	have (int): 1
			want (int): 2
	st.go:49:
		readme_test.go:65: actual should == expected
		 	have (int): 1
			want (int): 0
FAIL
FAIL	github.com/nbio/st/readme	0.010s
```

See [`package st`](https://godoc.org/github.com/nbio/st) documentation for more detail.
