package main

import "fmt"

func swap(array []int, firstIndex int, secondIndex int) {
	temp := array[firstIndex]
	array[firstIndex] = array[secondIndex]
	array[secondIndex] = temp
}

func partition(array []int, lb int, ub int) int {
	pivot := lb
	start := lb
	end := ub

	for start < end {

		for array[start] <= array[pivot] {
			start++
		}

		for array[end] > array[pivot] {
			end--
		}

		if start < end {
			swap(array, start, end)
		}
	}
	swap(array, lb, end)
	return end
}

func quickSort(array []int, lb int, ub int) {
	if lb < ub {
		loc := partition(array, lb, ub)
		quickSort(array, lb, loc-1)
		quickSort(array, loc+1, ub)
	}
}

func main() {
	array := []int{3, 5, 6, 1, 8, 2, 9}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
