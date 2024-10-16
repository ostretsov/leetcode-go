package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findTargetSumWays(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(5, findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
