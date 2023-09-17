package data_structures

import "fmt"

type bstNode struct {
	value int
	left  *bstNode
	right *bstNode
}

type BinarySearchTree struct {
	Root *bstNode
	Len  int
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{}
}

func InsertRec(rootNode *bstNode, v int) *bstNode {

	if rootNode == nil {
		return &bstNode{
			value: v,
		}
	}

	if v <= rootNode.value {
		rootNode.left = InsertRec(rootNode.left, v)
	} else if v > rootNode.value {
		rootNode.right = InsertRec(rootNode.right, v)
	}
	return rootNode
}

func InOrder(rootNode *bstNode) {
	if rootNode == nil {
		return
	}

	InOrder(rootNode.left)
	fmt.Println(rootNode.value)
	InOrder(rootNode.right)

}
