package main

import (
	g "github.com/grep-hammerspace/coding-curriculum/module6/graph"
)

func main() {
	myGraph := g.New[int](3)

	myGraph.AddNode(5)
	myGraph.AddNode(1)
	myGraph.AddNode(2)

	myGraph.AddEdge(5, 1)
	myGraph.AddEdge(5, 2)

	myGraph.DepthFirstSearch(5)
}
