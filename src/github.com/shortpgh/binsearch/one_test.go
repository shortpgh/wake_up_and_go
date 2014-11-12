// tests for one
package binsearch

import "testing"
import "fmt"

func TestSearch(t *testing.T) {
	list := []int{2, 3, 5, 6, 7, 11, 13}
	in1, out1 := 5, 2
	in2, out2 := 11, 5
	in3, out3 := 6, 3
	TrySearch(t, list, in1, out1)
	TrySearch(t, list, in2, out2)
	TrySearch(t, list, in3, out3)
	TrySearch(t, list, 2, 0)
	TrySearch(t, list, 13, 6)
	TrySearch(t, list, 0, 100)
	TrySearch(t, list, 25, 100)
}

func TrySearch(t *testing.T, list []int, in int, out int) {
	fmt.Println(in)
	if result := Search(list, in); result != out {
		t.Errorf("Search(%v) = %v, want %v", in, result, out)
	}
}
