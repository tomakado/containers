package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.tomakado.io/containers/tree"
)

func TestAppend(t *testing.T) {
	var (
		node  = &tree.Node[int]{Value: 1}
		child = &tree.Node[int]{Value: 2}
	)

	node.Append(child)
	assert.Equal(t, node, child.Parent)
	assert.Contains(t, node.Children, child)
}

func TestRemove(t *testing.T) {
	var (
		node  = &tree.Node[int]{Value: 1}
		child = &tree.Node[int]{Value: 2}
	)

	node.Append(child)
	child.Remove()

	assert.Nil(t, child.Parent)
	assert.NotContains(t, node.Children, child)
}

func TestDepth(t *testing.T) {
	var (
		node   = &tree.Node[int]{Value: 1}
		child1 = &tree.Node[int]{Value: 2}
		child2 = &tree.Node[int]{Value: 3}
		child3 = &tree.Node[int]{Value: 4}
	)

	node.Append(child1)
	node.Append(child2)
	child2.Append(child3)

	assert.Equal(t, 3, node.Depth())
}

func TestBFS(t *testing.T) {
	var (
		node   = &tree.Node[int]{Value: 1}
		child1 = &tree.Node[int]{Value: 2}
		child2 = &tree.Node[int]{Value: 3}
		child3 = &tree.Node[int]{Value: 4}
	)

	node.Append(child1)
	node.Append(child2)
	child2.Append(child3)

	foundNode, found := node.BFS(3)
	assert.True(t, found)
	assert.Equal(t, child2, foundNode)

	notFoundNode, notFound := node.BFS(5)
	assert.False(t, notFound)
	assert.Nil(t, notFoundNode)
}

func TestDFS(t *testing.T) {
	var (
		node   = &tree.Node[int]{Value: 1}
		child1 = &tree.Node[int]{Value: 2}
		child2 = &tree.Node[int]{Value: 3}
		child3 = &tree.Node[int]{Value: 4}
	)

	node.Append(child1)
	node.Append(child2)
	child2.Append(child3)

	foundNode, found := node.DFS(3)
	assert.True(t, found)
	assert.Equal(t, child2, foundNode)

	notFoundNode, notFound := node.DFS(5)
	assert.False(t, notFound)
	assert.Nil(t, notFoundNode)
}
