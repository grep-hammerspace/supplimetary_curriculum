package hashtable

import (
	"bytes"
	"fmt"
	"hash/fnv"
)

type HashTable[K comparable, V any] struct {
	Buckets    [][]HashNode[K, V]
	LoadFactor float64
	capacity   int
}

type HashNode[K comparable, V any] struct {
	Key   K
	Value V
	Hash  uint32 // Avoid recomputing each time
}

func NewHashTable[K comparable, V any]() *HashTable[K, V] {
	return &HashTable[K, V]{
		Buckets:    make([][]HashNode[K, V], 16),
		LoadFactor: 1.0,
		capacity:   16,
	}
}

func (h *HashTable[K, V]) Put(key K, value V) error {
	// Determine bucket
	keyHash := hash(key)
	bucketNumber := keyHash % 16
	bucket := h.Buckets[bucketNumber]

	for _, node := range bucket {
		if node.Key == key {
			return fmt.Errorf("Cannot put duplicate key")
		}
	}
	newHashNode := HashNode[K, V]{key, value, keyHash}
	h.Buckets[bucketNumber] = append(bucket, newHashNode)
	return nil
}

func (h *HashTable[K, V]) Get(key K) (V, error) {
	// Get correct bucket
	hash := hash(key)
	bucket := h.Buckets[hash%16]

	// Go to bucket, traverse until you find K
	for _, node := range bucket {
		if node.Hash == hash && node.Key == key {
			return node.Value, nil
		}
	}
	return *new(V), fmt.Errorf("No such key found")
}

func (h *HashTable[K, V]) Remove(key K) error {
	hash := hash(key)
	bucketNumber := hash % 16
	bucket := h.Buckets[bucketNumber]

	for i, node := range bucket {
		if node.Hash == hash && node.Key == key {
			// Everything from first item up to i and everything after i
			h.Buckets[bucketNumber] = append(bucket[:i], bucket[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("No such key found")
}

// Naive way of printing key value pairs, probably wont work for more complex types
func (h *HashTable[K, V]) Stringify() string {
	var buffer bytes.Buffer

	for _, bucket := range h.Buckets {
		for _, node := range bucket {
			buffer.Write([]byte(fmt.Sprint(node.Key)))
			buffer.Write([]byte(","))
			buffer.Write([]byte(fmt.Sprint(node.Value)))
			buffer.Write([]byte("\n"))
		}
	}
	return buffer.String()
}

func hash[K any](obj K) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprint(obj)))
	return h.Sum32()
}
