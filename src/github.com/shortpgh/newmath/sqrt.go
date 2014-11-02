// Day 2:  2014-11-02
//
// This is a trivial example package.
//
// use `go build github.com/shortpgh/newmath/sqrt.go' to build this package.
// use `go install github.com/shortpgh/newmath` to install this package.
// After adding it to the hello program, rerun `go install github.com/shortpgh/home` to reinstall that program.
// Installing a main package will statically link to the included packages.
//
package newmath

//Sqrt returns an approximation of the square root of x
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
