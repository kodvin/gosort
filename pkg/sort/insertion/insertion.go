package insertion

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
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && cmp(nums[j-1], nums[j]); {
			nums[j], nums[j-1] = nums[j-1], nums[j]
			j--
		}
	}

	return nums
}
