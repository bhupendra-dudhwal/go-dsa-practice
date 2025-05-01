package numbers

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "Non-palindrome: ends with 0",
			num:      10,
			expected: false,
		},
		{
			name:     "Palindrome: odd digits",
			num:      101,
			expected: true,
		},
		{
			name:     "Palindrome: even digits",
			num:      11,
			expected: true,
		},
		{
			name:     "Non-palindrome: negative number",
			num:      -121,
			expected: false,
		},
		{
			name:     "Single-digit number",
			num:      7,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.num)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%d)=%t; expected %t", tt.num, result, tt.expected)
			}
		})
	}
}
