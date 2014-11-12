// tests for two
package binsearch

import "testing"

func TestSearchTwo(t *testing.T) {

	list := []int{2, 3, 5, 6, 7, 11, 13}

	TrySearchTwo(t, list, 5, 2, false)
	TrySearchTwo(t, list, 11, 5, false)
	TrySearchTwo(t, list, 6, 3, false)
	TrySearchTwo(t, list, 2, 0, false)
	TrySearchTwo(t, list, 13, 6, false)
	TrySearchTwo(t, list, 0, 0, true)
	TrySearchTwo(t, list, 25, 7, true)
}

func TrySearchTwo(t *testing.T, list []int, in int, out int, isErr bool) {
	result, err := SearchTwo(list, in)
	switch {
	case isErr && err == nil:
		t.Errorf("Search(%v) = %v, want error", in, result)
	case !isErr && result != out:
		t.Errorf("Search(%v) = %v, want %v", in, result, out)
	}
}
