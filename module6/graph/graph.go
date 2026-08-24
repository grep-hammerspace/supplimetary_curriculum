package graph

// Graph represented internally as adjacency matrix.
// Adjacency matrix is a slices of slices, each must grow anytime we append a

type Graph[T any] struct {
	Capacity int
	Size     int
	Matrix   [32][32]*T
}

func New[T any](capacity int) *Graph[T] {
	return &Graph[T]{}
}

func (g *Graph[T]) Add(value T) {

}
