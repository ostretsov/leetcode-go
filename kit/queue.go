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

type ReuseMemQueue[T any] struct {
	data []T
	head int
	tail int
}

func NewReuseMemQueue[T any]() *ReuseMemQueue[T] {
	return &ReuseMemQueue[T]{head: -1, tail: -1}
}

func (q *ReuseMemQueue[T]) Enqueue(v T) {
	if q.head == -1 {
		q.head++
	}
	q.tail++
	q.data = append(q.data, v)
}

func (q *ReuseMemQueue[T]) Dequeue() (T, bool) {
	var v T
	if q.Empty() {
		return v, false
	}
	v = q.data[q.head]
	var e T
	q.data[q.head] = e
	if q.Len() == 1 {
		q.head = -1
		q.tail = -1
	} else {
		q.head++
	}
	return v, true
}

func (q *ReuseMemQueue[T]) Empty() bool {
	return q.head == -1
}

func (q *ReuseMemQueue[T]) Len() int {
	if q.Empty() {
		return 0
	}
	return q.tail - q.head + 1
}
