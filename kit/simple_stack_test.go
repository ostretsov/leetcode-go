package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleStack(t *testing.T) {
	ast := assert.New(t)

	stack := SimpleStack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	ast.False(stack.Empty())

	e, ok := stack.Pop()
	ast.True(ok)
	ast.Equal(3, e)

	e, ok = stack.Top()
	ast.True(ok)
	ast.Equal(2, e)

	_, ok = stack.Pop()
	ast.True(ok)

	_, ok = stack.Pop()
	ast.True(ok)

	_, ok = stack.Pop()
	ast.False(ok)
}
