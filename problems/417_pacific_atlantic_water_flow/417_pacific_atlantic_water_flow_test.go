package _17_pacific_atlantic_water_flow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pacificAtlantic(t *testing.T) {
	ast := assert.New(t)

	t.Run("Example1", func(t *testing.T) {
		got := pacificAtlantic([][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		})

		ast.Equal([]int{0, 4}, got[0])
	})

	t.Run("Example2", func(t *testing.T) {
		got := pacificAtlantic([][]int{
			{1, 2, 3},
			{3, 5, 2},
			{4, 4, 5},
		})

		ast.Equal([]int{0, 4}, got[0])
	})
}
