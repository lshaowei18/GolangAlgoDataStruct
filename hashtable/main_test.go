package hashtable

import "testing"

func TestHash(t *testing.T) {
	tests := []struct {
		key  string
		want int
	}{
		{"hello", 40},
		{"goodbye", 23},
		{"cyan", 23},
		{"hi", 45},
	}

	for _, tt := range tests {
		ht := HashTable{}
		got := ht.Hash(tt.key)

		if got != tt.want {
			t.Errorf("Got %v, want %v", got, tt.want)
		}
	}
}
