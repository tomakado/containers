package set

type Set[T comparable] map[T]struct{}

func New[T comparable](elements ...T) Set[T] {
	s := Set[T]{}

	for _, e := range elements {
		s.Add(e)
	}

	return s
}

func (s Set[T]) Add(e T) {
	s[e] = struct{}{}
}

func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Remove(e T) {
	delete(s, e)
}

func (s Set[T]) Slice() []T {
	var slice []T

	for e := range s {
		slice = append(slice, e)
	}

	return slice
}
