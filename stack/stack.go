package stack

type Stack[T any] struct {
	top      int
	elements []T
}

func New[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		top:      -1,
		elements: make([]T, 0, capacity),
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Push(e T) {
	s.elements = append(s.elements, e)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return e, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.elements)
}
