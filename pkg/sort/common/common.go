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
)

var AllSortTypes [4]SortType = [...]SortType{"bubble", "quick", "merge", "insertion"}

func DefaultCompare(a, b int) bool {
	return (a > b)
}
