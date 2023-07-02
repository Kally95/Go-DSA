package sorting

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

func QuickSort(array []int, lb int, ub int) {
	if lb < ub {
		loc := partition(array, lb, ub)
		QuickSort(array, lb, loc-1)
		QuickSort(array, loc+1, ub)
	}
}
