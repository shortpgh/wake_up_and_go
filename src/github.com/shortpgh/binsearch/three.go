// one is my first implementation of a binary search method in go.
// to run;
// 	go test github.com/shortpgh/binsearch &> out
//
// make sure your gopath is set appropriately before trying to run
// 	export GOPATH=$HOME/Projects/wake_up_and_go
package binsearch

type Error3 string

func (e Error3) Error() string {
	return string(e)
}

type SearchThree struct {
	list []int
	low  int
	high int
	mid  int
}

func (o SearchThree) BinSearch(find int) (int, error) {
	curlist := o.list
	o.low, o.high = 0, len(o.list)-1
	for o.low <= o.high {
		o.mid = (o.high + o.low) / 2
		if find == curlist[o.mid] {
			return o.mid, nil
		}
		if find > curlist[o.mid] {
			o.low = o.mid + 1
		} else {
			o.high = o.mid - 1
		}
	}

	return 0, Error3("No item found")

}
