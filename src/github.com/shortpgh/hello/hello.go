// Day 1: 2014-11-01
// to run:
// 	>> GOPATH=/Users/nmihalick/Projects/wake_up_and_go
// 	>> go run pkg/github.com/shortpgh/wake_up_and_go/hello.go
//
package main

import (
	"fmt"
	"github.com/shortpgh/newmath"
)

func main() {
	fmt.Printf("hello, world! Sqrt(2) = %v\n", newmath.Sqrt(2))
}
