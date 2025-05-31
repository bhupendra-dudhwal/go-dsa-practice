package str

// Problem Statement:
// Check if a given string containing just the characters '(', ')', '{', '}', '[' and ']'
// is balanced — meaning every open bracket has a corresponding closing bracket in correct order.

type parenthesesChecker struct {
	pairs map[rune]rune
}

func NewBalancedParentheses() *parenthesesChecker {
	return &parenthesesChecker{
		pairs: map[rune]rune{
			')': '(',
			']': '[',
			'}': '{',
		},
	}
}

// IsBalanced returns true if the input string has balanced parentheses
func (p *parenthesesChecker) IsBalanced(str string) bool {
	data := []rune{}
	for _, val := range str {
		switch val {
		case '(', '[', '{':
			data = append(data, val)
		case ')', ']', '}':
			opening, found := p.pairs[val]
			if !found || len(data) == 0 || opening != data[len(data)-1] {
				return false
			}
			data = data[:len(data)-1]
		default:
			continue
		}
	}
	return len(data) == 0
}
