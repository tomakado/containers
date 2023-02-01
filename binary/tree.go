package binary

import (
	"golang.org/x/exp/constraints"
)

type Node[K, V constraints.Ordered] struct {
	Key         K
	Value       V
	Left, Right *Node[K, V]
}

func (n *Node[K, V]) Append(node *Node[K, V]) {
	if node.Value < n.Value {
		if n.Left == nil {
			n.Left = node
			return
		}

		n.Left.Append(node)

		return
	}

	if n.Right == nil {
		n.Right = node
		return
	}

	n.Right.Append(node)
}

func (n *Node[K, V]) PreOrder(f func(node *Node[K, V])) {
	if n == nil {
		return
	}

	f(n)
	n.Left.PreOrder(f)
	n.Right.PreOrder(f)
}

func (n *Node[K, V]) InOrder(f func(node *Node[K, V])) {
	if n == nil {
		return
	}

	n.Left.InOrder(f)
	f(n)
	n.Right.InOrder(f)
}

func (n *Node[K, V]) PostOrder(f func(node *Node[K, V])) {
	if n == nil {
		return
	}

	n.Left.PostOrder(f)
	n.Right.PostOrder(f)
	f(n)
}

func (n *Node[K, V]) Search(key K) (*Node[K, V], bool) {
	if n == nil {
		return nil, false
	}

	if key < n.Key {
		return n.Left.Search(key)
	}

	if key > n.Key {
		return n.Right.Search(key)
	}

	return n, true
}

func (n *Node[K, V]) Remove(key K) *Node[K, V] {
	if n == nil {
		return nil
	}

	if key < n.Key {
		n.Left = n.Left.Remove(key)
		return n
	}

	if key > n.Key {
		n.Right = n.Right.Remove(key)
		return n
	}

	return n.removeSelf(key)
}

func (n *Node[K, V]) removeSelf(key K) *Node[K, V] {
	if n.Left == nil && n.Right == nil {
		n = nil
		return nil
	}

	if n.Left == nil {
		n = n.Right
		return n
	}

	if n.Right == nil {
		n = n.Left
		return n
	}

	leftMostRightSide := n.Right

	for {
		if leftMostRightSide == nil || leftMostRightSide.Left == nil {
			break
		}

		leftMostRightSide = leftMostRightSide.Left
	}

	n.Key, n.Value = leftMostRightSide.Key, leftMostRightSide.Value
	n.Right.Remove(n.Key)

	return n
}
