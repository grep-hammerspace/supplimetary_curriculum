package graph

import (
	"cmp"
	"fmt"
)

// Graph represented internally as adjacency matrix.
// Define the nodes as elements of a slice
// Define the adjacency matrix itself as Map<T comparable, []bool]> where we flip elements to true if there is connection between nodes

type Graph[T cmp.Ordered] struct {
	Size  int
	Nodes []T
	Edges map[T][]bool
}

func New[T cmp.Ordered](size int) *Graph[T] {

	edges := map[T][]bool{}
	graph := Graph[T]{
		Size:  size,
		Nodes: []T{},
		Edges: edges,
	}
	return &graph
}

func (graph *Graph[T]) AddNode(value T) {
	graph.Nodes = append(graph.Nodes, value)
	graph.Edges[value] = make([]bool, graph.Size)
}

func (graph *Graph[T]) AddEdge(from, to T) error {

	//Set sentinel value for invalid index
	indexFrom := -1
	indexTo := -1

	for i, node := range graph.Nodes {
		if node == from {
			indexFrom = i
		}

		if node == to {
			indexTo = i
		}
	}

	if indexFrom == -1 || indexTo == -1 {
		return fmt.Errorf("unable to create edge, invalid nodes")
	}

	if indexFrom == indexTo {
		return fmt.Errorf("Unable to create edge, from and to are the same")
	}

	// Create "edge" by setting to true, rememer the matrix is meant to be symmetric across the diagonal
	graph.Edges[from][indexTo] = true
	graph.Edges[to][indexFrom] = true

	return nil
}

func (graph *Graph[T]) RemoveEdge(from, to T) error {

	indexFrom := -1
	indexTo := -1

	for i, node := range graph.Nodes {
		if node == from {
			indexFrom = i
		}

		if node == to {
			indexTo = i
		}
	}

	if indexFrom == -1 || indexTo == -1 {
		return fmt.Errorf("unable to remove edge, invalid nodes")
	}

	// Remove "edge " by setting value to false, set both across the diagnoal
	graph.Edges[to][indexFrom] = false
	graph.Edges[from][indexTo] = false

	return nil
}

// PrintAdjacencyMatrix writes the matrix as a labelled grid, one row per node.
// Rows are the "from" node and columns the "to" node, matching AddEdge(from, to).
func (graph *Graph[T]) PrintAdjacencyMatrix() {

	if len(graph.Nodes) == 0 {
		fmt.Println("<empty graph>")
		return
	}

	labels := make([]string, len(graph.Nodes))
	width := 1
	for i, node := range graph.Nodes {
		labels[i] = fmt.Sprint(node)
		width = max(width, len(labels[i]))
	}

	// Header row: a blank corner cell followed by every column label.
	fmt.Printf("%-*s", width, "")
	for _, label := range labels {
		fmt.Printf(" %*s", width, label)
	}
	fmt.Println()

	for from := range graph.Nodes {
		fmt.Printf("%-*s", width, labels[from])
		for _, to := range graph.Nodes {
			// Edges[to] is indexed by the position of the "from" node.
			cell := 0
			if row, ok := graph.Edges[to]; ok && from < len(row) && row[from] {
				cell = 1
			}
			fmt.Printf(" %*d", width, cell)
		}
		fmt.Println()
	}
}
