package leetcode

import (
	"math/big"
	"testing"
)

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
