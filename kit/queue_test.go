package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleQueue(t *testing.T) {
	ast := assert.New(t)

	q := NewSimpleQueue[int]()
	ast.True(q.Empty())

	q.Enqueue(1)
	ast.False(q.Empty())

	q.Enqueue(2)
	ast.Equal(2, q.Len())

	v, ok := q.Dequeue()
	ast.True(ok)
	ast.Equal(1, v)

	v, ok = q.Dequeue()
	ast.True(ok)
	ast.Equal(2, v)

	v, ok = q.Dequeue()
	ast.False(ok)
	ast.Equal(0, v)

	ast.Equal(0, q.Len())

	q.Enqueue(3)
	ast.Equal(1, q.Len())

	v, ok = q.Dequeue()
	ast.True(ok)
	ast.Equal(3, v)
}

func TestReuseMemQueue(t *testing.T) {
	ast := assert.New(t)

	q := NewReuseMemQueue[int]()
	ast.True(q.Empty())

	q.Enqueue(1)
	ast.False(q.Empty())

	q.Enqueue(2)
	ast.Equal(2, q.Len())

	v, ok := q.Dequeue()
	ast.True(ok)
	ast.Equal(1, v)

	v, ok = q.Dequeue()
	ast.True(ok)
	ast.Equal(2, v)

	v, ok = q.Dequeue()
	ast.False(ok)
	ast.Equal(0, v)

	ast.Equal(0, q.Len())
}
