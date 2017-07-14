// Roman numerals are not a good candidate for programming examples, because
// there is no authority. Depending on the century, the location, and
// other factors, numbers were used one or the other way, even with regard to
// capitalization. MDCCCCLXXXXVIIII, MCMXCIX, and MIM are the same
// representation for 1999.

// On the other hand, maybe this vagueness makes it perfect for junior
// programmers - did you ever get a complete spec?

package leetcode

// Input is guaranteed to be within the range 1..3999

// romanToIntStandard converts roman numerals (I, V, X, L, C, D, M) to a number
// in base 10 using standard notation, i.e. VIII for 8 and XVI for 16.
// In absence of an error in the function spec by leetcode, -1 is used to
// indicate invalid roman numerals
// Empty strings are illegal and case a -1 return
// If input can be interpreted in standard and subtraction way, subtraction is
// chosen. XIV returns 14 instead of 16.
func romanToInt(s string) int {
	err := -1
	if len(s) == 0 {
		return err
	}
	val := func(roman rune) int {
		// These hardcoded switches are not very maintainable,
		// but changes to roman numerals are unlikely
		switch roman {
		case 'I':
			return 1
		case 'V':
			return 5
		case 'X':
			return 10
		case 'L':
			return 50
		case 'C':
			return 100
		case 'D':
			return 500
		case 'M':
			return 1000
		default:
			return err
		}
	}
	// subtraction rule: the next non-same character is higher
	isSub := func(rs []rune) bool {
		left := val(rs[0])
		for _, r := range rs {
			right := val(r)
			if left == right {
				continue
			}
			return left < right
		}
		return false
	}
	rs := []rune(s)
	sum := 0
	for i, r := range rs {
		n := val(r)
		if n == err {
			return err
		}
		if isSub(rs[i:]) {
			sum -= n
		} else {
			sum += n
		}
	}
	return sum
}
