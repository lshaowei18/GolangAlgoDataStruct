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
