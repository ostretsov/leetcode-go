package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_carFleet(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(3, carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	ast.Equal(1, carFleet(10, []int{3}, []int{3}))
	ast.Equal(1, carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
	ast.Equal(2, carFleet(10, []int{6, 8}, []int{3, 2}))
	ast.Equal(1, carFleet(10, []int{0, 4, 2}, []int{2, 1, 3}))
	ast.Equal(6, carFleet(10, []int{8, 3, 7, 4, 6, 5}, []int{4, 4, 4, 4, 4, 4}))
	ast.Equal(1, carFleet(10, []int{2, 4}, []int{3, 2}))
}
