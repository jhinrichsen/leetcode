// Roman numerals are not a good candidate for programming examples, because
// there is no authority. Depending on the century, the location, and
// other factors, numbers were used one or the other way, even with regard to
// capitalization. MDCCCCLXXXXVIIII, MCMXCIX, and MIM are the same
// representation for 1999.

// On the other hand, maybe this vagueness makes it perfect for junior
// programmers - did you ever get a complete spec?

package leetcode

import "errors"

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

var invalidNumeralError = errors.New("Invalid roman numeral, only " +
	"I, V, X, L, C, D and M allowed")

// Input is guaranteed to be within the range 1..3999
func romanToInt(s string) int {
	// Spec does not say which way to encode XIV, so we go for standard
	n, _ := romanToIntStandard(s)
	return n
}

// romanToIntStandard converts roman numerals (I, V, X, L, C, D, M) to a number
// in base 10 using standard notation, i.e. VIII for 8 and XIV for 16.
func romanToIntStandard(roman string) (base10 int, err error) {
	for _, c := range roman {
		// These hardcoded switches are not very maintainable, but in
		// the case of roman numerals changes are unlikely
		switch c {
		case 'I':
			base10 += I
		case 'V':
			base10 += V
		case 'X':
			base10 += X
		case 'L':
			base10 += L
		case 'C':
			base10 += C
		case 'D':
			base10 += D
		case 'M':
			base10 += M
		default:
			return -1, invalidNumeralError
		}
	}
	return
}

// romanToIntSubstraction converts roman numerals (I, V, X, L, C, D, M) to a
// number in base 10 using subtraction notation, i.e. IIX for 8 and XIV for 14.
func romanToIntSubtraction(roman string) (base10 int, err error) {
	// Need positional access
	rs := []rune(roman)
	subMode := false

	return
}
