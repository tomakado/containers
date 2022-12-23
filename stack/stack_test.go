package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomakado/containers/stack"
)

func TestStack(t *testing.T) {
	s := stack.New[int](10)

	t.Run("IsEmpty", func(t *testing.T) {
		assert.True(t, s.IsEmpty(), "Expected stack to be empty")
	})

	t.Run("Push and Peek", func(t *testing.T) {
		s.Push(1)
		s.Push(2)
		s.Push(3)

		val, ok := s.Peek()

		assert.Equal(t, 3, val, "Expected 3 at the top of the stack")
		assert.True(t, ok, "Expected Peek to return true")
	})

	t.Run("Peek empty stack", func(t *testing.T) {
		var (
			s       = stack.New[int](10)
			val, ok = s.Peek()
		)

		assert.False(t, ok, "Expected Pop to return false")
		assert.Equal(t, 0, val, "Expected 0 to be popped off the stack")
	})

	t.Run("Len", func(t *testing.T) {
		assert.Equal(t, 3, s.Len(), "Expected stack to have length 3")
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("non-empty stack", func(t *testing.T) {
			val, ok := s.Pop()

			assert.Equal(t, 3, val, "Expected 3 to be popped off the stack")
			assert.True(t, ok, "Expected Pop to return true")
			assert.Equal(t, 2, s.Len(), "Expected stack to have length 2")
		})

		t.Run("empty stack", func(t *testing.T) {
			var (
				s       = stack.New[int](10)
				val, ok = s.Pop()
			)

			assert.False(t, ok, "Expected Pop to return false")
			assert.Equal(t, 0, val, "Expected 0 to be popped off the stack")
		})
	})
}
