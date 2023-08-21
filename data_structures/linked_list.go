package data_structures

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(value int) {
	newNode := Node{Value: value}
	if l.Head != nil {
		newNode.Next = l.Head
		l.Head = &newNode
		l.Length++
	} else {
		l.Head = &newNode
		l.Length++
	}
}

func (l *LinkedList) PrintLinkedList() {
	if l.Head == nil {
		return
	}
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}

func (l *LinkedList) DeleteNodeWithVal(value int) {
	if l.Length == 0 {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	previousToDelete := l.Head
	for previousToDelete.Next.Value != value {
		if previousToDelete.Next.Next == nil {
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.Length--
}
