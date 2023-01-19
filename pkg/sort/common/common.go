package common

import "golang.org/x/exp/constraints"

type Options[T constraints.Ordered] struct {
	Fn func(a, b T) bool
}

func DefaultCompare[T constraints.Ordered](i, j T) bool {
	return i > j
}
