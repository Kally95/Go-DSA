package data_structures

import (
	"fmt"
	"strconv"
)

type BtNode struct {
	value int
	left  *BtNode
	right *BtNode
}

type BinaryTree struct {
	Root *BtNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (n *BtNode) Insert(value int, side byte) {
	newNode := &BtNode{value: value}
	if side != 'L' && side != 'R' {
		return
	} else if n == nil {

	} else if side == 'L' {
        if n.left == nil {
            n.left = newNode
        } else {
            n.left.Insert(value, side)
        }
    } else {
        if n.right == nil {
            n.right = newNode
        } else {
            n.right.Insert(value, side)
        }
    }   
}

func TraverseBT(node *BtNode) {
	if node != nil {
		TraverseBT(node.left)
		fmt.Print(" " + strconv.Itoa(node.value))
		TraverseBT(node.right)
	}
}