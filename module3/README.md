## Stacks and queues

- Implement a stack  with push, pop, peek, and a Queue with enque , deque, peek
- write a function that uses a Stack to check for parentheses being closed, ie [({})] is ok  but {[)(]}. the nesting should be correct as well, so [({})] is ok but [({)}] is not ok.


## Solutions
- Implementation of stack and queue can be found in `Stack.go` and `Queue.go` respectively under the `stacknqueue` package.
- The function to check for parentheses is implemented in `parentheses.go`.

#### Sample output for parentheses check:
```
asad@obsidian:~/Projects/Personal/supplimetary_curriculum$ go run ./module3/
{[()]} is balanced: true
{[]()} is balanced: true
{[(])} is balanced: false
{[(]]] is balanced: false
```

