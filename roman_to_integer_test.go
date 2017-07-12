package leetcode

import (
	"testing"
)

type cvtFn func(string) int

var standardTestTable = []string{
	"",
	"I",
	"II",
	"III",
	"IIII",
	"V",
	"VI",
	"VII",
	"VIII",
	"VIIII",
	"X",
}

var subtractionTestTable = []string{
	"",
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"IIX",
	"IX",
	"X",
}

func assertRoman(t *testing.T, val string, want int) {
	got := romanToInt(val)
	if want != got {
		t.Fatalf("Want %v for %v but got %v", want, val, got)
	}
}

func TestStandard(t *testing.T) {
	for want, val := range standardTestTable {
		if val == "" {
			continue
		}
		assertRoman(t, val, want)
	}
}

func TestSubtraction(t *testing.T) {
	for want, val := range subtractionTestTable {
		if val == "" {
			continue
		}
		assertRoman(t, val, want)
	}
}

func TestEdgeCasesStandard(t *testing.T) {
	// Poor man's currying
	a := func(roman string, want int) {
		assertRoman(t, roman, want)
	}
	a("VIII", 8)
	a("XLVIII", 48)
	a("XCVIII", 98)
	a("XCIX", 99)
	a("CMXC", 990)
	a("MDCCCCLXXXXVIIII", 1999)
}

func TestEdgeCasesSubtraction(t *testing.T) {
	// Poor man's currying
	a := func(roman string, want int) {
		assertRoman(t, roman, want)
	}
	a("IIX", 8)
	a("XL", 40)
	a("XLIX", 49)
	a("IIL", 48)
	a("IL", 49)
	a("IIC", 98)
	a("IC", 99)
	a("XM", 990)

	// subtraction only for certain numerals
	a("MCMXCIX", 1999)
	a("MIM", 1999)
}

func TestInvalidNumeral(t *testing.T) {
	bads := []string{"A", "X1", "VIVY"}
	for _, bad := range bads {
		n := romanToInt(bad)
		if n != -1 {
			t.Fatalf("Expected error but got %v", n)
		}
	}
}
