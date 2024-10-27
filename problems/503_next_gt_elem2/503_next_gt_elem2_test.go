package _03_next_gt_elem2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextGreaterElements(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{2, 3, 4, -1, 4}, nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
