package data_structures

import "fmt"

type node struct {
	Value    int
	Next     *node
	Previous *node
}

type DoubleLinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (dll *DoubleLinkedList) PrependDLL(value int) {
	newNode := node{Value: value}
	if dll.Length == 0 {
		dll.Head = &newNode
		dll.Tail = &newNode
		dll.Length++
	} else {
		dll.Head.Previous = &newNode
		newNode.Next = dll.Head
		dll.Head = &newNode
	}
}

func (dll *DoubleLinkedList) PrintDLL() {
	if dll.Length == 0 {
		return
	}

	currentNode := dll.Head
	for currentNode != nil {
		fmt.Print(currentNode.Value)
		if currentNode.Next != nil {
			fmt.Print("=>")
		}
		currentNode = currentNode.Next
	}
}

func (dll *DoubleLinkedList) PrintDLLReverse() {
	if dll.Length == 0 {
		return
	}

	currentNode := dll.Tail
	for currentNode != nil {
		fmt.Print(currentNode.Value)
		if currentNode.Previous != nil {
			fmt.Print("=>")
		}
		currentNode = currentNode.Previous
	}
}
