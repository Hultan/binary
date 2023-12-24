package main

import (
	"cmp"
	"fmt"
)

// Node structure to store the data and the
// pointers to the left and right children
type Node[T cmp.Ordered] struct {
	key         T
	left, right *Node[T]
	height      int
}

// NewTree creates a new AVL tree root node
func NewTree[T cmp.Ordered](key T) *Node[T] {
	return &Node[T]{key: key}
}

// Insert a new node into the AVL Tree
func (n *Node[T]) Insert(key T) *Node[T] {
	if n == nil {
		return &Node[T]{key: key, height: 1}
	}
	if key < n.key {
		n.left = n.left.Insert(key)
	} else if key > n.key {
		n.right = n.right.Insert(key)
	} else {
		return n
	}

	n.height = 1 + max(n.left.Height(), n.right.Height())
	balanceFactor := n.getBalanceFactor()

	if balanceFactor > 1 {
		if key < n.left.key {
			return n.rotateRight()
		} else if key > n.left.key {
			n.left = n.left.rotateLeft()
			return n.rotateRight()
		}
	}

	if balanceFactor < -1 {
		if key > n.right.key {
			return n.rotateLeft()
		} else if key < n.right.key {
			n.right = n.right.rotateRight()
			return n.rotateLeft()
		}
	}

	return n
}

// Delete a node from the AVL Tree
func (n *Node[T]) Delete(key T) *Node[T] {
	// Searching Node
	if n == nil {
		return n
	}
	if key < n.key {
		n.left = n.left.Delete(key)
	} else if key > n.key {
		n.right = n.right.Delete(key)
	} else {
		if n.left == nil || n.right == nil {
			temp := n.left
			if temp == nil {
				temp = n.right
			}
			if temp == nil {
				temp = n
				n = nil
			} else {
				*n = *temp
			}
		} else {
			temp := n.right.nodeWithMinimumValue()
			n.key = temp.key
			n.right = n.right.Delete(temp.key)
		}
	}
	if n == nil {
		return n
	}
	n.height = 1 + max(n.left.Height(), n.right.Height())
	balanceFactor := n.getBalanceFactor()

	if balanceFactor > 1 {
		if n.left.getBalanceFactor() >= 0 {
			return n.rotateRight()
		} else {
			n.left = n.left.rotateLeft()
			return n.rotateRight()
		}
	}
	if balanceFactor < -1 {
		if n.right.getBalanceFactor() <= 0 {
			return n.rotateLeft()
		} else {
			n.right = n.right.rotateRight()
			return n.rotateLeft()
		}
	}
	return n
}

// Print the AVL tree
func (n *Node[T]) Print(indent string, last bool) {
	if n != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "   "
		} else {
			fmt.Print("L----")
			indent += "|  "
		}
		fmt.Println(n.key)
		n.left.Print(indent, false)
		n.right.Print(indent, true)
	}
}

// Height calculates the height of the tree/node
func (n *Node[T]) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

// Performs a right rotation on the node
func (n *Node[T]) rotateRight() *Node[T] {
	l := n.left
	r := l.right
	l.right = n
	n.left = r
	n.height = max(n.left.Height(), n.right.Height()) + 1
	l.height = max(l.left.Height(), l.right.Height()) + 1
	return l
}

// Performs a left rotation on the node
func (n *Node[T]) rotateLeft() *Node[T] {
	r := n.right
	l := r.left
	r.left = n
	n.right = l
	r.height = max(r.left.Height(), r.right.Height()) + 1
	n.height = max(n.left.Height(), n.right.Height()) + 1
	return r
}

// Calculates the balance factor
// of the node
func (n *Node[T]) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.Height() - n.right.Height()
}

// Fetches the Node with minimum
// value from the AVL tree
func (n *Node[T]) nodeWithMinimumValue() *Node[T] {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}
