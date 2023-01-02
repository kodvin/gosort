package sort

import (
	"github.com/tobe-decided/gosort/pkg/sort/bubble"
	"github.com/tobe-decided/gosort/pkg/sort/common"
	"github.com/tobe-decided/gosort/pkg/sort/quick"
)

func Sort(s common.SortType, content []int) []int {
	switch s {
	case common.Bubble:
		return bubble.Sort(content, nil)
	case common.Quick:
		return quick.Sort(content, nil)
	}
	return content
}
