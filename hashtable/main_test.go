package hashtable

import "testing"

func TestHash(t *testing.T) {
	tests := []struct {
		key    string
		length int
		want   int
	}{
		{"pink", 10, 0},
		{"orangered", 10, 7},
		{"cyan", 10, 3},
	}

	for _, tt := range tests {
		got := Hash(tt.key, tt.length)

		if got != tt.want {
			t.Errorf("Got %v, want %v", got, tt.want)
		}
	}
}
