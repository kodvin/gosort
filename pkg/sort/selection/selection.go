package selection

import (
	"github.com/tobe-decided/gosort/pkg/sort/common"
)

func Sort(nums []int, opts *common.Options) []int {
	return sortNumbers(nums, opts)
}

func sortNumbers(nums []int, opts *common.Options) []int {
	cmp := common.DefaultCompare
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
