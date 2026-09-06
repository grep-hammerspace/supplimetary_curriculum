package main

import (
	"fmt"

	g "github.com/grep-hammerspace/coding-curriculum/module6/graph"
)

func main() {
	myGraph := g.New[int](9)

	myGraph.AddNode(1)
	myGraph.AddNode(2)
	myGraph.AddNode(3)
	myGraph.AddNode(4)
	myGraph.AddNode(5)
	myGraph.AddNode(6)
	myGraph.AddNode(7)
	myGraph.AddNode(8)
	myGraph.AddNode(9)

	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 3)

	myGraph.AddEdge(2, 4)
	myGraph.AddEdge(2, 5)
	myGraph.AddEdge(2, 6)

	myGraph.AddEdge(3, 7)
	myGraph.AddEdge(3, 8)
	myGraph.AddEdge(3, 9)

	dfsResults := myGraph.DepthFirstSearch(1)

	fmt.Println(dfsResults)
}
