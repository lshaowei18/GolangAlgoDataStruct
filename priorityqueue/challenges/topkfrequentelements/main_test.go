package topkfrequent

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{4, 1, -1, 2, -1, 2, 3},
			2,
			[]int{-1, 2},
		},
	}

	for _, tt := range tests {

		got := topKFrequent(tt.input, tt.k)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Want %v, got %v", tt.want, got)
		}
	}
}
