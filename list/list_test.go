package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomakado/containers/list"
)

func TestList(t *testing.T) {
	t.Run("single element list", func(t *testing.T) {
		var (
			list = list.NewList[string]()
			e    = list.PushBack("foo")
		)

		assert.Equal(t, 1, list.Len())
		assert.Equal(t, e, list.Front())

		list.MoveToFront(e)
		assert.Equal(t, 1, list.Len())
		assert.Equal(t, e, list.Front())

		list.MoveToBack(e)
		assert.Equal(t, 1, list.Len())
		assert.Equal(t, e, list.Front())

		list.Remove(e)
		assert.Equal(t, 0, list.Len())
		assert.Nil(t, list.Front())
	})

	t.Run("bigger list", func(t *testing.T) {
		var (
			list = list.NewList[int]()
			e1   = list.PushBack(2)
			e2   = list.PushBack(1)
			e3   = list.PushBack(1)
		)

		assert.Equal(t, 3, list.Len())
		assert.Equal(t, e1, list.Front())
		assert.Equal(t, e3, list.Back())

		list.MoveToFront(e2)
		assert.Equal(t, 3, list.Len())
		assert.Equal(t, e2, list.Front())
		assert.Equal(t, e3, list.Back())

		list.MoveToBack(e2)
		assert.Equal(t, 3, list.Len())
		assert.Equal(t, e1, list.Front())
		assert.Equal(t, e2, list.Back())

		list.Remove(e2)
		assert.Equal(t, 2, list.Len())
		assert.Equal(t, e1, list.Front())
		assert.Equal(t, e3, list.Back())
	})

	t.Run("iteration", func(t *testing.T) {
		var (
			list = list.NewList[int]()
		)

		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)

		var (
			i   = 0
			sum = 0
		)

		for e := list.Front(); e != nil; e = e.Next() {
			sum += e.Value
			i++
		}

		assert.Equal(t, 3, i)
		assert.Equal(t, 6, sum)
	})

	t.Run("clear all elements by iterating", func(t *testing.T) {
		var (
			list = list.NewList[int]()
		)

		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)

		for e := list.Front(); e != nil; e = e.Next() {
			toDelete := *e
			list.Remove(&toDelete)
		}

		assert.Equal(t, 0, list.Len())
	})

	t.Run("extend", func(t *testing.T) {
		var (
			l1 = list.NewList[int]()
			l2 = list.NewList[int]()
			l3 = list.NewList[int]()
		)

		l1.PushBack(1)
		l1.PushBack(2)
		l1.PushBack(3)

		l2.PushBack(4)
		l2.PushBack(5)

		l3.PushBackList(l1)
		assert.Equal(t, 3, l3.Len())

		l3.PushBackList(l2)
		assert.Equal(t, 5, l3.Len())

		l3 = list.NewList[int]()
		l3.PushFrontList(l1)
		assert.Equal(t, 3, l3.Len())

		l3.PushFrontList(l2)
		assert.Equal(t, 5, l3.Len())
	})

	t.Run("remove", func(t *testing.T) {
		var (
			list = list.NewList[int]()
			e1   = list.PushBack(1)
			e2   = list.PushBack(2)
		)

		list.Remove(e1)
		assert.Equal(t, 1, list.Len())
		assert.Equal(t, e2, list.Front())

		list.Remove(e2)
		assert.Equal(t, 0, list.Len())
		assert.Nil(t, list.Front())
	})

	t.Run("move", func(t *testing.T) {
		var (
			list = list.NewList[int]()
			e1   = list.PushBack(1)
			e2   = list.PushBack(2)
		)

		list.MoveToFront(e2)
		assert.Equal(t, 2, list.Len())
		assert.Equal(t, e2, list.Front())
		assert.Equal(t, e1, list.Back())

		list.MoveToBack(e2)
		assert.Equal(t, 2, list.Len())
		assert.Equal(t, e1, list.Front())
		assert.Equal(t, e2, list.Back())

		list.MoveBefore(e1, e2)
		assert.Equal(t, 2, list.Len())
		assert.Equal(t, e1, list.Front())
		assert.Equal(t, e2, list.Back())

		list.MoveAfter(e1, e2)
		assert.Equal(t, 2, list.Len())
		assert.Equal(t, e2, list.Front())
		assert.Equal(t, e1, list.Back())
	})

	t.Run("zero list", func(t *testing.T) {
		var (
			list = list.NewList[int]()
		)

		assert.Equal(t, 0, list.Len())
		assert.Nil(t, list.Front())
		assert.Nil(t, list.Back())
	})
}
