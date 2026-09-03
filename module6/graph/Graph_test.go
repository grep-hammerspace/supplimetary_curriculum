package graph

import (
	"fmt"
	"slices"
	"testing"
)

func TestAddEdgeBetweenTwoNodesIsSymmetric(t *testing.T) {

	graph := New[string](2)

	graph.AddNode("A")
	graph.AddNode("B")

	graph.AddEdge("A", "B")

	fmt.Println("--------- TestInsertIntoEmptyBinaryTree -----------------")
	graph.PrintAdjacencyMatrix()

	if graph.Edges["A"][1] != true {
		t.Errorf("Expected edge between A and B, not found")
	}

	if graph.Edges["B"][0] != true {
		t.Errorf("Expected edge between A and B, not found")
	}

}

func TestAddNodesIntoEmptyBinaryTree(t *testing.T) {

	graph := New[string](2)

	graph.AddNode("A")
	graph.AddNode("B")

	fmt.Println("--------- TestAddNodesIntoEmptyBinaryTree -----------------")
	graph.PrintAdjacencyMatrix()

	if len(graph.Nodes) != 2 {
		t.Errorf("Invalid number of rows in adjacency matrix, expected 2, got %d ", len(graph.Nodes))
	}

	if graph.Nodes[0] != "A" && graph.Nodes[1] != "B" {
		t.Errorf("Expected first and second edge between A and B, not found")
	}

	for _, nodeEdgeList := range graph.Edges {
		if len(nodeEdgeList) != 2 {
			t.Errorf("Invalid number of columns in matrix, expected 2, got %d ", len(nodeEdgeList))
		}
		for _, nodeEdge := range nodeEdgeList {
			if nodeEdge == true {
				t.Errorf("Expected 0 edges, found at least one")
			}
		}

	}
}

func TestGetNeighboursGetsNeighboursForGivenNode(t *testing.T) {
	graph := New[string](5)

	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("A", "D")
	graph.AddEdge("A", "E")

	graph.PrintAdjacencyMatrix()

	neighbours, err := graph.GetNeighbors("A")
	fmt.Println(neighbours)

	if err != nil {
		t.Errorf("Error while getting neighbours")
	}

	if len(neighbours) != 4 {
		t.Errorf("Expected 4 neighbours for A, got %d", len(neighbours))
	}

}

func TestRemoveEdgeFromGraph(t *testing.T) {

	graph := New[string](2)

	graph.AddNode("A")
	graph.AddNode("B")

	fmt.Println("--------- Add A-B -----------------")
	graph.AddEdge("A", "B")
	graph.PrintAdjacencyMatrix()

	if graph.Edges["A"][1] != true {
		t.Errorf("Expected edge between A and B, not found")
	}

	if graph.Edges["B"][0] != true {
		t.Errorf("Expected edge between A and B, not found")
	}

	fmt.Println("--------- Remove A-B -----------------")
	err := graph.RemoveEdge("A", "B")
	if err != nil {
		fmt.Errorf("Error while removing edge")
	}
	graph.PrintAdjacencyMatrix()

	for _, nodeEdgeList := range graph.Edges {

		for _, nodeEdge := range nodeEdgeList {
			if nodeEdge == true {
				t.Errorf("Expected 0 edges, found at least one")
			}
		}

	}

}

func TestDepthFirstSearch(t *testing.T) {
	graph := New[string](6)

	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")
	graph.AddNode("F")

	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("B", "D")
	//graph.AddEdge("D", "E")
	//graph.AddEdge("D", "F")

	graph.PrintAdjacencyMatrix()

	// We expect BFS to give us C,D,B, A - Expect to traverse from A down to B, then inverse input order visiting for any node who has more than one child
	// In this example we add node C, then node D, and, getNeighbours wil return ["A","C","D"], we will have seen A already so we first call dfsHelper for C,
	// whose only Neighbour is B, which we will have seen, so we add C to visited, then Dm then back up the stack for B, nad A
	visited := graph.DepthFirstSearch("A")

	if len(visited) != 3 {
		t.Errorf("Expected 3 nodes to be visited in DFS, got %d", len(visited))
	}

	if !slices.Equal(visited, []string{"C", "D", "B", "A"}) {
		t.Errorf("Incorrect order for DFS, expected 'C','D','B','A'")
	}
}
