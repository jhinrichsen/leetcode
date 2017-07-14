package leetcode

// In real world, i'd rather extract a func digit(n int), but leetcode expects
// the source code to be pasteable into a function body.

// isPalindrome returns true if a number is a number palindrome in base 10
// https://en.wikipedia.org/wiki/Palindromic_number
// Neither leetcode nor Wikipedia contain infos about negative numbers being
// palindromic, so negative numbers never return true
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// Number of digits
	digits := func(x int) int {
		if x < 0 {
			x = -x
		}
		n := 0
		for x != 0 {
			n++
			x /= 10
		}
		return n
	}

	// Power-of for int, avoid math.Pow(float64)
	pow := func(x, n int) int {
		p := 1
		for i := 0; i < n; i++ {
			p *= x
		}
		return p
	}

	// Digit d of base 10 integer n
	// 4 := digit(123456, 2) (0 based from right)
	// 123456 6 digits
	// 123456 / (10 ^ d) => 1234
	// 1234 % 10 => 4
	digit := func(n, d int) int {
		y := (n / pow(10, d)) % 10
		return y
	}

	left := digits(x) - 1
	right := 0
	for left >= right {
		l := digit(x, left)
		r := digit(x, right)
		if l != r {
			return false
		}
		left--
		right++
	}
	return true
}
