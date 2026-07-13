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
	node := &TreeNode[T]{
		Value: value,
		Left:  nil,
		Right: nil,
	}
	if bt.Root == nil {
		bt.Root = node
		return
	}

	currentNode := bt.Root

	for currentNode != nil {
		if node.Value < currentNode.Value {
			if currentNode.Left == nil {
				//Set the value of the next thing
				currentNode.Left = node
				return
			}
			currentNode = currentNode.Left
		} else if node.Value > currentNode.Value {
			if currentNode.Right == nil {
				currentNode.Right = node
				return
			}
			currentNode = currentNode.Right
		} else {
			return
		}
	}
}

func (bt *BinaryTree[T]) Search(target T) (found bool) {
	currentNode := bt.Root

	for currentNode != nil {
		if target < currentNode.Value {
			currentNode = currentNode.Left
		} else if target > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			return true
		}
	}

	return false
}

func (bt *BinaryTree[T]) InOrderTraverse() []T {
	elements := make([]T, 0)

	currentNode := bt.Root

	for currentNode != nil {
		elements = append(elements, currentNode.Value)
		currentNode = currentNode.Left
	}

	return elements
}

func inOrderTraverseHelper(elements []T, currentNode TreeNode[T]) {
	elements = append(elements, currentNode.Value)
	if currentNode.Left != nil {
		inOrderTraverseHelper(elements, *currentNode.Left)
	}

	if currentNode.Right != nil {
		inOrderTraverseHelper(elements, *currentNode.Right)
	}

	return
}
