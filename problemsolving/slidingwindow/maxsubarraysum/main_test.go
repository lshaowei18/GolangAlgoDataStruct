package maxsubarraysum

import "testing"

func TestMaxSubArraySum(t *testing.T) {
	tests := []struct {
		arr      []int
		num      int
		expected int
	}{
		{
			[]int{2, 6, 9, 2, 1, 8, 5, 6, 3}, 3, 19,
		},
		{
			[]int{-2, -6, -8, -1, -2, -1}, 3, -4,
		},
	}
	for _, tt := range tests {
		actual := maxSubArraySum(tt.arr, tt.num)
		if actual != tt.expected {
			t.Errorf("arr :%v, num: %v, actual: %v, expected: %v",
				tt.arr, tt.num, actual, tt.expected)
		}
	}
}
