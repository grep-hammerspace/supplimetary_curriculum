package hashtable

import (
	"bytes"
	"fmt"
	"hash/fnv"

	ll "github.com/grep-hammerspace/coding-curriculum/module2/linkedlist"
)

type HashTable[K string, V any] struct {
	Buckets    []ll.LinkedList[HashNode[string, V]]
	LoadFactor float64
	Capacity   uint32 // To allow bucket computation without type conversion
}

type HashNode[K string, V any] struct {
	Key   string
	Value V
	Hash  uint32 // Avoid recomputing each time
}

func New[K string, V any]() *HashTable[K, V] {
	var buckets []ll.LinkedList[HashNode[string, V]]
	for i := 0; i < 16; i++ {
		buckets = append(buckets, ll.LinkedList[HashNode[string, V]]{})
	}
	return &HashTable[K, V]{
		Buckets:    buckets,
		LoadFactor: 1.0,
		Capacity:   16,
	}
}

func (h *HashTable[K, V]) Put(key string, value V) {
	// Determine bucket
	keyHash := hash(key)
	bucketNumber := keyHash % h.Capacity

	current := h.Buckets[bucketNumber].Head

	newHashNode := HashNode[string, V]{
		Key:   key,
		Value: value,
		Hash:  keyHash,
	}

	// Traverse the rest of the linked list and check for dup keys in case of an update
	for current != nil {
		if current.Value.Key == key {
			current.Value = newHashNode
			return
		} else {
			current = current.Next
		}
	}

	//New values are added to end
	h.Buckets[bucketNumber].AddToEnd(newHashNode)
}

func (h *HashTable[K, V]) Get(key string) (V, error) {
	// Get correct bucket
	keyHash := hash(key)
	bucket := h.Buckets[keyHash%h.Capacity]

	//Traverse bucket till we find it
	current := bucket.Head

	for current != nil {
		if current.Value.Key == key {
			return current.Value.Value, nil
		} else {
			current = current.Next
		}
	}

	return *new(V), fmt.Errorf("No such key found")
}

func (h *HashTable[K, V]) Remove(key string) error {
	keyHash := hash(key)
	bucketNumber := keyHash % h.Capacity

	current := h.Buckets[bucketNumber].Head

	found := false
	deleteIndex := 0
	for current != nil {
		if current.Value.Hash == keyHash && current.Value.Key == key {
			found = true
			break
		} else {
			deleteIndex += 1
			current = current.Next
		}
	}

	if found {
		h.Buckets[bucketNumber].Delete(deleteIndex)
		return nil
	} else {
		return fmt.Errorf("No such key found")
	}
}

// Naive way of printing key value pairs, probably wont work for more complex types
func (h *HashTable[K, V]) Stringify() string {
	var buffer bytes.Buffer

	for _, bucket := range h.Buckets {
		current := bucket.Head
		for current != nil {
			buffer.Write([]byte(current.Value.Key))
			buffer.Write([]byte(","))
			buffer.Write([]byte(fmt.Sprint(current.Value)))
			buffer.Write([]byte("\n"))
			current = current.Next
		}
	}
	return buffer.String()
}

func hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}
