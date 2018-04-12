package leetcode

// Divide two integers without using multiplication, division and mod
// operator.
// If it is overflow, return MAX_INT.
func divide(dividend, divisor int) int {
	// brute force before SRT, Goldsmith, Newton-Raphson
	n := 0
	for dividend > 0 {
		dividend -= divisor
		n++
	}
	return n
}
