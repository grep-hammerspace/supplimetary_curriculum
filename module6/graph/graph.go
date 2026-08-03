package graph

type Node[T any] struct {
	Neighbours []*Node[T]
	Value      T
}

// Do graphs have root nodes?
// How could you print the shape of a graph to terminal.
