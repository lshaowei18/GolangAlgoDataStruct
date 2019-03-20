package priorityqueue

type Node struct {
	priority int
	value    string
}

type PriorityQueue struct {
	nodes []Node
}
