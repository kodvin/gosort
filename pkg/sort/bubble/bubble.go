package bubble

import (
	"github.com/tobe-decided/gosort/pkg/sort/common"
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](nums []T, opt *common.Options[T]) []T {
	for range nums {
		for i := 0; i < len(nums)-1; i++ {
			if opt != nil {
				if opt.Fn(nums[i], nums[i+1]) {
					nums[i], nums[i+1] = nums[i+1], nums[i]
				}
			} else {
				if common.DefaultCompare(nums[i], nums[i+1]) {
					nums[i], nums[i+1] = nums[i+1], nums[i]
				}

			}
		}
	}

	return nums
}
