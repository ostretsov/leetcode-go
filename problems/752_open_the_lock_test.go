package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_openLock(t *testing.T) {
	got := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	assert.Equal(t, 6, got)

	got = openLock([]string{"8888"}, "0009")
	assert.Equal(t, 1, got)

	got = openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	assert.Equal(t, -1, got)

	got = openLock([]string{"0000"}, "8888")
	assert.Equal(t, -1, got)

	got = openLock([]string{"8888"}, "0000")
	assert.Equal(t, 0, got)
}

func Test_allNextMoves(t *testing.T) {
	ast := assert.New(t)
	got := allNextMoves("0000")
	ast.Len(got, 8)
	ast.Contains(got, "0001")
	ast.Contains(got, "0009")
	ast.Contains(got, "0010")
	ast.Contains(got, "0090")
	ast.Contains(got, "0100")
	ast.Contains(got, "0900")
	ast.Contains(got, "1000")
	ast.Contains(got, "9000")
}

func Test_allNextMovesInt(t *testing.T) {
	ast := assert.New(t)
	got := allNextMovesInt(0)
	ast.Len(got, 8)
	ast.Contains(got, 1)
	ast.Contains(got, 9)
	ast.Contains(got, 10)
	ast.Contains(got, 90)
	ast.Contains(got, 100)
	ast.Contains(got, 900)
	ast.Contains(got, 1000)
	ast.Contains(got, 9000)

	got = allNextMovesInt(1234)
	ast.Len(got, 8)
	ast.Contains(got, 2234)
	ast.Contains(got, 234)
	ast.Contains(got, 1334)
	ast.Contains(got, 1134)
	ast.Contains(got, 1244)
	ast.Contains(got, 1224)
	ast.Contains(got, 1233)
	ast.Contains(got, 1235)
}
