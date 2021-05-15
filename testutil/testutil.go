package testutil

import (
	"fmt"

	"github.com/maja42/gotils/compare"
	"github.com/stretchr/testify/assert"
)

// TestingT is implemented by testing.T.
// The interface is used to ensure that the test-library us not pulled into non-test code.
type TestingT interface {
	// Name returns the name of the running test or benchmark.
	Name() string

	// Fail marks the function as having failed but continues execution.
	Fail()
	// Failed reports whether the function has failed.
	Failed() bool
	// FailNow marks the function as having failed and stops its execution
	// by calling runtime.Goexit (which then runs all deferred calls in the
	// current goroutine).
	// Execution will continue at the next test or benchmark.
	// FailNow must be called from the goroutine running the
	// test or benchmark function, not from other goroutines
	// created during the test. Calling FailNow does not stop
	// those other goroutines.
	FailNow()

	Log(args ...interface{})
	Logf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	// Fatal is equivalent to Log followed by FailNow.
	Fatal(args ...interface{})
	// Fatalf is equivalent to Logf followed by FailNow.
	Fatalf(format string, args ...interface{})

	// Skip is equivalent to Log followed by SkipNow.
	Skip(args ...interface{})
	// Skipf is equivalent to Logf followed by SkipNow.
	Skipf(format string, args ...interface{})
	// SkipNow marks the test as having been skipped and stops its execution
	// by calling runtime.Goexit.
	// If a test fails (see Error, Errorf, Fail) and is then skipped,
	// it is still considered to have failed.
	// Execution will continue at the next test or benchmark. See also FailNow.
	// SkipNow must be called from the goroutine running the test, not from
	// other goroutines created during the test. Calling SkipNow does not stop
	// those other goroutines.
	SkipNow()
	// Skipped reports whether the test was skipped.
	Skipped() bool

	Helper()
	Cleanup(f func())
	TempDir() string
	Parallel()
}

// AssertErrorContains verifies that the error message contains a given substring.
func AssertErrorContains(t TestingT, err error, contains string) bool {
	t.Helper()
	if assert.Error(t, err) {
		return assert.Contains(t, err.Error(), contains)
	}
	return false
}

// AssertElementsMatch verifies that the elements in the given arrays or slices match, ignoring their order.
// Duplicate elements are treated individually, meaning the number of occurrences is verified.
// Elements are compared using the provided equal function.
// Reports the list of mismatches in a humanly readable manner.
func AssertElementsMatch(t TestingT, msg string, expected, actual interface{}, compareFunc func(expected, actual interface{}) bool) bool {
	// based on testify.assert (ElementsMatch)
	t.Helper()

	if msg == "" {
		msg = "Unexpected elements\n"
	} else {
		msg += "\n"
	}
	okay, missing, unexpected := compare.DiffUnorderedSlice(expected, actual, compareFunc)
	if len(missing) > 0 || len(unexpected) > 0 {
		msg := msg
		for _, nw := range okay {
			msg += fmt.Sprintf("        [OK] %s\n", nw)
		}
		for _, nw := range missing {
			msg += fmt.Sprintf("   [MISSING] %s\n", nw)
		}
		for _, nw := range unexpected {
			msg += fmt.Sprintf("[UNEXPECTED] %s\n", nw)
		}
		t.Error(msg)
		return false
	}
	return true
}
