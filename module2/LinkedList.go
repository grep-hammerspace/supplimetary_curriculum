package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	size int
}

func (l *LinkedList[T]) addToStart(value T) {
	node := &Node[T]{value, nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	}

	node.Next = l.Head
	l.Head = node

	l.size++
}

func (l *LinkedList[T]) addToEnd(value T) {
	node := &Node[T]{value, nil}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	}

	l.Tail.Next = node
	l.Tail = node
	l.size++
}

func (l *LinkedList[T]) delete(target int) (T, error) {
	// Handle empty list or invalid index
	if l.Head == nil || target < 0 || target >= l.size {
		var zero T
		return zero, fmt.Errorf("invalid index or empty list")
	}

	// we are deleting head
	if target == 0 {
		valueToReturn := l.Head.Value
		l.Head = l.Head.Next
		return valueToReturn, nil
	}

	// Find node we are deleting
	currentNode := l.Head
	for i := 0; i < target-1; i++ {
		currentNode = currentNode.Next
	}

	nodeToDelete := currentNode.Next
	valueToReturn := nodeToDelete.Value

	// If we are deleting the tail, we need to update the tail pointer
	if nodeToDelete.Next == nil {
		currentNode.Next = nil
		l.Tail = currentNode
	}

	// otherwise just make the pointer skip over what we are deleting
	currentNode.Next = nodeToDelete.Next

	l.size--
	return valueToReturn, nil
}

// Reverse order in place, without creating a new list, just by manipulating next pointers
func (l *LinkedList[T]) reverse() {
	if l.Head == nil || l.Head.Next == nil {
		return // Nothing to reverse (empty or single-node list)
	}

	var previous *Node[T] = nil // Will track the reversed chain
	current := l.Head           // Start at the head
	l.Tail = l.Head             // The original head will become the new tail

	for current != nil {
		next := current.Next    // Save the next node before we change it
		current.Next = previous // Reverse: point current back to previous
		previous = current      // Move previous forward (to current)
		current = next          // Move current forward (to next)
	}

	// After the loop, previous is the new head
	l.Head = previous
}

func printList[T any](l *LinkedList[T]) {
	current := l.Head
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
	fmt.Println()
}
