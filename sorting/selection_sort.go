package main

import "fmt"

func swap(array []int, firstIndex int, secondIndex int) {
	temp := array[firstIndex]
	array[firstIndex] = array[secondIndex]
	array[secondIndex] = temp
}

func indexOfMin(array []int, startIndex int) int {
	minValue := array[startIndex]
	minIndex := startIndex

	for i := minIndex + 1; i < len(array); i++ {
		if array[i] < minValue {
			minIndex = i
			minValue = array[i]
		}
	}
	return minIndex
}

func selectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		minVal := indexOfMin(array, i)
		swap(array, i, minVal)
	}
}

func main() {
	array := []int{3, 5, 6, 1, 8, 2, 9}
	selectionSort(array)
	fmt.Println(array)
}
