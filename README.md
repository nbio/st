# st a simple test micro-framework for Go

Pronounced "ghost", this is the smallest possible test framework for making short, useful assertions in your Go tests.

`Assert(t, actual, expected)` and `Refute(t, actual, expected)` abort a test immediately with `t.Fatal`.

`Expect(t, actual, expected)` and `Reject(t, actual, expected)` allow a test to continue, reporting failure at the end with `t.Error`.

See GoDoc at godoc.org/github.com/nbio/st


Run benchmarks to see example output:

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

