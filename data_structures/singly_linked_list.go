package data_structures

import "fmt"

type SinglyLinkedListNode struct {
	Val  int
	Next *SinglyLinkedListNode
}

var head *SinglyLinkedListNode

func NewLinkedListNode(value int) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{
		Val:  value,
		Next: nil,
	}
}

func Push(value int) {
	// Create a new node with value as the input
	newNode := NewLinkedListNode(value)
	// if there is no head node assigned then assign the head with this newNode and return
	if head == nil {
		head = newNode
		return
	}
	// if head is present, that implies the list exists.
	// assign the head to a temp node because we do not want to change the original head pointer anywhere
	tempHead := head
	// Tracking a previous node so that once our loop ends, this would point to the last node in the linked list
	prev := head
	for tempHead != nil {
		// previous is now the current node
		prev = tempHead
		// change the current node to its next node
		tempHead = tempHead.Next
	}
	// the insertion of newNode is simply assignment to the Next of last node in the linked list, which is pointed by prev
	prev.Next = newNode
}

func Print() {
	tempHead := head
	fmt.Println("The list is as following:")
	for tempHead != nil {
		fmt.Print(tempHead.Val)
		tempHead = tempHead.Next
	}
	fmt.Println("")
}
