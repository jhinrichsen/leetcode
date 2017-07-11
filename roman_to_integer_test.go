package leetcode

import "testing"

type cvtFn func(string) (int, error)

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

var fStd cvtFn = func(roman string) (int, error) {
	return romanToIntStandard(roman)
}

var fSub cvtFn = func(roman string) (int, error) {
	return romanToIntSubtraction(roman)
}

func assert(t *testing.T, val string, want int, fn cvtFn) {
	got, err := fn(val)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if want != got {
		t.Fatalf("Want %v for %v but got %v", want, val, got)
	}
}

func TestStandard(t *testing.T) {
	for want, val := range standardTestTable {
		if val == "" {
			continue
		}
		assert(t, val, want, fStd)
	}
}

func TestSubtraction(t *testing.T) {
	for want, val := range subtractionTestTable {
		if val == "" {
			continue
		}
		assert(t, val, want, fSub)
	}
}

func TestEdgeCases(t *testing.T) {
	aStd := func(roman string, want int) {
		assert(t, roman, want, fStd)
	}
	aSub := func(roman string, want int) {
		assert(t, roman, want, fSub)
	}

	aStd("VIII", 8)
	aSub("IIX", 8)

	aSub("XL", 40)
	aSub("XL", 90)

	aStd("XLVIII", 48)
	aSub("IIL", 48)

	aStd("XLIX", 49)
	aSub("IL", 49)

	aStd("XCVIII", 98)
	aSub("IIC", 98)

	aStd("XCIX", 99)
	aSub("IC", 99)

	aStd("CMXC", 990)
	aSub("XM", 990)

	aStd("MDCCCCLXXXXVIIII", 1999)
	// subtraction only for certain numerals
	aSub("MCMXCIX", 1999)
	aSub("MIM", 1999)
}

func TestInvalidNumeral(t *testing.T) {
	bads := []string{"A", "X1", ""}
	fns := []func(string) (int, error){
		romanToIntStandard,
		romanToIntSubtraction,
	}
	for _, fn := range fns {
		for _, bad := range bads {
			n, err := fn(bad)
			if err != invalidNumeralError {
				t.Fatalf("Expected error but got %v, %v",
					n, err)
			}
		}
	}
}
