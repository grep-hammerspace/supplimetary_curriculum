package graph

import (
	"cmp"
	"fmt"
)

// Graph represented internally as adjacency matrix.
// Define the nodes as elements of a slice
// Define the adjacency matrix itself as Map<T comparable, []bool]> where we flip elements to true if there is connection between nodes

type Graph[T cmp.Ordered] struct {
	Size int
	Nodes []T
	Edges map[T][]bool
}

func New[T cmp.Ordered](size int) *Graph[T] {

	edges := map[T][]bool{}
	graph := Graph[T]{
		Size: size,
		Nodes: []T{},
		Edges: edges,
	}
	return &graph
}

func (graph *Graph[T]) AddNode(value T) {
	graph.Nodes = append(graph.Nodes, value)
	graph.Edges[value] = make([]bool, graph.Size)
}

func (graph *Graph[T]) AddEdge(from, to T) error err {

	var indexFrom, indexTo int
	for i, node := range graph.Nodes {
		if node == from {
			indexFrom = i
		}

		if node == to {
			indexTo = i
        }
	}

	if indexFrom == nil or indexTo == nil {
		return errors.New("unbable to create edge as a node is missing"
	}

	graph.Edges[to][indexFrom] = true

}

func (graph *Graph[T]) printAdjacencyMatrix() {}
