// tests for two
package binsearch

import "testing"

func TestSearchFour(t *testing.T) {

	list := SearchableArray{2, 3, 5, 6, 7, 11, 13}

	TrySearchFour(t, list, 5, 2, false)
	TrySearchFour(t, list, 11, 5, false)
	TrySearchFour(t, list, 6, 3, false)
	TrySearchFour(t, list, 2, 0, false)
	TrySearchFour(t, list, 13, 6, false)
	TrySearchFour(t, list, 0, 0, true)
	TrySearchFour(t, list, 25, 7, true)
}

func TrySearchFour(t *testing.T, list SearchableArray, in int, out int, isErr bool) {
	result, err := list.BinSearch(in)
	switch {
	case isErr && err == nil:
		t.Errorf("Search(%v) = %v, want error", in, result)
	case !isErr && result != out:
		t.Errorf("Search(%v) = %v, want %v", in, result, out)
	}
}
