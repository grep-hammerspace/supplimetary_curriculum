package main

import (
	"fmt"

	bst "github.com/grep-hammerspace/coding-curriculum/module5/binarytree"
)

func main() {
	binarytree := bst.New[int]()

	binarytree.Insert(4)
	binarytree.Insert(3)
	binarytree.Insert(5)
	binarytree.Insert(6)

	sortedElements := binarytree.InOrderTraverse()

	fmt.Println(sortedElements)
}
