package bst

import "testing"

func TestInsert(t *testing.T) {
	bst := BinarySearchTree{}

	want := 10

	bst.Insert(want)

	got := bst.root.Val

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
