package binary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomakado/containers/binary"
)

func TestAppend(t *testing.T) {
	var (
		root  = &binary.Node[int, int]{Key: 5, Value: 5}
		node1 = &binary.Node[int, int]{Key: 3, Value: 3}
		node2 = &binary.Node[int, int]{Key: 7, Value: 7}
	)

	root.Append(node1)
	root.Append(node2)

	assert.Equal(t, root.Left.Key, 3, "Expected left node key to be 3")
	assert.Equal(t, root.Right.Key, 7, "Expected right node key to be 7")
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
}
