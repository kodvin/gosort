package common

type Options struct {
	Fn func(a, b int) bool
}

func DefaultCompare(a, b int) bool {
	return (a > b)
}
