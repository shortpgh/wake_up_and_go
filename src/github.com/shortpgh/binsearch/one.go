// one is my first implementation of a binary search method in go.
// to run;
// 	go test github.com/shortpgh/binsearch &> out
//
// make sure your gopath is set appropriately before trying to run
// 	export GOPATH=$HOME/Projects/wake_up_and_go
package binsearch

func Search(list []int, find int) int {
	curlist := list
	low, high := 0, len(list)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		if find == curlist[mid] {
			return mid
		}
		if find > curlist[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 100
}
