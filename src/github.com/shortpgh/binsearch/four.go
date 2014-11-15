// one is my first implementation of a binary search method in go.
// to run;
// 	  &> out
//
// make sure your gopath is set appropriately before trying to run
// 	export GOPATH=$HOME/Projects/wake_up_and_go
package binsearch

type Error4 string

func (e Error4) Error() string {
	return string(e)
}

type SearchableArray []int

func (o SearchableArray) BinSearch(find int) (int, error) {
	curlist := o
	low, high := 0, len(curlist)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		if find == curlist[mid] {
			return mid, nil
		}
		if find > curlist[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0, Error4("No item found")
}
