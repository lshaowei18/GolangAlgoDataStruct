package bst

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(val int) {
	new := Node{val, nil, nil}

	if bst.root == nil {
		bst.root = &new
		return
	}

	var node *Node

	node = bst.root

	for {
		if val < node.Val {
			if node.Left != nil {
				node = node.Left
				continue
			}
			node.Left = &new
			break
		}

		if node.Right != nil {
			node = node.Right
			continue
		}
		node.Right = &new
		break

		if val < node.Val && node.Left != nil {
			node = node.Left
			continue
		} else if val > node.Val && node.Right != nil {
			node = node.Right
			continue
		}

	}
}
