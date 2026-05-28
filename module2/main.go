package main

import (
	"fmt"

	linkedlist2 "github.com/grep-hammerspace/coding-curriculum/module2/linkedlist"
)

// safeValue returns a string representation of the node's value or "nil" if node is nil.
func safeValue[T any](node *linkedlist2.Node[T]) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", node.Value)
}

func main() {
	// Create a new linked list of integers
	list := &linkedlist2.LinkedList[int]{}

	list.AddToStart(42)
	fmt.Println("add single element to start")
	linkedlist2.PrintList(list)
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.Size(), safeValue(list.Head), safeValue(list.Tail))

	// remove to empty the list
	deleted, err := list.Delete(0)
	if err != nil {
		fmt.Println("Error deleting:", err)
	} else {
		fmt.Printf("Deleted value: %v\n", deleted)
	}
	fmt.Println("delete single element to empty the list")
	linkedlist2.PrintList(list)
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.Size(), safeValue(list.Head), safeValue(list.Tail))

	// Add elements to the end
	list.AddToEnd(1)
	fmt.Println("After adding 1 to end:")
	linkedlist2.PrintList(list)
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.Size(), safeValue(list.Head), safeValue(list.Tail))

	list.AddToEnd(2)
	list.AddToEnd(3)
	list.AddToEnd(4)
	list.AddToEnd(5)
	fmt.Println("After adding 2, 3,4,5 to end:")
	linkedlist2.PrintList(list) // Should print: 1 2 3 4 5
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.Size(), safeValue(list.Head), safeValue(list.Tail))

	// Add an element to the start
	list.AddToStart(0)
	fmt.Println("After adding 0 to start:")
	linkedlist2.PrintList(list) // Should print: 0 1 2 3 4 5
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.Size(), safeValue(list.Head), safeValue(list.Tail))

	list.Delete(0) // delete head of a multi item list
	fmt.Println("Deleted 0 from head")
	linkedlist2.PrintList(list) // Should print: 1 2 3 4 5
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.Size(), safeValue(list.Head), safeValue(list.Tail))

	// Delete an element (e.g., index 2, which is value 3)
	deleted, err = list.Delete(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Deleted value: %v\n", deleted)
		fmt.Println("After deleting index 2:")
		linkedlist2.PrintList(list) // Should print: 1 2 4 5
		fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.Size(), safeValue(list.Head), safeValue(list.Tail))
	}

	// Reverse the list
	list.Reverse()
	fmt.Println("After reversing:")
	linkedlist2.PrintList(list) // Should print: 5 4 2 1
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.Size(), safeValue(list.Head), safeValue(list.Tail))
}
