package findmedian

import (
	"reflect"
	"testing"
)

func TestAddNum(t *testing.T) {
	tests := []struct {
		values []int
	}{
		{
			[]int{2, 3, 4},
		},
	}

	for _, tt := range tests {
		mf := Constructor()

		for _, n := range tt.values {
			mf.AddNum(n)
		}

		got := mf.elements

		if !reflect.DeepEqual(got, tt.values) {
			t.Errorf("Want %v, got %v", tt.values, got)
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
