package main

import (
	"fmt"
	"strings"

	ht "github.com/grep-hammerspace/coding-curriculum/module4/hashtable"
)

func main() {
	myMap := ht.New[string, int]()

	myMap.Put("apple", 5)
	myMap.Put("banana", 3)
	myMap.Put("cherry", 4)

	fmt.Printf(myMap.Stringify())

	appleInt, err := myMap.Get("apple")
	if err != nil {
		fmt.Printf("Key 'apple' not found")
	}

	bananaInt, err := myMap.Get("banana")
	if err != nil {
		fmt.Printf("Key 'banana' not found")
	}
	fmt.Printf("Key 'apple' has value %d\n", appleInt)
	fmt.Printf("Key 'banana' has value %d\n", bananaInt)

	fmt.Println("Removing by key 'apple' and 'cherry'")
	myMap.Remove("apple")
	myMap.Remove("cherry")

	fmt.Printf(myMap.Stringify())
}

func CountWordsInText(source string) string {
	sourceAsArray := strings.Split(source, " ")

	wordCountMap := ht.New[string, int]()

	for _, word := range sourceAsArray {
		wordCountMap.Put(word, 5)
	}

	return ""
}
