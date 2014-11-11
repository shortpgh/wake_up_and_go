// one is my first implementation of a binary search method in go.
// to run;
// 	go test github.com/shortpgh/binsearch &> out
// 
// make sure your gopath is set appropriately before trying to run
// 	export GOPATH=$HOME/Projects/wake_up_and_go
package binsearch

func Search(list []int, find int) int {
	curlist := list
	var nextpos int
	modifier := 0
	for len(curlist) > 1 {
		nextpos = len(curlist) / 2
		if find == curlist[nextpos] {
			return nextpos + modifier
		}
		if find > curlist[nextpos] {
			newlist := curlist[nextpos:]
			modifier += len(curlist) - len(newlist)
			curlist = newlist
		} else {
			newlist := curlist[:nextpos]
			modifier = len(curlist) - len(newlist)
			curlist = newlist
		}
	}
	return 100
}
