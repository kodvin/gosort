package heap

import (
	"github.com/tobe-decided/gosort/pkg/sort/common"
)

func Sort(nums []int, opts *common.Options) []int {
	return sortNumbers(nums, opts)
}

func sortNumbers(nums []int, opts *common.Options) []int {
	return heapSort(nums, opts)
}

func heapSort(nums []int, opts *common.Options) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, len(nums), i, opts)
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0, opts)
	}
	return nums
}

func heapify(nums []int, size int, rootIndex int, opts *common.Options) {
	cmp := common.DefaultCompare
	if opts != nil {
		cmp = opts.Fn
	}
	extremum := rootIndex

	left := 2*rootIndex + 1
	right := 2*rootIndex + 2
	if left < size && cmp(nums[left], nums[extremum]) {
		extremum = left
	}
	if right < size && cmp(nums[right], nums[extremum]) {
		extremum = right
	}

	if extremum != rootIndex {
		nums[extremum], nums[rootIndex] = nums[rootIndex], nums[extremum]
		heapify(nums, size, extremum, opts)
	}
}
