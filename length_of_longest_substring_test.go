package leetcode

import (
	"testing"
)

func assertLengthOfLongestSubstring(t *testing.T, arg string, want int) {
	got := LengthOfLongestSubstring(arg)
	if want != LengthOfLongestSubstring(arg) {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assertLengthOfLongestSubstring(t, "abcabcbb", 3)
	assertLengthOfLongestSubstring(t, "bbbbb", 1)
	assertLengthOfLongestSubstring(t, "pwwkew", 3)
	assertLengthOfLongestSubstring(t, "c", 1)
}
