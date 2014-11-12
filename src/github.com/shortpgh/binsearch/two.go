// one is my first implementation of a binary search method in go.
// to run;
// 	go test github.com/shortpgh/binsearch &> out
//
// make sure your gopath is set appropriately before trying to run
// 	export GOPATH=$HOME/Projects/wake_up_and_go
package binsearch

type Error string

func (e Error) Error() string {
	return string(e)
}

func SearchTwo(list []int, find int) (int, error) {
	curlist := list
	low, high := 0, len(list)-1

	var mid int
	mid = (high + low) / 2

	var ret int
	var err error

	if len(curlist) == 0 {
		return 0, Error("No item found")
	}

	if find == curlist[mid] {
		return mid, nil
	}
	if find > curlist[mid] {
		low = mid + 1
		ret, err = SearchTwo(curlist[low:], find)
		return ret + low, err
	} else {
		high = mid
		ret, err = SearchTwo(curlist[:mid], find)
		return ret, err
	}

	return 0, Error("No item found")
}
