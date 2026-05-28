package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	// Set up data for bench marking and data capture.
	linearSearchResults := make(map[int]time.Duration)
	binarySearchResults := make(map[int]time.Duration)

	benchmarkSizes := []int{1000, 100000, 10000000, 10000000, 1000000000}

	benchmarkingData := make(map[int][]int)

	for _, benchmarkSize := range benchmarkSizes {
		benchmarkingData[benchmarkSize] = produceSortedArray(benchmarkSize)
	}

	for arraySize, arr := range benchmarkingData {
		// Run linear search and binary search on arrays of different sizes and capture results.
		// We are searching for the last element in the array to ensure we are testing the worst-case scenario for linear search.
		linearSearchResults[arraySize] = benchmark(
			func() {
				linearSearch(arr, arraySize)
			})

		binarySearchResults[arraySize] = benchmark(
			func() {
				binarySearch(arr, arraySize)
			})
	}

	prettyPrintBenchmarkResults(linearSearchResults, binarySearchResults)

}

// Linear search - takes an array and returns index of target
func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

// Binary-search - takes a sorted array and returns index of target
func binarySearch(arr []int, target int) int {
	hi := len(arr) - 1
	lo := 0

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] > target {
			// target is in the left half
			hi = mid - 1
		} else if arr[mid] < target {
			// target in the right half
			lo = mid + 1
		} else {
			return mid
		}
	}

	// Target not in array
	return -1
}

func produceSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

// Closure to time method calls
func benchmark(function func()) time.Duration {
	start := time.Now()
	function()
	return time.Since(start)
}

func prettyPrintBenchmarkResults(linearResults, binaryResults map[int]time.Duration) {
	// Extract and sort the array sizes
	var sizes []int
	for size := range linearResults {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)

	fmt.Printf("%-15s %-20s %-20s %-20s\n", "Array Size", "Linear Search Time", "Binary Search Time", "Speed Factor (Binary is X times faster)")

	// Print results in sorted order
	for _, size := range sizes {
		linearTime := linearResults[size]
		binaryTime := binaryResults[size]

		// Calculate speed factor (how many times faster binary search is)
		var speedFactor float64
		if binaryTime > 0 {
			speedFactor = float64(linearTime) / float64(binaryTime)
		} else {
			speedFactor = 0
		}

		fmt.Printf("%-15d %-20s %-20s %-20.2f\n", size, linearTime, binaryTime, speedFactor)
	}
}
