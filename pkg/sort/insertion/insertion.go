package insertion

import (
	"github.com/tobe-decided/gosort/pkg/sort/common"
)

func Sort(nums []int, opts *common.Options) []int {
	return sortNumbers(nums, opts)
}

func sortNumbers(nums []int, opts *common.Options) []int {
	var cmp func(int, int) bool = common.DefaultCompare
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
