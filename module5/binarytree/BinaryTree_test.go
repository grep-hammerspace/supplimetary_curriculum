package binarytree

import (
	"testing"
)

func TestInsertIntoEmptyBinaryTree(t *testing.T) {
	bt := New[int]()

	bt.Insert(5)

	if bt.Root.Value != 5 {
		t.Errorf("InsertIntoEmptyBinaryTree failed, expected value for Head TreeNode = 5")
	}
}

func TestInsertIntoBinaryTreeLeftAndRight(t *testing.T) {
	bt := New[int]()
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(8)

	//   5
	//  / \
	// 3   7
	//      \
	//       8

	if bt.Root.Value != 5 {
		t.Errorf("Insert into empty tree should have head 5")
	}

	if bt.Root.Left != nil && bt.Root.Left.Value != 3 {
		t.Errorf("Insert into left tree should have head 3")
	}

	if bt.Root.Right != nil && bt.Root.Right.Value != 7 {
		t.Errorf("Insert into right tree should have head 7")
	}
}

func TestInOrderTraverse(t *testing.T) {
	bt := New[int]()
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(8)

	//   5
	//  / \
	// 3   7
	//      \
	//       8

	sortedElements := bt.InOrderTraverse()

	if len(sortedElements) != 4 {
		t.Errorf("InOrderTraverse failed, expected 4 elements, got %d", len(sortedElements))
	}

	expectedElementOrder := make([]int, 4)
	expectedElementOrder[0] = 3
	expectedElementOrder[1] = 5
	expectedElementOrder[2] = 7
	expectedElementOrder[3] = 8

	for i := 0; i < 4; i++ {
		if sortedElements[i] != expectedElementOrder[i] {
			t.Errorf("Invalid order for InOrder traverse")
		}
	}

}
