package hashtable

import (
	"testing"
)

func TestHashTableGet(t *testing.T) {
	ht := NewHashTable[string, int]()

	ht.Put("apple", 5)
	ht.Put("banana", 3)
	ht.Put("orange", 3)

	val, err := ht.Get("apple")
	if err != nil || val != 5 {
		t.Errorf("Expected (5, nil), got (%d, %v)", val, err)
	}

	val, err = ht.Get("banana")
	if err != nil || val != 3 {
		t.Errorf("Expected (3, nil), got (%d, %v)", val, err)
	}

	val, err = ht.Get("grape")
	if err == nil {
		t.Error("Expected error for non-existent key, got nil")
	}
}

//func TestHashTableRemove(t *testing.T) {
//	ht := NewHashTable[string, int]()
//
//	ht.Put("apple", 5)
//	ht.Put("banana", 3)
//
//	err := ht.Remove("apple")
//	if err != nil {
//		t.Errorf("Expected no error removing existing key, got %v", err)
//	}
//
//	_, err = ht.Get("apple")
//	if err == nil {
//		t.Error("Expected error after removing key, but key still exists")
//	}
//
//	err = ht.Remove("nonexistent")
//	if err == nil {
//		t.Error("Expected error removing non-existent key, got nil")
//	}
//}
//
//func TestHashTableWithDifferentTypes(t *testing.T) {
//	// Test with int keys and string values
//	ht := NewHashTable[int, string]()
//
//	ht.Put(1, "one")
//	ht.Put(2, "two")
//
//	val, err := ht.Get(1)
//	if err != nil || val != "one" {
//		t.Errorf("Expected (one, nil), got (%s, %v)", val, err)
//	}
//}
//
//func TestHashTableCollisions(t *testing.T) {
//	// Test behavior when multiple keys hash to same bucket
//	ht := NewHashTable[string, int]()
//
//	ht.Put("key1", 10)
//	ht.Put("key2", 20)
//	ht.Put("key3", 30)
//
//	val1, _ := ht.Get("key1")
//	val2, _ := ht.Get("key2")
//	val3, _ := ht.Get("key3")
//
//	if val1 != 10 || val2 != 20 || val3 != 30 {
//		t.Error("Values not retrieved correctly after collisions")
//	}
//}
