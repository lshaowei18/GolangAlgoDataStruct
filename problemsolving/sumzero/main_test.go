package sumzero

import (
	"reflect"
	"testing"
)

func TestSumZero(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{-3, -2, -1, 0, 1, 2, 3}, []int{-3, 3},
		},
		{
			[]int{-2, 0, 1, 3}, []int{},
		},
		{
			[]int{1, 2, 3}, []int{},
		},
	}

	for _, tt := range tests {
		t.Run("First attempt", func(t *testing.T) {
			actual := sumZero(tt.nums)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("nums: %v, actual : %v, expected : %v",
					tt.nums, actual, tt.expected)
			}
		})

		t.Run("Multiple Pointer Solution", func(t *testing.T) {
			actual := multiplePointer(tt.nums)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("nums: %v, actual : %v, expected : %v",
					tt.nums, actual, tt.expected)
			}
		})
	}
}
