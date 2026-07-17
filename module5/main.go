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

	//       4
	//      / \
	//     3   5
	//          \
	//           6

	preorderedElements := binarytree.PreOrderTraverse()
	sortedElements := binarytree.InOrderTraverse()
	postorderedElements := binarytree.PostOrderTraverse()

	fmt.Println(sortedElements)
	fmt.Println(preorderedElements)
	fmt.Println(postorderedElements)
}
