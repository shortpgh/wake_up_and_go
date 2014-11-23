// tests for the fifth binary search implementation
package binsearch

import "testing"

func TestSearchFive(t *testing.T) {

	list := SearchableArray5{2, 3, 5, 6, 7, 11, 13}

	TrySearchFive(t, list, 5, 2, false)
	TrySearchFive(t, list, 11, 5, false)
	TrySearchFive(t, list, 6, 3, false)
	TrySearchFive(t, list, 2, 0, false)
	TrySearchFive(t, list, 13, 6, false)
	TrySearchFive(t, list, 0, 0, true)
	TrySearchFive(t, list, 25, 7, true)
}

func TrySearchFive(t *testing.T, list SearchableArray5, in int, out int, isErr bool) {
	result, err := list.BinarySearch(in)
	switch {
	case isErr && err == nil:
		t.Errorf("Search(%v) = %v, want error", in, result)
	case !isErr && result != out:
		t.Errorf("Search(%v) = %v, want %v", in, result, out)
	}
}
