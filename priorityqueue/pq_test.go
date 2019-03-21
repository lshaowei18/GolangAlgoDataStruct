package priorityqueue

import (
	"reflect"
	"testing"
)

var enqueueTests = []struct {
	values   []Node
	expected []Node
}{
	{[]Node{{1, ""}, {0, ""}, {2, ""}}, []Node{{0, ""}, {1, ""}, {2, ""}}},
	{[]Node{{41, ""}, {39, ""}, {33, ""}, {18, ""}, {27, ""}, {12, ""}},
		[]Node{{12, ""}, {27, ""}, {18, ""}, {41, ""}, {33, ""}, {39, ""}},
	},
}

func TestEnqueue(t *testing.T) {
	for _, tt := range enqueueTests {
		pq := PriorityQueue{}

		//Enqueue nodes in test cases
		for _, node := range tt.values {
			pq.Enqueue(node.priority, node.value)
		}

		actual := pq.nodes
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Want %v, got %v.", tt.expected, actual)
		}
	}
}

var swapTests = []struct {
	index1   int
	index2   int
	values   []Node
	expected []Node
}{
	{
		1,
		3,
		[]Node{{0, ""}, {1, ""}, {2, ""}, {3, ""}, {4, ""}},
		[]Node{{0, ""}, {3, ""}, {2, ""}, {1, ""}, {4, ""}},
	},
}

func TestSwap(t *testing.T) {
	for _, tt := range swapTests {
		pq := PriorityQueue{tt.values}
		pq.swap(tt.index1, tt.index2)
		actual := pq.nodes

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Want %v, got %v", tt.expected, actual)
		}
	}
}

var lowerPriorityTests = []struct {
	i        int
	j        int
	values   []Node
	expected bool
}{
	{
		1,
		3,
		[]Node{{0, ""}, {3, ""}, {2, ""}, {1, ""}, {4, ""}},
		true,
	},
}

func TestLowerPriority(t *testing.T) {
	for _, tt := range lowerPriorityTests {
		pq := PriorityQueue{tt.values}
		actual := pq.lowerPriority(tt.i, tt.j)

		if actual != tt.expected {
			t.Errorf("Want %v, got %v", tt.expected, actual)
		}
	}
}

var dequeueTests = []struct {
	values       []Node
	expected     []Node
	expectedNode Node
}{
	{
		[]Node{{12, ""}, {27, ""}, {18, ""}, {41, ""}, {33, ""}, {39, ""}},
		[]Node{{18, ""}, {27, ""}, {39, ""}, {41, ""}, {33, ""}},
		Node{12, ""},
	},
	{
		[]Node{},
		[]Node{},
		Node{},
	},
}

func TestDequeue(t *testing.T) {
	for _, tt := range dequeueTests {
		actual := PriorityQueue{tt.values}

		node := actual.Dequeue()

		if node != tt.expectedNode {
			t.Errorf("Want %v, got %v", tt.expectedNode, node)
		}

		if !reflect.DeepEqual(actual.nodes, tt.expected) {
			t.Errorf("Want %v, got %v", tt.expected, actual.nodes)
		}
	}
}
