package main

import (
	"fmt"
	"strings"

	ht "github.com/grep-hammerspace/coding-curriculum/module4/hashtable"
)

func main() {
	myMap := ht.NewHashTable[string, int]()

	myMap.Put("apple", 5)
	myMap.Put("banana", 3)
	myMap.Put("cherry", 4)

	fmt.Printf(myMap.Stringify())
}

func CountWordsInText(source string) string {
	sourceAsArray := strings.Split(source, " ")

	wordCountMap := ht.NewHashTable[string, int]()

	for _, word := range sourceAsArray {
		wordCountMap.Put(word, 5)
	}

}
