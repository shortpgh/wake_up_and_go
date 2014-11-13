// tests for two
package binsearch

import "testing"

func TestSearchThree(t *testing.T) {

	list := []int{2, 3, 5, 6, 7, 11, 13}

	TrySearchThree(t, list, 5, 2, false)
	TrySearchThree(t, list, 11, 5, false)
	TrySearchThree(t, list, 6, 3, false)
	TrySearchThree(t, list, 2, 0, false)
	TrySearchThree(t, list, 13, 6, false)
	TrySearchThree(t, list, 0, 0, true)
	TrySearchThree(t, list, 25, 7, true)
}

func TrySearchThree(t *testing.T, list []int, in int, out int, isErr bool) {
	o := new(SearchThree)
	o.list = list
	result, err := o.BinSearch(in)
	switch {
	case isErr && err == nil:
		t.Errorf("Search(%v) = %v, want error", in, result)
	case !isErr && result != out:
		t.Errorf("Search(%v) = %v, want %v", in, result, out)
	}
}
