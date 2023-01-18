package sort

import (
	"github.com/tobe-decided/gosort/pkg/sort/bubble"
	"github.com/tobe-decided/gosort/pkg/sort/common"
	"github.com/tobe-decided/gosort/pkg/sort/heap"
	"github.com/tobe-decided/gosort/pkg/sort/insertion"
	"github.com/tobe-decided/gosort/pkg/sort/merge"
	"github.com/tobe-decided/gosort/pkg/sort/quick"
	"github.com/tobe-decided/gosort/pkg/sort/selection"
)

type SortType string

const (
	Bubble    SortType = "bubble"
	Quick     SortType = "quick"
	Merge     SortType = "merge"
	Insertion SortType = "insertion"
	Selection SortType = "selection"
	Heap      SortType = "heap"
)

var AllSortTypes [6]SortType = [...]SortType{
	"bubble", "quick", "merge", "insertion", "selection", "heap",
}

func Sort(s SortType, content []int, opts *common.Options) []int {
	switch s {
	case Bubble:
		return bubble.Sort(content, opts)
	case Quick:
		return quick.Sort(content, opts)
	case Merge:
		return merge.Sort(content, opts)
	case Insertion:
		return insertion.Sort(content, opts)
	case Selection:
		return selection.Sort(content, opts)
	case Heap:
		return heap.Sort(content, opts)
	default:
		panic(s + " method not implemented")
	}
}
