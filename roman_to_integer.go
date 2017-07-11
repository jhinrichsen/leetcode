// Roman numerals are not a good candidate for programming examples, because
// there is no authority. Depending on the century, the location, and
// other factors, numbers were used one or the other way, even with regard to
// capitalization. MDCCCCLXXXXVIIII, MCMXCIX, and MIM are the same
// representation for 1999.

// On the other hand, maybe this vagueness makes it perfect for junior
// programmers - did you ever get a complete spec?

package leetcode

import "errors"

var invalidNumeralError = errors.New("Invalid roman numeral, only " +
	"I, V, X, L, C, D and M allowed")

// Input is guaranteed to be within the range 1..3999

// romanToIntStandard converts roman numerals (I, V, X, L, C, D, M) to a number
// in base 10 using standard notation, i.e. VIII for 8 and XVI for 16.
// In absence of an error in the function spec by leetcode, -1 is used to
// indicate invalid roman numerals
// Empty strings are illegal and case a -1 return
// If input can be interpreted in standard and subtraction way, subtraction is
// chosen. XIV returns 14 instead of 16.
func romanToInt(s string) int {
	if len(s) == 0 {
		return -1
	}
	// Standard representation
	std := func([]rune) int {
		base10 := 0
		for _, c := range s {
			// These hardcoded switches are not very maintainable,
			// but in the case of roman numerals changes are
			// unlikely
			switch c {
			case 'I':
				base10 += 1
			case 'V':
				base10 += 5
			case 'X':
				base10 += 10
			case 'L':
				base10 += 50
			case 'C':
				base10 += 100
			case 'D':
				base10 += 500
			case 'M':
				base10 += 1000
			default:
				return -1
			}
		}
		return base10
	}

	rs := []rune(s)
	return std(rs)
}
