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

func MiddleNode() int {
	if head == nil {
		return -1
	}
	mid := LengthOfList() / 2
	tempHead := head
	counter := 0
	for counter < mid {
		tempHead = tempHead.Next
		counter++
	}
	return tempHead.Val
}

func LengthOfList() int {
	tempHead := head
	counter := 0
	for tempHead != nil {
		counter++
		tempHead = tempHead.Next
	}
	return counter
}

func EvenValues() int {
	tempHead := head
	counter := 0
	for tempHead != nil {
		if tempHead.Val%2 == 0 {
			counter++
		}
		tempHead = tempHead.Next
	}
	return counter
}

func Push(value int) {
	newNode := NewLinkedListNode(value)

	if head == nil {
		head = newNode
		return
	}

	tempHead := head

	prev := head
	for tempHead != nil {
		prev = tempHead
		tempHead = tempHead.Next
	}
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
