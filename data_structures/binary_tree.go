package data_structures

import (
	"fmt"
)

type BtNode struct {
	Data int
	Left *BtNode
	Right *BtNode
 }
 
 func (root *BtNode)PreOrderTraversal(){
	if root !=nil{
	   fmt.Printf("%d ", root.Data)
	   root.Left.PreOrderTraversal()
	   root.Right.PreOrderTraversal()
	}
	return
 }