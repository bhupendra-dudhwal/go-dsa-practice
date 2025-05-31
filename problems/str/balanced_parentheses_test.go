package str

import "testing"

func TestIsBalanced(t *testing.T) {
	checker := NewBalancedParentheses()

	tests := []struct {
		input    string
		expected bool
	}{
		// Basic valid cases
		{"()", true},
		{"[]", true},
		{"{}", true},
		{"()[]{}", true},
		{"([{}])", true},
		{"( [ { } ( ) ] )", true},

		// Nested and mixed valid cases
		{"(((())))", true},
		{"{[()()]}", true},
		{"{[(())]}", true},

		// Invalid cases
		{"(", false},
		{")", false},
		{"[)", false},
		{"(]", false},
		{"({[)]}", false},
		{"[", false},
		{"]", false},
		{"[ ( )", false},
		{"([)]", false},

		// Edge cases
		{"", true},              // empty input is balanced
		{"abc", true},           // no brackets = balanced
		{"a(b)c", true},         // non-bracket characters ignored
		{"{[a+b]*(x+y)}", true}, // complex expression with brackets
	}

	for _, test := range tests {
		result := checker.IsBalanced(test.input)
		if result != test.expected {
			t.Errorf("IsBalanced(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
