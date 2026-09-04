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

func (graph *Graph[T]) AddNode(value T) error {
	if len(graph.Nodes) == graph.Size {
		return fmt.Errorf("Max size of graph is %d", graph.Size)
	}
	graph.Nodes = append(graph.Nodes, value)
	graph.Edges[value] = make([]bool, graph.Size)
	return nil
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

func (graph *Graph[T]) GetNeighbors(value T) ([]T, error) {

	// make sure the node whose neighbours we want actually exist
	if graph.Edges[value] == nil {
		return nil, fmt.Errorf("Unable to get neighbors for invalid node %v", value)
	}
	edgesOnTargetNode := graph.Edges[value]
	var neighbours []T

	for i, nodeExists := range edgesOnTargetNode {
		if nodeExists {
			neighbours = append(neighbours, graph.Nodes[i])
		}
	}

	return neighbours, nil
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

	// Remove "edge " by setting value to false, set both across the diagonal
	graph.Edges[to][indexFrom] = false
	graph.Edges[from][indexTo] = false

	return nil
}

func (graph *Graph[T]) DepthFirstSearch(start T) []T {

	// The actual thing you want to do is to visit each node, and recursively call to its neighbours
	// For each node, we keep going until every one of its neighbours (for loop) has been visited

	//We use a map[T} for O(1) memebership checks, and string because its default value is nil, so membership is checked as visited[node] == nil or not
	visited := []T{}
	seen := map[T]struct{}{}
	dfsHelper(start, graph, &visited, seen)

	return visited
}

func dfsHelper[T cmp.Ordered](node T, graph *Graph[T], visited *[]T, seen map[T]struct{}) error {

	// Add current node to the list of seen nodes, so we dont call dfsHelper on it again
	seen[node] = struct{}{}

	neighbours, err := graph.GetNeighbors(node)

	if err != nil {
		return err
	}

	for _, node := range neighbours {
		if _, seenBefore := seen[node]; !seenBefore { // ie we havent been to this node before
			err := dfsHelper[T](node, graph, visited, seen)
			if err != nil {
				return err
			}
		}
	}

	// If we get here, when a node has no more neighbours that have never been seen, we mark it as visited
	*visited = append(*visited, node)
	fmt.Println(node)

	return nil
}

func (graph *Graph[T]) BreadthFirstSearch(start T) []T {

	visited := &[]T{}
	seen := map[T]struct{}{}

	*visited = append(*visited, start)
	seen[start] = struct{}{}

	// for each node, we get its neighbours, we add it to a queue, add its unseen neighbours to the queue (will make for a dfs?)
	// call sequentially? how do we get past the first set of neighbours,  just call helpre over and over until no neighbours?
	// in that case, how will ordering work?
	return nil
}

func bfsHelper() {}

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
