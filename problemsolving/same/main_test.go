package same

import "testing"

func TestSame(t *testing.T) {
	tests := []struct {
		nums     []int
		squares  []int
		expected bool
	}{
		{
			[]int{1, 2, 3}, []int{4, 1, 9}, true,
		},
		{
			[]int{1, 2, 3}, []int{1, 9}, false,
		},
		{
			[]int{1, 2, 1}, []int{4, 4, 1}, false,
		},
		{
			[]int{1, 4, 4}, []int{16, 1, 16}, true,
		},
		{
			[]int{1, 2, 2}, []int{1, 4}, false,
		},
	}
	for _, tt := range tests {
		t.Run("First Try", func(t *testing.T) {
			actual := same(tt.nums, tt.squares)
			if actual != tt.expected {
				t.Errorf("actual: %v, expected: %v", actual, tt.expected)
			}
		})
	}
}
