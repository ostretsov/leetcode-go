package _96_next_grater_elem1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextGreaterElement(t *testing.T) {
	assert.Equal(t, []int{-1, 3, -1}, nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	assert.Equal(t, []int{3, -1}, nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
