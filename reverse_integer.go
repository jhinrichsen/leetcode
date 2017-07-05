package leetcode

// reverse flips digits of an integer
// Return 0 for overflows
func reverse(x int) int {
	to := 0
	// Normalize
	neg := x < 0
	if neg {
		x = -x
	}

	for x > 0 {
		digit := x % 10
		x /= 10
		before := to
		to *= 10
		to += digit
		// Check for overflow
		if to < before {
			return 0
		}
	}
	if neg {
		to = -to
	}
	return to
}
