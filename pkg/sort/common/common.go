package common

type Options struct {
	Fn func(a, b int) bool
}

type SortType string

const (
	Bubble    SortType = "bubble"
	Quick     SortType = "quick"
	Merge     SortType = "merge"
	Insertion SortType = "insertion"
	Selection SortType = "selection"
	Heap      SortType = "heap"
)

var AllSortTypes [6]SortType = [...]SortType{
	"bubble", "quick", "merge", "insertion", "selection", "heap",
}

func DefaultCompare(a, b int) bool {
	return (a > b)
}
