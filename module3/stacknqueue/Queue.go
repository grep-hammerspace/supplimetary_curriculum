package stacknqueue

import ll "github.com/grep-hammerspace/coding-curriculum/module2/linkedlist"

type Queue[T any] struct {
	Elements ll.LinkedList[T]
}

func (q *Queue[T]) Enqueue(value T) {
	q.Elements.AddToEnd(value)
}

func (q *Queue[T]) Deque() (T, bool) {
	return q.Elements.GetFirst()
}

func (q *Queue[T]) Peek() (T, bool) {
	return q.Elements.PeekFirst()
}

func (q *Queue[T]) Size() int {
	return q.Elements.Size()
}
