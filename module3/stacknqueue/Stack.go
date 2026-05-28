package stacknqueue

import ll "github.com/grep-hammerspace/coding-curriculum/module2/linkedlist"

type Stack[T any] struct {
	Elements ll.LinkedList[T]
}

func (s *Stack[T]) Push(value T) {
	s.Elements.AddToEnd(value)
}

// Returns (firstElement, true) if element exists, (dummyValue,false) otherwise
func (s *Stack[T]) Pop() (T, bool) {
	return s.Elements.GetLast()
}

// Returns (copy of first element, true) if element exists, (dummyValue,false) otherwise
func (s *Stack[T]) Peek() (T, bool) {
	return s.Elements.PeekLast()
}
func (s *Stack[T]) Size() int {
	return s.Elements.Size()
}
