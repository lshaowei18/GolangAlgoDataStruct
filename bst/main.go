package bst

import "reflect"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	root Node
}

func (bst *BinarySearchTree) Insert(val int) {
	new := Node{val, nil, nil}

	if reflect.DeepEqual(bst.root, Node{}) {
		bst.root = new
	}
}
