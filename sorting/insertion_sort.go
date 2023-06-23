package main

import "fmt"

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
}

func main() {
	array := []int{5, 4, 10, 1, 6, 2}
	insertionSort(array)
	fmt.Println(array)
}
