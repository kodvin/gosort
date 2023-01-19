package merge

import (
	"github.com/tobe-decided/gosort/pkg/sort/common"
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](nums []T, opts *common.Options[T]) []T {
	res := mergesort(nums, 0, len(nums), opts)
	copy(nums, res)
	return res
}

func mergesort[T constraints.Ordered](nums []T, low int, high int, opts *common.Options[T]) []T {
	cmp := common.DefaultCompare[T]
	if opts != nil {
		cmp = opts.Fn
	}
	if high-low == 2 {
		if cmp(nums[low], nums[low+1]) {
			return []T{nums[low+1], nums[low]}
		} else {
			return []T{nums[low], nums[low+1]}
		}
	} else if high-low == 1 {
		return []T{nums[low]}
	} else if high-low == 0 {
		return []T{}
	}
	res := make([]T, high-low)

	a1 := mergesort(nums, low, low+(high-low)/2, opts)
	a2 := mergesort(nums, low+(high-low)/2, high, opts)

	l1 := len(a1)
	l2 := len(a2)
	lr := 0
	for l1 > 0 && l2 > 0 {
		l1i := len(a1) - l1
		l2i := len(a2) - l2
		if cmp(a1[l1i], a2[l2i]) {
			res[lr] = a2[l2i]
			l2--
		} else {
			res[lr] = a1[l1i]
			l1--
		}
		lr++
	}

	for l1 > 0 {
		l1i := len(a1) - l1
		res[lr] = a1[l1i]
		l1--
		lr++
	}

	for l2 > 0 {
		l2i := len(a2) - l2
		res[lr] = a2[l2i]
		l2--
		lr++
	}

	return res
}
