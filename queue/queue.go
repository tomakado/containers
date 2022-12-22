package queue

type Queue[T comparable] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func New[T comparable](items ...T) *Queue[T] {
	q := &Queue[T]{}
	for _, item := range items {
		q.Enqueue(item)
	}
	return q
}

func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
	q.length = 0
}

func (q *Queue[T]) Contains(item T) bool {
	n := q.head
	for n.next != nil {
		if n.value == item {
			return true
		}
		n = n.next
	}
	return false
}
func (q *Queue[T]) Dequeue() (T, bool) {
	n := q.head
	if n == nil {
		return node[T]{}.value, false
	}
	next := q.head.next
	q.head = next
	q.length--

	return n.value, true
}

func (q *Queue[T]) Enqueue(item T) {
	n := &node[T]{value: item}
	q.length++
	if q.head == nil {
		q.head = n
		q.tail = q.head
		return
	}
	q.tail.next = n
	q.tail = q.tail.next
}

func (q *Queue[T]) Peek() (T, bool) {
	n := q.head
	if n == nil {
		return node[T]{}.value, false
	}
	return n.value, true
}

func (q *Queue[T]) Len() int {
	return q.length
}

func (q *Queue[T]) Slice() []T {
	slice := make([]T, 0, q.Len())
	for q.Len() > 0 {
		item, _ := q.Dequeue()
		slice = append(slice, item)
	}
	return slice
}

type node[T comparable] struct {
	value T
	next  *node[T]
}
