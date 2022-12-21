package queue

type Queue[T comparable] struct {
	head *node[T]
	tail *node[T]
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

	return n.value, true
}

func (q *Queue[T]) Enqueue(item T) {
	n := &node[T]{value: item}
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

type node[T comparable] struct {
	value T
	next  *node[T]
}
