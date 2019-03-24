package mergesortedlists

import (
	"reflect"
	"testing"
)

func TestIntToListNode(t *testing.T) {
	want := 4
	actual := IntToListNode(want)

	//actual should be a ListNode pointer
	if reflect.ValueOf(actual).Kind() != reflect.Ptr {
		t.Errorf("Want a pointer, got %v instead", reflect.ValueOf(actual).Kind())
	}

	if actual.Val != want {
		t.Errorf("ListNode.Val should be %v, but got %v instead.", want, actual.Val)
	}
}

func TestIntSliceToSingly(t *testing.T) {
	sliceToSinglyTests := []struct {
		values []int
	}{
		{[]int{1, 4, 5}},
		{[]int{}},
	}

	for _, tt := range sliceToSinglyTests {

		actual := IntSliceToSingly(tt.values)

		//actual should be a ListNode pointer
		if reflect.ValueOf(actual).Kind() != reflect.Ptr {
			t.Errorf("Want a pointer, got %v instead", reflect.ValueOf(actual).Kind())
		}

		//actual.Next should not be nil if len > 1
		if len(tt.values) > 1 && actual.Next == nil {
			t.Errorf("actual.Next should not be nil when length of tt.values is more than 1.")
		}

		node := actual
		for i := 0; i < len(tt.values); i++ {
			if node.Val != tt.values[i] {
				t.Errorf("Expected %v, got %v ", tt.values[i], node.Val)
			}

			if i == len(tt.values)-1 && node.Next != nil {
				t.Errorf("Last element.Next should be nil, index:%v, len:%v", i, len(tt.values))
			}

			node = node.Next
		}
	}
}

func TestSinglyToIntList(t *testing.T) {
	tests := []struct {
		wanted []int
	}{
		{[]int{1, 4, 5}},
		{[]int{}},
	}

	for _, tt := range tests {
		head := IntSliceToSingly(tt.wanted)
		actual := SinglyToIntList(head)

		if !reflect.DeepEqual(actual, tt.wanted) {
			t.Errorf("Wanted %v, got %v", tt.wanted, actual)
		}
	}
}

func TestEnqueue(t *testing.T) {
	tests := []struct {
		values   []*ListNode
		expected []int
	}{
		{
			[]*ListNode{
				&ListNode{1, nil}, &ListNode{0, nil}, &ListNode{2, nil},
			},
			[]int{0, 1, 2},
		},
		{
			[]*ListNode{
				&ListNode{41, nil}, &ListNode{39, nil},
				&ListNode{33, nil}, &ListNode{18, nil},
				&ListNode{27, nil}, &ListNode{12, nil},
			},
			[]int{12, 27, 18, 41, 33, 39},
		},
	}
	for _, tt := range tests {
		pq := PriorityQueue{}

		//Enqueue nodes in test cases
		for _, node := range tt.values {
			pq.Enqueue(node)
		}

		actual := NodesToInts(pq.nodes)

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Want %v, got %v.", tt.expected, actual)
		}
	}
}

func TestNodesToInts(t *testing.T) {
	tests := []struct {
		values   []*ListNode
		expected []int
	}{
		{
			[]*ListNode{
				&ListNode{41, nil}, &ListNode{39, nil},
				&ListNode{33, nil}, &ListNode{18, nil},
				&ListNode{27, nil}, &ListNode{12, nil},
			},
			[]int{41, 39, 33, 18, 27, 12},
		},
	}

	for _, tt := range tests {
		actual := NodesToInts(tt.values)

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Wanted %v, got %v", tt.expected, actual)
		}
	}
}
