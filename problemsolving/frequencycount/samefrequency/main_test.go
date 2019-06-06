package samefrequency

import "testing"

func TestSameFrequency(t *testing.T) {
	tests := []struct {
		n1       int
		n2       int
		expected bool
	}{
		{182, 281, true},
		{34, 14, false},
		{3589578, 5879385, true},
		{22, 222, false},
	}

	for _, tt := range tests {
		actual := sameFrequency(tt.n1, tt.n2)
		if actual != tt.expected {
			t.Errorf("n1: %v, n2: %v, actual: %v, expected: %v",
				tt.n1, tt.n2, actual, tt.expected)
		}
	}
}
