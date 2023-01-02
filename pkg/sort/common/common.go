package common

type Options struct {
	Fn func(a, b int) bool
}

type SortType string

const (
	Bubble SortType = "bubble"
	Quick  SortType = "quick"
)

func DefaultCompare(a, b int) bool {
	return (a > b)
}
