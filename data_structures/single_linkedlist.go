package data_structures

import "fmt"

type node struct {
	Value int
	Next  *node
}

type LinkedListNew struct {
	Head   *node
	Tail   *node
	Length int
}

func (l *LinkedListNew) Prepend(value int) {
	newNode := node{Value: value}
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
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}

func (l *LinkedListNew) PrependTail(value int) {
	newNode := node{Value: value}
	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
	} else {
		l.Tail.Next = &newNode
		l.Tail = &newNode
	}
}
