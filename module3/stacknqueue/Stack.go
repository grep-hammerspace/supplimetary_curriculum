package stacknqueue

import ll "github.com/grep-hammerspace/coding-curriculum/module2/linkedlist"

type Stack[T any] struct {
	Elements ll.LinkedList[T]
}

func (q *Stack[T]) Push(value T) {
	q.Elements.AddToEnd(value)
}

// Returns (firstElement, true) if element exists, (dummyValue,false) otherwise
func (q *Stack[T]) Pop() (T, bool) {
	return q.Elements.GetFirst()
}

// Returns (copy of first element, true) if element exists, (dummyValue,false) otherwise
func (q *Stack[T]) Peek() (T, bool) {
	return q.Elements.Peek()
}
