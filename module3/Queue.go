package main

type Queue[T any] struct {
	Elements Linkedlist[T]
}

func (q *Queue[T]) Push(value T) {
	q.Elements.addToBack(value)
}
