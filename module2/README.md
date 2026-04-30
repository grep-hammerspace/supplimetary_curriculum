## Module 2 - Linked Lists

### Tasks
Implement linked list with the following methods:
- `addToStart(value)` - adds a new node with the given value to the end of the list
- `addToEnd(value)` - adds a new node with the given value to the beginning of the list
- `delete(position)` - deletes node at given index (0-based)
- `reverse()` - reverses the list in place with no copies

### Results
```
After adding 1, 2, 3,4,5 to end:
1 2 3 4 5 
Size: 5, Head: 1, Tail: 5

After adding 0 to start:
0 1 2 3 4 5 
Size: 6, Head: 0, Tail: 5

Deleted value: 2
After deleting index 2:
0 1 3 4 5 
Size: 5, Head: 0, Tail: 5

After reversing:
5 4 3 1 0 
Size: 5, Head: 5, Tail: 0
```