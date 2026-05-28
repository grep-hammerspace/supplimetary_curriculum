package linkedlist

import (
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	size int
}

func (l *LinkedList[T]) AddToStart(value T) {
	node := &Node[T]{value, nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.size++
}

func (l *LinkedList[T]) AddToEnd(value T) {
	node := &Node[T]{value, nil}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}

	l.size++
}

func (l *LinkedList[T]) Delete(target int) (T, error) {
	// Handle empty list or invalid index
	if l.Head == nil || target < 0 || target >= l.size {
		var zero T
		return zero, fmt.Errorf("invalid index or empty list")
	}

	// we are deleting head
	if target == 0 {

		if l.Head == l.Tail {
			// ie we are deleting the only node
			valueToReturn := l.Head.Value
			l.Head = nil
			l.Tail = nil
			l.size--
			return valueToReturn, nil
		}

		// we can just dereference head and let the GC clean it up
		valueToReturn := l.Head.Value
		l.Head = l.Head.Next
		l.size--
		return valueToReturn, nil
	}

	//  if not deleting head, find node we are deleting
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
	} else {
		// otherwise just make the pointer skip over what we are deleting
		currentNode.Next = nodeToDelete.Next
	}

	l.size--
	return valueToReturn, nil
}

// Reverse order in place, without creating a new list, just by manipulating next pointers
func (l *LinkedList[T]) Reverse() {
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

func (l *LinkedList[T]) Size() int {
	return l.size
}

// Returns (firstElement, true) if element exists, (dummyValue,false) otherwise
func (l *LinkedList[T]) GetFirst() (T, bool) {

	// Empty LinkedList
	if l.Head == nil {
		var dummy T
		return dummy, false
	}

	// One element only
	if l.Head == l.Tail {
		valueToReturn := l.Head.Value
		l.Head = nil
		l.Tail = nil
		l.size--
		return valueToReturn, true
	}

	// Return head value and update head pointer to next node
	valueToReturn := l.Head.Value
	l.Head = l.Head.Next
	return valueToReturn, true
}

// Returns (firstElement, true) if element exists, (dummyValue,false) otherwise
func (l *LinkedList[T]) GetLast() (T, bool) {

	// Empty LinkedList
	if l.Head == nil {
		var dummy T
		return dummy, false
	}

	// One element only
	if l.Head == l.Tail {
		valueToReturn := l.Head.Value
		l.Head = nil
		l.Tail = nil
		l.size--
		return valueToReturn, true
	}

	// Iterate up to second last
	currentNode := l.Head
	for currentNode.Next.Next != nil {
		currentNode = currentNode.Next
	}

	valueToReturn := currentNode.Next.Value
	l.Tail = currentNode
	return valueToReturn, true
}

// Return value and true if found, false if not
func (l *LinkedList[T]) Peek() (T, bool) {

	var dummy T
	if l.Head.Value == nil {
		return dummy, false
	}
	return l.Head.Value, true

}

func PrintList[T any](l *LinkedList[T]) {
	current := l.Head
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
	fmt.Println()
}
