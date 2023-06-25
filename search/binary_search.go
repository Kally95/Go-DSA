package main

import "fmt"

// Question
// You are given a sorted array of size N. You are given a target number to be found in the array.
// Return the index at which the target number occurs. If the target number is not found then return -1.

func binarySearch(array []int, target int) int {
	start := 0
	end := len(array) - 1

	for start <= end {
		middle := (start + end) / 2

		if target == array[middle] {
			return middle
		}
		if target < array[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return -1
}

func main() {
	array := []int{1, 3, 6, 8, 11, 13, 16, 19, 20}
	target := 11
	fmt.Println(binarySearch(array, target))
}
