package main

import (
	"fmt"

	"github.com/bhupendra-dudhwal/go-dsa-practice/problems/str"
)

func main() {
	// fmt.Println(str.IsPalindrome("abccbaZabccba"))

	brackets := "))"
	fmt.Printf("\n brackets '%s' - %t\n", brackets, str.NewBalancedParentheses().IsBalanced(brackets))
}
