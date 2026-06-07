package hashtable

import (
	"fmt"
	"testing"
)

// --- Put & Get ---

func TestPutAndGetBasic(t *testing.T) {
	ht := New[string, int]()
	ht.Put("apple", 5)
	ht.Put("banana", 3)
	ht.Put("orange", 8)

	cases := []struct {
		key      string
		expected int
	}{
		{"apple", 5},
		{"banana", 3},
		{"orange", 8},
	}

	for _, tc := range cases {
		val, err := ht.Get(tc.key)
		if err != nil || val != tc.expected {
			t.Errorf("Get(%q): expected (%d, nil), got (%d, %v)", tc.key, tc.expected, val, err)
		}
	}
}

func TestGetNonExistentKey(t *testing.T) {
	ht := New[string, int]()
	ht.Put("apple", 5)

	_, err := ht.Get("grape")
	if err == nil {
		t.Error("Expected error for non-existent key, got nil")
	}
}

func TestGetOnEmptyTable(t *testing.T) {
	ht := New[string, int]()
	_, err := ht.Get("anything")
	if err == nil {
		t.Error("Expected error on Get from empty table, got nil")
	}
}

// --- Upsert (Put updating existing key) ---

func TestPutUpdatesExistingKey(t *testing.T) {
	ht := New[string, int]()
	ht.Put("apple", 5)
	ht.Put("apple", 99)

	val, err := ht.Get("apple")
	if err != nil {
		t.Fatalf("Expected to find 'apple' after update, got error: %v", err)
	}
	if val != 99 {
		t.Errorf("Expected updated value 99, got %d", val)
	}
}

// --- Remove ---

func TestRemoveExistingKey(t *testing.T) {
	ht := New[string, int]()
	ht.Put("apple", 5)
	ht.Put("banana", 3)

	err := ht.Remove("apple")
	if err != nil {
		t.Fatalf("Expected no error removing existing key, got: %v", err)
	}

	_, err = ht.Get("apple")
	if err == nil {
		t.Error("Expected error after removing 'apple', but key still found")
	}

	// Other keys should be unaffected
	val, err := ht.Get("banana")
	if err != nil || val != 3 {
		t.Errorf("Expected banana=3 to still exist, got (%d, %v)", val, err)
	}
}

func TestRemoveOnEmptyTable(t *testing.T) {
	ht := New[string, int]()
	err := ht.Remove("anything")
	if err == nil {
		t.Error("Expected error removing from empty table, got nil")
	}
}

func TestRemoveAndReinsert(t *testing.T) {
	ht := New[string, int]()
	ht.Put("apple", 5)
	ht.Remove("apple")
	ht.Put("apple", 42)

	val, err := ht.Get("apple")
	if err != nil || val != 42 {
		t.Errorf("Expected (42, nil) after reinsert, got (%d, %v)", val, err)
	}
}

func TestRemoveAllKeys(t *testing.T) {
	ht := New[string, int]()
	keys := []string{"a", "b", "c", "d"}
	for i, k := range keys {
		ht.Put(k, i)
	}
	for _, k := range keys {
		if err := ht.Remove(k); err != nil {
			t.Errorf("Unexpected error removing key %q: %v", k, err)
		}
	}
	for _, k := range keys {
		if _, err := ht.Get(k); err == nil {
			t.Errorf("Expected error getting removed key %q, got nil", k)
		}
	}
}

// --- Collisions ---

func TestCollisionsGetAndPut(t *testing.T) {
	ht := New[string, int]()
	// Insert enough keys to statistically guarantee some collisions
	for i := 0; i < 32; i++ {
		key := fmt.Sprintf("key%d", i)
		ht.Put(key, i*10)
	}
	for i := 0; i < 32; i++ {
		key := fmt.Sprintf("key%d", i)
		val, err := ht.Get(key)
		if err != nil || val != i*10 {
			t.Errorf("Get(%q): expected (%d, nil), got (%d, %v)", key, i*10, val, err)
		}
	}
}

func TestCollisionRemove(t *testing.T) {
	ht := New[string, int]()
	for i := 0; i < 32; i++ {
		ht.Put(fmt.Sprintf("key%d", i), i)
	}
	// Remove every other key
	for i := 0; i < 32; i += 2 {
		ht.Remove(fmt.Sprintf("key%d", i))
	}
	// Removed keys should be gone
	for i := 0; i < 32; i += 2 {
		if _, err := ht.Get(fmt.Sprintf("key%d", i)); err == nil {
			t.Errorf("Expected key%d to be removed", i)
		}
	}
	// Remaining keys should still be present
	for i := 1; i < 32; i += 2 {
		val, err := ht.Get(fmt.Sprintf("key%d", i))
		if err != nil || val != i {
			t.Errorf("Expected key%d=%d to survive, got (%d, %v)", i, i, val, err)
		}
	}
}
