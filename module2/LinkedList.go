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
	l.Tail  = node
	l.size++
}

func (l *LinkedList[T]) delete(target int) T {
	node := l.Head
	i := 0
	for (node != nil) {
		if
	}
	return *new(T)
}

// Reverse order in place, without creating a new list by manipulating next pointers
func (l *LinkedList[T]) reverse() {}

func printList[T any](l *LinkedList[T]) {
	current := l.Head
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
	fmt.Println()
}
