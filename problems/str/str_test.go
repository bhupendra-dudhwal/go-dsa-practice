package str

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Odd length palindrome", "madam", true},
		{"Even length palindrome", "abba", true},
		{"Not a palindrome", "hello", false},
		{"Single character", "a", true},
		{"Empty string", "", true},
		{"Repeated emoji", "😊😊", true},
		{"Emoji palindrome", "😊🚀😊", true},
		{"Emoji non-palindrome", "😊hello😊", false},
		{"Japanese characters", "あいいあ", true},
		{"Emoji faces", "😀🙃😀", true},
		{"Case sensitive phrase", "A man a plan a canal Panama", false},
		{"Numeric palindrome", "12321", true},
		{"Palindrome with space", "12 21", true},
		{"False numeric with space", "12 321", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%s)=%t; expected %t", tt.input, result, tt.expected)
			}
		})
	}
}
