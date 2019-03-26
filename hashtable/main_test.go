package hashtable

import "testing"

func TestHash(t *testing.T) {
	tests := []struct {
		key    string
		length int
		want   int
	}{
		{"hello", 13, 7},
		{"goodbye", 13, 9},
		{"cyan", 13, 5},
		{"hi", 13, 10},
	}

	for _, tt := range tests {
		got := Hash(tt.key, tt.length)

		if got != tt.want {
			t.Errorf("Got %v, want %v", got, tt.want)
		}
	}
}
