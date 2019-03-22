package mergesortedlists

import (
	"reflect"
	"testing"
)

func TestIntToListNode(t *testing.T) {
	want := 4
	actual := IntToListNode(want)

	//Ln should be a ListNode pointer
	if reflect.ValueOf(actual).Kind() != reflect.Ptr {
		t.Errorf("Want a pointer, got %v instead", reflect.ValueOf(actual).Kind())
	}

	if actual.Val != want {
		t.Errorf("ListNode.Val should be %v, but got %v instead.", want, actual.Val)
	}
}
