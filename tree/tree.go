package tree

import "github.com/tomakado/containers/queue"

// Node is a node of N-ary tree.
// Implementation is based on code by Ilija Eftimov (URL: https://ieftimov.com/posts/golang-datastructures-trees/).
type Node[T comparable] struct {
	Value    T
	Parent   *Node[T]
	Children []*Node[T]
}

func (n *Node[T]) Append(c *Node[T]) {
	c.Parent = n
	n.Children = append(n.Children, c)
}

func (n *Node[T]) Remove() {
	for i, sibling := range n.Parent.Children {
		if sibling == n {
			n.Parent.Children = append(n.Parent.Children[:i], n.Parent.Children[i+1:]...)
			break
		}
	}

	if len(n.Children) != 0 {
		for _, child := range n.Children {
			child.Parent = nil
		}
	}

	n.Parent = nil
}

func (n *Node[T]) Depth() int {
	if len(n.Children) == 0 {
		return 1
	}

	max := 0

	for _, c := range n.Children {
		d := c.Depth()
		if d > max {
			max = d
		}
	}

	return max + 1
}

func (n *Node[T]) BFS(value T) (*Node[T], bool) {
	q := queue.New(n)

	for q.Len() > 0 {
		node, ok := q.Dequeue()
		if !ok {
			break
		}

		if node.Value == value {
			return node, true
		}

		for _, c := range node.Children {
			q.Enqueue(c)
		}
	}

	return nil, false
}

func (n *Node[T]) DFS(value T) (*Node[T], bool) {
	if n.Value == value {
		return n, true
	}

	if len(n.Children) == 0 {
		return nil, false
	}

	for _, c := range n.Children {
		node, ok := c.DFS(value)
		if ok {
			return node, true
		}
	}

	return nil, false
}
