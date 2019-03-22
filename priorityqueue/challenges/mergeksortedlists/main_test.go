package mergesortedlists

import (
	"reflect"
	"testing"
)

func TestIntToListNode(t *testing.T) {
	want := 4
	ln := IntToListNode(want)

	//Ln should be a ListNode pointer
	if reflect.ValueOf(ln).Kind() != reflect.Ptr {
		t.Errorf("Should return a pointer, got %v instead", reflect.ValueOf(ln).Kind())
	}
}
