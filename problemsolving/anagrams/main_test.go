package anagram

import "testing"

func TestAnagram(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{
			"", "", true,
		},
		{
			"aaz", "zza", false,
		},
		{
			"anagram", "nagaram", true,
		},
		{
			"rat", "car", false,
		},
		{
			"qwerty", "qeywrt", true,
		},
		{
			"qwerty", "qeyw", false,
		},
	}
	for _, tt := range tests {
		t.Run("First Attempt", func(t *testing.T) {
			actual := validAnagram(tt.s1, tt.s2)

			if actual != tt.expected {
				t.Errorf("s1: %v, s2: %v, actual: %v, expected: %v",
					tt.s1, tt.s2, actual, tt.expected)
			}
		})
		t.Run("Second Attempt", func(t *testing.T) {
			actual := alternative(tt.s1, tt.s2)

			if actual != tt.expected {
				t.Errorf("s1: %v, s2: %v, actual: %v, expected: %v",
					tt.s1, tt.s2, actual, tt.expected)
			}
		})
	}
}
