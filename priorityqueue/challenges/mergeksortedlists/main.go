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

func (pq *PriorityQueue) Dequeue() *ListNode {
	if len(pq.nodes) == 0 {
		return nil
	}
	dequeued := pq.nodes[0]

	pq.nodes[0] = pq.nodes[len(pq.nodes)-1] //Replace last node with first node
	pq.nodes = pq.nodes[:len(pq.nodes)-1]   //Remove last node

	i := 0 //  of previously last node

	for {
		n := len(pq.nodes)

		left := i*2 + 1

		//Once left child is out of range, both children will be
		if left >= n {
			break
		}

		swap := left

		// right is not out of range and less than left
		if right := left + 1; right < n && pq.less(right, left) {
			swap = right
		}

		//if i is higher than swap child, break
		if pq.less(i, swap) {
			break
		}

		pq.swap(i, swap)

		i = swap
	}
	return dequeued
}

//Higher number has lower priority
func (pq *PriorityQueue) less(i, j int) bool {
	return pq.nodes[i].Val < pq.nodes[j].Val
}

func (pq *PriorityQueue) swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
}
