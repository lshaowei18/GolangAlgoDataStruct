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

func NodesToInts(nodes []*ListNode) []int {
	var ints []int
	for _, node := range nodes {
		ints = append(ints, node.Val)
	}
	return ints
}

func (pq *PriorityQueue) Enqueue(node *ListNode) {
	pq.nodes = append(pq.nodes, node)
	i := len(pq.nodes) - 1 // index of child

	for {
		p := (i - 1) / 2 //index of parent

		if pq.nodes[i].Val >= pq.nodes[p].Val {
			break
		}
		//Swap
		pq.nodes[i], pq.nodes[p] = pq.nodes[p], pq.nodes[i]

		i = p
	}
}

func (pq *PriorityQueue) EnqueueSingly(head *ListNode) {
	node := head
	for {
		if node == nil {
			break
		}
		pq.Enqueue(node)
		node = node.Next
	}
}
