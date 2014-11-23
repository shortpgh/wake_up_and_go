// one is my first implementation of a binary search method in go.
// to run;
// 	  &> out
//
// make sure your gopath is set appropriately before trying to run
// 	export GOPATH=$HOME/Projects/wake_up_and_go
package binsearch

type Error5 string

func (e Error5) Error() string {
	return string(e)
}

type SearchableArray5 []int

func (o SearchableArray5) BinarySearch(find int) (int, error) {
	curlist := o
	low, high := 0, len(curlist)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		switch {
		case find == curlist[mid]:
			return mid, nil
		case find > curlist[mid]:
			low = mid + 1
		default:
			high = mid - 1
		}

	}
	return 0, Error5("No item found")
}
