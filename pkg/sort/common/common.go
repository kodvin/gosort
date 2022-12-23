package common

type Options struct {
	Fn func(a, b int) bool
}

type SortType string

const (
	Bubble    SortType = "bubble"
	Quicksort SortType = "quicksort"
)

func DefaultCompare(a, b int) bool {
	return (a > b)
}
