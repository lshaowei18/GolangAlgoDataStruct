package countuniquevalues

import "testing"

func TestCountUniqueValues(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{1, 1, 1, 1, 1, 2}, 2,
		},
		{
			[]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}, 7,
		},
		{
			[]int{}, 0,
		},
		{
			[]int{-2, -1, -1, 0, 1}, 4,
		},
	}

	for _, tt := range tests {
		actual := countUniqueValues(tt.nums)
		if actual != tt.expected {
			t.Errorf("%v, actual: %v, expected: %v", tt.nums, actual, tt.expected)
		}
	}
}
