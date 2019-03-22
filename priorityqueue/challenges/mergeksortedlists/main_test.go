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
	values := []int{1, 4, 5}
	actual := IntSliceToSingly(values)

	//actual should be a ListNode pointer
	if reflect.ValueOf(actual).Kind() != reflect.Ptr {
		t.Errorf("Want a pointer, got %v instead", reflect.ValueOf(actual).Kind())
	}

	//actual.Next should not be nil if len > 1
	if len(values) > 1 && actual.Next == nil {
		t.Errorf("actual.Next should not be nil when length of values is more than 1.")
	}

	node := actual
	for i := 0; i < len(values); i++ {
		if node.Val != values[i] {
			t.Errorf("Expected %v, got %v ", values[i], node.Val)
		}

		if i == len(values)-1 && node.Next != nil {
			t.Errorf("Last element.Next should be nil, index:%v, len:%v", i, len(values))
		}

		node = node.Next
	}
}
