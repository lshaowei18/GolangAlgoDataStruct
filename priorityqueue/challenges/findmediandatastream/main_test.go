package findmedian

import (
	"reflect"
	"testing"
)

func TestAddNum(t *testing.T) {
	tests := []struct {
		values []int
		want   []int
	}{
		{
			[]int{3, 2, 4},
			[]int{2, 3, 4},
		},
		{
			[]int{6, 10, 2, 6, 5, 0, 6, 3, 1, 0, 0},
			[]int{0, 0, 2, 3, 0, 6, 6, 10, 5, 6, 1},
		},
	}

	for _, tt := range tests {
		mf := Constructor()

		for _, n := range tt.values {
			mf.AddNum(n)
		}

		got := mf.elements

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Want %v, got %v", tt.want, got)
		}
	}
}

func TestFindMedian(t *testing.T) {
	tests := []struct {
		values []int
		want   float64
	}{
		{
			[]int{2, 3, 4},
			3,
		},
		{
			[]int{2, 3},
			2.5,
		},
		{
			[]int{6, 10, 2, 6, 5, 0, 6, 3, 1, 0, 0},
			3,
		},
	}

	for _, tt := range tests {
		mf := Constructor()
		mf.elements = tt.values

		got := mf.FindMedian()

		if got != tt.want {
			t.Errorf("Got %v, want %v", got, tt.want)
		}
	}
}
