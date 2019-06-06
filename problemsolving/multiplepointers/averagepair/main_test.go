package averagepair

import "testing"

func TestAveragePair(t *testing.T) {
	tests := []struct {
		arr      []int
		avg      float64
		expected bool
	}{
		{[]int{1, 2, 3}, 2.5, true},
		{[]int{1, 3, 3, 5, 6, 7, 10, 12, 19}, 8, true},
		{[]int{-1, 0, 3, 4, 5, 6}, 4.1, false},
		{[]int{}, 4, false},
	}
	for _, tt := range tests {
		actual := averagePair(tt.arr, tt.avg)
		if actual != tt.expected {
			t.Errorf("arr:%v, actual:%v, expected:%v", tt.arr, actual, tt.expected)
		}
	}
}
