package sorting

func Bubblesort(items []int) {

	n := len(items)

	for j := 0; j < n; j++ {
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
			}
		}
	}
}

func Bubblesort2(items []int) {

	n := len(items)

	for j := 0; j < n; j++ {
		for i := 1; i < n; i++ {
			if items[i] < items[i-1] {
				items[i-1], items[i] = items[i], items[i-1]
			}
		}
	}
}
