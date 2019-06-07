package issubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		sub      string
		str      string
		expected bool
	}{
		{"hello", "hello world", true},
		{"sing", "sting", true},
		{"abc", "abracadabra", true},
		{"abc", "acb", false},
		{"abc", "abc", true},
		{"abc", "abdc", true},
		{"abc", "dabdc", true},
		{"bc", "b", false},
		{"bc", "bbb", false},
		{"bc", "bbbc", true},
	}
	for _, tt := range tests {
		actual := isSubsequence(tt.sub, tt.str)
		if actual != tt.expected {
			t.Errorf("s1:%v, s2:%v, actual:%v, expected:%v",
				tt.sub, tt.str, actual, tt.expected)
		}
	}
}
