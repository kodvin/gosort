package sort

import (
	"github.com/tobe-decided/gosort/pkg/sort/bubble"
	"github.com/tobe-decided/gosort/pkg/sort/common"
	"github.com/tobe-decided/gosort/pkg/sort/merge"
	"github.com/tobe-decided/gosort/pkg/sort/quick"
)

func Sort(s common.SortType, content []int, opts *common.Options) []int {
	switch s {
	case common.Bubble:
		return bubble.Sort(content, opts)
	case common.Quick:
		return quick.Sort(content, opts)
	case common.Merge:
		return merge.Sort(content, opts)
	}
	return content
}
