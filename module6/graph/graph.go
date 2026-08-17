package graph

type Node[T any] struct {
	Neighbours []*Node[T]
	Value      T
}

// Link a node to an existing one, returns pointer to new node
func (n *Node[T]) Link(value T) *Node[T] {
	newNode := Node[T]{
		Neighbours: make([]*Node[T], 0),
		Value:      value,
	}

	n.Neighbours = append(n.Neighbours, &newNode)

	return &newNode
}

type Graph[T any] struct {
	Root *Node[T]
}



// Create an empty graph with a single root node.
func New[T any](value T) *Graph[T] {
	root := Node[T]{
		Neighbours: make([]*Node[T], 0),
		Value: value,
	}

	graph := Graph[T]{
		Root: &root,
	}

	return &graph
}

// Do graphs have root nodes?
// yes they do, otherwise how would you do a depth first search or a breadth first search, you gotta start somewhere
// How could you print the shape of a graph to terminal.
// This isnt as easy, drawing lines would be really tricky, so maybe some kinda numerical represnetation. - surprisingly enough, and adjacecny matrix is what they suggested, whod have thunk it?
