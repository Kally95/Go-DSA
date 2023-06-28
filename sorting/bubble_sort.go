package main

// func main() {
// 	slice := []int{301, 34, 583, 124, 246, 767, 139, -416, 415, 877, -57, 372, -61, -55, -34, 353, 292, -644, 461, 258}
// 	fmt.Println("Unsorted: ", slice)
// 	bubblesort2(slice)
// 	fmt.Println("Sorted: ", slice)
// }

func bubblesort(items []int) {

	n := len(items)

	for j := 0; j < n; j++ {
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
			}
		}
	}
}

func bubblesort2(items []int) {

	n := len(items)

	for j := 0; j < n; j++ {
		for i := 1; i < n; i++ {
			if items[i] < items[i-1] {
				items[i-1], items[i] = items[i], items[i-1]
			}
		}
	}
}
