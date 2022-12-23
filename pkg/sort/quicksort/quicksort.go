package quicksort

import "github.com/tobe-decided/gosort/pkg/sort/common"

func Sort(nums []int, opts *common.Options) []int {
	return sortNumbers(nums, opts)
}

func sortNumbers(nums []int, opts *common.Options) []int {
	quicksort(nums, 0, len(nums)-1, opts)
	return nums
}

func quicksort(nums []int, low int, high int, opts *common.Options) {
	if low < high {
		pi := partition(nums, low, high, opts)

		quicksort(nums, low, pi-1, opts)
		quicksort(nums, pi+1, high, opts)
	}
}

func partition(nums []int, low int, high int, opts *common.Options) int {
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
