package merge

import (
	"github.com/tobe-decided/gosort/pkg/sort/common"
)

func Sort(nums []int, opts *common.Options) []int {
	return sortNumbers(nums, opts)
}

func sortNumbers(nums []int, opts *common.Options) []int {
	res := mergesort(nums, 0, len(nums), opts)
	copy(nums, res)
	return res
}

func mergesort(nums []int, low int, high int, opts *common.Options) []int {
	var cmp func(int, int) bool

	if opts != nil {
		cmp = opts.Fn
	} else {
		cmp = common.DefaultCompare
	}
	if high-low == 2 {
		if cmp(nums[low], nums[low+1]) {
			return []int{nums[low+1], nums[low]}
		} else {
			return []int{nums[low], nums[low+1]}
		}
	} else if high-low == 1 {
		return []int{nums[low]}
	} else if high-low == 0 {
		return []int{}
	}
	res := make([]int, high-low)

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
