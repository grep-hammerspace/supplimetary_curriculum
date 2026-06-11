package main

import (
	"fmt"

	snq "github.com/grep-hammerspace/coding-curriculum/module3/stacknqueue"
)

func main() {
	// Unicode pointers for the brackets as int32 because runes are aliases for int32
	// This is ok becuase we only care about ASCII characters here
	var curly_open_rune int32 = 123
	var square_open_rune int32 = 91
	var round_open_rune int32 = 40
	var round_close_rune int32 = 41
	var square_close_rune int32 = 93
	var curly_close_rune int32 = 125

	curly_open := bracket{shape: "{", isOpen: true}
	square_open := bracket{shape: "[", isOpen: true}
	round_open := bracket{shape: "(", isOpen: true}
	round_close := bracket{shape: ")", isOpen: false}
	square_close := bracket{shape: "]", isOpen: false}
	curly_close := bracket{shape: "}", isOpen: false}

	runeMappings := make(map[rune]bracket)
	runeMappings[curly_open_rune] = curly_open
	runeMappings[curly_close_rune] = curly_close
	runeMappings[square_open_rune] = square_open
	runeMappings[square_close_rune] = square_close
	runeMappings[round_open_rune] = round_open
	runeMappings[round_close_rune] = round_close

	openCloseMappings := make(map[bracket]bracket)
	openCloseMappings[curly_open] = curly_close
	openCloseMappings[curly_close] = curly_open
	openCloseMappings[square_open] = square_close
	openCloseMappings[square_close] = square_open
	openCloseMappings[round_open] = round_close
	openCloseMappings[round_close] = round_open

	fmt.Printf("{[()]} is balanced: %t\n", areBracketsBalanced("{[()]}", runeMappings, openCloseMappings))
	fmt.Printf("{[]()} is balanced: %t\n", areBracketsBalanced("{[]()}", runeMappings, openCloseMappings))
	fmt.Printf("{[(])} is balanced: %t\n", areBracketsBalanced("{[(])}", runeMappings, openCloseMappings))
	fmt.Printf("{[(]]] is balanced: %t\n", areBracketsBalanced("{[(]]]", runeMappings, openCloseMappings))
	fmt.Printf("}[](){ is balanced: %t\n", areBracketsBalanced("}[](){", runeMappings, openCloseMappings))
	fmt.Printf("{] is balanced: %t\n", areBracketsBalanced("{]", runeMappings, openCloseMappings))
}

type bracket struct {
	shape  string
	isOpen bool
}

func areBracketsBalanced(brackets string, runeMap map[rune]bracket, openCloseMap map[bracket]bracket) bool {
	// Create a slice of the inputs
	bracketsAsSlice := []rune(brackets)
	if len(bracketsAsSlice) == 0 {
		return true
	}

	// create stack
	stack := snq.Stack[bracket]{}

	// Add first element manually
	firstElement := runeMap[bracketsAsSlice[0]]
	if firstElement.isOpen {
		stack.Push(firstElement)
	} else {
		return false
	}

	// Iterate over the rest, if rune at current bracket matches the top of the stack, pop, otherwise push
	for i := 1; i < len(bracketsAsSlice); i++ {
		currentBracket := runeMap[bracketsAsSlice[i]]
		topOfStack, _ := stack.Peek()

		if openCloseMap[currentBracket].shape == topOfStack.shape && !currentBracket.isOpen {
			stack.Pop()
		} else {
			stack.Push(currentBracket)
		}
	}

	return stack.Size() == 0
}
