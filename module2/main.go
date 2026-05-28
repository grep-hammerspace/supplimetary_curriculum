package main

import "fmt"

// safeValue returns a string representation of the node's value or "nil" if node is nil.
func safeValue[T any](node *Node[T]) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", node.Value)
}

func main() {
	// Create a new linked list of integers
	list := &LinkedList[int]{}

	list.addToStart(42)
	fmt.Println("add single element to start")
	printList(list)
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.size, safeValue(list.Head), safeValue(list.Tail))

	// remove to empty the list
	deleted, err := list.delete(0)
	if err != nil {
		fmt.Println("Error deleting:", err)
	} else {
		fmt.Printf("Deleted value: %v\n", deleted)
	}
	fmt.Println("delete single element to empty the list")
	printList(list)
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.size, safeValue(list.Head), safeValue(list.Tail))

	// Add elements to the end
	list.addToEnd(1)
	fmt.Println("After adding 1 to end:")
	printList(list)
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.size, safeValue(list.Head), safeValue(list.Tail))

	list.addToEnd(2)
	list.addToEnd(3)
	list.addToEnd(4)
	list.addToEnd(5)
	fmt.Println("After adding 2, 3,4,5 to end:")
	printList(list) // Should print: 1 2 3 4 5
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.size, safeValue(list.Head), safeValue(list.Tail))

	// Add an element to the start
	list.addToStart(0)
	fmt.Println("After adding 0 to start:")
	printList(list) // Should print: 0 1 2 3 4 5
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.size, safeValue(list.Head), safeValue(list.Tail))

	list.delete(0) // delete head of a multi item list
	fmt.Println("Deleted 0 from head")
	printList(list) // Should print: 1 2 3 4 5
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.size, safeValue(list.Head), safeValue(list.Tail))

	// Delete an element (e.g., index 2, which is value 3)
	deleted, err = list.delete(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Deleted value: %v\n", deleted)
		fmt.Println("After deleting index 2:")
		printList(list) // Should print: 1 2 4 5
		fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.size, safeValue(list.Head), safeValue(list.Tail))
	}

	// Reverse the list
	list.reverse()
	fmt.Println("After reversing:")
	printList(list) // Should print: 5 4 2 1
	fmt.Printf("Size: %d, Head: %s, Tail: %s\n\n", list.size, safeValue(list.Head), safeValue(list.Tail))
}
