// Copyright 2014 nb.io, LLC
// Author: Cameron Walters <cameron@nb.io>

// Package st, pronounced "ghost", is a tiny test framework for
// making short, useful assertions in your Go tests.
//
// To abort a test immediately with t.Fatal, use
// Assert(t, actual, expected) and Refute(t, actual, expected)
//
// To allow a test to continue, reporting failure at the end with t.Error, use
// Expect(t, actual, expected) and Reject(t, actual, expected)
package st

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"reflect"
)

const (
	equal   = "\n%s:%d: actual should == expected\n%s \thave (%T): %+v\n\twant (%T): %+v"
	unequal = "\n%s:%d: actual should != expected\n%s \thave (%T): %+v\n\tand  (%T): %+v"
)

// Expect calls t.Error and prints a nice comparison message when act != exp.
// Especially useful in table-based tests when passing the loop index as iter.
func Expect(t testing.TB, act interface{}, exp interface{}, iter ...int) {
	if !reflect.DeepEqual(act, exp) {
		file, line := caller()
		t.Errorf(equal, file, line, exampleNum(iter), exp, exp, act, act)
	}
}

// Reject calls t.Error and prints a nice comparison message when act == exp.
// Especially useful in table-based tests when passing the loop index as iter.
func Reject(t testing.TB, act interface{}, exp interface{}, iter ...int) {
	if reflect.DeepEqual(act, exp) {
		file, line := caller()
		t.Errorf(unequal, file, line, exampleNum(iter), exp, exp, act, act)
	}
}

// Assert calls t.Fatal to abort the test immediately and prints a nice
// comparison message when act != exp.
func Assert(t testing.TB, act interface{}, exp interface{}) {
	if !reflect.DeepEqual(act, exp) {
		file, line := caller()
		t.Fatalf(equal, file, line, "", exp, exp, act, act)
	}
}

// Refute calls t.Fatal to abort the test immediately and prints a nice
// comparison message when act != exp.
func Refute(t testing.TB, act interface{}, exp interface{}) {
	if reflect.DeepEqual(act, exp) {
		file, line := caller()
		t.Fatalf(unequal, file, line, "", exp, exp, act, act)
	}
}

// returns file and line two stack frames above its invocation
func caller() (file string, line int) {
	var ok bool
	_, file, line, ok = runtime.Caller(2)
	if !ok {
		file = "???"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return
}

// returns an example number from the optional zero-based loop iterator n, if
// provided.
func exampleNum(n []int) string {
	if len(n) == 1 {
		return fmt.Sprintf("%d.", n[0])
	}
	return ""
}
