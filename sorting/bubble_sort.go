package main

import (
	"fmt"
)

func main() {
	slice := []int{301, 34, 583, 124, 246, 767, 139, -416, 415, 877, -57, 372, -61, -55, -34, 353, 292, -644, 461, 258}
	fmt.Println("Unsorted: ", slice)
	bubblesort(slice)
	fmt.Println("Sorted: ", slice)
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
