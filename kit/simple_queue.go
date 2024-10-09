package kit

type SimpleQueue[T any] struct {
	data []T
}

func NewSimpleQueue[T any]() *SimpleQueue[T] {
	return &SimpleQueue[T]{}
}

func (q *SimpleQueue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

func (q *SimpleQueue[T]) Dequeue() (T, bool) {
	var v T
	if q.Empty() {
		return v, false
	}
	v = q.data[0]
	q.data = q.data[1:]
	return v, true
}

func (q *SimpleQueue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *SimpleQueue[T]) Len() int {
	return len(q.data)
}
