package priorityqueue

type Node struct {
	priority int
	value    string
}

type PriorityQueue struct {
	nodes []Node
}

func (pq *PriorityQueue) Enqueue(priority int, value string) {
	node := Node{priority, value}
	pq.nodes = append(pq.nodes, node)

	index := len(pq.nodes) - 1

	for {
		parentIndex := (index - 1) / 2
		if index == parentIndex || pq.lowerPriority(index, parentIndex) {
			break
		}
		//Swap Values
		pq.swap(index, parentIndex)

		//Increment
		index = parentIndex
	}
}

func (pq *PriorityQueue) Dequeue() Node {
	//Returns empty node if pq is empty
	if len(pq.nodes) == 0 {
		return Node{}
	}

	dequeued := pq.nodes[0]

	pq.nodes[0] = pq.nodes[len(pq.nodes)-1] //Replace last node with first node
	pq.nodes = pq.nodes[:len(pq.nodes)-1]   //Remove last node

	index := 0

	for {
		n := len(pq.nodes)

		left := index*2 + 1
		//Check if children are out of range, once left is out right will be too
		if left >= n {
			break
		}
		swap := left

		// right is not out of range and has higher priority than left
		if right := left + 1; right < n && !pq.lowerPriority(right, left) {
			swap = right
		}

		//if index has higher priority than child, break
		if !pq.lowerPriority(index, swap) {
			break
		}

		pq.swap(index, swap)

		index = swap
	}

	return dequeued
}

func (pq *PriorityQueue) swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
}

//Higher number has lower priority
func (pq *PriorityQueue) lowerPriority(i, j int) bool {
	return pq.nodes[i].priority > pq.nodes[j].priority
}
