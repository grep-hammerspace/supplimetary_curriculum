package hashtable

import (
	"fmt"
	"hash/fnv"

	ll "github.com/grep-hammerspace/coding-curriculum/module2/linkedlist"
)

type HashTable[K string, V any] struct {
	Buckets    []ll.LinkedList[HashNode[K, V]]
	LoadFactor float64
	capacity   int
}

type HashNode[K string, V any] struct {
	Key   K
	Value V
	Hash  uint32 // Avoid recomputing each time
}

func NewHashTable[K string, V any]() *HashTable[K, V] {
	var buckets []ll.LinkedList[HashNode[K, V]]
	for i := 0; i < 16; i++ {
		buckets = append(buckets, ll.LinkedList[HashNode[K, V]]{})
	}
	return &HashTable[K, V]{
		Buckets:    buckets,
		LoadFactor: 1.0,
		capacity:   16,
	}
}

func (h *HashTable[K, V]) Put(key K, value V) error {
	// Determine bucket
	keyHash := hash(key)
	bucketNumber := keyHash % 16

	newHashNode := HashNode[K, V]{
		Key:   key,
		Value: value,
		Hash:  keyHash,
	}

	current := h.Buckets[bucketNumber].Head

	// Empty bucket
	if current == nil {
		h.Buckets[bucketNumber].AddToEnd(newHashNode)
		return nil
	}

	// exactly one element, check not duplicate
	if current.Value.Key == key {
		return fmt.Errorf("Cannot put duplicated key")
	}

	// Traverse the rest of the linked list and check for dupes
	for current.Next != nil {
		if current.Value.Key == key {
			return fmt.Errorf("Cannot put duplicated key")
		} else {
			current = current.Next
		}
	}

	h.Buckets[bucketNumber].AddToEnd(newHashNode)
	return nil
}

func (h *HashTable[K, V]) Get(key K) (V, error) {
	// Get correct bucket
	hash := hash(key)
	bucket := h.Buckets[hash%16]

	//Traverse bucket till we find it
	current := bucket.Head
	if current == nil {
		return *new(V), fmt.Errorf("No such key found")
	}

	// exactly one element, check not duplicate
	if current.Value.Hash != hash {
		return *new(V), fmt.Errorf("No such key found")
	}

	// Traverse the rest of the linked list and check for dupes
	for current.Next != nil {
		if current.Value.Key == key {
			return *new(V), fmt.Errorf("No such key found")
		} else {
			current = current.Next
		}
	}
	for current.Value.Hash != hash && current.Next != nil {

	}
	return *new(V), fmt.Errorf("No such key found")
}

//func (h *HashTable[K, V]) Remove(key K) error {
//	hash := hash(key)
//	bucketNumber := hash % 16
//	bucket := h.Buckets[bucketNumber]
//
//	for i, node := range bucket {
//		if node.Hash == hash && node.Key == key {
//			// Everything from first item up to i and everything after i
//			h.Buckets[bucketNumber] = append(bucket[:i], bucket[i+1:]...)
//			return nil
//		}
//	}
//
//	return fmt.Errorf("No such key found")
//}
//
//// Naive way of printing key value pairs, probably wont work for more complex types
//func (h *HashTable[K, V]) Stringify() string {
//	var buffer bytes.Buffer
//
//	for _, bucket := range h.Buckets {
//		for _, node := range bucket {
//			buffer.Write([]byte(fmt.Sprint(node.Key)))
//			buffer.Write([]byte(","))
//			buffer.Write([]byte(fmt.Sprint(node.Value)))
//			buffer.Write([]byte("\n"))
//		}
//	}
//	return buffer.String()
//}

func hash[K any](obj K) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprint(obj)))
	return h.Sum32()
}
