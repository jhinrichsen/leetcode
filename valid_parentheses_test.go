package leetcode

import "testing"

var testTable = []struct {
	value    string
	expected bool
}{
	{"{}", true},
	{"[]", true},
	{"()", true},
	{"(){}[]", true},
	{"({})", true},
	{"{[()]}", true},

	{"(}", false},
	{"([]}", false},
	{"[(}]", false},
}

func TestAll(t *testing.T) {
	for _, prospect := range testTable {
		want := prospect.expected
		got := isValid(prospect.value)
		if want != got {
			t.Fatalf("Expected %v but got %v\n", want, got)
		}
	}
}
