package _41_keys_rooms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canVisitAllRooms(t *testing.T) {
	ast := assert.New(t)

	ast.True(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	ast.False(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
