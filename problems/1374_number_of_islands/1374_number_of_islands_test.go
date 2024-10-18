package _374_number_of_islands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numIslands(t *testing.T) {
	in1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	assert.Equal(t, 1, numIslands(in1))

	in2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	assert.Equal(t, 3, numIslands(in2))

	in3 := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	assert.Equal(t, 1, numIslands(in3))
}
