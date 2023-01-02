package merge

import "github.com/tobe-decided/gosort/pkg/sort/common"

func Sort(nums []int, opts *common.Options) []int {
	return sortNumbers(nums, opts)
}

func sortNumbers(nums []int, opts *common.Options) []int {
	mergesort(nums, 0, len(nums)-1, opts)
	return nums
}

func mergesort(nums []int, low int, high int, opts *common.Options) {
	if low < high {
		pi := low + (high-low)/2

		mergesort(nums, low, pi-1, opts)
		mergesort(nums, pi+1, high, opts)
	}
}
