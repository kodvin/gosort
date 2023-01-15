package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tobe-decided/gosort/pkg/sort/common"
)

var expectedAscending []int = []int{1, 2, 3, 4, 5, 6, 10, 100, 101, 150, 150, 200}
var expectedDescending []int = []int{200, 150, 150, 101, 100, 10, 6, 5, 4, 3, 2, 1}
var initial []int = []int{1, 101, 3, 150, 4, 6, 10, 2, 100, 150, 200, 5}

func TestNumbers_DefaultSorting(t *testing.T) {
	unordered := make([]int, len(initial))

	for _, sType := range common.AllSortTypes {
		copy(unordered, initial)
		Sort(sType, unordered, nil)
		assert.Equal(t, expectedAscending, unordered, "SortType: "+sType)
	}
}

func TestNumbers_CustomSorting_Ascending(t *testing.T) {
	unordered := make([]int, len(initial))
	for _, sType := range common.AllSortTypes {
		copy(unordered, initial)
		Sort(sType, unordered, &common.Options{Fn: func(a, b int) bool { return (a > b) }})
		assert.Equal(t, expectedAscending, unordered, "SortType: "+sType)
	}
}

func TestNumbers_CustomSorting_Descending(t *testing.T) {
	unordered := make([]int, len(initial))

	for _, sType := range common.AllSortTypes {
		copy(unordered, initial)
		Sort(sType, unordered, &common.Options{Fn: func(a, b int) bool { return (a < b) }})
		assert.Equal(t, expectedDescending, unordered, "SortType: "+sType)
	}
}
