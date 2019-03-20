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
