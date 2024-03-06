package binary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.tomakado.io/containers/binary"
)

func TestAppend(t *testing.T) {
	t.Run("no children", func(t *testing.T) {
		var (
			root  = &binary.Node[int, int]{Key: 5, Value: 5}
			node1 = &binary.Node[int, int]{Key: 3, Value: 3}
			node2 = &binary.Node[int, int]{Key: 7, Value: 7}
		)

		root.Append(node1)
		root.Append(node2)

		assert.Equal(t, root.Left.Key, 3, "Expected left node key to be 3")
		assert.Equal(t, root.Right.Key, 7, "Expected right node key to be 7")
	})

	t.Run("left children", func(t *testing.T) {
		var (
			root  = &binary.Node[int, int]{Key: 5, Value: 5}
			node1 = &binary.Node[int, int]{Key: 3, Value: 3}
			node2 = &binary.Node[int, int]{Key: 1, Value: 1}
		)

		root.Append(node1)
		root.Append(node2)

		assert.Equal(t, root.Left.Key, 3, "Expected left node key to be 3")
		assert.Nil(t, root.Right, "Expected right node to be nil")
	})

	t.Run("right children", func(t *testing.T) {
		var (
			root = &binary.Node[int, int]{Key: 5, Value: 5}

			node1 = &binary.Node[int, int]{Key: 3, Value: 3}
			node2 = &binary.Node[int, int]{Key: 7, Value: 7}
			node3 = &binary.Node[int, int]{Key: 9, Value: 9}
		)

		root.Append(node1)
		root.Append(node2)
		root.Append(node3)

		assert.Equal(t, root.Left.Key, 3, "Expected left node key to be 3")
		assert.Equal(t, root.Right.Key, 7, "Expected left node key to be 7")
		assert.Equal(t, root.Right.Right.Key, 9, "Expected left node key to be 9")
	})
}

func TestPreOrder(t *testing.T) {
	var (
		root  = &binary.Node[int, int]{Key: 5, Value: 5}
		node1 = &binary.Node[int, int]{Key: 3, Value: 3}
		node2 = &binary.Node[int, int]{Key: 7, Value: 7}
	)

	root.Append(node1)
	root.Append(node2)

	var keys []int

	root.PreOrder(func(node *binary.Node[int, int]) {
		keys = append(keys, node.Key)
	})

	assert.Equal(t, []int{5, 3, 7}, keys, "Expected pre-order traversal to be [5, 3, 7]")
}

func TestInOrder(t *testing.T) {
	var (
		root  = &binary.Node[int, int]{Key: 5, Value: 5}
		node1 = &binary.Node[int, int]{Key: 3, Value: 3}
		node2 = &binary.Node[int, int]{Key: 7, Value: 7}
	)

	root.Append(node1)
	root.Append(node2)

	var keys []int

	root.InOrder(func(node *binary.Node[int, int]) {
		keys = append(keys, node.Key)
	})

	assert.Equal(t, []int{3, 5, 7}, keys, "Expected pre-order traversal to be [3, 5, 7]")
}

func TestPostOrder(t *testing.T) {
	var (
		root  = &binary.Node[int, int]{Key: 5, Value: 5}
		node1 = &binary.Node[int, int]{Key: 3, Value: 3}
		node2 = &binary.Node[int, int]{Key: 7, Value: 7}
	)

	root.Append(node1)
	root.Append(node2)

	var keys []int

	root.PostOrder(func(node *binary.Node[int, int]) {
		keys = append(keys, node.Key)
	})

	assert.Equal(t, []int{3, 7, 5}, keys, "Expected pre-order traversal to be [3, 7, 5]")
}

func TestSearch(t *testing.T) {
	var (
		root  = &binary.Node[int, int]{Key: 5, Value: 5}
		node1 = &binary.Node[int, int]{Key: 3, Value: 3}
		node2 = &binary.Node[int, int]{Key: 7, Value: 7}
	)

	root.Append(node1)
	root.Append(node2)

	node, found := root.Search(7)
	assert.True(t, found, "Expected to find node with key 7")
	assert.Equal(t, 7, node.Key, "Expected key to be 7")
	assert.Equal(t, 7, node.Value, "Expected value to be 7")

	node, found = root.Search(10)
	assert.False(t, found, "Expected not to find node with key 10")
	assert.Nil(t, node, "Expected node to be nil")
}

func TestRemove(t *testing.T) {
	t.Run("no children", func(t *testing.T) {
		var (
			root  = &binary.Node[int, int]{Key: 5, Value: 5}
			node1 = &binary.Node[int, int]{Key: 3, Value: 3}
			node2 = &binary.Node[int, int]{Key: 7, Value: 7}
		)

		root.Append(node1)
		root.Append(node2)

		root.Remove(3)
		node, found := root.Search(3)
		assert.False(t, found, "Expected not to find node with key 3")
		assert.Nil(t, node, "Expected node to be nil")

		root.Remove(7)
		node, found = root.Search(7)
		assert.False(t, found, "Expected not to find node with key 7")
		assert.Nil(t, node, "Expected node to be nil")
	})

	t.Run("left is nil", func(t *testing.T) {
		var (
			root  = &binary.Node[int, int]{Key: 5, Value: 5}
			node1 = &binary.Node[int, int]{Key: 6, Value: 6}
			node2 = &binary.Node[int, int]{Key: 7, Value: 7}
		)

		root.Append(node1)
		root.Append(node2)

		root.Remove(6)
		node, found := root.Search(6)
		assert.False(t, found, "Expected not to find node with key 6")
		assert.Nil(t, node, "Expected node to be nil")

		root.Remove(7)
		node, found = root.Search(7)
		assert.False(t, found, "Expected not to find node with key 7")
		assert.Nil(t, node, "Expected node to be nil")
	})

	t.Run("right is nil", func(t *testing.T) {
		var (
			root  = &binary.Node[int, int]{Key: 5, Value: 5}
			node1 = &binary.Node[int, int]{Key: 4, Value: 4}
			node2 = &binary.Node[int, int]{Key: 3, Value: 3}
		)

		root.Append(node1)
		root.Append(node2)

		root.Remove(4)
		node, found := root.Search(4)
		assert.False(t, found, "Expected not to find node with key 4")
		assert.Nil(t, node, "Expected node to be nil")
	})

	t.Run("both children are not nil", func(t *testing.T) {
		var (
			root  = &binary.Node[int, int]{Key: 5, Value: 5}
			node1 = &binary.Node[int, int]{Key: 4, Value: 4}
			node2 = &binary.Node[int, int]{Key: 6, Value: 6}
		)

		root.Append(node1)
		root.Append(node2)

		root.Remove(5)
		node, found := root.Search(5)
		assert.False(t, found, "Expected not to find node with key 5")
		assert.Nil(t, node, "Expected node to be nil")
	})
}
