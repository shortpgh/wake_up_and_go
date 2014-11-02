// Day2: 2014-11-02
//
// This is a test for the Sqrt package written in this directory.
// run using `go test github.com/shortpgh/newmath` -- this technically will run all tests.
//
// You write a test by creating a file with a name ending in _test.go that contains functions
// named TestXXX with signature func (t *testing.T). The test framework runs each such
// function; if the function calls a failure function such as t.Error or t.Fail, the test is
// considered to have failed.
//
package newmath

// this package has to match for all files in this folder.

import "testing"

func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}
