package leetcode

import "testing"

func assert(t *testing.T, x int, want bool) {
	got := isPalindrome(x)
	if want != got {
		t.Fatalf("Expected %v but got %v\n", want, got)
	}
}

// Negative numbers are never palindromic
func TestNegativeNumber(t *testing.T) {
	assert(t, -131, false)
	assert(t, -1331, false)
}

// Zero is palindromic by definition
func TestZer0(t *testing.T) {
	assert(t, 0, true)
}

func TestNo(t *testing.T) {
	assert(t, 130, false)
	assert(t, 132, false)
	assert(t, 1311, false)
	assert(t, 1131, false)
}

func TestYes(t *testing.T) {
	assert(t, 131, true)
	assert(t, 1331, true)
}
