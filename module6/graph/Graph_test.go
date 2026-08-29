package graph

import (
	"fmt"
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
