package data_structures

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type LinkedListNew struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *LinkedListNew) Prepend(value int) {
	newNode := Node{Data: value}
	if l.Length == 0 {
		l.Head = &newNode
		l.Tail = &newNode
		l.Length++
	} else {
		newNode.Next = l.Head
		l.Head = &newNode
		l.Length++
	}
}

func (l *LinkedListNew) PrintLL() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Print(currentNode.Data, "-> ")
		currentNode = currentNode.Next
	}
}

func (l *LinkedListNew) PrependTail(value int) {
	newNode := Node{Data: value}
	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
	} else {
		l.Tail.Next = &newNode
		l.Tail = &newNode
	}
}

func (l *LinkedListNew) InsertAtElement(value int, element int) {
	// 3 -> 5 -> 9 -> 8 -> 12
	// size is 5
	newNode := Node{Data: value}
	currentNode := l.Head
	counter := 0
	for counter <= element-1 {
		if counter+1 == element {
			newNode.Next = currentNode.Next
			currentNode.Next = &newNode
			return
		}
		counter++
		currentNode = currentNode.Next
	}

}
