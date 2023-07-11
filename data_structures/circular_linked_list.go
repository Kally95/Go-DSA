package data_structures

import "fmt"

type node struct {
	Value    int
	Next     *node
	Previous *node
}

type CircularLinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (cll *CircularLinkedList) PrependCLL(value int) {
	newNode := node{Value: value}
	if cll.Length == 0 {
		cll.Head = &newNode
		cll.Tail = &newNode
		cll.Length++
	} else {
		cll.Tail.Next = &newNode
		cll.Tail = &newNode
		newNode.Next = cll.Head
		cll.Length++
	}
}

func (cll *CircularLinkedList) DisplayCircularLL() {
	if cll.Length == 0 {
		return
	}

	currentNode := cll.Head
	for {
		fmt.Print(currentNode.Value)
		currentNode = currentNode.Next

		if currentNode == cll.Head {
			break
		}

		fmt.Print(" => ")
	}
	fmt.Println()
}

func (cll *CircularLinkedList) HasCycle() bool {
	if cll.Head == nil {
		return false
	}
	fast := cll.Head.Next
	slow := cll.Head
	for fast != nil && slow != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
