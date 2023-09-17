package data_structures

func FibRecursion(n int) int {
	if n < 2 {
		return n
	}

	return FibRecursion(n-1) + FibRecursion(n-2)
}

func RecursiveBinarySearch(arr []int, start int, end int, n int) int {
	mid := (start + end) / 2

	if start > end {
		return -1
	}

	if arr[mid] == n {
		return mid
	}

	if n < arr[mid] {
		return RecursiveBinarySearch(arr, start, mid-1, n)
	}

	return RecursiveBinarySearch(arr, mid+1, end, n)

}
