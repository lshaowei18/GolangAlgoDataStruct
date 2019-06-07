package minsubarraylen

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		arr      []int
		num      int
		expected int
	}{
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		{[]int{2, 1, 6, 5, 4}, 9, 2},
		{[]int{3, 1, 7, 11, 2, 9, 8, 21, 62, 33}, 52, 1},
		{[]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 95, 0},
	}
	for _, tt := range tests {
		actual := minSubArrayLen(tt.arr, tt.num)
		if actual != tt.expected {
			t.Errorf("arr: %v, num: %v, actual:%v, expected:%v",
				tt.arr, tt.num, actual, tt.expected)
		}
	}
}
