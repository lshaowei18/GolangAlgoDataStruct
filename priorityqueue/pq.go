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

func (pq *PriorityQueue) swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
}

//Higher number has lower priority
func (pq *PriorityQueue) lowerPriority(i, j int) bool {
	return pq.nodes[i].priority > pq.nodes[j].priority
}
