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
	parentIndex := (index - 1) / 2

	for pq.nodes[index].priority < pq.nodes[parentIndex].priority {
		//Swap Values
		pq.swap(index, parentIndex)

		//Increment
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

func (pq *PriorityQueue) swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
}
