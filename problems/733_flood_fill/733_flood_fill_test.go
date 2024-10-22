package _33_flood_fill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_floodFill(t *testing.T) {
	ast := assert.New(t)

	img1 := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	got := floodFill(img1, 1, 1, 2)
	ast.Equal([]int{2, 2, 2}, got[0])
	ast.Equal([]int{2, 2, 0}, got[1])
	ast.Equal([]int{2, 0, 1}, got[2])
}
