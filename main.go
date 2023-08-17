package main

import (
	"fmt"

	ds "github.com/Kally95/Go-DSA/data_structures"
)

func main() {
	// myList := ds.LinkedListNew{}
	// myList.Prepend(1)
	// myList.Prepend(2)
	// myList.Prepend(3)
	// myList.Prepend(154)
	// myList.Prepend(123)
	// myList.Prepend(122133)
	// fmt.Println(myList.Length)
	// myList.InsertAtElement(111, 4)
	// myList.PrintLL()

	// ds.Push(1)
	// ds.Push(2)
	// ds.Push(3)
	// ds.Push(4)
	// ds.Push()
	// ds.Print()
	// fmt.Println(ds.EvenValues())
	// fmt.Println(ds.MiddleNode())

	// dll := ds.DoubleLinkedList{}
	// dll.PrependDLL(5)
	// dll.PrependDLL(6)
	// dll.PrependDLL(7)
	// dll.PrependDLL(9)
	// dll.PrintDLL()
	// dll.PrintDLLReverse()

	// cll := ds.CircularLinkedList{}
	// cll.PrependCLL(5)
	// cll.PrependCLL(7)
	// cll.PrependCLL(8)
	// cll.PrependCLL(12)
	// cll.DisplayCircularLL()
	// fmt.Println(cll.HasCycle())

	stack := ds.NewStack[int]()
	stack.Push(10)
	stack.Push(15)
	stack.Push(1150)
	stack.Push(121)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())

}
