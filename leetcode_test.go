package leetcode

import (
	"math/big"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	want := []int{0, 1}
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func TestListNode(t *testing.T) {
	got := listNode(big.NewInt(42))
	v := got.Val
	if v != 2 {
		t.Fatalf("want %v but got %v", 2, got)
	}
	v = got.Next.Val
	if v != 4 {
		t.Fatalf("want %v but got %v", 4, got)
	}
	if got.Next.Next != nil {
		t.Fatalf("want %v but got %v", nil, got)
	}
}

func TestVal(t *testing.T) {
	want := big.NewInt(2017)
	got := val(listNode(want))
	if want.Cmp(got) == 0 {
		t.Fatalf("want %+v but got %v", want, got)
	}
}

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
func TestAddTwoNumbers(t *testing.T) {
	l1 := listNode(big.NewInt(342))
	l2 := listNode(big.NewInt(465))

	want := big.NewInt(807)
	got := val(AddTwoNumbers(l1, l2))
	if want.Cmp(got) != 0 {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

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
