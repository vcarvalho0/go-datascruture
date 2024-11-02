package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	right *TreeNode
	left  *TreeNode
}

func New(value int) *TreeNode {
	tree := &TreeNode{}
	tree.value = value

	return tree
}

// pre-order traversal
func (root *TreeNode) printTree() {
	if root == nil {
		return
	}

	root.left.printTree()
	fmt.Printf("%d ", root.value)
	root.right.printTree()
}

func (root *TreeNode) invert() *TreeNode {
	if root == nil {
		return nil
	}

	// swap
	tmp := root.left
	root.left = root.right
	root.right = tmp

	root.left.invert()
	root.right.invert()

	return root
}

func main() {
	root := New(0)
	root.left = New(1)
	root.right = New(2)
	root.left.left = New(3)
	root.left.right = New(4)
	root.right.left = New(5)
	root.right.right = New(6)

	root.invert()
	root.printTree()
}
