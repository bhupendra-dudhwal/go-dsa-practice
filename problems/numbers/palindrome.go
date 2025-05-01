package numbers

func IsPalindrome(x int) bool {
	// Negative numbers and numbers ending in 0 (but not 0 itself) are not palindromes
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	// For even length: x == reversedHalf
	// For odd length: x == reversedHalf/10 (middle digit doesn't matter)
	return x == reversedHalf || x == reversedHalf/10
}
