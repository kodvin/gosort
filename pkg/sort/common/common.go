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
)

var AllSortTypes [5]SortType = [...]SortType{
	"bubble", "quick", "merge", "insertion", "selection",
}

func DefaultCompare(a, b int) bool {
	return (a > b)
}
