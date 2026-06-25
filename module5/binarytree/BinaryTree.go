package binarytree

import "cmp"

type TreeNode[T cmp.Ordered] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BinaryTree[T cmp.Ordered] struct {
	Root *TreeNode[T]
}

func New[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{
		Root: nil,
	}
}

func (bt *BinaryTree[T]) Insert(value T) {
	bt.insertHelper(value)
}

func (bt *BinaryTree[T]) insertHelper(value T) {
	node := &TreeNode[T]{
		Value: value,
		Left:  nil,
		Right: nil,
	}
	if bt.Root == nil {
		bt.Root = node
	}
	currentNode := bt.Root

	for currentNode != nil {
		if node.Value < currentNode.Value {
			// Smaller so go left
			currentNode = currentNode.Left
		} else if node.Value > currentNode.Value {
			// Bigger so go right
			currentNode = currentNode.Right
		} else {
			return
		}
	}

	currentNode := node

}
