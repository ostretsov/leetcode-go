package kit

type SimpleStack[T any] struct {
	data []T
}

func (s *SimpleStack[T]) Clear() {
	for !s.Empty() {
		s.Pop()
	}
}

func (s *SimpleStack[T]) Push(e T) {
	s.data = append(s.data, e)
}

func (s *SimpleStack[T]) Pop() (T, bool) {
	t, ok := s.Top()
	if ok {
		s.data = s.data[:len(s.data)-1]
	}
	return t, ok
}

func (s *SimpleStack[T]) MustPop() T {
	t, ok := s.Pop()
	if !ok {
		panic("can't pop")
	}
	return t
}

func (s *SimpleStack[T]) Top() (T, bool) {
	var t T
	if s.Empty() {
		return t, false
	}
	t = s.data[len(s.data)-1]
	return t, true
}

func (s *SimpleStack[T]) MustTop() T {
	t, ok := s.Top()
	if !ok {
		panic("can't pop")
	}
	return t
}

func (s *SimpleStack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *SimpleStack[T]) Len() int {
	return len(s.data)
}
