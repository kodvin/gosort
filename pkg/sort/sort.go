package sort

import (
	Bubble "github.com/tobe-decided/gosort/pkg/sort/bubble"
	"github.com/tobe-decided/gosort/pkg/sort/common"
	Quicksort "github.com/tobe-decided/gosort/pkg/sort/quicksort"
)

func Sort(s common.SortType, content []int) []int {
	switch s {
	case common.Bubble:
		return Bubble.Sort(content, nil)
	case common.Quicksort:
		return Quicksort.Sort(content, nil)
	}
	return content
}
