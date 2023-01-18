package quick

import (
	"github.com/tobe-decided/gosort/pkg/sort/common"
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](nums []T, opts *common.Options[T]) []T {
	return sortNumbers(nums, opts)
}

func sortNumbers[T constraints.Ordered](nums []T, opts *common.Options[T]) []T {
	quicksort(nums, 0, len(nums)-1, opts)
	return nums
}

func quicksort[T constraints.Ordered](nums []T, low int, high int, opts *common.Options[T]) {
	if low < high {
		pi := partition(nums, low, high, opts)

		quicksort(nums, low, pi-1, opts)
		quicksort(nums, pi+1, high, opts)
	}
}

func partition[T constraints.Ordered](nums []T, low int, high int, opts *common.Options[T]) int {
	pivot := nums[high]
	min := low

	for i := low; i < len(nums)-1; i++ {
		if opts != nil {
			if opts.Fn(pivot, nums[i]) {
				nums[min], nums[i] = nums[i], nums[min]
				min++
			}
		} else {
			if common.DefaultCompare(pivot, nums[i]) {
				nums[min], nums[i] = nums[i], nums[min]
				min++
			}
		}
	}
	nums[min], nums[high] = nums[high], nums[min]

	return min
}
