package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Clear(t *testing.T) {
	tests := []struct {
		name  string
		items []int
	}{
		{
			name:  "empty_queue",
			items: []int{},
		},
		{
			name:  "filled_queue",
			items: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(tt.items...)
			q.Clear()
			assert.Equal(t, 0, q.Len())
			assert.Nil(t, q.head)
			assert.Nil(t, q.tail)
		})
	}
}

func TestQueue_Contains(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		item     int
		contains bool
	}{
		{
			name:     "doesn't contain",
			items:    []int{1, 2, 3, 4},
			item:     5,
			contains: false,
		},
		{
			name:     "contains",
			items:    []int{1, 2, 3, 4},
			item:     1,
			contains: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(tt.items...)
			assert.Equal(t, tt.contains, q.Contains(tt.item))
		})
	}
}

func TestQueue_Len(t *testing.T) {
	tests := []struct {
		name        string
		items       []int
		expectedLen int
	}{
		{
			name:        "empty_queue",
			items:       []int{},
			expectedLen: 0,
		},
		{
			name:        "filled_queue_with_1_item",
			items:       []int{1},
			expectedLen: 1,
		},
		{
			name:        "filled_queue_with_2_items",
			items:       []int{1, 2},
			expectedLen: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(tt.items...)
			assert.Equal(t, tt.expectedLen, q.Len())
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name        string
		items       []int
		dequeued    int
		ok          bool
		expectedLen int
	}{
		{
			name:        "empty_queue",
			items:       []int{},
			dequeued:    0,
			ok:          false,
			expectedLen: 0,
		},
		{
			name:        "filled_queue_with_1_item",
			items:       []int{1},
			dequeued:    1,
			ok:          true,
			expectedLen: 0,
		},
		{
			name:        "filled_queue_with_2_item",
			items:       []int{1, 2},
			dequeued:    1,
			ok:          true,
			expectedLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(tt.items...)
			item, ok := q.Dequeue()

			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.dequeued, item)
			assert.Equal(t, tt.expectedLen, q.Len())
		})
	}

	t.Run("dequeue_several_times", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		q := New(items...)
		for _, item := range items {
			got, ok := q.Dequeue()
			assert.True(t, ok)
			assert.Equal(t, item, got)
		}
		_, ok := q.Dequeue()
		assert.False(t, ok)
	})
}

func TestQueue_Enqueue(t *testing.T) {
	tests := []struct {
		name  string
		items []int
	}{
		{
			name:  "one time",
			items: []int{1},
		},
		{
			name:  "two times",
			items: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New[int]()
			for i, item := range tt.items {
				q.Enqueue(item)
				assert.Equal(t, i+1, q.Len())
				assert.Equal(t, item, q.tail.value)
			}
			assert.Equal(t, tt.items[0], q.head.value)
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		ok       bool
		expected int
	}{
		{
			name:     "empty_queue",
			items:    []int{},
			ok:       false,
			expected: 0,
		},
		{
			name:     "filled_queue",
			items:    []int{1, 2, 3},
			ok:       true,
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(tt.items...)
			actual, ok := q.Peek()
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestQueue_Slice(t *testing.T) {
	tests := []struct {
		name  string
		items []int
	}{
		{
			name:  "empty_queue",
			items: []int{},
		},
		{
			name:  "filled_queue",
			items: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(tt.items...)
			actual := q.Slice()
			assert.Equal(t, tt.items, actual)
		})
	}
}
