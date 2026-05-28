## Module 2 - Linked Lists

### Tasks
Implement linked list with the following methods:
- `addToStart(value)` - adds a new node with the given value to the end of the list
- `addToEnd(value)` - adds a new node with the given value to the beginning of the list
- `delete(position)` - deletes node at given index (0-based)
- `reverse()` - reverses the list in place with no copies

### Results
```
asad@obsidian:~/Projects/Personal/supplimetary_curriculum$ go run ./module2/
add single element to start
42 
Size: 1, Head: 42, Tail: 42

Deleted value: 42
delete single element to empty the list

Size: 0, Head: nil, Tail: nil

After adding 1 to end:
1 
Size: 1, Head: 1, Tail: 1

After adding 1, 2, 3,4,5 to end:
1 2 3 4 5 
Size: 5, Head: 1, Tail: 5

After adding 0 to start:
0 1 2 3 4 5 
Size: 6, Head: 0, Tail: 5

Deleted 0 from head
1 2 3 4 5 
Size: 5, Head: 1, Tail: 5

Deleted value: 3
After deleting index 2:
1 2 4 5 
Size: 4, Head: 1, Tail: 5

After reversing:
5 4 2 1 
Size: 4, Head: 5, Tail: 1
```