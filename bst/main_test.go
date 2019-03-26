package bst

import "testing"

func TestInsert(t *testing.T) {
	bst := BinarySearchTree{}

	values := []int{10, 6, 15}

	want := values[0]
	bst.Insert(want)
	got := bst.root.Val

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}

	want = values[1]
	bst.Insert(want)
	got = bst.root.Left.Val

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
