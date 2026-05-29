package main

import (
	"fmt"

	snq "github.com/grep-hammerspace/coding-curriculum/module3/stacknqueue"
)

func main() {
	// Unicode pointers for the brackets as int32 because runes are aliases for int32
	// This is ok becuase we only care about ASCII characters here
	var curly_open int32 = 123
	var square_open int32 = 91
	var round_open int32 = 40
	var round_close int32 = 41
	var square_close int32 = 93
	var curly_close int32 = 125
	openCloseMappings := make(map[int32]int32)
	openCloseMappings[curly_open] = curly_close
	openCloseMappings[curly_close] = curly_open
	openCloseMappings[square_open] = square_close
	openCloseMappings[square_close] = square_open
	openCloseMappings[round_open] = round_close
	openCloseMappings[round_close] = round_open

	fmt.Printf("{[()]} is balanced: %t\n", areBracketsBalanced("{[()]}", openCloseMappings))
	fmt.Printf("{[]()} is balanced: %t\n", areBracketsBalanced("{[]()}", openCloseMappings))
	fmt.Printf("{[(])} is balanced: %t\n", areBracketsBalanced("{[(])}", openCloseMappings))
	fmt.Printf("{[(]]] is balanced: %t\n", areBracketsBalanced("{[(]]]", openCloseMappings))
}

func areBracketsBalanced(brackets string, mappings map[int32]int32) bool {

	// Create a slice of the inputs and a stack to match them
	bracketsAsSlice := []rune(brackets)
	runeStack := snq.Stack[rune]{}

	// Add first element manually
	runeStack.Push(bracketsAsSlice[0])

	// Iterate over the rest, if rune at current index matches the top of the stack, pop, otherwise push
	for i := 1; i < len(bracketsAsSlice); i++ {
		currentRune := bracketsAsSlice[i]
		topOfStack, _ := runeStack.Peek()

		if mappings[currentRune] == topOfStack {
			runeStack.Pop()
		} else {
			runeStack.Push(currentRune)
		}
	}

	return runeStack.Size() == 0

}
