package _25_stack_using_queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStack_Push(t *testing.T) {
	ast := assert.New(t)

	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	ast.Equal(3, stack.Pop())
	ast.Equal(2, stack.Top())

	stack.Push(4)
	ast.Equal(4, stack.Top())
	ast.Equal(4, stack.Pop())
	ast.Equal(2, stack.Pop())
	ast.Equal(1, stack.Pop())
}
