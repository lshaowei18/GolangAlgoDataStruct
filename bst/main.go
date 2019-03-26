package bst

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	root Node
}
