package main

import (
	"fmt"

	ds "github.com/Kally95/Go-DSA/data_structures"
)

func main() {
	myList := ds.LinkedListNew{}
	myList.Prepend(1)
	myList.Prepend(2)
	myList.Prepend(3)
	myList.Prepend(154)
	myList.Prepend(123)
	fmt.Println(myList.Length)
	myList.InsertAtElement(111, 4)
	myList.PrintLL()

	// ds.Push(1)
	// ds.Push(2)
	// ds.Push(3)
	// ds.Push(4)
	// ds.Push()
	// ds.Print()
	// fmt.Println(ds.EvenValues())
	// fmt.Println(ds.MiddleNode())
}
