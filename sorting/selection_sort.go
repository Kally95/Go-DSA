package sorting

func swap1(array []int, firstIndex int, secondIndex int) {
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

func SelectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		minVal := indexOfMin(array, i)
		swap1(array, i, minVal)
	}
}
