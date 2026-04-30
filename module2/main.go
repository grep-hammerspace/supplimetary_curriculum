package main

import "fmt"

func main() {
	// Create a new linked list of integers
	list := &LinkedList[int]{}

	// Add elements to the end
	list.addToEnd(1)
	list.addToEnd(2)
	list.addToEnd(3)
	list.addToEnd(4)
	list.addToEnd(5)
	fmt.Println("After adding 1, 2, 3,4,5 to end:")
	printList(list) // Should print: 1 2 3 4. 5
	fmt.Printf("Size: %d, Head: %v, Tail: %v\n\n", list.size, list.Head.Value, list.Tail.Value)

	// Add an element to the start
	list.addToStart(0)
	fmt.Println("After adding 0 to start:")
	printList(list) // Should print: 0 1 2 3 4 5
	fmt.Printf("Size: %d, Head: %v, Tail: %v\n\n", list.size, list.Head.Value, list.Tail.Value)

	// Delete an element (e.g., index 2, which is value 2)
	deleted, err := list.delete(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Deleted value: %v\n", deleted)
		fmt.Println("After deleting index 2:")
		printList(list) // Should print: 0 1 3 4 5
		fmt.Printf("Size: %d, Head: %v, Tail: %v\n\n", list.size, list.Head.Value, list.Tail.Value)
	}

	// Reverse the list
	list.reverse()
	fmt.Println("After reversing:")
	printList(list) // Should print: 5 4 3 1 0
	fmt.Printf("Size: %d, Head: %v, Tail: %v\n\n", list.size, list.Head.Value, list.Tail.Value)

}
