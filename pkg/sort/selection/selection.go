package selection

import (
	"github.com/tobe-decided/gosort/pkg/sort/common"
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](nums []T, opts *common.Options[T]) []T {
	return sortNumbers(nums, opts)
}

func sortNumbers[T constraints.Ordered](nums []T, opts *common.Options[T]) []T {
	cmp := common.DefaultCompare[T]
	if opts != nil {
		cmp = opts.Fn
	}

	for i := range nums {
		extremumIndex := i
		for j := i; j < len(nums); j++ {
			if cmp(nums[extremumIndex], nums[j]) {
				extremumIndex = j
			}
		}
		nums[i], nums[extremumIndex] = nums[extremumIndex], nums[i]
	}
	return nums
}
