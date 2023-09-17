package data_structures

import (
	"fmt"
)

type BtNode struct {
	Data  int
	Left  *BtNode
	Right *BtNode
}

func (root *BtNode) PreOrderTraversal() {
	if root != nil {
		fmt.Printf("%d ", root.Data)
		root.Left.PreOrderTraversal()
		root.Right.PreOrderTraversal()
	}
}

func (root *BtNode) InOrderTraversal() {
	if root != nil {
		root.Left.PreOrderTraversal()
		fmt.Printf("%d ", root.Data)
		root.Right.PreOrderTraversal()
	}
}

func (root *BtNode) PostOrderTraversal() {
	if root != nil {
		root.Left.PreOrderTraversal()
		root.Right.PreOrderTraversal()
		fmt.Printf("%d ", root.Data)
	}
}
