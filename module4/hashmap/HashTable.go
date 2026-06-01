package hashmap

import (
	"fmt"
	"hash/fnv"
)

type HashTable[K any, V any] struct {
	Buckets    [][]HashNode[K, V]
	LoadFactor float64
	capacity   int
}

type HashNode[K any, V any] struct {
	Key   K
	Value V
	Hash  uint32 // Avoid recomputing each time
}

func (h *HashTable[K, V]) Put(key K, value V) {
	// Determine bucket
	keyHash := hash(key)
	bucketNumber := keyHash % 16

	//write to bucket
	bucket := h.Buckets[bucketNumber]
	newHashNode := HashNode[K, V]{key, value, keyHash}
	bucket = append(bucket, newHashNode)
}

func (h *HashTable[K, V]) Get(key K) V {
	// Get correct bucket
	bucket := h.Buckets[hash(key)%16]

	// Go to bucket, traverse until you find K
	for _, node := range bucket {
		if node.Key == key {
			return node.Value
		}
	}

}

func hash[K any](obj K) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprint(obj)))
	return h.Sum32()
}
