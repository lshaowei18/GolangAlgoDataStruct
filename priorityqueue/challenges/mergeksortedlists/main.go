package mergesortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue struct {
	nodes []*ListNode
}

func IntToListNode(n int) *ListNode {
	return &ListNode{4, nil}
}

func IntSliceToSingly(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	head := &ListNode{s[0], nil}

	node := head
	for i := 1; i < len(s); i++ {
		newNode := &ListNode{s[i], nil}
		node.Next = newNode
		node = newNode
	}
	return head
}

func SinglyToIntList(head *ListNode) []int {
	node := head
	s := []int{}
	for {
		if node == nil {
			break
		}
		s = append(s, node.Val)
		node = node.Next
	}
	return s
}
