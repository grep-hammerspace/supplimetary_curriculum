package hashmap

import (
	"hash/fnv"

	ll "github.com/grep-hammerspace/coding-curriculum/module2/linkedlist"
)

type HashTable[K any, V any] struct {
	Buckets    []ll.LinkedList[V]
	LoadFactor float64
	capacity   int
}

type HashNode[K any, V any] struct {
	Key   K
	Value V
	Hash  uint32 // Avoid recomputing each time
	Next  *HashNode[K, V]
}

func (h *HashTable[K, V]) Put(key K, value V) {
	bucket := hash(key) % 16
	//Get corresponding bucket
	//write to bucket
}

func (h *HashTable[K, V]) Get(K) V {
	// Hash input
	// Go to bucket, traverse until you find K

}

func hash[K ~string](obj K) uint32 {
	h := fnv.New32a()
	h.Write([]byte(obj))
	return h.Sum32()
}
