package _55_min_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	ast.Equal(-3, obj.GetMin())

	obj.Pop()
	ast.Equal(0, obj.Top())
	ast.Equal(-2, obj.GetMin())
}
