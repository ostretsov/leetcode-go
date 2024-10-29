package _762_buildins_with_ocean_view

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findBuildings(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([]int{0, 2, 3}, findBuildings([]int{4, 2, 3, 1}))
	ast.Equal([]int{0, 1, 2, 3}, findBuildings([]int{4, 3, 2, 1}))
	ast.Equal([]int{3}, findBuildings([]int{1, 3, 2, 4}))
}
