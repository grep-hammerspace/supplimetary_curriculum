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

	//   5
	//  / \
	// 3   7

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
