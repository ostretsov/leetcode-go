package _56_132_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_find132pattern(t *testing.T) {
	assert.False(t, find132pattern([]int{1, 2, 3, 4}))

	assert.True(t, find132pattern([]int{3, 1, 4, 2}))
	assert.True(t, find132pattern([]int{3, 1, 5, 4}))
	assert.True(t, find132pattern([]int{-1, 3, 2, 0}))
	assert.True(t, find132pattern([]int{3, 5, 0, 3, 4}))
}
