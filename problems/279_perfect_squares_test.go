package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"n=1", 1, 1},
		{"n=12", 12, 3},
		{"n=13", 13, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numSquares(tt.n), "numSquares(%v)", tt.n)
		})
	}
}
