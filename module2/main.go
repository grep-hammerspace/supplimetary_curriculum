package main

func main() {

}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}

func (l *LinkedList[T]) addToStart(value T) {}

func (l *LinkedList[T]) addToEnd(value T) {}

func (l *LinkedList[T]) delete(target int) T {
	return *new(T)
}

// Reverse order in place, without creating a new list by manipulating next pointers
func (l *LinkedList[T]) reverse() {}
