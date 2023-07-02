package main

import (
	"fmt"

	ds "github.com/Kally95/Go-DSA/data_structures"
	"github.com/Kally95/Go-DSA/search"
)

func main() {
	myList := ds.LinkedList{}
	myList.Prepend(1)
	myList.Prepend(2)
	myList.Prepend(3)
	myList.Prepend(154)
	myList.Prepend(123)
	myList.PrintLinkedList()

	array := []int{1, 3, 6, 8, 11, 13, 16, 19, 20}
	target := 11
	fmt.Println(search.BinarySearch(array, target))
}
