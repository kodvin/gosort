package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tobe-decided/gosort/pkg/sort/common"
)

var expectedAscending []int = []int{1, 2, 3, 4, 5, 6, 10, 100, 101, 150, 150, 200}
var expectedDescending []int = []int{200, 150, 150, 101, 100, 10, 6, 5, 4, 3, 2, 1}
var initial []int = []int{1, 101, 3, 150, 4, 6, 10, 2, 100, 150, 200, 5}

var expectedAscendingString []string = []string{"animals", "cat", "dog", "mice", "other"}
var expectedDescendingString []string = []string{"other", "mice", "dog", "cat", "animals"}
var initialString []string = []string{"dog", "cat", "mice", "other", "animals"}

var expectedAscendingFloat []float32 = []float32{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 10.1, 100.1, 101.1, 150.1, 150.1, 200.1}
var expectedDescendingFloat []float32 = []float32{200.1, 150.1, 150.1, 101.1, 100.1, 10.1, 6.1, 5.1, 4.1, 3.1, 2.1, 1.1}
var initialFloat []float32 = []float32{1.1, 101.1, 3.1, 150.1, 4.1, 6.1, 10.1, 2.1, 100.1, 150.1, 200.1, 5.1}

func TestNumbers_DefaultSorting(t *testing.T) {
	unordered := make([]int, len(initial))

	for _, sType := range AllSortTypes {
		copy(unordered, initial)
		Sort(sType, unordered, nil)
		assert.Equal(t, expectedAscending, unordered, "Sorting int. SortType: "+sType)
	}
}

func TestNumbers_CustomSorting_Ascending(t *testing.T) {
	unordered := make([]int, len(initial))
	for _, sType := range AllSortTypes {
		copy(unordered, initial)
		Sort(sType, unordered, &common.Options[int]{Fn: func(a, b int) bool { return (a > b) }})
		assert.Equal(t, expectedAscending, unordered, "Sorting int. SortType: "+sType)
	}
}

func TestNumbers_CustomSorting_Descending(t *testing.T) {
	unordered := make([]int, len(initial))

	for _, sType := range AllSortTypes {
		copy(unordered, initial)
		Sort(sType, unordered, &common.Options[int]{Fn: func(a, b int) bool { return (a < b) }})
		assert.Equal(t, expectedDescending, unordered, "Sorting int. SortType: "+sType)
	}
}

func TestString_DefaultSorting(t *testing.T) {
	unordered := make([]string, len(initialString))

	for _, sType := range AllSortTypes {
		copy(unordered, initialString)
		Sort(sType, unordered, nil)
		assert.Equal(t, expectedAscendingString, unordered, "Sorting string. SortType: "+sType)
	}
}
func TestString_CustomSorting_Ascending(t *testing.T) {
	unordered := make([]string, len(initialString))
	for _, sType := range AllSortTypes {
		copy(unordered, initialString)
		Sort(sType, unordered, &common.Options[string]{Fn: func(a, b string) bool { return (a > b) }})
		assert.Equal(t, expectedAscendingString, unordered, "Sorting string. SortType: "+sType)
	}
}
func TestString_CustomSorting_Descending(t *testing.T) {
	unordered := make([]string, len(initialString))

	for _, sType := range AllSortTypes {
		copy(unordered, initialString)
		Sort(sType, unordered, &common.Options[string]{Fn: func(a, b string) bool { return (a < b) }})
		assert.Equal(t, expectedDescendingString, unordered, "Sorting string. SortType: "+sType)
	}
}

func TestFloat_DefaultSorting(t *testing.T) {
	unordered := make([]float32, len(initialFloat))

	for _, sType := range AllSortTypes {
		copy(unordered, initialFloat)
		Sort(sType, unordered, nil)
		assert.Equal(t, expectedAscendingFloat, unordered, "Sorting string. SortType: "+sType)
	}
}
func TestFloat_CustomSorting_Ascending(t *testing.T) {
	unordered := make([]float32, len(initialFloat))
	for _, sType := range AllSortTypes {
		copy(unordered, initialFloat)
		Sort(sType, unordered, &common.Options[float32]{Fn: func(a, b float32) bool { return (a > b) }})
		assert.Equal(t, expectedAscendingFloat, unordered, "Sorting string. SortType: "+sType)
	}
}
func TestFloat_CustomSorting_Descending(t *testing.T) {
	unordered := make([]float32, len(initialFloat))

	for _, sType := range AllSortTypes {
		copy(unordered, initialFloat)
		Sort(sType, unordered, &common.Options[float32]{Fn: func(a, b float32) bool { return (a < b) }})
		assert.Equal(t, expectedDescendingFloat, unordered, "Sorting string. SortType: "+sType)
	}
}
