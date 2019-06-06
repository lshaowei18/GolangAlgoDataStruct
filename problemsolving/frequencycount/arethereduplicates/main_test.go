package arethereduplicate

import "testing"

func TestAreThereDuplicates(t *testing.T) {
	t.Run("Integers", func(t *testing.T) {
		tests := []struct {
			arr      []int
			expected bool
		}{
			{[]int{1, 2, 3}, false},
			{[]int{1, 2, 2}, true},
		}

		for _, tt := range tests {
			actual, _ := areThereDuplicates(tt.arr)
			if actual != tt.expected {
				t.Errorf("arr: %v, actual: %v, expected: %v",
					tt.arr, actual, tt.expected)
			}
		}
	})

	t.Run("Integers", func(t *testing.T) {
		tests := []struct {
			arr      []string
			expected bool
		}{
			{[]string{"a", "b", "c", "a"}, true},
		}

		for _, tt := range tests {
			actual, _ := areThereDuplicates(tt.arr)
			if actual != tt.expected {
				t.Errorf("arr: %v, actual: %v, expected: %v",
					tt.arr, actual, tt.expected)
			}
		}
	})
}
